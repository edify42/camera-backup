package filewalk

import (
	"regexp"

	"github.com/karrick/godirwalk"
	"go.uber.org/zap"
)

// WalkerConfig will store working config
type WalkerConfig struct {
	Location string   `help:"Where the walker will begin searching"`
	Exclude  []string `help:"Which path/file patterns will be excluded"`
	Include  []string `help:"Which path/file patterns will be included"`
}

// FileWalk interface for mocking.
type FileWalk interface {
	Walker() ([]string, error)
}

// NewWalker returns a WalkerConfig
func NewWalker(location string, exclude, include []string) *WalkerConfig {
	return &WalkerConfig{
		Location: location,
		Exclude:  exclude,
		Include:  include,
	}
}

// Walker will return your files. It's responsible for filtering the files based on
// Include and Exclude.
func (w *WalkerConfig) Walker() ([]string, error) {
	var buff []string
	helper := &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			zap.S().Debugf("lets try match %s %s\n", w.Exclude, osPathname)
			w.returnMatch(osPathname)
			matched, err := regexp.MatchString(".*png", osPathname)
			if err != nil {
				return err
			}
			if matched {
				zap.S().Infof("we matched! %s\n", osPathname)
				buff = append(buff, osPathname)
			}
			return nil
		},
		PostChildrenCallback: func(osPathname string, de *godirwalk.Dirent) error {
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	}
	_ = godirwalk.Walk(w.Location, helper)
	zap.S().Infof("all the things: %v", buff)
	return buff, nil
}

// returnMatch will check the Include and Exclude options
func (w *WalkerConfig) returnMatch(file string) {
	return
}
