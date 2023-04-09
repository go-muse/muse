package muse

import (
	"reflect"
	"testing"
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
		exp := Scale([]*Note{
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

		exp := Scale([]*Note{
			{
				name: D,
			},
			{
				name: E,
			},
			{
				name: F,
			},
			{
				name: G,
			},
			{
				name: A,
			},
			{
				name: B,
			},
			{
				name: C,
			},
		})
		got = m.GenerateScale(false)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("Unexpected results for GenerateScale; expected %v but got %v", exp, got)
		}

		exp = Scale([]*Note{
			{
				name: C,
			},
			{
				name: B,
			},
			{
				name: A,
			},
			{
				name: G,
			},
			{
				name: F,
			},
			{
				name: E,
			},
			{
				name: D,
			},
		})
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
		newNote(C),
		newNote(DFLAT),
		newNote(CSHARP),
		newNote(D),
		newNote(EFLAT),
		newNote(DSHARP),
		newNote(E),
		newNote(F),
		newNote(GFLAT),
		newNote(FSHARP),
		newNote(G),
		newNote(AFLAT),
		newNote(GSHARP),
		newNote(A),
		newNote(BFLAT),
		newNote(ASHARP),
		newNote(B),
	}

	actual := GetFullChromaticScale()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetFullChromaticScale() returned unexpected result.\nExpected: %v\nActual: %v", expected, actual)
	}
}
