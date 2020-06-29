package localstore

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/huandu/go-sqlbuilder"
	"go.uber.org/zap"
)

// Sqlstore does stuff
type Sqlstore interface {
	createConn() string
	testConn(db *sql.DB) bool
	CreateFile(string) error
	CreateDB(db *sql.DB) error
	GetSqliteDB(string) (*sql.DB, error)
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

// StoredFileRecord is the FileRecord which is stored in the database
type StoredFileRecord struct {
	FileRecord
	id int
}

func (c *Config) createConn() string {
	return c.location
}

// testConn will check the database
func (c *Config) testConn(db *sql.DB) bool {
	return true
}

// GetSqliteDB returns the sql.DB object
func (c *Config) GetSqliteDB(database string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		zap.S().Errorf("Could not open the database %v", err)
		return nil, err
	}
	return db, nil
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

// WriteFileRecord will write the file info to the data table
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

// ReadFileRecord will return a set of results based on the input parameters
func (c *Config) ReadFileRecord(record FileRecord, db *sql.DB) ([]StoredFileRecord, error) {
	var result []StoredFileRecord
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", sb.As("COUNT(*)", "c"))
	sb.From("user")
	sb.Where(sb.In("status", 1, 2, 5))
	sb.Where(sb.In("Hello", "yeah"))
	sql, args := sb.Build()
	fmt.Println(sql)
	fmt.Println(args)

	query := `
	SELECT * FROM main.data (filename, filepath, sha1sum, etag, lastCheckTimeStamp)
	VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP);`
	resp, err := db.Exec(query, record.Filename, record.FilePath, record.Sha1sum, record.Etag)

	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", resp)
	return result, nil
}
