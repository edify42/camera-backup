package localstore

import (
	"database/sql"
	"fmt"
)

// Sqlstore does stuff
type Sqlstore interface {
	createConn() string
	testConn(db *sql.DB) bool
	CreateDB(db *sql.DB) error
	UpdateMetadata(db *sql.DB) error
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

	query := fmt.Sprintf(`
	INSERT INTO main.metadata (id, name, location, lastModified, absolute)
	VALUES (1, '%s', '%s', CURRENT_TIMESTAMP, true);`, c.name, c.location)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
