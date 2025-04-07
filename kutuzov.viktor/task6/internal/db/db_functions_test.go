package db_test

import (
	"errors"
	"example_mock/internal/db"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	errExpected error
}

var testTable = []rowTestDb{
	{
		names: []string{"Ivan, Gena228"},
	},
	{
		names:       nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	defer func() {
		_ = mockDB.Close()
	}()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected jsondata", err)
	}

	Service := db.Service{DB: mockDB}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).WillReturnError(row.errExpected)
		names, err := Service.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func TestSelectUniqueValues_ScanConversion(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer func() {
		if err := mockDB.Close(); err != nil {
			log.Printf("failed to close mockDB: %v", err)
			// или возврат ошибки, если это уместно
		}
	}()

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
