package degree

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

const degreesInTonality = Number(17)

func TestDegreeNode_Number(t *testing.T) {
	expected := Number(3)
	dn := Node{degree: Degree{number: expected}}
	assert.Equal(t, expected, dn.Number())
}

func TestDegreeNode_HalfTonesFromPrime(t *testing.T) {
	expected := halftone.HalfTones(3)
	dn := Node{degree: Degree{halfTonesFromPrime: expected}}

	result := dn.HalfTonesFromPrime()

	if result != expected {
		t.Errorf("Node.HalfTonesFromPrime() = %d; want %d", result, expected)
	}
}

func TestDegreeNode_GetNext(t *testing.T) {
	type testCase struct {
		dn   *Node
		want *Node
	}

	constructTestCase := func(dn *Node, next *Node) *testCase {
		if dn != nil {
			dn.next = next
		}

		return &testCase{
			dn:   dn,
			want: next,
		}
	}

	testCases := []*testCase{
		constructTestCase(&Node{degree: Degree{number: 1}}, &Node{degree: Degree{number: 2}}),
		constructTestCase(&Node{degree: Degree{number: 1}}, nil),
		constructTestCase(nil, nil),
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.dn.GetNext())
	}
}

func TestDegreeNode_SetNext(t *testing.T) {
	dn1 := &Node{degree: Degree{number: 1}}
	dn2 := &Node{degree: Degree{number: 2}}
	dn1.SetNext(dn2)

	if dn1.next != dn2 {
		t.Errorf("dn1.Next = %v; want %v", dn1.next, dn2)
	}
}

func TestDegreeNode_GetPrevious(t *testing.T) {
	t.Run("GetPrevious: get previous node", func(t *testing.T) {
		dn := &Node{previous: &Node{}}

		previousNode := dn.GetPrevious()

		if previousNode != dn.previous {
			t.Errorf("GetPrevious returned incorrect result: got %v, want %v", previousNode, dn.previous)
		}
	})

	t.Run("GetPrevious: get previous node from nil", func(t *testing.T) {
		var nilNode *Node
		assert.Nil(t, nilNode.GetPrevious())
	})
}

func TestDegreeNode_SetPrevious(t *testing.T) {
	dn := &Node{}
	prev := &Node{}

	dn.SetPrevious(prev)

	if dn.previous != prev {
		t.Errorf("Expected previous node to be %v, but got %v", prev, dn.previous)
	}

	dn2 := &Node{}
	dn2.SetPrevious(nil)

	if dn2.previous != nil {
		t.Errorf("Expected previous node to be nil, but got %v", dn2.previous)
	}
}

func TestDegreeNode_Note(t *testing.T) {
	t.Run("get note from node", func(t *testing.T) {
		expectedNote := note.C.NewNote()
		dn := Node{degree: Degree{note: expectedNote}}

		if !dn.Note().Equal(expectedNote) {
			t.Errorf("Expected Note %v but got %v", expectedNote, dn.Note())
		}
	})

	t.Run("get note from nil node returns zero value", func(t *testing.T) {
		var nilNode *Node
		// Note: nilNode.Note() returns note.Note{} which is equivalent to C natural
		// because Letter's zero value is LetterC
		assert.Equal(t, note.Note{}, nilNode.Note())
	})
}

func TestDegreeNode_SetNote(t *testing.T) {
	t.Run("TestDegreeNode_SetNote: set note", func(t *testing.T) {
		dn := &Node{}
		n := note.C.NewNote()

		dn.SetNote(n)

		if !dn.degree.note.Equal(n) {
			t.Errorf("Expected note %v, but got %v", n, dn.degree.note)
		}
	})

	t.Run("TestDegreeNode_SetNote: set zero value note equals C natural", func(t *testing.T) {
		dn := &Node{degree: Degree{note: note.D.NewNote()}}

		// Note: note.Note{} is equivalent to C natural
		dn.SetNote(note.Note{})

		// The zero value of Note equals C natural
		assert.Equal(t, note.Note{}, dn.degree.note)
		assert.True(t, dn.degree.note.Equal(note.C.NewNote()))
	})

	t.Run("TestDegreeNode_SetNote: overwrite note", func(t *testing.T) {
		dn := &Node{degree: Degree{note: note.C.NewNote()}}

		n := note.D.NewNote()
		dn.SetNote(n)

		if !dn.degree.note.Equal(n) {
			t.Errorf("Expected note %v, but got %v", n, dn.degree.note)
		}
	})
}

func TestDegreeNode_ModalCharacteristics(t *testing.T) {
	mc := ModalCharacteristics{{
		name:   CharacteristicClean,
		degree: &Node{degree: Degree{number: 1, note: note.C.NewNote()}},
	}, {
		name:   CharacteristicAug,
		degree: &Node{degree: Degree{number: 2, note: note.DSHARP.NewNote()}},
	}}
	dn := Node{degree: Degree{modalCharacteristics: mc}}
	result := dn.ModalCharacteristics()
	if !reflect.DeepEqual(mc, result) {
		t.Errorf("Expected modal characteristics %v but got %v", mc, result)
	}
}

