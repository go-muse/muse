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

// There is a possibility to iterate one cycle through the degrees of the mode.
// The iteration starts with the first note of the mode (tonic).
// You can specify the direction of iteration.
// The function returns the iterator that you can use to iterate through mode's degrees.
func ExampleMode_IterateOneRound() {
	mode := muse.MustMakeNewMode(muse.ModeNameAeolian, muse.A)

	iteratorForward := mode.IterateOneRound(false)

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

// There is a possibility to iterate one cycle through the degrees of the mode.
// The iteration starts with the first note of the mode (tonic).
// You can specify the direction of iteration.
// The function returns the iterator with functionality.
// GetAllNotes returns degrees's notes as a slice.
func ExampleDegreesIterator_GetAllNotes() {
	mode := muse.MustMakeNewMode(muse.ModeNameAeolian, muse.A)

	iteratorForward := mode.IterateOneRound(false)
	iteratorBackward := mode.IterateOneRound(true)

	fmt.Printf("%+v\n%+v", iteratorForward.GetAllNotes(), iteratorBackward.GetAllNotes())
	// Output: [{name:A} {name:B} {name:C} {name:D} {name:E} {name:F} {name:G}]
	// [{name:A} {name:G} {name:F} {name:E} {name:D} {name:C} {name:B}]
}

// There is a possibility to iterate one cycle through the degrees of the mode.
// The iteration starts with the first note of the mode (tonic).
// You can specify the direction of iteration.
// The function returns the iterator with functionality.
// GetAll returns mode's degrees as a slice.
func ExampleDegreesIterator_GetAllDegrees() {
	mode := muse.MustMakeNewMode(muse.ModeNameAeolian, muse.A)

	iteratorForward := mode.IterateOneRound(false)

	for _, degree := range iteratorForward.GetAllDegrees() {
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
