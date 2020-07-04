package check

import (
	"fmt"

	"github.com/edify42/camera-backup/config"
	"github.com/edify42/camera-backup/localstore"
	"go.uber.org/zap"
)

// Config contains...
type Config struct {
	Location string
}

// New will create a new instance...
func (c *Config) New(checkDir string) {
	if len(checkDir) == 0 {
		c.Location = "." // default look at current working directory
		return
	}

	c.Location = checkDir
	sql := localstore.NewLocalStore(c.Location)
	dbfile := fmt.Sprintf("%s/%s", c.Location, config.DbFile)
	db, _ := sql.GetSqliteDB(dbfile)

	metadata, err := sql.ReadMetadata(db)
	if err != nil {
		zap.S().Fatalf("Could not read metadata from table: %v", err)
		return
	}
	// TODO: check logic that metadata.Location should ALWAYS equal c.Location?
	zap.S().Infof("Using the sqlstore in location: %s", metadata.GetLocation())
	zap.S().Infof("Last update to the sqlstore was done on: %s", metadata.GetLastModified())
}