func TestDegreeNode_AbsoluteModalPosition(t *testing.T) {
	amp := ModalPosition{name: ModalPositionNameNeutral, weight: 0}
	dn := Node{degree: Degree{absoluteModalPosition: amp}}
	result := dn.AbsoluteModalPosition()
	if !reflect.DeepEqual(amp, result) {
		t.Errorf("Expected absolute modal position %v but got %v", amp, result)
	}
}

func TestDegreeNode_NoteReturnsNote(t *testing.T) {
	n := note.C.NewNote()
	dn := &Node{degree: Degree{note: n}}

	assert.Equal(t, n, dn.Note())
}

func TestDegreeNode_getDegreeByDegreeNum(t *testing.T) {
	testingFunc := func(t *testing.T, firstNode *Node) {
		t.Helper()
		for dn := range firstNode.IterateOneRound(false) {
			for d := range firstNode.IterateOneRound(false) {
				resultNode := dn.GetDegreeByDegreeNum(d.Number())
				assert.Equal(t, d.Number(), resultNode.Number(), "node number: %d, expected: %+v, actual: %+v", dn.Number(), dn, resultNode)
			}
		}
	}

	t.Run("TestDegreeNode_getDegreeByDegreeNum: uncycled chain", func(t *testing.T) {
		testingFunc(t, generateDegreeNodes(3, false))
	})

	t.Run("TestDegreeNode_getDegreeByDegreeNum: cycled chain", func(t *testing.T) {
		testingFunc(t, generateDegreeNodes(3, true))
	})

	t.Run("TestDegreeNode_getDegreeByDegreeNum: nil node", func(t *testing.T) {
		var nilNode *Node
		assert.Nil(t, nilNode.GetDegreeByDegreeNum(5))
	})

	t.Run("TestDegreeNode_getDegreeByDegreeNum: negative case", func(t *testing.T) {
		assert.Nil(t, generateDegreeNodes(3, true).GetDegreeByDegreeNum(5))
	})
}

func TestDegreeNode_GetForwardDegreeByDegreeNum(t *testing.T) {
	t.Run("GetForwardDegreeByDegreeNum: positive", func(t *testing.T) {
		firstNodeNum := Number(1)
		firstNode := &Node{
			degree: Degree{number: firstNodeNum},
		}

		currentNode := firstNode
		amountOfNodes := Number(7)
		for i := Number(2); i <= amountOfNodes; i++ {
			newNode := &Node{
				degree:   Degree{number: i},
				previous: currentNode,
			}
			currentNode.next = newNode
			currentNode = newNode
		}
		currentNode.next = firstNode

		const forward = Number(103)
		expectedNodeNum := forward - (forward/amountOfNodes)*amountOfNodes + 1
		result := firstNode.GetForwardDegreeByDegreeNum(forward)
		assert.Equal(t, expectedNodeNum, result.Number(), "expected: %d, actual: %d", expectedNodeNum, result.Number())
	})

	t.Run("GetForwardDegreeByDegreeNum: get from nil node", func(t *testing.T) {
		var nilNode *Node
		assert.Nil(t, nilNode.GetForwardDegreeByDegreeNum(5))
	})
}

func TestGetAllDegrees(t *testing.T) {
	t.Run("GetAllDegrees: non-empty chain", func(t *testing.T) {
		firstNode := generateDegreeNodes(4, false)
		result := firstNode.IterateOneRound(false).GetAllDegrees()

		if len(result) != 4 {
			t.Errorf("expected 4 nodes, got %d", len(result))
		}
		for i, dn := range result {
			if dn.Number() != Number(i+1) {
				t.Errorf("expected node number %d, got %d", i+1, dn.Number())
			}
		}
	})

	t.Run("GetAllDegrees: empty iterator", func(t *testing.T) {
		var nilNode *Node
		result := nilNode.IterateOneRound(false).GetAllDegrees()
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %d elements", len(result))
		}
	})
}

