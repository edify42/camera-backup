package filewalk

import (
	"github.com/karrick/godirwalk"
	"go.uber.org/zap"
)

// WalkerConfig will store working config
type WalkerConfig struct {
	Location string `help:"Where the walker will begin searching"`
}

// FileWalk interface for mocking.
type FileWalk interface {
	Walker() ([]string, error)
}

// NewWalker returns a WalkerConfig
func NewWalker(location string) *WalkerConfig {
	return &WalkerConfig{
		Location: location,
	}
}

// Walker will return your files.
func (w *WalkerConfig) Walker() ([]string, error) {
	var buff []string
	buff = append(buff, "string")
	helper := &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			zap.S().Infof("%s %s\n", de.ModeType(), osPathname)
			buff = append(buff, osPathname)
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	}
	_ = godirwalk.Walk(w.Location, helper)
	zap.S().Infof("all the things: %v", buff)
	return buff, nil
}
