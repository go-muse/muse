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

	constructTestCase := func(modeTemplate halftone.Template, modeName string, firstNote note.Note, expectedNotes note.Names) testCase {
		return testCase{
			modeTemplate:  modeTemplate,
			builder:       NewBuilderHeptatonic(modeTemplate, firstNote),
			expectedNotes: expectedNotes,
			name:          modeName,
		}
	}

	testCases := []testCase{
		// Heptatonics
		constructTestCase(halftone.Template{2, 2, 2, 2, 2, 1, 1}, "custom mode with 7 degrees", note.C.NewNote(), []note.Name{note.C, note.D, note.E, note.FSHARP, note.GSHARP, note.ASHARP, note.B}),
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
		assert.Nil(t, templateNotesInstance.getTemplateNote(note.C.NewNote()))
	})
}

func Test_NextBaseNote(t *testing.T) {
	tni := getTemplateNotesHeptatonic()

	testCases := []struct {
		tns  note.Notes
		want note.Note
	}{
		{
			tns:  note.Notes{note.C.NewNote(), note.CSHARP.NewNote(), note.CFLAT.NewNote(), note.CSHARP2.NewNote(), note.CFLAT2.NewNote()},
			want: note.D.NewNote(),
		},
		{
			tns:  note.Notes{note.D.NewNote(), note.DSHARP.NewNote(), note.DFLAT.NewNote(), note.DSHARP2.NewNote(), note.DFLAT2.NewNote()},
			want: note.E.NewNote(),
		},
		{
			tns:  note.Notes{note.E.NewNote(), note.ESHARP.NewNote(), note.EFLAT.NewNote(), note.ESHARP2.NewNote(), note.EFLAT2.NewNote()},
			want: note.F.NewNote(),
		},
		{
			tns:  note.Notes{note.F.NewNote(), note.FSHARP.NewNote(), note.FFLAT.NewNote(), note.FSHARP2.NewNote(), note.FFLAT2.NewNote()},
			want: note.G.NewNote(),
		},
		{
			tns:  note.Notes{note.G.NewNote(), note.GSHARP.NewNote(), note.GFLAT.NewNote(), note.GSHARP2.NewNote(), note.GFLAT2.NewNote()},
			want: note.A.NewNote(),
		},
		{
			tns:  note.Notes{note.A.NewNote(), note.ASHARP.NewNote(), note.AFLAT.NewNote(), note.ASHARP2.NewNote(), note.AFLAT2.NewNote()},
			want: note.B.NewNote(),
		},
		{
			tns:  note.Notes{note.B.NewNote(), note.BSHARP.NewNote(), note.BFLAT.NewNote(), note.BSHARP2.NewNote(), note.BFLAT2.NewNote()},
			want: note.C.NewNote(),
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
