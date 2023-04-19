package muse

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestDegree_Number(t *testing.T) {
	expected := DegreeNum(3)
	d := Degree{number: expected}
	assert.Equal(t, expected, d.Number())
}

func TestDegree_HalfTonesFromPrime(t *testing.T) {
	expected := HalfTones(3)
	degree := Degree{halfTonesFromPrime: expected}

	result := degree.HalfTonesFromPrime()

	if result != expected {
		t.Errorf("Degree.HalfTonesFromPrime() = %d; want %d", result, expected)
	}
}

func TestDegree_GetNext(t *testing.T) {
	testCases := []struct {
		degree *Degree
	}{
		{&Degree{number: 1, next: &Degree{number: 2}}},
		{&Degree{number: 1, next: nil}},
		{&Degree{}},
		{nil},
	}

	for _, testCase := range testCases {
		if testCase.degree == nil {
			assert.Panics(t, func() { _ = testCase.degree.GetNext() }) //nolint:scopelint

			continue
		}
		if testCase.degree.next == nil {
			assert.Nil(t, testCase.degree.GetNext())
		}
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
	// Create a new instance of Degree with a Previous degree.
	degree := &Degree{previous: &Degree{}}

	// Call GetPrevious on the degree and store the result in a variable.
	previousDegree := degree.GetPrevious()

	// Check if the result of GetPrevious is equal to the degree's Previous degree.
	if previousDegree != degree.previous {
		t.Errorf("GetPrevious returned incorrect result: got %v, want %v", previousDegree, degree.previous)
	}
}

func TestDegree_SetPrevious(t *testing.T) {
	// Create a degree and its previous degree
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
	expectedNote := newNote(C)
	degree := Degree{note: expectedNote}

	if degree.Note() != expectedNote {
		t.Errorf("Expected Note %v but got %v", expectedNote, degree.Note())
	}
}

func TestDegree_SetNote(t *testing.T) {
	t.Run("TestGetDegreeByDegreeNum: uncycled degrees chain", func(t *testing.T) {
		// create a degree and a note
		d := &Degree{}
		n := newNote(C)

		// set note on degree
		d.SetNote(n)

		// check if note was set correctly
		if d.note != n {
			t.Errorf("Expected note %v, but got %v", n, d.note)
		}
	})

	t.Run("TestGetDegreeByDegreeNum: uncycled degrees chain", func(t *testing.T) {
		// create a degree with a note
		d := &Degree{note: newNote(C)}

		// set nil note on degree
		d.SetNote(nil)

		// check if note was set to nil
		if d.note != nil {
			t.Errorf("Expected nil note, but got %v", d.note)
		}
	})

	t.Run("TestGetDegreeByDegreeNum: uncycled degrees chain", func(t *testing.T) {
		// create a degree with a note
		d := &Degree{note: newNote(C)}

		// create a new note and set it on degree
		n := newNote(D)
		d.SetNote(n)

		// check if note was overwritten correctly
		if d.note != n {
			t.Errorf("Expected note %v, but got %v", n, d.note)
		}
	})
}

func TestDegree_ModalCharacteristics(t *testing.T) {
	mc := ModalCharacteristics{{
		name:   DegreeCharacteristicClean,
		degree: &Degree{number: 1, note: newNote(C)},
	}, {
		name:   DegreeCharacteristicAug,
		degree: &Degree{number: 2, note: newNote(DSHARP)},
	}}
	degree := Degree{modalCharacteristics: mc}
	result := degree.ModalCharacteristics()
	if !reflect.DeepEqual(mc, result) {
		t.Errorf("Expected modal characteristics %v but got %v", mc, result)
	}
}

func TestDegree_AbsoluteModalPosition(t *testing.T) {
	amp := &ModalPosition{name: ModalPositionNameNeutral, weight: 0}
	degree := Degree{absoluteModalPosition: amp}
	result := degree.AbsoluteModalPosition()
	if !reflect.DeepEqual(amp, result) {
		t.Errorf("Expected absolute modal position %v but got %v", amp, result)
	}
}

func TestDegree_NoteReturnsNotePointer(t *testing.T) {
	note := &Note{name: C}
	degree := &Degree{note: note}

	assert.Equal(t, note, degree.Note())
}

func TestGetDegreeByDegreeNum(t *testing.T) {
	testingFunc := func(t *testing.T, firstDegree *Degree) {
		t.Helper()
		for degree := range firstDegree.IterateOneRound(false) {
			for d := range firstDegree.IterateOneRound(false) {
				resultDegree := degree.GetDegreeByDegreeNum(d.Number())
				assert.Equal(t, d.Number(), resultDegree.Number(), "degree number: %d, expected: %+v, actual: %+v", degree.Number(), degree, resultDegree)
			}
		}
	}

	t.Run("TestGetDegreeByDegreeNum: uncycled degrees chain", func(t *testing.T) {
		testingFunc(t, generateDegrees(3, false))
	})

	t.Run("TestGetDegreeByDegreeNum: cycled degrees chain", func(t *testing.T) {
		testingFunc(t, generateDegrees(3, true))
	})
}

func TestDegree_GetForwardDegreeByDegreeNum(t *testing.T) {
	firstDegreeNum := DegreeNum(1)
	firstDegree := &Degree{
		number: firstDegreeNum,
	}

	currentDegree := firstDegree
	amountOfDegrees := DegreeNum(7)
	for i := DegreeNum(2); i <= amountOfDegrees; i++ {
		newDegree := &Degree{
			number:   i,
			previous: currentDegree,
		}
		currentDegree.next = newDegree
		currentDegree = newDegree
	}
	currentDegree.next = firstDegree

	const forward = DegreeNum(103)
	expectedDegreeNum := forward - (forward/amountOfDegrees)*amountOfDegrees + 1
	result := firstDegree.GetForwardDegreeByDegreeNum(forward)
	assert.Equal(t, expectedDegreeNum, result.Number(), "expected: %d, actual: %d", expectedDegreeNum, result.Number())
}

func TestDegreesIterator_GetAll(t *testing.T) {
	t.Run("GetAll: non-empty input channel", func(t *testing.T) {
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

		iter := DegreesIterator(input)

		result := iter.GetAllDegrees()
		if !reflect.DeepEqual(result, expected) {
			assert.Equal(t, expected, result, "expected: %+v, result: %+v", expected, result)
		}
	})

	t.Run("GetAll: empty input channel", func(t *testing.T) {
		input := make(chan *Degree)
		var expected []*Degree

		go func(input chan *Degree, expected []*Degree) {
			for _, degree := range expected {
				input <- degree
			}
			close(input)
		}(input, expected)

		iter := DegreesIterator(input)

		result := iter.GetAllDegrees()
		if !reflect.DeepEqual(result, expected) {
			assert.Equal(t, expected, result, "expected: %+v, result: %+v", expected, result)
		}
	})
}

func TestDegreesIterator_GetAllNotes(t *testing.T) {
	t.Run("GetAllNotes: get all notes from degrees iterator", func(t *testing.T) {
		firstDegree := generateDegreesWithNotes(true, TemplateAeolian(), newNote(C))
		iter := firstDegree.IterateOneRound(false)
		notes := iter.GetAllNotes()

		currentDegree := firstDegree
		for _, note := range notes {
			if !reflect.DeepEqual(*currentDegree.Note(), note) {
				t.Errorf("expected: %+v, result: %+v", *currentDegree.Note(), note)
			}
			currentDegree = currentDegree.GetNext()
		}
	})
}

func TestDegree_GetDegrees(t *testing.T) {
	firstDegreeNum := DegreeNum(1)
	firstDegree := &Degree{
		number: firstDegreeNum,
	}

	currentDegree := firstDegree
	amountOfDegrees := DegreeNum(7)
	for i := DegreeNum(2); i <= amountOfDegrees; i++ {
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

		for i := DegreeNum(2); i <= amountOfDegrees; i++ {
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
}

func TestDegree_IterateOneRound(t *testing.T) {
	mode0, err := MakeNewCustomMode(ModeTemplate{2, 2, 2, 3, 2, 1}, "B", "Custom Mode 0")
	assert.NoError(t, err, "error: %+v", err)
	mode1, err := MakeNewCustomMode(ModeTemplate{2, 2, 2, 2, 2, 2}, "C", "Custom Mode 1")
	assert.NoError(t, err)
	mode2, err := MakeNewCustomMode(ModeTemplate{1, 2, 2, 2, 2, 2, 1}, "C#", "Custom Mode 2")
	assert.NoError(t, err)
	mode3, err := MakeNewCustomMode(ModeTemplate{1, 1, 1, 2, 2, 2, 2, 1}, "C#", "Custom Mode 3")
	assert.NoError(t, err)
	mode4, err := MakeNewCustomMode(ModeTemplate{12}, "A", "Custom Mode 4")
	assert.NoError(t, err)
	modes := []*Mode{mode0, mode1, mode2, mode3, mode4}

	mts := InitModeTemplatesStore()
	notes := GetFullChromaticScale()
	for modeName, modeTemplate := range mts {
		for _, note := range notes {
			mode, err := MakeNewCustomMode(modeTemplate, note.Name().String(), modeName)
			assert.NoError(t, err)
			modes = append(modes, mode)
		}
	}

	t.Run("IterateOneRound Iterating through cycled modes to right", func(t *testing.T) {
		for _, mode := range modes {
			firstDegree := mode.GetFirstDegree()
			var degreesAmount DegreeNum
			var degree *Degree
			for degree = range mode.degree.IterateOneRound(false) {
				degreesAmount++
			}
			assert.Equal(t, unsafe.Pointer(degree.GetNext()), unsafe.Pointer(firstDegree), "next pointer to last degree expected: %+v, actual: %+v", unsafe.Pointer(firstDegree), unsafe.Pointer(degree.GetNext()))
			assert.Equal(t, unsafe.Pointer(degree), unsafe.Pointer(firstDegree.GetPrevious()), "previous pointer to first degree expected: %+v, actual: %+v", unsafe.Pointer(firstDegree.GetPrevious()), unsafe.Pointer(degree))
			assert.Equal(t, mode.Length(), degreesAmount, "degrees amount expected: %d, actual: %d", mode.Length(), degreesAmount)
		}
	})

	t.Run("IterateOneRound Iterating through cycled modes to left", func(t *testing.T) {
		for _, mode := range modes {
			firstDegree := mode.GetFirstDegree()
			var degreesAmount DegreeNum
			var degree *Degree
			for degree = range mode.degree.IterateOneRound(true) {
				degreesAmount++
			}
			assert.Equal(t, unsafe.Pointer(degree.GetPrevious()), unsafe.Pointer(firstDegree), "previous pointer to first degree expected: %+v, actual: %+v", unsafe.Pointer(firstDegree), unsafe.Pointer(degree.GetPrevious()))
			assert.Equal(t, unsafe.Pointer(degree), unsafe.Pointer(firstDegree.GetNext()), "next pointer to second degree expected: %+v, actual: %+v", unsafe.Pointer(firstDegree.GetNext()), unsafe.Pointer(degree))
			assert.Equal(t, mode.Length(), degreesAmount, "degrees amount expected: %d, actual: %d", mode.Length(), degreesAmount)
		}
	})

	t.Run("IterateOneRound Iterating through not cycled modes", func(t *testing.T) {
		firstDegree := generateDegrees(7, false)
		mode0 := &Mode{"Custom Mode 0", firstDegree}
		var degrees0 []*Degree
		degrees0 = append(degrees0, firstDegree.IterateOneRound(false).GetAllDegrees()...)

		degree1 := &Degree{}
		mode1 := &Mode{"Custom Mode 1", degree1}
		degrees1 := []*Degree{degree1}

		testCases := []struct {
			mode    *Mode
			degrees []*Degree
		}{
			{
				mode0,
				degrees0,
			},
			{
				mode1,
				degrees1,
			},
		}

		for _, testCase := range testCases {
			firstDegree := testCase.mode.GetFirstDegree()
			var i uint
			for degree := range firstDegree.IterateOneRound(false) {
				assert.True(t, degree.EqualByDegreeNum(testCase.degrees[i]), "expected degree: %+v, actual: %+v", testCase.degrees[i], degree)
				i++
			}
		}
	})
}

func TestDegree_sortByAbsoluteModalPositions(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	n10 := rand.Intn(10) - 5  //nolint:gosec
	n20 := rand.Intn(20) - 10 //nolint:gosec

	getDegrees := func() (*Degree, *Degree) {
		firstDegree := &Degree{number: 1, absoluteModalPosition: NewModalPositionByWeight(Weight(n10))}
		lastDegree := firstDegree
		for i := DegreeNum(2); i <= DegreesInTonality; i++ {
			degree := &Degree{number: i, absoluteModalPosition: NewModalPositionByWeight(Weight(n20))}
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

	t.Run("test sort by AMP of not cycled degrees chain", func(t *testing.T) {
		firstDegree, _ := getDegrees()
		firstSortedDegree := firstDegree.sortByAbsoluteModalPositions()
		testingFunc(t, firstSortedDegree)
	})

	t.Run("test sort by AMP of cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		firstSortedDegree := firstDegree.sortByAbsoluteModalPositions()
		testingFunc(t, firstSortedDegree)
	})

	t.Run("test sort by AMP in case of degree without AMP", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		firstDegree.GetNext().absoluteModalPosition = nil // just one random degree without set absolute modal position
		firstSortedDegree := firstDegree.sortByAbsoluteModalPositions()
		assert.Nil(t, firstSortedDegree)
	})
}

func TestDegree_String(t *testing.T) {
	testCases := []*Degree{
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  &Degree{},
			note:                  &Note{},
			modalCharacteristics:  []ModalCharacteristic{},
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              nil,
			next:                  &Degree{},
			note:                  &Note{},
			modalCharacteristics:  []ModalCharacteristic{},
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  nil,
			note:                  &Note{},
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
			note:                  &Note{},
			modalCharacteristics:  nil,
			absoluteModalPosition: &ModalPosition{},
		},
		{
			number:                0,
			halfTonesFromPrime:    0,
			previous:              &Degree{},
			next:                  &Degree{},
			note:                  &Note{},
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
			note: &Note{
				name: C,
			},
			modalCharacteristics: []ModalCharacteristic{{
				name:     DegreeCharacteristic2xAug,
				degree:   &Degree{},
				interval: &Interval{},
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
		assert.NotPanics(t, func() { _ = degree.String() }, "degree: %+v", degree) //nolint:scopelint
	}
}

func TestDegree_Equal(t *testing.T) {
	testCases := []struct {
		degree1, degree2 *Degree
		want             bool
		testNumber       uint8
	}{
		{
			degree1:    &Degree{number: 1, note: &Note{name: C}},
			degree2:    &Degree{number: 1, note: &Note{name: C}},
			want:       true,
			testNumber: 0,
		},
		{
			degree1:    &Degree{number: 1, note: &Note{name: C}},
			degree2:    &Degree{number: 2, note: &Note{name: C}},
			want:       false,
			testNumber: 2,
		},
		{
			degree1:    &Degree{number: 1, note: &Note{name: C}},
			degree2:    &Degree{number: 1, note: &Note{name: D}},
			want:       false,
			testNumber: 2,
		},
		{
			degree1:    &Degree{number: 1, note: &Note{name: C}},
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
		assert.Equal(t, testCase.want, testCase.degree1.Equal(testCase.degree2), "test number: %d, expected: %+v, actual: %+v", testCase.testNumber, testCase.degree1, testCase.degree2)
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
		note:                  MustNewNote(C),
		modalCharacteristics:  ModalCharacteristics{},
		absoluteModalPosition: &ModalPosition{name: ModalPositionNameHigh, weight: -5},
	}

	// Act
	copiedDegree := d.Copy()

	// Assert
	if copiedDegree.Number() != d.Number() {
		t.Errorf("DegreeNum not copied properly. Expected: %d, Actual: %d", d.Number(), copiedDegree.Number())
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
	if copiedDegree.absoluteModalPosition.GetWeight() != d.absoluteModalPosition.GetWeight() ||
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
		note:                  MustNewNote(C),
		modalCharacteristics:  ModalCharacteristics{},
		absoluteModalPosition: &ModalPosition{name: ModalPositionNameHigh, weight: -5},
	}

	// Act
	copiedDegree := d.CopyCut()

	// Assert
	if copiedDegree.Number() != d.Number() {
		t.Errorf("DegreeNum not copied properly. Expected: %d, Actual: %d", d.Number(), copiedDegree.Number())
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
	if copiedDegree.absoluteModalPosition.GetWeight() != d.absoluteModalPosition.GetWeight() ||
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
		t.Errorf("Expected degree1 to have no previous degree, but got %v", degree1.previous)
	}
	if degree2.next != nil {
		t.Errorf("Expected degree2 to have no next degree, but got %v", degree2.next)
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
	// Create the first degree
	rootDegree := &Degree{number: 1}

	// Create the second degree
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
	// assert.Panics(t, func() { degree1.InsertPrevious(degree2) })

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

	// Create another degree and attach it as next to the previous degree
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

	// Create another degree and attach it as previous to the previous degree
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
		for i := DegreeNum(2); i <= DegreesInTonality; i++ {
			degree := &Degree{number: i}
			lastDegree.AttachNext(degree)
			lastDegree = degree
		}

		return firstDegree, lastDegree
	}

	t.Run("get ending degrees of the not cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		var degree, result *Degree
		for i := DegreeNum(0); i < DegreesInTonality; i++ {
			degree = firstDegree.GetForwardDegreeByDegreeNum(i)
			result = degree.GetLast(false)
			assert.True(t, result.EqualByDegreeNum(lastDegree), "expected lastDegree: %+v, actual result: %+v", lastDegree, result)
			result = degree.GetLast(true)
			assert.True(t, result.EqualByDegreeNum(firstDegree), "expected firstDegree: %+v, actual result: %+v", firstDegree, result)
		}
	})

	t.Run("get ending degrees of the cycled degrees chain", func(t *testing.T) {
		firstDegree, lastDegree := getDegrees()
		lastDegree.AttachNext(firstDegree)
		var degree, result *Degree
		for i := DegreeNum(0); i < DegreesInTonality; i++ {
			degree = firstDegree.GetForwardDegreeByDegreeNum(i)
			result = degree.GetLast(false)
			assert.True(t, result.EqualByDegreeNum(degree.GetPrevious()), "expected lastDegree: %+v, actual result: %+v", lastDegree, result)
			result = degree.GetLast(true)
			assert.True(t, result.EqualByDegreeNum(degree.GetNext()), "expected firstDegree: %+v, actual result: %+v", firstDegree, result)
		}
	})

	t.Run("get ending degrees from nil degree", func(t *testing.T) {
		var degree *Degree
		result := degree.GetLast(true)
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
