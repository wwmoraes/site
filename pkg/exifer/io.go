package exifer

import "io"

type Truncater interface {
	Truncate(size int64) error
}

type Descriptor interface {
	io.Reader
	io.Writer
	io.Seeker
	Truncater
}
