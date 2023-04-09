package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertNote(t *testing.T) {
	mode := Mode{}
	note1 := &Note{
		name: "note 1",
	}

	mode.InsertNote(note1, 0)
	assert.True(t, mode.GetFirstDegree().note.IsEqualByName(note1))
	assert.Nil(t, mode.GetFirstDegree().previous)
	assert.Nil(t, mode.GetFirstDegree().next)

	note2 := &Note{
		name: "note 2",
	}
	mode.InsertNote(note2, 2)
	assert.True(t, mode.GetFirstDegree().note.IsEqualByName(note1))
	assert.Nil(t, mode.GetFirstDegree().previous)
	assert.NotNil(t, mode.GetFirstDegree().next)

	assert.True(t, mode.GetFirstDegree().GetNext().note.IsEqualByName(note2))
	assert.NotNil(t, mode.GetFirstDegree().GetNext().previous)
	assert.True(t, mode.GetFirstDegree().GetNext().GetPrevious().note.IsEqualByName(note1))
	assert.Nil(t, mode.GetFirstDegree().GetNext().next)
}

func TestMakeNewMode(t *testing.T) {
	t.Run("make natural minor mode", func(t *testing.T) {
		tonalCenter := C
		mode, err := MakeNewMode(ModeNameNaturalMinor, tonalCenter)
		assert.NoError(t, err)
		assert.NotNil(t, mode)

		firstDegree := mode.GetFirstDegree()

		firstNote, err := NewNote(tonalCenter)
		assert.NoError(t, err)
		assert.True(t, firstDegree.note.IsEqualByName(firstNote))

		naturalMinorFromC := []*Note{
			MustNewNote(D),
			MustNewNote(EFLAT),
			MustNewNote(F),
			MustNewNote(G),
			MustNewNote(AFLAT),
			MustNewNote(BFLAT),
		}

		for _, note := range naturalMinorFromC {
			firstDegree = firstDegree.GetNext()
			assert.NotNil(t, firstDegree)
			assert.NotNil(t, firstDegree.note)
			assert.True(t, firstDegree.note.IsEqualByName(note), "expected: %s, actual: %s", note.Name(), firstDegree.note.Name())
		}
	})

	t.Run("make natural major mode", func(t *testing.T) {
		tonalCenter := C
		mode, err := MakeNewMode(ModeNameNaturalMajor, tonalCenter)
		assert.NoError(t, err)
		assert.NotNil(t, mode)

		firstDegree := mode.GetFirstDegree()
		assert.NotNil(t, mode)

		firstNote, err := NewNote(tonalCenter)
		assert.NoError(t, err)
		assert.True(t, firstDegree.note.IsEqualByName(firstNote))

		naturalMajorFromC := []*Note{
			MustNewNote(D),
			MustNewNote(E),
			MustNewNote(F),
			MustNewNote(G),
			MustNewNote(A),
			MustNewNote(B),
		}

		for _, note := range naturalMajorFromC {
			firstDegree = firstDegree.GetNext()
			assert.NotNil(t, firstDegree)
			assert.NotNil(t, firstDegree.note)
			assert.True(t, firstDegree.note.IsEqualByName(note))
		}
	})

	t.Run("make natural major mode B", func(t *testing.T) {
		tonalCenter := B
		mode, err := MakeNewMode(ModeNameNaturalMajor, tonalCenter)
		assert.NoError(t, err)
		assert.NotNil(t, mode)

		firstDegree := mode.GetFirstDegree()
		assert.NotNil(t, mode)

		firstNote, err := NewNote(tonalCenter)
		assert.NoError(t, err)
		assert.True(t, firstDegree.note.IsEqualByName(firstNote))

		naturalMajorFromB := []*Note{
			MustNewNote(CSHARP),
			MustNewNote(DSHARP),
			MustNewNote(E),
			MustNewNote(FSHARP),
			MustNewNote(GSHARP),
			MustNewNote(ASHARP),
		}

		for _, note := range naturalMajorFromB {
			firstDegree = firstDegree.GetNext()
			assert.NotNil(t, firstDegree)
			assert.NotNil(t, firstDegree.note)
			assert.True(t, firstDegree.note.IsEqualByName(note), "expect note %s, actual: %s", firstDegree.Note().Name(), note.Name())
		}
	})
}

func TestOpenCircleOfDegreesMethod(t *testing.T) {
	// Create a Mode with a circle of degrees
	degree1 := &Degree{number: 1}
	degree2 := &Degree{number: 2}
	degree3 := &Degree{number: 3}
	degree1.SetNext(degree2)
	degree2.SetNext(degree3)
	degree3.SetNext(degree1)
	degree1.SetPrevious(degree3)
	degree2.SetPrevious(degree1)
	degree3.SetPrevious(degree2)
	mode := &Mode{
		name:   "Test Mode",
		degree: degree1,
	}

	// Call OpenCircleOfDegrees() to break the circle
	mode.OpenCircleOfDegrees()

	// Ensure that the circle is broken
	if mode.GetFirstDegree().GetPrevious() != nil {
		t.Errorf("First degree still has a previous degree")
	}
	if mode.GetLastDegree().GetNext() != nil {
		t.Errorf("Last degree still has a next degree")
	}
}

func TestIsClosedCircleOfDegrees(t *testing.T) {
	// Creating degrees
	degree1 := &Degree{number: 1}
	degree2 := &Degree{number: 2}
	degree3 := &Degree{number: 3}

	// Setting links between degrees
	degree1.SetNext(degree2)
	degree2.SetPrevious(degree1)

	degree2.SetNext(degree3)
	degree3.SetPrevious(degree2)

	// Creating a mode instance
	mode, _ := MakeNewCustomModeWithDegree("test mode", degree1)

	if mode.IsClosedCircleOfDegrees() {
		t.Errorf("IsClosedCircleOfDegrees failed: expected false, got true")
	}

	// Circle the mode
	degree3.SetNext(degree1)
	degree1.SetPrevious(degree3)

	if !mode.IsClosedCircleOfDegrees() {
		t.Errorf("IsClosedCircleOfDegrees failed: expected true, got false")
	}
}
