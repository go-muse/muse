package degree

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

const degreesInTonality = Number(17)

func TestDegree_Number(t *testing.T) {
	expected := Number(3)
	deg := Degree{number: expected}
	assert.Equal(t, expected, deg.Number())
}

func TestDegree_HalfTonesFromPrime(t *testing.T) {
	expected := halftone.HalfTones(3)
	deg := Degree{halfTonesFromPrime: expected}

	result := deg.HalfTonesFromPrime()

	if result != expected {
		t.Errorf("Degree.HalfTonesFromPrime() = %d; want %d", result, expected)
	}
}

func TestDegree_GetNext(t *testing.T) {
	type testCase struct {
		deg  *Degree
		want *Degree
	}

	constructTestCase := func(deg *Degree, next *Degree) *testCase {
		if deg != nil {
			deg.next = next
		}

		return &testCase{
			deg:  deg,
			want: next,
		}
	}

	testCases := []*testCase{
		constructTestCase(&Degree{number: 1}, &Degree{number: 2}),
		constructTestCase(&Degree{number: 1}, nil),
		constructTestCase(nil, nil),
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.deg.GetNext())
	}
}

func TestDegree_SetNext(t *testing.T) {
	d1 := &Degree{number: 1}
	d2 := &Degree{number: 2}
	d1.SetNext(d2)

	if d1.next != d2 {
		t.Errorf("d1.Next = %v; want %v", d1.next, d2)
	}
}

func TestDegree_GetPrevious(t *testing.T) {
	t.Run("GetPrevious: get previous deg", func(t *testing.T) {
		// Create a new instance of Degree with a Previous degree.
		deg := &Degree{previous: &Degree{}}

		// Call GetPrevious on the degree and store the result in a variable.
		previousDegree := deg.GetPrevious()

		// Check if the result of GetPrevious is equal to the degree's Previous degree.
		if previousDegree != deg.previous {
			t.Errorf("GetPrevious returned incorrect result: got %v, want %v", previousDegree, deg.previous)
		}
	})

	t.Run("GetPrevious: get previous degree from nil deg", func(t *testing.T) {
		var nilDegree *Degree
		assert.Nil(t, nilDegree.GetPrevious())
	})
}

func TestDegree_SetPrevious(t *testing.T) {
	// Create a degree and its previous deg
	d := &Degree{}
	prev := &Degree{}

	// Set the previous degree using the SetPrevious method
	d.SetPrevious(prev)

	// Check that the previous degree was set correctly
	if d.previous != prev {
		t.Errorf("Expected previous degree to be %v, but got %v", prev, d.previous)
	}

	// Create a new degree without a previous degree
	d2 := &Degree{}

	// Set the previous degree to nil
	d2.SetPrevious(nil)

	// Check that the previous degree was set to nil
	if d2.previous != nil {
		t.Errorf("Expected previous degree to be nil, but got %v", d2.previous)
	}
}

func TestDegree_Note(t *testing.T) {
	t.Run("get note from deg", func(t *testing.T) {
		expectedNote := note.MustNewNote(note.C)
		deg := Degree{note: expectedNote}

		if deg.Note() != expectedNote {
			t.Errorf("Expected Note %v but got %v", expectedNote, deg.Note())
		}
	})

	t.Run("get note from empty deg", func(t *testing.T) {
		var nilDegree *Degree
		assert.Nil(t, nilDegree.Note())
	})
}

func TestDegree_SetNote(t *testing.T) {
	t.Run("TestDegree_SetNote: uncycled degrees chain", func(t *testing.T) {
		// create a degree and a note
		d := &Degree{}
		n := note.C.MustNewNote()

		// set note on degree
		d.SetNote(n)

		// check if note was set correctly
		if d.note != n {
			t.Errorf("Expected note %v, but got %v", n, d.note)
		}
	})

	t.Run("TestDegree_SetNote: uncycled degrees chain", func(t *testing.T) {
		// create a degree with a note
		d := &Degree{note: note.MustNewNote(note.C)}

		// set nil note on degree
		d.SetNote(nil)

		// check if note was set to nil
		if d.note != nil {
			t.Errorf("Expected nil note, but got %v", d.note)
		}
	})

	t.Run("TestDegree_SetNote: uncycled degrees chain", func(t *testing.T) {
		// create a degree with a note
		d := &Degree{note: note.MustNewNote(note.C)}

		// create a new note and set it on degree
		n := note.MustNewNote(note.D)
		d.SetNote(n)

		// check if note was overwritten correctly
		if d.note != n {
			t.Errorf("Expected note %v, but got %v", n, d.note)
		}
	})
}

func TestDegree_ModalCharacteristics(t *testing.T) {
	mc := ModalCharacteristics{{
		name:   CharacteristicClean,
		degree: &Degree{number: 1, note: note.MustNewNote(note.C)},
	}, {
		name:   CharacteristicAug,
		degree: &Degree{number: 2, note: note.MustNewNote(note.DSHARP)},
	}}
	degree := Degree{modalCharacteristics: mc}
	result := degree.ModalCharacteristics()
	if !reflect.DeepEqual(mc, result) {
		t.Errorf("Expected modal characteristics %v but got %v", mc, result)
	}
}

