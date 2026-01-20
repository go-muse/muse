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
		nextDegreeNode     *Node
		halfTonesFromPrime halftone.HalfTones
		expectedName       CharacteristicName
	}

	testCases := []testCase{
		{
			degreeNum:          1,
			nextDegreeNode:     &Node{degree: Degree{note: note.DFLAT.NewNote()}},
			halfTonesFromPrime: 1,
			expectedName:       CharacteristicAug,
		},
		{
			degreeNum:          1,
			nextDegreeNode:     &Node{degree: Degree{note: note.CSHARP.NewNote()}},
			halfTonesFromPrime: 1,
			expectedName:       CharacteristicAug,
		},
		{
			degreeNum:          1,
			nextDegreeNode:     &Node{degree: Degree{note: note.D.NewNote()}},
			halfTonesFromPrime: 2,
			expectedName:       Characteristic2xAug,
		},
		{
			degreeNum:          1,
			nextDegreeNode:     &Node{degree: Degree{note: note.EFLAT.NewNote()}},
			halfTonesFromPrime: 3,
			expectedName:       Characteristic3xAug,
		},
		{
			degreeNum:          1,
			nextDegreeNode:     &Node{degree: Degree{note: note.DSHARP.NewNote()}},
			halfTonesFromPrime: 3,
			expectedName:       Characteristic3xAug,
		},

		// TODO: more cases
	}

	for _, tc := range testCases {
		t.Run(string(tc.expectedName), func(t *testing.T) {
			mc, err := CalculateRelativeMC(tc.degreeNum, tc.nextDegreeNode, tc.halfTonesFromPrime)
			require.NoError(t, err)
			require.NotNil(t, mc)
			assert.Equal(t, tc.expectedName, mc.Name())
		})
	}
}
