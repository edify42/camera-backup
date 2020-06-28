package localstore

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
)

func TestConfig_CreateDB(t *testing.T) {
	// mock the *sql.DB interface
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec("CREATE TABLE metadata").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("CREATE TABLE data").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
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
			name: "First test case",
			fields: fields{
				location: "nowhere",
				name:     "here",
			},
			args: args{
				db: db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				location: tt.fields.location,
				name:     tt.fields.name,
			}
			if err := c.CreateDB(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Config.CreateDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Positive test args
type conf struct {
	location string
}

func (c *conf) createConn() string {
	return "/tmp/no_location"
}

func (c *conf) testConn(db *sql.DB) bool {
	return true
}

func (c *conf) CreateDB(db *sql.DB) error {
	return nil
}

func (c *conf) UpdateMetadata(db *sql.DB) error {
	return nil
}

func (c *conf) CreateFile(input string) error {
	return nil
}

func TestInitDB(t *testing.T) {
	type args struct {
		i Sqlstore
	}

	c := &conf{location: "aoeu"}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test",
			args: args{
				i: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitDB(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("InitDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
