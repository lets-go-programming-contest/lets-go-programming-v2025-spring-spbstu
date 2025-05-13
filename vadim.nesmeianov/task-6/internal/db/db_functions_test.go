package db_test

import (
	"errors"
	"example_mock/internal/db"
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
		names:       []string{"Ivan", "Gena228"},
		errExpected: nil,
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

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	defer func() {
		_ = mockDB.Close()
	}()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected jsondata", err)
	}

	service := db.New(mockDB)

	column := "name"
	table := "users"
	query := "SELECT DISTINCT " + column + " FROM " + table

	for _, row := range testTable {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(mockDbRows(row.names)).WillReturnError(row.errExpected)
		names, err := service.SelectUniqueValues(column, table)

		if row.errExpected != nil {
			require.Error(t, err, "expected query error")
			require.Nil(t, names)
			require.NoError(t, mock.ExpectationsWereMet())

			continue
		}

		require.NoError(t, err)
		require.Equal(t, row.names, names)
		require.NoError(t, mock.ExpectationsWereMet())
	}
}
