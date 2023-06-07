package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoteMIDINumber(t *testing.T) {
	type testCase struct {
		note *Note
		want uint8
	}

	t.Run("TestMIDINumber: main cases", func(t *testing.T) {
		var testCases []testCase

		for i := minOctaveNumber; i <= maxOctaveNumber; i++ {
			for _, note := range GetFullChromaticScale() {
				octaveNumberFromZero := uint8(i + 1)
				testCases = append(testCases, testCase{
					note: note.Copy().SetOctave(MustNewOctave(OctaveNumber(i))),
					want: octaveNumberFromZero*NotesInOctave + uint8(note.getNoteNumberWithinOctave()),
				})

				// 127th note is G
				if i == maxOctaveNumber && note.Name() == G {
					break
				}
			}
		}

		for _, testCase := range testCases {
			assert.Equal(t, testCase.want, testCase.note.MIDINumber(),
				"note: %s, octave: %s, expected MIDI number: %d, actual: %d, getNoteNumberWithinOctave():%d", testCase.note.Name(), testCase.note.Octave().Name(), testCase.want, testCase.note.MIDINumber(), testCase.note.getNoteNumberWithinOctave())
		}
	})

	t.Run("TestMIDINumber: corner cases", func(t *testing.T) {
		testCases := []testCase{
			{
				note: nil,
				want: 0,
			},
			{
				note: MustNewNoteWithoutOctave(C),
				want: 0,
			},
			{
				note: newNoteWithOctave(C, MustNewOctave(-1)),
				want: 0,
			},
			{
				note: newNoteWithOctave(C, MustNewOctave(-1)).AlterDown(),
				want: 0,
			},
			{
				note: newNoteWithOctave(C, MustNewOctave(-1)).AlterDownBy(2),
				want: 0,
			},
			{
				note: newNoteWithOctave(C, MustNewOctave(-1)).AlterDownBy(3),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, MustNewOctave(-1)).AlterDown(),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, MustNewOctave(-1)).AlterDownBy(2),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, MustNewOctave(-1)).AlterDownBy(3),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP2, MustNewOctave(-1)).AlterDownBy(2),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP2, MustNewOctave(-1)).AlterDownBy(3),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, MustNewOctave(-1)),
				want: 1,
			},
			{
				note: newNoteWithOctave(DFLAT, MustNewOctave(-1)),
				want: 1,
			},
			{
				note: newNoteWithOctave(CSHARP2, MustNewOctave(-1)).AlterDown(),
				want: 1,
			},
			{
				note: newNoteWithOctave(G, MustNewOctave(9)),
				want: 127,
			},
			{
				note: newNoteWithOctave(A, MustNewOctave(9)),
				want: 127,
			},
			{
				note: newNoteWithOctave(B, MustNewOctave(9)),
				want: 127,
			},
			{
				note: newNoteWithOctave(G, MustNewOctave(9)).AlterUp(),
				want: 127,
			},
			{
				note: newNoteWithOctave(G, MustNewOctave(9)).AlterUpBy(2),
				want: 127,
			},
			{
				note: newNoteWithOctave(BSHARP, MustNewOctave(9)).AlterUpBy(2),
				want: 127,
			},
		}

		for _, testCase := range testCases {
			assert.Equal(t, testCase.want, testCase.note.MIDINumber(),
				"note: %s, octave: %s, expected MIDI number: %d, actual: %d, getNoteNumberWithinOctave():%d", testCase.note.Name(), testCase.note.Octave().Name(), testCase.want, testCase.note.MIDINumber(), testCase.note.getNoteNumberWithinOctave())
		}
	})
}

func Test_NewNoteFromMIDINumber(t *testing.T) {
	t.Run("NewNoteFromMIDINumber: positive cases", func(t *testing.T) {
		testCases := []struct {
			midiNumber uint8
			want       *Note
		}{
			{midiNumber: 0, want: MustNewNote(C, -1)},
			{midiNumber: 1, want: MustNewNote(CSHARP, -1)},
			{midiNumber: 2, want: MustNewNote(D, -1)},
			{midiNumber: 3, want: MustNewNote(DSHARP, -1)},
			{midiNumber: 4, want: MustNewNote(E, -1)},
			{midiNumber: 5, want: MustNewNote(F, -1)},
			{midiNumber: 6, want: MustNewNote(FSHARP, -1)},
			{midiNumber: 7, want: MustNewNote(G, -1)},
			{midiNumber: 8, want: MustNewNote(GSHARP, -1)},
			{midiNumber: 9, want: MustNewNote(A, -1)},
			{midiNumber: 10, want: MustNewNote(ASHARP, -1)},
			{midiNumber: 11, want: MustNewNote(B, -1)},
			{midiNumber: 12, want: MustNewNote(C, 0)},
			{midiNumber: 24, want: MustNewNote(C, 1)},
			{midiNumber: 36, want: MustNewNote(C, 2)},
			{midiNumber: 126, want: MustNewNote(FSHARP, 9)},
			{midiNumber: 127, want: MustNewNote(G, 9)},
		}

		var midiNumber *Note
		var err error
		for _, testCase := range testCases {
			midiNumber, err = NewNoteFromMIDINumber(testCase.midiNumber)
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, midiNumber)
		}
	})

	t.Run("NewNoteFromMIDINumber: negative cases", func(t *testing.T) {
		testCases := []struct {
			midiNumber uint8
			want       error
		}{
			{128, ErrMIDINumberUnknown},
			{129, ErrMIDINumberUnknown},
			{255, ErrMIDINumberUnknown},
		}

		var midiNumber *Note
		var err error
		for _, testCase := range testCases {
			midiNumber, err = NewNoteFromMIDINumber(testCase.midiNumber)
			assert.Nil(t, midiNumber)
			assert.ErrorIs(t, err, testCase.want)
		}
	})
}
