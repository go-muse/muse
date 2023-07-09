package muse_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse"
)

// Creating a new relative duration and assigning it to the note.
func ExampleNewDurationWithRelativeValue() {
	// half note duration
	duration := muse.NewDurationWithRelativeValue(muse.DurationNameHalf)

	// creating note and setting duration
	note := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber3).SetDuration(*duration)

	fmt.Println(note.Duration().Name())
	// Output: Half
}

// Creating a new absolute duration and assigning it to the note.
func ExampleNewDurationWithAbsoluteValue() {
	// one second duration
	duration := muse.NewDurationWithAbsoluteValue(time.Second)

	// creating note and setting duration
	note := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber3).SetDuration(*duration)

	fmt.Println(note.GetAbsoluteDuration())
	// Output: 1s
}

// Getting time.Duration from duration.
func ExampleDuration_GetTimeDuration() {
	// musical settings
	trackSettings := muse.TrackSettings{
		BPM:           uint64(120),
		Unit:          muse.Fraction{1, 2},
		TimeSignature: muse.Fraction{4, 4},
	}

	// half note duration
	duration := muse.NewDurationWithRelativeValue(muse.DurationNameHalf)

	fmt.Println(duration.GetTimeDuration(trackSettings))
	// Output: 500ms
}

// Knowing the BPM, unit and the time signature(meter), calculate the number of bars in a minute.
func ExampleGetAmountOfBars() {
	trackSettings := muse.TrackSettings{
		BPM:           uint64(120),
		Unit:          muse.Fraction{1, 4},
		TimeSignature: muse.Fraction{4, 4},
	}

	amountOfBars := muse.GetAmountOfBars(trackSettings)

	fmt.Println(amountOfBars)
	// Output: 30
}
