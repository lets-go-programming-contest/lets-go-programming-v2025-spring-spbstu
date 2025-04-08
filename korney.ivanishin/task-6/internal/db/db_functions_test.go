package db_test

import (
	"errors"
	"example_mock/internal/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestGetNames struct {
        names       []string
        errExpected error
}

var testTableGetNames = []rowTestGetNames{
        {
                names:       []string{"name1, name2"},
                errExpected: nil,
        },
        {
                names:       nil,
                errExpected: errors.New("ExpectedError"),
        },
}

func TestGetNames(t *testing.T) {
        mockDB, mock, err := sqlmock.New()
        if err != nil {
                t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
        }

        dbService := db.Service{DB: mockDB}

        for i, row := range testTableGetNames {
                mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDBRows(row.names)).WillReturnError(row.errExpected)
                names, err := dbService.GetNames()

                if row.errExpected != nil {
                        require.ErrorIs(t, err, joinedDBError(row.errExpected), "row: %d, expected error: %w, actual error: %w", i, joinedDBError(row.errExpected).Error(), err.Error())
                        require.Nil(t, names, "row: %d, names must be nil", i)
                        continue
                }

                require.NoError(t, err, "row: %d, error must be nil", i)
                require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
        }
}

type dataTestSelectUniqueValues struct {
        colName     string
        tblName     string
        values      []string
        errExpected error
}

var testTableSelectUniqueValues = []dataTestSelectUniqueValues{
        {
                colName:     "ExpectedColName",
                tblName:     "ExpectedTblName",
                values:      []string{"uniqueVal1, uniqueVal2"},
                errExpected: nil,
        },
        {
                colName:     "ExpectedColName",
                tblName:     "ExpectedTblName",
                values:      nil,
                errExpected: errors.New("ExpectedValueError"),
        },
        {
                colName:     "TrashColName",
                tblName:     "ExpectedTblName",
                values:      nil,
                errExpected: errors.New("ExpectedColNameError"),
        },
        {
                colName:     "ExpectedColName",
                tblName:     "TrashTblName",
                values:      nil,
                errExpected: errors.New("ExpectedTblNameError"),
        },
}

func TestSelectUniqueValues(t *testing.T) {
        mockDB, mock, err := sqlmock.New()
        if err != nil {
                t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
        }

        dbService := db.Service{DB: mockDB}
        for i, data := range testTableSelectUniqueValues {
                mock.ExpectQuery("SELECT DISTINCT " + data.colName + " FROM " + data.tblName).WillReturnRows(mockDBRows(data.values)).WillReturnError(data.errExpected)
                values, err := dbService.SelectUniqueValues(data.colName, data.tblName)

                if data.errExpected != nil {
                        require.ErrorIs(t, err, joinedDBError(data.errExpected), "row: %d, expected error: %w, actual error: %w", i, joinedDBError(data.errExpected).Error(), err.Error())
                        require.Nil(t, values, "row: %d, values must be nil", i)
                        continue
                }

                require.NoError(t, err, "row: %d, error must be nil", i)
                require.Equal(t, data.values, values, "row: %d, expected unique values: %s, actual unique values: %s", i, data.values, values)
        }
}

func mockDBRows(names []string) *sqlmock.Rows {
        rows := sqlmock.NewRows([]string{"name"})
        for _, name := range names {
                rows = rows.AddRow(name)
        }

        return rows
}

func joinedDBError(errDB error) error {
        return errors.Join(db.ErrGetRowsFailed, errDB)
}
