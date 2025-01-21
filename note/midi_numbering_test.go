package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/common/convert"
	"github.com/go-muse/muse/octave"
)

func TestNoteMIDINumber(t *testing.T) {
	type testCase struct {
		note *Note
		want uint8
	}

	t.Run("TestMIDINumber: main cases", func(t *testing.T) {
		var testCases []testCase

		for i := octave.MinOctaveNumber; i <= octave.MaxOctaveNumber; i++ {
			for _, note := range GetSetFullChromatic() {
				octaveNumberFromZero := convert.AddUint8Int8(1, int8(i))

				noteNum := note.mustGetNoteNumberWithinOctave()

				midiNumber := octaveNumberFromZero*octave.NotesInOctave + noteNum
				if midiNumber < minMIDINumber || midiNumber > maxMIDINumber {
					t.Fatalf("MIDI number out of uint8 range: %d", midiNumber)
				}

				testCases = append(testCases, testCase{
					note: note.Copy().SetOctave(octave.MustNewByNumber(i)),
					want: midiNumber,
				})

				// 127th note is G
				if i == octave.MaxOctaveNumber && note.Name() == G {
					break
				}
			}
		}

		for _, testCase := range testCases {
			assert.Equal(t, testCase.want, testCase.note.MIDINumber(),
				"note: %s, octave: %s, expected MIDI number: %d, actual: %d, mustGetNoteNumberWithinOctave():%d", testCase.note.Name(), testCase.note.Octave().Name(), testCase.want, testCase.note.MIDINumber(), testCase.note.mustGetNoteNumberWithinOctave())
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
				note: newNoteWithOctave(C, octave.MustNewByNumber(-1)),
				want: 0,
			},
			{
				note: newNoteWithOctave(C, octave.MustNewByNumber(-1)).AlterDown(),
				want: 0,
			},
			{
				note: newNoteWithOctave(C, octave.MustNewByNumber(-1)).AlterDownBy(2),
				want: 0,
			},
			{
				note: newNoteWithOctave(C, octave.MustNewByNumber(-1)).AlterDownBy(3),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, octave.MustNewByNumber(-1)).AlterDown(),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, octave.MustNewByNumber(-1)).AlterDownBy(2),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, octave.MustNewByNumber(-1)).AlterDownBy(3),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP2, octave.MustNewByNumber(-1)).AlterDownBy(2),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP2, octave.MustNewByNumber(-1)).AlterDownBy(3),
				want: 0,
			},
			{
				note: newNoteWithOctave(CSHARP, octave.MustNewByNumber(-1)),
				want: 1,
			},
			{
				note: newNoteWithOctave(DFLAT, octave.MustNewByNumber(-1)),
				want: 1,
			},
			{
				note: newNoteWithOctave(CSHARP2, octave.MustNewByNumber(-1)).AlterDown(),
				want: 1,
			},
			{
				note: newNoteWithOctave(C, octave.MustNewByNumber(4)),
				want: 60,
			},
			{
				note: newNoteWithOctave(CSHARP, octave.MustNewByNumber(4)),
				want: 61,
			},
			{
				note: newNoteWithOctave(D, octave.MustNewByNumber(4)),
				want: 62,
			},
			{
				note: newNoteWithOctave(DSHARP, octave.MustNewByNumber(4)),
				want: 63,
			},
			{
				note: newNoteWithOctave(E, octave.MustNewByNumber(4)),
				want: 64,
			},
			{
				note: newNoteWithOctave(F, octave.MustNewByNumber(4)),
				want: 65,
			},
			{
				note: newNoteWithOctave(FSHARP, octave.MustNewByNumber(4)),
				want: 66,
			},
			{
				note: newNoteWithOctave(G, octave.MustNewByNumber(4)),
				want: 67,
			},
			{
				note: newNoteWithOctave(GSHARP, octave.MustNewByNumber(4)),
				want: 68,
			},
			{
				note: newNoteWithOctave(A, octave.MustNewByNumber(4)),
				want: 69,
			},
			{
				note: newNoteWithOctave(ASHARP, octave.MustNewByNumber(4)),
				want: 70,
			},
			{
				note: newNoteWithOctave(B, octave.MustNewByNumber(4)),
				want: 71,
			},
			{
				note: newNoteWithOctave(G, octave.MustNewByNumber(9)),
				want: 127,
			},
			{
				note: newNoteWithOctave(A, octave.MustNewByNumber(9)),
				want: 127,
			},
			{
				note: newNoteWithOctave(B, octave.MustNewByNumber(9)),
				want: 127,
			},
			{
				note: newNoteWithOctave(G, octave.MustNewByNumber(9)).AlterUp(),
				want: 127,
			},
			{
				note: newNoteWithOctave(G, octave.MustNewByNumber(9)).AlterUpBy(2),
				want: 127,
			},
			{
				note: newNoteWithOctave(BSHARP, octave.MustNewByNumber(9)).AlterUpBy(2),
				want: 127,
			},
		}

		for _, testCase := range testCases {
			assert.Equal(t, testCase.want, testCase.note.MIDINumber(),
				"note: %s, octave: %s, expected MIDI number: %d, actual: %d, mustGetNoteNumberWithinOctave():%d", testCase.note.Name(), testCase.note.Octave().Name(), testCase.want, testCase.note.MIDINumber(), testCase.note.mustGetNoteNumberWithinOctave())
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
			require.NoError(t, err)
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
			require.ErrorIs(t, err, testCase.want)
		}
	})
}
