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
		// Tonal modes

		{
			resultMode:  mb.build7DegreeMode(ModeNameNaturalMinor, TemplateNaturalMinor(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, G, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameMelodicMinor, TemplateMelodicMinor(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, G, A, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameHarmonicMinor, TemplateHarmonicMinor(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, G, AFLAT, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameNaturalMajor, TemplateNaturalMajor(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, A, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameMelodicMajor, TemplateMelodicMajor(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameHarmonicMajor, TemplateHarmonicMajor(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, AFLAT, B}),
		},

		// Modes of the Major scale

		{
			resultMode:  mb.build7DegreeMode(ModeNameIonian, TemplateIonian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, A, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameDorian, TemplateDorian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, G, A, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameDorian, TemplateDorian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, G, A, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameAeolian, TemplateAeolian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, G, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameAeolian, TemplateAeolian(), MustNewNote(A)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{A, B, C, D, E, F, G}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameLydian, TemplateLydian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, FSHARP, G, A, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameMixoLydian, TemplateMixoLydian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, A, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameMixoLydian, TemplateMixoLydian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, A, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNamePhrygian, TemplatePhrygian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, F, G, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameLocrian, TemplateLocrian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, F, GFLAT, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameUltraLocrian, TemplateUltraLocrian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT2}),
		},

		// Modes Of The Melodic Minor scale

		{
			resultMode:  mb.build7DegreeMode(ModeNameIonianFlat3, TemplateIonianFlat3(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, G, A, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNamePhrygoDorian, TemplatePhrygoDorian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, F, G, A, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameLydianAugmented, TemplateLydianAugmented(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, FSHARP, GSHARP, A, B}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameLydianDominant, TemplateLydianDominant(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, FSHARP, G, A, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameIonianAeolian, TemplateIonianAeolian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, E, F, G, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameAeolianLydian, TemplateAeolianLydian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, D, EFLAT, F, GFLAT, AFLAT, BFLAT}),
		},
		{
			resultMode:  mb.build7DegreeMode(ModeNameSuperLocrian, TemplateSuperLocrian(), MustNewNote(C)),
			controlMode: generateModeWithNotes(TemplateIonian(), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT}),
		},
	}

	testingFunc := func(resultMode, controlMode *Mode) {
		d1chan := controlMode.IterateOneRound(false)
		for d2 := range resultMode.IterateOneRound(false) {
			d1 := <-d1chan
			assert.True(t, d1.Equal(d2), "mode name: %s, expected note: %s, actual: %s", resultMode.Name(), d1.Note().Name(), d2.Note().Name())
		}
	}

	for _, testCase := range testCases {
		testingFunc(testCase.resultMode, testCase.controlMode)
	}
}
