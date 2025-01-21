package mode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
)

func Test_IsDiatonic(t *testing.T) {
	testCases := []struct {
		modeTemplate Template
		isDiatonic   bool
	}{
		// Tonal modes

		{
			modeTemplate: TemplateNaturalMinor(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateMelodicMinor(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateHarmonicMinor(),
			isDiatonic:   false,
		},
		{
			modeTemplate: TemplateNaturalMajor(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateMelodicMajor(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateHarmonicMajor(),
			isDiatonic:   false,
		},

		// Modes of the Major scale

		{
			modeTemplate: TemplateIonian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateDorian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateAeolian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateLydian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateMixoLydian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplatePhrygian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateLocrian(),
			isDiatonic:   true,
		},

		// Modes Of The Melodic Minor scale

		{
			modeTemplate: TemplateIonianFlat3(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplatePhrygoDorian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateLydianAugmented(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateLydianDominant(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateIonianAeolian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateAeolianLydian(),
			isDiatonic:   true,
		},
		{
			modeTemplate: TemplateSuperLocrian(),
			isDiatonic:   true,
		},

		// Modes of the Harmonic Minor scale

		{
			modeTemplate: TemplateAeolianRais7(),
			isDiatonic:   false,
		},
		{
			modeTemplate: TemplateLocrianRais6(),
			isDiatonic:   false,
		},
		{
			modeTemplate: TemplateIonianRais5(),
			isDiatonic:   false,
		},
		{
			modeTemplate: TemplateUkrainianDorian(),
			isDiatonic:   false,
		},
		{
			modeTemplate: TemplatePhrygianDominant(),
			isDiatonic:   false,
		},
		{
			modeTemplate: TemplateLydianRais9(),
			isDiatonic:   false,
		},
		{
			modeTemplate: TemplateUltraLocrian(),
			isDiatonic:   false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.isDiatonic, testCase.modeTemplate.IsDiatonic(), "expects isDiatonic: %t, template: %+v", testCase.isDiatonic, testCase.modeTemplate)
	}
}

// TODO: make more intellectual test.
func TestRearrangeFromDegree(t *testing.T) {
	modeTemplate := Template{1, 2, 3, 4, 5}
	result := modeTemplate.RearrangeFromDegree(3)
	assert.Equal(t, result[0], modeTemplate[2])
	assert.Equal(t, result[1], modeTemplate[3])
	assert.Equal(t, result[2], modeTemplate[4])
	assert.Equal(t, result[3], modeTemplate[0])
	assert.Equal(t, result[4], modeTemplate[1])
}

// TODO: make more intellectual test.
func TestGetHalftonesByDegreeNum(t *testing.T) {
	modeTemplate := Template{1, 2, 3, 4, 5}
	result := modeTemplate.GetHalftonesByDegreeNum(3)
	assert.Equal(t, halftone.HalfTones(6), result)
}

// TODO: make more intellectual test.
func TestIterateModeTemplateOneRound(t *testing.T) {
	modeTemplate := Template{1, 2, 3, 4, 5}
	var halfTonesFromPrimeExpected halftone.HalfTones
	degreeNumExpected := degree.Number(1)
	for f := range modeTemplate.IterateOneRound(true) {
		degreeNumExpected++
		halfTonesFromPrimeExpected += modeTemplate[int(degreeNumExpected)-2]
		halfTonesExpected := modeTemplate[int(degreeNumExpected)-2]

		degreeNum, halfTones, halfTonesFromPrime := f()

		assert.Equal(t, degreeNumExpected, degreeNum, "degreeNumExpected: %d, degreeNum: %d", degreeNumExpected, degreeNum)
		assert.Equal(t, halfTonesExpected, halfTones, "halfTonesExpected: %d, halfTones: %d", halfTonesExpected, halfTones)
		assert.Equal(t, halfTonesFromPrimeExpected, halfTonesFromPrime, "halfTonesFromPrimeExpected: %d, halfTonesFromPrime: %d", halfTonesFromPrimeExpected, halfTonesFromPrime)
	}
}

func TestModeTemplateIterateOneRound(t *testing.T) {
	m := Template{2, 2, 1, 2, 2, 2, 1}

	c := m.IterateOneRound(false)
	var expected []TemplateIteratorResult

	expected = []TemplateIteratorResult{
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 2, 2, 2 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 3, 2, 4 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 4, 1, 5 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 5, 2, 7 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 6, 2, 9 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 7, 2, 11 },
	}

	var res TemplateIteratorResult
	var expDegreenum, resDegreenum degree.Number
	var expHalfTones, resHalfTones, expHalfTonesFromPrime, resHalfTonesFromPrime halftone.HalfTones
	for i, exp := range expected {
		res = <-c
		expDegreenum, expHalfTones, expHalfTonesFromPrime = exp()
		resDegreenum, resHalfTones, resHalfTonesFromPrime = res()
		assert.Equal(t, expDegreenum, resDegreenum, "on iteration %d expected degreeNum: %d but got: %d", i, expDegreenum, resDegreenum)
		assert.Equal(t, expHalfTones, resHalfTones, "on iteration %d expected halfTones: %d but got: %d", i, expHalfTones, resHalfTones)
		assert.Equal(t, expHalfTonesFromPrime, resHalfTonesFromPrime, "on iteration %d expected halfTones from prime: %d but got: %d", i, expHalfTonesFromPrime, resHalfTonesFromPrime)
	}

	_, isOpen := <-c
	assert.False(t, isOpen)

	expected = []TemplateIteratorResult{
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 2, 2, 2 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 3, 2, 4 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 4, 1, 5 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 5, 2, 7 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 6, 2, 9 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 7, 2, 11 },
		func() (degree.Number, halftone.HalfTones, halftone.HalfTones) { return 8, 1, 12 },
	}
	c = m.IterateOneRound(true)

	for i, exp := range expected {
		res = <-c
		expDegreenum, expHalfTones, expHalfTonesFromPrime = exp()
		resDegreenum, resHalfTones, resHalfTonesFromPrime = res()
		assert.Equal(t, expDegreenum, resDegreenum, "on iteration %d expected degreeNum: %d but got: %d", i, expDegreenum, resDegreenum)
		assert.Equal(t, expHalfTones, resHalfTones, "on iteration %d expected halfTones: %d but got: %d", i, expHalfTones, resHalfTones)
		assert.Equal(t, expHalfTonesFromPrime, resHalfTonesFromPrime, "on iteration %d expected halfTones from prime: %d but got: %d", i, expHalfTonesFromPrime, resHalfTonesFromPrime)
	}

	_, isOpen = <-c
	assert.False(t, isOpen)
}
