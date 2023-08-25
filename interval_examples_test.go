package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Interval can be calculated between diatonic degrees.
func ExampleNewIntervalByDegrees() {
	// For example, degree1 is a second degree in a mode, and in contains note "D"
	degree1 := muse.NewDegree(2, 2, nil, nil, muse.MustNewNoteWithOctave(muse.D, muse.OctaveNumber0), nil, nil)
	// degree2 is the fourth degree in a mode, and it contains note "F"
	degree2 := muse.NewDegree(4, 5, nil, nil, muse.MustNewNoteWithOctave(muse.F, muse.OctaveNumber0), nil, nil)

	interval, err := muse.NewIntervalByDegrees(degree1, degree2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("HalfTones: %d, interval name: %s, short name: %s\n", interval.HalfTones(), interval.Name(), interval.ShortName())
	// Output: HalfTones: 3, interval name: MinorThird, short name: m3
}
