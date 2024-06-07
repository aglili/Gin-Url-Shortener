package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// initializeDb initializes the SQLite database
func initializeDb() {
	var err error
	// Open or create the SQLite database file
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create 'urls' table if it doesn't exist
	createTable := `CREATE TABLE IF NOT EXISTS urls(
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"original_url" TEXT NOT NULL,
		"short_url" TEXT NOT NULL
	)`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
