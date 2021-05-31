package filewalk

import (
	"testing"
)

func TestHandle_getMd5(t *testing.T) {
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
			if got := h.getMd5(tt.input); got != tt.want {
				t.Errorf("Handle.getMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandle_getSha1sum(t *testing.T) {
	tests := []struct {
		name  string
		h     *Handle
		input []byte
		want  string
	}{
		{
			name:  "A test case for sha1",
			h:     &Handle{},
			input: []byte("These pretzels are making me thirsty."),
			want:  "7a0f82aac45ddc67ac3652f01fb5f731ec8f64a6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handle{}
			if got := h.getSha1sum(tt.input); got != tt.want {
				t.Errorf("Handle.getSha1sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandle_getEtag(t *testing.T) {
	// Local testing - to be removed and switched out with something more...testable...

	dat := make([]byte, 185*1024*256)
	tests := []struct {
		name  string
		h     *Handle
		input []byte
		want  string
	}{
		{
			name:  "A test case for small file == md5sum",
			h:     &Handle{},
			input: []byte("These pretzels are making me thirsty."),
			want:  "b0804ec967f48520697662a204f5fe72",
		},
		{
			name:  "A test case for large file == etag",
			h:     &Handle{},
			input: dat,
			want:  "8079245890315d5bdcfcbbf7a6977ac1-6",
		},
	}

	// makes a file which is roughly 46MB
	// _ = ioutil.WriteFile("file.txt", dat, 0644)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handle{}
			if got := h.getEtag(tt.input); got != tt.want {
				t.Errorf("Handle.etag() = %v, want %v", got, tt.want)
			}
		})
	}
}
