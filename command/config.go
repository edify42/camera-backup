package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
)

const (
	// ConfigFile is the filename we give which stores assets for how our app runs.
	ConfigFile string = "config.yaml"

	// DbFile is the local sql store of our file backups
	DbFile string = "sqlstore.db"

	// HiddenDir is our hidden directory name inside $HOME
	HiddenDir string = ".backup-genie"
)

// Config is the object we write to `config.yaml`
type Config struct {
	exclude      string
	include      string
	location     string
	lastModified uint64
	dbshasum     string
	config       []byte
}

// writeConfig will always write the ConfigFile to the current working directory.
func (c *Config) writeConfig() error {
	filePerms := os.FileMode(06660) // default ug+rw, a+r
	c.genYaml()
	ioutil.WriteFile(ConfigFile, c.config, filePerms)
	return nil
}

// genYaml creates the config as a []byte type
func (c *Config) genYaml() {
	data := fmt.Sprintf(
		`location: %s
lastModifiedDBStore: %s
\n`, c.location, "never")
	zap.S().Debug(data)
	// very simple yaml syntax, no need to overcomplicate it.
	// err := yaml.Unmarshal([]byte(data), &t)
	// if err != nil {
	// 	zap.S().Error("could not unmarshal data to yaml")
	// 	zap.S().Error(err)
	// }
	c.config = []byte(data)
}
