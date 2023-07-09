package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Creating an octave.
func ExampleOctave() {
	// create octave #3
	octave, err := muse.NewOctave(muse.OctaveNumber3)
	if err != nil {
		panic(err)
	}

	fmt.Println(octave.Number(), octave.Name())
	// Output: 3 SmallOctave
}

// Create an octave and assign it to a note by passing the note as an argument.
func ExampleOctave_SetToNote() {
	// create octave #3
	octave, err := muse.NewOctave(muse.OctaveNumber4)
	if err != nil {
		panic(err)
	}

	// create a new note without any octave
	note := muse.C.NewNote()

	// set the octave to the note
	octave.SetToNote(note)

	fmt.Println(note.Octave().Name())
	// Output: FirstOctave
}
