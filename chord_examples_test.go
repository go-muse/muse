package muse_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse"
)

// Creating a new chord.
func ExampleNewChord() {
	// notes with different durations
	note1 := muse.MustNewNoteWithOctave(muse.E, 4).SetDurationRel(muse.NewDurationRel(muse.DurationNameWhole))
	note2 := muse.MustNewNoteWithOctave(muse.G, 4).SetDurationRel(muse.NewDurationRel(muse.DurationNameHalf))
	note3 := muse.MustNewNoteWithOctave(muse.B, 4).SetDurationRel(muse.NewDurationRel(muse.DurationNameQuarter))
	note4 := muse.MustNewNoteWithOctave(muse.D, 5).SetDurationRel(muse.NewDurationRel(muse.DurationNameSixteenth))

	// duration for the chord
	duration := muse.NewDurationRel(muse.DurationNameHalf)

	// notes will be added in the new chord with the specified duration. It will be the same for all the notes in the chord.
	chord := muse.NewChord(*note1, *note2, *note3, *note4).SetDurationRel(duration)

	fmt.Println(chord.String())
	// Output: notes: [E G B D], duration name: Half, custom duration: 0s
}

// Creating a new empty chord and adding notes.
func ExampleNewChordEmpty() {
	duration := muse.NewDurationRel(muse.DurationNameHalf)

	chord := muse.NewChordEmpty()
	chord.SetDurationRel(duration)

	note1 := muse.MustNewNoteWithOctave(muse.E, 4)
	note2 := muse.MustNewNoteWithOctave(muse.G, 4)
	note3 := muse.MustNewNoteWithOctave(muse.B, 4)
	note4 := muse.MustNewNoteWithOctave(muse.D, 5)

	chord.AddNotes(*note1, *note2, *note3, *note4)

	fmt.Println(chord.String())
	// Output: notes: [E G B D], duration name: Half, custom duration: 0s
}

// Adding notes to the chord. Identical notes (equal by Name and Octave) will not be added.
func ExampleChord_AddNotes() {
	duration := muse.NewDurationRel(muse.DurationNameHalf)

	chord := muse.NewChordEmpty()
	chord.SetDurationRel(duration)

	note1 := muse.MustNewNoteWithOctave(muse.E, 4)
	note2 := muse.MustNewNoteWithOctave(muse.G, 4)
	note3 := muse.MustNewNoteWithOctave(muse.B, 4)
	note4 := muse.MustNewNoteWithOctave(muse.D, 5)
	note5 := muse.MustNewNoteWithOctave(muse.D, 5) // existing note
	note6 := muse.MustNewNoteWithOctave(muse.E, 4) // existing note

	chord.AddNotes(*note1, *note2, *note3, *note4, *note5, *note6)

	existingNote := muse.MustNewNoteWithOctave(muse.E, 4)
	chord.AddNote(*existingNote)

	fmt.Println(chord.String())
	// Output: notes: [E G B D], duration name: Half, custom duration: 0s
}

// Setting relative duration to the chord.
func ExampleChord_SetDurationRel() {
	chord := muse.NewChordEmpty()

	note1 := muse.MustNewNoteWithOctave(muse.E, 4)
	note2 := muse.MustNewNoteWithOctave(muse.G, 4)
	note3 := muse.MustNewNoteWithOctave(muse.B, 4)
	note4 := muse.MustNewNoteWithOctave(muse.D, 5)

	chord.AddNotes(*note1, *note2, *note3, *note4)

	duration := muse.NewDurationRel(muse.DurationNameHalf)
	chord.SetDurationRel(duration)

	var resultStr string
	for _, chordNote := range chord.Notes() {
		resultStr += fmt.Sprintf("note: %s duration: %s\n", chordNote.Name(), chordNote.DurationRel().Name())
	}

	fmt.Println(resultStr)
	// Output: note: E duration: Half
	// note: G duration: Half
	// note: B duration: Half
	// note: D duration: Half
}

// Setting custom absolute duration to the chord.
func ExampleChord_SetDurationAbs() {
	chord := muse.NewChordEmpty()

	note1 := muse.MustNewNoteWithOctave(muse.E, 4)
	note2 := muse.MustNewNoteWithOctave(muse.G, 4)
	note3 := muse.MustNewNoteWithOctave(muse.B, 4)
	note4 := muse.MustNewNoteWithOctave(muse.D, 5)

	chord.AddNotes(*note1, *note2, *note3, *note4)

	chord.SetDurationAbs(time.Second)

	var resultStr string
	for _, chordNote := range chord.Notes() {
		resultStr += fmt.Sprintf("note: %s custom duration: %s\n", chordNote.Name(), chordNote.DurationAbs())
	}

	fmt.Println(resultStr)
	// Output: note: E custom duration: 1s
	// note: G custom duration: 1s
	// note: B custom duration: 1s
	// note: D custom duration: 1s
}
