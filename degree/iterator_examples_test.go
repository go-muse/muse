package degree

import (
	"fmt"

	"github.com/go-muse/muse/note"
)

// There is a possibility to iterate one cycle through the chained degrees.
// The iteration starts with the first degree (tonic).
// You can specify the direction of iteration.
// The function returns the iterator with functionality.
// GetAll returns mode's degrees as a slice.
func ExampleIterator_GetAllDegrees() {
	// In real life you can just build a mode by one line from mode package:
	// mode := mode.MustMakeNewMode(mode.NameAeolian, note.A)
	//
	// or build degrees manually:
	deg7 := New(7, 10, nil, nil, note.G.MustMakeNote(), nil, nil)
	deg6 := New(6, 8, nil, deg7, note.F.MustMakeNote(), nil, nil)
	deg5 := New(5, 7, nil, deg6, note.E.MustMakeNote(), nil, nil)
	deg4 := New(4, 5, nil, deg5, note.D.MustMakeNote(), nil, nil)
	deg3 := New(3, 3, nil, deg4, note.C.MustMakeNote(), nil, nil)
	deg2 := New(2, 2, nil, deg3, note.B.MustMakeNote(), nil, nil)
	deg1 := New(1, 0, nil, deg2, note.A.MustMakeNote(), nil, nil)

	// it works even the chain is cycled
	deg7.SetNext(deg1)

	// iteration in forward direction, by degree.fields
	iteratorForward := deg1.IterateOneRound(false)

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

// There is a possibility to iterate one cycle through the degrees.
// The iteration starts with the first degree (tonic).
// You can specify the direction of iteration.
// The function returns the iterator with functionality.
// GetAllNotes returns degrees' notes as a slice.
func ExampleIterator_GetAllNotes() {
	// In real life you can just build a mode by one line from mode package:
	// mode := mode.MustMakeNewMode(mode.NameAeolian, note.A)
	//
	// or build degrees manually:
	deg7 := New(7, 10, nil, nil, note.G.MustMakeNote(), nil, nil)
	deg6 := New(6, 8, nil, deg7, note.F.MustMakeNote(), nil, nil)
	deg5 := New(5, 7, nil, deg6, note.E.MustMakeNote(), nil, nil)
	deg4 := New(4, 5, nil, deg5, note.D.MustMakeNote(), nil, nil)
	deg3 := New(3, 3, nil, deg4, note.C.MustMakeNote(), nil, nil)
	deg2 := New(2, 2, nil, deg3, note.B.MustMakeNote(), nil, nil)
	deg1 := New(1, 0, deg7, deg2, note.A.MustMakeNote(), nil, nil)

	deg7.SetNext(deg1).SetPrevious(deg6)
	deg6.SetPrevious(deg5)
	deg5.SetPrevious(deg4)
	deg4.SetPrevious(deg3)
	deg3.SetPrevious(deg2)
	deg2.SetPrevious(deg1)

	// you can specify the direction of iteration from any degree
	iteratorForward := deg3.IterateOneRound(false)
	iteratorBackward := deg3.IterateOneRound(true)

	fmt.Printf("%+v\n%+v", iteratorForward.GetAllNotes(), iteratorBackward.GetAllNotes())
	// Output: [C D E F G A B]
	// [C B A G F E D]
}
