package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	service := New(db)

	expectedNames := []string{"AAAA", "bbb"}

	rows := sqlmock.NewRows([]string{"name"})
	
	for _, name := range expectedNames {
		rows.AddRow(name)
	}

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, expectedNames, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	service := New(db)

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrConnDone)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	service := New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSelectUniqueValues(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	service := New(db)

	rows := sqlmock.NewRows([]string{"status"}).
		AddRow("pending").
		AddRow("completed")

	mock.ExpectQuery("SELECT DISTINCT status FROM orders").WillReturnRows(rows)

	values, err := service.SelectUniqueValues("status", "orders")

	require.NoError(t, err)
	require.Equal(t, []string{"pending", "completed"}, values)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSelectUniqueValuesQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	service := New(db)

	mock.ExpectQuery("SELECT DISTINCT status FROM orders").WillReturnError(sql.ErrConnDone)

	values, err := service.SelectUniqueValues("status", "orders")

	require.Error(t, err)
	require.Nil(t, values)
	require.NoError(t, mock.ExpectationsWereMet())
}