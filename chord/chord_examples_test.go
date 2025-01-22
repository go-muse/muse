package chord_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse/chord"
	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
)

// Creating a new chord.
func ExampleNewChord() {
	// notes with different durations
	note1 := note.MustNewNoteWithOctave(note.E, 4).SetValue(duration.NewRelative(duration.NameWhole))
	note2 := note.MustNewNoteWithOctave(note.G, 4).SetValue(duration.NewRelative(duration.NameHalf))
	note3 := note.MustNewNoteWithOctave(note.B, 4).SetValue(duration.NewRelative(duration.NameQuarter))
	note4 := note.MustNewNoteWithOctave(note.D, 5).SetValue(duration.NewRelative(duration.NameSixteenth))

	// duration for the chord
	duration := duration.NewRelative(duration.NameHalf)

	// notes will be added in the new chord with the specified duration. It will be the same for all the notes in the chord.
	chord := chord.NewChord(note1, note2, note3, note4).SetValue(duration)

	fmt.Println(chord.String())
	// Output: notes: [E G B D], duration name: Half, custom duration: 0s
}

// Creating a new empty chord and adding notes.
func ExampleNewChordEmpty() {
	duration := duration.NewRelative(duration.NameHalf)

	chord := chord.NewChordEmpty()
	chord.SetValue(duration)

	note1 := note.MustNewNoteWithOctave(note.E, 4)
	note2 := note.MustNewNoteWithOctave(note.G, 4)
	note3 := note.MustNewNoteWithOctave(note.B, 4)
	note4 := note.MustNewNoteWithOctave(note.D, 5)

	chord.AddNotes(note1, note2, note3, note4)

	fmt.Println(chord.String())
	// Output: notes: [E G B D], duration name: Half, custom duration: 0s
}

// Adding notes to the chord. Identical notes (equal by Name and Octave) will not be added.
func ExampleChord_AddNotes() {
	duration := duration.NewRelative(duration.NameHalf)

	chord := chord.NewChordEmpty()
	chord.SetValue(duration)

	note1 := note.MustNewNoteWithOctave(note.E, 4)
	note2 := note.MustNewNoteWithOctave(note.G, 4)
	note3 := note.MustNewNoteWithOctave(note.B, 4)
	note4 := note.MustNewNoteWithOctave(note.D, 5)
	note5 := note.MustNewNoteWithOctave(note.D, 5) // existing note
	note6 := note.MustNewNoteWithOctave(note.E, 4) // existing note

	chord.AddNotes(note1, note2, note3, note4, note5, note6)

	existingNote := note.MustNewNoteWithOctave(note.E, 4)
	chord.AddNote(existingNote)

	fmt.Println(chord.String())
	// Output: notes: [E G B D], duration name: Half, custom duration: 0s
}

// Setting relative duration to the chord.
func ExampleChord_SetValue() {
	chord := chord.NewChordEmpty()

	note1 := note.MustNewNoteWithOctave(note.E, 4)
	note2 := note.MustNewNoteWithOctave(note.G, 4)
	note3 := note.MustNewNoteWithOctave(note.B, 4)
	note4 := note.MustNewNoteWithOctave(note.D, 5)

	chord.AddNotes(note1, note2, note3, note4)

	duration := duration.NewRelative(duration.NameHalf)
	chord.SetValue(duration)

	var resultStr string
	for _, chordNote := range chord.Notes() {
		resultStr += fmt.Sprintf("note: %s duration: %s\n", chordNote.Name(), chordNote.Value().Name())
	}

	fmt.Println(resultStr)
	// Output: note: E duration: Half
	// note: G duration: Half
	// note: B duration: Half
	// note: D duration: Half
}

// Setting custom absolute duration to the chord.
func ExampleChord_SetDuration() {
	chord := chord.NewChordEmpty()

	note1 := note.MustNewNoteWithOctave(note.E, 4)
	note2 := note.MustNewNoteWithOctave(note.G, 4)
	note3 := note.MustNewNoteWithOctave(note.B, 4)
	note4 := note.MustNewNoteWithOctave(note.D, 5)

	chord.AddNotes(note1, note2, note3, note4)

	chord.SetDuration(time.Second)

	var resultStr string
	for _, chordNote := range chord.Notes() {
		resultStr += fmt.Sprintf("note: %s custom duration: %s\n", chordNote.Name(), chordNote.Duration())
	}

	fmt.Println(resultStr)
	// Output: note: E custom duration: 1s
	// note: G custom duration: 1s
	// note: B custom duration: 1s
	// note: D custom duration: 1s
}
