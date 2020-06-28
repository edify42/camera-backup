package filewalk

import (
	"testing"
)

func TestHandle_md5(t *testing.T) {
	tests := []struct {
		name  string
		h     *Handle
		input []byte
		want  string
	}{
		{
			name:  "A test case for md5",
			h:     &Handle{},
			input: []byte("These pretzels are making me thirsty."),
			want:  "b0804ec967f48520697662a204f5fe72",
		},
		{
			name:  "Another test case for md5",
			h:     &Handle{},
			input: []byte("62c0fa56b4b0bf8bf08f639a9ff1ec754f8ed78a"),
			want:  "ea3ef55bab35315e1604fa8fc4e342a4",
		},
		{
			name:  "Third test case for md5",
			h:     &Handle{},
			input: []byte("!HN(!*&#(*GHNOTHEUCROEFUOEUNHTONTEHU"),
			want:  "68302ba36a9d59ca5e7ec14cb912d273",
		},
		{
			name:  "Not so short test case for md5",
			h:     &Handle{},
			input: []byte("a"),
			want:  "0cc175b9c0f1b6a831c399e269772661",
		},
		{
			name:  "A final test case for md5",
			h:     &Handle{},
			input: []byte(""),
			want:  "d41d8cd98f00b204e9800998ecf8427e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handle{}
			if got := h.md5(tt.input); got != tt.want {
				t.Errorf("Handle.md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandle_sha1sum(t *testing.T) {
	tests := []struct {
		name  string
		h     *Handle
		input []byte
		want  string
	}{
		{
			name:  "A test case for md5",
			h:     &Handle{},
			input: []byte("These pretzels are making me thirsty."),
			want:  "7a0f82aac45ddc67ac3652f01fb5f731ec8f64a6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handle{}
			if got := h.sha1sum(tt.input); got != tt.want {
				t.Errorf("Handle.sha1sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
