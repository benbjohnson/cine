package iconv_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/benbjohnson/cine/iconv"
)

func TestReader_Read(t *testing.T) {
	// iso-8859-1 encoded "él"
	b := []byte{0xE9, 0x6C}

	// Create reader.
	r, err := iconv.NewReader(bytes.NewReader(b), "ISO-8859-1", "UTF-8")
	if err != nil {
		t.Fatal(err)
	}

	// Read all bytes as UTF-8.
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if string(buf) != `él` {
		t.Fatalf("unexpected bytes: %s", buf)
	}
}
