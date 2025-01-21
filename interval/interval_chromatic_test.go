package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

func TestChromaticInterval_Names(t *testing.T) {
	testCases := map[halftone.HalfTones]*nameExtended{
		HalfTones0:  PerfectUnison().names,
		HalfTones1:  MinorSecond().names,
		HalfTones2:  MajorSecond().names,
		HalfTones3:  MinorThird().names,
		HalfTones4:  MajorThird().names,
		HalfTones5:  PerfectFourth().names,
		HalfTones6:  Tritone().names,
		HalfTones7:  PerfectFifth().names,
		HalfTones8:  MinorSixth().names,
		HalfTones9:  MajorSixth().names,
		HalfTones10: MinorSeventh().names,
		HalfTones11: MajorSeventh().names,
		HalfTones12: PerfectOctave().names,
		HalfTones13: MinorNinth().names,
		HalfTones14: MajorNinth().names,
		HalfTones15: MinorTenth().names,
		HalfTones16: MajorTenth().names,
		HalfTones17: PerfectEleventh().names,
		HalfTones18: OctaveWithTritone().names,
		HalfTones19: PerfectTwelfth().names,
		HalfTones20: MinorThirteenth().names,
		HalfTones21: MajorThirteenth().names,
		HalfTones22: MinorFourteenth().names,
		HalfTones23: MajorFourteenth().names,
		HalfTones24: PerfectFifteenth().names,
	}

	var intervalName, intervalShortName Name
	var interval *Chromatic
	var err error
	for halfTones, names := range testCases {
		interval, err = NewChromatic(halfTones)
		require.NoError(t, err)
		intervalName, intervalShortName = interval.Name(), interval.ShortName()
		assert.Equal(t, names.name, intervalName)
		assert.Equal(t, names.shortName, intervalShortName)
	}
}

func TestChromaticInterval_HalfTones(t *testing.T) {
	expectedHalfTones := halftone.HalfTones(12)
	// Create a test interval with a known number of halftone
	testInterval := Chromatic{
		halfTones: expectedHalfTones,
	}

	// Ensure that the HalfTones() method returns the expected number of halftone
	if testInterval.HalfTones() != expectedHalfTones {
		t.Errorf("HalfTones() returned %d, expected %d", testInterval.HalfTones(), expectedHalfTones)
	}
}

func TestMakeNoteByName(t *testing.T) {
	firstNote := note.MustNewNote(note.C)
	n, err := MakeNoteByName(firstNote, NameTritone)
	require.NoError(t, err)
	assert.NotNil(t, n)
	assert.Equal(t, note.FSHARP, n.Name())
}

func TestMakeDegreeByName(t *testing.T) {
	firstDegree := degree.New(1, 0, nil, nil, note.MustNewNote(note.C), nil, nil)
	interval, err := NewChromatic(6)
	require.NoError(t, err)
	secondDegree, err := MakeDegreeByName(firstDegree, interval.Name())
	require.NoError(t, err)
	assert.NotNil(t, secondDegree)
	assert.Equal(t, note.FSHARP, secondDegree.Note().Name())
	assert.Equal(t, firstDegree.Number()+1, secondDegree.Number())
	assert.Equal(t, interval.HalfTones(), secondDegree.HalfTonesFromPrime())
}