func TestDegreeNode_GetDegrees(t *testing.T) {
	firstNodeNum := Number(1)
	firstNode := &Node{
		degree: Degree{number: firstNodeNum},
	}

	currentNode := firstNode
	amountOfNodes := Number(7)
	for i := Number(2); i <= amountOfNodes; i++ {
		newNode := &Node{
			degree:   Degree{number: i},
			previous: currentNode,
		}
		currentNode.next = newNode
		currentNode = newNode
	}

	lastNode := currentNode

	t.Run("get next node till the last node num", func(t *testing.T) {
		currentNode := firstNode
		if currentNode.Number() != firstNodeNum {
			t.Errorf("first node num: %d, firstNodeNum: %d", currentNode.Number(), firstNodeNum)
		}

		for i := Number(2); i <= amountOfNodes; i++ {
			currentNode = currentNode.GetNext()
			if i != currentNode.Number() {
				t.Errorf("node in cycle: %d, node number: %d", i, currentNode.Number())
			}
		}
	})

	t.Run("get previous node till the first node num", func(t *testing.T) {
		currentNode := lastNode
		if currentNode.Number() != amountOfNodes {
			t.Errorf("last node num: %d, amount of nodes: %d", currentNode.Number(), amountOfNodes)
		}

		for i := amountOfNodes - 1; i >= firstNodeNum; i-- {
			currentNode = currentNode.GetPrevious()
			if i != currentNode.Number() {
				t.Errorf("node in cycle: %d, node number: %d", i, currentNode.Number())
			}
		}
	})

	t.Run("get next node till the end", func(t *testing.T) {
		currentNode := firstNode
		if currentNode.Number() != firstNodeNum {
			t.Errorf("first node num: %d, firstNodeNum: %d", currentNode.Number(), firstNodeNum)
		}

		for currentNode.next != nil {
			currentNode = currentNode.GetNext()
		}

		if currentNode.Number() != amountOfNodes {
			t.Errorf("last node num: %d, amount of nodes: %d", currentNode.Number(), amountOfNodes)
		}
	})

	t.Run("get previous node till the end", func(t *testing.T) {
		currentNode := lastNode
		if currentNode.Number() != amountOfNodes {
			t.Errorf("last node num: %d, amount of nodes: %d", currentNode.Number(), amountOfNodes)
		}

		for currentNode.previous != nil {
			currentNode = currentNode.GetPrevious()
		}

		if currentNode.Number() != firstNodeNum {
			t.Errorf("last node num: %d, firstNodeNum: %d", currentNode.Number(), firstNodeNum)
		}
	})

	t.Run("get previous node from nil node", func(t *testing.T) {
		var nilNode *Node
		assert.Nil(t, nilNode)
	})
}

func TestDegreeNode_IterateOneRound(t *testing.T) {
	t.Run("IterateOneRound Iterating through cycled nodes to right", func(t *testing.T) {
		dn1 := generateDegreeNodes(7, true)
		var i Number
		for dn := range dn1.IterateOneRound(false) {
			i++
			assert.Equal(t, dn.Number(), i)
		}
	})

	t.Run("IterateOneRound Iterating through cycled nodes to left", func(t *testing.T) {
		dn1 := generateDegreeNodes(7, true)
		i := Number(1)
		for dn := range dn1.IterateOneRound(true) {
			assert.Equal(t, dn.Number(), i)
			if i == 1 {
				i = 7
			} else {
				i--
			}
		}
	})

	t.Run("IterateOneRound Iterating through not cycled nodes to right", func(t *testing.T) {
		dn1 := generateDegreeNodes(7, false)
		var i Number
		for dn := range dn1.IterateOneRound(false) {
			i++
			assert.Equal(t, dn.Number(), i)
		}
	})
}

