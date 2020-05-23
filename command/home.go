package command

import (
	"fmt"
	"os"
	
	"github.com/edify42/camera-backup/config"
	"go.uber.org/zap"
)

// CheckHome will evaluate $HOME and see if any config exists in `${HOME}/.backup-genie/config.yaml`
func CheckHome() string {
	home := os.Getenv("HOME")
	if home == "" {
		zap.S().Errorf("No value for $HOME found")
		return ""
	}
	// TODO: regex to strip trailing slash in $HOME if exists.
	// create the directory in home
	location := fmt.Sprintf("%s/%s", home, config.HiddenDir)
	err := os.Mkdir(location, 0700)

	if err != nil && err.Error() != fmt.Sprintf("mkdir %s: file exists", location) {
		zap.S().Infof("Could not create file in %s; %v", location, err.Error())
	}
	file := fmt.Sprintf("%s/%s", location, config.ConfigFile)
	zap.S().Infof("Checking for %s", file)
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		return location
	}
	return ""
}
