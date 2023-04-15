package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Creating mode with textual mode name and note name will validate them dirung mode building.
func ExampleMakeNewMode() {
	mode, err := muse.MakeNewMode(muse.ModeName("UltraLocrian"), muse.NoteName("C"))
	if err != nil {
		panic(err)
	}

	ultraLocrianScale := mode.GenerateScale(false)
	fmt.Println(ultraLocrianScale)
	// Output: [{C} {Db} {Eb} {Fb} {Gb} {Ab} {Bbb}]
}

// Creating mode with a mode's name from muse and note's name from muse name guarantees the absence of errors,
// so we can use MustMakeMode function to avoid error checking.
func ExampleMustMakeNewMode() {
	mode := muse.MustMakeNewMode(muse.ModeNameSuperLocrian, muse.C)
	superLocrianScale := mode.GenerateScale(false)

	fmt.Println(superLocrianScale)
	// Output: [{C} {Db} {Eb} {Fb} {Gb} {Ab} {Bb}]
}

// You can create your custom mode with own mode template.
// Note, that amount of half tones in the template must be equal 12 (octave).
func ExampleMakeNewCustomMode() {
	customModeTemplate := muse.ModeTemplate{1, 2, 2, 1, 2, 2, 2}
	customMode, err := muse.MakeNewCustomMode(customModeTemplate, "C", "Custom mode name")
	if err != nil {
		panic(err)
	}

	fmt.Println(customMode.GenerateScale(false))
	// Output: [{C} {Db} {Eb} {F} {Gb} {Ab} {Bb}]
}

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
