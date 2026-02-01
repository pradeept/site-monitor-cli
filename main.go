/*
This file is used as a playground. Will be trashed later.
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" //driver
)

// STORE: SQlit
// Requests: net/http
// CLI : flag
// STD out: log / fmt

// add website (website string, time time.Seconds)
// remove website (website string)
// list website stats () []websites

// LIVE VIEW: live refresh of websites request status
// S. NO. | Website Name | Request history | total requests | fail rate | success rate
func main() {
	fmt.Println("A")
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(pwd, "test.db")
	fmt.Println("DB path:", dbPath)

	db, err := sql.Open("sqlite3", dbPath+"?_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'users' created successfully")
}
