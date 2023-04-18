package muse

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateScale(t *testing.T) {
	t.Run("Test a mode with just one note", func(t *testing.T) {
		var m *Mode
		var got Scale

		m = &Mode{
			name: "Custom mode",
			degree: &Degree{
				number: 1,
				note: &Note{
					name: C,
				},
			},
		}
		exp := Scale([]Note{
			{
				name: C,
			},
		})
		got = m.GenerateScale(false)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Unexpected results for GenerateScale; expected %v but got %v", exp, got)
		}
		got = m.GenerateScale(true)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Unexpected results for GenerateScale; expected %v but got %v", exp, got)
		}
	})

	t.Run("Test a mode with multiple notes", func(t *testing.T) {
		var m *Mode
		var got Scale
		firstDegree := generateDegreesWithNotes(true, TemplateDorian(), &Note{name: D})
		m = &Mode{name: ModeNameDorian, degree: firstDegree}

		exp := NewScaleFromNoteNames(D, E, F, G, A, B, C)

		got = m.GenerateScale(false)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Unexpected results for GenerateScale; expected %v but got %v", exp, got)
		}

		exp = NewScaleFromNoteNames(C, B, A, G, F, E, D)

		got = m.GenerateScale(true)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Unexpected results for GenerateScale; expected %v but got %v", exp, got)
		}
	})

	t.Run("Test an invalid mode", func(t *testing.T) {
		var m *Mode
		exp := Scale(nil)
		got := m.GenerateScale(false)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Unexpected results for GenerateScale; expected %v but got %v", exp, got)
		}
	})
}

func TestGetFullChromaticScale(t *testing.T) {
	expected := Scale{
		*newNote(C),
		*newNote(DFLAT),
		*newNote(CSHARP),
		*newNote(D),
		*newNote(EFLAT),
		*newNote(DSHARP),
		*newNote(E),
		*newNote(F),
		*newNote(GFLAT),
		*newNote(FSHARP),
		*newNote(G),
		*newNote(AFLAT),
		*newNote(GSHARP),
		*newNote(A),
		*newNote(BFLAT),
		*newNote(ASHARP),
		*newNote(B),
	}

	actual := GetFullChromaticScale()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetFullChromaticScale() returned unexpected result.\nExpected: %v\nActual: %v", expected, actual)
	}
}

func TestGetAllPossibleNotes(t *testing.T) {
	testCases := []struct {
		name          string
		alterations   uint8
		expectedScale Scale
	}{
		{"No Alterations", 0, NewScaleFromNoteNames(C, D, E, F, G, A, B)},
		{"One Alteration", 1, NewScaleFromNoteNames(C, CFLAT, CSHARP, D, DFLAT, DSHARP, E, EFLAT, ESHARP, F, FFLAT, FSHARP, G, GFLAT, GSHARP, A, AFLAT, ASHARP, B, BFLAT, BSHARP)},
		{"Two Alterations", 2, NewScaleFromNoteNames(C, CFLAT, CFLAT2, CSHARP, CSHARP2, D, DFLAT, DFLAT2, DSHARP, DSHARP2, E, EFLAT, EFLAT2, ESHARP, ESHARP2, F, FFLAT, FFLAT2, FSHARP, FSHARP2, G, GFLAT, GFLAT2, GSHARP, GSHARP2, A, AFLAT, AFLAT2, ASHARP, ASHARP2, B, BFLAT, BFLAT2, BSHARP, BSHARP2)},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(testCase.name, func(t *testing.T) {
			resultingScale := GetAllPossibleNotes(tc.alterations)
			assert.ElementsMatch(t, tc.expectedScale, resultingScale)
		})
	}
}
