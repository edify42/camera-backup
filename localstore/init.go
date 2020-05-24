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
}

// NewLocalStore will create stuff...
func NewLocalStore(location string) *Config {
	return &Config{location}
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
	name 		varchar(30),
	absolute	boolean,
	location 	varchar(255),
	exclude 	varchar(255),
	include 	varchar(255)
)
`
	data string = `
CREATE TABLE data (
	filename	varchar(255),
	filepath	varchar(255),
	sha1sum		varchar(255),
	lastCheckTimestamp	varchar(50)
)
`
)