func TestDegree_AbsoluteModalPosition(t *testing.T) {
	amp := &ModalPosition{name: ModalPositionNameNeutral, weight: 0}
	deg := Degree{absoluteModalPosition: amp}
	result := deg.AbsoluteModalPosition()
	if !reflect.DeepEqual(amp, result) {
		t.Errorf("Expected absolute modal position %v but got %v", amp, result)
	}
}

func TestDegree_NoteReturnsNotePointer(t *testing.T) {
	n := note.C.MustNewNote()
	deg := &Degree{note: n}

	assert.Equal(t, n, deg.Note())
}

func TestDegree_getDegreeByDegreeNum(t *testing.T) {
	testingFunc := func(t *testing.T, firstDegree *Degree) {
		t.Helper()
		for deg := range firstDegree.IterateOneRound(false) {
			for d := range firstDegree.IterateOneRound(false) {
				resultDegree := deg.GetDegreeByDegreeNum(d.Number())
				assert.Equal(t, d.Number(), resultDegree.Number(), "degree number: %d, expected: %+v, actual: %+v", deg.Number(), deg, resultDegree)
			}
		}
	}

	t.Run("TestDegree_getDegreeByDegreeNum: uncycled degrees chain", func(t *testing.T) {
		testingFunc(t, generateDegrees(3, false))
	})

	t.Run("TestDegree_getDegreeByDegreeNum: cycled degrees chain", func(t *testing.T) {
		testingFunc(t, generateDegrees(3, true))
	})

	t.Run("TestDegree_getDegreeByDegreeNum: cycled degrees chain", func(t *testing.T) {
		var nilDegree *Degree
		assert.Nil(t, nilDegree.GetDegreeByDegreeNum(5))
	})

	t.Run("TestDegree_getDegreeByDegreeNum: negative case", func(t *testing.T) {
		assert.Nil(t, generateDegrees(3, true).GetDegreeByDegreeNum(5))
	})
}

func TestDegree_GetForwardDegreeByDegreeNum(t *testing.T) {
	t.Run("GetForwardDegreeByDegreeNum: positive", func(t *testing.T) {
		firstDegreeNum := Number(1)
		firstDegree := &Degree{
			number: firstDegreeNum,
		}

		currentDegree := firstDegree
		amountOfDegrees := Number(7)
		for i := Number(2); i <= amountOfDegrees; i++ {
			newDegree := &Degree{
				number:   i,
				previous: currentDegree,
			}
			currentDegree.next = newDegree
			currentDegree = newDegree
		}
		currentDegree.next = firstDegree

		const forward = Number(103)
		expectedDegreeNum := forward - (forward/amountOfDegrees)*amountOfDegrees + 1
		result := firstDegree.GetForwardDegreeByDegreeNum(forward)
		assert.Equal(t, expectedDegreeNum, result.Number(), "expected: %d, actual: %d", expectedDegreeNum, result.Number())
	})

	t.Run("GetForwardDegreeByDegreeNum: get from nil deg", func(t *testing.T) {
		var nilDegree *Degree
		assert.Nil(t, nilDegree.GetForwardDegreeByDegreeNum(5))
	})
}

func TestDegreesIterator_GetAllDegrees(t *testing.T) {
	t.Run("GetAllDegrees: non-empty input channel", func(t *testing.T) {
		input := make(chan *Degree)
		expected := []*Degree{
			{number: 1},
			{number: 3},
			{number: 6},
			{number: 9},
		}

		go func(input chan *Degree, expected []*Degree) {
			for _, degree := range expected {
				input <- degree
			}
			close(input)
		}(input, expected)

		iter := Iterator(input)

		result := iter.GetAllDegrees()
		if !reflect.DeepEqual(result, expected) {
			assert.Equal(t, expected, result, "expected: %+v, result: %+v", expected, result)
		}
	})

	t.Run("GetAllDegrees: empty input channel", func(t *testing.T) {
		input := make(chan *Degree)
		var expected []*Degree

		go func(input chan *Degree, expected []*Degree) {
			for _, deg := range expected {
				input <- deg
			}
			close(input)
		}(input, expected)

		iter := Iterator(input)

		result := iter.GetAllDegrees()
		if !reflect.DeepEqual(result, expected) {
			assert.Equal(t, expected, result, "expected: %+v, result: %+v", expected, result)
		}
	})

	t.Run("GetAllNotes: get all degrees from nil degrees iterator", func(t *testing.T) {
		var nilDI Iterator
		assert.Nil(t, nilDI.GetAllDegrees())
	})
}

