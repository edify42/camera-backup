package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/edify42/camera-backup/config"
	"go.uber.org/zap"
)

// Config is the object we write to `config.yaml`
type Config struct {
	exclude      []string  `help:"regex of anything that we do not want to search/backup"`
	include      []string  `help:"regex of anything we DO want to search/backup"`
	location     string    `help:"fully qualified path of where the backup location is"`
	lastModified time.Time `help:"last write to the datastore. should match file metadata"` // TODO: write a functional test for this
	dbshasum     string    `help:"sha1sum of the datastore"`                                // TODO: write a test for this
	dryRun       bool      `help:"enable dry run mode so nothing is actually created"`
	config       []byte
}

// writeConfig will always write the ConfigFile to the current working directory.
func (c *Config) writeConfig() error {
	if c.dryRun {
		return nil // early return for dry run
	}
	filePerms := os.FileMode(06660) // default ug+rw, a+r
	c.genYaml()
	ioutil.WriteFile(config.ConfigFile, c.config, filePerms)
	return nil
}

// DryRun will prevent any of the config from actually running (testing purposes)
func (c *Config) DryRun() {
	c.dryRun = true
}

// genYaml creates the config as a []byte type
func (c *Config) genYaml() {
	data := fmt.Sprintf(
		`location: %s
lastModifiedDBStore: %s
`,
		c.location, "never")
	zap.S().Debug(data)
	// very simple yaml syntax, no need to overcomplicate it.
	// err := yaml.Unmarshal([]byte(data), &t)
	// if err != nil {
	// 	zap.S().Error("could not unmarshal data to yaml")
	// 	zap.S().Error(err)
	// }
	c.config = []byte(data)
}
