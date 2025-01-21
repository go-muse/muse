package mode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/scale"
)

func TestModeTemplatesStore(t *testing.T) {
	mts := InitTemplatesStore()
	notes := scale.GetFullChromaticScale()
	var modeNormal *Mode
	var modeCustom *Mode
	var err error
	for modeName, modeTemplate := range mts {
		for _, note := range notes {
			modeNormal, err = MakeNewMode(modeName, note.Name())
			require.NoError(t, err)
			assert.Equal(t, modeTemplate.Length(), modeNormal.Length())
			modeCustom, err = MakeNewCustomMode(modeTemplate, note.Name().String(), modeName)
			require.NoError(t, err)
			assert.Equal(t, modeTemplate.Length(), modeCustom.Length())
			assert.True(t, modeNormal.IsEqual(modeCustom))
		}
	}
}

func TestFindModeTemplatesByPattern(t *testing.T) {
	mts := make(TemplatesStore)

	for i := uint8(1); i < 18; i++ {
		customModeTemplate := make(Template, i+1)
		customName := Name(fmt.Sprintf("%s %d", "custom abstract mode template", i))
		for j := uint8(0); j <= i; j++ {
			customModeTemplate[j] = halftone.HalfTones(j)
		}
		mts.AddTemplate(customName, customModeTemplate)
	}

	for customName, modeTemplate := range mts {
		testCases := generateAbstractPatterns(modeTemplate[1:])
		for _, testCase := range testCases {
			result := mts.FindModeTemplatesByPattern(testCase)
			var exist bool
			for modeName := range result {
				if modeName == customName {
					exist = true
				}
			}
			assert.True(t, exist, "customName: %s, testCase: %+v", customName, testCase)
		}
	}
}

func TestFindModeTemplatesByPatternCases(t *testing.T) {
	mts := InitTemplatesStore()
	modeTemplates := mts.FindModeTemplatesByPattern(TemplateNaturalMinor())
	if _, ok := modeTemplates[NameAeolian]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", NameAeolian)
	}
	if _, ok := modeTemplates[NameNaturalMinor]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", NameNaturalMinor)
	}

	modeTemplates = mts.FindModeTemplatesByPattern(Template{1, 2, 2, 1, 2, 2})
	if _, ok := modeTemplates[NameAeolian]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", NameAeolian)
	}
	if _, ok := modeTemplates[NameNaturalMinor]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", NameNaturalMinor)
	}

	modeTemplates = mts.FindModeTemplatesByPattern(Template{2})
	if _, ok := modeTemplates[NameAeolian]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", NameAeolian)
	}
	if _, ok := modeTemplates[NameNaturalMinor]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", NameNaturalMinor)
	}
}

func TestFindModeTemplatesByNotes(t *testing.T) {
	testCases := []struct {
		notes         note.Notes
		expectedModes []Name
	}{
		{
			notes:         note.MustNewNotesFromNoteNames(note.C, note.D, note.E, note.F, note.G, note.A, note.B),
			expectedModes: []Name{NameAeolian, NameIonian, NamePhrygian, NameLocrian, NameDorian, NameLydian, NameMixoLydian, NameNaturalMajor, NameNaturalMinor},
		},
	}

	mts := InitTemplatesStore()
	for _, testCase := range testCases {
		var result TemplatesWithPrime
		assert.NotPanics(t, func() { result = mts.FindModeTemplatesByNotes(testCase.notes) }, "testCase.notes: %+v", testCase.notes) //nolint:scopelint

		for _, modeName := range testCase.expectedModes {
			assert.True(t, result.Contains(modeName), "expected mode name: %s", modeName)
		}

		assert.Equal(t, len(result), len(testCase.expectedModes))
	}
}

func generateAbstractPatterns(modeTemplate []halftone.HalfTones) [][]halftone.HalfTones {
	var res [][]halftone.HalfTones
	for j := len(modeTemplate); j > 0; j-- {
		k := j
		for i := 0; i+j <= len(modeTemplate); i++ {
			res = append(res, modeTemplate[i:k])
			k++
		}
	}

	return res
}
