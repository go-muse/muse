package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTemplateByModeNames(t *testing.T) {
	testCases := []struct {
		modeName ModeName
		expected ModeTemplate
	}{
		{
			modeName: ModeNameNaturalMinor,
			expected: TemplateNaturalMinor(),
		},
		{
			modeName: ModeNameHarmonicMinor,
			expected: TemplateHarmonicMinor(),
		},
		{
			modeName: ModeNameMelodicMinor,
			expected: TemplateMelodicMinor(),
		},
		{
			modeName: ModeNameNaturalMajor,
			expected: TemplateNaturalMajor(),
		},
		{
			modeName: ModeNameHarmonicMajor,
			expected: TemplateHarmonicMajor(),
		},
		{
			modeName: ModeNameMelodicMajor,
			expected: TemplateMelodicMajor(),
		},
	}

	var modeTemplate ModeTemplate
	var err error
	for _, testCase := range testCases {
		modeTemplate, err = GetTemplateByModeName(testCase.modeName)
		assert.NoError(t, err)
		assert.NoError(t, modeTemplate.Validate())
		assert.Equal(t, testCase.expected, modeTemplate)
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		modeTemplate *ModeTemplate
		isValid      bool
	}{
		{
			&ModeTemplate{1, 2, 3, 5, 1}, // 12 is ok
			true,
		}, {
			&ModeTemplate{3, 3, 3, 3}, // 12 is ok
			true,
		}, {
			&ModeTemplate{12}, // 12 is ok
			true,
		}, {
			&ModeTemplate{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // 12 is ok
			true,
		}, {
			&ModeTemplate{1, 2, 3, 5, 2}, // 13 is not ok
			false,
		}, {
			&ModeTemplate{1, 2, 3, 3, 2}, // 11 is not ok
			false,
		}, {
			&ModeTemplate{0, 1, 2, 3, 3, 2}, // tonal centre must not to be zero
			false,
		}, {
			&ModeTemplate{1, 2, 0, 3, 2}, // intervals must not to be zero
			false,
		},
	}

	var err error
	for _, testCase := range testCases {
		err = testCase.modeTemplate.Validate()
		switch testCase.isValid {
		case true:
			assert.NoErrorf(t, err, "expected validation result: %t, template: %+v", testCase.isValid, *testCase.modeTemplate)
		case false:
			assert.Errorf(t, err, "expected validation result: %t, template: %+v", testCase.isValid, *testCase.modeTemplate)
		}
	}
}
