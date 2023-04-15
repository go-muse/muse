package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild7DegreeMode(t *testing.T) {
	mb := newModeBuilder()

	testCases := []struct {
		resultMode  *Mode
		controlMode *Mode
	}{
		{
			resultMode:  mb.build7DegreeMode(ModeNameIonian, TemplateIonian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, A, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameIonian, TemplateAeolian(), MustNewNote(A)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{A, B, C, D, E, F, G}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameIonian, TemplateSuperLocrian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameIonian, TemplateUltraLocrian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT2}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameIonian, TemplatePhrygian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, F, G, AFLAT, BFLAT}),
		},
	}
	testingFunc := func(resultMode, controlMode *Mode) {
		d1chan := controlMode.IterateOneRound(false)
		for d2 := range resultMode.IterateOneRound(false) {
			assert.True(t, (<-d1chan).Equal(d2))
		}
	}

	for _, testCase := range testCases {
		testingFunc(testCase.resultMode, testCase.controlMode)
	}
}
