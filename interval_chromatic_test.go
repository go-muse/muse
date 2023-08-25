package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChromaticInterval_Names(t *testing.T) {
	testCases := map[HalfTones]*intervalNameExtended{
		IntervalHalfTones0:  IntervalPerfectUnison().names,
		IntervalHalfTones1:  IntervalMinorSecond().names,
		IntervalHalfTones2:  IntervalMajorSecond().names,
		IntervalHalfTones3:  IntervalMinorThird().names,
		IntervalHalfTones4:  IntervalMajorThird().names,
		IntervalHalfTones5:  IntervalPerfectFourth().names,
		IntervalHalfTones6:  IntervalTritone().names,
		IntervalHalfTones7:  IntervalPerfectFifth().names,
		IntervalHalfTones8:  IntervalMinorSixth().names,
		IntervalHalfTones9:  IntervalMajorSixth().names,
		IntervalHalfTones10: IntervalMinorSeventh().names,
		IntervalHalfTones11: IntervalMajorSeventh().names,
		IntervalHalfTones12: IntervalPerfectOctave().names,
		IntervalHalfTones13: IntervalMinorNinth().names,
		IntervalHalfTones14: IntervalMajorNinth().names,
		IntervalHalfTones15: IntervalMinorTenth().names,
		IntervalHalfTones16: IntervalMajorTenth().names,
		IntervalHalfTones17: IntervalPerfectEleventh().names,
		IntervalHalfTones18: IntervalOctaveWithTritone().names,
		IntervalHalfTones19: IntervalPerfectTwelfth().names,
		IntervalHalfTones20: IntervalMinorThirteenth().names,
		IntervalHalfTones21: IntervalMajorThirteenth().names,
		IntervalHalfTones22: IntervalMinorFourteenth().names,
		IntervalHalfTones23: IntervalMajorFourteenth().names,
		IntervalHalfTones24: IntervalPerfectFifteenth().names,
	}

	var intervalName, intervalShortName IntervalName
	var interval *IntervalChromatic
	var err error
	for halfTones, names := range testCases {
		interval, err = NewIntervalChromatic(halfTones)
		assert.NoError(t, err)
		intervalName, intervalShortName = interval.Name(), interval.ShortName()
		assert.Equal(t, names.name, intervalName)
		assert.Equal(t, names.shortName, intervalShortName)
	}
}

func TestChromaticInterval_HalfTones(t *testing.T) {
	expectedHalfTones := HalfTones(12)
	// Create a test interval with a known number of half tones
	testInterval := IntervalChromatic{
		halfTones: expectedHalfTones,
	}

	// Ensure that the HalfTones() method returns the expected number of half tones
	if testInterval.HalfTones() != expectedHalfTones {
		t.Errorf("HalfTones() returned %d, expected %d", testInterval.HalfTones(), expectedHalfTones)
	}
}

func TestMakeNoteByIntervalName(t *testing.T) {
	firstNote := MustNewNote(C)
	note, err := MakeNoteByIntervalName(firstNote, IntervalNameTritone)
	assert.NoError(t, err)
	assert.NotNil(t, note)
	assert.Equal(t, note.Name(), FSHARP)
}

func TestMakeDegreeByIntervalName(t *testing.T) {
	firstDegree := NewDegree(1, 0, nil, nil, MustNewNote(C), nil, nil)
	interval, err := NewIntervalChromatic(6)
	assert.NoError(t, err)
	secondDegree, err := MakeDegreeByIntervalName(firstDegree, interval.Name())
	assert.NoError(t, err)
	assert.NotNil(t, secondDegree)
	assert.Equal(t, secondDegree.Note().Name(), FSHARP)
	assert.Equal(t, secondDegree.Number(), firstDegree.Number()+1)
	assert.Equal(t, secondDegree.HalfTonesFromPrime(), interval.HalfTones())
}

func TestNewIntervalChromatic(t *testing.T) {
	testCases := map[HalfTones]*IntervalChromatic{
		0:  IntervalPerfectUnison(),
		1:  IntervalMinorSecond(),
		2:  IntervalMajorSecond(),
		3:  IntervalMinorThird(),
		4:  IntervalMajorThird(),
		5:  IntervalPerfectFourth(),
		6:  IntervalTritone(),
		7:  IntervalPerfectFifth(),
		8:  IntervalMinorSixth(),
		9:  IntervalMajorSixth(),
		10: IntervalMinorSeventh(),
		11: IntervalMajorSeventh(),
		12: IntervalPerfectOctave(),
		13: IntervalMinorNinth(),
		14: IntervalMajorNinth(),
		15: IntervalMinorTenth(),
		16: IntervalMajorTenth(),
		17: IntervalPerfectEleventh(),
		18: IntervalOctaveWithTritone(),
		19: IntervalPerfectTwelfth(),
		20: IntervalMinorThirteenth(),
		21: IntervalMajorThirteenth(),
		22: IntervalMinorFourteenth(),
		23: IntervalMajorFourteenth(),
		24: IntervalPerfectFifteenth(),
	}

	var result *IntervalChromatic
	var err error
	for halfTones, expected := range testCases {
		result, err = NewIntervalChromatic(halfTones)
		assert.NoError(t, err)
		assert.Equal(t, result, expected)
	}
}

