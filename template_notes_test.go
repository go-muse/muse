package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTemplateNotesInstance(t *testing.T) {
	templateNotesInstance := getTemplateNotesInstance()

	t.Run("getTemplateNotesInstance positive cases", func(t *testing.T) {
		testCases := []struct {
			*Note
			isAltered bool
		}{
			{
				&Note{name: C},
				false,
			},
			{
				&Note{name: DFLAT},
				true,
			},
			{
				&Note{name: CSHARP},
				true,
			},
			{
				&Note{name: D},
				false,
			},
			{
				&Note{name: EFLAT},
				true,
			},
			{
				&Note{name: DSHARP},
				true,
			},
			{
				&Note{name: E},
				false,
			},
			{
				&Note{name: F},
				false,
			},
			{
				&Note{name: GFLAT},
				true,
			},
			{
				&Note{name: FSHARP},
				true,
			},
			{
				&Note{name: G},
				false,
			},
			{
				&Note{name: AFLAT},
				true,
			},
			{
				&Note{name: GSHARP},
				true,
			},
			{
				&Note{name: A},
				false,
			},
			{
				&Note{name: BFLAT},
				true,
			},
			{
				&Note{name: ASHARP},
				true,
			},
			{
				&Note{name: B},
				false,
			},
		}

		for _, testCase := range testCases {
			firstTemplateNote := templateNotesInstance.getTemplateNote(testCase.Note)
			assert.Equal(t, testCase.isAltered, firstTemplateNote.isAltered)
			assert.Nil(t, firstTemplateNote.resultingNote)
			assert.NotNil(t, firstTemplateNote.next)
			assert.NotNil(t, firstTemplateNote.previous)

			if testCase.isAltered {
				assert.Nil(t, firstTemplateNote.notAltered)
				assert.NotNil(t, firstTemplateNote.alteredNotes)
			} else {
				assert.NotNil(t, firstTemplateNote.notAltered)
				assert.Nil(t, firstTemplateNote.alteredNotes)
			}
		}
	})

	t.Run("getTemplateNotesInstance negative cases", func(t *testing.T) {
		// not existent note
		assert.Nil(t, templateNotesInstance.getTemplateNote(&Note{name: NoteName("HELLO!")}))
		// impossible case
		templateNotesInstance.templateNote = nil
		assert.Nil(t, templateNotesInstance.getTemplateNote(&Note{name: C}))
	})
}
