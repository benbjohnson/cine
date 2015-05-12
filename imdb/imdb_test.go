package imdb_test

import (
	"reflect"
	"testing"

	"github.com/benbjohnson/cine/imdb"
)

func TestSplitMovie(t *testing.T) {
	var tests = []struct {
		s string
		m *imdb.Movie
	}{
		{
			s: `"#1 Single" (2006) {Cats and Dogs (#1.4)}`,
			m: &imdb.Movie{
				Title:         "#1 Single",
				Year:          2006,
				EpisodeName:   "Cats and Dogs",
				EpisodeSeason: 1,
				EpisodeNumber: 4,
			},
		},
	}

	for i, tt := range tests {
		m, err := imdb.SplitMovie(tt.s)
		if err != nil {
			t.Errorf("%d. %s: error: %s", i, tt.s, err)
		} else if !reflect.DeepEqual(tt.m, m) {
			t.Errorf("%d. %s: mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.m, m)
		}
	}
}