func TestNewIntervalChromatic(t *testing.T) {
	testCases := map[halftone.HalfTones]*Chromatic{
		0:  PerfectUnison(),
		1:  MinorSecond(),
		2:  MajorSecond(),
		3:  MinorThird(),
		4:  MajorThird(),
		5:  PerfectFourth(),
		6:  Tritone(),
		7:  PerfectFifth(),
		8:  MinorSixth(),
		9:  MajorSixth(),
		10: MinorSeventh(),
		11: MajorSeventh(),
		12: PerfectOctave(),
		13: MinorNinth(),
		14: MajorNinth(),
		15: MinorTenth(),
		16: MajorTenth(),
		17: PerfectEleventh(),
		18: OctaveWithTritone(),
		19: PerfectTwelfth(),
		20: MinorThirteenth(),
		21: MajorThirteenth(),
		22: MinorFourteenth(),
		23: MajorFourteenth(),
		24: PerfectFifteenth(),
	}

	var result *Chromatic
	var err error
	for halfTones, expected := range testCases {
		result, err = NewChromatic(halfTones)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestNewIntervalByHalfTonesAndDegrees(t *testing.T) {
	type testCase struct {
		halfTones halftone.HalfTones
		degrees   degree.Number
		expected  *Chromatic
	}

	testCases := []testCase{
		{
			halfTones: HalfTones0,
			degrees:   0,
			expected:  PerfectUnison(),
		},
		{
			halfTones: HalfTones0,
			degrees:   1,
			expected:  DiminishedSecond(),
		},
		{
			halfTones: HalfTones1,
			degrees:   0,
			expected:  AugmentedUnison(),
		},
		{
			halfTones: HalfTones1,
			degrees:   1,
			expected:  MinorSecond(),
		},
		{
			halfTones: HalfTones2,
			degrees:   1,
			expected:  MajorSecond(),
		},
		{
			halfTones: HalfTones2,
			degrees:   2,
			expected:  DiminishedThird(),
		},
		{
			halfTones: HalfTones3,
			degrees:   1,
			expected:  AugmentedSecond(),
		},
		{
			halfTones: HalfTones3,
			degrees:   2,
			expected:  MinorThird(),
		},
		{
			halfTones: HalfTones4,
			degrees:   2,
			expected:  MajorThird(),
		},
		{
			halfTones: HalfTones4,
			degrees:   3,
			expected:  DiminishedFourth(),
		},
		{
			halfTones: HalfTones5,
			degrees:   2,
			expected:  AugmentedThird(),
		},
		{
			halfTones: HalfTones5,
			degrees:   3,
			expected:  PerfectFourth(),
		},
		{
			halfTones: HalfTones6,
			degrees:   3,
			expected:  AugmentedFourth(),
		},
		{
			halfTones: HalfTones6,
			degrees:   4,
			expected:  DiminishedFifth(),
		},
		{
			halfTones: HalfTones7,
			degrees:   4,
			expected:  PerfectFifth(),
		},
		{
			halfTones: HalfTones7,
			degrees:   5,
			expected:  DiminishedSixth(),
		},
		{
			halfTones: HalfTones8,
			degrees:   4,
			expected:  AugmentedFifth(),
		},
		{
			halfTones: HalfTones8,
			degrees:   5,
			expected:  MinorSixth(),
		},
		{
			halfTones: HalfTones9,
			degrees:   5,
			expected:  MajorSixth(),
		},
		{
			halfTones: HalfTones9,
			degrees:   6,
			expected:  DiminishedSeventh(),
		},
		{
			halfTones: HalfTones10,
			degrees:   5,
			expected:  AugmentedSixth(),
		},
		{
			halfTones: HalfTones10,
			degrees:   6,
			expected:  MinorSeventh(),
		},
		{
			halfTones: HalfTones11,
			degrees:   6,
			expected:  MajorSeventh(),
		},
		{
			halfTones: HalfTones11,
			degrees:   7,
			expected:  DiminishedOctave(),
		},
		{
			halfTones: HalfTones12,
			degrees:   6,
			expected:  AugmentedSeventh(),
		},
		{
			halfTones: HalfTones12,
			degrees:   7,
			expected:  PerfectOctave(),
		},
	}

	var result *Chromatic
	var err error
	for _, testCase := range testCases {
		result, err = NewIntervalByHalfTonesAndDegrees(testCase.halfTones, testCase.degrees)
		require.NoError(t, err)
		assert.Equal(t, testCase.expected, result)
	}
}
