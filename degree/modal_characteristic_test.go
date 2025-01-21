package degree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

func TestCalculateRelativeMC(t *testing.T) {
	type testCase struct {
		degreeNum          Number
		nextDegree         *Degree
		halfTonesFromPrime halftone.HalfTones
		expectedName       CharacteristicName
	}

	testCases := []testCase{
		{
			degreeNum:          1,
			nextDegree:         &Degree{note: note.MustNewNote(note.DFLAT)},
			halfTonesFromPrime: 1,
			expectedName:       CharacteristicAug,
		},
		{
			degreeNum:          1,
			nextDegree:         &Degree{note: note.MustNewNote(note.CSHARP)},
			halfTonesFromPrime: 1,
			expectedName:       CharacteristicAug,
		},
		{
			degreeNum:          1,
			nextDegree:         &Degree{note: note.MustNewNote(note.D)},
			halfTonesFromPrime: 2,
			expectedName:       Characteristic2xAug,
		},
		{
			degreeNum:          1,
			nextDegree:         &Degree{note: note.MustNewNote(note.EFLAT)},
			halfTonesFromPrime: 3,
			expectedName:       Characteristic3xAug,
		},
		{
			degreeNum:          1,
			nextDegree:         &Degree{note: note.MustNewNote(note.DSHARP)},
			halfTonesFromPrime: 3,
			expectedName:       Characteristic3xAug,
		},

		// TODO: more cases
	}

	for _, tc := range testCases {
		t.Run(string(tc.expectedName), func(t *testing.T) {
			mc, err := CalculateRelativeMC(tc.degreeNum, tc.nextDegree, tc.halfTonesFromPrime)
			require.NoError(t, err)
			require.NotNil(t, mc)
			assert.Equal(t, tc.expectedName, mc.Name())
		})
	}
}
