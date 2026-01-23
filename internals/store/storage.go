package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //driver
	"github.com/pradeept/site-monitor-cli/internals/logger"
)

type SiteStatus struct {
	Id         int
	SiteUrl    string
	StatusCode int
	StatusText string
}

type Store struct {
	db  *sql.DB
	log *log.Logger
}

func NewStore(dbString string) (*Store, error) {

	db, err := sql.Open("sqlite3", dbString)

	if err != nil {
		return nil, err
	}

	// tell sqlite explicitely to handle foreign keys (by default it ignores)
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	// site table stores the site information
	siteTable := `
 		CREATE TABLE IF NOT EXISTS site (
  		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  		site_name TEXT ,
		site_url TEXT NOT NULL UNIQUE
 	);`

	if _, err := db.Exec(siteTable); err != nil {
		return nil, err
	}

	// a request is a http call made to the site
	// every call is realted to a site
	requestTable := `
		CREATE TABLE IF NOT EXISTS request(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		site_id INTEGER NOT NULL,
		status_code INTEGER NOT NULL,
		status_text TEXT,
		FOREIGN KEY (site_id) REFERENCES site(id)
	);`

	if _, err := db.Exec(requestTable); err != nil {
		return nil, err
	}

	return &Store{
		db,
		logger.Logger(),
	}, nil
}

// Insert method

// Delete method
