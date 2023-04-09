package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Create and compare notes and print their names.
//
//nolint:dupword
func ExampleNote() {
	note1, err := muse.NewNote(muse.C)
	if err != nil {
		panic(err)
	}

	note2 := note1.Copy()

	note3, err := muse.NewNoteFromString("C#")
	if err != nil {
		panic(err)
	}

	note4 := muse.MustNewNote(muse.CSHARP)

	fmt.Println(note1.IsEqualByName(note2), note3.IsEqualByName(note4), note1.IsEqualByName(note3))
	// Output: true true false
}
