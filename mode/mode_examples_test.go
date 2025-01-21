package mode_test

import (
	"fmt"

	"github.com/go-muse/muse/mode"
	"github.com/go-muse/muse/note"
)

// Creating mode with textual mode name and note name will validate them during mode building.
func ExampleMakeNewMode() {
	mode, err := mode.MakeNewMode(mode.Name("UltraLocrian"), note.Name("C"))
	if err != nil {
		panic(err)
	}

	ultraLocrianScale := mode.GenerateScale(false)

	fmt.Println(ultraLocrianScale)
	// Output: [C Db Eb Fb Gb Ab Bbb]
}

// Creating mode with a mode's name from muse and note's name from muse name guarantees the absence of errors,
// so we can use MustMakeMode function to avoid error checking.
func ExampleMustMakeNewMode() {
	mode := mode.MustMakeNewMode(mode.NameSuperLocrian, note.C)

	superLocrianScale := mode.GenerateScale(false)

	fmt.Println(superLocrianScale)
	// Output: [C Db Eb Fb Gb Ab Bb]
}

// You can create your custom mode with own mode template.
// Note, that amount of halftone in the template must be equal 12 (octave).
func ExampleMakeNewCustomMode() {
	customModeTemplate := mode.Template{1, 2, 2, 1, 2, 2, 2}

	customMode, err := mode.MakeNewCustomMode(customModeTemplate, "C", "Custom mode name")
	if err != nil {
		panic(err)
	}

	fmt.Println(customMode.GenerateScale(false))
	// Output: [C Db Eb F Gb Ab Bb]
}

// Getting from a mode it's notes from degrees, sorted by absolute modal positions.
func ExampleMode_SortByAbsoluteModalPositions() {
	firstNote := note.C
	m := mode.MustMakeNewMode(mode.NameDorian, firstNote)

	dorianScale := m.GenerateScale(false)

	m.SortByAbsoluteModalPositions(false) // argument 'false' means descending order of modal positions
	notes := m.IterateOneRound(false).GetAllNotes()

	fmt.Printf("The Dorian scale with tonal centre '%s' is: %v\nSorted by Modal Positions: %v\n", firstNote, dorianScale, notes)
	// Output: The Dorian scale with tonal centre 'C' is: [C D Eb F G A Bb]
	// Sorted by Modal Positions: [A D G C F Bb Eb]
}

// There is a possibility to iterate one cycle through the degrees of the mode.
// The iteration starts with the first note of the mode (tonic).
// You can specify the direction of iteration.
// The function returns the iterator that you can use to iterate through mode's degrees.
func ExampleMode_IterateOneRound() {
	m := mode.MustMakeNewMode(mode.NameAeolian, note.A)

	iteratorForward := m.IterateOneRound(false)

	for degree := range iteratorForward {
		fmt.Printf("Degree Number: %d, Half tones from prime: %d, Note: %s\n", degree.Number(), degree.HalfTonesFromPrime(), degree.Note().Name())
	}
	// Output: Degree Number: 1, Half tones from prime: 0, Note: A
	// Degree Number: 2, Half tones from prime: 2, Note: B
	// Degree Number: 3, Half tones from prime: 3, Note: C
	// Degree Number: 4, Half tones from prime: 5, Note: D
	// Degree Number: 5, Half tones from prime: 7, Note: E
	// Degree Number: 6, Half tones from prime: 8, Note: F
	// Degree Number: 7, Half tones from prime: 10, Note: G
}
