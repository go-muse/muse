package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Create and compare notes and print their names.
// Creating a note by its name, taken from the Muse library, guarantees panic-free execution of this method.
func ExampleNoteName_NewNote() {
	// Creating a new note from specified note name.
	note1 := muse.C.NewNote()

	// Creating another note from specified note name.
	note2 := muse.AFLAT.NewNote()

	fmt.Println(note1.Name(), note2.Name())
	// Output: C Ab
}