func TestDegree_GetDegrees(t *testing.T) {
	firstDegreeNum := Number(1)
	firstDegree := &Degree{
		number: firstDegreeNum,
	}

	currentDegree := firstDegree
	amountOfDegrees := Number(7)
	for i := Number(2); i <= amountOfDegrees; i++ {
		newDegree := &Degree{
			number:   i,
			previous: currentDegree,
		}
		currentDegree.next = newDegree
		currentDegree = newDegree
	}

	lastDegree := currentDegree

	t.Run("get next degree till the last degree num", func(t *testing.T) {
		currentDegree := firstDegree
		if currentDegree.Number() != firstDegreeNum {
			t.Errorf("first degree num: %d, firstDegreeNum: %d", currentDegree.Number(), firstDegreeNum)
		}

		for i := Number(2); i <= amountOfDegrees; i++ {
			currentDegree = currentDegree.GetNext()
			if i != currentDegree.Number() {
				t.Errorf("degree in cycle: %d, degree of note: %d", i, currentDegree.Number())
			}
		}
	})

	t.Run("get previous degree till the first degree num", func(t *testing.T) {
		currentDegree := lastDegree
		if currentDegree.Number() != amountOfDegrees {
			t.Errorf("last degree num: %d, amount of degrees: %d", currentDegree.Number(), amountOfDegrees)
		}

		for i := amountOfDegrees - 1; i >= firstDegreeNum; i-- {
			currentDegree = currentDegree.GetPrevious()
			if i != currentDegree.Number() {
				t.Errorf("degree in cycle: %d, degree of note: %d", i, currentDegree.Number())
			}
		}
	})

	t.Run("get next degree till the end", func(t *testing.T) {
		currentDegree := firstDegree
		if currentDegree.Number() != firstDegreeNum {
			t.Errorf("first degree num: %d, firstDegreeNum: %d", currentDegree.Number(), firstDegreeNum)
		}

		for currentDegree.next != nil {
			currentDegree = currentDegree.GetNext()
		}

		if currentDegree.Number() != amountOfDegrees {
			t.Errorf("last degree num: %d, amount of degrees: %d", currentDegree.Number(), amountOfDegrees)
		}
	})

	t.Run("get previous degree till the end", func(t *testing.T) {
		currentDegree := lastDegree
		if currentDegree.Number() != amountOfDegrees {
			t.Errorf("last degree num: %d, amount of degrees: %d", currentDegree.Number(), amountOfDegrees)
		}

		for currentDegree.previous != nil {
			currentDegree = currentDegree.GetPrevious()
		}

		if currentDegree.Number() != firstDegreeNum {
			t.Errorf("last degree num: %d, firstDegreeNum: %d", currentDegree.Number(), firstDegreeNum)
		}
	})

	t.Run("get previous degree from the empty degree", func(t *testing.T) {
		var nilDegree *Degree
		assert.Nil(t, nilDegree)
	})
}

func TestDegree_IterateOneRound(t *testing.T) {
	t.Run("IterateOneRound Iterating through cycled degrees to right", func(t *testing.T) {
		deg1 := generateDegrees(7, true)
		var i Number
		for deg := range deg1.IterateOneRound(false) {
			i++
			assert.Equal(t, deg.Number(), i)
		}
	})

	t.Run("IterateOneRound Iterating through cycled degrees to left", func(t *testing.T) {
		deg1 := generateDegrees(7, true)
		i := Number(1)
		for deg := range deg1.IterateOneRound(true) {
			assert.Equal(t, deg.Number(), i)
			if i == 1 {
				i = 7
			} else {
				i--
			}
		}
	})

	t.Run("IterateOneRound Iterating through not cycled modes to right", func(t *testing.T) {
		deg1 := generateDegrees(7, false)
		var i Number
		for deg := range deg1.IterateOneRound(false) {
			i++
			assert.Equal(t, deg.Number(), i)
		}
	})
}

