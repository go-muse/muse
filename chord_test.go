package muse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewChord(t *testing.T) {
	t.Run("NewChord: creating a new chord with different notes", func(t *testing.T) {
		notes := Notes{
			{C, nil, nil},
			{E, nil, nil},
			{G, nil, nil},
		}

		duration := &Duration{4, relativeDuration{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}}}

		chord := NewChord(*duration, notes...)

		assert.Equal(t, len(notes), len(chord.notes), "expected %v notes, got %v", len(notes), len(chord.notes))
		assert.Equal(t, duration, chord.duration, "expected chord duration to be %v, got %v", duration, chord.duration)

		for i, chordNote := range chord.notes {
			assert.Equal(t, duration, chordNote.duration, "expected note %v to have duration %v, got %v", i, duration, chordNote.duration)
		}
	})

	t.Run("NewChord: creating a new chord with existing notes", func(t *testing.T) {
		notes := Notes{
			{C, MustNewOctave(4), nil},
			{E, MustNewOctave(4), nil},
			{G, MustNewOctave(4), nil},
			{C, MustNewOctave(4), nil}, // existing note, won't be added
			{E, MustNewOctave(4), nil}, // existing note, won't be added
			{G, MustNewOctave(4), nil}, // existing note, won't be added
		}

		duration := &Duration{4, relativeDuration{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}}}

		chord := NewChord(*duration, notes...)

		differentNotes := 3
		assert.Equal(t, differentNotes, len(chord.notes), "expected %v notes, got %v", differentNotes, len(chord.notes))
		assert.Equal(t, duration, chord.duration, "expected chord duration to be %v, got %v", duration, chord.duration)

		for i, chordNote := range chord.notes {
			assert.Equal(t, duration, chordNote.duration, "expected note %v to have duration %v, got %v", i, duration, chordNote.duration)
		}
	})
}

func TestNewChordEmpty(t *testing.T) {
	chord := NewChordEmpty()
	assert.Equal(t, 0, len(chord.notes), "expected %v notes, got %v", 0, len(chord.notes))
	assert.Nil(t, chord.duration, "expected chord duration to be nil, got %v", chord.duration)
}

func TestChord_AddNote(t *testing.T) {
	t.Run("Chord_AddNote: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes: nil,
			duration: &Duration{
				absoluteDuration: time.Second,
				relativeDuration: relativeDuration{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
			},
		}

		testCases := []struct {
			note *Note
		}{
			{&Note{C, MustNewOctave(1), nil}},
			{&Note{E, MustNewOctave(2), NewDuration(DurationNameHalf)}},
			{&Note{G, MustNewOctave(3), &Duration{
				absoluteDuration: time.Hour,
				relativeDuration: relativeDuration{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}},
			}}},
		}

		for _, testCase := range testCases {
			chord.AddNote(*testCase.note)
		}

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.duration, chordNote.duration, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.duration, chordNote.duration)
			assert.True(t, chordNote.IsEqualByName(testCases[i].note))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNote: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes: nil,
			duration: &Duration{
				absoluteDuration: time.Second,
				relativeDuration: relativeDuration{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
			},
		}

		testCases := []struct {
			note *Note
		}{
			{&Note{C, MustNewOctave(1), nil}},
			{&Note{E, MustNewOctave(2), NewDuration(DurationNameHalf)}},
			{&Note{G, MustNewOctave(3), &Duration{
				absoluteDuration: time.Hour,
				relativeDuration: relativeDuration{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}},
			}}},
		}

		// adding notes
		for _, testCase := range testCases {
			chord.AddNote(*testCase.note)
		}

		// adding existing notes
		for _, testCase := range testCases {
			chord.AddNote(*testCase.note)
		}

		for _, chordNote := range chord.notes {
			assert.Equal(t, chord.duration, chordNote.duration, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.duration, chordNote.duration)
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNote: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		note := MustNewNote(C, 1)
		chord.AddNote(*note)
		assert.Nil(t, chord)
	})
}

