package localstore

import (
	"database/sql"
)

// Sqlstore does stuff
type Sqlstore interface {
	createConn() string
	testConn(db *sql.DB) bool
	CreateDB(db *sql.DB) error
	UpdateMetadata(db *sql.DB) error
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
