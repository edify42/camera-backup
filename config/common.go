package config

import (
	"math/rand"
	"strings"
	"time"
)

// Config is the object we write to `config.yaml`
type Config struct {
	exclude      string
	include      string
	location     string `help:"fully qualified path of where the backup location is"`
	lastModified uint64
	dbshasum     string
	config       []byte
}

// RandomName will return a random name that can be used
func RandomName(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
