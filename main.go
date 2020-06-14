package main

import (
	"github.com/edify42/camera-backup/check"
	"github.com/edify42/camera-backup/command"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/jpillora/opts"
)

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}

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
	loggerMgr := initZapLog()
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync() // flushes buffer, if any
	logger := loggerMgr.Sugar()
	logger.Debug("Started the zap logger!")
	opts.New(&c).
		AddCommand(
			opts.New(&Init{}),
		).
		AddCommand(
			opts.New(&Check{}),
		).
		AddCommand(
			opts.New(&Scan{}),
		).
		Complete().
		Version(version).
		Parse().
		Run()
	zap.S().Infof("%+v", c)

	if len(c.Dir) != 0 {
		zap.S().Infof("hello")
	}

}

// Check is a placeholder...
type Check struct {
	Location string `opts:"help=location of the config.yaml file,default=."`
	ScanDir  string `opts:"help=target check directory,default=/tmp"`
}

// Run check - comprende?
func (c *Check) Run() {
	var config check.Config
	config.New(c.ScanDir)
	zap.S().Infof("Running the check against %s", config.ScanDir)
}

// Init does things to initialise the configuration
type Init struct {
	Location string   `opts:"help=specify where the config.yaml file will be dropped"`
	Include  []string `opts:"help=specify which file extensions should be included,default=.*"`
	Exclude  []string `opts:"help=exclude certain file extensions,default=nil"`
}

// Run will run init...yeah!
func (f *Init) Run() {
	var config command.Config
	if len(f.Location) > 0 {
		zap.S().Infof("Location of the config will be stored in %s", f.Location)
		config.NewLocation(f.Location)
	}

	if len(f.Include) > 0 {
		zap.S().Infof("Will look for files with extensions %s", f.Include)
		config.AddInclude(f.Include)
	} else {
		defaultInclude := []string{".*"}
		zap.S().Infof("Including all files by default %s", defaultInclude)
		config.AddInclude(defaultInclude)
	}

	if len(f.Exclude) > 0 {
		zap.S().Infof("Will exclude the following file and path matches %s", f.Exclude)
		config.AddExclude(f.Exclude)
	}

	err := config.RunInit()

	if err != nil {
		zap.S().Errorf("Not able to init - exiting because: %v", err)
	}
}

// Scan type for looking through a new directory structure.
type Scan struct {
	Location string `opts:"help=location of the config.yaml file,default=."`
	ScanDir  string `opts:"help=target scan directory,default=/tmp"`
}

func (s *Scan) Run() {
	zap.S().Infof("Running the scan against %s with config in location %s", s.ScanDir, s.Location)
}
