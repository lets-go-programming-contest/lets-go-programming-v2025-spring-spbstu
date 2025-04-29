package db_test

import (
	"errors"
	"regexp"
	"testing"

	"example_mock/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetNames_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	query := "SELECT name FROM users"
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob")
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice", "Bob"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	query := "SELECT name FROM users"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("query error"))

	names, err := service.GetNames()
	require.Error(t, err, "expected query error")
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanConversion(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	query := "SELECT name FROM users"
	// Simulate a row where the 'name' field is an integer (123)
	rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	names, err := service.GetNames()
	require.NoError(t, err, "expected no error during scan")
	require.Equal(t, []string{"123"}, names, "expected names to contain '123' as a string")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSelectUniqueValues_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	column := "city"
	table := "locations"
	query := "SELECT DISTINCT " + column + " FROM " + table
	rows := sqlmock.NewRows([]string{column}).AddRow("NY").AddRow("LA")
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	values, err := service.SelectUniqueValues(column, table)
	require.NoError(t, err)
	require.Equal(t, []string{"NY", "LA"}, values)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSelectUniqueValues_Error(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	column := "city"
	table := "locations"
	query := "SELECT DISTINCT " + column + " FROM " + table
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("query error"))

	values, err := service.SelectUniqueValues(column, table)
	require.Error(t, err, "expected query error")
	require.Nil(t, values)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSelectUniqueValues_ScanConversion(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	column := "city"
	table := "locations"
	query := "SELECT DISTINCT " + column + " FROM " + table
	// Simulate a row where the 'column' field is an integer (456)
	rows := sqlmock.NewRows([]string{column}).AddRow(456)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	values, err := service.SelectUniqueValues(column, table)
	require.NoError(t, err, "expected no error during scan")
	require.Equal(t, []string{"456"}, values, "expected names to contain '456' as a string")
	require.NoError(t, mock.ExpectationsWereMet())
}
