package chord

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
	"github.com/go-muse/muse/tuplet"
)

func TestNewChord(t *testing.T) {
	t.Run("NewChord: creating a new chord with different notes", func(t *testing.T) {
		notes := note.Notes{
			note.C.NewNote(),
			note.E.NewNote(),
			note.G.NewNote(),
		}

		dur := duration.NewRelative("custom").
			SetDots(2).
			SetTuplet(tuplet.New(2, 3))

		chord := NewChord(notes...).SetValue(dur)

		assert.Equal(t, len(notes), len(chord.notes), "expected %v notes, got %v", len(notes), len(chord.notes))
		assert.Equal(t, dur, chord.Value(), "expected chord duration to be %v, got %v", dur, chord.Value())

		for i, chordNote := range chord.notes {
			assert.Equal(t, dur, chordNote.Value(), "expected note %v to have duration %v, got %v", i, dur, chordNote.Value())
		}
	})

	t.Run("NewChord: creating a new chord with existing notes", func(t *testing.T) {
		notes := note.Notes{
			note.C.NewNote().SetOctave(octave.MustNewByNumber(octave.Number4)),
			note.E.NewNote().SetOctave(octave.MustNewByNumber(octave.Number4)),
			note.G.NewNote().SetOctave(octave.MustNewByNumber(octave.Number4)),
			note.C.NewNote().SetOctave(octave.MustNewByNumber(octave.Number4)), // existing note, won't be added
			note.E.NewNote().SetOctave(octave.MustNewByNumber(octave.Number4)), // existing note, won't be added
			note.G.NewNote().SetOctave(octave.MustNewByNumber(octave.Number4)), // existing note, won't be added
		}

		dur := duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3))
		chord := NewChord(notes...).SetValue(dur)

		differentNotes := 3
		assert.Len(t, chord.notes, differentNotes, "expected %v notes, got %v", differentNotes, len(chord.notes))
		assert.Equal(t, dur, chord.Value(), "expected chord duration to be %v, got %v", dur, chord.Value())

		for i, chordNote := range chord.notes {
			assert.Equal(t, dur, chordNote.Value(), "expected note %v to have duration %v, got %v", i, dur, chordNote.Value())
		}
	})
}

