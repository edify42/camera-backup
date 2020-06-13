package command

import (
	"testing"

	_ "github.com/edify42/camera-backup/filewalk"
)

func TestConfig_RunInit(t *testing.T) {
	type fields struct {
		exclude      string
		include      string
		location     string
		lastModified uint64
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
				exclude:  "",
				include:  "",
				location: "yeah",
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
