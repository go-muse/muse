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
			note.C.MustNewNote(),
			note.E.MustNewNote(),
			note.G.MustNewNote(),
		}

		dur := duration.NewRelative("custom").
			SetDots(2).
			SetTuplet(tuplet.New(2, 3))

		chord := NewChord(notes...).SetDurationRel(dur)

		assert.Equal(t, len(notes), len(chord.notes), "expected %v notes, got %v", len(notes), len(chord.notes))
		assert.Equal(t, dur, chord.DurationRel(), "expected chord duration to be %v, got %v", dur, chord.DurationRel())

		for i, chordNote := range chord.notes {
			assert.Equal(t, dur, chordNote.DurationRel(), "expected note %v to have duration %v, got %v", i, dur, chordNote.DurationRel())
		}
	})

	t.Run("NewChord: creating a new chord with existing notes", func(t *testing.T) {
		notes := note.Notes{
			note.C.MustNewNote().SetOctave(octave.MustNewByNumber(octave.Number4)),
			note.E.MustNewNote().SetOctave(octave.MustNewByNumber(octave.Number4)),
			note.G.MustNewNote().SetOctave(octave.MustNewByNumber(octave.Number4)),
			note.C.MustNewNote().SetOctave(octave.MustNewByNumber(octave.Number4)), // existing note, won't be added
			note.E.MustNewNote().SetOctave(octave.MustNewByNumber(octave.Number4)), // existing note, won't be added
			note.G.MustNewNote().SetOctave(octave.MustNewByNumber(octave.Number4)), // existing note, won't be added
		}

		dur := duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3))
		chord := NewChord(notes...).SetDurationRel(dur)

		differentNotes := 3
		assert.Len(t, chord.notes, differentNotes, "expected %v notes, got %v", differentNotes, len(chord.notes))
		assert.Equal(t, dur, chord.DurationRel(), "expected chord duration to be %v, got %v", dur, chord.DurationRel())

		for i, chordNote := range chord.notes {
			assert.Equal(t, dur, chordNote.DurationRel(), "expected note %v to have duration %v, got %v", i, dur, chordNote.DurationRel())
		}
	})
}

func TestChord_String(t *testing.T) {
	t.Run("Chord_String: positive", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		assert.Equal(t,
			fmt.Sprintf("notes: %+v, duration name: %+v, custom duration: %+v", chord.notes, chord.DurationRel().Name(), chord.durationAbs),
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
	assert.Nil(t, chord.DurationRel(), "expected chord duration to be nil, got %v", chord.DurationRel())
}

func TestChord_AddNote(t *testing.T) {
	t.Run("Chord_AddNote: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := []struct {
			note *note.Note
		}{
			{note.MustNewNoteWithOctave(note.C, 1)},
			{note.MustNewNote(note.E).SetOctave(octave.MustNewByNumber(2)).SetDurationRel(duration.NewRelative(duration.NameHalf))},
			{note.MustNewNote(note.E).SetOctave(octave.MustNewByNumber(3)).SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletTriplet())},
		}

		for _, testCase := range testCases {
			chord.AddNote(testCase.note)
		}

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.DurationRel(), chordNote.DurationRel(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.DurationRel(), chordNote.DurationRel())
			assert.True(t, chordNote.IsEqualByName(testCases[i].note))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNote: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := []struct {
			note *note.Note
		}{
			{note.MustNewNoteWithOctave(note.C, 1)},
			{note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(duration.NewRelative(duration.NameHalf))},
			{note.MustNewNoteWithOctave(note.G, 3).SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletTriplet()).SetDurationAbs(time.Hour)},
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
			assert.Equal(t, chord.DurationRel(), chordNote.DurationRel(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.DurationRel(), chordNote.DurationRel())
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNote: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		note := note.MustNewNoteWithOctave(note.C, 1)
		chord.AddNote(note)
		assert.Nil(t, chord)
	})
}

