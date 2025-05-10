package db

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
        errOpenFailed = errors.New("failed to open")
        errPingFailed = errors.New("ping failed")
)

type DataBase struct {
        Postgres *sql.DB
}

func New() *DataBase {
        return &DataBase{}
}

func (database *DataBase) Open(port string, pswd string) error {
        connStr := fmt.Sprintf("host=localhost port=%s user=postgres password=%s dbname=contacts sslmode=disable",
                               port, pswd)

        db, err := sql.Open("postgres", connStr)
        if err != nil {
                return errors.Join(errOpenFailed, err)
        }
        database.Postgres = db

        err = database.Postgres.Ping()
        if err != nil {
                return errors.Join(errPingFailed, err)
        }
        return nil
}

func (database *DataBase) Close() {
        database.Postgres.Close()
}
