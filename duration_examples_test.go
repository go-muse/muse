package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Creating a new relative duration and assigning it to the note.
func ExampleNewDurationWithRelativeValue() {
	// half note duration
	duration := muse.NewDurationRel(muse.DurationNameHalf)

	// creating note and setting duration
	note := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber3).SetDurationRel(duration)

	fmt.Println(note.DurationRel().Name())
	// Output: Half
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
	duration := muse.NewDurationRel(muse.DurationNameHalf)

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