func TestDegree_sortByAbsoluteModalPositions(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	n10 := rand.Intn(3) //nolint:gosec
	n20 := n10
	for n20 == n10 {
		n20 = rand.Intn(2) //nolint:gosec
	}

	getDegrees := func() (*Degree, *Degree) {
		firstDegree := &Degree{number: 1, absoluteModalPosition: NewModalPositionByWeight(Weight(n10))} //nolint:gosec
		lastDegree := firstDegree
		for i := Number(2); i <= degreesInTonality; i++ {
			deg := &Degree{number: i, absoluteModalPosition: NewModalPositionByWeight(Weight(n20))} //nolint:gosec
			lastDegree.AttachNext(deg)
			lastDegree = deg
		}

		return firstDegree, lastDegree
	}

	testingFunc := func(t *testing.T, firstSortedDegree *Degree) {
		t.Helper()
		iterator := firstSortedDegree.IterateOneRound(false)
		firstDegree := <-iterator
		var comparison bool
		for deg := range iterator {
			if deg.NextExists() {
				comparison = deg.absoluteModalPosition.Weight() <= deg.GetNext().absoluteModalPosition.Weight()
				if unsafe.Pointer(deg.GetNext()) != unsafe.Pointer(firstDegree) {
					assert.True(t, comparison, "current - degree Num: %d, w: %d, next - degree Num: %d, w: %d", deg.Number(), deg.absoluteModalPosition.Weight(), deg.GetNext().Number, deg.GetNext().absoluteModalPosition.Weight())
				} else {
					assert.False(t, comparison, "current - degree Num: %d, w: %d, next - degree Num: %d, w: %d", deg.Number(), deg.absoluteModalPosition.Weight(), deg.GetNext().Number, deg.GetNext().absoluteModalPosition.Weight())
				}
			}
		}
	}

	t.Run("test sort by AMP of not cycled degrees chain", func(t *testing.T) {
		firstDegree, _ := getDegrees()
		firstSortedDegree := firstDegree.SortByAbsoluteModalPositions(true)
		testingFunc(t, firstSortedDegree)
	})

	t.Run("test sort by AMP of cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		firstSortedDegree := firstDegree.SortByAbsoluteModalPositions(true)
		testingFunc(t, firstSortedDegree)
	})

	t.Run("test sort by AMP in case of degree without AMP", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		firstDegree.GetNext().absoluteModalPosition = nil // just one random degree without set absolute modal position
		firstSortedDegree := firstDegree.SortByAbsoluteModalPositions(true)
		assert.Nil(t, firstSortedDegree)
	})

	t.Run("test sort by AMP of cycled degrees chain Dorian mode", func(t *testing.T) {
		seventhDegree := New(7, 0, nil, nil, note.BFLAT.MustMakeNote(), nil, &ModalPosition{"", 2})
		sixthDegree := New(6, 0, nil, seventhDegree, note.A.MustMakeNote(), nil, &ModalPosition{"", 7})
		fifthDegree := New(5, 0, nil, sixthDegree, note.G.MustMakeNote(), nil, &ModalPosition{"", 6})
		fourthDegree := New(4, 0, nil, fifthDegree, note.F.MustMakeNote(), nil, &ModalPosition{"", 3})
		thirdDegree := New(3, 0, nil, fourthDegree, note.EFLAT.MustMakeNote(), nil, &ModalPosition{"", 1})
		secondDegree := New(2, 0, nil, thirdDegree, note.D.MustMakeNote(), nil, &ModalPosition{"", 5})
		firstDegree := New(1, 0, nil, secondDegree, note.C.MustMakeNote(), nil, &ModalPosition{"", 4})

		sixthDegreeSorted := New(6, 0, nil, nil, note.A.MustMakeNote(), nil, &ModalPosition{"", 7})
		fifthDegreeSorted := New(5, 0, nil, sixthDegreeSorted, note.G.MustMakeNote(), nil, &ModalPosition{"", 6})
		secondDegreeSorted := New(2, 0, nil, fifthDegreeSorted, note.D.MustMakeNote(), nil, &ModalPosition{"", 5})
		firstDegreeSorted := New(1, 0, nil, secondDegreeSorted, note.C.MustMakeNote(), nil, &ModalPosition{"", 4})
		fourthDegreeSorted := New(4, 0, nil, firstDegreeSorted, note.F.MustMakeNote(), nil, &ModalPosition{"", 3})
		seventhDegreeSorted := New(7, 0, nil, fourthDegreeSorted, note.BFLAT.MustMakeNote(), nil, &ModalPosition{"", 2})
		thirdDegreeSorted := New(3, 0, nil, seventhDegreeSorted, note.EFLAT.MustMakeNote(), nil, &ModalPosition{"", 1})

		seventhDegree.AttachNext(firstDegree)

		sortedDegree := firstDegree.SortByAbsoluteModalPositions(true)

		thirdDegreeIterator := thirdDegreeSorted.IterateOneRound(false)
		var expectedDegree *Degree
		for resultDegree := range sortedDegree.IterateOneRound(false) {
			expectedDegree = <-thirdDegreeIterator
			assert.True(t, expectedDegree.EqualByDegreeNum(resultDegree), "resulting degree number: %d, expectedDegree number: %d", resultDegree.Number(), expectedDegree.Number())
		}
	})
}

func TestDegree_String(t *testing.T) {
	testCases := []*Degree{
		nil,
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  &Degree{},
			note:                  &note.Note{},
			modalCharacteristics:  []ModalCharacteristic{},
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              nil,
			next:                  &Degree{},
			note:                  &note.Note{},
			modalCharacteristics:  []ModalCharacteristic{},
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  nil,
			note:                  &note.Note{},
			modalCharacteristics:  []ModalCharacteristic{},
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  &Degree{},
			note:                  nil,
			modalCharacteristics:  []ModalCharacteristic{},
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  &Degree{},
			note:                  &note.Note{},
			modalCharacteristics:  nil,
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  &Degree{},
			note:                  &note.Note{},
			modalCharacteristics:  []ModalCharacteristic{},
			absoluteModalPosition: nil,
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              nil,
			next:                  nil,
			note:                  nil,
			modalCharacteristics:  nil,
			absoluteModalPosition: nil,
		},
		{
			number:             1,
			halfTonesFromPrime: 1,
			previous: &Degree{
				number:                7,
				halfTonesFromPrime:    12,
				previous:              nil,
				next:                  nil,
				note:                  nil,
				modalCharacteristics:  nil,
				absoluteModalPosition: nil,
			},
			next: &Degree{
				number:                2,
				halfTonesFromPrime:    3,
				previous:              nil,
				next:                  nil,
				note:                  nil,
				modalCharacteristics:  nil,
				absoluteModalPosition: nil,
			},
			note: note.C.MustNewNote(),
			modalCharacteristics: []ModalCharacteristic{{
				name:   Characteristic2xAug,
				degree: &Degree{},
				relativeModalPosition: &ModalPosition{
					name:   ModalPositionNameNeutral,
					weight: 0,
				},
			}},
			absoluteModalPosition: &ModalPosition{
				name:   ModalPositionNameNeutral,
				weight: 0,
			},
		},
	}

	for _, degree := range testCases {
		assert.NotPanics(t, func() { _ = degree.String() }, "deg: %+v", degree) //nolint:scopelint
	}
}

