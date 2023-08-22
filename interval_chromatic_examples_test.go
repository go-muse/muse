package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Chromatic (acoustic) interval is just a name for a certain quantity of intervals.
func ExampleNewIntervalChromatic() {
	interval, err := muse.NewIntervalChromatic(11)
	if err != nil {
		panic(err)
	}

	fmt.Println(interval.Name(), interval.ShortName())
	// Output: MajorSeventh M7
}

// Diatonic interval is defined by the number of semitones and steps in a scale.
func ExampleNewIntervalByHalfTonesAndDegrees() {
	halfTones := muse.HalfTones(4)
	degrees := muse.DegreeNum(3)

	// Four halftones between three degrees
	diatonicInterval, err := muse.NewIntervalByHalfTonesAndDegrees(halfTones, degrees)
	if err != nil {
		panic(err)
	}

	// Just four halftones
	interval, err := muse.NewIntervalChromatic(halfTones)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Diatonic interval of %d halfTones (between %d degrees): %s\nChromatic interval of %d halfTones: %s\n", halfTones, degrees, diatonicInterval.Name(), halfTones, interval.Name())
	// Output: Diatonic interval of 4 halfTones (between 3 degrees): DiminishedFourth
	// Chromatic interval of 4 halfTones: MajorThird
}

// Get chromatic interval by it's name.
func ExampleNewIntervalByName() {
	interval, err := muse.NewIntervalByName(muse.IntervalNameMajorSixth)
	if err != nil {
		panic(err)
	}

	fmt.Printf("HalfTones: %d, interval name: %s\n", interval.HalfTones(), interval.Name())
	// Output: HalfTones: 9, interval name: MajorSixth
}

// Making note by existing note and known interval.
func ExampleMakeNoteByIntervalName() {
	// The note from which we will make the second note by the interval.
	firstNote := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumberDefault)

	// Needed interval
	interval, err := muse.NewIntervalChromatic(6) // six halftones means tritone interval
	if err != nil {
		panic(err)
	}

	// We can get the second note from the name extracted from interval
	secondNote1, err := muse.MakeNoteByIntervalName(firstNote, interval.Name())
	if err != nil {
		panic(err)
	}

	// or we can get it directly by the interval name
	secondNote2, err := muse.MakeNoteByIntervalName(firstNote, muse.IntervalNameTritone) // tritone interval
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
func ExampleMakeDegreeByIntervalName() {
	firstDegree := muse.NewDegree(1, 0, nil, nil, muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber1), nil, nil)

	secondDegree, err := muse.MakeDegreeByIntervalName(firstDegree, muse.IntervalNameTritone)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Second Interval Degree Number: %d, half tones from prime: %d, Note name: %s\n",
		secondDegree.Number(),
		secondDegree.HalfTonesFromPrime(),
		secondDegree.Note().Name(),
	)
	// Output: Second Interval Degree Number: 2, half tones from prime: 6, Note name: F#
}
