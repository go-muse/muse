package muse_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse"
)

// Create and compare notes and print their names.
//
//nolint:dupword
func ExampleNote() {
	note1, err := muse.NewNote(muse.C, muse.OctaveNumberDefault)
	if err != nil {
		panic(err)
	}

	note2 := note1.Copy()

	note3, err := muse.NewNoteFromString("C#")
	if err != nil {
		panic(err)
	}

	note4 := muse.MustNewNote(muse.CSHARP, muse.OctaveNumberDefault)

	fmt.Println(note1.IsEqualByName(note2), note3.IsEqualByName(note4), note1.IsEqualByName(note3))
	// Output: true true false
}

// Create a note and set octave.
func ExampleNote_SetOctave() {
	// create octave #3
	octave, err := muse.NewOctave(muse.OctaveNumber4)
	if err != nil {
		panic(err)
	}

	// create a new note without any octave
	note := muse.MustNewNoteWithoutOctave(muse.C)

	// set the octave to the note
	note.SetOctave(octave)

	fmt.Println(note.Octave().Name())
	// Output: FirstOctave
}

// Setting and Getting duration.
func ExampleNote_SetDuration() {
	// half note duration
	duration := muse.NewDurationWithRelativeValue(muse.DurationNameHalf)

	// creating note and setting duration
	note := muse.MustNewNote(muse.C, muse.OctaveNumber3).SetDuration(*duration)

	fmt.Println(note.Duration().Name())
	// Output: Half
}

// Getting note's time.Duration from note's duration.
func ExampleNote_TimeDuration() {
	// musical settings
	trackSettings := muse.TrackSettings{
		BPM:           uint64(80),
		Unit:          muse.Fraction{1, 2},
		TimeSignature: muse.Fraction{4, 4},
	}

	// half note duration
	duration := muse.NewDurationWithRelativeValue(muse.DurationNameHalf)

	// creating note and setting duration
	note := muse.MustNewNote(muse.C, muse.OctaveNumber3).SetDuration(*duration)

	fmt.Println(note.TimeDuration(trackSettings))
	// Output: 750ms
}

// Setting and Getting custom duration.
func ExampleNote_SetAbsoluteDuration() {
	// creating note and setting custom duration
	note := muse.MustNewNote(muse.C, muse.OctaveNumber3).SetAbsoluteDuration(time.Second)

	fmt.Println(note.GetAbsoluteDuration())
	// Output: 1s
}
