package config

const (
	// ConfigFile is the filename we give which stores assets for how our app runs.
	ConfigFile string = "config.yaml"

	// DbFile is the local sql store of our file backups
	DbFile string = "sqlstore.db"

	// HiddenDir is our hidden directory name inside $HOME
	HiddenDir string = ".backup-genie"
)