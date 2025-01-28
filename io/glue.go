package io

import (
	gio "io"
)

type GlueReadWriteCloser struct {
	Reader gio.Reader
	Writer gio.Writer
}

func (grwc GlueReadWriteCloser) Write(b []byte) (int, error) {
	return grwc.Writer.Write(b)
}
func (grwc GlueReadWriteCloser) Read(p []byte) (int, error) {
	return grwc.Reader.Read(p)
}
func (grwc GlueReadWriteCloser) Close() error {

	if readClose, ok := (grwc.Reader).(gio.Closer); ok {
		e := readClose.Close()
		if e != nil {
			return e
		}
	}
	if writeClose, ok := (grwc.Writer).(gio.Closer); ok {
		e := writeClose.Close()
		if e != nil {
			return e
		}
	}
	return nil
}

func Glue2(read gio.Reader, write gio.Writer) *GlueReadWriteCloser {
	return &GlueReadWriteCloser{
		Reader: read,
		Writer: write,
	}
}
