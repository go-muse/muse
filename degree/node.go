package degree

import (
	"fmt"
	"sort"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

// Node is a node in a doubly-linked list of degrees.
// It wraps a Degree and provides list navigation.
// The degree number is stored in the Degree itself, not in the Node.
type Node struct {
	degree   Degree
	previous *Node
	next     *Node
}

// New creates a new degree node with specified parameters.
// Deprecated: Use NewNode instead for clarity.
func New(
	number Number,
	halfTonesFromPrime halftone.HalfTones,
	previous, next *Node,
	n note.Note,
	modalCharacteristics ModalCharacteristics,
	absoluteModalPosition ModalPosition,
) *Node {
	return &Node{
		degree: Degree{
			number:                number,
			halfTonesFromPrime:    halfTonesFromPrime,
			note:                  n,
			modalCharacteristics:  modalCharacteristics,
			absoluteModalPosition: absoluteModalPosition,
		},
		previous: previous,
		next:     next,
	}
}

// NewNode creates a new degree node with the given degree.
// The degree number is stored in the Degree itself.
func NewNode(degree Degree) *Node {
	return &Node{
		degree: degree,
	}
}

// Number returns degree node's number (its position in a mode).
// Delegates to the underlying Degree.
func (dn *Node) Number() Number {
	if dn == nil {
		return 0
	}

	return dn.degree.number
}

// SetNumber sets degree node's number.
// Delegates to the underlying Degree.
func (dn *Node) SetNumber(number Number) {
	if dn != nil {
		dn.degree.number = number
	}
}

// Degree returns the underlying degree.
func (dn *Node) Degree() Degree {
	if dn == nil {
		return Degree{}
	}

	return dn.degree
}

// SetDegree sets the underlying degree.
func (dn *Node) SetDegree(d Degree) {
	if dn != nil {
		dn.degree = d
	}
}

// HalfTonesFromPrime returns degree's distance from prime in halftones.
func (dn *Node) HalfTonesFromPrime() halftone.HalfTones {
	if dn == nil {
		return 0
	}

	return dn.degree.halfTonesFromPrime
}

// SetHalfTonesFromPrime sets halftone from prime value in halftones.
func (dn *Node) SetHalfTonesFromPrime(halfTones halftone.HalfTones) {
	if dn != nil {
		dn.degree.halfTonesFromPrime = halfTones
	}
}

// GetNext returns next degree node.
func (dn *Node) GetNext() *Node {
	if dn == nil {
		return nil
	}

	return dn.next
}

// SetNext sets next degree node for the current node.
func (dn *Node) SetNext(nextNode *Node) *Node {
	if dn == nil {
		return nil
	}
	dn.next = nextNode

	return dn
}

// GetPrevious returns previous degree node.
func (dn *Node) GetPrevious() *Node {
	if dn == nil {
		return nil
	}

	return dn.previous
}

// SetPrevious sets previous degree node for the current node.
func (dn *Node) SetPrevious(previousNode *Node) *Node {
	if dn == nil {
		return nil
	}
	dn.previous = previousNode

	return dn
}

// Note returns the note lying on this degree.
func (dn *Node) Note() note.Note {
	if dn == nil {
		return note.Note{}
	}

	return dn.degree.note
}

// SetNote sets note for the degree node.
func (dn *Node) SetNote(n note.Note) {
	if dn != nil {
		dn.degree.note = n
	}
}

// ModalCharacteristics returns modal characteristics of the degree.
func (dn *Node) ModalCharacteristics() ModalCharacteristics {
	if dn == nil {
		return nil
	}

	return dn.degree.modalCharacteristics
}

// SetModalCharacteristics sets modal characteristics to the degree node.
func (dn *Node) SetModalCharacteristics(modalCharacteristics ModalCharacteristics) {
	if dn != nil {
		dn.degree.modalCharacteristics = modalCharacteristics
	}
}

// AbsoluteModalPosition returns absolute modal position of the degree.
func (dn *Node) AbsoluteModalPosition() ModalPosition {
	if dn == nil {
		return ModalPosition{}
	}

	return dn.degree.absoluteModalPosition
}

// SetAbsoluteModalPosition sets absolute modal position to the degree node.
func (dn *Node) SetAbsoluteModalPosition(modalPosition ModalPosition) {
	if dn != nil {
		dn.degree.absoluteModalPosition = modalPosition
	}
}

// GetDegreeByDegreeNum returns the degree node from the chain by its number, if it exists.
func (dn *Node) GetDegreeByDegreeNum(degreeNum Number) *Node {
	if dn == nil {
		return nil
	}

	if dn.Number() == degreeNum {
		return dn
	}

	firstNode := dn.GetLast(true)
	for node := range firstNode.IterateOneRound(false) {
		if node.Number() == degreeNum {
			return node
		}
	}

	return nil
}

// GetForwardDegreeByDegreeNum returns a degree node that is a few positions ahead of the current node.
func (dn *Node) GetForwardDegreeByDegreeNum(forwardDegrees Number) *Node {
	if dn == nil {
		return nil
	}

	currentNode := dn
	for ; currentNode.GetNext() != nil && forwardDegrees != 0; currentNode = currentNode.GetNext() {
		forwardDegrees--
	}

	return currentNode
}

// IterateOneRound iterates through a chain of degree nodes forwards or backwards depending on the argument.
// If the sequence is closed, the last element will be the previous one.
// If it is not closed, it will be the last in the chain.
func (dn *Node) IterateOneRound(left bool) Iterator {
	return func(yield func(*Node) bool) {
		if dn == nil {
			return
		}

		var next func(node *Node) *Node
		if left {
			next = func(node *Node) *Node { return node.GetPrevious() }
		} else {
			next = func(node *Node) *Node { return node.GetNext() }
		}

		current := dn
		if !yield(current) {
			return
		}

		// Use pointer comparison instead of Number comparison for cycle detection
		for next(current) != nil && next(current) != dn {
			current = next(current)
			if !yield(current) {
				return
			}
		}
	}
}

// SortByAbsoluteModalPositions sorts the chain of degree nodes by their absolute modal positions.
func (dn *Node) SortByAbsoluteModalPositions(asc bool) *Node {
	if dn == nil {
		return nil
	}

	// Check for the existence of absolute modal positions
	for node := range dn.IterateOneRound(false) {
		if !node.AbsoluteModalPosition().IsSet() {
			return nil
		}
	}

	// Create a slice to store all degree nodes
	nodes := make([]*Node, 0)
	for node := range dn.IterateOneRound(false) {
		nodes = append(nodes, node.CopyCut())
	}

	// Sort the slice by the weight of the absolute modal position
	sort.Slice(nodes, func(i, j int) bool {
		if asc {
			return nodes[i].AbsoluteModalPosition().Weight() < nodes[j].AbsoluteModalPosition().Weight()
		}
		return nodes[i].AbsoluteModalPosition().Weight() > nodes[j].AbsoluteModalPosition().Weight()
	})

	// Link the sorted nodes
	for i := range len(nodes) - 1 {
		nodes[i].AttachNext(nodes[i+1])
	}
	nodes[len(nodes)-1].AttachNext(nodes[0])

	return nodes[0]
}

// String is stringer for degree node object.
func (dn *Node) String() string {
	if dn == nil {
		return "nil degree node"
	}

	nodeString := fmt.Sprintf("Number: %d, HalfTonesFromPrime: %d, previous exist: %t, next exist: %t",
		dn.degree.number, dn.degree.halfTonesFromPrime, dn.PreviousExists(), dn.NextExists())

	nodeString = fmt.Sprintf("%s, note: %s", nodeString, dn.degree.note.Name())

	if dn.degree.absoluteModalPosition.IsSet() {
		nodeString = fmt.Sprintf("%s, absolute modal position: %s (Weight:%d)", nodeString, dn.degree.absoluteModalPosition.name, dn.degree.absoluteModalPosition.Weight())
	}

	return nodeString
}

// Equal compares degree nodes by degree number and by contained notes.
func (dn *Node) Equal(other *Node) bool {
	if dn == nil || other == nil {
		return false
	}
	if dn.Number() != other.Number() {
		return false
	}

	return dn.Note().EqualByName(other.Note())
}

// EqualByDegreeNum compares degree nodes by degree number only.
func (dn *Node) EqualByDegreeNum(other *Node) bool {
	if dn == nil || other == nil {
		return false
	}

	return dn.Number() == other.Number()
}

// Copy creates full copy of current degree node.
func (dn *Node) Copy() *Node {
	if dn == nil {
		return nil
	}

	return &Node{
		degree:   dn.degree.Copy(),
		previous: dn.previous,
		next:     dn.next,
	}
}

// CopyCut creates copy of current degree node without links to next and previous nodes.
func (dn *Node) CopyCut() *Node {
	if dn == nil {
		return nil
	}

	return &Node{
		degree:   dn.degree.Copy(),
		previous: nil,
		next:     nil,
	}
}

// InsertBetween inserts the degree node between given nodes, e.g. attaches current node after the first and before the second node.
func (dn *Node) InsertBetween(node1, node2 *Node) {
	node1.AttachNext(dn)
	node2.AttachPrevious(dn)
}

// AttachNext adds degree node as next to current node with mutual reference.
func (dn *Node) AttachNext(node *Node) {
	dn.SetNext(node)
	node.SetPrevious(dn)
}

// InsertNext inserts the specified degree node between the current one and the next one, if it exists.
func (dn *Node) InsertNext(node *Node) {
	if dn.NextExists() {
		dn.GetNext().AttachPrevious(node)
	}

	dn.AttachNext(node)
}

// AttachPrevious adds degree node as previous to current node with mutual reference.
func (dn *Node) AttachPrevious(node *Node) {
	dn.SetPrevious(node)
	node.SetNext(dn)
}

// InsertPrevious inserts the specified degree node between the current one and the previous one, if it exists.
func (dn *Node) InsertPrevious(node *Node) {
	if dn.PreviousExists() {
		dn.GetPrevious().AttachNext(node)
	}
	dn.AttachPrevious(node)
}

// NextExists checks for the existence of the next degree node.
func (dn *Node) NextExists() bool {
	return dn.next != nil
}

// PreviousExists checks for the existence of the previous degree node.
func (dn *Node) PreviousExists() bool {
	return dn.previous != nil
}

// GetLast iterates to the specified end (left or right) and returns the last degree node.
// If the chain is cycled, the method returns the last node before the current or next to the current node depending on the specified argument.
// Otherwise, it returns the last node in the chain.
func (dn *Node) GetLast(left bool) *Node {
	if dn == nil {
		return nil
	}

	var next func(node *Node) *Node

	switch left {
	case true:
		if !dn.PreviousExists() {
			return dn
		}
		next = func(node *Node) *Node {
			return node.GetPrevious()
		}
	default:
		if !dn.NextExists() {
			return dn
		}
		next = func(node *Node) *Node {
			return node.GetNext()
		}
	}

	firstNode := dn
	var currentNode *Node

	currentNode = next(firstNode)
	// Use pointer comparison instead of Number comparison for cycle detection
	for next(currentNode) != nil && next(currentNode) != firstNode {
		currentNode = next(currentNode)
	}

	return currentNode
}

// AttachToTheEnd attaches the degree node to the last node in the chain.
// If the chain is cycled, the method will attach the given node
// before the current or next to the current node depending on the specified argument.
// Otherwise, the node will be attached to the last node in the chain by the specified side.
func (dn *Node) AttachToTheEnd(node *Node, left bool) {
	lastNode := dn.GetLast(left)
	switch left {
	case true:
		lastNode.AttachPrevious(node)
	default:
		lastNode.AttachNext(node)
	}
}

// ReverseSequence creates new reversed sequence of degree nodes and returns the first node.
func (dn *Node) ReverseSequence() *Node {
	var firstNodeResult *Node
	lastNode := dn.GetLast(false)

	var iterateFrom *Node
	// in case of cycled sequence - use pointer comparison
	if lastNode.GetNext() != nil && lastNode.GetNext() == dn &&
		dn.GetPrevious() != nil && dn.GetPrevious() == lastNode {
		iterateFrom = dn.GetPrevious()
		// in case of the sequence is not cyclic
	} else if lastNode.GetNext() == nil && dn.GetPrevious() == nil {
		iterateFrom = lastNode
	}

	isFirst := true
	var lastNodeResult *Node
	for node := range iterateFrom.IterateOneRound(true) {
		if isFirst {
			firstNodeResult = node.CopyCut()
			isFirst = false
			lastNodeResult = firstNodeResult

			continue
		}
		lastNodeResult.AttachNext(node.CopyCut())
		lastNodeResult = lastNodeResult.GetNext()
	}
	firstNodeResult.AttachPrevious(lastNodeResult)

	return firstNodeResult
}
