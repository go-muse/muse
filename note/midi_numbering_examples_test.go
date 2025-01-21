package note_test

import (
	"fmt"

	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

// Getting MIDI number of the note.
func ExampleNote_MIDINumber() {
	// creating a note with octave
	n := note.MustNewNoteWithOctave(note.GSHARP, octave.Number2)

	fmt.Println(n.MIDINumber())
	// Output: 44
}

// Creating a note from midi number.
func ExampleNewNoteFromMIDINumber() {
	n, err := note.NewNoteFromMIDINumber(76)
	if err != nil {
		panic(err)
	}

	fmt.Println(n.Name(), n.Octave().Name())
	// Output: E SecondOctave
}
