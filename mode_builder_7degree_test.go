package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild7DegreeMode(t *testing.T) {
	type testCase struct {
		modeTemplate ModeTemplate
		resultMode   *Mode
		controlMode  *Mode
	}

	constructTestCase := func(modeTemplate ModeTemplate, modeName ModeName, firstNote *Note, expectedNotes []NoteName) testCase {
		return testCase{
			modeTemplate: modeTemplate,
			resultMode:   newModeBuilder(modeTemplate).build(modeName, modeTemplate, firstNote),
			controlMode:  generateModeWithNotes(modeTemplate, expectedNotes),
		}
	}

	testCases := []testCase{
		// Tonal modes
		constructTestCase(TemplateNaturalMinor(), ModeNameNaturalMinor, MustNewNote(C), []NoteName{C, D, EFLAT, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateMelodicMinor(), ModeNameMelodicMinor, MustNewNote(C), []NoteName{C, D, EFLAT, F, G, A, B}),
		constructTestCase(TemplateHarmonicMinor(), ModeNameHarmonicMinor, MustNewNote(C), []NoteName{C, D, EFLAT, F, G, AFLAT, B}),
		constructTestCase(TemplateNaturalMajor(), ModeNameNaturalMajor, MustNewNote(C), []NoteName{C, D, E, F, G, A, B}),
		constructTestCase(TemplateMelodicMajor(), ModeNameMelodicMajor, MustNewNote(C), []NoteName{C, D, E, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateHarmonicMajor(), ModeNameHarmonicMajor, MustNewNote(C), []NoteName{C, D, E, F, G, AFLAT, B}),

		// Modes of the Major scale
		constructTestCase(TemplateIonian(), ModeNameIonian, MustNewNote(C), []NoteName{C, D, E, F, G, A, B}),
		constructTestCase(TemplateDorian(), ModeNameDorian, MustNewNote(C), []NoteName{C, D, EFLAT, F, G, A, BFLAT}),
		constructTestCase(TemplateAeolian(), ModeNameAeolian, MustNewNote(C), []NoteName{C, D, EFLAT, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLydian(), ModeNameLydian, MustNewNote(C), []NoteName{C, D, E, FSHARP, G, A, B}),
		constructTestCase(TemplateMixoLydian(), ModeNameMixoLydian, MustNewNote(C), []NoteName{C, D, E, F, G, A, BFLAT}),
		constructTestCase(TemplatePhrygian(), ModeNamePhrygian, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLocrian(), ModeNameLocrian, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, F, GFLAT, AFLAT, BFLAT}),

		// Modes Of The Melodic Minor scale
		constructTestCase(TemplateIonianFlat3(), ModeNameIonianFlat3, MustNewNote(C), []NoteName{C, D, EFLAT, F, G, A, B}),
		constructTestCase(TemplatePhrygoDorian(), ModeNamePhrygoDorian, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, F, G, A, BFLAT}),
		constructTestCase(TemplateLydianAugmented(), ModeNameLydianAugmented, MustNewNote(C), []NoteName{C, D, E, FSHARP, GSHARP, A, B}),
		constructTestCase(TemplateLydianDominant(), ModeNameLydianDominant, MustNewNote(C), []NoteName{C, D, E, FSHARP, G, A, BFLAT}),
		constructTestCase(TemplateIonianAeolian(), ModeNameIonianAeolian, MustNewNote(C), []NoteName{C, D, E, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateAeolianLydian(), ModeNameAeolianLydian, MustNewNote(C), []NoteName{C, D, EFLAT, F, GFLAT, AFLAT, BFLAT}),
		constructTestCase(TemplateSuperLocrian(), ModeNameSuperLocrian, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT}),

		// Modes of the Harmonic Minor scale
		constructTestCase(TemplateAeolianRais7(), ModeNameAeolianRais7, MustNewNote(C), []NoteName{C, D, EFLAT, F, G, AFLAT, B}),
		constructTestCase(TemplateLocrianRais6(), ModeNameLocrianRais6, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, F, GFLAT, A, BFLAT}),
		constructTestCase(TemplateIonianRais5(), ModeNameIonianRais5, MustNewNote(C), []NoteName{C, D, E, F, GSHARP, A, B}),
		constructTestCase(TemplateUkrainianDorian(), ModeNameUkrainianDorian, MustNewNote(C), []NoteName{C, D, EFLAT, FSHARP, G, A, BFLAT}),
		constructTestCase(TemplatePhrygianDominant(), ModeNamePhrygianDominant, MustNewNote(C), []NoteName{C, DFLAT, E, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLydianRais9(), ModeNameLydianRais9, MustNewNote(C), []NoteName{C, DSHARP, E, FSHARP, G, A, B}),
		constructTestCase(TemplateUltraLocrian(), ModeNameUltraLocrian, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT2}),

		// Modes Of The Harmonic Major scale
		constructTestCase(TemplateIonianFlat6(), ModeNameIonianFlat6, MustNewNote(C), []NoteName{C, D, E, F, G, AFLAT, B}),
		constructTestCase(TemplateDorianDiminished(), ModeNameDorianDiminished, MustNewNote(C), []NoteName{C, D, EFLAT, F, GFLAT, A, BFLAT}),
		constructTestCase(TemplatePhrygianDiminished(), ModeNamePhrygianDiminished, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, FFLAT, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLydianDiminished(), ModeNameLydianDiminished, MustNewNote(C), []NoteName{C, D, EFLAT, FSHARP, G, A, B}),
		constructTestCase(TemplateMixolydianFlat2(), ModeNameMixolydianFlat2, MustNewNote(C), []NoteName{C, DFLAT, E, F, G, A, BFLAT}),
		constructTestCase(TemplateLydianAugmented2(), ModeNameLydianAugmented2, MustNewNote(C), []NoteName{C, DSHARP, E, FSHARP, GSHARP, A, B}),
		constructTestCase(TemplateLocrianDoubleFlat7(), ModeNameLocrianDoubleFlat7, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, F, GFLAT, AFLAT, BFLAT2}),

		// Double Harmonic Major Modes
		constructTestCase(TemplateHungarianMajor(), ModeNameHungarianMajor, MustNewNote(C), []NoteName{C, DFLAT, E, F, G, AFLAT, B}),
		constructTestCase(TemplateLydianRais2Rais6(), ModeNameLydianRais2Rais6, MustNewNote(C), []NoteName{C, DSHARP, E, FSHARP, G, ASHARP, B}),
		constructTestCase(TemplateUltraphrygian(), ModeNameUltraphrygian, MustNewNote(C), []NoteName{C, DFLAT, EFLAT, FFLAT, G, AFLAT, BFLAT2}),
		constructTestCase(TemplateHungarianMinor(), ModeNameHungarianMinor, MustNewNote(C), []NoteName{C, D, EFLAT, FSHARP, G, AFLAT, B}),
		constructTestCase(TemplateOriental(), ModeNameOriental, MustNewNote(C), []NoteName{C, DFLAT, E, F, GFLAT, A, BFLAT}),
		constructTestCase(TemplateIonianAugmented2(), ModeNameIonianAugmented2, MustNewNote(C), []NoteName{C, DSHARP, E, F, GSHARP, A, B}),
		constructTestCase(TemplateLocrianDoubleFlat3DoubleFlat7(), ModeNameLocrianDoubleFlat3DoubleFlat7, MustNewNote(C), []NoteName{C, DFLAT, EFLAT2, F, GFLAT, AFLAT, BFLAT2}),
	}

	testingFunc := func(modeTemplate ModeTemplate, resultMode, controlMode *Mode) {
		d1chan := controlMode.IterateOneRound(false)
		for d2 := range resultMode.IterateOneRound(false) {
			d1 := <-d1chan
			assert.True(t, d1.Equal(d2),
				"%+v, mode name: %s, expected note: %s, actual: %s,\nexpected scale: %+v,\nactual scale: %+v\n",
				modeTemplate,
				resultMode.Name(), d1.Note().Name(), d2.Note().Name(),
				controlMode.IterateOneRound(false).GetAllNotes(),
				resultMode.IterateOneRound(false).GetAllNotes())
		}
	}

	for _, testCase := range testCases {
		testingFunc(testCase.modeTemplate, testCase.resultMode, testCase.controlMode)
	}
}
