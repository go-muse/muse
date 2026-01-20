package note_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

// Create notes and compare them by name.
func ExampleNote() {
	// Create note with octave using constructor
	note1, err := note.NewWithOctave(note.C, octave.NumberDefault)
	if err != nil {
		panic(err)
	}

	// Copy the note
	note2 := note1.Copy()

	// Create note from string
	note3, err := note.NewFromString("C#")
	if err != nil {
		panic(err)
	}

	// Create note using MustNewWithOctave
	note4 := note.MustNewWithOctave(note.CSHARP, octave.NumberDefault)

	// Create note from predefined name
	note5 := note.C.NewNote()

	fmt.Println(note1.EqualByName(note2), note3.EqualByName(note4), note1.EqualByName(note3), note1.EqualByName(note5))
	// Output: true true false true
}

// Create a note and set its octave.
func ExampleNote_SetOctave() {
	// Create octave #4 (First Octave)
	oct, err := octave.NewByNumber(octave.Number4)
	if err != nil {
		panic(err)
	}

	// Create a new note without any octave
	n := note.C.NewNote()

	// Set the octave to the note
	n = n.WithOctave(oct)

	fmt.Println(n.Octave().Name())
	// Output: FirstOctave
}

// Setting and getting relative duration.
func ExampleNote_SetValue() {
	// Half note duration
	dur := duration.NewRelative(duration.NameHalf)

	// Create note and set duration
	n := note.MustNewWithOctave(note.C, octave.Number3).WithValue(dur)

	fmt.Println(n.Value().Name())
	// Output: Half
}

// Getting note's relative value name.
func ExampleNote_Value() {
	// Half note duration
	dur := duration.NewRelative(duration.NameHalf)

	// Create note and set duration
	n := note.MustNewWithOctave(note.C, octave.Number3).WithValue(dur)

	fmt.Println(n.Value().Name())
	// Output: Half
}

// Setting and getting absolute (custom) duration.
func ExampleNote_SetDuration() {
	// Create note and set custom duration
	n := note.MustNewWithOctave(note.C, octave.Number3).WithDuration(time.Second)

	fmt.Println(n.Duration())
	// Output: 1s
}

// Create notes from strings.
func ExampleNewFromString() {
	// Parse various note strings
	n1, _ := note.NewFromString("C")
	n2, _ := note.NewFromString("D#")
	n3, _ := note.NewFromString("Eb")

	fmt.Println(n1.String(), n2.String(), n3.String())
	// Output: C D# Eb
}

// Create notes with octave.
func ExampleNewWithOctave() {
	// Create notes at different octaves
	c4, _ := note.NewWithOctave(note.C, 4) // Middle C
	a4, _ := note.NewWithOctave(note.A, 4) // A440

	fmt.Println(c4.String(), c4.Octave().Number())
	fmt.Println(a4.String(), a4.Octave().Number())
	// Output:
	// C 4
	// A 4
}

// Altering notes.
func ExampleNote_AlterUp() {
	// Start with C
	n := note.C.NewNote()

	// Alter up to C#
	nSharp := n.AlterUp()

	// Alter up again to C##
	nDoubleSharp := nSharp.AlterUp()

	fmt.Println(n.String(), "->", nSharp.String(), "->", nDoubleSharp.String())
	// Output: C -> C# -> C##
}

// Comparing notes by octave.
func ExampleNote_EqualByOctave() {
	c4 := note.MustNewWithOctave(note.C, 4)
	d4 := note.MustNewWithOctave(note.D, 4)
	c5 := note.MustNewWithOctave(note.C, 5)

	// Same octave, different notes
	fmt.Println("C4 and D4 same octave:", c4.EqualByOctave(d4))

	// Different octaves
	fmt.Println("C4 and C5 same octave:", c4.EqualByOctave(c5))

	// Output:
	// C4 and D4 same octave: true
	// C4 and C5 same octave: false
}

// Getting the base name without accidentals.
func ExampleNote_BaseName() {
	notes := []note.Note{
		note.C.NewNote(),
		note.CSHARP.NewNote(),
		note.CFLAT.NewNote(),
		note.CSHARP2.NewNote(),
	}

	for _, n := range notes {
		fmt.Printf("%s -> %s\n", n.String(), n.BaseName())
	}
	// Output:
	// C -> C
	// C# -> C
	// Cb -> C
	// C## -> C
}

// Getting the alteration shift.
func ExampleNote_AlterationShift() {
	notes := []note.Note{
		note.CFLAT2.NewNote(),
		note.CFLAT.NewNote(),
		note.C.NewNote(),
		note.CSHARP.NewNote(),
		note.CSHARP2.NewNote(),
	}

	for _, n := range notes {
		fmt.Printf("%s: %d\n", n.String(), n.AlterationShift())
	}
	// Output:
	// Cbb: -2
	// Cb: -1
	// C: 0
	// C#: 1
	// C##: 2
}

// Get frequency of a note.
func ExampleNote_Frequency440() {
	// A4 is the reference pitch (440 Hz)
	a4 := note.MustNewWithOctave(note.A, 4)

	// Middle C
	c4 := note.MustNewWithOctave(note.C, 4)

	fmt.Printf("A4: %.2f Hz\n", a4.Frequency440())
	fmt.Printf("C4: %.2f Hz\n", c4.Frequency440())
	// Output:
	// A4: 440.00 Hz
	// C4: 261.63 Hz
}
