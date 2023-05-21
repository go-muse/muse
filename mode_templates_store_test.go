package muse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModeTemplatesStore(t *testing.T) {
	mts := InitModeTemplatesStore()
	notes := GetFullChromaticScale()
	var modeNormal *Mode
	var modeCustom *Mode
	var err error
	for modeName, modeTemplate := range mts {
		for _, note := range notes {
			modeNormal, err = MakeNewMode(modeName, note.Name())
			assert.NoError(t, err)
			assert.Equal(t, modeTemplate.Length(), modeNormal.Length())
			modeCustom, err = MakeNewCustomMode(modeTemplate, note.Name().String(), modeName)
			assert.NoError(t, err)
			assert.Equal(t, modeTemplate.Length(), modeCustom.Length())
			assert.True(t, modeNormal.Equal(modeCustom))
		}
	}
}

func TestFindModeTemplatesByPattern(t *testing.T) {
	mts := make(ModeTemplatesStore)

	for i := 1; i < 18; i++ {
		customModeTemplate := make(ModeTemplate, i+1)
		customModeName := ModeName(fmt.Sprintf("%s %d", "custom abstract mode template", i))
		for j := 0; j <= i; j++ {
			customModeTemplate[j] = HalfTones(j)
		}
		mts.AddTemplate(customModeName, customModeTemplate)
	}

	for customModeName, modeTemplate := range mts {
		testCases := generateAbstractPatterns(modeTemplate[1:])
		for _, testCase := range testCases {
			result := mts.FindModeTemplatesByPattern(testCase)
			var exist bool
			for modeName := range result {
				if modeName == customModeName {
					exist = true
				}
			}
			assert.True(t, exist, "customModeName: %s, testCase: %+v", customModeName, testCase)
		}
	}
}

func TestFindModeTemplatesByPatternCases(t *testing.T) {
	mts := InitModeTemplatesStore()
	modeTemplates := mts.FindModeTemplatesByPattern(TemplateNaturalMinor())
	if _, ok := modeTemplates[ModeNameAeolian]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", ModeNameAeolian)
	}
	if _, ok := modeTemplates[ModeNameNaturalMinor]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", ModeNameNaturalMinor)
	}

	modeTemplates = mts.FindModeTemplatesByPattern(ModeTemplate{1, 2, 2, 1, 2, 2})
	if _, ok := modeTemplates[ModeNameAeolian]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", ModeNameAeolian)
	}
	if _, ok := modeTemplates[ModeNameNaturalMinor]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", ModeNameNaturalMinor)
	}

	modeTemplates = mts.FindModeTemplatesByPattern(ModeTemplate{2})
	if _, ok := modeTemplates[ModeNameAeolian]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", ModeNameAeolian)
	}
	if _, ok := modeTemplates[ModeNameNaturalMinor]; !ok {
		assert.Failf(t, "absent result", "expected mode name: %s", ModeNameNaturalMinor)
	}
}

func TestFindModeTemplatesByNotes(t *testing.T) {
	testCases := []struct {
		notes         []*Note
		expectedModes []ModeName
	}{
		{
			notes:         newNotesFromNoteNames(C, D, E, F, G, A, B),
			expectedModes: []ModeName{ModeNameAeolian, ModeNameIonian, ModeNamePhrygian, ModeNameLocrian, ModeNameDorian, ModeNameLydian, ModeNameMixoLydian, ModeNameNaturalMajor, ModeNameNaturalMinor},
		},
	}

	mts := InitModeTemplatesStore()
	for _, testCase := range testCases {
		var result ModeTemplatesWithPrime
		assert.NotPanics(t, func() { result = mts.FindModeTemplatesByNotes(testCase.notes) }, "testCase.notes: %+v", testCase.notes) //nolint:scopelint

		for _, modeName := range testCase.expectedModes {
			assert.True(t, result.Contains(modeName), "expected modename: %s", modeName)
		}

		assert.Equal(t, len(result), len(testCase.expectedModes))
	}
}
