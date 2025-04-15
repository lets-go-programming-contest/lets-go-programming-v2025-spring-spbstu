package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*sql.DB, *sqlmock.Sqlmock, Service) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	service := New(db)
	return db, &mock, service
}

func TestGetNamesBasic(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("name1").AddRow("name2")

	(*mock).ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	assert.NoError(t, err)
	assert.Equal(t, []string{"name1", "name2"}, names)
}

func TestGetNamesError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	(*mock).ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrConnDone)

	names, err := service.GetNames()

	assert.Error(t, err)
	assert.Nil(t, names)
}

func TestGetNamesScanError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	(*mock).ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	assert.Error(t, err)
	assert.Nil(t, names)
}

func TestSelectUniqueValuesBasic(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"category"}).AddRow("cat1").AddRow("cat2").AddRow("cat3")

	(*mock).ExpectQuery("SELECT DISTINCT category FROM products").WillReturnRows(rows)

	values, err := service.SelectUniqueValues("category", "products")

	assert.NoError(t, err)
	assert.Equal(t, []string{"cat1", "cat2", "cat3"}, values)
}

func TestSelectUniqueValuesQueryError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	(*mock).ExpectQuery("SELECT DISTINCT category FROM products").WillReturnError(sql.ErrConnDone)

	values, err := service.SelectUniqueValues("category", "products")

	assert.Error(t, err)
	assert.Nil(t, values)
}

func TestSelectUniqueValuesParsingError(t *testing.T) {
	db, mock, service := setupTestDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"category"}).AddRow(nil)
	(*mock).ExpectQuery("SELECT DISTINCT category FROM products").WillReturnRows(rows)

	values, err := service.SelectUniqueValues("category", "products")

	assert.Error(t, err)
	assert.Nil(t, values)
}
