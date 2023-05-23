package muse_test

import (
	"fmt"

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
