package imdb

import (
	"bufio"
	"io"
	"regexp"
)

var tabRegexp = regexp.MustCompile(`\t+`)

// Reader represents a parser for the IMDB file format.
type Reader struct {
	scanner *bufio.Scanner
	header  string
}

// NewReader returns a new instance of Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		scanner: bufio.NewScanner(r),
	}
}

// Header returns the header from the reader.
func (r *Reader) Header() (string, error) {
	if err := r.readHeader(); err != nil {
		return "", err
	}

	return r.header, nil
}

// Read reads the next record from the underlying reader.
func (r *Reader) Read() (record []string, err error) {
	// Read and ignore header.
	if err := r.readHeader(); err != nil {
		return nil, err
	}

	// Read single line.
	if !r.scanner.Scan() {
		return nil, r.scanner.Err()
	}

	// Split by tabs.
	record = tabRegexp.Split(r.scanner.Text(), -1)
	return
}

// readHeader reads in the entire file header up to a set of equal signs ("====").
func (r *Reader) readHeader() error {
	// Ignore if we already read the header.
	if r.header != "" {
		return nil
	}

	// Read all lines up to the equals signs.
	re := regexp.MustCompile(`^=+$`)
	for r.scanner.Scan() {
		line := r.scanner.Text()

		// If line matches the regex then read the next line and exit.
		// Otherwise append line to header.
		if re.MatchString(line) {
			r.scanner.Scan()
			return r.scanner.Err()
		} else {
			r.header += line + "\n"
		}
	}

	return r.scanner.Err()
}
