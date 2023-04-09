package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewNote(t *testing.T) {
	expectedNoteName := C
	assert.Equal(t, newNote(C).name, expectedNoteName)
}

func TestNewNote(t *testing.T) {
	t.Run("TestNewNote: valid note name", func(t *testing.T) {
		noteNames := []NoteName{
			C,
			DFLAT,
			CSHARP,
			D,
			EFLAT,
			DSHARP,
			E,
			F,
			GFLAT,
			FSHARP,
			G,
			AFLAT,
			GSHARP,
			A,
			BFLAT,
			ASHARP,
			B,
		}
		var newNote *Note
		var err error
		for _, noteName := range noteNames {
			// setup: create a Note with a valid name
			newNote, err = NewNote(noteName)
			assert.NoError(t, err)

			// assert that the returned name matches the expected name
			if newNote.name != noteName {
				t.Errorf("Expected note name to be '%s', but got '%s'", noteName, newNote.name)
			}
		}
	})

	t.Run("TestNewNote: invalid note name", func(t *testing.T) {
		// setup: create a Note with invalid name
		expectedName := NoteName("Hello")
		newNote, err := NewNote(expectedName)
		// assert that the returned error matches the expected error
		assert.ErrorIs(t, err, ErrUnknownNoteName)
		assert.Nil(t, newNote)
	})
}

func TestMustNewNote(t *testing.T) {
	t.Run("TestNewNote: valid note name", func(t *testing.T) {
		noteNames := []NoteName{
			C,
			DFLAT,
			CSHARP,
			D,
			EFLAT,
			DSHARP,
			E,
			F,
			GFLAT,
			FSHARP,
			G,
			AFLAT,
			GSHARP,
			A,
			BFLAT,
			ASHARP,
			B,
		}
		var newNote *Note
		for _, noteName := range noteNames {
			// assert that the function works without panic
			assert.NotPanics(t, func() { newNote = MustNewNote(noteName) }) //nolint:scopelint

			// assert that the returned name matches the expected name
			if newNote.name != noteName {
				t.Errorf("Expected note name to be '%s', but got '%s'", noteName, newNote.name)
			}
		}
	})

	t.Run("MustNewNote: invalid note name", func(t *testing.T) {
		// setup: create a Note with invalid name
		expectedName := NoteName("Hello")
		// assert that the function works without panic
		assert.Panics(t, func() { _ = MustNewNote(expectedName) })
	})
}

func TestNoteName(t *testing.T) {
	// setup: create a Note with a known name
	expectedName := C
	note := &Note{name: expectedName}

	// execute the Name() method
	actualName := note.Name()

	// assert that the returned name matches the expected name
	assert.Equal(t, expectedName, actualName, "expected note name: %s, actual: %s", expectedName, actualName)
}

func TestNoteIsEqualByName(t *testing.T) {
	testCases := []struct {
		note1, note2 *Note
		want         bool
	}{
		{
			note1: &Note{name: C},
			note2: &Note{name: C},
			want:  true,
		},
		{
			note1: &Note{name: C},
			note2: &Note{name: D},
			want:  false,
		},
		{
			note1: &Note{name: C},
			note2: nil,
			want:  false,
		},
		{
			note1: nil,
			note2: nil,
			want:  false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note1.IsEqualByName(testCase.note2))
	}
}

func TestNoteCopy(t *testing.T) {
	note1 := &Note{name: C}
	note2 := note1.Copy()

	// Test the pointer address is not the same
	if note1 == note2 {
		t.Error("Expected new note instance, but got the same pointer address")
	}

	// Test the notes have the same name
	if note1.Name() != note2.Name() {
		t.Error("Expected new note with same name, but got different names")
	}

	// Test nil input returns nil
	if nilNote := (*Note)(nil).Copy(); nilNote != nil {
		t.Error("Expected nil output for nil input, but got non-nil result")
	}
}
