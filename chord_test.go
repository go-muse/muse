package muse

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewChord(t *testing.T) {
	t.Run("NewChord: creating a new chord with different notes", func(t *testing.T) {
		notes := Notes{
			*C.NewNote(),
			*E.NewNote(),
			*G.NewNote(),
		}

		duration := &DurationRel{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}}

		chord := NewChord(notes...).SetDurationRel(duration)

		assert.Equal(t, len(notes), len(chord.notes), "expected %v notes, got %v", len(notes), len(chord.notes))
		assert.Equal(t, duration, chord.GetDurationRel(), "expected chord duration to be %v, got %v", duration, chord.GetDurationRel())

		for i, chordNote := range chord.notes {
			assert.Equal(t, duration, chordNote.durationRel, "expected note %v to have duration %v, got %v", i, duration, chordNote.durationRel)
		}
	})

	t.Run("NewChord: creating a new chord with existing notes", func(t *testing.T) {
		notes := Notes{
			*C.NewNote().SetOctave(MustNewOctave(OctaveNumber4)),
			*E.NewNote().SetOctave(MustNewOctave(OctaveNumber4)),
			*G.NewNote().SetOctave(MustNewOctave(OctaveNumber4)),
			*C.NewNote().SetOctave(MustNewOctave(OctaveNumber4)), // existing note, won't be added
			*E.NewNote().SetOctave(MustNewOctave(OctaveNumber4)), // existing note, won't be added
			*G.NewNote().SetOctave(MustNewOctave(OctaveNumber4)), // existing note, won't be added
		}

		duration := &DurationRel{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}}

		chord := NewChord(notes...).SetDurationRel(duration)

		differentNotes := 3
		assert.Equal(t, differentNotes, len(chord.notes), "expected %v notes, got %v", differentNotes, len(chord.notes))
		assert.Equal(t, duration, chord.GetDurationRel(), "expected chord duration to be %v, got %v", duration, chord.GetDurationRel())

		for i, chordNote := range chord.notes {
			assert.Equal(t, duration, chordNote.durationRel, "expected note %v to have duration %v, got %v", i, duration, chordNote.durationRel)
		}
	})
}

func TestChord_String(t *testing.T) {
	t.Run("Chord_String: positive", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: &DurationRel{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
		}

		assert.Equal(t,
			fmt.Sprintf("notes: %+v, duration name: %+v, custom duration: %+v", chord.notes, chord.durationRel.Name(), chord.durationAbs),
			chord.String(),
		)
	})

	t.Run("Chord_String: negative", func(t *testing.T) {
		var chord *Chord
		assert.Equal(t, "nil chord", chord.String())
	})

}

func TestNewChordEmpty(t *testing.T) {
	chord := NewChordEmpty()
	assert.Equal(t, 0, len(chord.notes), "expected %v notes, got %v", 0, len(chord.notes))
	assert.Nil(t, chord.durationRel, "expected chord duration to be nil, got %v", chord.durationRel)
}

func TestChord_AddNote(t *testing.T) {
	t.Run("Chord_AddNote: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: &DurationRel{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
		}

		testCases := []struct {
			note *Note
		}{
			{&Note{C, MustNewOctave(1), 0, nil}},
			{&Note{E, MustNewOctave(2), 0, NewDurationRel(DurationNameHalf)}},
			{&Note{G, MustNewOctave(3), time.Hour, &DurationRel{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}}}},
		}

		for _, testCase := range testCases {
			chord.AddNote(*testCase.note)
		}

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.GetDurationRel(), chordNote.durationRel, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.GetDurationRel(), chordNote.durationRel)
			assert.True(t, chordNote.IsEqualByName(testCases[i].note))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNote: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: &DurationRel{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
		}

		testCases := []struct {
			note *Note
		}{
			{&Note{C, MustNewOctave(1), 0, nil}},
			{&Note{E, MustNewOctave(2), 0, NewDurationRel(DurationNameHalf)}},
			{&Note{G, MustNewOctave(3), time.Hour, &DurationRel{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}}}},
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
			assert.Equal(t, chord.GetDurationRel(), chordNote.durationRel, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.GetDurationRel(), chordNote.durationRel)
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNote: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		note := MustNewNoteWithOctave(C, 1)
		chord.AddNote(*note)
		assert.Nil(t, chord)
	})
}

