package config

const (
	// ConfigFile is the filename we give which stores assets for how our app runs.
	ConfigFile string = "config.yaml"

	// DbFile is the local sql store of our file backups
	DbFile string = "sqlstore.db"

	// HiddenDir is our hidden directory name inside $HOME
	HiddenDir string = ".backup-genie"

	// DataTable is what is referenced in our sql queries.
	DataTable string = "main.data"

	// MetadataTable is where we store our search settings!
	MetadataTable string = "main.metadata"

	// RegexDivide is the custom character sequence to divide regex for include and exclude
	RegexDivide string = "||"

	// Sqlstore is the name we write into the metadata
	Sqlstore string = "sqlstore-v1"
)
