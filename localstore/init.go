package localstore

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/edify42/camera-backup/config"
	_ "github.com/mattn/go-sqlite3" // coment
	"go.uber.org/zap"
)

// Config is the SQL DB Config object
type Config struct {
	location string
	name     string `help:"default to sqlstore-v1"`
}

// NewLocalStore will create stuff...
func NewLocalStore(location string) *Config {
	return &Config{
		location,
		"sqlstore-v1",
	}
}

// InitDB create the datastore file and run CreateDB
func InitDB(i Sqlstore) error {
	location := i.createConn()
	database := fmt.Sprintf("%s/%s", location, config.DbFile)
	zap.S().Infof("Creating the data store file at: %s", database)
	os.MkdirAll(location, 0755) // `mkdir -p $location` #first
	os.Create(database)
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		zap.S().Errorf("Could not create the database %v", err)
		return err
	}
	defer db.Close()

	err = i.CreateDB(db)
	if err != nil {
		zap.S().Errorf("Could not initialise the database within InitDB method %v", err)
		return err
	}

	err = i.UpdateMetadata(db)
	if err != nil {
		zap.S().Errorf("Update table metadata within InitDB method %v", err)
		return err
	}
	return nil
}

// CreateDB will bootstrap the datastore file with the schema
func (c *Config) CreateDB(db *sql.DB) error {

	// metadata table definition
	_, err := db.Exec(metadata)

	if err != nil {
		return err
	}

	// data table definition
	_, err = db.Exec(data)
	if err != nil {
		return err
	}

	// everything worked!
	return nil
}

const (
	// metadata table looks a bit like how we store things in the Config struct
	metadata string = `
CREATE TABLE metadata (
	id				int NOT NULL PRIMARY KEY,
	name 			varchar(30),
	location 		varchar(255),
	lastModified	TIMESTAMP,
	absolute		boolean,
	exclude 		varchar(255),
	include 		varchar(255)
)
`
	data string = `
CREATE TABLE data (
	filename	varchar(255) NOT NULL,
	filepath	varchar(255) NOT NULL,
	sha1sum		varchar(255) NOT NULL,
	etag		varchar(255) NOT NULL,
	lastCheckTimestamp	TIMESTAMP NOT NULL
)
`
)
