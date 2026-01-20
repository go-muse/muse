package degree

import (
	"fmt"

	"github.com/go-muse/muse/note"
)

// There is a possibility to iterate one cycle through the chained degree nodes.
// The iteration starts with the first node (tonic).
// You can specify the direction of iteration.
// GetAllDegrees returns mode's degree nodes as a slice.
func ExampleGetAllDegrees() {
	// In real life you can just build a mode by one line from mode package:
	// mode := mode.MustMakeNewMode(mode.NameAeolian, note.A)
	//
	// or build degree nodes manually:
	deg7 := New(7, 10, nil, nil, note.G.NewNote(), nil, ModalPosition{})
	deg6 := New(6, 8, nil, deg7, note.F.NewNote(), nil, ModalPosition{})
	deg5 := New(5, 7, nil, deg6, note.E.NewNote(), nil, ModalPosition{})
	deg4 := New(4, 5, nil, deg5, note.D.NewNote(), nil, ModalPosition{})
	deg3 := New(3, 3, nil, deg4, note.C.NewNote(), nil, ModalPosition{})
	deg2 := New(2, 2, nil, deg3, note.B.NewNote(), nil, ModalPosition{})
	deg1 := New(1, 0, nil, deg2, note.A.NewNote(), nil, ModalPosition{})

	// it works even the chain is cycled
	deg7.SetNext(deg1)

	// iteration in forward direction
	for _, node := range deg1.IterateOneRound(false).GetAllDegrees() {
		fmt.Printf("Degree Number: %d, Half tones from prime: %d, Note: %s\n", node.Number(), node.HalfTonesFromPrime(), node.Note().Name())
	}
	// Output: Degree Number: 1, Half tones from prime: 0, Note: A
	// Degree Number: 2, Half tones from prime: 2, Note: B
	// Degree Number: 3, Half tones from prime: 3, Note: C
	// Degree Number: 4, Half tones from prime: 5, Note: D
	// Degree Number: 5, Half tones from prime: 7, Note: E
	// Degree Number: 6, Half tones from prime: 8, Note: F
	// Degree Number: 7, Half tones from prime: 10, Note: G
}

// There is a possibility to iterate one cycle through the degree nodes.
// The iteration starts with the first node (tonic).
// You can specify the direction of iteration.
// GetAllNotes returns nodes' notes as a slice.
func ExampleGetAllNotes() {
	// In real life you can just build a mode by one line from mode package:
	// mode := mode.MustMakeNewMode(mode.NameAeolian, note.A)
	//
	// or build degree nodes manually:
	deg7 := New(7, 10, nil, nil, note.G.NewNote(), nil, ModalPosition{})
	deg6 := New(6, 8, nil, deg7, note.F.NewNote(), nil, ModalPosition{})
	deg5 := New(5, 7, nil, deg6, note.E.NewNote(), nil, ModalPosition{})
	deg4 := New(4, 5, nil, deg5, note.D.NewNote(), nil, ModalPosition{})
	deg3 := New(3, 3, nil, deg4, note.C.NewNote(), nil, ModalPosition{})
	deg2 := New(2, 2, nil, deg3, note.B.NewNote(), nil, ModalPosition{})
	deg1 := New(1, 0, deg7, deg2, note.A.NewNote(), nil, ModalPosition{})

	deg7.SetNext(deg1).SetPrevious(deg6)
	deg6.SetPrevious(deg5)
	deg5.SetPrevious(deg4)
	deg4.SetPrevious(deg3)
	deg3.SetPrevious(deg2)
	deg2.SetPrevious(deg1)

	// you can specify the direction of iteration from any node
	notesForward := deg3.IterateOneRound(false).GetAllNotes()
	notesBackward := deg3.IterateOneRound(true).GetAllNotes()

	fmt.Printf("%+v\n%+v", notesForward, notesBackward)
	// Output: [C D E F G A B]
	// [C B A G F E D]
}
