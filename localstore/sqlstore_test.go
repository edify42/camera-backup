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
		exclude  string
		include  string
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
				exclude:  ".*",
				include:  ".*",
			},
			args: args{
				db: db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		mock.ExpectExec("INSERT INTO main.metadata").
			WithArgs(tt.fields.name, tt.fields.location, tt.fields.include, tt.fields.exclude).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				location: tt.fields.location,
				name:     tt.fields.name,
				exclude:  tt.fields.exclude,
				include:  tt.fields.include,
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
			want: []StoredFileRecord{
				{
					FileRecord: FileRecord{
						Filename: "1",
						FilePath: "hello",
						Etag:     "0",
						Sha1sum:  "0",
					},
					ID: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		query := fmt.Sprintf(`SELECT \* FROM %s WHERE etag IN \(\?\) AND filename IN \(\?\)`, config.DataTable)
		rows := sqlmock.NewRows([]string{"id", "filename", "filepath", "sha1sum", "etag"}).
			AddRow(1, "1", "hello", "0", "0")
		mock.ExpectQuery(query).
			WillReturnRows(rows).
			WithArgs(tt.args.record.Etag, tt.args.record.Filename)
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				location: tt.fields.location,
				name:     tt.fields.name,
			}
			got, err := c.ReadFileRecord(tt.args.record, config.DataTable, tt.args.db)
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

func TestConfig_ReadMetadata(t *testing.T) {
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
		want    Metadata
		wantErr bool
	}{
		{
			name: "Positive test case to read Metadata",
			fields: fields{
				location: "a location",
				name:     "something that i can delete?",
			},
			args: args{
				db: db,
			},
			want:    Metadata{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		query := fmt.Sprintf(`SELECT \* FROM %s WHERE id IN \(\?\)`, config.MetadataTable)
		rows := sqlmock.NewRows([]string{"id", "filename"})
		mock.ExpectQuery(query).
			WillReturnRows(rows).
			WithArgs(1)
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				location: tt.fields.location,
				name:     tt.fields.name,
			}
			got, err := c.ReadMetadata(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ReadMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.ReadMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
