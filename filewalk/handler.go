package filewalk

// Handler is my best fwiend
type Handler interface {
	md5() string
}

// Handle struct...
type Handle struct{}

// NewHandler returns a famous struct
func NewHandler() *Handle {
	return &Handle{}
}

func (h *Handle) md5() string {
	return ""
}
