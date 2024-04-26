package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import your SQL driver
	"io/ioutil"
	"log"
	"os"
)

func executeSQLFile(db *sql.DB, filepath string) error {

	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	sqlCommands := string(fileContent)
	_, err = db.Exec(sqlCommands)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	DATABASE_URL := os.Getenv("DATABASE_URL")

	if DATABASE_URL == "" {
		DATABASE_URL = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable" // Default port if not specified
		log.Println("DATABASE_URL", DATABASE_URL)
	}
	db, err := sql.Open("postgres", DATABASE_URL)

	if err != nil {
		log.Fatal(err)
	}
	if err := executeSQLFile(db, "./allowance.sql"); err != nil {
		log.Fatalf("Failed to execute master_deduct.sql: %v", err)
	}

	fmt.Println("Database initialized successfully!")
}
