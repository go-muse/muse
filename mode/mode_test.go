package mode

import (
	"crypto/rand"
	"math/big"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

func TestInsertNote(t *testing.T) {
	mode := Mode{}
	note1 := note.C.MustMakeNote()

	mode.InsertNote(note1, 0)
	assert.True(t, mode.GetFirstDegree().Note().IsEqualByName(note1))
	assert.Nil(t, mode.GetFirstDegree().GetPrevious())
	assert.Nil(t, mode.GetFirstDegree().GetNext())

	note2 := note.D.MustMakeNote()
	mode.InsertNote(note2, 2)
	assert.True(t, mode.GetFirstDegree().Note().IsEqualByName(note1))
	assert.Nil(t, mode.GetFirstDegree().GetPrevious())
	assert.NotNil(t, mode.GetFirstDegree().GetNext())

	assert.True(t, mode.GetFirstDegree().GetNext().Note().IsEqualByName(note2))
	assert.NotNil(t, mode.GetFirstDegree().GetNext().GetPrevious())
	assert.True(t, mode.GetFirstDegree().GetNext().GetPrevious().Note().IsEqualByName(note1))
	assert.Nil(t, mode.GetFirstDegree().GetNext().GetNext())
}

func TestMakeNewMode(t *testing.T) {
	t.Run("make natural minor mode", func(t *testing.T) {
		tonalCenter := note.C
		mode, err := MakeNewMode(NameNaturalMinor, tonalCenter)
		require.NoError(t, err)
		assert.NotNil(t, mode)

		firstDegree := mode.GetFirstDegree()

		firstNote, err := note.NewNoteWithOctave(tonalCenter, octave.NumberDefault)
		require.NoError(t, err)
		assert.True(t, firstDegree.Note().IsEqualByName(firstNote))

		naturalMinorFromC := note.Notes{
			note.MustNewNote(note.D),
			note.MustNewNote(note.EFLAT),
			note.MustNewNote(note.F),
			note.MustNewNote(note.G),
			note.MustNewNote(note.AFLAT),
			note.MustNewNote(note.BFLAT),
		}

		for _, note := range naturalMinorFromC {
			firstDegree = firstDegree.GetNext()
			assert.NotNil(t, firstDegree)
			assert.NotNil(t, firstDegree.Note())
			assert.True(t, firstDegree.Note().IsEqualByName(note), "expected: %s, actual: %s", note.Name(), firstDegree.Note().Name())
		}
	})

	t.Run("make natural major mode", func(t *testing.T) {
		tonalCenter := note.C
		mode, err := MakeNewMode(NameNaturalMajor, tonalCenter)
		require.NoError(t, err)
		assert.NotNil(t, mode)

		firstDegree := mode.GetFirstDegree()
		assert.NotNil(t, mode)

		firstNote, err := note.NewNoteWithOctave(tonalCenter, octave.NumberDefault)
		require.NoError(t, err)
		assert.True(t, firstDegree.Note().IsEqualByName(firstNote))

		naturalMajorFromC := note.Notes{
			note.MustNewNote(note.D),
			note.MustNewNote(note.E),
			note.MustNewNote(note.F),
			note.MustNewNote(note.G),
			note.MustNewNote(note.A),
			note.MustNewNote(note.B),
		}

		for _, note := range naturalMajorFromC {
			firstDegree = firstDegree.GetNext()
			assert.NotNil(t, firstDegree)
			assert.NotNil(t, firstDegree.Note())
			assert.True(t, firstDegree.Note().IsEqualByName(note))
		}
	})

	t.Run("make natural major mode B", func(t *testing.T) {
		tonalCenter := note.B
		mode, err := MakeNewMode(NameNaturalMajor, tonalCenter)
		require.NoError(t, err)
		assert.NotNil(t, mode)

		firstDegree := mode.GetFirstDegree()
		assert.NotNil(t, mode)

		firstNote, err := note.NewNoteWithOctave(tonalCenter, octave.NumberDefault)
		require.NoError(t, err)
		assert.True(t, firstDegree.Note().IsEqualByName(firstNote))

		naturalMajorFromB := note.Notes{
			note.MustNewNote(note.CSHARP),
			note.MustNewNote(note.DSHARP),
			note.MustNewNote(note.E),
			note.MustNewNote(note.FSHARP),
			note.MustNewNote(note.GSHARP),
			note.MustNewNote(note.ASHARP),
		}

		for _, note := range naturalMajorFromB {
			firstDegree = firstDegree.GetNext()
			assert.NotNil(t, firstDegree)
			assert.NotNil(t, firstDegree.Note())
			assert.True(t, firstDegree.Note().IsEqualByName(note), "expect note %s, actual: %s", firstDegree.Note().Name(), note.Name())
		}
	})
}

func TestMustMakeNewMode(t *testing.T) {
	// Test case with valid inputs
	var mode *Mode
	store := InitTemplatesStore()
	for modeName := range store {
		assert.NotPanics(t, func() { mode = MustMakeNewMode(modeName, note.A) }) //nolint:scopelint
		if mode == nil {
			t.Errorf("Expected mode to be created successfully, but got nil")
		}
	}

	// Negative test case with invalid mode name
	withPanic := func() {
		MustMakeNewMode(Name(""), note.Name("C"))
	}
	assert.Panics(t, withPanic)

	// Negative test case with invalid note name
	withPanic = func() {
		MustMakeNewMode(Name("testMode"), note.Name(""))
	}
	assert.Panics(t, withPanic)
}

func TestOpenCircleOfDegreesMethod(t *testing.T) {
	// Create a Mode with a circle of degrees
	degree1 := newDegreeWithNum(1)
	degree2 := newDegreeWithNum(2)
	degree3 := newDegreeWithNum(3)
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
	degree1 := newDegreeWithNum(1)
	degree2 := newDegreeWithNum(2)
	degree3 := newDegreeWithNum(3)

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
	require.NoError(t, err)
	n20, err := rand.Int(rand.Reader, big.NewInt(10))
	require.NoError(t, err)

	getDegrees := func() (*degree.Degree, *degree.Degree) {
		firstDegree := newDegreeWithNum(1)
		firstDegree.SetAbsoluteModalPosition(degree.NewModalPositionByWeight(degree.Weight(n10.Int64() - 5))) //nolint:gosec
		lastDegree := firstDegree
		for i := degree.Number(2); i <= DegreesInTonality; i++ {
			nextDegree := newDegreeWithNum(i)
			nextDegree.SetAbsoluteModalPosition(degree.NewModalPositionByWeight(degree.Weight(n20.Int64() - 10))) //nolint:gosec
			lastDegree.AttachNext(nextDegree)
			lastDegree = nextDegree
		}

		return firstDegree, lastDegree
	}

	testingFunc := func(t *testing.T, firstSortedDegree *degree.Degree) {
		t.Helper()
		iterator := firstSortedDegree.IterateOneRound(false)
		firstDegree := <-iterator
		var comparison bool
		for degree := range iterator {
			if degree.NextExists() {
				comparison = degree.AbsoluteModalPosition().Weight() <= degree.GetNext().AbsoluteModalPosition().Weight()
				if unsafe.Pointer(degree.GetNext()) != unsafe.Pointer(firstDegree) {
					assert.True(t, comparison, "current - degree Num: %d, w: %d, next - degree Num: %d, w: %d", degree.Number(), degree.AbsoluteModalPosition().Weight(), degree.GetNext().Number, degree.GetNext().AbsoluteModalPosition().Weight())
				} else {
					assert.False(t, comparison, "current - degree Num: %d, w: %d, next - degree Num: %d, w: %d", degree.Number(), degree.AbsoluteModalPosition().Weight(), degree.GetNext().Number, degree.GetNext().AbsoluteModalPosition().Weight())
				}
			}
		}
	}

	t.Run("test sort by AMP of cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		mode, err := MakeNewCustomModeWithDegree("custom mode", firstDegree)
		require.NoError(t, err)
		mode.SortByAbsoluteModalPositions(false)
		testingFunc(t, mode.GetFirstDegree())
	})

	t.Run("test sort by AMP of uncycled degrees chain", func(t *testing.T) {
		firstDegree, _ := getDegrees()
		mode, err := MakeNewCustomModeWithDegree("custom mode", firstDegree)
		require.NoError(t, err)
		mode.SortByAbsoluteModalPositions(false)
		testingFunc(t, mode.GetFirstDegree())
	})

	t.Run("test sort by AMP in case of degree without AMP", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		firstDegree.GetNext().SetAbsoluteModalPosition(nil) // just one random degree without set absolute modal position
		mode, err := MakeNewCustomModeWithDegree("custom mode", firstDegree)
		require.NoError(t, err)
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
	mode, err := MakeNewCustomMode(Template{2, 2, 1, 2, 2, 2, 1}, "C", NameNaturalMajor)
	require.NoError(t, err)

	testCases := []struct {
		note *note.Note
		want bool
	}{
		{note.C.MustMakeNote(), true},
		{note.D.MustMakeNote(), true},
		{note.E.MustMakeNote(), true},
		{note.F.MustMakeNote(), true},
		{note.G.MustMakeNote(), true},
		{note.A.MustMakeNote(), true},
		{note.B.MustMakeNote(), true},

		{note.CFLAT.MustMakeNote(), false},
		{note.CFLAT2.MustMakeNote(), false},
		{note.CSHARP.MustMakeNote(), false},
		{note.CSHARP2.MustMakeNote(), false},
		{note.DFLAT.MustMakeNote(), false},
		{note.DFLAT2.MustMakeNote(), false},
		{note.DSHARP.MustMakeNote(), false},
		{note.DSHARP2.MustMakeNote(), false},
		{note.EFLAT.MustMakeNote(), false},
		{note.EFLAT2.MustMakeNote(), false},
		{note.ESHARP.MustMakeNote(), false},
		{note.ESHARP2.MustMakeNote(), false},
		{note.FFLAT.MustMakeNote(), false},
		{note.FFLAT2.MustMakeNote(), false},
		{note.FSHARP.MustMakeNote(), false},
		{note.FSHARP2.MustMakeNote(), false},
		{note.GFLAT.MustMakeNote(), false},
		{note.GFLAT2.MustMakeNote(), false},
		{note.GSHARP.MustMakeNote(), false},
		{note.GSHARP2.MustMakeNote(), false},
		{note.AFLAT.MustMakeNote(), false},
		{note.AFLAT2.MustMakeNote(), false},
		{note.ASHARP.MustMakeNote(), false},
		{note.ASHARP2.MustMakeNote(), false},
		{note.BFLAT.MustMakeNote(), false},
		{note.BFLAT2.MustMakeNote(), false},
		{note.BSHARP.MustMakeNote(), false},
		{note.BSHARP2.MustMakeNote(), false},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, mode.Contains(testCase.note), "expected note in mode: %f, actual: %f", testCase.want, mode.Contains(testCase.note))
	}
}

