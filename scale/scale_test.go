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
		note.C.NewNote(),
		note.DFLAT.NewNote(),
		note.CSHARP.NewNote(),
		note.D.NewNote(),
		note.EFLAT.NewNote(),
		note.DSHARP.NewNote(),
		note.E.NewNote(),
		note.F.NewNote(),
		note.GFLAT.NewNote(),
		note.FSHARP.NewNote(),
		note.G.NewNote(),
		note.AFLAT.NewNote(),
		note.GSHARP.NewNote(),
		note.A.NewNote(),
		note.BFLAT.NewNote(),
		note.ASHARP.NewNote(),
		note.B.NewNote(),
	}

	actual := GetFullChromaticScale()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetFullChromaticScale() returned unexpected result.\nExpected: %v\nActual: %v", expected, actual)
	}
}
