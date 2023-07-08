package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCommonMode(t *testing.T) {
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
		// Pentatonics

		// Main pentatonics
		constructTestCase(TemplatePentatonicMajor(), ModeNamePentatonicMajor, MustNewNoteWithoutOctave(C), []NoteName{C, D, E, G, A}),
		constructTestCase(TemplatePentatonicSustained(), ModeNamePentatonicSustained, MustNewNoteWithoutOctave(D), []NoteName{D, E, G, A, C}),
		constructTestCase(TemplatePentatonicBluesMinor(), ModeNamePentatonicBluesMinor, MustNewNoteWithoutOctave(E), []NoteName{E, G, A, C, D}),
		constructTestCase(TemplatePentatonicBluesMajor(), ModeNamePentatonicBluesMajor, MustNewNoteWithoutOctave(G), []NoteName{G, A, C, D, E}),
		constructTestCase(TemplatePentatonicMinor(), ModeNamePentatonicBluesMinor, MustNewNoteWithoutOctave(A), []NoteName{A, C, D, E, G}),

		// Japanese pentatonics
		constructTestCase(TemplatePentatonicHirajoshi(), ModeNamePentatonicHirajoshi, MustNewNoteWithoutOctave(C), []NoteName{C, D, EFLAT, G, AFLAT}),
		constructTestCase(TemplatePentatonicIwato(), ModeNamePentatonicIwato, MustNewNoteWithoutOctave(D), []NoteName{D, EFLAT, G, AFLAT, C}),
		constructTestCase(TemplatePentatonicHonKumoiShiouzhi(), ModeNamePentatonicHonKumoiShiouzhi, MustNewNoteWithoutOctave(EFLAT), []NoteName{EFLAT, G, AFLAT, C, D}),
		constructTestCase(TemplatePentatonicHonKumoiJoshi(), ModeNamePentatonicHonKumoiJoshi, MustNewNoteWithoutOctave(G), []NoteName{G, AFLAT, C, D, EFLAT}),
		constructTestCase(TemplatePentatonicLydianPentatonic(), ModeNamePentatonicLydianPentatonic, MustNewNoteWithoutOctave(AFLAT), []NoteName{AFLAT, C, D, EFLAT, G}),

		// Modes with other amount of degrees
		constructTestCase(ModeTemplate{12}, "custom mode with 1 degree", MustNewNoteWithoutOctave(C), []NoteName{C}),
		constructTestCase(ModeTemplate{6, 6}, "custom mode with 2 degrees", MustNewNoteWithoutOctave(C), []NoteName{C, FSHARP}),
		constructTestCase(ModeTemplate{6, 6}, "custom mode with 2 degrees", MustNewNoteWithoutOctave(BSHARP), []NoteName{BSHARP, FSHARP}),
		constructTestCase(ModeTemplate{6, 6}, "custom mode with 2 degrees", MustNewNoteWithoutOctave(CFLAT), []NoteName{CFLAT, F}),
		constructTestCase(ModeTemplate{4, 4, 4}, "custom mode with 3 degrees", MustNewNoteWithoutOctave(C), []NoteName{C, E, GSHARP}),
		constructTestCase(ModeTemplate{3, 3, 3, 3}, "custom mode with 4 degrees", MustNewNoteWithoutOctave(C), []NoteName{C, DSHARP, FSHARP, A}),
		constructTestCase(ModeTemplate{3, 3, 3, 3}, "custom mode with 5 degrees", MustNewNoteWithoutOctave(C), []NoteName{C, DSHARP, FSHARP, A}),
		constructTestCase(ModeTemplate{2, 2, 2, 2, 2, 2}, "custom mode with 6 degrees", MustNewNoteWithoutOctave(C), []NoteName{C, D, E, FSHARP, GSHARP, ASHARP}),
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

func Test_templateNoteCommon_getNext(t *testing.T) {
	tnc1 := &templateNoteCommon{}
	tnc2 := &templateNoteCommon{}
	tnc1.next = tnc2

	assert.Equal(t, tnc2, tnc1.getNext())
	assert.Nil(t, tnc2.getNext())
}

func Test_templateNotesCommon_getTemplateNote(t *testing.T) {
	templateNotesInstance := getTemplateNotesCommon()

	t.Run("getTemplateNotesCommon positive cases", func(t *testing.T) {
		testCases := GetFullChromaticScale()
		for i := range testCases {
			firstTemplateNote := templateNotesInstance.getTemplateNote(&testCases[i])
			assert.NotNil(t, firstTemplateNote, "expected common template note from note: %+v, actual: nil", &testCases[i])
			if firstTemplateNote.isAltered {
				assert.NotNil(t, firstTemplateNote.alteredNotes)
				var exist bool
				for _, note := range firstTemplateNote.alteredNotes {
					if note.realNote().IsEqualByName(&testCases[i]) {
						exist = true
					}
				}
				assert.True(t, exist)
			} else {
				assert.NotNil(t, firstTemplateNote.notAltered)
				assert.Equal(t, &testCases[i], firstTemplateNote.notAltered)
			}

			assert.NotNil(t, firstTemplateNote.next)
			assert.NotNil(t, firstTemplateNote.previous)
		}
	})

	t.Run("getTemplateNotesCommon negative cases", func(t *testing.T) {
		// not existent note
		assert.Nil(t, templateNotesInstance.getTemplateNote(&Note{name: NoteName("HELLO!")}))
		// impossible case
		templateNotesInstance.templateNoteCommon = nil
		assert.Nil(t, templateNotesInstance.getTemplateNote(&Note{name: C}))
	})
}
