package main

import (
	"database/sql"
	"errors"
	"fmt"

	dbPack "example_mock/internal/db"
)

var (
	errSQLOpenFailed  = errors.New("failed to open SQL")
	errGetNamesFailed = errors.New("failed getting names")
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(errors.Join(errSQLOpenFailed, err))
	}
	defer db.Close()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()
	if err != nil {
		panic(errors.Join(errGetNamesFailed, err))
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
