package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Creating a new duration and assigning it to the note.
func ExampleNewDuration() {
	// half note duration
	duration := muse.NewDuration(muse.DurationNameHalf)

	// creating note and setting duration
	note := muse.MustNewNote(muse.C, muse.OctaveNumber3).SetDuration(duration)

	fmt.Println(note.GetDuration().Name())
	// Output: Half
}

// Getting time.Duration from duration.
func ExampleDuration_TimeDuration() {
	// musical settings
	bpm := uint64(120)
	unit := &muse.Fraction{1, 2}
	timeSignature := &muse.Fraction{4, 4}

	// half note duration
	duration := muse.NewDuration(muse.DurationNameHalf)

	fmt.Println(duration.TimeDuration(bpm, unit, timeSignature))
	// Output: 500ms
}
