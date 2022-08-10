package utils

import (
	"bytes"
	"io"
)

func IoReaderChangeBytes(reader io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.Bytes()
}
