package command

import (
	"testing"
	"time"

	"github.com/edify42/camera-backup/config"
	_ "github.com/edify42/camera-backup/filewalk"
)

func TestConfig_RunInit(t *testing.T) {
	type fields struct {
		exclude      []string
		include      []string
		location     string
		lastModified time.Time
		dbshasum     string
		dryRun       bool
		config       []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "DryRun mode of init",
			fields: fields{
				exclude:  []string{},
				include:  []string{},
				location: "RunInitFunction",
				dryRun:   true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				exclude:      tt.fields.exclude,
				include:      tt.fields.include,
				location:     tt.fields.location,
				lastModified: tt.fields.lastModified,
				dbshasum:     tt.fields.dbshasum,
				dryRun:       tt.fields.dryRun,
				config:       tt.fields.config,
			}
			if err := c.RunInit(); (err != nil) != tt.wantErr {
				t.Errorf("Config.RunInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_AddExclude(t *testing.T) {
	type fields struct {
		exclude      []string
		include      []string
		location     string
		lastModified time.Time
		dbshasum     string
		dryRun       bool
		config       []byte
	}
	type args struct {
		exclude []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test bare minimum excludes files",
			fields: fields{},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				exclude:      tt.fields.exclude,
				include:      tt.fields.include,
				location:     tt.fields.location,
				lastModified: tt.fields.lastModified,
				dbshasum:     tt.fields.dbshasum,
				dryRun:       tt.fields.dryRun,
				config:       tt.fields.config,
			}
			c.AddExclude(tt.args.exclude)
			if c.exclude[0] != config.DbFile {
				t.Errorf("Oh man it didn't equal %s", config.DbFile)
			}
		})
	}
}