func TestDegree_Equal(t *testing.T) {
	testCases := []struct {
		degree1, degree2 *Degree
		want             bool
		testNumber       uint8
	}{
		{
			degree1:    &Degree{number: 1, note: note.C.MustNewNote()},
			degree2:    &Degree{number: 1, note: note.C.MustNewNote()},
			want:       true,
			testNumber: 0,
		},
		{
			degree1:    &Degree{number: 1, note: note.C.MustNewNote()},
			degree2:    &Degree{number: 2, note: note.C.MustNewNote()},
			want:       false,
			testNumber: 2,
		},
		{
			degree1:    &Degree{number: 1, note: note.C.MustNewNote()},
			degree2:    &Degree{number: 1, note: note.D.MustNewNote()},
			want:       false,
			testNumber: 2,
		},
		{
			degree1:    &Degree{number: 1, note: note.C.MustNewNote()},
			degree2:    &Degree{number: 1, note: nil},
			want:       false,
			testNumber: 3,
		},
		{
			degree1:    &Degree{number: 1, note: nil},
			degree2:    &Degree{number: 1, note: nil},
			want:       true,
			testNumber: 4,
		},
		{
			degree1:    &Degree{number: 1, note: nil},
			degree2:    nil,
			want:       false,
			testNumber: 5,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.degree1.IsEqual(testCase.degree2), "test number: %d, expected: %+v, actual: %+v", testCase.testNumber, testCase.degree1, testCase.degree2)
	}
}

func TestDegree_EqualByDegreeNum(t *testing.T) {
	// Test the case where both degrees are nil
	if degree1, degree2 := (*Degree)(nil), (*Degree)(nil); degree1.EqualByDegreeNum(degree2) {
		t.Error("Expected nil degrees to be equal")
	}

	// Test the case where one degree is nil
	degree1 := &Degree{number: 1}
	if degree1.EqualByDegreeNum(nil) {
		t.Error("Expected non-nil degree to not be equal to nil")
	}

	// Test the case where degree numbers are equal
	degree2 := &Degree{number: 1}
	if !degree1.EqualByDegreeNum(degree2) {
		t.Error("Expected degrees with equal degree numbers to be equal")
	}

	// Test the case where degree numbers are not equal
	degree3 := &Degree{number: 2}
	if degree1.EqualByDegreeNum(degree3) {
		t.Error("Expected degrees with different degree numbers to not be equal")
	}
}

func TestDegree_Copy(t *testing.T) {
	// Arrange
	d := &Degree{
		number:                1,
		halfTonesFromPrime:    2,
		previous:              &Degree{},
		next:                  &Degree{},
		note:                  note.MustNewNoteWithOctave(note.C, octave.Number1),
		modalCharacteristics:  ModalCharacteristics{},
		absoluteModalPosition: &ModalPosition{name: ModalPositionNameHigh, weight: -5},
	}

	// Act
	copiedDegree := d.Copy()

	// Assert
	if copiedDegree.Number() != d.Number() {
		t.Errorf("Number not copied properly. Expected: %d, Actual: %d", d.Number(), copiedDegree.Number())
	}
	if copiedDegree.halfTonesFromPrime != d.halfTonesFromPrime {
		t.Errorf("HalfTonesFromPrime not copied properly. Expected: %d, Actual: %d", d.halfTonesFromPrime, copiedDegree.halfTonesFromPrime)
	}
	if copiedDegree.previous != d.previous {
		t.Errorf("Previous not copied properly")
	}
	if copiedDegree.next != d.next {
		t.Errorf("Next not copied properly")
	}
	if copiedDegree.note.Name() != d.note.Name() {
		t.Errorf("Note not copied properly. Expected: %+v, Actual: %+v", d.note, copiedDegree.note)
	}
	// TODO: ModalCharacteristics comparation
	// if copy.ModalCharacteristics != d.ModalCharacteristics {
	// 	t.Errorf("ModalCharacteristics not copied properly")
	// }
	if copiedDegree.absoluteModalPosition.Weight() != d.absoluteModalPosition.Weight() ||
		copiedDegree.absoluteModalPosition.name != d.absoluteModalPosition.name {
		t.Errorf("AbsoluteModalPosition not copied properly. Expected: %+v, Actual: %+v", d.absoluteModalPosition, copiedDegree.absoluteModalPosition)
	}
}

