package db_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"example_mock/internal/db"
)

func getGetNamesQuery() string {
	return "SELECT name FROM users"
}

func getRowsAndExpected(columns []string) (*sqlmock.Rows, []string) {
	expectedNames := []string{"Ksusha", "Matvei"}
	rows := sqlmock.NewRows(columns)
	for _, name := range expectedNames {
		rows.AddRow(name)
	}
	return rows, expectedNames
}

func TestGetNamesSuccess(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		_ = mockDB.Close()
	}()

	query := getGetNamesQuery()
	rows, expectedNames := getRowsAndExpected([]string{"name"})
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, expectedNames, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		_ = mockDB.Close()
	}()

	query := getGetNamesQuery()
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))

	service := db.New(mockDB)
	names, err := service.GetNames()
	require.Error(t, err, "expected error")
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func getSelectUniqueValuesColTblQuery() (string, string, string) {
	columnName := "name"
	tableName := "names"
	query := "SELECT DISTINCT " + columnName + " FROM " + tableName
	return columnName, tableName, query
}

func TestSelectUniqueValuesSuccess(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		_ = mockDB.Close()
	}()

	columnName, tableName, query := getSelectUniqueValuesColTblQuery()
	rows, expectedNames := getRowsAndExpected([]string{columnName})
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	service := db.New(mockDB)
	values, err := service.SelectUniqueValues(columnName, tableName)
	require.NoError(t, err)
	require.Equal(t, expectedNames, values)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSelectUniqueValuesError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		_ = mockDB.Close()
	}()

	columnName, tableName, query := getSelectUniqueValuesColTblQuery()
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))

	service := db.New(mockDB)
	values, err := service.SelectUniqueValues(columnName, tableName)
	require.Error(t, err, "query error")
	require.Nil(t, values)
	require.NoError(t, mock.ExpectationsWereMet())
}
