package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

func TestBuildHeptatonicMode(t *testing.T) {
	type testCase struct {
		modeTemplate  halftone.Template
		builder       Builder
		expectedNotes note.Names
		name          string
	}

	constructTestCase := func(modeTemplate halftone.Template, modeName string, firstNote *note.Note, expectedNotes note.Names) testCase {
		return testCase{
			modeTemplate:  modeTemplate,
			builder:       NewBuilderHeptatonic(modeTemplate, firstNote),
			expectedNotes: expectedNotes,
			name:          modeName,
		}
	}

	testCases := []testCase{
		// Heptatonics
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
				assert.Equal(t, expectedNotes[i], n.Name(), "test case: '%s' unexpected note name, expected: '%s', actual: '%s'", testCaseName, expectedNotes[i], n.Name())
			}

			// increment index of expected note
			i++
		}
	}

	for _, testCase := range testCases {
		testingFunc(testCase.modeTemplate, testCase.builder, testCase.expectedNotes, testCase.name)
	}
}

func Test_templateNotes7degree_getTemplateNote(t *testing.T) {
	templateNotesInstance := getTemplateNotesHeptatonic()

	t.Run("getTemplateNotesHeptatonic positive cases", func(t *testing.T) {
		testCases := note.GetNotesWithAlterations(
			note.Notes{note.C.MustMakeNote(), note.D.MustMakeNote(), note.E.MustMakeNote(), note.F.MustMakeNote(), note.G.MustMakeNote(), note.A.MustMakeNote(), note.B.MustMakeNote()},
			2,
		)
		for i := range testCases {
			firstTemplateNote := templateNotesInstance.getTemplateNote(testCases[i])
			assert.NotNil(t, firstTemplateNote.allNotes)
			assert.NotNil(t, firstTemplateNote.next)
		}
	})

	t.Run("getTemplateNotesHeptatonic negative cases", func(t *testing.T) {
		// impossible case
		templateNotesInstance.templateNoteHeptatonic = nil
		assert.Nil(t, templateNotesInstance.getTemplateNote(note.C.MustNewNote()))
	})
}

func Test_NextBaseNote(t *testing.T) {
	tni := getTemplateNotesHeptatonic()

	testCases := []struct {
		tns  note.Notes
		want *note.Note
	}{
		{
			tns:  note.Notes{note.MustNewNote(note.C), note.MustNewNote(note.CSHARP), note.MustNewNote(note.CFLAT), note.MustNewNote(note.CSHARP2), note.MustNewNote(note.CFLAT2)},
			want: note.MustNewNote(note.D),
		},
		{
			tns:  note.Notes{note.MustNewNote(note.D), note.MustNewNote(note.DSHARP), note.MustNewNote(note.DFLAT), note.MustNewNote(note.DSHARP2), note.MustNewNote(note.DFLAT2)},
			want: note.MustNewNote(note.E),
		},
		{
			tns:  note.Notes{note.MustNewNote(note.E), note.MustNewNote(note.ESHARP), note.MustNewNote(note.EFLAT), note.MustNewNote(note.ESHARP2), note.MustNewNote(note.EFLAT2)},
			want: note.MustNewNote(note.F),
		},
		{
			tns:  note.Notes{note.MustNewNote(note.F), note.MustNewNote(note.FSHARP), note.MustNewNote(note.FFLAT), note.MustNewNote(note.FSHARP2), note.MustNewNote(note.FFLAT2)},
			want: note.MustNewNote(note.G),
		},
		{
			tns:  note.Notes{note.MustNewNote(note.G), note.MustNewNote(note.GSHARP), note.MustNewNote(note.GFLAT), note.MustNewNote(note.GSHARP2), note.MustNewNote(note.GFLAT2)},
			want: note.MustNewNote(note.A),
		},
		{
			tns:  note.Notes{note.MustNewNote(note.A), note.MustNewNote(note.ASHARP), note.MustNewNote(note.AFLAT), note.MustNewNote(note.ASHARP2), note.MustNewNote(note.AFLAT2)},
			want: note.MustNewNote(note.B),
		},
		{
			tns:  note.Notes{note.MustNewNote(note.B), note.MustNewNote(note.BSHARP), note.MustNewNote(note.BFLAT), note.MustNewNote(note.BSHARP2), note.MustNewNote(note.BFLAT2)},
			want: note.MustNewNote(note.C),
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
