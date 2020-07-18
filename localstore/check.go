package localstore

import (
	"database/sql"
	"fmt"

	"github.com/edify42/camera-backup/config"
	"go.uber.org/zap"
)

// Check will compare the input table to the backup location table
// Using 'EXCEPT', will find all records missing between the tables.
// Unoptimised for large record sets according to https://stackoverflow.com/questions/45069655/is-there-a-faster-way-to-compare-two-sqlite3-tables-in-python
func (c *Config) Check(table string, db *sql.DB) error {
	var result StoredFileRecord
	var missingRecords []StoredFileRecord
	query := "SELECT %s FROM %s %s SELECT %s FROM %s"
	attr := "filename, filepath, sha1sum, etag" // naturally checks all the file metadata values
	exceptQuery := fmt.Sprintf(query, attr, config.DataTable, "EXCEPT", attr, table)
	zap.S().Infof("Except query is: %s", exceptQuery)
	resp, err := db.Query(exceptQuery)

	if err != nil {
		zap.S().Error("Could not run exceptQuery")
		return err
	}

	for resp.Next() {
		err = resp.Scan(&result.Filename, &result.FilePath, &result.Sha1sum, &result.Etag)
		if err != nil {
			zap.S().Errorf("Could not scan exceptQuery result")
			return err
		}
		missingRecords = append(missingRecords, result)
		zap.S().Debugf("found this record in Check: %v which was different %v written at %v", result.Filename, result.Sha1sum, result.Timestamp)
	}

	return nil
}