func TestNewIntervalByHalfTonesAndDegrees(t *testing.T) {
	type testCase struct {
		halfTones HalfTones
		degrees   DegreeNum
		expected  *IntervalChromatic
	}

	testCases := []testCase{
		{
			halfTones: IntervalHalfTones0,
			degrees:   0,
			expected:  IntervalPerfectUnison(),
		},
		{
			halfTones: IntervalHalfTones0,
			degrees:   1,
			expected:  IntervalDiminishedSecond(),
		},
		{
			halfTones: IntervalHalfTones1,
			degrees:   0,
			expected:  IntervalAugmentedUnison(),
		},
		{
			halfTones: IntervalHalfTones1,
			degrees:   1,
			expected:  IntervalMinorSecond(),
		},
		{
			halfTones: IntervalHalfTones2,
			degrees:   1,
			expected:  IntervalMajorSecond(),
		},
		{
			halfTones: IntervalHalfTones2,
			degrees:   2,
			expected:  IntervalDiminishedThird(),
		},
		{
			halfTones: IntervalHalfTones3,
			degrees:   1,
			expected:  IntervalAugmentedSecond(),
		},
		{
			halfTones: IntervalHalfTones3,
			degrees:   2,
			expected:  IntervalMinorThird(),
		},
		{
			halfTones: IntervalHalfTones4,
			degrees:   2,
			expected:  IntervalMajorThird(),
		},
		{
			halfTones: IntervalHalfTones4,
			degrees:   3,
			expected:  IntervalDiminishedFourth(),
		},
		{
			halfTones: IntervalHalfTones5,
			degrees:   2,
			expected:  IntervalAugmentedThird(),
		},
		{
			halfTones: IntervalHalfTones5,
			degrees:   3,
			expected:  IntervalPerfectFourth(),
		},
		{
			halfTones: IntervalHalfTones6,
			degrees:   3,
			expected:  IntervalAugmentedFourth(),
		},
		{
			halfTones: IntervalHalfTones6,
			degrees:   4,
			expected:  IntervalDiminishedFifth(),
		},
		{
			halfTones: IntervalHalfTones7,
			degrees:   4,
			expected:  IntervalPerfectFifth(),
		},
		{
			halfTones: IntervalHalfTones7,
			degrees:   5,
			expected:  IntervalDiminishedSixth(),
		},
		{
			halfTones: IntervalHalfTones8,
			degrees:   4,
			expected:  IntervalAugmentedFifth(),
		},
		{
			halfTones: IntervalHalfTones8,
			degrees:   5,
			expected:  IntervalMinorSixth(),
		},
		{
			halfTones: IntervalHalfTones9,
			degrees:   5,
			expected:  IntervalMajorSixth(),
		},
		{
			halfTones: IntervalHalfTones9,
			degrees:   6,
			expected:  IntervalDiminishedSeventh(),
		},
		{
			halfTones: IntervalHalfTones10,
			degrees:   5,
			expected:  IntervalAugmentedSixth(),
		},
		{
			halfTones: IntervalHalfTones10,
			degrees:   6,
			expected:  IntervalMinorSeventh(),
		},
		{
			halfTones: IntervalHalfTones11,
			degrees:   6,
			expected:  IntervalMajorSeventh(),
		},
		{
			halfTones: IntervalHalfTones11,
			degrees:   7,
			expected:  IntervalDiminishedOctave(),
		},
		{
			halfTones: IntervalHalfTones12,
			degrees:   6,
			expected:  IntervalAugmentedSeventh(),
		},
		{
			halfTones: IntervalHalfTones12,
			degrees:   7,
			expected:  IntervalPerfectOctave(),
		},
	}

	var result *IntervalChromatic
	var err error
	for _, testCase := range testCases {
		result, err = NewIntervalByHalfTonesAndDegrees(testCase.halfTones, testCase.degrees)
		assert.NoError(t, err)
		assert.Equal(t, result, testCase.expected)
	}
}
