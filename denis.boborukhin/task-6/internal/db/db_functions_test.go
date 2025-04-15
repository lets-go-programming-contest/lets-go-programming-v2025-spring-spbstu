package db

import (
	"database/sql"
	"errors"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, Service) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	service := New(db)
	return db, mock, service
}

func TestGetNames(t *testing.T) {
	tests := []struct {
		name          string
		setupMock     func(sqlmock.Sqlmock)
		expectedNames []string
		expectedErr   error
	}{
		{
			name: "successful case",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("name1").
					AddRow("name2")
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(rows)
			},
			expectedNames: []string{"name1", "name2"},
			expectedErr:   nil,
		},
		{
			name: "connection error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnError(sql.ErrConnDone)
			},
			expectedNames: nil,
			expectedErr:   sql.ErrConnDone,
		},
		{
			name: "scan error",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(rows)
			},
			expectedNames: nil,
			expectedErr:   errors.New("sql: Scan error on column index 0, name \"name\": converting NULL to string is unsupported"),
		},
		{
			name: "empty result",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"})
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(rows)
			},
			expectedNames: nil,
			expectedErr:   nil,
		},
		{
			name: "rows error after iteration",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("name1").
					AddRow("name2")
				rows.RowError(1, errors.New("iteration error"))
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(rows)
			},
			expectedNames: nil,
			expectedErr:   errors.New("iteration error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, service := setupTestDB(t)
			defer func() {
				db.Close()
				require.NoError(t, mock.ExpectationsWereMet())
			}()

			tt.setupMock(mock)

			names, err := service.GetNames()

			if tt.expectedErr != nil {
				assert.Error(t, err)
				if tt.expectedErr == sql.ErrConnDone {
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.True(t, strings.Contains(err.Error(), tt.expectedErr.Error()))
				}
				assert.Nil(t, names)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedNames, names)
		})
	}
}

func TestSelectUniqueValues(t *testing.T) {
	tests := []struct {
		name           string
		column         string
		table          string
		setupMock      func(sqlmock.Sqlmock)
		expectedValues []string
		expectedErr    error
	}{
		{
			name:   "successful case",
			column: "category",
			table:  "products",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"category"}).
					AddRow("cat1").
					AddRow("cat2").
					AddRow("cat3")
				mock.ExpectQuery("SELECT DISTINCT category FROM products").
					WillReturnRows(rows)
			},
			expectedValues: []string{"cat1", "cat2", "cat3"},
			expectedErr:    nil,
		},
		{
			name:   "connection error",
			column: "category",
			table:  "products",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT category FROM products").
					WillReturnError(sql.ErrConnDone)
			},
			expectedValues: nil,
			expectedErr:    sql.ErrConnDone,
		},
		{
			name:   "scan error",
			column: "category",
			table:  "products",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"category"}).AddRow(nil)
				mock.ExpectQuery("SELECT DISTINCT category FROM products").
					WillReturnRows(rows)
			},
			expectedValues: nil,
			expectedErr:    errors.New("sql: Scan error on column index 0, name \"category\": converting NULL to string is unsupported"),
		},
		{
			name:   "empty result",
			column: "category",
			table:  "products",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"category"})
				mock.ExpectQuery("SELECT DISTINCT category FROM products").
					WillReturnRows(rows)
			},
			expectedValues: nil,
			expectedErr:    nil,
		},
		{
			name:   "invalid column name",
			column: "invalid; DROP TABLE products; --",
			table:  "products",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT invalid; DROP TABLE products; -- FROM products").
					WillReturnError(sql.ErrNoRows)
			},
			expectedValues: nil,
			expectedErr:    sql.ErrNoRows,
		},
		{
			name:   "rows error after iteration",
			column: "category",
			table:  "products",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"category"}).
					AddRow("electronics").
					AddRow("books")
				rows.RowError(1, errors.New("iteration error"))
				mock.ExpectQuery("SELECT DISTINCT category FROM products").
					WillReturnRows(rows)
			},
			expectedValues: nil,
			expectedErr:    errors.New("iteration error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, service := setupTestDB(t)
			defer func() {
				db.Close()
				require.NoError(t, mock.ExpectationsWereMet())
			}()

			tt.setupMock(mock)

			values, err := service.SelectUniqueValues(tt.column, tt.table)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				if tt.expectedErr == sql.ErrConnDone || tt.expectedErr == sql.ErrNoRows {
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.True(t, strings.Contains(err.Error(), tt.expectedErr.Error()))
				}
				assert.Nil(t, values)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedValues, values)
		})
	}
}
