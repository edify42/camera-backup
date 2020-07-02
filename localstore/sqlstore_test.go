package localstore

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/edify42/camera-backup/config"
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

func TestConfig_ReadFileRecord(t *testing.T) {
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
		record FileRecord
		db     *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []StoredFileRecord
		wantErr bool
	}{
		{
			name: "First test case",
			fields: fields{
				location: "here",
				name:     "there",
			},
			args: args{
				record: FileRecord{
					Filename: "1",
					Etag:     "0",
				},
				db: db,
			},
			want:    []StoredFileRecord{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		query := fmt.Sprintf("SELECT (.+) FROM %s WHERE etag IN (?) AND filename IN (?)", config.DataTable)
		rows := sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "one").
			AddRow(2, "two")
		mock.ExpectQuery(query).
			WillReturnRows(rows).
			WithArgs(tt.args.record.Etag, tt.args.record.Filename)
		// mock.ExpectCommit()
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				location: tt.fields.location,
				name:     tt.fields.name,
			}
			got, err := c.ReadFileRecord(tt.args.record, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ReadFileRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.ReadFileRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
