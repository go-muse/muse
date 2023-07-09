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
				note: MustNewNote(C),
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
				note: newNoteWithOctave(C, MustNewOctave(4)),
				want: 60,
			},
			{
				note: newNoteWithOctave(CSHARP, MustNewOctave(4)),
				want: 61,
			},
			{
				note: newNoteWithOctave(D, MustNewOctave(4)),
				want: 62,
			},
			{
				note: newNoteWithOctave(DSHARP, MustNewOctave(4)),
				want: 63,
			},
			{
				note: newNoteWithOctave(E, MustNewOctave(4)),
				want: 64,
			},
			{
				note: newNoteWithOctave(F, MustNewOctave(4)),
				want: 65,
			},
			{
				note: newNoteWithOctave(FSHARP, MustNewOctave(4)),
				want: 66,
			},
			{
				note: newNoteWithOctave(G, MustNewOctave(4)),
				want: 67,
			},
			{
				note: newNoteWithOctave(GSHARP, MustNewOctave(4)),
				want: 68,
			},
			{
				note: newNoteWithOctave(A, MustNewOctave(4)),
				want: 69,
			},
			{
				note: newNoteWithOctave(ASHARP, MustNewOctave(4)),
				want: 70,
			},
			{
				note: newNoteWithOctave(B, MustNewOctave(4)),
				want: 71,
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
			{midiNumber: 0, want: MustNewNoteWithOctave(C, -1)},
			{midiNumber: 1, want: MustNewNoteWithOctave(CSHARP, -1)},
			{midiNumber: 2, want: MustNewNoteWithOctave(D, -1)},
			{midiNumber: 3, want: MustNewNoteWithOctave(DSHARP, -1)},
			{midiNumber: 4, want: MustNewNoteWithOctave(E, -1)},
			{midiNumber: 5, want: MustNewNoteWithOctave(F, -1)},
			{midiNumber: 6, want: MustNewNoteWithOctave(FSHARP, -1)},
			{midiNumber: 7, want: MustNewNoteWithOctave(G, -1)},
			{midiNumber: 8, want: MustNewNoteWithOctave(GSHARP, -1)},
			{midiNumber: 9, want: MustNewNoteWithOctave(A, -1)},
			{midiNumber: 10, want: MustNewNoteWithOctave(ASHARP, -1)},
			{midiNumber: 11, want: MustNewNoteWithOctave(B, -1)},
			{midiNumber: 12, want: MustNewNoteWithOctave(C, 0)},
			{midiNumber: 24, want: MustNewNoteWithOctave(C, 1)},
			{midiNumber: 36, want: MustNewNoteWithOctave(C, 2)},
			{midiNumber: 126, want: MustNewNoteWithOctave(FSHARP, 9)},
			{midiNumber: 127, want: MustNewNoteWithOctave(G, 9)},
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
