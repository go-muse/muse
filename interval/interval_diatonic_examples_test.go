package interval_test

import (
	"fmt"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/interval"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

// Diatonic interval can be calculated between diatonic degrees.
func ExampleNewDiatonic() {
	// For example, degree1 is a second degree in a mode, and in contains note "D"
	degree1 := degree.New(2, 2, nil, nil, note.MustNewNoteWithOctave(note.D, octave.Number0), nil, nil)
	// degree2 is the fourth degree in a mode, and it contains note "F"
	degree2 := degree.New(4, 5, nil, nil, note.MustNewNoteWithOctave(note.F, octave.Number0), nil, nil)

	interval, err := interval.NewDiatonic(degree1, degree2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("HalfTones: %d, interval name: %s, short name: %s\n", interval.HalfTones(), interval.Name(), interval.ShortName())
	// Output: HalfTones: 3, interval name: MinorThird, short name: m3
}
