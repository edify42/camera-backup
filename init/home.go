package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/edify42/camera-backup/config"
	"github.com/edify42/camera-backup/localstore"
	"go.uber.org/zap"
)

// CheckHome will evaluate $HOME and see if any config exists in `${HOME}/.backup-genie/config.yaml`
func CheckHome() string {
	home := os.Getenv("HOME")
	if home == "" {
		zap.S().Errorf("No value for $HOME found")
		return "No $HOME set"
	}

	// We always try to create a HiddenDir in $HOME
	home = filepath.Clean(home)
	location := fmt.Sprintf("%s/%s", home, config.HiddenDir)
	err := os.Mkdir(location, 0700)

	if err != nil && err.Error() != fmt.Sprintf("mkdir %s: file exists", location) {
		zap.S().Errorf("Could not create file in %s; %v", location, err.Error())
		return fmt.Sprintf("Unable to create %s in $HOME", config.HiddenDir)
	}
	file := fmt.Sprintf("%s/%s", location, config.ConfigFile)
	zap.S().Infof("Checking for %s", file)
	if localstore.CheckFileExists(file) {
		return "Found" // Don't return location, this is known by convention.
	}
	return "No config found in $HOME"
}