func TestDegree_CopyCut(t *testing.T) {
	// Arrange
	d := &Degree{
		number:                1,
		halfTonesFromPrime:    2,
		previous:              &Degree{},
		next:                  &Degree{},
		note:                  note.MustNewNoteWithOctave(note.C, octave.Number0),
		modalCharacteristics:  ModalCharacteristics{},
		absoluteModalPosition: &ModalPosition{name: ModalPositionNameHigh, weight: -5},
	}

	// Act
	copiedDegree := d.CopyCut()

	// Assert
	if copiedDegree.Number() != d.Number() {
		t.Errorf("Number not copied properly. Expected: %d, Actual: %d", d.Number(), copiedDegree.Number())
	}
	if copiedDegree.halfTonesFromPrime != d.halfTonesFromPrime {
		t.Errorf("HalfTonesFromPrime not copied properly. Expected: %d, Actual: %d", d.halfTonesFromPrime, copiedDegree.halfTonesFromPrime)
	}
	if copiedDegree.previous != nil {
		t.Errorf("Previous not set to nil")
	}
	if copiedDegree.next != nil {
		t.Errorf("Next not set to nil")
	}
	if copiedDegree.note.Name() != d.note.Name() {
		t.Errorf("Note not copied properly. Expected: %+v, Actual: %+v", d.note, copiedDegree.note)
	}
	// TODO: ModalCharacteristics comparation
	// if copy.ModalCharacteristics != d.ModalCharacteristics {
	// 	t.Errorf("ModalCharacteristics not copied properly")
	// }
	if copiedDegree.absoluteModalPosition.Weight() != d.absoluteModalPosition.Weight() ||
		copiedDegree.absoluteModalPosition.name != d.absoluteModalPosition.name {
		t.Errorf("AbsoluteModalPosition not copied properly. Expected: %+v, Actual: %+v", d.absoluteModalPosition, copiedDegree.absoluteModalPosition)
	}
}

func TestDegree_InsertBetween(t *testing.T) {
	degree1 := &Degree{number: 1}
	degree2 := &Degree{number: 2}
	degree3 := &Degree{number: 3}

	degree2.InsertBetween(degree1, degree3)

	if unsafe.Pointer(degree1.next) != unsafe.Pointer(degree2) {
		t.Errorf("degree1 is not correctly attached to degree2. Next Degree is: %d", degree1.GetNext().Number())
	}
	if degree3.previous != degree2 {
		t.Errorf("degree3 is not correctly attached to degree2. Previous Degree is: %d", degree3.GetPrevious().Number())
	}
	if unsafe.Pointer(degree2.next) != unsafe.Pointer(degree3) {
		t.Errorf("degree2 is not correctly attached to degree3. Next Degree is: %d", degree2.GetNext().Number())
	}
	if degree2.previous != degree1 {
		t.Errorf("degree2 is not correctly attached to degree1. Previous Degree is: %d", degree2.GetPrevious().Number())
	}
}

func TestDegree_AttachNext(t *testing.T) {
	// Set up initial Degrees
	degree1 := &Degree{number: 1}
	degree2 := &Degree{number: 2}

	// Attach degree2 as next degree after degree1
	degree1.AttachNext(degree2)

	// Assertions to ensure references are correctly set
	if degree1.next != degree2 {
		t.Errorf("Expected degree2 to be the next degree after degree1, but got %v", degree1.next)
	}
	if degree2.previous != degree1 {
		t.Errorf("Expected degree1 to be the previous degree before degree2, but got %v", degree2.previous)
	}

	// Ensure mutual references are not set
	if degree1.previous != nil {
		t.Errorf("Expected degree1 to have no previous deg, but got %v", degree1.previous)
	}
	if degree2.next != nil {
		t.Errorf("Expected degree2 to have no next deg, but got %v", degree2.next)
	}

	// Attach degree1 as next degree after degree2
	degree2.AttachNext(degree1)

	// Assertions to ensure references are correctly set
	if degree2.next != degree1 {
		t.Errorf("Expected degree1 to be the next degree after degree2, but got %v", degree2.next)
	}
	if degree1.previous != degree2 {
		t.Errorf("Expected degree2 to be the previous degree before degree1, but got %v", degree1.previous)
	}

	// Ensure mutual references are set
	if degree2.previous != degree1 {
		t.Errorf("Expected degree1 to be the previous degree before degree2, but got %v", degree2.previous)
	}
	if degree1.next != degree2 {
		t.Errorf("Expected degree2 to be the next degree after degree1, but got %v", degree1.next)
	}

	// Attach degree2 as next degree after degree2 (invalid input)
	degree2.AttachNext(degree2)

	// Assertions to ensure references are not set
	if degree2.next != degree2 {
		t.Errorf("Expected degree2 to be the next degree after degree2, but got %v", degree1.next)
	}
	if degree2.previous != degree2 {
		t.Errorf("Expected degree2 to be the previous degree before degree2, but got %v", degree1.next)
	}
	if degree1.next != degree2 {
		t.Errorf("Expected degree2 to be the next degree after degree1, but got %v", degree1.next)
	}
	if degree1.previous != degree2 {
		t.Errorf("Expected degree2 to be the previous degree before degree1, but got %v", degree1.next)
	}
}

