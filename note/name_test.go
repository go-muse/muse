package note

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoteName_NewNote(t *testing.T) {
	t.Run("NoteName_NewNote: positive cases", func(t *testing.T) {
		notes := GetNotesWithAlterations(GetSetFullChromatic(), 2)
		for _, note := range notes {
			assert.Equal(t, note.name, note.name.MustNewNote().Name())
		}
	})

	t.Run("NoteName_NewNote: negative cases", func(t *testing.T) {
		names := Names{"", "incorrect_name", "Abbb#", "C##b#", "a"}
		for _, name := range names {
			assert.Panics(t, func() { _ = name.MustNewNote() }, "expected panic on note name %s", name)
		}
	})
}

func TestNoteName_String(t *testing.T) {
	expectedNoteName := "C"
	nn := Name(expectedNoteName)

	result := nn.String()
	if _, ok := interface{}(result).(string); !ok {
		t.Errorf("note name Stringer return value: %s is not string: %s", result, reflect.TypeOf(result))
	}

	if expectedNoteName != result {
		t.Errorf("note name Stringer expected: %s actual: %s", expectedNoteName, result)
	}
}

func TestNoteName_MakeNote(t *testing.T) {
	var newNote *Note
	for _, note := range GetNotesWithAlterations(GetSetFullChromatic(), 2) {
		newNote = note.Name().MustMakeNote()
		assert.Equal(t, note, newNote, "expected note: %+v, actual note: %+v", note, newNote)
	}
}

func TestNoteName_MustMakeNote(t *testing.T) {
	// Test case 1: Valid note name
	if C.MustMakeNote() == nil {
		t.Errorf("MustMakeNote(%s) returned nil, expected *Note", Name("C"))
	}

	// Test case 2: Invalid note name
	nn2 := Name("X")
	assert.Panics(t, func() { _ = nn2.MustMakeNote() }, "note name: %s", nn2)
}

func TestNoteName_Validate(t *testing.T) {
	type testCase struct {
		noteName Name
		want     bool
	}

	testCases := []testCase{
		{noteName: "AB", want: false},
		{noteName: "Ac", want: false},
		{noteName: "A#b", want: false},
		{noteName: "Ab#", want: false},
		{noteName: "Ab#b", want: false},
		{noteName: "Ab##", want: false},
		{noteName: "A##b", want: false},
		{noteName: "A#bb", want: false},
		{noteName: "b", want: false},
		{noteName: "bb", want: false},
		{noteName: "bbb", want: false},
		{noteName: "bb#", want: false},
		{noteName: "b##", want: false},
		{noteName: "#", want: false},
		{noteName: "##", want: false},
		{noteName: "##b", want: false},
		{noteName: "#bb", want: false},
		{noteName: "###", want: false},
	}

	for _, note := range GetNotesWithAlterations(GetSetFullChromatic(), 2) {
		testCases = append(testCases, testCase{note.Name(), true})
	}

	for _, testCase := range testCases {
		if testCase.want {
			require.NoError(t, testCase.noteName.Validate(), "note name %s validation expected: true, actual: %+v", testCase.noteName, testCase.noteName.Validate())
		} else {
			require.ErrorIs(t, ErrNoteNameUnknown, testCase.noteName.Validate(), "note name %s validation expected: false, actual: %+v", testCase.noteName, testCase.noteName.Validate())
		}
	}
}
