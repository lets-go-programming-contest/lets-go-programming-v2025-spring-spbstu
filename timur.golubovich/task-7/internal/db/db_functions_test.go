package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// setupTestDB initializes a new sqlmock database and returns the mock and service
func setupTestDB(t *testing.T) (*sql.DB, *sqlmock.Sqlmock, Service) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	service := New(db)
	return db, &mock, service
}

func TestGetNames(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")

	(*mock).ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	assert.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob"}, names)
	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestGetNamesQueryError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	(*mock).ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrConnDone)

	names, err := service.GetNames()

	assert.Error(t, err)
	assert.Nil(t, names)
	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestGetNamesScanError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil) // Causes scan error
	(*mock).ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	assert.Error(t, err)
	assert.Nil(t, names)
	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestSelectUniqueValues(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"category"}).
		AddRow("Electronics").
		AddRow("Furniture")

	(*mock).ExpectQuery("SELECT DISTINCT category FROM products").WillReturnRows(rows)

	values, err := service.SelectUniqueValues("category", "products")

	assert.NoError(t, err)
	assert.Equal(t, []string{"Electronics", "Furniture"}, values)
	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestSelectUniqueValuesQueryError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	(*mock).ExpectQuery("SELECT DISTINCT category FROM products").WillReturnError(sql.ErrConnDone)

	values, err := service.SelectUniqueValues("category", "products")

	assert.Error(t, err)
	assert.Nil(t, values)
	assert.NoError(t, (*mock).ExpectationsWereMet())
}

func TestSelectUniqueValuesScanError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"category"}).AddRow(nil) // Causes scan error
	(*mock).ExpectQuery("SELECT DISTINCT category FROM products").WillReturnRows(rows)

	values, err := service.SelectUniqueValues("category", "products")

	assert.Error(t, err)
	assert.Nil(t, values)
	assert.NoError(t, (*mock).ExpectationsWereMet())
}
