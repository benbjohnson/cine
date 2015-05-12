package imdb

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Movie struct {
	Title         string
	Year          int
	EpisodeName   string
	EpisodeSeason int
	EpisodeNumber int
}

var movieRegexp = regexp.MustCompile(`^"(.+)" \((\d+)\) \{(.+?) \(#(\d+)\.(\d+)\)\}$`)

// SplitMovie returns parts of a movie title.
func SplitMovie(s string) (*Movie, error) {
	a := movieRegexp.FindStringSubmatch(s)
	if a == nil {
		return nil, errors.New("cannot match movie title")
	}

	// Parse year.
	year, err := strconv.Atoi(a[2])
	if err != nil {
		return nil, fmt.Errorf("parse year: %s", err)
	}

	// Parse episode season.
	season, err := strconv.Atoi(a[4])
	if err != nil {
		return nil, fmt.Errorf("parse episode season: %s", err)
	}

	// Parse episode number.
	num, err := strconv.Atoi(a[5])
	if err != nil {
		return nil, fmt.Errorf("parse episode number: %s", err)
	}

	// Create movie object.
	m := &Movie{
		Title:         a[1],
		Year:          year,
		EpisodeName:   a[3],
		EpisodeSeason: season,
		EpisodeNumber: num,
	}

	// TODO: Parse episode season & number.

	return m, nil
}
