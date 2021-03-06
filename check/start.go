package check

import (
	"fmt"

	"github.com/edify42/camera-backup/config"
	"github.com/edify42/camera-backup/filewalk"
	"github.com/edify42/camera-backup/localstore"
	"go.uber.org/zap"
)

// Check interface is our mockable thing...
type Check interface {
	GetFiles(string) []string
}

// StartNewFileCheck implements a filewalk.
func (c *Config) StartNewFileCheck(checker Check) error {
	checker.GetFiles(c.Location)
	return nil
}

// StartDB

// GetFiles will call the filewalk functions to grab me some files!
// It creates a temp table in the datastore and compares the two tables.
func (c *Config) GetFiles(location string) ([]string, error) {
	var results []string
	// Try to run the filewalk...
	handler := filewalk.Handle{}
	walker := filewalk.NewWalker(c.Location, c.exclude, c.include)
	array, err := walker.Walker(&handler)

	if err != nil {
		zap.S().Errorf("GetFiles failed to call walker: %v", err)
		return results, err
	}

	// Attempt to create the database after the config is initialised
	sqlConf := localstore.NewLocalStore(c.Location, "noTable", c.include, c.exclude)

	// Database stuff
	database := fmt.Sprintf("%s/%s", c.Location, config.DbFile)
	db, err := sqlConf.GetSqliteDB(database)
	if err != nil {
		zap.S().Errorf("Error while getting database in GetFiles command: %s", c.Location)
		return nil, err
	}

	table, _ := sqlConf.CreateTempTable(db)

	for _, fileObject := range array {
		zap.S().Debugf("single record dump %v", fileObject)
		record := localstore.FileRecord{
			Filename: fileObject.Name,
			FilePath: fileObject.Path,
			Sha1sum:  fileObject.Sha1sum,
			Etag:     fileObject.Etag,
		}
		results = append(results, record.FilePath)

		zap.S().Infof("hey there grumpy: %v", record)
		err := sqlConf.WriteFileRecordTempTable(record, db)
		if err != nil {
			zap.S().Errorf("Error while writing to the database in GetFiles command: %s", c.Location)
			return nil, err
		}
	}

	// Big logic dump...lets break this up later.
	var newFiles []localstore.FileRecord
	var changedFiles []localstore.FileRecord
	var missingFiles []localstore.FileRecord

	// Check returns changed or new files.
	arr, err := sqlConf.Check(true, table, db)
	if err != nil {
		zap.S().Fatalf("first check could not be run: %v", err)
	}

	// Check returns changed or missing files.
	arr2, err := sqlConf.Check(false, table, db)
	if err != nil {
		zap.S().Fatalf("first check could not be run: %v", err)
	}

	// Join arr and arr2, excluding duplicates
	for _, record := range arr2 {
		arr = localstore.AppendFileRecordIfMissing(arr, record)
	}

	for _, record := range arr {
		zap.S().Debugf("Determining what happened to %v", record)
		thing := localstore.FileRecord{
			Filename: record.Filename,
		}
		// Attempt read on the data table - confirm it's a new file.
		readOldResults, err := sqlConf.ReadFileRecord(thing, config.DataTable, db)
		zap.S().Debugf("hello? %v", readOldResults)
		if err != nil {
			zap.S().Errorf("what happened to my result? %v", err)
			return nil, err
		}
		readNewResults, err := sqlConf.ReadFileRecord(record.FileRecord, table, db)
		if len(readOldResults) == 0 {
			zap.S().Debug("we got a new file here")
			newFiles = append(newFiles, record.FileRecord)
		} else if len(readNewResults) == 0 {
			zap.S().Debug("we got a missing file here")
			missingFiles = append(missingFiles, record.FileRecord)
		} else {
			// Must have been a changed file!
			zap.S().Debugf("Made it here. - something changed!")
			changedFiles = append(changedFiles, record.FileRecord)
		}
	}

	// sqlConf.DropTempTable(table, db)
	zap.S().Debugf("new files: %v", newFiles)
	zap.S().Debugf("changed files: %v", changedFiles)
	zap.S().Debugf("missing files: %v", missingFiles)

	return results, nil
}
