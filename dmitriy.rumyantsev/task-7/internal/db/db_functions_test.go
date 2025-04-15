package db_test

import (
	"errors"
	"testing"

	"github.com/dmitriy.rumyantsev/task-7/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func mockRows(column string, values []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{column})
	for _, val := range values {
		rows.AddRow(val)
	}
	return rows
}

func TestGetNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		expected := []string{"Alice", "Bob"}
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockRows("name", expected))

		result, err := service.GetNames()
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("query error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnError(errors.New("query failed"))

		result, err := service.GetNames()
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("scan error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		result, err := service.GetNames()
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("rows error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		rows := sqlmock.NewRows([]string{"name"}).AddRow("Ivan")
		rows.RowError(0, errors.New("rows error"))
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		result, err := service.GetNames()
		require.Error(t, err)
		require.Nil(t, result)
	})
}

func TestSelectUniqueValues(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		expected := []string{"admin", "user"}
		mock.ExpectQuery("SELECT DISTINCT role FROM employees").
			WillReturnRows(mockRows("role", expected))

		result, err := service.SelectUniqueValues("role", "employees")
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("query error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		mock.ExpectQuery("SELECT DISTINCT role FROM employees").
			WillReturnError(errors.New("query error"))

		result, err := service.SelectUniqueValues("role", "employees")
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("scan error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		rows := sqlmock.NewRows([]string{"role"}).AddRow(nil)
		mock.ExpectQuery("SELECT DISTINCT role FROM employees").
			WillReturnRows(rows)

		result, err := service.SelectUniqueValues("role", "employees")
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("rows error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		service := db.Service{DB: mockDB}
		rows := sqlmock.NewRows([]string{"role"}).AddRow("admin")
		rows.RowError(0, errors.New("rows error"))
		mock.ExpectQuery("SELECT DISTINCT role FROM employees").
			WillReturnRows(rows)

		result, err := service.SelectUniqueValues("role", "employees")
		require.Error(t, err)
		require.Nil(t, result)
	})
}
