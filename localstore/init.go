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

// SQLInit does stuff to initialise our code...
type SQLInit interface {
	CreateDB()
}

// NewLocalStore will create stuff...
func NewLocalStore(location string) Config {
	return Config{location}
}

// InitDB will be exposed externally for all to use!
func InitDB(i SQLInit) error {
	i.CreateDB()
	return nil
}

// CreateDB will simply touch the database file to ensure we can write to it.
func (c *Config) CreateDB() error {
	database := fmt.Sprintf("%s/%s", c.location, config.DbFile)
	zap.S().Infof("Creating the data store file at: %s", database)
	os.MkdirAll(c.location, 0755)
	os.Create(database)

	db, err := sql.Open("sqlite3", database)
	defer db.Close()
	if err != nil {
		zap.S().Errorf("Could not create the database %v", err)
		return err
	}

	// metadata table definition
	_, err = db.Exec(metadata)

	if err != nil {
		zap.S().Errorf("Could not load metadata schema %v", err)
		return err
	}

	// data table definition
	_, err = db.Exec(data)
	if err != nil {
		zap.S().Errorf("Could not load data schema %v", err)
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
