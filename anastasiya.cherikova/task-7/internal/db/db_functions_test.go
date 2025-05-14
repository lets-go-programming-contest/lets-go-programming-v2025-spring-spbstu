package db_test

import (
	"errors"
	"regexp"
	"testing"

	"example_mock/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

// TestGetNames_ValidResults verifies that the GetNames method correctly retrieves
// and returns user names from the database when valid results are present.
func TestGetNames_ValidResults(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	svc := db.New(mockDB)
	expectedQuery := regexp.QuoteMeta("SELECT name FROM users")
	rows := sqlmock.NewRows([]string{"name"}).AddRow("John").AddRow("Doe")
	mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

	names, err := svc.GetNames()
	require.NoError(t, err)
	require.ElementsMatch(t, []string{"John", "Doe"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

// TestGetNames_HandlesQueryError verifies that the GetNames method properly handles
// and propagates errors that occur during database query execution.
func TestGetNames_HandlesQueryError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	svc := db.New(mockDB)
	expectedQuery := regexp.QuoteMeta("SELECT name FROM users")
	mock.ExpectQuery(expectedQuery).WillReturnError(errors.New("connection failed"))

	names, err := svc.GetNames()
	require.ErrorContains(t, err, "connection failed")
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

// TestGetNames_TypeConversion verifies that the GetNames method correctly converts
// non-string database values (e.g., integers) into strings during row scanning.
func TestGetNames_TypeConversion(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	svc := db.New(mockDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow(456) // Simulate integer value in a string column
	mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM users")).WillReturnRows(rows)

	names, err := svc.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"456"}, names)
}

// TestSelectUniqueValues_ScanConversion verifies that the SelectUniqueValues method
// correctly converts non-string database values into strings when scanning rows.
func TestSelectUniqueValues_ScanConversion(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	column := "city"
	table := "locations"
	query := "SELECT DISTINCT " + column + " FROM " + table

	// Simulate a row with an integer value in a string column
	rows := sqlmock.NewRows([]string{column}).AddRow(456)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	values, err := service.SelectUniqueValues(column, table)
	require.NoError(t, err, "expected no error during scan")
	require.Equal(t, []string{"456"}, values, "expected string conversion of integer value")
	require.NoError(t, mock.ExpectationsWereMet())
}

// TestSelectUniqueValues_EmptyResult verifies that the SelectUniqueValues method
// correctly handles empty result sets by returning an empty slice without errors.
func TestSelectUniqueValues_EmptyResult(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	svc := db.New(mockDB)
	query := regexp.QuoteMeta("SELECT DISTINCT city FROM locations")
	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"city"}))

	values, err := svc.SelectUniqueValues("city", "locations")
	require.NoError(t, err)
	require.Empty(t, values)
}

// TestSelectUniqueValues_InvalidTable verifies that the SelectUniqueValues method
// correctly propagates database errors (e.g., querying a non-existent table).
func TestSelectUniqueValues_InvalidTable(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	svc := db.New(mockDB)
	invalidTable := "non_existent_table"
	query := regexp.QuoteMeta("SELECT DISTINCT id FROM " + invalidTable)
	mock.ExpectQuery(query).WillReturnError(errors.New("table not found"))

	values, err := svc.SelectUniqueValues("id", invalidTable)
	require.Error(t, err)
	require.Nil(t, values)
}
