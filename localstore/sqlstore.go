package localstore

import (
	"database/sql"
	"os"
)

// Sqlstore does stuff
type Sqlstore interface {
	createConn() string
	testConn(db *sql.DB) bool
	CreateFile(string) error
	CreateDB(db *sql.DB) error
	UpdateMetadata(db *sql.DB) error
}

// CreateFile will create a blank file to disk
func (c *Config) CreateFile(file string) error {
	_, err := os.Create(file)
	if err != nil {
		return err
	}
	return nil
}

// FileRecord Type is what we want to update/modify
type FileRecord struct {
	Filename string
	FilePath string
	Sha1sum  string
	Etag     string
}

func (c *Config) createConn() string {
	return c.location
}

// testConn will check the database
func (c *Config) testConn(db *sql.DB) bool {
	return true
}

// UpdateMetadata will simply update the metadata table. Should only be called after write to table.
func (c *Config) UpdateMetadata(db *sql.DB) error {

	query := `
	INSERT INTO main.metadata (id, name, location, lastModified, absolute)
	VALUES (1, ?, ?, CURRENT_TIMESTAMP, true);`
	_, err := db.Exec(query, c.name, c.location)
	if err != nil {
		return err
	}
	return nil
}

// WriteFileRecord will write the file and associated metadata to the data table
func (c *Config) WriteFileRecord(record FileRecord, db *sql.DB) error {
	query := `
	INSERT INTO main.data (filename, filepath, sha1sum, etag, lastCheckTimeStamp)
	VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP);`
	_, err := db.Exec(query, record.Filename, record.FilePath, record.Sha1sum, record.Etag)
	if err != nil {
		return err
	}
	return nil
}
