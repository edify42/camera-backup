package localstore

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConfig_WriteFileRecord(t *testing.T) {
	// mock the *sql.DB interface
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	record := FileRecord{
		Filename: "filename",
		FilePath: "filepath/filepath",
		Sha1sum:  "aoeuaoeu",
		Etag:     "e-tag-23",
	}
	defer db.Close()
	type fields struct {
		location string
		name     string
	}
	type args struct {
		record FileRecord
		db     *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid Write File test",
			fields: fields{
				location: "a_location",
				name:     "a_name",
			},
			args: args{
				record: record,
				db:     db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectExec("INSERT INTO main.data").
				WithArgs(tt.args.record.Filename, tt.args.record.FilePath, tt.args.record.Sha1sum, tt.args.record.Etag).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			c := &Config{
				location: tt.fields.location,
				name:     tt.fields.name,
			}
			if err := c.WriteFileRecord(tt.args.record, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Config.WriteFileRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_UpdateMetadata(t *testing.T) {
	// mock the *sql.DB interface
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	type fields struct {
		location string
		name     string
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid Update testcase",
			fields: fields{
				location: "here",
				name:     "there",
			},
			args: args{
				db: db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		mock.ExpectExec("INSERT INTO main.metadata").
			WithArgs(tt.fields.name, tt.fields.location).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				location: tt.fields.location,
				name:     tt.fields.name,
			}
			if err := c.UpdateMetadata(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Config.UpdateMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}