package check

import (
	"fmt"

	"github.com/edify42/camera-backup/config"
	"github.com/edify42/camera-backup/localstore"
)

type Config struct {
	Location string
}

func (c *Config) Store(checkDir string) {
	c.Location = checkDir
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
	sql.ReadMetadata(db)
}
