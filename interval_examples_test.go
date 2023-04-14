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

func ExampleNewIntervalByDegrees() {
	degree1 := &muse.NewDegree()
}
