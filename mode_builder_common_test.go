package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
