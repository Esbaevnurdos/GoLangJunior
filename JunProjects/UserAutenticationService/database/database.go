package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect! Open")
	}

	DB = db
    
	err = createTables()

	if err != nil {
		panic("Database could not connect! CreateTables")
	}

	fmt.Println("Tables created succesfully")
	

}

func createTables() error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create users table! Exec")
	}

	return nil
}