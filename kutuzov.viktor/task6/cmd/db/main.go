package main

import (
	"database/sql"
	dbPack "example_mock/internal/db"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close db: %v", err)
		}
	}()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()

	if err != nil {
		log.Printf("failed to get names: %v", err)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