func TestChord_AddNotes(t *testing.T) {
	t.Run("Chord_AddNotes: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: &DurationRel{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
		}

		testCases := Notes{
			Note{C, MustNewOctave(1), 0, nil},
			Note{E, MustNewOctave(2), 0, NewDurationRel(DurationNameHalf)},
			Note{G, MustNewOctave(3), time.Hour, &DurationRel{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}}},
		}

		chord.AddNotes(testCases...)

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.GetDurationRel(), chordNote.durationRel, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.GetDurationRel(), chordNote.durationRel)
			assert.True(t, chordNote.IsEqualByName(&testCases[i]))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNotes: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: &DurationRel{name: "custom", dots: 2, tuplet: &Tuplet{n: 2, m: 3}},
		}

		testCases := Notes{
			Note{C, MustNewOctave(1), 0, nil},
			Note{E, MustNewOctave(2), 0, NewDurationRel(DurationNameHalf)},
			Note{G, MustNewOctave(3), time.Hour, &DurationRel{name: "custom 2", dots: 0, tuplet: &Tuplet{n: 3, m: 2}}},
		}

		// adding notes
		chord.AddNotes(testCases...)

		// adding existing notes
		chord.AddNotes(testCases...)

		for _, chordNote := range chord.notes {
			assert.Equal(t, chord.GetDurationRel(), chordNote.durationRel, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, chord.GetDurationRel(), chordNote.durationRel)
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNotes: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		chord.AddNotes(*MustNewNoteWithOctave(C, 1), *MustNewNoteWithOctave(D, 2))
		assert.Nil(t, chord)
	})
}

func TestChord_GetNotes(t *testing.T) {
	t.Run("Chord_AddNotes: get notes", func(t *testing.T) {
		duration := NewDurationRel(DurationNameHalf)
		chord := &Chord{
			notes: Notes{
				Note{C, MustNewOctave(1), 0, duration},
				Note{E, MustNewOctave(2), 0, duration},
				Note{G, MustNewOctave(3), 0, duration},
			},
			durationRel: duration,
		}

		notes := chord.Notes()
		assert.Equal(t, len(chord.notes), len(notes))
		assert.Equal(t, chord.notes, notes)
	})

	t.Run("Chord_AddNotes: getting from nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.Notes())
	})
}

func TestChord_SetDurationRel(t *testing.T) {
	t.Run("Chord_SetDurationRel: set duration to the chord", func(t *testing.T) {
		duration := NewDurationRel(DurationNameHalf)
		chord := &Chord{
			notes: Notes{
				Note{C, MustNewOctave(1), 0, duration},
				Note{E, MustNewOctave(2), 0, duration},
				Note{G, MustNewOctave(3), 0, duration},
			},
			durationRel: duration,
		}

		newDuration := NewDurationRel(DurationNameWhole)
		chord.SetDurationRel(newDuration)
		assert.Equal(t, newDuration, chord.GetDurationRel(), "expected chord duration: %+v, actual chord duration: %+v", newDuration, chord.GetDurationRel())

		for _, chordNote := range chord.notes {
			assert.Equal(t, newDuration, chordNote.durationRel, "note: %s, expected duration: %+v, actual: %+v", chordNote.name, newDuration, chordNote.durationRel)
		}
	})

	t.Run("Chord_SetDurationRel: set duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetDurationRel(NewDurationRel(DurationNameLong)))
	})
}

func TestChord_GetDuration(t *testing.T) {
	t.Run("Chord_GetDurationRel: getting duration from the chord", func(t *testing.T) {
		duration := NewDurationRel(DurationNameHalf)
		chord := NewChord(
			Note{C, MustNewOctave(1), 0, duration},
			Note{E, MustNewOctave(2), 0, duration},
			Note{G, MustNewOctave(3), 0, duration},
		).SetDurationRel(duration)

		assert.Equal(t, duration, chord.GetDurationRel(), " expected duration: %+v, actual: %+v", duration, chord.GetDurationRel())
	})

	t.Run("Chord_GetDurationRel: getting duration from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.GetDurationRel())
	})
}

