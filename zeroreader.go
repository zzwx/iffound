package iffound

import (
	"io"
)

// ZeroReader implements an io.Reader that will always return 0, nil.
// It also satisfies Wrapper with ZeroReader.Unwrap to unwraps the error that
// caused ZeroReader to be used instead of actual reader.
//
// Using ZeroReader can be used in the environments where only one value
// is desired to be returned, but potentially allowing to check for error,
// by attempting to assert to ZeroReader.
type ZeroReader struct {
	Err error
}

// NewZeroReader returns a new ZeroReader.
func NewZeroReader(err error) ZeroReader {
	return ZeroReader{
		Err: err,
	}
}

// Read implements the io.Reader interface.
func (r ZeroReader) Read(b []byte) (n int, err error) {
	return 0, io.EOF
}

// Unwrap for ZeroReader returns an error that potentially caused
// ZeroReader to be returned.
func (r ZeroReader) Unwrap() error {
	return r.Err
}
