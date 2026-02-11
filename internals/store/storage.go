package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" //driver
	"github.com/pradeept/site-monitor-cli/internals/logger"
)

type Site struct {
	Id          int
	SiteName    string
	SiteUrl     string
	RequestTime int64
}

type SiteStatus struct {
	Id         int
	SiteId     int
	StatusCode int
	StatusText string
}

type Store struct {
	db  *sql.DB
	log *log.Logger
}

func NewStore(dbString string) (*Store, error) {
	log.Println(dbString)
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
		site_url TEXT NOT NULL UNIQUE,
		request_time INT DEFAULT 1
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

// Find method
func (s *Store) FindSite(siteName string, siteURL string) ([]Site, error) {
	findStatement := `SELECT * FROM site WHERE site_name=? OR site_url=?;`

	rows, err := s.db.Query(findStatement, siteName, siteURL)
	if err != nil {
		return nil, fmt.Errorf("[Error] failed to fetch the site: ", err)
	}
	defer rows.Close()
	var sites []Site
	for rows.Next() {
		var site Site
		if err := rows.Scan(
			&site.Id,
			&site.SiteName,
			&site.SiteUrl,
			&site.RequestTime,
		); err != nil {
			return nil, fmt.Errorf("[Error] Failed to scan the rows: ", err)
		}
		sites = append(sites, site)
	}
	return sites, nil
}

// Insert method
func (s *Store) InsertSite(site *Site) error {
	insertStatement := `INSERT INTO site(site_name, site_url, request_time) VALUES(?,?,?);`

	exists, err := s.FindSite(site.SiteName, site.SiteUrl)
	if err != nil {
		return fmt.Errorf("[Error] validating new site: ", err)
	}
	if len(exists) > 0 {
		return fmt.Errorf("[Error] Site already exists")
	}

	if _, err := s.db.Exec(insertStatement, site.SiteName, site.SiteUrl, site.RequestTime); err != nil {
		return err
	}
	return nil
}

// Update a site
func (s *Store) UpdateSite(site *Site) error {
	updateStatement := `UPDATE site SET site_name=?, site_url=?, request_time=? WHERE id=?;`
	if _, err := s.db.Exec(updateStatement, site.SiteName, site.SiteUrl, site.RequestTime, site.Id); err != nil {
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
	queryString := `SELECT id, site_name, site_url, request_time FROM site;`
	var siteList []Site
	rows, err := s.db.Query(queryString)
	if err != nil {
		return siteList, err
	}
	defer rows.Close()
	// for every row create a site struct
	// and store it into site list
	for rows.Next() {
		var site Site
		if err := rows.Scan(
			&site.Id,
			&site.SiteName,
			&site.SiteUrl,
			&site.RequestTime,
		); err != nil {
			return nil, err
		}
		siteList = append(siteList, site)
	}
	return siteList, nil
}

// -------- request --------------

// List all requests of a site
func (s *Store) ListSiteRequests(siteId string) ([]SiteStatus, error) {
	queryString := `SELECT id, site_id, status_code, status_text FROM request r WHERE site_id=?;`

	rows, err := s.db.Query(queryString, siteId)
	if err != nil {
		return nil, err
	}

	var siteList []SiteStatus

	defer rows.Close()

	for rows.Next() {
		var siteStatus SiteStatus
		if err := rows.Scan(
			&siteStatus.Id,
			&siteStatus.SiteId,
			&siteStatus.StatusCode,
			&siteStatus.StatusText,
		); err != nil {
			return nil, err
		}
	}
	return siteList, nil
}

// Insert a sitestatus
func (s *Store) InsertSiteRequest(st SiteStatus) error {
	queryString := `INSERT into request(site_id, status_code, status_text) VALUES(?,?,?);`

	if _, err := s.db.Exec(queryString, st.SiteId, st.StatusCode, st.StatusText); err != nil {
		return err
	}

	return nil
}
