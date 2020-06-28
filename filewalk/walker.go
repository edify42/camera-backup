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

// FileObject might be something we use.
type FileObject struct {
	name    string
	md5     string
	sha1sum string
	etag    string
}

// ReturnObject is a group of file objects
type ReturnObject []FileObject

// Walker will return your files. It's responsible for filtering the files based on
// Include and Exclude. Use a Handler interface of your choosing to calculate additional
// properties of the file including `sha1sum, etag and md5`
func (w *WalkerConfig) Walker(fh Handler) (ReturnObject, error) {
	var buff []string
	var returnObject ReturnObject
	var fileObject FileObject
	helper := &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			zap.S().Debugf("lets try match %s %s\n", w.Exclude, osPathname)
			if de.IsDir() {
				return nil
			}

			matched := w.returnMatch(osPathname)
			if matched {
				zap.S().Debugf("matched: %s\n", osPathname)
				file := fh.loadFile(osPathname)
				sha1sum := fh.sha1sum(file)
				zap.S().Debugf("sha1sum of the file is %s", sha1sum) // TODO: make the logging better
				md5 := fh.md5(file)
				zap.S().Debugf("md5sum of the file is %s", md5)
				etag := fh.etag(file)
				zap.S().Debugf("etag of the file is %s", etag)
				fileObject.name = osPathname
				fileObject.md5 = md5
				fileObject.etag = etag
				fileObject.sha1sum = sha1sum
				buff = append(buff, osPathname)
				returnObject = append(returnObject, fileObject)
			}
			return nil
		},
		PostChildrenCallback: func(osPathname string, de *godirwalk.Dirent) error {
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	}
	_ = godirwalk.Walk(w.Location, helper)
	zap.S().Debugf("all the things: %v", buff)
	return returnObject, nil
}

// returnMatch will check the Include and Exclude options
// Exclude has higher priority (executed first)
func (w *WalkerConfig) returnMatch(input string) bool {
	for _, regex := range w.Exclude {
		match, err := regexp.MatchString(regex, input)
		if err != nil {
			zap.S().Errorf("failed to execute exclude match for string %s - regex was: %s", input, regex)
		}
		if match {
			return false
		}
	}
	for _, regex := range w.Include {
		match, err := regexp.MatchString(regex, input)
		if err != nil {
			zap.S().Errorf("failed to execute include match for string %s - regex was: %s", input, regex)
		}
		if match {
			return true
		}
	}
	return false
}
