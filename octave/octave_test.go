package octave

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewOctave(t *testing.T) {
	t.Run("TestNewOctave: positive cases", func(t *testing.T) {
		type testCase struct {
			*Octave
		}

		testCases := []testCase{
			{&Octave{-1, NameSubSubContraOctave}},
			{&Octave{0, NameSubContraOctave}},
			{&Octave{1, NameContraOctave}},
			{&Octave{2, NameGreatOctave}},
			{&Octave{3, SmallOctave}},
			{&Octave{4, NameFirstOctave}},
			{&Octave{5, NameSecondOctave}},
			{&Octave{6, NameThirdOctave}},
			{&Octave{7, NameFourthOctave}},
			{&Octave{8, NameFifthOctave}},
			{&Octave{9, NameSixthOctave}},
		}

		var octave *Octave
		var err error
		for _, testCase := range testCases {
			octave, err = NewByNumber(testCase.Octave.number)
			require.NoError(t, err)
			assert.Equal(t, testCase.Octave.Name(), octave.Name())
			assert.Equal(t, testCase.Octave.Number(), octave.Number())
		}
	})

	t.Run("TestNewOctave: negative cases", func(t *testing.T) {
		testCases := []Number{
			-2, -3, -4, -5, -6, -7, -8, -9, -10, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		}

		var octave *Octave
		var err error
		for _, testCase := range testCases {
			octave, err = NewByNumber(testCase)
			require.ErrorIs(t, err, ErrOctaveNumberUnknown)
			assert.Nil(t, octave)
		}
	})
}

func TestOctaveName(t *testing.T) {
	// setup: create an Octave with a known name
	expectedName := NameContraOctave
	octave := &Octave{name: expectedName}

	// execute the Name() method
	actualName := octave.Name()

	// assert that the returned name matches the expected name
	assert.Equal(t, expectedName, actualName, "expected octave name: %s, actual: %s", expectedName, actualName)
}

func TestOctaveNumber(t *testing.T) {
	// setup: create an Octave with a known number
	expectedNumber := Number3
	octave := &Octave{number: expectedNumber}

	// execute the Number() method
	actualNumber := octave.Number()

	// assert that the returned number matches the expected number
	assert.Equal(t, expectedNumber, actualNumber, "expected octave number: %s, actual: %s", expectedNumber, actualNumber)
}

func TestOctaveIsEqual(t *testing.T) {
	t.Run("TestOctaveIsEqual: positive cases", func(t *testing.T) {
		type testCase struct {
			octave1 *Octave
			octave2 *Octave
		}

		testCases := []testCase{
			{&Octave{-1, NameSubSubContraOctave}, &Octave{-1, NameSubSubContraOctave}},
			{&Octave{0, NameSubContraOctave}, &Octave{0, NameSubContraOctave}},
			{&Octave{1, NameContraOctave}, &Octave{1, NameContraOctave}},
			{&Octave{2, NameGreatOctave}, &Octave{2, NameGreatOctave}},
			{&Octave{3, SmallOctave}, &Octave{3, SmallOctave}},
			{&Octave{4, NameFirstOctave}, &Octave{4, NameFirstOctave}},
			{&Octave{5, NameSecondOctave}, &Octave{5, NameSecondOctave}},
			{&Octave{6, NameThirdOctave}, &Octave{6, NameThirdOctave}},
			{&Octave{7, NameFourthOctave}, &Octave{7, NameFourthOctave}},
			{&Octave{8, NameFifthOctave}, &Octave{8, NameFifthOctave}},
			{&Octave{9, NameSixthOctave}, &Octave{9, NameSixthOctave}},
		}

		for _, testCase := range testCases {
			assert.True(t, testCase.octave1.IsEqual(testCase.octave2))
		}
	})

	t.Run("TestOctaveIsEqual: negative cases", func(t *testing.T) {
		type testCase struct {
			octave1 *Octave
			octave2 *Octave
		}

		testCases := []testCase{
			{&Octave{-1, NameSubSubContraOctave}, &Octave{1, NameContraOctave}},
			{&Octave{0, NameSubContraOctave}, &Octave{5, NameSecondOctave}},
			{&Octave{1, NameContraOctave}, &Octave{-1, NameSubSubContraOctave}},
		}

		for _, testCase := range testCases {
			assert.False(t, testCase.octave1.IsEqual(testCase.octave2))
		}
	})
}

func TestOctaveValidate(t *testing.T) {
	testCases := []struct {
		octave *Octave
		want   error
	}{
		{octave: &Octave{number: -5, name: ""}, want: ErrOctaveNumberUnknown},
		{octave: &Octave{number: -4, name: ""}, want: ErrOctaveNumberUnknown},
		{octave: &Octave{number: -3, name: ""}, want: ErrOctaveNumberUnknown},
		{octave: &Octave{number: -2, name: ""}, want: ErrOctaveNumberUnknown},
		{octave: &Octave{number: -1, name: ""}, want: nil},
		{octave: &Octave{number: 0, name: ""}, want: nil},
		{octave: &Octave{number: 1, name: ""}, want: nil},
		{octave: &Octave{number: 2, name: ""}, want: nil},
		{octave: &Octave{number: 3, name: ""}, want: nil},
		{octave: &Octave{number: 4, name: ""}, want: nil},
		{octave: &Octave{number: 5, name: ""}, want: nil},
		{octave: &Octave{number: 6, name: ""}, want: nil},
		{octave: &Octave{number: 7, name: ""}, want: nil},
		{octave: &Octave{number: 8, name: ""}, want: nil},
		{octave: &Octave{number: 9, name: ""}, want: nil},
		{octave: &Octave{number: 10, name: ""}, want: ErrOctaveNumberUnknown},
		{octave: &Octave{number: 11, name: ""}, want: ErrOctaveNumberUnknown},
		{octave: &Octave{number: 12, name: ""}, want: ErrOctaveNumberUnknown},
	}

	for _, testCase := range testCases {
		require.ErrorIs(t, testCase.octave.Validate(), testCase.want)
	}
}
