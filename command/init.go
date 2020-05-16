package command

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
	"go.uber.org/zap"
)

// Reader is a generic interface to read our config file.
type Reader interface {
	readFile() []byte
}

// readFile for Input struct
func (c *Config) readFile() []byte {
	content, err := ioutil.ReadFile(c.location)
	if err != nil {
		zap.S().Fatal(err)
	}

	zap.S().Debugf("File contents: %s", content)
	return content
}

// NewLocation will write the location path to config
func (c *Config) NewLocation(location string) {
	c.location = location
}

// AddInclude for the files we want to include
func (c *Config) AddInclude(include string) {
	c.include = include
}

// AddExclude for the files we want to ignore
func (c *Config) AddExclude(exclude string) {
	c.exclude = exclude
}

// RunInit will take care of stuff
func (c *Config) RunInit() error {

	// Early exit of RunInit if we want to use config found in $HOME
	configExistsHome := CheckHome()
	if configExistsHome != "" {
		promptHome := promptui.Prompt{
			Label:     "Config found in $HOME - use this instead of running init?",
			IsConfirm: true,
		}
		useHome, err := promptHome.Run()

		if err != nil {
			zap.S().Errorf("Prompt for home check failed %v\n", err)
			return err
		}

		zap.S().Infof("Using home? %b", useHome)

		if strings.ToLower(useHome) == "y" {
			return nil
		}

	}

	// TODO: Early exit of RunInit if we find config in current working directory

	prompt := promptui.Prompt{
		Label: "Enter the directory of the datastore (leave blank for current working directory): ",
	}

	location, err := prompt.Run()
	// TODO: Add more path validation around location? - for now just run path.Clean
	cleanLocation := path.Clean(location)

	c.location = fmt.Sprintf("%s", cleanLocation)

	if err != nil {
		zap.S().Errorf("Prompt failed %v\n", err)
		return err
	}

	zap.S().Debugf("%v", c) // should point to config with the value of the input
	err = c.writeConfig()
	if err != nil {
		zap.S().Errorf("Could not write to location %s\n", c.location)
		return err
	}

	// Attempt to create the database after the config is initialised
	c.CreateDB()

	return nil

}