func TestDegree_InsertNext(t *testing.T) {
	// Create three-degree objects
	degree1 := &Degree{}
	degree2 := &Degree{}
	degree3 := &Degree{}

	// Call the InsertNext function on degree2 and passing degree1 as the new Degree object to be inserted
	degree2.InsertNext(degree1)

	// Check that degree1 is attached as next to degree2
	if degree2.next != degree1 {
		t.Errorf("degree1 is not correctly attached as next to degree2")
	}
	if degree1.previous != degree2 {
		t.Errorf("degree2 is not correctly attached as previous to degree1")
	}

	// Call the InsertNext function on degree2 (which already has a next) and passing degree3 as the new Degree object to be inserted
	degree2.InsertNext(degree3)

	// Check that degree3 is attached as next to degree2 and degree2 is attached as previous to degree3 (degree1 is in between)
	if degree2.next != degree3 {
		t.Errorf("degree3 is not correctly attached as next to degree2")
	}
	if degree3.previous != degree2 {
		t.Errorf("degree2 is not correctly attached as previous to degree3")
	}
}

func TestDegree_AttachPrevious(t *testing.T) {
	// Create the first deg
	rootDegree := &Degree{number: 1}

	// Create the second deg
	secondDegree := &Degree{number: 2}

	// Attach secondDegree as previous to rootDegree
	rootDegree.AttachPrevious(secondDegree)

	// Test that AttachPrevious sets the previous degree correctly
	if rootDegree.GetPrevious() != secondDegree {
		t.Errorf("AttachPrevious doesn't set previous degree correctly, expected %v but got %v",
			secondDegree, rootDegree.GetPrevious())
	}

	// Test that AttachPrevious sets the next degree correctly
	if secondDegree.GetNext() != rootDegree {
		t.Errorf("AttachPrevious doesn't set next degree correctly, expected %v but got %v",
			rootDegree, secondDegree.GetNext())
	}

	// Create thirdDegree
	thirdDegree := &Degree{number: 3}

	// Attach thirdDegree as previous to secondDegree
	secondDegree.AttachPrevious(thirdDegree)

	// Test that the previous degree of secondDegree is now thirdDegree
	if secondDegree.GetPrevious() != thirdDegree {
		t.Errorf("AttachPrevious doesn't update previous degree correctly, expected %v but got %v",
			thirdDegree, secondDegree.GetPrevious())
	}

	// Test that AttachPrevious sets the next degree correctly
	if thirdDegree.GetNext() != secondDegree {
		t.Errorf("AttachPrevious doesn't set next degree correctly, expected %v but got %v",
			secondDegree, thirdDegree.GetNext())
	}
}

func TestDegree_InsertPrevious(t *testing.T) {
	// create degrees
	degree1 := &Degree{number: 1}
	degree2 := &Degree{number: 2}
	degree3 := &Degree{number: 3}

	// inserting previous degree to a nil degree should result in panicking
	var nilDegree *Degree
	assert.Panics(t, func() { nilDegree.InsertPrevious(degree2) })

	// insert previous degree and test for mutual reference
	degree1.InsertPrevious(degree2)
	assert.Equal(t, degree1.previous, degree2)
	assert.Equal(t, degree2.next, degree1)

	// insert another degree as the previous degree and test for mutual reference
	degree1.InsertPrevious(degree3)
	assert.Equal(t, degree1.previous, degree3)
	assert.Equal(t, degree3.next, degree1)
	assert.Equal(t, degree3.previous, degree2)
	assert.Equal(t, degree2.next, degree3)
}

func TestDegree_NextExists(t *testing.T) {
	// Create a degree with no next degree attached
	degree := &Degree{number: 1}

	// Test that NextExists returns false
	if degree.NextExists() {
		t.Errorf("NextExists should return false when next degree doesn't exist")
	}

	// Create another degree and attach it as next to the previous deg
	nextDegree := &Degree{number: 2}
	degree.AttachNext(nextDegree)

	// Test that NextExists returns true
	if !degree.NextExists() {
		t.Errorf("NextExists should return true when next degree exists")
	}
}

func TestDegree_PreviousExists(t *testing.T) {
	// Create a degree with no previous degree attached
	degree := &Degree{number: 1}

	// Test that PreviousExists returns false
	if degree.PreviousExists() {
		t.Errorf("PreviousExists should return false when previous degree doesn't exist")
	}

	// Create another degree and attach it as previous to the previous deg
	previousDegree := &Degree{number: 2}
	degree.AttachPrevious(previousDegree)

	// Test that PreviousExists returns true
	if !degree.PreviousExists() {
		t.Errorf("PreviousExists should return true when previous degree exists")
	}
}

