package muse

import (
	"crypto/rand"
	"math/big"
	"testing"
	"unsafe"

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

		firstNote, err := NewNote(tonalCenter, OctaveNumberDefault)
		assert.NoError(t, err)
		assert.True(t, firstDegree.note.IsEqualByName(firstNote))

		naturalMinorFromC := []*Note{
			MustNewNoteWithoutOctave(D),
			MustNewNoteWithoutOctave(EFLAT),
			MustNewNoteWithoutOctave(F),
			MustNewNoteWithoutOctave(G),
			MustNewNoteWithoutOctave(AFLAT),
			MustNewNoteWithoutOctave(BFLAT),
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

		firstNote, err := NewNote(tonalCenter, OctaveNumberDefault)
		assert.NoError(t, err)
		assert.True(t, firstDegree.note.IsEqualByName(firstNote))

		naturalMajorFromC := []*Note{
			MustNewNoteWithoutOctave(D),
			MustNewNoteWithoutOctave(E),
			MustNewNoteWithoutOctave(F),
			MustNewNoteWithoutOctave(G),
			MustNewNoteWithoutOctave(A),
			MustNewNoteWithoutOctave(B),
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

		firstNote, err := NewNote(tonalCenter, OctaveNumberDefault)
		assert.NoError(t, err)
		assert.True(t, firstDegree.note.IsEqualByName(firstNote))

		naturalMajorFromB := []*Note{
			MustNewNoteWithoutOctave(CSHARP),
			MustNewNoteWithoutOctave(DSHARP),
			MustNewNoteWithoutOctave(E),
			MustNewNoteWithoutOctave(FSHARP),
			MustNewNoteWithoutOctave(GSHARP),
			MustNewNoteWithoutOctave(ASHARP),
		}

		for _, note := range naturalMajorFromB {
			firstDegree = firstDegree.GetNext()
			assert.NotNil(t, firstDegree)
			assert.NotNil(t, firstDegree.note)
			assert.True(t, firstDegree.note.IsEqualByName(note), "expect note %s, actual: %s", firstDegree.Note().Name(), note.Name())
		}
	})
}

func TestMustMakeNewMode(t *testing.T) {
	// Test case with valid inputs
	var mode *Mode
	store := InitModeTemplatesStore()
	for modeName := range store {
		assert.NotPanics(t, func() { mode = MustMakeNewMode(modeName, A) }) //nolint:scopelint
		if mode == nil {
			t.Errorf("Expected mode to be created successfully, but got nil")
		}
	}

	// Negative test case with invalid mode name
	withPanic := func() {
		MustMakeNewMode(ModeName(""), NoteName("C"))
	}
	assert.Panics(t, withPanic)

	// Negative test case with invalid note name
	withPanic = func() {
		MustMakeNewMode(ModeName("testMode"), NoteName(""))
	}
	assert.Panics(t, withPanic)
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

func TestSortByAbsoluteModalPositions(t *testing.T) {
	n10, err := rand.Int(rand.Reader, big.NewInt(10))
	assert.NoError(t, err)
	n20, err := rand.Int(rand.Reader, big.NewInt(10))
	assert.NoError(t, err)

	getDegrees := func() (*Degree, *Degree) {
		firstDegree := &Degree{number: 1, absoluteModalPosition: NewModalPositionByWeight(Weight(n10.Int64() - 5))}
		lastDegree := firstDegree
		for i := DegreeNum(2); i <= DegreesInTonality; i++ {
			degree := &Degree{number: i, absoluteModalPosition: NewModalPositionByWeight(Weight(n20.Int64() - 10))}
			lastDegree.AttachNext(degree)
			lastDegree = degree
		}

		return firstDegree, lastDegree
	}

	testingFunc := func(t *testing.T, firstSortedDegree *Degree) {
		t.Helper()
		iterator := firstSortedDegree.IterateOneRound(false)
		firstDegree := <-iterator
		var comparison bool
		for degree := range iterator {
			if degree.NextExists() {
				comparison = degree.absoluteModalPosition.GetWeight() <= degree.GetNext().absoluteModalPosition.GetWeight()
				if unsafe.Pointer(degree.GetNext()) != unsafe.Pointer(firstDegree) {
					assert.True(t, comparison, "current - degree Num: %d, w: %d, next - degree Num: %d, w: %d", degree.Number(), degree.absoluteModalPosition.GetWeight(), degree.GetNext().Number, degree.GetNext().absoluteModalPosition.GetWeight())
				} else {
					assert.True(t, !comparison, "current - degree Num: %d, w: %d, next - degree Num: %d, w: %d", degree.Number(), degree.absoluteModalPosition.GetWeight(), degree.GetNext().Number, degree.GetNext().absoluteModalPosition.GetWeight())
				}
			}
		}
	}

	t.Run("test sort by AMP of cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		mode, err := MakeNewCustomModeWithDegree("custom mode", firstDegree)
		assert.NoError(t, err)
		mode.SortByAbsoluteModalPositions(false)
		testingFunc(t, mode.GetFirstDegree())
	})

	t.Run("test sort by AMP of uncycled degrees chain", func(t *testing.T) {
		firstDegree, _ := getDegrees()
		mode, err := MakeNewCustomModeWithDegree("custom mode", firstDegree)
		assert.NoError(t, err)
		mode.SortByAbsoluteModalPositions(false)
		testingFunc(t, mode.GetFirstDegree())
	})

	t.Run("test sort by AMP in case of degree without AMP", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		firstDegree.GetNext().absoluteModalPosition = nil // just one random degree without set absolute modal position
		mode, err := MakeNewCustomModeWithDegree("custom mode", firstDegree)
		assert.NoError(t, err)
		mode.SortByAbsoluteModalPositions(false)
		assert.Equal(t, firstDegree, mode.GetFirstDegree())

		// check for the old order
		currentDegree := firstDegree
		for modeDegree := range mode.IterateOneRound(false) {
			assert.Equal(t, currentDegree, modeDegree)
			currentDegree = currentDegree.GetNext()
		}
	})
}

