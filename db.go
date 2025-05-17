package main

import (
	"database/sql"
	"log"
)

import _ "github.com/mattn/go-sqlite3"

var DB *sql.DB

func initDB() {
	// Try to open the database
	var err error
	DB, err = sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the table
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT NOT NULL,
        year INTEGER
    )`)
	if err != nil {
		log.Fatal(err)
	}
}
