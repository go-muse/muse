package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Getting from a mode it's notes from degrees, sorted by absolute modal positions.
func ExampleMode_SortByAbsoluteModalPositions() {
	mode := muse.MustMakeNewMode(muse.ModeNameDorian, muse.B)
	dorianScale := mode.GenerateScale(false)
	mode.SortByAbsoluteModalPositions(false) // argument 'false' means descending order of modal positions
	notes := mode.IterateOneRound(false).GetAllNotes()
	fmt.Printf("The Dorian scale with tonal centre 'B' is: %v\nSorted by Modal Positions: %v\n", dorianScale, notes)
	// Output: The Dorian scale with tonal centre 'B' is: [{B} {C#} {D} {E} {F#} {G#} {A}]
	// Sorted by Modal Positions: [{G#} {C#} {F#} {B} {E} {A} {D}]
}