func TestChord_SetDurationAbs(t *testing.T) {
	t.Run("Chord_SetDurationAbs: set custom duration to the chord", func(t *testing.T) {
		duration := NewDurationRel(DurationNameHalf)
		chord := &Chord{
			notes: Notes{
				Note{C, MustNewOctave(1), 0, duration},
				Note{E, MustNewOctave(2), 0, duration},
				Note{G, MustNewOctave(3), 0, duration},
			},
			durationRel: duration,
		}

		customDuration := time.Second
		chord.SetDurationAbs(customDuration)
		assert.Equal(t, customDuration, chord.durationAbs, "expected chord custom duration: %+v, actual chord custom duration: %+v", customDuration, chord.durationAbs)

		for _, chordNote := range chord.notes {
			assert.Equal(t, customDuration, chordNote.durationAbs, "note: %s, expected duration: %+v, actual duration: %+v", chordNote.name, customDuration, chordNote.durationAbs)
		}
	})

	t.Run("Chord_SetDurationAbs: set custom duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetDurationAbs(time.Second))
	})
}

func TestChord_DurationAbs(t *testing.T) {
	t.Run("Chord_DurationAbs: getting custom duration from the chord", func(t *testing.T) {
		duration := NewDurationRel(DurationNameHalf)
		chord := NewChord(
			Note{C, MustNewOctave(1), 0, duration},
			Note{E, MustNewOctave(2), 0, duration},
			Note{G, MustNewOctave(3), 0, duration},
		).SetDurationRel(duration)

		customDuration := time.Second
		chord.SetDurationAbs(customDuration)

		assert.Equal(t, customDuration, chord.GetDurationAbs(), " expected custom duration: %d, actual custom duration: %d", customDuration, chord.GetDurationAbs())
	})

	t.Run("Chord_DurationAbs: getting custom duration from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Zero(t, chord.GetDurationAbs())
	})
}

func TestChord_Empty(t *testing.T) {
	t.Run("Chord_Empty: clearing the chord", func(t *testing.T) {
		chord := NewChordEmpty().AddNotes(
			Note{C, MustNewOctave(1), 0, nil},
			Note{E, MustNewOctave(2), 0, nil},
			Note{G, MustNewOctave(3), 0, nil},
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
			Note{C, MustNewOctave(1), 0, nil},
			Note{D, MustNewOctave(2), 0, nil},
			Note{E, MustNewOctave(3), 0, nil},
			Note{F, MustNewOctave(4), 0, nil},
			Note{G, MustNewOctave(5), 0, nil},
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		assert.Equal(t, testCases, chord.notes)

		length := len(testCases)
		for i := 0; i < length; i++ {
			chord.RemoveNote(&testCases[0])
			testCases = testCases[1:]
			assert.Equal(t, testCases, chord.notes)
		}
	})

	t.Run("Chord_RemoveNote: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.RemoveNote(&Note{C, MustNewOctave(1), 0, nil}))
	})
}

func TestChord_RemoveNotes(t *testing.T) {
	t.Run("Chord_RemoveNotes: remove from the chord", func(t *testing.T) {
		testCases := Notes{
			Note{C, MustNewOctave(1), 0, nil},
			Note{D, MustNewOctave(2), 0, nil},
			Note{E, MustNewOctave(3), 0, nil},
			Note{F, MustNewOctave(4), 0, nil},
			Note{G, MustNewOctave(5), 0, nil},
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
			Note{C, MustNewOctave(1), 0, nil},
			Note{D, MustNewOctave(2), 0, nil},
			Note{E, MustNewOctave(3), 0, nil},
			Note{F, MustNewOctave(4), 0, nil},
			Note{G, MustNewOctave(5), 0, nil},
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		for _, testCase := range testCases {
			assert.True(t, chord.Exists(&testCase))
		}

		testCasesNotExist := Notes{
			Note{C, MustNewOctave(2), 0, nil},
			Note{D, MustNewOctave(1), 0, nil},
			Note{E, MustNewOctave(1), 0, nil},
			Note{CFLAT, MustNewOctave(1), 0, nil},
			Note{DSHARP, MustNewOctave(2), 0, nil},
			Note{EFLAT2, MustNewOctave(3), 0, nil},
			Note{FSHARP2, MustNewOctave(4), 0, nil},
			Note{A, MustNewOctave(5), 0, nil},
			Note{B, MustNewOctave(5), 0, nil},
		}

		for _, testCase := range testCasesNotExist {
			assert.False(t, chord.Exists(&testCase))
		}
	})

	t.Run("Chord_Exists: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.False(t, chord.Exists(&Note{C, MustNewOctave(1), 0, nil}))
	})
}
