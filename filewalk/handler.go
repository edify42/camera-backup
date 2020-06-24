package filewalk

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

// Handler is my best fwiend
type Handler interface {
	md5([]byte) string
	etag([]byte) string
	loadFile(string) []byte
}

// Handle struct...
type Handle struct{}

// NewHandler returns a famous struct
func NewHandler() *Handle {
	return &Handle{}
}

func (h *Handle) md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// TODO: finish this off...
func (h *Handle) etag(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Required function to actually do the work of reading a file.
// should not ever be mocked!
func (h *Handle) loadFile(file string) []byte {
	dat, err := ioutil.ReadFile(file)
	check(err)
	return dat
}

// copied from gobyexample.com docs!
func check(e error) {
	if e != nil {
		panic(e)
	}
}
