package iffound

import (
	"io"
	"io/ioutil"
	"os"
	"runtime"
)

type ifFound struct {
	fileName string
}

// IfFound creates a pathway Reader and Content
// methods, handy functions that return a file content or emptiness
// if file not found.
func IfFound(fileName string) *ifFound {
	return &ifFound{
		fileName: fileName,
	}
}

// Reader returns the file of ifFound as reader, if found. It's designed to return no error.
// It registers a finalizer to close the file as its reference becomes unattended.
//
// Returns a ZeroReader with the error wrapped if the file can't be opened for any reason.
//
// Usage of ZeroReader allows for asserting to it if a need rises to
// confirm error condition as well as unwrap, because ZeroReader
// contains the error that cased it to be used instead of actual file reader.
//
// if _, ok := reader.(file.ZeroReader); ok {
func (iff ifFound) Reader() io.Reader {
	f, err := os.Open(iff.fileName)
	if err == nil {
		runtime.SetFinalizer(f, closeFileFinalizer)
		return f
		//defer f.Close()
		//var buf bytes.Buffer
		//io.Copy(&buf, f)
		//return &buf
	}
	return NewZeroReader(err)
}

func closeFileFinalizer(f *os.File) {
	if f != nil {
		_ = f.Close()
		runtime.SetFinalizer(f, nil)
	}
}

// Bytes returns the content of the ifFound file as []byte.
// If file is not found, returns an empty []byte{}.
func (iff ifFound) Bytes() []byte {
	file, err := ioutil.ReadFile(iff.fileName)
	if err == nil {
		return file
	}
	return []byte{}
}

// Content returns the content of the ifFound file as string.
// If file is not found, returns "".
func (iff ifFound) String() string {
	file, err := ioutil.ReadFile(iff.fileName)
	if err == nil {
		return string(file)
	}
	return ""
}
