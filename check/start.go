package check

import (
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

	// // Database stuff
	// database := fmt.Sprintf("%s/%s", c.Location, config.DbFile)
	// db, err := sqlConf.GetSqliteDB(database)
	// if err != nil {
	// 	zap.S().Errorf("Error while getting database in init command: %s", c.location)
	// 	return err
	// }

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
		// err := sqlConf.WriteFileRecord(record, db)
		// if err != nil {
		// 	zap.S().Errorf("Error while writing to the database in init command: %s", c.location)
		// 	return err
		// }
	}

	return results, nil
}
