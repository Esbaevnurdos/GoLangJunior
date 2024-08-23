package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDatabase() {
    var err error

	DB, err = sql.Open("sqlite", "/movie_reviews.db")

	if err != nil {
		log.Fatal(err)
	}


	initTables()
	
}


func initTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
	);`

	createMovieTable := `CREATE TABLE IF NOT EXISTS movies (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT,
	release_year INTEGER
	);`

	createReviewTable := `CREATE TABLE IF NOT EXISTS reviews (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	movie_id INTEGER,
	user_id INTEGER,
	rating INTEGER CHECK(rating >= 1 AND rating <= 5),
	review_text TEXT,
	FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	DB.Exec(createUserTable)
	DB.Exec(createMovieTable)
	DB.Exec(createReviewTable)

}