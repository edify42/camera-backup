package localstore

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/edify42/camera-backup/config"
	"github.com/manifoldco/promptui"
	_ "github.com/mattn/go-sqlite3" // coment
	"go.uber.org/zap"
)

// Config is the SQL DB Config object
type Config struct {
	location string
	name     string `help:"default to sqlstore-v1"`
	exclude  string `help:"need to write and convert the excluded regex patterns"`
	include  string `help:"handle the array to string conversion here"`
	table    string `help:"which table we should read/write from"`
}

// NewLocalStore will create stuff...
func NewLocalStore(location, table string, include, exclude []string) *Config {

	inc := strings.Join(include[:], config.RegexDivide)
	exc := strings.Join(exclude[:], config.RegexDivide)
	return &Config{
		location,
		config.Sqlstore,
		inc,
		exc,
		table,
	}
}

// InitDB create the datastore file and run CreateDB
func InitDB(i Sqlstore) error {
	location := i.createConn()
	database := fmt.Sprintf("%s/%s", location, config.DbFile)
	zap.S().Infof("Creating the data store file at: %s", database)
	os.MkdirAll(location, 0755) // `mkdir -p $location` #first
	// Check if a current sqlstore exists
	if CheckFileExists(database) {
		prompt := promptui.Prompt{
			Label: "Overwrite the existing sqlstore file (Y/n)? ",
		}

		answer, err := prompt.Run()
		if err != nil {
			zap.S().Errorf("Prompt failed %v\n", err)
			return err
		}
		if answer == "n" {
			return nil // early return
		}
	}

	err := i.CreateFile(database)
	if err != nil {
		zap.S().Errorf("Could not create the database file %v", err)
		return err
	}

	db, err := i.GetSqliteDB(database)
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
	dataTable := fmt.Sprintf(data, "data")
	_, err = db.Exec(dataTable)
	if err != nil {
		return err
	}

	// everything worked!
	return nil
}

// CreateTempTable will create and return a temporary table to store records
func (c *Config) CreateTempTable(db *sql.DB) (string, error) {

	// give us a name
	name := config.RandomName(8)
	c.table = name
	dataTable := fmt.Sprintf(data, name)
	_, err := db.Exec(dataTable)
	if err != nil {
		return "", err
	}

	// everything worked!
	return name, nil
}

// DropTempTable will drop the above temp table
func (c *Config) DropTempTable(table string, db *sql.DB) error {
	// maybe we should blacklist a set of tables which shouldn't ever be dropped.
	// TODO: the above... simply write a blacklist array...
	query := fmt.Sprintf("DROP  TABLE %s", table)
	_, err := db.Exec(query)
	if err != nil {
		zap.S().Errorf("DropTempTable could not drop table: %s", table)
		return err
	}
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
CREATE TABLE %s (
	filename	varchar(255) NOT NULL,
	filepath	varchar(255) NOT NULL,
	sha1sum		varchar(255) NOT NULL,
	etag		varchar(255) NOT NULL,
	lastCheckTimestamp	TIMESTAMP NOT NULL
)
`
)

// CheckFileExists is a reuseable function which returns true if a file exists at a known location
func CheckFileExists(file string) bool {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		return true // Don't return location, this is known by convention.
	}
	return false
}
