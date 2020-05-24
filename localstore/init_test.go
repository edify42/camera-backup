package localstore

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
)

func TestConfig_CreateDB(t *testing.T) {
	//
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
