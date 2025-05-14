package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/alexander-romanov-edu/phonebook/pkg/models"
)

type PhonebookRepository struct {
	DB *sql.DB
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func NewPhonebookRepository() *PhonebookRepository {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "phonebook")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &PhonebookRepository{DB: db}
}

func (r *PhonebookRepository) DeleteMe() {
	r.DB.Close()
}

func (r *PhonebookRepository) Initialize() error {
	// First check if the table exists
	var tableExists bool
	err := r.DB.QueryRow(`
        SELECT EXISTS (
            SELECT FROM information_schema.tables 
            WHERE table_name = 'contacts'
        )
    `).Scan(&tableExists)
	if err != nil {
		return fmt.Errorf("failed to check if table exists: %w", err)
	}

	if tableExists {
		// Check if the id column exists
		var columnExists bool
		err = r.DB.QueryRow(`
            SELECT EXISTS (
                SELECT FROM information_schema.columns 
                WHERE table_name = 'contacts' AND column_name = 'id'
            )
        `).Scan(&columnExists)
		if err != nil {
			return fmt.Errorf("failed to check if column exists: %w", err)
		}

		if !columnExists {
			// Migrate the existing table
			_, err = r.DB.Exec(`
                ALTER TABLE contacts 
                ADD COLUMN id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                ADD CONSTRAINT unique_name_phone UNIQUE (name, phone)
            `)
			if err != nil {
				return fmt.Errorf("failed to migrate table: %w", err)
			}
			log.Println("Successfully migrated contacts table")
		}
	} else {
		// Create new table with the correct schema
		_, err = r.DB.Exec(`
            CREATE TABLE contacts (
                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                name VARCHAR(100) NOT NULL,
                phone VARCHAR(20) NOT NULL,
                email VARCHAR(100),
                CONSTRAINT unique_name_phone UNIQUE (name, phone)
            )
        `)
		if err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
		log.Println("Successfully created contacts table")
	}

	return nil
}

func (r *PhonebookRepository) DeleteContact(id string) error {
	query := `DELETE FROM contacts WHERE id = $1`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("contact not found")
	}

	return nil
}

func (r *PhonebookRepository) AddContact(contact models.Contact) error {
	if contact.ID == "" {
		// Let the database generate the UUID
		query := `
        INSERT INTO contacts (name, phone, email)
        VALUES ($1, $2, $3)
        ON CONFLICT (name, phone) DO UPDATE 
        SET email = EXCLUDED.email
        RETURNING id`

		err := r.DB.QueryRow(query, contact.Name, contact.Phone, contact.Email).Scan(&contact.ID)
		return err
	} else {
		query := `
        INSERT INTO contacts (id, name, phone, email)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (name, phone) DO UPDATE 
        SET email = EXCLUDED.email`
		_, err := r.DB.Exec(query, contact.ID, contact.Name, contact.Phone, contact.Email)
		return err
	}
}

func (r *PhonebookRepository) GetContact(id string) (models.Contact, error) {
	var contact models.Contact
	query := `SELECT id, name, phone, email FROM contacts WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email)
	return contact, err
}

func (r *PhonebookRepository) UpdateContact(contact models.Contact) error {
	query := `UPDATE contacts SET name = $1, phone = $2, email = $3 WHERE id = $4`
	result, err := r.DB.Exec(query, contact.Name, contact.Phone, contact.Email, contact.ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("contact not found")
	}

	return nil
}

func (r *PhonebookRepository) GetAllContacts() ([]models.Contact, error) {
	query := `SELECT id, name, phone, email FROM contacts ORDER BY name`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
