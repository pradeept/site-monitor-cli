package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //driver
	"github.com/pradeept/site-monitor-cli/internals/logger"
)

type Site struct {
	Id       int
	SiteName string
	SiteUrl  string
}

type SiteStatus struct {
	Id         int
	Site       *Site
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
func (s *Store) InsertSite(site *Site) error {
	insertStatement := `INSERT INTO site(site_name, site_url) VALUES(?,?);`
	if _, err := s.db.Exec(insertStatement, site.SiteName, site.SiteUrl); err != nil {
		return err
	}
	return nil
}

// Update a site
func (s *Store) UpdateSite(site *Site) error {
	updateStatement := `UPDATE site SET site_name=?, site_url=? WHERE id=?;`
	if _, err := s.db.Exec(updateStatement, site.SiteName, site.SiteUrl, site.Id); err != nil {
		return err
	}
	return nil
}

// Delete method
func (s *Store) DeleteSite(site *Site) error {
	deleteStatement := `DELETE FROM site WHERE siteId=?;`
	if _, err := s.db.Exec(deleteStatement, site.Id); err != nil {
		return err
	}
	return nil
}

// List all sites
func (s *Store) ListSites() ([]Site, error) {
	queryString := `SELECT id, site_name, site_url FROM site;`
	var siteList []Site
	rows, err := s.db.Query(queryString)
	if err != nil {
		return siteList, err
	}
	defer rows.Close()
	for rows.Next() {
		var site Site
		if err := rows.Scan(
			&site.Id,
			&site.SiteName,
			&site.SiteUrl,
		); err != nil {
			return nil, err
		}
		siteList = append(siteList, site)
	}
	return siteList, nil
}
