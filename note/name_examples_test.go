package note_test

import (
	"fmt"

	"github.com/go-muse/muse/note"
)

// Create and compare notes and print their names.
// Creating a note by its name, taken from the Muse library, guarantees panic-free execution of this method.
func ExampleName_NewNote() {
	// Creating a new note from specified note name.
	note1 := note.C.MustNewNote()

	// Creating another note from specified note name.
	note2 := note.AFLAT.MustNewNote()

	fmt.Println(note1.Name(), note2.Name())
	// Output: C Ab
}