func TestChord_AddNotes(t *testing.T) {
	t.Run("Chord_AddNotes: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes: nil,
			duration: &Duration{
				absoluteDuration: time.Second,
				relativeDuration: relativeDuration{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
			},
		}

		testCases := Notes{
			Note{C, MustNewOctave(1), nil},
			Note{E, MustNewOctave(2), NewDuration(DurationNameHalf)},
			Note{G, MustNewOctave(3), &Duration{
				absoluteDuration: time.Hour,
				relativeDuration: relativeDuration{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}},
			}},
		}

		chord.AddNotes(testCases...)

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.duration, chordNote.duration, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.duration, chordNote.duration)
			assert.True(t, chordNote.IsEqualByName(&testCases[i]))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNotes: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes: nil,
			duration: &Duration{
				absoluteDuration: time.Second,
				relativeDuration: relativeDuration{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
			},
		}

		testCases := Notes{
			Note{C, MustNewOctave(1), nil},
			Note{E, MustNewOctave(2), NewDuration(DurationNameHalf)},
			Note{G, MustNewOctave(3), &Duration{
				absoluteDuration: time.Hour,
				relativeDuration: relativeDuration{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}},
			}},
		}

		// adding notes
		chord.AddNotes(testCases...)

		// adding existing notes
		chord.AddNotes(testCases...)

		for _, chordNote := range chord.notes {
			assert.Equal(t, chord.duration, chordNote.duration, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.duration, chordNote.duration)
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNotes: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		chord.AddNotes(*MustNewNote(C, 1), *MustNewNote(D, 2))
		assert.Nil(t, chord)
	})
}

func TestChord_GetNotes(t *testing.T) {
	t.Run("Chord_AddNotes: get notes", func(t *testing.T) {
		duration := NewDuration(DurationNameHalf)
		chord := &Chord{
			notes: Notes{
				Note{C, MustNewOctave(1), duration},
				Note{E, MustNewOctave(2), duration},
				Note{G, MustNewOctave(3), duration},
			},
			duration: duration,
		}

		notes := chord.GetNotes()
		assert.Equal(t, len(chord.notes), len(notes))
		assert.Equal(t, chord.notes, notes)
	})

	t.Run("Chord_AddNotes: getting from nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.GetNotes())
	})
}

func TestChord_SetDuration(t *testing.T) {
	t.Run("Chord_SetDuration: set duration to the chord", func(t *testing.T) {
		duration := NewDuration(DurationNameHalf)
		chord := &Chord{
			notes: Notes{
				Note{C, MustNewOctave(1), duration},
				Note{E, MustNewOctave(2), duration},
				Note{G, MustNewOctave(3), duration},
			},
			duration: duration,
		}

		newDuration := NewDuration(DurationNameWhole)
		chord.SetDuration(*newDuration)
		assert.Equal(t, newDuration, chord.duration, "expected chord duration: %+v, actual chord duration: %+v", newDuration, chord.duration)

		for _, chordNote := range chord.notes {
			assert.Equal(t, newDuration, chordNote.duration, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, newDuration, chordNote.duration)
		}
	})

	t.Run("Chord_SetDuration: set duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetDuration(*NewDuration(DurationNameLong)))
	})
}

func TestChord_GetDuration(t *testing.T) {
	t.Run("Chord_GetDuration: getting duration from the chord", func(t *testing.T) {
		duration := NewDuration(DurationNameHalf)
		chord := NewChord(
			*duration,
			Note{C, MustNewOctave(1), duration},
			Note{E, MustNewOctave(2), duration},
			Note{G, MustNewOctave(3), duration},
		)

		assert.Equal(t, duration, chord.GetDuration(), " expected duration: %+v, actual: %+v", duration, chord.GetDuration())
	})

	t.Run("Chord_GetDuration: getting duration from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.GetDuration())
	})
}

func TestChord_SetAbsoluteDuration(t *testing.T) {
	t.Run("Chord_SetAbsoluteDuration: set custom duration to the chord", func(t *testing.T) {
		duration := NewDuration(DurationNameHalf)
		chord := &Chord{
			notes: Notes{
				Note{C, MustNewOctave(1), duration},
				Note{E, MustNewOctave(2), duration},
				Note{G, MustNewOctave(3), duration},
			},
			duration: duration,
		}

		customDuration := time.Second
		chord.SetAbsoluteDuration(customDuration)
		assert.Equal(t, customDuration, chord.duration.absoluteDuration, "expected chord custom duration: %+v, actual chord custom duration: %+v", customDuration, chord.duration.absoluteDuration)

		for _, chordNote := range chord.notes {
			assert.Equal(t, customDuration, chordNote.duration.absoluteDuration, "note: %s, expected duration: %+v, actual duration: %+v", chordNote.name, customDuration, chordNote.duration.absoluteDuration)
		}
	})

	t.Run("Chord_SetAbsoluteDuration: set custom duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetAbsoluteDuration(time.Second))
	})
}

