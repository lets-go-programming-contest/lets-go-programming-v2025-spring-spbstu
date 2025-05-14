package db

import (
	"database/sql"
	"fmt"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type Service struct {
	DB Database
}

func New(db Database) Service {
	return Service{DB: db}
}

func (s Service) GetNames() ([]string, error) {
	query := "SELECT name FROM users"

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetNames query: %w", err)
	}
	defer rows.Close()

	var names []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("GetNames scan: %w", err)
		}

		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetNames rows error: %w", err)
	}

	return names, nil
}

func (s Service) SelectUniqueValues(columnName string, tableName string) ([]string, error) {
	query := "SELECT DISTINCT " + columnName + " FROM " + tableName

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("SelectUniqueValues query: %w", err)
	}
	defer rows.Close()

	var values []string

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, fmt.Errorf("SelectUniqueValues scan: %w", err)
		}

		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SelectUniqueValues rows error: %w", err)
	}

	return values, nil
}