func TestChord_String(t *testing.T) {
	t.Run("Chord_String: positive", func(t *testing.T) {
		chord := &Chord{
			notes:    nil,
			duration: time.Second,
			value:    duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		assert.Equal(t,
			fmt.Sprintf("notes: %+v, duration name: %+v, custom duration: %+v", chord.notes, chord.Value().Name(), chord.duration),
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
	assert.Empty(t, chord.notes, "expected %v notes, got %v", 0, len(chord.notes))
	assert.Nil(t, chord.Value(), "expected chord duration to be nil, got %v", chord.Value())
}

func TestChord_AddNote(t *testing.T) {
	t.Run("Chord_AddNote: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes:    nil,
			duration: time.Second,
			value:    duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := []struct {
			note note.Note
		}{
			{WithOctave(note.C, 1)},
			{(note.E).SetOctave(octave.MustNewByNumber(2)).SetValue(duration.NewRelative(duration.NameHalf))},
			{(note.E).SetOctave(octave.MustNewByNumber(3)).SetValue(duration.NewRelative(duration.NameWhole).SetTupletTriplet())},
		}

		for _, testCase := range testCases {
			chord.AddNote(testCase.note)
		}

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.Value(), chordNote.Value(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.Value(), chordNote.Value())
			assert.True(t, chordNote.EqualByName(testCases[i].note))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNote: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes:    nil,
			duration: time.Second,
			value:    duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := []struct {
			note note.Note
		}{
			{WithOctave(note.C, 1)},
			{WithOctave(note.E, 2).SetValue(duration.NewRelative(duration.NameHalf))},
			{WithOctave(note.G, 3).SetValue(duration.NewRelative(duration.NameWhole).SetTupletTriplet()).SetDuration(time.Hour)},
		}

		// adding notes
		for _, testCase := range testCases {
			chord.AddNote(testCase.note)
		}

		// adding existing notes
		for _, testCase := range testCases {
			chord.AddNote(testCase.note)
		}

		for _, chordNote := range chord.notes {
			assert.Equal(t, chord.Value(), chordNote.Value(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.Value(), chordNote.Value())
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNote: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		note := note.MustNewWithOctave(note.C, 1)
		chord.AddNote(note)
		assert.Nil(t, chord)
	})
}

func TestChord_AddNotes(t *testing.T) {
	t.Run("Chord_AddNotes: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes:    nil,
			duration: time.Second,
			value:    duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := note.Notes{
			WithOctave(note.C, 1),
			WithOctave(note.E, 2).SetValue(duration.NewRelative(duration.NameHalf)),
			WithOctave(note.E, 3).SetValue(duration.NewRelative(duration.NameWhole).SetTupletTriplet()),
		}

		chord.AddNotes(testCases...)

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.Value(), chordNote.Value(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.Value(), chordNote.Value())
			assert.True(t, chordNote.EqualByName(testCases[i]))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNotes: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes:    nil,
			duration: time.Second,
			value:    duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := note.Notes{
			note.C.NewNote().SetOctave(octave.MustNewByNumber(1)),
			note.E.NewNote().SetValue(duration.NewRelative(duration.NameHalf)).SetOctave(octave.MustNewByNumber(2)),
			note.E.NewNote().SetValue(duration.NewRelative(duration.NameWhole).SetTupletTriplet()).SetOctave(octave.MustNewByNumber(3)),
		}

		// adding notes
		chord.AddNotes(testCases...)

		// adding existing notes
		chord.AddNotes(testCases...)

		for _, chordNote := range chord.notes {
			assert.Equal(t, chord.Value(), chordNote.Value(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.Value(), chordNote.Value())
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNotes: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		chord.AddNotes(WithOctave(note.C, 1), note.MustNewWithOctave(note.D, 2))
		assert.Nil(t, chord)
	})
}

func TestChord_GetNotes(t *testing.T) {
	t.Run("Chord_AddNotes: get notes", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := &Chord{
			notes: note.Notes{
				WithOctave(note.C, 1).SetValue(dur),
				WithOctave(note.E, 2).SetValue(dur),
				WithOctave(note.G, 3).SetValue(dur),
			},
			value: dur,
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
		dur := duration.NewRelative(duration.NameHalf)
		chord := &Chord{
			notes: note.Notes{
				WithOctave(note.C, 1).SetValue(dur),
				WithOctave(note.E, 2).SetValue(dur),
				WithOctave(note.G, 3).SetValue(dur),
			},
			value: dur,
		}

		newDuration := duration.NewRelative(duration.NameWhole)
		chord.SetValue(newDuration)
		assert.Equal(t, newDuration, chord.Value(), "expected chord duration: %+v, actual chord duration: %+v", newDuration, chord.Value())

		for _, chordNote := range chord.notes {
			assert.Equal(t, newDuration, chordNote.Value(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), newDuration, chordNote.Value())
		}
	})

	t.Run("Chord_SetDurationRel: set duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetValue(duration.NewRelative(duration.NameLong)))
	})
}

func TestChord_GetDuration(t *testing.T) {
	t.Run("Chord_GetDurationRel: getting duration from the chord", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := NewChord(
			WithOctave(note.C, 1).SetValue(dur),
			WithOctave(note.E, 2).SetValue(dur),
			WithOctave(note.G, 3).SetValue(dur),
		).SetValue(dur)

		assert.Equal(t, dur, chord.Value(), " expected duration: %+v, actual: %+v", dur, chord.Value())
	})

	t.Run("Chord_GetDurationRel: getting duration from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.Value())
	})
}

func TestChord_SetDurationAbs(t *testing.T) {
	t.Run("Chord_SetDurationAbs: set custom duration to the chord", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := &Chord{
			notes: note.Notes{
				WithOctave(note.C, 1).SetValue(dur),
				WithOctave(note.E, 2).SetValue(dur),
				WithOctave(note.G, 3).SetValue(dur),
			},
			value: dur,
		}

		customDuration := time.Second
		chord.SetDuration(customDuration)
		assert.Equal(t, customDuration, chord.duration, "expected chord custom duration: %+v, actual chord custom duration: %+v", customDuration, chord.duration)

		for _, chordNote := range chord.notes {
			assert.Equal(t, customDuration, chordNote.Duration(), "note: %s, expected duration: %+v, actual duration: %+v", chordNote.Name(), customDuration, chordNote.Duration())
		}
	})

	t.Run("Chord_SetDurationAbs: set custom duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetDuration(time.Second))
	})
}

func TestChord_DurationAbs(t *testing.T) {
	t.Run("Chord_DurationAbs: getting custom duration from the chord", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := NewChord(
			WithOctave(note.C, 1).SetValue(dur),
			WithOctave(note.E, 2).SetValue(dur),
			WithOctave(note.G, 3).SetValue(dur),
		).SetValue(dur)

		customDuration := time.Second
		chord.SetDuration(customDuration)

		assert.Equal(t, customDuration, chord.Duration(), " expected custom duration: %d, actual custom duration: %d", customDuration, chord.Duration())
	})

	t.Run("Chord_DurationAbs: getting custom duration from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Zero(t, chord.Duration())
	})
}

