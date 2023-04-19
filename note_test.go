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
			CFLAT,
			CFLAT2,
			CSHARP,
			CSHARP2,
			D,
			DFLAT,
			DFLAT2,
			DSHARP,
			DSHARP2,
			E,
			EFLAT,
			EFLAT2,
			ESHARP,
			ESHARP2,
			F,
			FFLAT,
			FFLAT2,
			FSHARP,
			FSHARP2,
			G,
			GFLAT,
			GFLAT2,
			GSHARP,
			GSHARP2,
			A,
			AFLAT,
			AFLAT2,
			ASHARP,
			ASHARP2,
			B,
			BFLAT,
			BFLAT2,
			BSHARP,
			BSHARP2,
		}
		var newNote *Note
		var err error
		for _, noteName := range noteNames {
			// setup: create a Note with a valid name
			newNote, err = NewNote(noteName)
			assert.NoError(t, err)
			assert.NotNil(t, newNote, "expected note from note name %s", noteName)

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
		assert.ErrorIs(t, err, ErrNoteNameUnknown)
		assert.Nil(t, newNote)
	})
}

func TestMustNewNote(t *testing.T) {
	t.Run("TestNewNote: valid note name", func(t *testing.T) {
		noteNames := []NoteName{
			C,
			CFLAT,
			CFLAT2,
			CSHARP,
			CSHARP2,
			D,
			DFLAT,
			DFLAT2,
			DSHARP,
			DSHARP2,
			E,
			EFLAT,
			EFLAT2,
			ESHARP,
			ESHARP2,
			F,
			FFLAT,
			FFLAT2,
			FSHARP,
			FSHARP2,
			G,
			GFLAT,
			GFLAT2,
			GSHARP,
			GSHARP2,
			A,
			AFLAT,
			AFLAT2,
			ASHARP,
			ASHARP2,
			B,
			BFLAT,
			BFLAT2,
			BSHARP,
			BSHARP2,
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

func TestNote_AlterUp(t *testing.T) {
	// Test case when note is nil
	var n *Note
	if n.AlterUp() != nil {
		t.Errorf("AlterUp on nil note should return nil")
	}

	// Test case when NoteName doesn't end with any symbols
	n = &Note{name: B}
	if n.AlterUp().name != BSHARP {
		t.Errorf("AlterUp on note B should result in B#, got: %s", n.name)
	}

	// Test case when NoteName ends with AlterSymbolSharp
	n = &Note{name: BSHARP}
	if n.AlterUp().name != BSHARP2 {
		t.Errorf("AlterUp on note B# should result in B##, got: %s", n.name)
	}

	// Test case when NoteName ends with AlterSymbolFlat
	n = &Note{name: BFLAT}
	if n.AlterUp().name != B {
		t.Errorf("AlterUp on note Bb should result in B, got: %s", n.name)
	}

	// Test case when NoteName ends with AlterSymbolFlat twice
	n = &Note{name: BFLAT2}
	if n.AlterUp().name != BFLAT {
		t.Errorf("AlterUp on note Bbb should result in Bb, got: %s", n.name)
	}
}

func TestNote_AlterDown(t *testing.T) {
	// Test case when note is nil
	var n *Note
	if n.AlterDown() != nil {
		t.Errorf("AlterDown on nil note should return nil")
	}

	// Test case when NoteName doesn't end with any symbols
	n = &Note{name: B}
	if n.AlterDown().name != BFLAT {
		t.Errorf("AlterDown on note B should result in Bb, got: %s", n.name)
	}

	// Test case when NoteName ends with AlterSymbolSharp
	n = &Note{name: BSHARP}
	if n.AlterDown().name != B {
		t.Errorf("AlterDown on note B# should result in B, got: %s", n.name)
	}

	// Test case when NoteName ends with AlterSymbolFlat
	n = &Note{name: BFLAT}
	if n.AlterDown().name != BFLAT2 {
		t.Errorf("AlterDown on note Bb should result in Bbb, got: %s", n.name)
	}

	// Test case when NoteName ends with AlterSymbolSharp twice
	n = &Note{name: BSHARP2}
	if n.AlterDown().name != BSHARP {
		t.Errorf("AlterDown on note B## should result in Bb# got: %s", n.name)
	}
}

func TestNote_AlterUpBy(t *testing.T) {
	t.Run("Note_AlterUpBy: positive case 1", func(t *testing.T) {
		n := newNote(C)
		alteredNote := n.AlterUpBy(2)
		expectedNote := newNote(CSHARP2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Error("AlterUpBy should alter the note up by the provided value")
		}
	})

	t.Run("Note_AlterUpBy: positive case 2", func(t *testing.T) {
		n := newNote(CFLAT2)
		alteredNote := n.AlterUpBy(4)
		expectedNote := newNote(CSHARP2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Error("AlterUpBy should alter the note up by the provided value")
		}
	})

	t.Run("Note_AlterUpBy: zero times", func(t *testing.T) {
		n := newNote(C)
		if n.AlterUpBy(0) != n {
			t.Error("AlterUpBy should return the same note for 0 alterations")
		}
	})

	t.Run("Note_AlterUpBy: nil note", func(t *testing.T) {
		var n *Note
		if n.AlterUpBy(1) != nil {
			t.Error("AlterUpBy should return nil for nil note")
		}
	})
}

func TestNote_AlterDownBy(t *testing.T) {
	t.Run("Note_AlterDownBy: positive case 1", func(t *testing.T) {
		n := newNote(C)
		alteredNote := n.AlterDownBy(2)
		expectedNote := newNote(CFLAT2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Errorf("AlterDownBy expected: %s, actual: %s", expectedNote.Name(), alteredNote.Name())
		}
	})

	t.Run("Note_AlterDownBy: positive case 2", func(t *testing.T) {
		n := newNote(CSHARP2)
		alteredNote := n.AlterDownBy(4)
		expectedNote := newNote(CFLAT2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Errorf("AlterDownBy expected: %s, actual: %s", expectedNote.Name(), alteredNote.Name())
		}
	})

	t.Run("Note_AlterDownBy: zero times", func(t *testing.T) {
		n := newNote(C)
		if n.AlterDownBy(0) != n {
			t.Error("AlterDownBy should return the same note for 0 alterations")
		}
	})

	t.Run("Note_AlterDownBy: nil note", func(t *testing.T) {
		var n *Note
		if n.AlterDownBy(1) != nil {
			t.Error("AlterDownBy should return nil for nil note")
		}
	})
}

func TestNote_BaseName(t *testing.T) {
	testCases := []struct {
		note *Note
		want NoteName
	}{
		{
			note: newNote(CSHARP2),
			want: C,
		},
		{
			note: newNote(CSHARP),
			want: C,
		},
		{
			note: newNote(C),
			want: C,
		},
		{
			note: newNote(CFLAT),
			want: C,
		},
		{
			note: newNote(CFLAT2),
			want: C,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.BaseName(), "expected note name: %s, result: %s", testCase.want, testCase.note.BaseName())
	}
}

func TestNote_BaseNote(t *testing.T) {
	testCases := []struct {
		note *Note
		want *Note
	}{
		{
			note: newNote(CSHARP2),
			want: C.MustMakeNote(),
		},
		{
			note: newNote(CSHARP),
			want: C.MustMakeNote(),
		},
		{
			note: newNote(C),
			want: C.MustMakeNote(),
		},
		{
			note: newNote(CFLAT),
			want: C.MustMakeNote(),
		},
		{
			note: newNote(CFLAT2),
			want: C.MustMakeNote(),
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.baseNote(), "expected note: %+v, result: %+v", testCase.want, testCase.note.baseNote())
	}
}
