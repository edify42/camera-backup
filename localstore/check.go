package localstore

import "database/sql"

// Check will compare the input table to the backup location table
func (c *Config) Check(table string, db *sql.DB) error {
	return nil
}