func TestDegreeNode_sortByAbsoluteModalPositions(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	n10 := rand.Intn(3) //nolint:gosec
	n20 := n10
	for n20 == n10 {
		n20 = rand.Intn(2) //nolint:gosec
	}

	getNodes := func() (*Node, *Node) {
		firstNode := &Node{degree: Degree{number: 1, absoluteModalPosition: NewModalPositionByWeight(Weight(n10))}} //nolint:gosec
		lastNode := firstNode
		for i := Number(2); i <= degreesInTonality; i++ {
			dn := &Node{degree: Degree{number: i, absoluteModalPosition: NewModalPositionByWeight(Weight(n20))}} //nolint:gosec
			lastNode.AttachNext(dn)
			lastNode = dn
		}

		return firstNode, lastNode
	}

	testingFunc := func(t *testing.T, firstSortedNode *Node) {
		t.Helper()
		firstNode := firstSortedNode
		var comparison bool
		isFirst := true
		for dn := range firstSortedNode.IterateOneRound(false) {
			if isFirst {
				isFirst = false
				continue
			}
			if dn.NextExists() {
				comparison = dn.degree.absoluteModalPosition.Weight() <= dn.GetNext().degree.absoluteModalPosition.Weight()
				if dn.GetNext().Number() != firstNode.Number() {
					assert.True(t, comparison, "current - node Num: %d, w: %d, next - node Num: %d, w: %d", dn.Number(), dn.degree.absoluteModalPosition.Weight(), dn.GetNext().Number(), dn.GetNext().degree.absoluteModalPosition.Weight())
				} else {
					assert.False(t, comparison, "current - node Num: %d, w: %d, next - node Num: %d, w: %d", dn.Number(), dn.degree.absoluteModalPosition.Weight(), dn.GetNext().Number(), dn.GetNext().degree.absoluteModalPosition.Weight())
				}
			}
		}
	}

	t.Run("test sort by AMP of not cycled chain", func(t *testing.T) {
		firstNode, _ := getNodes()
		firstSortedNode := firstNode.SortByAbsoluteModalPositions(true)
		testingFunc(t, firstSortedNode)
	})

	t.Run("test sort by AMP of cycled chain", func(t *testing.T) {
		firstNode, lastNode := getNodes()
		lastNode.AttachNext(firstNode)
		firstSortedNode := firstNode.SortByAbsoluteModalPositions(true)
		testingFunc(t, firstSortedNode)
	})

	t.Run("test sort by AMP in case of node without AMP", func(t *testing.T) {
		firstNode, lastNode := getNodes()
		lastNode.AttachNext(firstNode)
		firstNode.GetNext().degree.absoluteModalPosition = ModalPosition{} // just one random node without set absolute modal position
		firstSortedNode := firstNode.SortByAbsoluteModalPositions(true)
		assert.Nil(t, firstSortedNode)
	})

	t.Run("test sort by AMP of cycled chain Dorian mode", func(t *testing.T) {
		seventhNode := New(7, 0, nil, nil, note.BFLAT.NewNote(), nil, ModalPosition{"", 2})
		sixthNode := New(6, 0, nil, seventhNode, note.A.NewNote(), nil, ModalPosition{"", 7})
		fifthNode := New(5, 0, nil, sixthNode, note.G.NewNote(), nil, ModalPosition{"", 6})
		fourthNode := New(4, 0, nil, fifthNode, note.F.NewNote(), nil, ModalPosition{"", 3})
		thirdNode := New(3, 0, nil, fourthNode, note.EFLAT.NewNote(), nil, ModalPosition{"", 1})
		secondNode := New(2, 0, nil, thirdNode, note.D.NewNote(), nil, ModalPosition{"", 5})
		firstNode := New(1, 0, nil, secondNode, note.C.NewNote(), nil, ModalPosition{"", 4})

		sixthNodeSorted := New(6, 0, nil, nil, note.A.NewNote(), nil, ModalPosition{"", 7})
		fifthNodeSorted := New(5, 0, nil, sixthNodeSorted, note.G.NewNote(), nil, ModalPosition{"", 6})
		secondNodeSorted := New(2, 0, nil, fifthNodeSorted, note.D.NewNote(), nil, ModalPosition{"", 5})
		firstNodeSorted := New(1, 0, nil, secondNodeSorted, note.C.NewNote(), nil, ModalPosition{"", 4})
		fourthNodeSorted := New(4, 0, nil, firstNodeSorted, note.F.NewNote(), nil, ModalPosition{"", 3})
		seventhNodeSorted := New(7, 0, nil, fourthNodeSorted, note.BFLAT.NewNote(), nil, ModalPosition{"", 2})
		thirdNodeSorted := New(3, 0, nil, seventhNodeSorted, note.EFLAT.NewNote(), nil, ModalPosition{"", 1})

		seventhNode.AttachNext(firstNode)

		sortedNode := firstNode.SortByAbsoluteModalPositions(true)

		expectedNodes := thirdNodeSorted.IterateOneRound(false).GetAllDegrees()
		resultNodes := sortedNode.IterateOneRound(false).GetAllDegrees()
		for i, resultNode := range resultNodes {
			assert.True(t, expectedNodes[i].EqualByDegreeNum(resultNode), "resulting node number: %d, expectedNode number: %d", resultNode.Number(), expectedNodes[i].Number())
		}
	})
}

func TestDegreeNode_String(t *testing.T) {
	testCases := []*Node{
		nil,
		{
			degree: Degree{
				number:                0,
				halfTonesFromPrime:    0,
				note:                  note.Note{},
				modalCharacteristics:  []ModalCharacteristic{},
				absoluteModalPosition: ModalPosition{},
			},
			previous: &Node{},
			next:     &Node{},
		},
		{
			degree: Degree{
				number:                0,
				halfTonesFromPrime:    0,
				note:                  note.Note{},
				modalCharacteristics:  []ModalCharacteristic{},
				absoluteModalPosition: ModalPosition{},
			},
			previous: nil,
			next:     &Node{},
		},
		{
			degree: Degree{
				number:                0,
				halfTonesFromPrime:    0,
				note:                  note.Note{},
				modalCharacteristics:  []ModalCharacteristic{},
				absoluteModalPosition: ModalPosition{},
			},
			previous: &Node{},
			next:     nil,
		},
		{
			degree: Degree{
				number:                0,
				halfTonesFromPrime:    0,
				note:                  note.Note{},
				modalCharacteristics:  nil,
				absoluteModalPosition: ModalPosition{},
			},
			previous: &Node{},
			next:     &Node{},
		},
		{
			degree: Degree{
				number:                0,
				halfTonesFromPrime:    0,
				note:                  note.Note{},
				modalCharacteristics:  []ModalCharacteristic{},
				absoluteModalPosition: ModalPosition{},
			},
			previous: nil,
			next:     nil,
		},
		{
			degree: Degree{
				number:             1,
				halfTonesFromPrime: 1,
				note:               note.C.NewNote(),
				modalCharacteristics: []ModalCharacteristic{{
					name:   Characteristic2xAug,
					degree: &Node{},
					relativeModalPosition: ModalPosition{
						name:   ModalPositionNameNeutral,
						weight: 0,
					},
				}},
				absoluteModalPosition: ModalPosition{
					name:   ModalPositionNameNeutral,
					weight: 0,
				},
			},
			previous: &Node{
				degree: Degree{
					number:             7,
					halfTonesFromPrime: 12,
				},
			},
			next: &Node{
				degree: Degree{
					number:             2,
					halfTonesFromPrime: 3,
				},
			},
		},
	}

	for _, dn := range testCases {
		assert.NotPanics(t, func() { _ = dn.String() }, "node: %+v", dn) //nolint:scopelint
	}
}