func TestChord_GetAbsoluteDuration(t *testing.T) {
	t.Run("Chord_GetAbsoluteDuration: getting custom duration from the chord", func(t *testing.T) {
		duration := NewDuration(DurationNameHalf)
		chord := NewChord(
			*duration,
			Note{C, MustNewOctave(1), duration},
			Note{E, MustNewOctave(2), duration},
			Note{G, MustNewOctave(3), duration},
		)

		customDuration := time.Second
		chord.SetAbsoluteDuration(customDuration)

		assert.Equal(t, customDuration, chord.GetAbsoluteDuration(), " expected custom duration: %d, actual custom duration: %d", customDuration, chord.GetAbsoluteDuration())
	})

	t.Run("Chord_GetAbsoluteDuration: getting custom duration from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Zero(t, chord.GetAbsoluteDuration())
	})
}

func TestChord_Empty(t *testing.T) {
	t.Run("Chord_Empty: clearing the chord", func(t *testing.T) {
		chord := NewChordEmpty().AddNotes(
			Note{C, MustNewOctave(1), nil},
			Note{E, MustNewOctave(2), nil},
			Note{G, MustNewOctave(3), nil},
		)

		assert.Equal(t, 0, len(chord.Empty().notes))
	})

	t.Run("Chord_Empty: clearing the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.Empty())
	})
}

func TestChord_RemoveNote(t *testing.T) {
	t.Run("Chord_RemoveNote: remove from the chord", func(t *testing.T) {
		testCases := Notes{
			Note{C, MustNewOctave(1), nil},
			Note{D, MustNewOctave(2), nil},
			Note{E, MustNewOctave(3), nil},
			Note{F, MustNewOctave(4), nil},
			Note{G, MustNewOctave(5), nil},
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		assert.Equal(t, testCases, chord.notes)

		length := len(testCases)
		for i := 0; i < length; i++ {
			chord.RemoveNote(testCases[0])
			testCases = testCases[1:]
			assert.Equal(t, testCases, chord.notes)
		}
	})

	t.Run("Chord_RemoveNote: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.RemoveNote(Note{C, MustNewOctave(1), nil}))
	})
}

func TestChord_RemoveNotes(t *testing.T) {
	t.Run("Chord_RemoveNotes: remove from the chord", func(t *testing.T) {
		testCases := Notes{
			Note{C, MustNewOctave(1), nil},
			Note{D, MustNewOctave(2), nil},
			Note{E, MustNewOctave(3), nil},
			Note{F, MustNewOctave(4), nil},
			Note{G, MustNewOctave(5), nil},
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		assert.Equal(t, testCases, chord.notes)

		chord.RemoveNotes(Notes{testCases[1], testCases[3]})
		testCases = append(testCases[0:1], testCases[2:]...)
		testCases = append(testCases[0:2], testCases[3:]...)
		assert.Equal(t, testCases, chord.notes)
	})

	t.Run("Chord_RemoveNotes: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.RemoveNotes(Notes{}))
	})
}

func TestChord_Exists(t *testing.T) {
	t.Run("Chord_Exists: remove from the chord", func(t *testing.T) {
		testCases := Notes{
			Note{C, MustNewOctave(1), nil},
			Note{D, MustNewOctave(2), nil},
			Note{E, MustNewOctave(3), nil},
			Note{F, MustNewOctave(4), nil},
			Note{G, MustNewOctave(5), nil},
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		for _, testCase := range testCases {
			assert.True(t, chord.Exists(testCase))
		}

		testCasesNotExist := Notes{
			Note{C, MustNewOctave(2), nil},
			Note{D, MustNewOctave(1), nil},
			Note{E, MustNewOctave(1), nil},
			Note{CFLAT, MustNewOctave(1), nil},
			Note{DSHARP, MustNewOctave(2), nil},
			Note{EFLAT2, MustNewOctave(3), nil},
			Note{FSHARP2, MustNewOctave(4), nil},
			Note{A, MustNewOctave(5), nil},
			Note{B, MustNewOctave(5), nil},
		}

		for _, testCase := range testCasesNotExist {
			assert.False(t, chord.Exists(testCase))
		}
	})

	t.Run("Chord_Exists: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.False(t, chord.Exists(Note{C, MustNewOctave(1), nil}))
	})
}
