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
