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
		constructTestCase(TemplateNaturalMinor(), ModeNameNaturalMinor, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateMelodicMinor(), ModeNameMelodicMinor, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, G, A, B}),
		constructTestCase(TemplateHarmonicMinor(), ModeNameHarmonicMinor, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, G, AFLAT, B}),
		constructTestCase(TemplateNaturalMajor(), ModeNameNaturalMajor, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, G, A, B}),
		constructTestCase(TemplateMelodicMajor(), ModeNameMelodicMajor, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateHarmonicMajor(), ModeNameHarmonicMajor, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, G, AFLAT, B}),

		// Modes of the Major scale
		constructTestCase(TemplateIonian(), ModeNameIonian, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, G, A, B}),
		constructTestCase(TemplateDorian(), ModeNameDorian, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, G, A, BFLAT}),
		constructTestCase(TemplateAeolian(), ModeNameAeolian, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLydian(), ModeNameLydian, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, FSHARP, G, A, B}),
		constructTestCase(TemplateMixoLydian(), ModeNameMixoLydian, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, G, A, BFLAT}),
		constructTestCase(TemplatePhrygian(), ModeNamePhrygian, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLocrian(), ModeNameLocrian, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, F, GFLAT, AFLAT, BFLAT}),

		// Modes Of The Melodic Minor scale
		constructTestCase(TemplateIonianFlat3(), ModeNameIonianFlat3, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, G, A, B}),
		constructTestCase(TemplatePhrygoDorian(), ModeNamePhrygoDorian, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, F, G, A, BFLAT}),
		constructTestCase(TemplateLydianAugmented(), ModeNameLydianAugmented, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, FSHARP, GSHARP, A, B}),
		constructTestCase(TemplateLydianDominant(), ModeNameLydianDominant, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, FSHARP, G, A, BFLAT}),
		constructTestCase(TemplateIonianAeolian(), ModeNameIonianAeolian, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateAeolianLydian(), ModeNameAeolianLydian, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, GFLAT, AFLAT, BFLAT}),
		constructTestCase(TemplateSuperLocrian(), ModeNameSuperLocrian, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT}),

		// Modes of the Harmonic Minor scale
		constructTestCase(TemplateAeolianRais7(), ModeNameAeolianRais7, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, G, AFLAT, B}),
		constructTestCase(TemplateLocrianRais6(), ModeNameLocrianRais6, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, F, GFLAT, A, BFLAT}),
		constructTestCase(TemplateIonianRais5(), ModeNameIonianRais5, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, GSHARP, A, B}),
		constructTestCase(TemplateUkrainianDorian(), ModeNameUkrainianDorian, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, FSHARP, G, A, BFLAT}),
		constructTestCase(TemplatePhrygianDominant(), ModeNamePhrygianDominant, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, E, F, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLydianRais9(), ModeNameLydianRais9, MustNewNoteWithoutOctave(C), []NoteName{C, DSHARP, E, FSHARP, G, A, B}),
		constructTestCase(TemplateUltraLocrian(), ModeNameUltraLocrian, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, FFLAT, GFLAT, AFLAT, BFLAT2}),

		// Modes Of The Harmonic Major scale
		constructTestCase(TemplateIonianFlat6(), ModeNameIonianFlat6, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, F, G, AFLAT, B}),
		constructTestCase(TemplateDorianDiminished(), ModeNameDorianDiminished, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, F, GFLAT, A, BFLAT}),
		constructTestCase(TemplatePhrygianDiminished(), ModeNamePhrygianDiminished, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, FFLAT, G, AFLAT, BFLAT}),
		constructTestCase(TemplateLydianDiminished(), ModeNameLydianDiminished, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, FSHARP, G, A, B}),
		constructTestCase(TemplateMixolydianFlat2(), ModeNameMixolydianFlat2, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, E, F, G, A, BFLAT}),
		constructTestCase(TemplateLydianAugmented2(), ModeNameLydianAugmented2, MustNewNoteWithoutOctave(C), []NoteName{C, DSHARP, E, FSHARP, GSHARP, A, B}),
		constructTestCase(TemplateLocrianDoubleFlat7(), ModeNameLocrianDoubleFlat7, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, F, GFLAT, AFLAT, BFLAT2}),

		// Double Harmonic Major Modes
		constructTestCase(TemplateHungarianMajor(), ModeNameHungarianMajor, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, E, F, G, AFLAT, B}),
		constructTestCase(TemplateLydianRais2Rais6(), ModeNameLydianRais2Rais6, MustNewNoteWithoutOctave(C), []NoteName{C, DSHARP, E, FSHARP, G, ASHARP, B}),
		constructTestCase(TemplateUltraphrygian(), ModeNameUltraphrygian, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT, FFLAT, G, AFLAT, BFLAT2}),
		constructTestCase(TemplateHungarianMinor(), ModeNameHungarianMinor, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, FSHARP, G, AFLAT, B}),
		constructTestCase(TemplateOriental(), ModeNameOriental, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, E, F, GFLAT, A, BFLAT}),
		constructTestCase(TemplateIonianAugmented2(), ModeNameIonianAugmented2, MustNewNoteWithoutOctave(C), []NoteName{C, DSHARP, E, F, GSHARP, A, B}),
		constructTestCase(TemplateLocrianDoubleFlat3DoubleFlat7(), ModeNameLocrianDoubleFlat3DoubleFlat7, MustNewNoteWithoutOctave(C), []NoteName{C, DFLAT, EFLAT2, F, GFLAT, AFLAT, BFLAT2}),
	}

	testingFunc := func(modeTemplate ModeTemplate, resultMode, controlMode *Mode) {
		d1chan := controlMode.IterateOneRound(false)
		for d2 := range resultMode.IterateOneRound(false) {
			d1 := <-d1chan
			assert.True(t, d1.IsEqual(d2),
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

func Test_templateNotes7degree_getTemplateNote(t *testing.T) {
	templateNotesInstance := getTemplateNotes7degree()

	t.Run("getTemplateNotes7degree positive cases", func(t *testing.T) {
		testCases := GetAllPossibleNotes(2)
		for i := range testCases {
			firstTemplateNote := templateNotesInstance.getTemplateNote(&testCases[i])
			assert.NotNil(t, firstTemplateNote.allNotes)
			assert.NotNil(t, firstTemplateNote.next)
		}
	})

	t.Run("getTemplateNotes7degree negative cases", func(t *testing.T) {
		// not existent note
		assert.Nil(t, templateNotesInstance.getTemplateNote(&Note{name: NoteName("HELLO!")}))
		// impossible case
		templateNotesInstance.templateNote7degree = nil
		assert.Nil(t, templateNotesInstance.getTemplateNote(&Note{name: C}))
	})
}

func Test_NextBaseNote(t *testing.T) {
	tni := getTemplateNotes7degree()

	testCases := []struct {
		tns  []*Note
		want *Note
	}{
		{
			tns:  []*Note{newNote(C), newNote(CSHARP), newNote(CFLAT), newNote(CSHARP2), newNote(CFLAT2)},
			want: newNote(D),
		},
		{
			tns:  []*Note{newNote(D), newNote(DSHARP), newNote(DFLAT), newNote(DSHARP2), newNote(DFLAT2)},
			want: newNote(E),
		},
		{
			tns:  []*Note{newNote(E), newNote(ESHARP), newNote(EFLAT), newNote(ESHARP2), newNote(EFLAT2)},
			want: newNote(F),
		},
		{
			tns:  []*Note{newNote(F), newNote(FSHARP), newNote(FFLAT), newNote(FSHARP2), newNote(FFLAT2)},
			want: newNote(G),
		},
		{
			tns:  []*Note{newNote(G), newNote(GSHARP), newNote(GFLAT), newNote(GSHARP2), newNote(GFLAT2)},
			want: newNote(A),
		},
		{
			tns:  []*Note{newNote(A), newNote(ASHARP), newNote(AFLAT), newNote(ASHARP2), newNote(AFLAT2)},
			want: newNote(B),
		},
		{
			tns:  []*Note{newNote(B), newNote(BSHARP), newNote(BFLAT), newNote(BSHARP2), newNote(BFLAT2)},
			want: newNote(C),
		},
	}

	for _, testCase := range testCases {
		for _, tn := range testCase.tns {
			tni.setLastUsedBaseNote(tn)
			nextBase := tni.nextBaseNote()
			assert.Equal(t, testCase.want, nextBase, "expected: %+v, actual: %+v", testCase.want, nextBase)
		}
	}
}