func TestChord_AddNotes(t *testing.T) {
	t.Run("Chord_AddNotes: adding new notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := note.Notes{
			note.MustNewNoteWithOctave(note.C, 1),
			note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(duration.NewRelative(duration.NameHalf)),
			note.MustNewNoteWithOctave(note.E, 3).SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletTriplet()),
		}

		chord.AddNotes(testCases...)

		for i, chordNote := range chord.notes {
			assert.Equal(t, chord.DurationRel(), chordNote.DurationRel(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.DurationRel(), chordNote.DurationRel())
			assert.True(t, chordNote.IsEqualByName(testCases[i]))
		}

		assert.Equal(t, len(testCases), len(chord.notes))
	})

	t.Run("Chord_AddNotes: adding existing notes", func(t *testing.T) {
		chord := &Chord{
			notes:       nil,
			durationAbs: time.Second,
			durationRel: duration.NewRelative(duration.NameWhole).SetDots(2).SetTuplet(tuplet.New(2, 3)),
		}

		testCases := note.Notes{
			note.C.MustNewNote().SetOctave(octave.MustNewByNumber(1)),
			note.E.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameHalf)).SetOctave(octave.MustNewByNumber(2)),
			note.E.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletTriplet()).SetOctave(octave.MustNewByNumber(3)),
		}

		// adding notes
		chord.AddNotes(testCases...)

		// adding existing notes
		chord.AddNotes(testCases...)

		for _, chordNote := range chord.notes {
			assert.Equal(t, chord.DurationRel(), chordNote.DurationRel(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), chord.DurationRel(), chordNote.DurationRel())
		}

		assert.Equal(t, len(testCases), len(chord.notes), "%+v", chord.notes)
	})

	t.Run("Chord_AddNotes: adding to nil chord", func(t *testing.T) {
		var chord *Chord
		chord.AddNotes(note.MustNewNoteWithOctave(note.C, 1), note.MustNewNoteWithOctave(note.D, 2))
		assert.Nil(t, chord)
	})
}

func TestChord_GetNotes(t *testing.T) {
	t.Run("Chord_AddNotes: get notes", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := &Chord{
			notes: note.Notes{
				note.MustNewNoteWithOctave(note.C, 1).SetDurationRel(dur),
				note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(dur),
				note.MustNewNoteWithOctave(note.G, 3).SetDurationRel(dur),
			},
			durationRel: dur,
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
				note.MustNewNoteWithOctave(note.C, 1).SetDurationRel(dur),
				note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(dur),
				note.MustNewNoteWithOctave(note.G, 3).SetDurationRel(dur),
			},
			durationRel: dur,
		}

		newDuration := duration.NewRelative(duration.NameWhole)
		chord.SetDurationRel(newDuration)
		assert.Equal(t, newDuration, chord.DurationRel(), "expected chord duration: %+v, actual chord duration: %+v", newDuration, chord.DurationRel())

		for _, chordNote := range chord.notes {
			assert.Equal(t, newDuration, chordNote.DurationRel(), "note: %s, expected duration: %+v, actual: %+v", chordNote.Name(), newDuration, chordNote.DurationRel())
		}
	})

	t.Run("Chord_SetDurationRel: set duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetDurationRel(duration.NewRelative(duration.NameLong)))
	})
}

func TestChord_GetDuration(t *testing.T) {
	t.Run("Chord_GetDurationRel: getting duration from the chord", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := NewChord(
			note.MustNewNoteWithOctave(note.C, 1).SetDurationRel(dur),
			note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(dur),
			note.MustNewNoteWithOctave(note.G, 3).SetDurationRel(dur),
		).SetDurationRel(dur)

		assert.Equal(t, dur, chord.DurationRel(), " expected duration: %+v, actual: %+v", dur, chord.DurationRel())
	})

	t.Run("Chord_GetDurationRel: getting duration from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.DurationRel())
	})
}

