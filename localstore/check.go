package localstore

import (
	"database/sql"
	"fmt"

	"github.com/edify42/camera-backup/config"
	"go.uber.org/zap"
)

// Check will compare the input table to the backup location table.
// It always returns files that have changed their filename, filepath, etag or sha1sum.
// Use newRecords==true looks for new files.
// Use newRecords==false to look for missing/deleted files.
// Program needs to go through the results set to understand if file is missing/new or changed.
// Unoptimised for large record sets according to https://stackoverflow.com/questions/45069655/is-there-a-faster-way-to-compare-two-sqlite3-tables-in-python
func (c *Config) Check(newRecords bool, table string, db *sql.DB) ([]StoredFileRecord, error) {
	// TODO: Run checkTable type function to ensure it exists and has data...
	// Should check both `table` input and c.table...
	var result StoredFileRecord
	var checkRecords []StoredFileRecord
	var exceptQuery string
	query := "SELECT %s FROM %s %s SELECT %s FROM %s"
	attr := "filename, filepath, sha1sum, etag" // naturally checks all the file metadata values

	// switch `missing`: operation A-B vs B-A
	// Using 'EXCEPT', will find all records missing between the tables.
	if newRecords {
		exceptQuery = fmt.Sprintf(query, attr, table, "EXCEPT", attr, config.DataTable)
	} else {
		exceptQuery = fmt.Sprintf(query, attr, config.DataTable, "EXCEPT", attr, table)
	}
	zap.S().Infof("Except query is: %s", exceptQuery)
	resp, err := db.Query(exceptQuery)

	if err != nil {
		zap.S().Error("Could not run exceptQuery")
		return nil, err
	}

	for resp.Next() {
		err = resp.Scan(&result.Filename, &result.FilePath, &result.Sha1sum, &result.Etag)
		if err != nil {
			zap.S().Errorf("Could not scan exceptQuery result")
			return nil, err // No partial record returns - fix yo shit before continuing.
		}
		checkRecords = append(checkRecords, result)
		zap.S().Debugf("found this record in Check: %v which was different %v written at %v", result.Filename, result.Sha1sum, result.Timestamp)
	}

	return checkRecords, nil
}
