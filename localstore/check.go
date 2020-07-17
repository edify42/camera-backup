package localstore

import (
	"database/sql"
	"fmt"

	"github.com/edify42/camera-backup/config"
	"go.uber.org/zap"
)

// Check will compare the input table to the backup location table
func (c *Config) Check(table string, db *sql.DB) error {
	var result StoredFileRecord
	var records []StoredFileRecord
	query := fmt.Sprintf("SELECT filename, filepath, sha1sum, etag FROM %s EXCEPT SELECT filename, filepath, sha1sum, etag FROM main.%s", config.DataTable, table)
	zap.S().Infof(query)
	resp, err := db.Query(query)

	if err != nil {
		return err
	}

	for resp.Next() {
		err = resp.Scan(&result.Filename, &result.FilePath, &result.Sha1sum, &result.Etag)
		if err != nil {
			zap.S().Errorf("error happened: %v", err)
			return err
		}
		records = append(records, result)
		zap.S().Infof("found this record in Check: %v which was different %v written at %v", result.Filename, result.Sha1sum, result.Timestamp)
	}

	return nil
}
