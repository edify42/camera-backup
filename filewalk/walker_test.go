package filewalk

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"reflect"
	"testing"

	"github.com/edify42/camera-backup/config"
)

func TestWalkerConfig_returnMatch(t *testing.T) {
	type fields struct {
		Location string
		Exclude  []string
		Include  []string
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "basic test case",
			fields: fields{
				Exclude: []string{},
				Include: []string{".*"},
			},
			args: args{
				input: "basic-test-case",
			},
			want: true,
		},
		{
			name: "exclude all",
			fields: fields{
				Exclude: []string{".*"},
				Include: []string{".*"},
			},
			args: args{
				input: "basic-test-case",
			},
			want: false,
		},
		{
			name: "exclude sqlstore.db",
			fields: fields{
				Exclude: []string{config.DbFile},
				Include: []string{".*"},
			},
			args: args{
				input: config.DbFile,
			},
			want: false,
		},
		{
			name: "exclude path match",
			fields: fields{
				Exclude: []string{"/home/test/.*"},
				Include: []string{".*"},
			},
			args: args{
				input: "/home/test/my/file.test",
			},
			want: false,
		},
		{
			name: "include path match",
			fields: fields{
				Exclude: []string{"*.png"},
				Include: []string{"/home/hello/.*"},
			},
			args: args{
				input: "/home/hello/newfile.jpg",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalkerConfig{
				Location: tt.fields.Location,
				Exclude:  tt.fields.Exclude,
				Include:  tt.fields.Include,
			}
			if got := w.returnMatch(tt.args.input); got != tt.want {
				t.Errorf("WalkerConfig.returnMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Mocked out interface functions.

type testHandle struct{}

func (h *testHandle) sha1sum(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func (h *testHandle) md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (h *testHandle) etag(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (h *testHandle) loadFile(path string) []byte {
	return []byte(path)
}

func TestWalkerConfig_Walker(t *testing.T) {
	type fields struct {
		Location string
		Exclude  []string
		Include  []string
	}
	type args struct {
		fh Handler
	}

	myHandle := &testHandle{}
	var returnObject ReturnObject

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ReturnObject
		wantErr bool
	}{
		{
			name: "First test case",
			fields: fields{
				Location: ".",
				Exclude:  []string{".*"},
				Include:  []string{""},
			},
			args: args{
				fh: myHandle,
			},
			want:    returnObject,
			wantErr: false,
		},
		{
			name: "Get this test file only", // THIS TEST IS CURRENTLY BROKEN
			fields: fields{
				Location: ".",
				Exclude:  []string{""},
				Include:  []string{".*"},
			},
			args: args{
				fh: myHandle,
			},
			want:    returnObject,
			wantErr: false,
		},
		{
			name: "Error testing",
			fields: fields{
				Location: "here",
				Exclude:  []string{""},
				Include:  []string{""},
			},
			args: args{
				fh: myHandle,
			},
			want:    returnObject,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WalkerConfig{
				Location: tt.fields.Location,
				Exclude:  tt.fields.Exclude,
				Include:  tt.fields.Include,
			}
			got, err := w.Walker(tt.args.fh)
			if (err != nil) != tt.wantErr {
				t.Errorf("WalkerConfig.Walker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalkerConfig.Walker() = %v, want %v", got, tt.want)
			}
		})
	}
}