func TestDegree_GetLast(t *testing.T) {
	getDegrees := func() (*Degree, *Degree) {
		firstDegree := &Degree{number: 1}
		lastDegree := firstDegree
		for i := Number(2); i <= degreesInTonality; i++ {
			deg := &Degree{number: i}
			lastDegree.AttachNext(deg)
			lastDegree = deg
		}

		return firstDegree, lastDegree
	}

	t.Run("get ending degrees of the not cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		var deg, result *Degree
		for i := Number(0); i < degreesInTonality; i++ {
			deg = firstDegree.GetForwardDegreeByDegreeNum(i)
			result = deg.GetLast(false)
			assert.True(t, result.EqualByDegreeNum(lastDegree), "expected lastDegree: %+v, actual result: %+v", lastDegree, result)
			result = deg.GetLast(true)
			assert.True(t, result.EqualByDegreeNum(firstDegree), "expected firstDegree: %+v, actual result: %+v", firstDegree, result)
		}
	})

	t.Run("get ending degrees of the cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		var deg, result *Degree
		for i := Number(0); i < degreesInTonality; i++ {
			deg = firstDegree.GetForwardDegreeByDegreeNum(i)
			result = deg.GetLast(false)
			assert.True(t, result.EqualByDegreeNum(deg.GetPrevious()), "expected lastDegree: %+v, actual result: %+v", lastDegree, result)
			result = deg.GetLast(true)
			assert.True(t, result.EqualByDegreeNum(deg.GetNext()), "expected firstDegree: %+v, actual result: %+v", firstDegree, result)
		}
	})

	t.Run("get ending degrees from nil deg", func(t *testing.T) {
		var deg *Degree
		result := deg.GetLast(true)
		assert.Nil(t, result, "expected nil, but got: %+v", result)
	})
}

func TestDegree_AttachToTheEnd(t *testing.T) {
	t.Run("TestAttachToTheEnd test case when degrees chain is not cycled", func(t *testing.T) {
		d1 := &Degree{number: 1}
		d2 := &Degree{number: 2}
		d3 := &Degree{number: 3}
		d4 := &Degree{number: 4}

		d1.AttachNext(d2)
		d2.AttachNext(d3)
		d3.AttachNext(d4)

		d5 := &Degree{number: 5}
		d1.AttachToTheEnd(d5, false)

		expected := []*Degree{d1, d2, d3, d4, d5}

		var i int
		for degree := range d1.IterateOneRound(false) {
			assert.Equal(t, expected[i], degree)
			i++
		}

		d0 := &Degree{number: 0}
		d5.AttachToTheEnd(d0, true)

		expected = []*Degree{d0, d1, d2, d3, d4, d5}

		var j int
		for degree := range d0.IterateOneRound(false) {
			assert.Equal(t, expected[j], degree)
			j++
		}
	})

	t.Run("TestAttachToTheEnd test case when degrees chain is cycled", func(t *testing.T) {
		d1 := &Degree{number: 1}
		d2 := &Degree{number: 2}
		d3 := &Degree{number: 3}
		d4 := &Degree{number: 4}
		d1.AttachPrevious(d4)
		d1.AttachNext(d2)
		d2.AttachNext(d3)
		d3.AttachNext(d4)

		d5 := &Degree{number: 5}
		d1.AttachToTheEnd(d5, false)

		expected := []*Degree{d1, d2, d3, d4, d5}

		var i int
		for degree := range d1.IterateOneRound(false) {
			assert.Equal(t, expected[i], degree)
			i++
		}

		d6 := &Degree{number: 6}
		d1.GetPrevious().AttachToTheEnd(d6, true)

		expected = append(expected, d6)

		var j int
		for degree := range d1.IterateOneRound(false) {
			assert.Equal(t, expected[j], degree)
			j++
		}
	})
}

func TestDegree_ReverseSequence(t *testing.T) {
	t.Run("TestReverseSequence with cycled sequence", func(t *testing.T) {
		d := generateDegrees(7, true)
		res := d.ReverseSequence()
		exp := d.GetPrevious().IterateOneRound(true)
		for degree := range res.IterateOneRound(false) {
			expectedDegree := <-exp
			assert.Equal(t, degree.Number(), expectedDegree.Number(), "expected: %d, actual: %d", degree.Number(), expectedDegree.Number())
		}
	})

	t.Run("TestReverseSequence with not cycled sequence", func(t *testing.T) {
		d := generateDegrees(7, false)
		res := d.ReverseSequence()
		exp := d.GetLast(false).IterateOneRound(true)
		for degree := range res.IterateOneRound(false) {
			expectedDegree := <-exp
			assert.Equal(t, degree.Number(), expectedDegree.Number(), "expected: %d, actual: %d", degree.Number(), expectedDegree.Number())
		}
	})
}

func generateDegrees(degreesAmount uint8, isCycled bool) *Degree {
	if degreesAmount < 1 {
		return nil
	}
	degreeNum := Number(1)
	firstDegree := newDegreeWithNum(degreeNum)
	currentDegree := firstDegree
	for degreeNum++; degreeNum <= Number(degreesAmount); degreeNum++ {
		newDegree := newDegreeWithNum(degreeNum)
		currentDegree.AttachNext(newDegree)
		currentDegree = newDegree
	}

	if isCycled {
		firstDegree.AttachPrevious(currentDegree)
	}

	return firstDegree
}

func newDegreeWithNum(
	number Number,
) *Degree {
	return New(
		number,
		0,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}