func TestChord_Empty(t *testing.T) {
	t.Run("Chord_Empty: clearing the chord", func(t *testing.T) {
		chord := NewChordEmpty().AddNotes(
			WithOctave(note.C, 1).SetValue(nil),
			WithOctave(note.E, 2).SetValue(nil),
			WithOctave(note.G, 3).SetValue(nil),
		)

		assert.Empty(t, chord.Empty().notes)
	})

	t.Run("Chord_Empty: clearing the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.Empty())
	})
}

func TestChord_RemoveNote(t *testing.T) {
	t.Run("Chord_RemoveNote: remove from the chord", func(t *testing.T) {
		testCases := note.Notes{
			WithOctave(note.C, 1),
			WithOctave(note.D, 2),
			WithOctave(note.E, 3),
			WithOctave(note.F, 4),
			WithOctave(note.G, 5),
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		assert.Equal(t, testCases, chord.notes)

		for range testCases {
			chord.RemoveNote(testCases[0])
			testCases = testCases[1:]
			assert.Equal(t, testCases, chord.notes)
		}
	})

	t.Run("Chord_RemoveNote: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.RemoveNote(WithOctave(note.C, 1)))
	})
}

func TestChord_RemoveNotes(t *testing.T) {
	t.Run("Chord_RemoveNotes: remove from the chord", func(t *testing.T) {
		testCases := note.Notes{
			WithOctave(note.C, 1),
			WithOctave(note.D, 2),
			WithOctave(note.E, 3),
			WithOctave(note.F, 4),
			WithOctave(note.G, 5),
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		assert.Equal(t, testCases, chord.notes)

		chord.RemoveNotes(note.Notes{testCases[1], testCases[3]})
		testCases = append(testCases[0:1], testCases[2:]...)
		testCases = append(testCases[0:2], testCases[3:]...)
		assert.Equal(t, testCases, chord.notes)
	})

	t.Run("Chord_RemoveNotes: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.RemoveNotes(note.Notes{}))
	})
}

func TestChord_Exists(t *testing.T) {
	t.Run("Chord_Exists: remove from the chord", func(t *testing.T) {
		testCases := note.Notes{
			WithOctave(note.C, 1),
			WithOctave(note.D, 2),
			WithOctave(note.E, 3),
			WithOctave(note.F, 4),
			WithOctave(note.G, 5),
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		for _, testCase := range testCases {
			assert.True(t, chord.Exists(testCase))
		}

		testCasesNotExist := note.Notes{
			WithOctave(note.C, 2),
			WithOctave(note.D, 1),
			WithOctave(note.E, 1),
			WithOctave(note.CFLAT, 1),
			WithOctave(note.DSHARP, 2),
			WithOctave(note.EFLAT2, 3),
			WithOctave(note.FSHARP2, 4),
			WithOctave(note.A, 5),
			WithOctave(note.B, 5),
		}

		for _, testCase := range testCasesNotExist {
			assert.False(t, chord.Exists(testCase))
		}
	})

	t.Run("Chord_Exists: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.False(t, chord.Exists(WithOctave(note.C, 1)))
	})
}
