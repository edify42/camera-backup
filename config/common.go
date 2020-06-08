package config

// Config is the object we write to `config.yaml`
type Config struct {
	exclude      string
	include      string
	location     string `help:"fully qualified path of where the backup location is"`
	lastModified uint64
	dbshasum     string
	config       []byte
}
