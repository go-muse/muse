package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

func TestBuildCommonMode(t *testing.T) {
	type testCase struct {
		modeTemplate  halftone.Template
		builder       Builder
		expectedNotes note.Names
		name          string
	}

	constructTestCase := func(modeTemplate halftone.Template, modeName string, firstNote *note.Note, expectedNotes note.Names) testCase {
		return testCase{
			modeTemplate:  modeTemplate,
			builder:       NewBuilderCommon(modeTemplate, firstNote),
			expectedNotes: expectedNotes,
			name:          modeName,
		}
	}

	testCases := []testCase{
		// Pentatonics

		// Main pentatonics
		constructTestCase(halftone.Template{2, 2, 3, 2, 3}, "PentatonicMajor", note.MustNewNote(note.C), []note.Name{note.C, note.D, note.E, note.G, note.A}),
		constructTestCase(halftone.Template{2, 3, 2, 3, 2}, "PentatonicSustained", note.MustNewNote(note.D), []note.Name{note.D, note.E, note.G, note.A, note.C}),
		constructTestCase(halftone.Template{3, 2, 3, 2, 2}, "PentatonicBluesMinor", note.MustNewNote(note.E), []note.Name{note.E, note.G, note.A, note.C, note.D}),
		constructTestCase(halftone.Template{2, 3, 2, 2, 3}, "PentatonicBluesMajor", note.MustNewNote(note.G), []note.Name{note.G, note.A, note.C, note.D, note.E}),
		constructTestCase(halftone.Template{3, 2, 2, 3, 2}, "PentatonicBluesMinor", note.MustNewNote(note.A), []note.Name{note.A, note.C, note.D, note.E, note.G}),

		// Japanese pentatonics
		constructTestCase(halftone.Template{2, 1, 4, 1, 4}, "PentatonicHirajoshi", note.MustNewNote(note.C), []note.Name{note.C, note.D, note.EFLAT, note.G, note.AFLAT}),
		constructTestCase(halftone.Template{1, 4, 1, 4, 2}, "PentatonicIwato", note.MustNewNote(note.D), []note.Name{note.D, note.EFLAT, note.G, note.AFLAT, note.C}),
		constructTestCase(halftone.Template{4, 1, 4, 2, 1}, "PentatonicHonKumoiShiouzhi", note.MustNewNote(note.EFLAT), []note.Name{note.EFLAT, note.G, note.AFLAT, note.C, note.D}),
		constructTestCase(halftone.Template{1, 4, 2, 1, 4}, "PentatonicHonKumoiJoshi", note.MustNewNote(note.G), []note.Name{note.G, note.AFLAT, note.C, note.D, note.EFLAT}),
		constructTestCase(halftone.Template{4, 2, 1, 4, 1}, "PentatonicLydianPentatonic", note.MustNewNote(note.AFLAT), []note.Name{note.AFLAT, note.C, note.D, note.EFLAT, note.G}),

		// Modes with other amount of degrees
		constructTestCase(halftone.Template{12}, "custom mode with 1 degree", note.MustNewNote(note.C), []note.Name{note.C}),
		constructTestCase(halftone.Template{6, 6}, "custom mode with 2 degrees", note.MustNewNote(note.C), []note.Name{note.C, note.FSHARP}),
		constructTestCase(halftone.Template{6, 6}, "custom mode with 2 degrees", note.MustNewNote(note.BSHARP), []note.Name{note.BSHARP, note.FSHARP}),
		constructTestCase(halftone.Template{6, 6}, "custom mode with 2 degrees", note.MustNewNote(note.CFLAT), []note.Name{note.CFLAT, note.F}),
		constructTestCase(halftone.Template{4, 4, 4}, "custom mode with 3 degrees", note.MustNewNote(note.C), []note.Name{note.C, note.E, note.GSHARP}),
		constructTestCase(halftone.Template{3, 3, 3, 3}, "custom mode with 4 degrees", note.MustNewNote(note.C), []note.Name{note.C, note.DSHARP, note.FSHARP, note.A}),
		constructTestCase(halftone.Template{3, 3, 3, 3}, "custom mode with 5 degrees", note.MustNewNote(note.C), []note.Name{note.C, note.DSHARP, note.FSHARP, note.A}),
		constructTestCase(halftone.Template{2, 2, 2, 2, 2, 2}, "custom mode with 6 degrees", note.MustNewNote(note.C), []note.Name{note.C, note.D, note.E, note.FSHARP, note.GSHARP, note.ASHARP}),
		constructTestCase(halftone.Template{2, 2, 2, 2, 2, 1, 1}, "custom mode with 7 degrees", note.MustNewNote(note.C), []note.Name{note.C, note.D, note.E, note.FSHARP, note.GSHARP, note.ASHARP, note.B}),
	}

	testingFunc := func(modeTemplate halftone.Template, builder Builder, expectedNotes note.Names, testCaseName string) {
		// length of halftone template should be equal to length of expected notes
		require.Equal(t, modeTemplate.Length(), expectedNotes.Length())

		// builder returns results from 2nd note
		expectedNotes = expectedNotes[1:]

		// halftonesFromPrimeCheck is the way to check halftones "from prime" from builder and from mode's template
		var halftonesFromPrimeCheck halftone.HalfTones

		// index of expected note
		var i int

		// iterating through the given template (expected values)
		for res := range modeTemplate.IterateOneRound(false) {
			halftonesFromMode, halftonesFromPrimeExpected := res()
			halftonesFromPrimeCheck += halftonesFromMode

			// self-check for equality of halftones from prime from builder and from mode's template
			assert.Equal(t, halftonesFromPrimeExpected, halftonesFromPrimeCheck)

			// If buildResult is available, it means there are more notes to build
			// it can be unavailable if mode contains just one note, and we can't build next note
			buildResult, ok := <-builder
			if ok {
				n, halftonesFromPrime := buildResult()
				assert.Equal(t, halftonesFromPrimeCheck, halftonesFromPrime)
				assert.Equal(t, halftonesFromPrimeExpected, halftonesFromPrime)
				assert.Equal(t, n.Name(), expectedNotes[i], "test case: '%s' unexpected note name, expected: '%s', actual: '%s'", testCaseName, expectedNotes[i], n.Name())
			}

			// increment index of expected note
			i++
		}
	}

	for _, testCase := range testCases {
		testingFunc(testCase.modeTemplate, testCase.builder, testCase.expectedNotes, testCase.name)
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
		testCases := note.GetSetFullChromatic()
		for i := range testCases {
			firstTemplateNote := templateNotesInstance.getTemplateNote(testCases[i])
			assert.NotNil(t, firstTemplateNote, "expected common template note from note: %+v, actual: nil", &testCases[i])
			if firstTemplateNote.isAltered {
				assert.NotNil(t, firstTemplateNote.alteredNotes)
				var exist bool
				for _, note := range firstTemplateNote.alteredNotes {
					if note.realNote().IsEqualByName(testCases[i]) {
						exist = true
					}
				}
				assert.True(t, exist)
			} else {
				assert.NotNil(t, firstTemplateNote.notAltered)
				assert.Equal(t, testCases[i], firstTemplateNote.notAltered)
			}

			assert.NotNil(t, firstTemplateNote.next)
			assert.NotNil(t, firstTemplateNote.previous)
		}
	})

	t.Run("getTemplateNotesCommon negative cases", func(t *testing.T) {
		// impossible case
		templateNotesInstance.templateNoteCommon = nil
		assert.Nil(t, templateNotesInstance.getTemplateNote(note.C.MustMakeNote()))
	})
}
