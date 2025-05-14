package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	dbPack "example_mock/internal/db"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	db.Close()
}