func TestChord_SetDurationAbs(t *testing.T) {
	t.Run("Chord_SetDurationAbs: set custom duration to the chord", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := &Chord{
			notes: note.Notes{
				note.MustNewNoteWithOctave(note.C, 1).SetDurationRel(dur),
				note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(dur),
				note.MustNewNoteWithOctave(note.G, 3).SetDurationRel(dur),
			},
			durationRel: dur,
		}

		customDuration := time.Second
		chord.SetDurationAbs(customDuration)
		assert.Equal(t, customDuration, chord.durationAbs, "expected chord custom duration: %+v, actual chord custom duration: %+v", customDuration, chord.durationAbs)

		for _, chordNote := range chord.notes {
			assert.Equal(t, customDuration, chordNote.DurationAbs(), "note: %s, expected duration: %+v, actual duration: %+v", chordNote.Name(), customDuration, chordNote.DurationAbs())
		}
	})

	t.Run("Chord_SetDurationAbs: set custom duration to the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.Nil(t, chord.SetDurationAbs(time.Second))
	})
}

func TestChord_DurationAbs(t *testing.T) {
	t.Run("Chord_DurationAbs: getting custom duration from the chord", func(t *testing.T) {
		dur := duration.NewRelative(duration.NameHalf)
		chord := NewChord(
			note.MustNewNoteWithOctave(note.C, 1).SetDurationRel(dur),
			note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(dur),
			note.MustNewNoteWithOctave(note.G, 3).SetDurationRel(dur),
		).SetDurationRel(dur)

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
			note.MustNewNoteWithOctave(note.C, 1).SetDurationRel(nil),
			note.MustNewNoteWithOctave(note.E, 2).SetDurationRel(nil),
			note.MustNewNoteWithOctave(note.G, 3).SetDurationRel(nil),
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
			note.MustNewNoteWithOctave(note.C, 1),
			note.MustNewNoteWithOctave(note.D, 2),
			note.MustNewNoteWithOctave(note.E, 3),
			note.MustNewNoteWithOctave(note.F, 4),
			note.MustNewNoteWithOctave(note.G, 5),
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
		assert.Nil(t, chord.RemoveNote(note.MustNewNoteWithOctave(note.C, 1)))
	})
}

func TestChord_RemoveNotes(t *testing.T) {
	t.Run("Chord_RemoveNotes: remove from the chord", func(t *testing.T) {
		testCases := note.Notes{
			note.MustNewNoteWithOctave(note.C, 1),
			note.MustNewNoteWithOctave(note.D, 2),
			note.MustNewNoteWithOctave(note.E, 3),
			note.MustNewNoteWithOctave(note.F, 4),
			note.MustNewNoteWithOctave(note.G, 5),
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
			note.MustNewNoteWithOctave(note.C, 1),
			note.MustNewNoteWithOctave(note.D, 2),
			note.MustNewNoteWithOctave(note.E, 3),
			note.MustNewNoteWithOctave(note.F, 4),
			note.MustNewNoteWithOctave(note.G, 5),
		}

		chord := NewChordEmpty().AddNotes(testCases...)
		for _, testCase := range testCases {
			assert.True(t, chord.Exists(testCase))
		}

		testCasesNotExist := note.Notes{
			note.MustNewNoteWithOctave(note.C, 2),
			note.MustNewNoteWithOctave(note.D, 1),
			note.MustNewNoteWithOctave(note.E, 1),
			note.MustNewNoteWithOctave(note.CFLAT, 1),
			note.MustNewNoteWithOctave(note.DSHARP, 2),
			note.MustNewNoteWithOctave(note.EFLAT2, 3),
			note.MustNewNoteWithOctave(note.FSHARP2, 4),
			note.MustNewNoteWithOctave(note.A, 5),
			note.MustNewNoteWithOctave(note.B, 5),
		}

		for _, testCase := range testCasesNotExist {
			assert.False(t, chord.Exists(testCase))
		}
	})

	t.Run("Chord_Exists: remove from the nil chord", func(t *testing.T) {
		var chord *Chord
		assert.False(t, chord.Exists(note.MustNewNoteWithOctave(note.C, 1)))
	})
}
