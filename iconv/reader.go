package iconv

import (
	"bytes"
	"io"
	"os/exec"
)

// Reader converts data from one character encoding to another.
type Reader struct {
	cmd    *exec.Cmd
	stdout io.ReadCloser
	stderr bytes.Buffer
}

// NewReader returns a new instance of Reader.
func NewReader(r io.Reader, fromEncoding, toEncoding string) (*Reader, error) {
	// Create reader.
	rd := &Reader{}

	// Create command.
	rd.cmd = exec.Command(`iconv`, `-f`, fromEncoding, `-t`, toEncoding)
	rd.cmd.Stdin = r
	rd.cmd.Stderr = &rd.stderr

	// Save stdout.
	stdout, err := rd.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	rd.stdout = stdout

	// Execute command.
	if err := rd.cmd.Start(); err != nil {
		return nil, err
	}

	// Return reader.
	return rd, nil
}

// Read reads p bytes from the underlying reader.
func (r *Reader) Read(p []byte) (n int, err error) {
	return r.stdout.Read(p)
}
