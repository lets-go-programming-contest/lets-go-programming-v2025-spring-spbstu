package storage

import (
	"database/sql"
	"fmt"

	"task-9/internal/config"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *sql.DB
}

func NewDatabase(config config.DBConfig) (*Database, error) {
	var db *sql.DB
	var err error

	switch config.Driver {
	case "postgres":
		connStr := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.Name, config.SSLMode,
		)
		db, err = sql.Open("postgres", connStr)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", config.Driver)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{conn: db}, nil
}

func (d *Database) GetConnection() *sql.DB {
	return d.conn
}

func (d *Database) Close() error {
	return d.conn.Close()
}

func (d *Database) RunMigrations() error {

	var queries = []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,

		`CREATE TABLE IF NOT EXISTS contacts (
				id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
				name VARCHAR(100) NOT NULL,
				phone VARCHAR(20) NOT NULL,
				email VARCHAR(100),
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				CONSTRAINT unique_name_phone UNIQUE (name, phone)
			)`,

		`CREATE OR REPLACE FUNCTION update_modified_column()
			RETURNS TRIGGER AS $$
			BEGIN
				NEW.updated_at = now();
				RETURN NEW;
			END;
			$$ language 'plpgsql'`,

		`DROP TRIGGER IF EXISTS update_contacts_timestamp ON contacts`,

		`CREATE TRIGGER update_contacts_timestamp
			BEFORE UPDATE ON contacts
			FOR EACH ROW EXECUTE PROCEDURE update_modified_column()`}

	for _, query := range queries {
		_, err := d.conn.Exec(query)
		if err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	return nil
}
