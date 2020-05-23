package command

import "testing"

func TestCheckHome(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Successful testing",
			want: "",
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
