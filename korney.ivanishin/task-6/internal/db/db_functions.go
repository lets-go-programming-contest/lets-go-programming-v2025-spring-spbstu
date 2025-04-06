package db

import (
	"database/sql"
	"errors"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type Service struct {
	DB Database
}

var (
	errGetRowsFailed  = errors.New("failed getting rows")
	errScanRowsFailed = errors.New("failed scanning rows")
	errIterRowsFailed = errors.New("failed iterating rows")
)

func New(db Database) Service {
	return Service{DB: db}
}

func (service Service) GetNames() ([]string, error) {
	query := "SELECT name FROM users"

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, errors.Join(errGetRowsFailed, err)
	}
	defer rows.Close()

	var names []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, errors.Join(errScanRowsFailed, err)
		}

		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Join(errIterRowsFailed, err)
	}

	return names, nil
}

func (service Service) SelectUniqueValues(columnName string, tableName string) ([]string, error) {
	query := "SELECT DISTINCT " + columnName + " FROM " + tableName

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, errors.Join(errGetRowsFailed, err)
	}

	defer rows.Close()

	var values []string

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, errors.Join(errScanRowsFailed, err)
		}

		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Join(errIterRowsFailed, err)
	}

	return values, nil
}
