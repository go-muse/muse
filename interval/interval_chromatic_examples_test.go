package interval_test

import (
	"fmt"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/interval"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

// Chromatic (acoustic) interval is just a name for a certain quantity of intervals.
func ExampleNewChromatic() {
	interval, err := interval.NewChromatic(11)
	if err != nil {
		panic(err)
	}

	fmt.Println(interval.Name(), interval.ShortName())
	// Output: MajorSeventh M7
}

// Diatonic interval is defined by the number of semitones and steps in a scale.
func ExampleNewIntervalByHalfTonesAndDegrees() {
	halfTones := halftone.HalfTones(4)
	degrees := degree.Number(3)

	// Four halftone between three degrees
	diatonicInterval, err := interval.NewIntervalByHalfTonesAndDegrees(halfTones, degrees)
	if err != nil {
		panic(err)
	}

	// Just four halftone
	interval, err := interval.NewChromatic(halfTones)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Diatonic interval of %d halfTones (between %d degrees): %s\nChromatic interval of %d halfTones: %s\n", halfTones, degrees, diatonicInterval.Name(), halfTones, interval.Name())
	// Output: Diatonic interval of 4 halfTones (between 3 degrees): DiminishedFourth
	// Chromatic interval of 4 halfTones: MajorThird
}

// Get chromatic interval by it's name.
func ExampleNewIntervalByName() {
	interval, err := interval.NewIntervalByName(interval.NameMajorSixth)
	if err != nil {
		panic(err)
	}

	fmt.Printf("HalfTones: %d, interval name: %s\n", interval.HalfTones(), interval.Name())
	// Output: HalfTones: 9, interval name: MajorSixth
}

// Making note by existing note and known interval.
func ExampleMakeNoteByName() {
	// The note from which we will make the second note by the interval.
	firstNote := note.MustNewNoteWithOctave(note.C, octave.NumberDefault)

	// Needed interval
	chromaticInterval, err := interval.NewChromatic(6) // six halftone means tritone interval
	if err != nil {
		panic(err)
	}

	// We can get the second note from the name extracted from interval
	secondNote1, err := interval.MakeNoteByName(firstNote, chromaticInterval.Name())
	if err != nil {
		panic(err)
	}

	// or we can get it directly by the interval name
	secondNote2, err := interval.MakeNoteByName(firstNote, interval.NameTritone) // tritone interval
	if err != nil {
		panic(err)
	}

	// The notes made from the same note and interval must be equal
	if !secondNote1.IsEqualByName(secondNote2) {
		panic("notes aren't equal")
	}

	fmt.Printf("the second note is: %s", secondNote1.Name())
	// Output: the second note is: F#
}

// Making degree by existing degree and known interval.
func ExampleMakeDegreeByName() {
	firstDegree := degree.New(1, 0, nil, nil, note.MustNewNoteWithOctave(note.C, octave.Number1), nil, nil)

	secondDegree, err := interval.MakeDegreeByName(firstDegree, interval.NameTritone)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Second Diatonic Degree Number: %d, halftone from prime: %d, Note name: %s\n",
		secondDegree.Number(),
		secondDegree.HalfTonesFromPrime(),
		secondDegree.Note().Name(),
	)
	// Output: Second Diatonic Degree Number: 2, halftone from prime: 6, Note name: F#
}
