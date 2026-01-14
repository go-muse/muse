package note_test

import (
	"fmt"

	"github.com/go-muse/muse/note"
)

// Creating note names from letters.
func ExampleLetter_Natural() {
	// Create a natural note name from a letter
	name := note.LetterC.Natural()
	fmt.Println(name.String())
	// Output: C
}

// Creating sharp note names from letters.
func ExampleLetter_Sharp() {
	// Create a sharp note name from a letter
	name := note.LetterF.Sharp()
	fmt.Println(name.String())
	// Output: F#
}

// Creating flat note names from letters.
func ExampleLetter_Flat() {
	// Create a flat note name from a letter
	name := note.LetterB.Flat()
	fmt.Println(name.String())
	// Output: Bb
}

// Creating double-sharp note names from letters.
func ExampleLetter_DoubleSharp() {
	// Create a double-sharp note name from a letter
	name := note.LetterG.DoubleSharp()
	fmt.Println(name.String())
	// Output: G##
}

// Creating double-flat note names from letters.
func ExampleLetter_DoubleFlat() {
	// Create a double-flat note name from a letter
	name := note.LetterA.DoubleFlat()
	fmt.Println(name.String())
	// Output: Abb
}

// Creating notes directly from predefined note names.
func ExampleName_NewNote() {
	// Creating notes from predefined note names
	noteC := note.C.NewNote()
	noteAFlat := note.AFLAT.NewNote()

	fmt.Println(noteC.Name(), noteAFlat.Name())
	// Output: C Ab
}

// Parsing note names from strings.
func ExampleNewNameFromString() {
	// Parse various note name strings
	name1, _ := note.NewNameFromString("C")
	name2, _ := note.NewNameFromString("F#")
	name3, _ := note.NewNameFromString("Bb")
	name4, _ := note.NewNameFromString("E##")

	fmt.Println(name1, name2, name3, name4)
	// Output: C F# Bb E##
}

// Altering note names.
func ExampleName_AlterUp() {
	// Alter a note name up (add one sharp / remove one flat)
	c := note.C
	cSharp := c.AlterUp()
	cDoubleSharp := cSharp.AlterUp()

	fmt.Println(c, "->", cSharp, "->", cDoubleSharp)
	// Output: C -> C# -> C##
}

// Altering note names down.
func ExampleName_AlterDown() {
	// Alter a note name down (add one flat / remove one sharp)
	c := note.C
	cFlat := c.AlterDown()
	cDoubleFlat := cFlat.AlterDown()

	fmt.Println(c, "->", cFlat, "->", cDoubleFlat)
	// Output: C -> Cb -> Cbb
}

// Getting the base name (letter without accidentals).
func ExampleName_BaseName() {
	names := []note.Name{note.C, note.CSHARP, note.CFLAT, note.CSHARP2}

	for _, n := range names {
		fmt.Printf("%s -> %s\n", n.String(), n.BaseName())
	}
	// Output:
	// C -> C
	// C# -> C
	// Cb -> C
	// C## -> C
}

// Comparing note name spellings.
func ExampleName_EqualSpelling() {
	// Same note names are equal
	fmt.Println("C == C:", note.C.EqualSpelling(note.C))

	// Different accidentals are not equal (even if enharmonic)
	fmt.Println("C# == Db:", note.CSHARP.EqualSpelling(note.DFLAT))

	// Different letters are not equal
	fmt.Println("C == D:", note.C.EqualSpelling(note.D))

	// Output:
	// C == C: true
	// C# == Db: false
	// C == D: false
}