func TestDegreeNode_Equal(t *testing.T) {
	testCases := []struct {
		node1, node2 *Node
		want         bool
		testNumber   uint8
	}{
		{
			node1:      &Node{degree: Degree{number: 1, note: note.C.NewNote()}},
			node2:      &Node{degree: Degree{number: 1, note: note.C.NewNote()}},
			want:       true,
			testNumber: 0,
		},
		{
			node1:      &Node{degree: Degree{number: 1, note: note.C.NewNote()}},
			node2:      &Node{degree: Degree{number: 2, note: note.C.NewNote()}},
			want:       false,
			testNumber: 1,
		},
		{
			node1:      &Node{degree: Degree{number: 1, note: note.C.NewNote()}},
			node2:      &Node{degree: Degree{number: 1, note: note.D.NewNote()}},
			want:       false,
			testNumber: 2,
		},
		{
			// Note: note.Note{} is equivalent to C natural, so these are equal
			node1:      &Node{degree: Degree{number: 1, note: note.C.NewNote()}},
			node2:      &Node{degree: Degree{number: 1, note: note.Note{}}},
			want:       true,
			testNumber: 3,
		},
		{
			node1:      &Node{degree: Degree{number: 1, note: note.Note{}}},
			node2:      &Node{degree: Degree{number: 1, note: note.Note{}}},
			want:       true,
			testNumber: 4,
		},
		{
			node1:      &Node{degree: Degree{number: 1, note: note.Note{}}},
			node2:      nil,
			want:       false,
			testNumber: 5,
		},
		{
			// Different notes should not be equal
			node1:      &Node{degree: Degree{number: 1, note: note.D.NewNote()}},
			node2:      &Node{degree: Degree{number: 1, note: note.Note{}}},
			want:       false,
			testNumber: 6,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.node1.Equal(testCase.node2), "test number: %d, expected: %+v, actual: %+v", testCase.testNumber, testCase.node1, testCase.node2)
	}
}

func TestDegreeNode_EqualByDegreeNum(t *testing.T) {
	// Test the case where both nodes are nil
	if node1, node2 := (*Node)(nil), (*Node)(nil); node1.EqualByDegreeNum(node2) {
		t.Error("Expected nil nodes to be equal")
	}

	// Test the case where one node is nil
	node1 := &Node{degree: Degree{number: 1}}
	if node1.EqualByDegreeNum(nil) {
		t.Error("Expected non-nil node to not be equal to nil")
	}

	// Test the case where node numbers are equal
	node2 := &Node{degree: Degree{number: 1}}
	if !node1.EqualByDegreeNum(node2) {
		t.Error("Expected nodes with equal numbers to be equal")
	}

	// Test the case where node numbers are not equal
	node3 := &Node{degree: Degree{number: 2}}
	if node1.EqualByDegreeNum(node3) {
		t.Error("Expected nodes with different numbers to not be equal")
	}
}

func TestDegreeNode_Copy(t *testing.T) {
	// Arrange
	dn := &Node{
		degree: Degree{
			number:                1,
			halfTonesFromPrime:    2,
			note:                  note.MustNewWithOctave(note.C, octave.Number1),
			modalCharacteristics:  ModalCharacteristics{},
			absoluteModalPosition: ModalPosition{name: ModalPositionNameHigh, weight: -5},
		},
		previous: &Node{},
		next:     &Node{},
	}

	// Act
	copiedNode := dn.Copy()

	// Assert
	if copiedNode.Number() != dn.Number() {
		t.Errorf("Number not copied properly. Expected: %d, Actual: %d", dn.Number(), copiedNode.Number())
	}
	if copiedNode.degree.halfTonesFromPrime != dn.degree.halfTonesFromPrime {
		t.Errorf("HalfTonesFromPrime not copied properly. Expected: %d, Actual: %d", dn.degree.halfTonesFromPrime, copiedNode.degree.halfTonesFromPrime)
	}
	if copiedNode.previous != dn.previous {
		t.Errorf("Previous not copied properly")
	}
	if copiedNode.next != dn.next {
		t.Errorf("Next not copied properly")
	}
	if copiedNode.degree.note.Name() != dn.degree.note.Name() {
		t.Errorf("Note not copied properly. Expected: %+v, Actual: %+v", dn.degree.note, copiedNode.degree.note)
	}
	if copiedNode.degree.absoluteModalPosition.Weight() != dn.degree.absoluteModalPosition.Weight() ||
		copiedNode.degree.absoluteModalPosition.name != dn.degree.absoluteModalPosition.name {
		t.Errorf("AbsoluteModalPosition not copied properly. Expected: %+v, Actual: %+v", dn.degree.absoluteModalPosition, copiedNode.degree.absoluteModalPosition)
	}
}

