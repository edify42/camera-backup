package filewalk

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
)

// Handler is my best fwiend
type Handler interface {
	etag([]byte) string
	loadFile(string) []byte
	md5([]byte) string
	sha1sum([]byte) string
}

// Handle struct...
type Handle struct{}

// NewHandler returns a famous struct
func NewHandler() *Handle {
	return &Handle{}
}

func (h *Handle) sha1sum(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func (h *Handle) md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// TODO: finish this off...
func (h *Handle) etag(data []byte) string {
	// Always be splicing 8MB chucks
	chunkSize := 8 * 1024 * 1024
	if len(data) < chunkSize {
		return fmt.Sprintf("%x", md5.Sum(data))
	}

	var md5s []byte
	chunks := split(data, chunkSize)
	for _, v := range chunks {
		md5 := md5.Sum(v)
		a := md5[:]
		md5s = append(md5s, a...)
	}
	b := fmt.Sprintf("%x", md5.Sum(md5s))
	return fmt.Sprintf("%s-%d", b, len(chunks))
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

// copied from https://gist.github.com/xlab/6e204ef96b4433a697b3
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
