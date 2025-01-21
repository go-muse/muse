package scale

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/note"
)

func TestString(t *testing.T) {
	testCases := []struct {
		name     string
		notes    note.Names
		expected string
	}{
		{"No Alterations", note.Names{note.C}, "[C]"},
		{"No Alterations", note.Names{note.C, note.D}, "[C D]"},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(testCase.name, func(t *testing.T) {
			result := MustNewScaleFromNoteNames(testCase.notes...).String()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetFullChromaticScale(t *testing.T) {
	expected := Scale{
		note.C.MustNewNote(),
		note.DFLAT.MustNewNote(),
		note.CSHARP.MustNewNote(),
		note.D.MustNewNote(),
		note.EFLAT.MustNewNote(),
		note.DSHARP.MustNewNote(),
		note.E.MustNewNote(),
		note.F.MustNewNote(),
		note.GFLAT.MustNewNote(),
		note.FSHARP.MustNewNote(),
		note.G.MustNewNote(),
		note.AFLAT.MustNewNote(),
		note.GSHARP.MustNewNote(),
		note.A.MustNewNote(),
		note.BFLAT.MustNewNote(),
		note.ASHARP.MustNewNote(),
		note.B.MustNewNote(),
	}

	actual := GetFullChromaticScale()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetFullChromaticScale() returned unexpected result.\nExpected: %v\nActual: %v", expected, actual)
	}
}
