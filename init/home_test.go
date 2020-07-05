package command

import (
	"os"
	"testing"
)

func TestAccCheckHome(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Successful testing with no config in $HOME",
			want: "No config found in $HOME",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckHome(); got != tt.want {
				t.Errorf("CheckHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccCheckHomeMissingEnv(t *testing.T) {
	os.Setenv("HOME", "")
	want := "No $HOME set"
	t.Run("Missing $HOME directory in ENV", func(t *testing.T) {
		if got := CheckHome(); got != want {
			t.Errorf("CheckHome() = %v, want %s", got, want)
		}
	})
}
