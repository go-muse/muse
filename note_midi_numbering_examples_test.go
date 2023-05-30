package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Getting MIDI number of the note.
func ExampleNote_MIDINumber() {
	// creating a note with octave
	note := muse.MustNewNote(muse.GSHARP, muse.OctaveNumber2)

	fmt.Println(note.MIDINumber())
	// Output: 44
}