func TestDegreeNode_CopyCut(t *testing.T) {
	// Arrange
	dn := &Node{
		degree: Degree{
			number:                1,
			halfTonesFromPrime:    2,
			note:                  note.MustNewWithOctave(note.C, octave.Number0),
			modalCharacteristics:  ModalCharacteristics{},
			absoluteModalPosition: ModalPosition{name: ModalPositionNameHigh, weight: -5},
		},
		previous: &Node{},
		next:     &Node{},
	}

	// Act
	copiedNode := dn.CopyCut()

	// Assert
	if copiedNode.Number() != dn.Number() {
		t.Errorf("Number not copied properly. Expected: %d, Actual: %d", dn.Number(), copiedNode.Number())
	}
	if copiedNode.degree.halfTonesFromPrime != dn.degree.halfTonesFromPrime {
		t.Errorf("HalfTonesFromPrime not copied properly. Expected: %d, Actual: %d", dn.degree.halfTonesFromPrime, copiedNode.degree.halfTonesFromPrime)
	}
	if copiedNode.previous != nil {
		t.Errorf("Previous not set to nil")
	}
	if copiedNode.next != nil {
		t.Errorf("Next not set to nil")
	}
	if copiedNode.degree.note.Name() != dn.degree.note.Name() {
		t.Errorf("Note not copied properly. Expected: %+v, Actual: %+v", dn.degree.note, copiedNode.degree.note)
	}
	if copiedNode.degree.absoluteModalPosition.Weight() != dn.degree.absoluteModalPosition.Weight() ||
		copiedNode.degree.absoluteModalPosition.name != dn.degree.absoluteModalPosition.name {
		t.Errorf("AbsoluteModalPosition not copied properly. Expected: %+v, Actual: %+v", dn.degree.absoluteModalPosition, copiedNode.degree.absoluteModalPosition)
	}
}

func TestDegreeNode_InsertBetween(t *testing.T) {
	node1 := &Node{degree: Degree{number: 1}}
	node2 := &Node{degree: Degree{number: 2}}
	node3 := &Node{degree: Degree{number: 3}}

	node2.InsertBetween(node1, node3)

	if node1.next.Number() != node2.Number() {
		t.Errorf("node1 is not correctly attached to node2. Next Node is: %d", node1.GetNext().Number())
	}
	if node3.previous != node2 {
		t.Errorf("node3 is not correctly attached to node2. Previous Node is: %d", node3.GetPrevious().Number())
	}
	if node2.next.Number() != node3.Number() {
		t.Errorf("node2 is not correctly attached to node3. Next Node is: %d", node2.GetNext().Number())
	}
	if node2.previous != node1 {
		t.Errorf("node2 is not correctly attached to node1. Previous Node is: %d", node2.GetPrevious().Number())
	}
}

func TestDegreeNode_AttachNext(t *testing.T) {
	// Set up initial Nodes
	node1 := &Node{degree: Degree{number: 1}}
	node2 := &Node{degree: Degree{number: 2}}

	// Attach node2 as next node after node1
	node1.AttachNext(node2)

	// Assertions to ensure references are correctly set
	if node1.next != node2 {
		t.Errorf("Expected node2 to be the next node after node1, but got %v", node1.next)
	}
	if node2.previous != node1 {
		t.Errorf("Expected node1 to be the previous node before node2, but got %v", node2.previous)
	}

	// Ensure mutual references are not set
	if node1.previous != nil {
		t.Errorf("Expected node1 to have no previous node, but got %v", node1.previous)
	}
	if node2.next != nil {
		t.Errorf("Expected node2 to have no next node, but got %v", node2.next)
	}

	// Attach node1 as next node after node2
	node2.AttachNext(node1)

	// Assertions to ensure references are correctly set
	if node2.next != node1 {
		t.Errorf("Expected node1 to be the next node after node2, but got %v", node2.next)
	}
	if node1.previous != node2 {
		t.Errorf("Expected node2 to be the previous node before node1, but got %v", node1.previous)
	}

	// Ensure mutual references are set
	if node2.previous != node1 {
		t.Errorf("Expected node1 to be the previous node before node2, but got %v", node2.previous)
	}
	if node1.next != node2 {
		t.Errorf("Expected node2 to be the next node after node1, but got %v", node1.next)
	}

	// Attach node2 as next node after node2 (invalid input)
	node2.AttachNext(node2)

	// Assertions to ensure references are not set
	if node2.next != node2 {
		t.Errorf("Expected node2 to be the next node after node2, but got %v", node1.next)
	}
	if node2.previous != node2 {
		t.Errorf("Expected node2 to be the previous node before node2, but got %v", node1.next)
	}
	if node1.next != node2 {
		t.Errorf("Expected node2 to be the next node after node1, but got %v", node1.next)
	}
	if node1.previous != node2 {
		t.Errorf("Expected node2 to be the previous node before node1, but got %v", node1.next)
	}
}

func TestDegreeNode_InsertNext(t *testing.T) {
	// Create three-node objects
	node1 := &Node{}
	node2 := &Node{}
	node3 := &Node{}

	// Call the InsertNext function on node2 and passing node1 as the new Node object to be inserted
	node2.InsertNext(node1)

	// Check that node1 is attached as next to node2
	if node2.next != node1 {
		t.Errorf("node1 is not correctly attached as next to node2")
	}
	if node1.previous != node2 {
		t.Errorf("node2 is not correctly attached as previous to node1")
	}

	// Call the InsertNext function on node2 (which already has a next) and passing node3 as the new Node object to be inserted
	node2.InsertNext(node3)

	// Check that node3 is attached as next to node2 and node2 is attached as previous to node3 (node1 is in between)
	if node2.next != node3 {
		t.Errorf("node3 is not correctly attached as next to node2")
	}
	if node3.previous != node2 {
		t.Errorf("node2 is not correctly attached as previous to node3")
	}
}

