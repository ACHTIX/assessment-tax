package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgresSQL driver
)

func ConnectDB() error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("connect to database error: %w", err)
	}
	defer db.Close()

	log.Println("Connection to database established successfully")

	createTb := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT);`
	if _, err = db.Exec(createTb); err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	fmt.Println("Create table success")
	return nil
}