func TestMode_Contains(t *testing.T) {
	// Create a new mode
	mode, err := MakeNewCustomMode(ModeTemplate{2, 2, 1, 2, 2, 2, 1}, "C", ModeNameNaturalMajor)
	assert.NoError(t, err)

	testCases := []struct {
		note *Note
		want bool
	}{
		{C.MustMakeNote(), true},
		{D.MustMakeNote(), true},
		{E.MustMakeNote(), true},
		{F.MustMakeNote(), true},
		{G.MustMakeNote(), true},
		{A.MustMakeNote(), true},
		{B.MustMakeNote(), true},

		{CFLAT.MustMakeNote(), false},
		{CFLAT2.MustMakeNote(), false},
		{CSHARP.MustMakeNote(), false},
		{CSHARP2.MustMakeNote(), false},
		{DFLAT.MustMakeNote(), false},
		{DFLAT2.MustMakeNote(), false},
		{DSHARP.MustMakeNote(), false},
		{DSHARP2.MustMakeNote(), false},
		{EFLAT.MustMakeNote(), false},
		{EFLAT2.MustMakeNote(), false},
		{ESHARP.MustMakeNote(), false},
		{ESHARP2.MustMakeNote(), false},
		{FFLAT.MustMakeNote(), false},
		{FFLAT2.MustMakeNote(), false},
		{FSHARP.MustMakeNote(), false},
		{FSHARP2.MustMakeNote(), false},
		{GFLAT.MustMakeNote(), false},
		{GFLAT2.MustMakeNote(), false},
		{GSHARP.MustMakeNote(), false},
		{GSHARP2.MustMakeNote(), false},
		{AFLAT.MustMakeNote(), false},
		{AFLAT2.MustMakeNote(), false},
		{ASHARP.MustMakeNote(), false},
		{ASHARP2.MustMakeNote(), false},
		{BFLAT.MustMakeNote(), false},
		{BFLAT2.MustMakeNote(), false},
		{BSHARP.MustMakeNote(), false},
		{BSHARP2.MustMakeNote(), false},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, mode.Contains(testCase.note), "expected note in mode: %f, actual: %f", testCase.want, mode.Contains(testCase.note))
	}
}

