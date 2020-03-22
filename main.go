package main

import (
	"github.com/edify42/temp-golang/command"
	log "github.com/sirupsen/logrus"

	"github.com/jpillora/opts"
)

//set this via ldflags (see https://stackoverflow.com/q/11354518)
var version = "0.0.0"

// Config is our main thing...
type Config struct {
	Dir            string `opts:"help=target directory for search"`
	CheckBackup    bool   `opts:"help=check every file in the database exists in their saved location"`
	Config         string `opts:"help=location of the config.yaml file"`
	DataSourceFile string `opts:"help=local sqlite datasource (default ~/.backup-genie/sqlite.db)"`
	Lines          int    `opts:"help=number of lines to show"`
	SyncBackup     bool   `opts:"help=use CheckBackup first to see if any files are missing. Requires a search path"`
	UpdateBackup   bool   `opts:"help=forced update of the backup"`
}

func main() {
	c := Config{}
	opts.New(&c).
		AddCommand(
			opts.New(&Init{}),
		).
		Complete().
		Version(version).
		Parse().
		Run()
	log.Printf("%+v", c)

	if len(c.Dir) != 0 {
		log.Info("hello")
	}

}

// Init does things to initialise the configuration
type Init struct {
	Location string `opts:"help=specify where the config.yaml file will be dropped"`
}

// Run will run init...yeah!
func (f *Init) Run() {
	if len(f.Location) > 0 {
		log.Infof("Location of the config will be stored in %s", f.Location)
	}

	var input command.Input
	input.Location = f.Location

	err := input.RunInit()

	if err != nil {
		log.Errorf("Not able to init - exiting because: %v", err)
	}
}
