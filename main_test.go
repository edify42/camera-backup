package main

import "testing"

func TestInit_Run(t *testing.T) {
	type fields struct {
		Location string
		Include  []string
		Exclude  []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "testing",
			fields: fields{
				Include: []string{},
				Exclude: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Init{
				Location: tt.fields.Location,
				Include:  tt.fields.Include,
				Exclude:  tt.fields.Exclude,
			}
			f.Run()
		})
	}
}