func TestMode_IsEqual(t *testing.T) {
	type testCase struct {
		mode1, mode2 *Mode
	}

	t.Run("TestMode_IsEqual testing equal modes", func(t *testing.T) {
		testCases := []testCase{
			// same mode, different tonics
			{MustMakeNewMode(ModeNameAeolian, C), MustMakeNewMode(ModeNameAeolian, C)},
			{MustMakeNewMode(ModeNameAeolian, D), MustMakeNewMode(ModeNameAeolian, D)},
			{MustMakeNewMode(ModeNameAeolian, E), MustMakeNewMode(ModeNameAeolian, E)},
			{MustMakeNewMode(ModeNameAeolian, F), MustMakeNewMode(ModeNameAeolian, F)},
			{MustMakeNewMode(ModeNameAeolian, G), MustMakeNewMode(ModeNameAeolian, G)},
			{MustMakeNewMode(ModeNameAeolian, A), MustMakeNewMode(ModeNameAeolian, A)},
			{MustMakeNewMode(ModeNameAeolian, B), MustMakeNewMode(ModeNameAeolian, B)},

			// same tonics, different modes
			{MustMakeNewMode(ModeNameIonian, C), MustMakeNewMode(ModeNameIonian, C)},
			{MustMakeNewMode(ModeNameAeolian, C), MustMakeNewMode(ModeNameAeolian, C)},
			{MustMakeNewMode(ModeNameLydian, C), MustMakeNewMode(ModeNameLydian, C)},
			{MustMakeNewMode(ModeNameDorian, C), MustMakeNewMode(ModeNameDorian, C)},
			{MustMakeNewMode(ModeNamePhrygian, C), MustMakeNewMode(ModeNamePhrygian, C)},
			{MustMakeNewMode(ModeNameLocrian, C), MustMakeNewMode(ModeNameLocrian, C)},
			{MustMakeNewMode(ModeNameMixoLydian, C), MustMakeNewMode(ModeNameMixoLydian, C)},
		}

		for _, testCase := range testCases {
			assert.True(t, testCase.mode1.IsEqual(testCase.mode2))
		}
	})

	t.Run("TestMode_IsEqual testing unequal modes", func(t *testing.T) {
		testCases := []testCase{
			// same mode, unequal tonics
			{MustMakeNewMode(ModeNameAeolian, C), MustMakeNewMode(ModeNameAeolian, D)},
			{MustMakeNewMode(ModeNameAeolian, D), MustMakeNewMode(ModeNameAeolian, E)},
			{MustMakeNewMode(ModeNameAeolian, E), MustMakeNewMode(ModeNameAeolian, F)},
			{MustMakeNewMode(ModeNameAeolian, F), MustMakeNewMode(ModeNameAeolian, G)},
			{MustMakeNewMode(ModeNameAeolian, G), MustMakeNewMode(ModeNameAeolian, A)},
			{MustMakeNewMode(ModeNameAeolian, A), MustMakeNewMode(ModeNameAeolian, B)},
			{MustMakeNewMode(ModeNameAeolian, B), MustMakeNewMode(ModeNameAeolian, C)},

			// same tonics, unequal modes
			{MustMakeNewMode(ModeNameIonian, C), MustMakeNewMode(ModeNameAeolian, C)},
			{MustMakeNewMode(ModeNameAeolian, C), MustMakeNewMode(ModeNameLydian, C)},
			{MustMakeNewMode(ModeNameLydian, C), MustMakeNewMode(ModeNameDorian, C)},
			{MustMakeNewMode(ModeNameDorian, C), MustMakeNewMode(ModeNamePhrygian, C)},
			{MustMakeNewMode(ModeNamePhrygian, C), MustMakeNewMode(ModeNameLocrian, C)},
			{MustMakeNewMode(ModeNameLocrian, C), MustMakeNewMode(ModeNameMixoLydian, C)},
			{MustMakeNewMode(ModeNameMixoLydian, C), MustMakeNewMode(ModeNameIonian, C)},
		}

		for _, testCase := range testCases {
			assert.False(t, testCase.mode1.IsEqual(testCase.mode2), "mode1: %+v, mode2: %+v", testCase.mode1, testCase.mode2)
		}
	})
}
