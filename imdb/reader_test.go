package imdb_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/benbjohnson/cine/imdb"
)

// Ensure the header of a IMDB file can be retrieved.
func TestReader_Header(t *testing.T) {
	data := `` +
		`foo bar` + "\n" +
		`-------` + "\n" +
		`` + "\n" +
		`TITLE` + "\n" +
		`=====` + "\n" +
		`` + "\n" +
		`this is not the header` + "\n"

	// Read the first row of data.
	r := imdb.NewReader(strings.NewReader(data))
	header, err := r.Header()
	if err != nil {
		t.Fatal(err)
	} else if header != "foo bar\n-------\n\nTITLE\n" {
		t.Fatalf("unexpected header: %s", header)
	}
}

func TestReader_Read(t *testing.T) {
	// Mock data with weird header.
	data := `CRC: 0x475E5845  File: movies.list  Date: Fri Apr 24 00:00:00 2015` + "\n" +
		`-----------------------------------------------------------------------------` + "\n" +
		`` + "\n" +
		`MOVIES LIST` + "\n" +
		`===========` + "\n" +
		`` + "\n" +
		`"!Next?" (1994)						1994-1995` + "\n"

	// Read the first row of data.
	r := imdb.NewReader(strings.NewReader(data))
	row, err := r.Read()
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(row, []string{`"!Next?" (1994)`, `1994-1995`}) {
		t.Fatalf("unexpected row: %#v", row)
	}
}
