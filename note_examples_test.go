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
	note1, err := muse.NewNoteWithOctave(muse.C, muse.OctaveNumberDefault)
	if err != nil {
		panic(err)
	}

	note2 := note1.Copy()

	note3, err := muse.NewNoteFromString("C#")
	if err != nil {
		panic(err)
	}

	note4 := muse.MustNewNoteWithOctave(muse.CSHARP, muse.OctaveNumberDefault)

	note5, err := muse.NewNote(muse.C)
	if err != nil {
		panic(err)
	}

	fmt.Println(note1.IsEqualByName(note2), note3.IsEqualByName(note4), note1.IsEqualByName(note3), note1.IsEqualByName(note5))
	// Output: true true false true
}

// Create a note and set octave.
func ExampleNote_SetOctave() {
	// create octave #3
	octave, err := muse.NewOctave(muse.OctaveNumber4)
	if err != nil {
		panic(err)
	}

	// create a new note without any octave
	note := muse.C.NewNote()

	// set the octave to the note
	note.SetOctave(octave)

	fmt.Println(note.Octave().Name())
	// Output: FirstOctave
}

// Setting and Getting relative duration.
func ExampleNote_SetDurationRel() {
	// half note duration
	duration := muse.NewDurationRel(muse.DurationNameHalf)

	// creating note and setting duration
	note := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber3).SetDurationRel(duration)

	fmt.Println(note.DurationRel().Name())
	// Output: Half
}

// Getting note's time.Duration from note's duration.
func ExampleNote_GetTimeDuration() {
	// musical settings
	trackSettings := muse.TrackSettings{
		BPM:           uint64(80),
		Unit:          muse.Fraction{1, 2},
		TimeSignature: muse.Fraction{4, 4},
	}

	// half note duration
	duration := muse.NewDurationRel(muse.DurationNameHalf)

	// creating note and setting duration
	note := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber3).SetDurationRel(duration)

	fmt.Println(note.GetTimeDuration(trackSettings))
	// Output: 750ms
}

// Setting and Getting custom duration.
func ExampleNote_SetDurationAbs() {
	// creating note and setting custom duration
	note := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber3).SetDurationAbs(time.Second)

	fmt.Println(note.DurationAbs())
	// Output: 1s
}
