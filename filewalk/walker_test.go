package filewalk

import "testing"

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
