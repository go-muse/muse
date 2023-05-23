package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOctave(t *testing.T) {
	t.Run("TestNewOctave: positive cases", func(t *testing.T) {
		type testCase struct {
			*Octave
		}

		testCases := []testCase{
			{&Octave{-1, OctaveNameSubSubContraOctave}},
			{&Octave{0, OctaveNameSubContraOctave}},
			{&Octave{1, OctaveNameContraOctave}},
			{&Octave{2, OctaveNameGreatOctave}},
			{&Octave{3, OctaveNameSmallOctave}},
			{&Octave{4, OctaveNameFirstOctave}},
			{&Octave{5, OctaveNameSecondOctave}},
			{&Octave{6, OctaveNameThirdOctave}},
			{&Octave{7, OctaveNameFourthOctave}},
			{&Octave{8, OctaveNameFifthOctave}},
			{&Octave{9, OctaveNameSixthOctave}},
		}

		var octave *Octave
		var err error
		for _, testCase := range testCases {
			octave, err = NewOctave(testCase.Octave.number)
			assert.NoError(t, err)
			assert.Equal(t, testCase.Octave.Name(), octave.Name())
			assert.Equal(t, testCase.Octave.Number(), octave.Number())
		}
	})

	t.Run("TestNewOctave: negative cases", func(t *testing.T) {
		testCases := []OctaveNumber{
			-2, -3, -4, -5, -6, -7, -8, -9, -10, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		}

		var octave *Octave
		var err error
		for _, testCase := range testCases {
			octave, err = NewOctave(testCase)
			assert.ErrorIs(t, err, ErrOctaveNumberUnknown)
			assert.Nil(t, octave)
		}
	})
}

func TestOctaveName(t *testing.T) {
	// setup: create an Octave with a known name
	expectedName := OctaveNameContraOctave
	octave := &Octave{name: expectedName}

	// execute the Name() method
	actualName := octave.Name()

	// assert that the returned name matches the expected name
	assert.Equal(t, expectedName, actualName, "expected octave name: %s, actual: %s", expectedName, actualName)
}

func TestOctaveNumber(t *testing.T) {
	// setup: create an Octave with a known number
	expectedNumber := OctaveNumber3
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
			{&Octave{-1, OctaveNameSubSubContraOctave}, &Octave{-1, OctaveNameSubSubContraOctave}},
			{&Octave{0, OctaveNameSubContraOctave}, &Octave{0, OctaveNameSubContraOctave}},
			{&Octave{1, OctaveNameContraOctave}, &Octave{1, OctaveNameContraOctave}},
			{&Octave{2, OctaveNameGreatOctave}, &Octave{2, OctaveNameGreatOctave}},
			{&Octave{3, OctaveNameSmallOctave}, &Octave{3, OctaveNameSmallOctave}},
			{&Octave{4, OctaveNameFirstOctave}, &Octave{4, OctaveNameFirstOctave}},
			{&Octave{5, OctaveNameSecondOctave}, &Octave{5, OctaveNameSecondOctave}},
			{&Octave{6, OctaveNameThirdOctave}, &Octave{6, OctaveNameThirdOctave}},
			{&Octave{7, OctaveNameFourthOctave}, &Octave{7, OctaveNameFourthOctave}},
			{&Octave{8, OctaveNameFifthOctave}, &Octave{8, OctaveNameFifthOctave}},
			{&Octave{9, OctaveNameSixthOctave}, &Octave{9, OctaveNameSixthOctave}},
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
			{&Octave{-1, OctaveNameSubSubContraOctave}, &Octave{1, OctaveNameContraOctave}},
			{&Octave{0, OctaveNameSubContraOctave}, &Octave{5, OctaveNameSecondOctave}},
			{&Octave{1, OctaveNameContraOctave}, &Octave{-1, OctaveNameSubSubContraOctave}},
		}

		for _, testCase := range testCases {
			assert.False(t, testCase.octave1.IsEqual(testCase.octave2))
		}
	})
}

func TestOctaveSetToNote(t *testing.T) {
	expectedOctave := MustNewOctave(OctaveNumberDefault)

	t.Run("TestOctaveSetToNote: setting octave to the note without octave", func(t *testing.T) {
		// create a note without octave
		note1 := newNote(C)
		// set the octave no the note
		expectedOctave.SetToNote(note1)
		// check that they are the same
		assert.True(t, expectedOctave.IsEqual(note1.Octave()))
	})

	t.Run("TestOctaveSetToNote: setting octave to the note that already has an octave", func(t *testing.T) {
		// create a note with an octave
		note1 := newNoteWithOctave(C, MustNewOctave(OctaveNumber9))
		// set new octave no the note
		expectedOctave.SetToNote(note1)
		// check that they are the same
		assert.True(t, expectedOctave.IsEqual(note1.Octave()))
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
		assert.ErrorIs(t, testCase.octave.Validate(), testCase.want)
	}
}
