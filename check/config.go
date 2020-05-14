package check

type Config struct {
	ScanDir string
}

func (c *Config) Store(checkDir string) {
	c.ScanDir = checkDir
}

// New will create a new instance...
func (c *Config) New(checkDir string) {
	var a string = "/tmp"
	if len(checkDir) == 0 {
		c.ScanDir = a
		return
	}
	c.ScanDir = checkDir
}