func TestDegreeNode_AttachPrevious(t *testing.T) {
	// Create the first node
	rootNode := &Node{degree: Degree{number: 1}}

	// Create the second node
	secondNode := &Node{degree: Degree{number: 2}}

	// Attach secondNode as previous to rootNode
	rootNode.AttachPrevious(secondNode)

	// Test that AttachPrevious sets the previous node correctly
	if rootNode.GetPrevious() != secondNode {
		t.Errorf("AttachPrevious doesn't set previous node correctly, expected %v but got %v",
			secondNode, rootNode.GetPrevious())
	}

	// Test that AttachPrevious sets the next node correctly
	if secondNode.GetNext() != rootNode {
		t.Errorf("AttachPrevious doesn't set next node correctly, expected %v but got %v",
			rootNode, secondNode.GetNext())
	}

	// Create thirdNode
	thirdNode := &Node{degree: Degree{number: 3}}

	// Attach thirdNode as previous to secondNode
	secondNode.AttachPrevious(thirdNode)

	// Test that the previous node of secondNode is now thirdNode
	if secondNode.GetPrevious() != thirdNode {
		t.Errorf("AttachPrevious doesn't update previous node correctly, expected %v but got %v",
			thirdNode, secondNode.GetPrevious())
	}

	// Test that AttachPrevious sets the next node correctly
	if thirdNode.GetNext() != secondNode {
		t.Errorf("AttachPrevious doesn't set next node correctly, expected %v but got %v",
			secondNode, thirdNode.GetNext())
	}
}

func TestDegreeNode_InsertPrevious(t *testing.T) {
	// create nodes
	node1 := &Node{degree: Degree{number: 1}}
	node2 := &Node{degree: Degree{number: 2}}
	node3 := &Node{degree: Degree{number: 3}}

	// inserting previous node to a nil node should result in panicking
	var nilNode *Node
	assert.Panics(t, func() { nilNode.InsertPrevious(node2) })

	// insert previous node and test for mutual reference
	node1.InsertPrevious(node2)
	assert.Equal(t, node1.previous, node2)
	assert.Equal(t, node2.next, node1)

	// insert another node as the previous node and test for mutual reference
	node1.InsertPrevious(node3)
	assert.Equal(t, node1.previous, node3)
	assert.Equal(t, node3.next, node1)
	assert.Equal(t, node3.previous, node2)
	assert.Equal(t, node2.next, node3)
}

func TestDegreeNode_NextExists(t *testing.T) {
	// Create a node with no next node attached
	node := &Node{degree: Degree{number: 1}}

	// Test that NextExists returns false
	if node.NextExists() {
		t.Errorf("NextExists should return false when next node doesn't exist")
	}

	// Create another node and attach it as next to the previous node
	nextNode := &Node{degree: Degree{number: 2}}
	node.AttachNext(nextNode)

	// Test that NextExists returns true
	if !node.NextExists() {
		t.Errorf("NextExists should return true when next node exists")
	}
}

func TestDegreeNode_PreviousExists(t *testing.T) {
	// Create a node with no previous node attached
	node := &Node{degree: Degree{number: 1}}

	// Test that PreviousExists returns false
	if node.PreviousExists() {
		t.Errorf("PreviousExists should return false when previous node doesn't exist")
	}

	// Create another node and attach it as previous to the previous node
	previousNode := &Node{degree: Degree{number: 2}}
	node.AttachPrevious(previousNode)

	// Test that PreviousExists returns true
	if !node.PreviousExists() {
		t.Errorf("PreviousExists should return true when previous node exists")
	}
}

func TestDegreeNode_GetLast(t *testing.T) {
	getNodes := func() (*Node, *Node) {
		firstNode := &Node{degree: Degree{number: 1}}
		lastNode := firstNode
		for i := Number(2); i <= degreesInTonality; i++ {
			dn := &Node{degree: Degree{number: i}}
			lastNode.AttachNext(dn)
			lastNode = dn
		}

		return firstNode, lastNode
	}

	t.Run("get ending nodes of the not cycled chain", func(t *testing.T) {
		firstNode, lastNode := getNodes()
		var dn, result *Node
		for i := Number(0); i < degreesInTonality; i++ {
			dn = firstNode.GetForwardDegreeByDegreeNum(i)
			result = dn.GetLast(false)
			assert.True(t, result.EqualByDegreeNum(lastNode), "expected lastNode: %+v, actual result: %+v", lastNode, result)
			result = dn.GetLast(true)
			assert.True(t, result.EqualByDegreeNum(firstNode), "expected firstNode: %+v, actual result: %+v", firstNode, result)
		}
	})

	t.Run("get ending nodes of the cycled chain", func(t *testing.T) {
		firstNode, lastNode := getNodes()
		lastNode.AttachNext(firstNode)
		var dn, result *Node
		for i := Number(0); i < degreesInTonality; i++ {
			dn = firstNode.GetForwardDegreeByDegreeNum(i)
			result = dn.GetLast(false)
			assert.True(t, result.EqualByDegreeNum(dn.GetPrevious()), "expected lastNode: %+v, actual result: %+v", lastNode, result)
			result = dn.GetLast(true)
			assert.True(t, result.EqualByDegreeNum(dn.GetNext()), "expected firstNode: %+v, actual result: %+v", firstNode, result)
		}
	})

	t.Run("get ending nodes from nil node", func(t *testing.T) {
		var dn *Node
		result := dn.GetLast(true)
		assert.Nil(t, result, "expected nil, but got: %+v", result)
	})
}

