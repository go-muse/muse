package muse

import (
	"testing"
	"unsafe"

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

func Test_saveTonic(t *testing.T) {
	tni := getTemplateNotesInstance()
	note := MustNewNote(BFLAT)
	tn := tni.getTemplateNote(note)
	tn.saveTonic(note)
	assert.Equal(t, HalfTones(0), tn.halfTonesFromPrime, "expected: 0, actual: %d", tn.halfTonesFromPrime)
	assert.True(t, tn.isTonic)

	currentNote := tn
	var halfTonesFromPrime HalfTones
	for unsafe.Pointer(currentNote.next) != unsafe.Pointer(tn) {
		halfTonesFromPrime++
		currentNote = currentNote.getNext()
		assert.Equal(t, halfTonesFromPrime, currentNote.halfTonesFromPrime, "expected: %d, actual: %d", halfTonesFromPrime, currentNote.halfTonesFromPrime)
		assert.False(t, currentNote.isTonic)
	}
}

func Test_getTonic(t *testing.T) {
	scale := GetFullChromaticScale()
	for i := range scale {
		tni := getTemplateNotesInstance()
		tn := tni.getTemplateNote(&scale[i])
		tn.saveTonic(&scale[i])
		assert.Equal(t, tn, tn.getTonic(), "expected: %v, actual: %v", tn, tn.getTonic())

		currentNote := tn
		for unsafe.Pointer(currentNote.next) != unsafe.Pointer(tn) {
			currentNote = currentNote.getNext()
			tonic := currentNote.getTonic()
			assert.Equal(t, tn, tonic, "expected: %v, actual: %v", tn, tonic)
		}
	}
}

func Test_getPreviousUsedBaseNote(t *testing.T) {
	scale := GetFullChromaticScale()
	for i, note := range scale {
		tni := getTemplateNotesInstance()
		tn := tni.getTemplateNote(&scale[i])
		tn.saveResultingNote(&scale[i])

		currentNote := tn
		for unsafe.Pointer(currentNote.next) != unsafe.Pointer(tn) {
			currentNote = currentNote.getNext()
			pbn := currentNote.getPreviousUsedBaseNote()
			assert.Equal(t, note.baseNote(), pbn, "expected: %v, actual: %v", note.baseNote(), pbn)
		}
	}
}

func Test_NextBaseNote(t *testing.T) {
	tni := getTemplateNotesInstance()

	testCases := []struct {
		tns  []*Note
		want *Note
	}{
		{
			tns:  []*Note{newNote(C), newNote(CSHARP), newNote(CFLAT)},
			want: newNote(D),
		},
		{
			tns:  []*Note{newNote(D), newNote(DSHARP), newNote(DFLAT)},
			want: newNote(E),
		},
		{
			tns:  []*Note{newNote(E), newNote(ESHARP), newNote(EFLAT)},
			want: newNote(F),
		},
		{
			tns:  []*Note{newNote(F), newNote(FSHARP), newNote(FFLAT)},
			want: newNote(G),
		},
		{
			tns:  []*Note{newNote(G), newNote(GSHARP), newNote(GFLAT)},
			want: newNote(A),
		},
		{
			tns:  []*Note{newNote(A), newNote(ASHARP), newNote(AFLAT)},
			want: newNote(B),
		},
		{
			tns:  []*Note{newNote(B), newNote(BSHARP), newNote(BFLAT)},
			want: newNote(C),
		},
	}

	for _, testCase := range testCases {
		for _, tn := range testCase.tns {
			tni.SetLastUsedBaseNote(tn)
			nextBase := tni.NextBaseNote()
			assert.Equal(t, testCase.want, nextBase, "expected: %+v, actual: %+v", testCase.want, nextBase)
		}
	}
}
