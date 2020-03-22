package localstore

// Sqlstore does stuff
type Sqlstore interface {
	createConn() string
	testConn() bool
}

// func main() {
// 	fmt.Printf("hi there you've hit this package")
// }

// Connection is the thing that will do stuff.
type Connection struct {
	file  string
	debug bool
}

func (c *Connection) createConn() error {

}
