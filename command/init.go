package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
)

const (
	// ConfigFile is the filename we give
	ConfigFile string = "config.yaml"
	// HiddenDir is our hidden directory name inside $HOME
	HiddenDir string = ".backup-genie"
)

// Input is the allowed type...
type Input struct {
	Location string
	config   []byte
}

// Config is the object we write to `config.yaml`
type Config struct {
	location     string
	lastModified uint64
	dbshasum     string
}

// Reader is a generic interface to read our file.
type Reader interface {
	readFile() []byte
}

// readFile for Input struct
func (i *Input) readFile() []byte {
	content, err := ioutil.ReadFile(i.Location)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("File contents: %s", content)
	return content
}

// RunInit will take care of stuff
func (i *Input) RunInit() error {

	// Early exit of RunInit if we want to use config found in $HOME
	configExistsHome := CheckHome()
	if configExistsHome != "" {
		promptHome := promptui.Prompt{
			Label:     "Config found in $HOME - use this instead of running init?",
			IsConfirm: true,
		}
		useHome, err := promptHome.Run()

		if err != nil {
			log.Errorf("Prompt for home check failed %v\n", err)
			return err
		}

		log.Infof("Using home? %b", useHome)

		if strings.ToLower(useHome) == "y" {
			return nil
		}
	}

	// Early exit of RunInit if we find config in current working directory

	prompt := promptui.Prompt{
		Label: "Enter the location of the datastore (leave blank for current working directory): ",
	}

	location, err := prompt.Run()

	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return err
	}

	i.Location = fmt.Sprintf("%s", location)

	log.Info(i) // should point to i with the value of the input
	err = i.writeConfig()
	if err != nil {
		log.Errorf("Could not write to location %s\n", i.Location)
		return err
	}

	return nil

}

// writeConfig will
func (i *Input) writeConfig() error {
	// check if config already exists
	filePerms := os.FileMode(06660) // default ug+rw, a+r
	i.genYaml()
	ioutil.WriteFile("config.yaml", i.config, filePerms)
	return nil
}

// genYaml creates the config as a []byte type
func (i *Input) genYaml() {
	data := fmt.Sprintf("location: %s\n", i.Location)
	log.Debug(data)
	// very simple yaml syntax, no need to overcomplicate it.
	// err := yaml.Unmarshal([]byte(data), &t)
	// if err != nil {
	// 	log.Error("could not unmarshal data to yaml")
	// 	log.Error(err)
	// }
	i.config = []byte(data)
}
