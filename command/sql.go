package command

import (
	"os"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // coment
	"go.uber.org/zap"
)

// CreateDB will simply touch the database file to ensure we can write to it.
func (c *Config)CreateDB() error {
	database := fmt.Sprintf("%s/%s", c.location, DbFile)
	zap.S().Infof("Creating the data store file at: %s", database)
	os.MkdirAll(c.location, 0755)
	os.Create(database)

	db, err := sql.Open("sqlite3", database)
	defer db.Close()
	if err != nil {
		zap.S().Errorf("Could not create the database %v", err)
		return err
	}

	_, err = db.Exec("CREATE TABLE `metadata` (`name` INTEGER PRIMARY KEY AUTOINCREMENT)")

	if err != nil {
		zap.S().Errorf("Could not metadata schema %v", err)
	}
	
	// everything worked!
	return nil
}