func TestDegreeNode_AttachToTheEnd(t *testing.T) {
	t.Run("TestAttachToTheEnd test case when chain is not cycled", func(t *testing.T) {
		dn1 := &Node{degree: Degree{number: 1}}
		dn2 := &Node{degree: Degree{number: 2}}
		dn3 := &Node{degree: Degree{number: 3}}
		dn4 := &Node{degree: Degree{number: 4}}

		dn1.AttachNext(dn2)
		dn2.AttachNext(dn3)
		dn3.AttachNext(dn4)

		dn5 := &Node{degree: Degree{number: 5}}
		dn1.AttachToTheEnd(dn5, false)

		expected := []*Node{dn1, dn2, dn3, dn4, dn5}

		var i int
		for node := range dn1.IterateOneRound(false) {
			assert.Equal(t, expected[i], node)
			i++
		}

		dn0 := &Node{degree: Degree{number: 0}}
		dn5.AttachToTheEnd(dn0, true)

		expected = []*Node{dn0, dn1, dn2, dn3, dn4, dn5}

		var j int
		for node := range dn0.IterateOneRound(false) {
			assert.Equal(t, expected[j], node)
			j++
		}
	})

	t.Run("TestAttachToTheEnd test case when chain is cycled", func(t *testing.T) {
		dn1 := &Node{degree: Degree{number: 1}}
		dn2 := &Node{degree: Degree{number: 2}}
		dn3 := &Node{degree: Degree{number: 3}}
		dn4 := &Node{degree: Degree{number: 4}}
		dn1.AttachPrevious(dn4)
		dn1.AttachNext(dn2)
		dn2.AttachNext(dn3)
		dn3.AttachNext(dn4)

		dn5 := &Node{degree: Degree{number: 5}}
		dn1.AttachToTheEnd(dn5, false)

		expected := []*Node{dn1, dn2, dn3, dn4, dn5}

		var i int
		for node := range dn1.IterateOneRound(false) {
			assert.Equal(t, expected[i], node)
			i++
		}

		dn6 := &Node{degree: Degree{number: 6}}
		dn1.GetPrevious().AttachToTheEnd(dn6, true)

		expected = append(expected, dn6)

		var j int
		for node := range dn1.IterateOneRound(false) {
			assert.Equal(t, expected[j], node)
			j++
		}
	})
}

func TestDegreeNode_ReverseSequence(t *testing.T) {
	t.Run("TestReverseSequence with cycled sequence", func(t *testing.T) {
		dn := generateDegreeNodes(7, true)
		res := dn.ReverseSequence()
		expectedNodes := dn.GetPrevious().IterateOneRound(true).GetAllDegrees()
		resultNodes := res.IterateOneRound(false).GetAllDegrees()
		for i, node := range resultNodes {
			assert.Equal(t, node.Number(), expectedNodes[i].Number(), "expected: %d, actual: %d", expectedNodes[i].Number(), node.Number())
		}
	})

	t.Run("TestReverseSequence with not cycled sequence", func(t *testing.T) {
		dn := generateDegreeNodes(7, false)
		res := dn.ReverseSequence()
		expectedNodes := dn.GetLast(false).IterateOneRound(true).GetAllDegrees()
		resultNodes := res.IterateOneRound(false).GetAllDegrees()
		for i, node := range resultNodes {
			assert.Equal(t, node.Number(), expectedNodes[i].Number(), "expected: %d, actual: %d", expectedNodes[i].Number(), node.Number())
		}
	})
}

func generateDegreeNodes(nodesAmount uint8, isCycled bool) *Node {
	if nodesAmount < 1 {
		return nil
	}
	nodeNum := Number(1)
	firstNode := newDegreeNodeWithNum(nodeNum)
	currentNode := firstNode
	for nodeNum++; nodeNum <= Number(nodesAmount); nodeNum++ {
		newNode := newDegreeNodeWithNum(nodeNum)
		currentNode.AttachNext(newNode)
		currentNode = newNode
	}

	if isCycled {
		firstNode.AttachPrevious(currentNode)
	}

	return firstNode
}

func newDegreeNodeWithNum(number Number) *Node {
	return New(
		number,
		0,
		nil,
		nil,
		note.Note{},
		nil,
		ModalPosition{},
	)
}