//nolint:dupl
func TestMode_IsEqual(t *testing.T) {
	type testCase struct {
		mode1, mode2 *Mode
	}

	t.Run("TestMode_IsEqual testing equal modes", func(t *testing.T) {
		testCases := []testCase{
			// same mode, same tonics
			{MustMakeNewMode(NameAeolian, note.C), MustMakeNewMode(NameAeolian, note.C)},
			{MustMakeNewMode(NameAeolian, note.D), MustMakeNewMode(NameAeolian, note.D)},
			{MustMakeNewMode(NameAeolian, note.E), MustMakeNewMode(NameAeolian, note.E)},
			{MustMakeNewMode(NameAeolian, note.F), MustMakeNewMode(NameAeolian, note.F)},
			{MustMakeNewMode(NameAeolian, note.G), MustMakeNewMode(NameAeolian, note.G)},
			{MustMakeNewMode(NameAeolian, note.A), MustMakeNewMode(NameAeolian, note.A)},
			{MustMakeNewMode(NameAeolian, note.B), MustMakeNewMode(NameAeolian, note.B)},

			// same tonics, different modes
			{MustMakeNewMode(NameIonian, note.C), MustMakeNewMode(NameIonian, note.C)},
			{MustMakeNewMode(NameAeolian, note.C), MustMakeNewMode(NameAeolian, note.C)},
			{MustMakeNewMode(NameLydian, note.C), MustMakeNewMode(NameLydian, note.C)},
			{MustMakeNewMode(NameDorian, note.C), MustMakeNewMode(NameDorian, note.C)},
			{MustMakeNewMode(NamePhrygian, note.C), MustMakeNewMode(NamePhrygian, note.C)},
			{MustMakeNewMode(NameLocrian, note.C), MustMakeNewMode(NameLocrian, note.C)},
			{MustMakeNewMode(NameMixoLydian, note.C), MustMakeNewMode(NameMixoLydian, note.C)},
		}

		for _, testCase := range testCases {
			assert.True(t, testCase.mode1.IsEqual(testCase.mode2))
		}
	})

	t.Run("TestMode_IsEqual testing unequal modes", func(t *testing.T) {
		testCases := []testCase{
			// same mode, different tonics
			{MustMakeNewMode(NameAeolian, note.C), MustMakeNewMode(NameAeolian, note.D)},
			{MustMakeNewMode(NameAeolian, note.D), MustMakeNewMode(NameAeolian, note.E)},
			{MustMakeNewMode(NameAeolian, note.E), MustMakeNewMode(NameAeolian, note.F)},
			{MustMakeNewMode(NameAeolian, note.F), MustMakeNewMode(NameAeolian, note.G)},
			{MustMakeNewMode(NameAeolian, note.G), MustMakeNewMode(NameAeolian, note.A)},
			{MustMakeNewMode(NameAeolian, note.A), MustMakeNewMode(NameAeolian, note.B)},
			{MustMakeNewMode(NameAeolian, note.B), MustMakeNewMode(NameAeolian, note.C)},

			// same tonics, different modes
			{MustMakeNewMode(NameIonian, note.C), MustMakeNewMode(NameAeolian, note.C)},
			{MustMakeNewMode(NameAeolian, note.C), MustMakeNewMode(NameLydian, note.C)},
			{MustMakeNewMode(NameLydian, note.C), MustMakeNewMode(NameDorian, note.C)},
			{MustMakeNewMode(NameDorian, note.C), MustMakeNewMode(NamePhrygian, note.C)},
			{MustMakeNewMode(NamePhrygian, note.C), MustMakeNewMode(NameLocrian, note.C)},
			{MustMakeNewMode(NameLocrian, note.C), MustMakeNewMode(NameMixoLydian, note.C)},
			{MustMakeNewMode(NameMixoLydian, note.C), MustMakeNewMode(NameIonian, note.C)},
		}

		for _, testCase := range testCases {
			assert.False(t, testCase.mode1.IsEqual(testCase.mode2), "mode1: %+v, mode2: %+v", testCase.mode1, testCase.mode2)
		}
	})
}

func newDegreeWithNum(degreeNum degree.Number) *degree.Degree {
	return degree.New(
		degreeNum,
		0,
		nil, nil,
		nil,
		nil,
		nil,
	)
}
