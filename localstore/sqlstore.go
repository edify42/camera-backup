package localstore

import "database/sql"

// Sqlstore does stuff
type Sqlstore interface {
	createConn() string
	// testConn() bool
	CreateDB(db *sql.DB) error
}

// func main() {
// 	fmt.Printf("hi there you've hit this package")
// }

func (c *Config) createConn() string {
	return c.location
}
