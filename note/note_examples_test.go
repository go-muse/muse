package note_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
	"github.com/go-muse/muse/track"
)

// Create and compare notes and print their names.
func ExampleNote() {
	note1, err := note.NewNoteWithOctave(note.C, octave.NumberDefault)
	if err != nil {
		panic(err)
	}

	note2 := note1.Copy()

	note3, err := note.NewNoteFromString("C#")
	if err != nil {
		panic(err)
	}

	note4 := note.MustNewNoteWithOctave(note.CSHARP, octave.NumberDefault)

	note5, err := note.New(note.C)
	if err != nil {
		panic(err)
	}

	fmt.Println(note1.IsEqualByName(note2), note3.IsEqualByName(note4), note1.IsEqualByName(note3), note1.IsEqualByName(note5))
	// Output: true true false true
}

// Create a note and set octave.
func ExampleNote_SetOctave() {
	// create octave #3
	octave, err := octave.NewByNumber(octave.Number4)
	if err != nil {
		panic(err)
	}

	// create a new note without any octave
	note := note.C.MustNewNote()

	// set the octave to the note
	note.SetOctave(octave)

	fmt.Println(note.Octave().Name())
	// Output: FirstOctave
}

// Setting and Getting relative duration.
func ExampleNote_SetDurationRel() {
	// half note duration
	duration := duration.NewRelative(duration.NameHalf)

	// creating note and setting duration
	note := note.MustNewNoteWithOctave(note.C, octave.Number3).SetDurationRel(duration)

	fmt.Println(note.DurationRel().Name())
	// Output: Half
}

// Getting note's time.Duration from note's duration.
func ExampleNote_GetTimeDuration() {
	// musical settings
	trackSettings := track.Settings{
		BPM:           uint64(80),
		Unit:          *fraction.New(1, 2),
		TimeSignature: *fraction.New(4, 4),
	}

	// half note duration
	duration := duration.NewRelative(duration.NameHalf)

	// creating note and setting duration
	note := note.MustNewNoteWithOctave(note.C, octave.Number3).SetDurationRel(duration)

	fmt.Println(note.GetTimeDuration(trackSettings.GetAmountOfBars()))
	// Output: 750ms
}

// Setting and Getting custom duration.
func ExampleNote_SetDurationAbs() {
	// creating note and setting custom duration
	note := note.MustNewNoteWithOctave(note.C, octave.Number3).SetDurationAbs(time.Second)

	fmt.Println(note.DurationAbs())
	// Output: 1s
}
