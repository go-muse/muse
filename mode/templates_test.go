package mode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTemplateByNames(t *testing.T) {
	testCases := []struct {
		modeName Name
		expected Template
	}{
		{
			modeName: NameNaturalMinor,
			expected: TemplateNaturalMinor(),
		},
		{
			modeName: NameHarmonicMinor,
			expected: TemplateHarmonicMinor(),
		},
		{
			modeName: NameMelodicMinor,
			expected: TemplateMelodicMinor(),
		},
		{
			modeName: NameNaturalMajor,
			expected: TemplateNaturalMajor(),
		},
		{
			modeName: NameHarmonicMajor,
			expected: TemplateHarmonicMajor(),
		},
		{
			modeName: NameMelodicMajor,
			expected: TemplateMelodicMajor(),
		},
	}

	var modeTemplate Template
	var err error
	for _, testCase := range testCases {
		modeTemplate, err = GetTemplateByName(testCase.modeName)
		require.NoError(t, err)
		require.NoError(t, modeTemplate.Validate())
		assert.Equal(t, testCase.expected, modeTemplate)
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		modeTemplate *Template
		isValid      bool
	}{
		{
			&Template{1, 2, 3, 5, 1}, // 12 is ok
			true,
		}, {
			&Template{3, 3, 3, 3}, // 12 is ok
			true,
		}, {
			&Template{12}, // 12 is ok
			true,
		}, {
			&Template{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // 12 is ok
			true,
		}, {
			&Template{1, 2, 3, 5, 2}, // 13 is not ok
			false,
		}, {
			&Template{1, 2, 3, 3, 2}, // 11 is not ok
			false,
		}, {
			&Template{0, 1, 2, 3, 3, 2}, // tonal centre must not to be zero
			false,
		}, {
			&Template{1, 2, 0, 3, 2}, // intervals must not to be zero
			false,
		},
	}

	var err error
	for _, testCase := range testCases {
		err = testCase.modeTemplate.Validate()
		switch testCase.isValid {
		case true:
			require.NoErrorf(t, err, "expected validation result: %t, template: %+v", testCase.isValid, *testCase.modeTemplate)
		case false:
			assert.Errorf(t, err, "expected validation result: %t, template: %+v", testCase.isValid, *testCase.modeTemplate)
		}
	}
}
