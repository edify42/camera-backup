package localstore

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/edify42/camera-backup/config"
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
	ID int
}

// Metadata stores the returned metadata from the sqlstore
type Metadata struct {
	id           int
	name         string
	location     string
	lastModified string
	absolute     bool
	exclude      sql.NullString
	include      sql.NullString
}

// GetLocation will return the internal struct location
func (m *Metadata) GetLocation() string {
	return m.location
}

// GetExclude will fetch the excluded regex lookup
func (m *Metadata) GetExclude() []string {
	exc := strings.Split(m.exclude.String, config.RegexDivide)
	return exc
}

// GetInclude will fetch the include regex lookup
func (m *Metadata) GetInclude() []string {
	inc := strings.Split(m.include.String, config.RegexDivide)
	return inc
}

// GetLastModified will fetch the timestamp of the last modified datetime
func (m *Metadata) GetLastModified() string {
	return m.lastModified
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
	INSERT INTO main.metadata (id, name, location, lastModified, absolute, exclude, include)
	VALUES (1, ?, ?, CURRENT_TIMESTAMP, true, ?, ?);`
	_, err := db.Exec(query, c.name, c.location, c.include, c.exclude)
	if err != nil {
		return err
	}
	return nil
}

// ReadMetadata returns the metadata from the table
func (c *Config) ReadMetadata(db *sql.DB) (Metadata, error) {

	var r Metadata
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*")
	sb.From(config.MetadataTable)
	sb.Where(sb.In("id", 1))
	sql, args := sb.Build()
	resp, err := db.Query(sql, args...)
	if err != nil {
		return Metadata{}, err
	}
	defer resp.Close()
	if resp.Next() {
		err = resp.Scan(&r.id, &r.name, &r.location, &r.lastModified, &r.absolute, &r.exclude, &r.include)
		if err != nil {
			zap.S().Fatal(err)
			return Metadata{}, err
		}
	}

	fmt.Printf("hey hey hey: %v", r)

	return r, nil
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
	var result StoredFileRecord
	var records []StoredFileRecord
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*")
	sb.From(config.DataTable)
	// bit ugly but ok...
	if record.Etag != "" {
		sb.Where(sb.In("etag", record.Etag))
	}
	if record.Filename != "" {
		sb.Where(sb.In("filename", record.Filename))
	}
	if record.FilePath != "" {
		sb.Where(sb.In("filepath", record.FilePath))
	}
	if record.Sha1sum != "" {
		sb.Where(sb.In("sha1sum", record.Sha1sum))
	}
	sql, args := sb.Build()
	// fmt.Println(sql)
	// fmt.Println(args...)

	// query := `
	// SELECT * FROM main.data (filename, filepath, sha1sum, etag, lastCheckTimeStamp)
	// VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP);`
	resp, err := db.Query(sql, args...)

	if err != nil {
		return nil, err
	}

	for resp.Next() {
		resp.Scan(&result.ID, &result.Filename, &result.FilePath, &result.Etag, &result.Sha1sum)
		records = append(records, result)
	}

	// fmt.Printf("%v", resp)
	return records, nil
}
