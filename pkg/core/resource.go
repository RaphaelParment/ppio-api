package core

import "io"

// Resource represents an item exchange via HTTP
type Resource interface {
	ToJSON(w io.Writer) error
	FromJSON(r io.Reader) error
	Validate() error
	GetKey() interface{}
}
