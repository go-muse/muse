package degree

import (
	"fmt"
	"sort"
	"unsafe"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

// Number is a position of a note in a mode, from 1 to last.
type Number uint8

// Degree is a position of a note in a mode that has specific characteristics.
type Degree struct {
	number                Number
	halfTonesFromPrime    halftone.HalfTones
	previous              *Degree
	next                  *Degree
	note                  *note.Note
	modalCharacteristics  ModalCharacteristics // Modal positions of another degrees
	absoluteModalPosition *ModalPosition       // Absolute modal position of this deg
}

// New creates a new degree with specified parameters.
func New(
	number Number,
	halfTonesFromPrime halftone.HalfTones,
	previous, next *Degree,
	note *note.Note,
	modalCharacteristics ModalCharacteristics,
	absoluteModalPosition *ModalPosition,
) *Degree {
	return &Degree{
		number:                number,
		halfTonesFromPrime:    halfTonesFromPrime,
		previous:              previous,
		next:                  next,
		note:                  note,
		modalCharacteristics:  modalCharacteristics,
		absoluteModalPosition: absoluteModalPosition,
	}
}

// Number returns degree's number (it's number in a mode).
func (d *Degree) Number() Number {
	return d.number
}

// HalfTonesFromPrime returns degree's distance from prime in halftone.
func (d *Degree) HalfTonesFromPrime() halftone.HalfTones {
	return d.halfTonesFromPrime
}

// SetHalfTonesFromPrime sets halftone from prime value in halftone.
func (d *Degree) SetHalfTonesFromPrime(halfTones halftone.HalfTones) {
	d.halfTonesFromPrime = halfTones
}

// GetNext returns next degree.
func (d *Degree) GetNext() *Degree {
	if d == nil {
		return nil
	}

	return d.next
}

// SetNext sets next degree for the current degree.
func (d *Degree) SetNext(nextDegree *Degree) *Degree {
	if d == nil {
		return nil
	}

	d.next = nextDegree

	return d
}

// GetPrevious returns previous degree.
func (d *Degree) GetPrevious() *Degree {
	if d == nil {
		return nil
	}

	return d.previous
}

// SetPrevious sets previous degree for the current degree.
func (d *Degree) SetPrevious(previousDegree *Degree) *Degree {
	if d == nil {
		return nil
	}

	d.previous = previousDegree

	return d
}

// Note returns a pointer to the note lying on this degree.
func (d *Degree) Note() *note.Note {
	if d == nil {
		return nil
	}

	return d.note
}

// SetNote sets note for the current degree.
func (d *Degree) SetNote(note *note.Note) {
	if d != nil {
		d.note = note
	}
}

// ModalCharacteristics returns modal characteristics of the degree.
func (d *Degree) ModalCharacteristics() ModalCharacteristics {
	return d.modalCharacteristics
}

// SetModalCharacteristics sets modal characteristics to the degree.
func (d *Degree) SetModalCharacteristics(modalCharacteristics ModalCharacteristics) {
	d.modalCharacteristics = modalCharacteristics
}

// AbsoluteModalPosition returns absolute modal position of the degree.
func (d *Degree) AbsoluteModalPosition() *ModalPosition {
	if d == nil {
		return nil
	}

	return d.absoluteModalPosition
}

// SetAbsoluteModalPosition sets absolute modal position to the degree.
func (d *Degree) SetAbsoluteModalPosition(modalPosition *ModalPosition) {
	d.absoluteModalPosition = modalPosition
}

// GetDegreeByDegreeNum returns the degree from the chain of degrees by its number, if it exists.
func (d *Degree) GetDegreeByDegreeNum(degreeNum Number) *Degree {
	if d == nil {
		return nil
	}

	if d.Number() == degreeNum {
		return d
	}

	firstDegree := d.GetLast(true)
	for degree := range firstDegree.IterateOneRound(false) {
		if degree.Number() == degreeNum {
			return degree
		}
	}

	return nil
}

// GetForwardDegreeByDegreeNum returns a degree that is a few degrees ahead of the current degree.
func (d *Degree) GetForwardDegreeByDegreeNum(forwardDegrees Number) *Degree {
	if d == nil {
		return nil
	}

	currentDegree := d
	for ; currentDegree.GetNext() != nil && forwardDegrees != 0; currentDegree = currentDegree.GetNext() {
		forwardDegrees--
	}

	return currentDegree
}

// IterateOneRound iterates through a chain of steps forwards or backwards depending on the argument.
// If the sequence is closed, the last element will be the previous one.
// If it is not closed, it will be the last in the chain.
func (d *Degree) IterateOneRound(left bool) Iterator {
	type nextFunc func(d *Degree) *Degree
	var next nextFunc
	switch left {
	case true:
		next = func(d *Degree) *Degree { return d.GetPrevious() }
	default:
		next = func(d *Degree) *Degree { return d.GetNext() }
	}
	current := d

	c := make(chan *Degree)

	go func(ch chan *Degree, d *Degree, next nextFunc) {
		defer close(ch)
		if d == nil {
			return
		}
		ch <- d
		for next(current) != nil && unsafe.Pointer(next(current)) != unsafe.Pointer(d) {
			current = next(current)
			ch <- current
		}
	}(c, d, next)

	return c
}

// SortByAbsoluteModalPositions sorts the chain of degrees by their absolute modal positions.
func (d *Degree) SortByAbsoluteModalPositions(asc bool) *Degree {
	if d == nil {
		return nil
	}

	// Check for the existence of absolute modal positions
	for degree := range d.IterateOneRound(false) {
		if degree.AbsoluteModalPosition() == nil {
			return nil
		}
	}

	// Create a slice to store all degrees
	degrees := make([]*Degree, 0)
	for degree := range d.IterateOneRound(false) {
		degrees = append(degrees, degree.CopyCut())
	}

	// Sort the slice by the weight of the absolute modal position
	sort.Slice(degrees, func(i, j int) bool {
		if asc {
			return degrees[i].AbsoluteModalPosition().Weight() < degrees[j].AbsoluteModalPosition().Weight()
		}

		return degrees[i].AbsoluteModalPosition().Weight() > degrees[j].AbsoluteModalPosition().Weight()
	})

	// Link the sorted degrees
	for i := range len(degrees) - 1 {
		degrees[i].AttachNext(degrees[i+1])
	}
	degrees[len(degrees)-1].AttachNext(degrees[0])

	return degrees[0]
}

// String is stringer for degree object.
func (d *Degree) String() string {
	if d == nil {
		return "nil degree"
	}

	degreeString := fmt.Sprintf("Number: %d, HalfTonesFromPrime: %d, previous exist: %t, next exist: %t",
		d.Number(), d.halfTonesFromPrime, d.PreviousExists(), d.NextExists())

	if d.note != nil {
		degreeString = fmt.Sprintf("%s, note: %s", degreeString, d.note.Name())
	}

	if d.absoluteModalPosition != nil {
		degreeString = fmt.Sprintf("%s, absolute modal position: %s (Weight:%d)", degreeString, d.absoluteModalPosition.name, d.absoluteModalPosition.Weight())
	}

	return degreeString
}

// IsEqual compares degrees by degree number and by contained notes.
func (d *Degree) IsEqual(degree *Degree) bool {
	if d == nil || degree == nil {
		return false
	}
	if d.Number() != degree.Number() {
		return false
	}
	if (d.note == nil) != (degree.note == nil) {
		return false
	}
	if (d.note != nil) && (degree.note != nil) {
		if !d.note.IsEqualByName(degree.note) {
			return false
		}
	}

	return true
}

// EqualByDegreeNum compares degrees by degree number only.
func (d *Degree) EqualByDegreeNum(degree *Degree) bool {
	if d == nil || degree == nil {
		return false
	}

	return d.Number() == degree.Number()
}

// Copy creates full copy of current degree.
// The method returns a pointer to a new Degree containing the same attribute values as the original Degree that the function was called on.
func (d *Degree) Copy() *Degree {
	return &Degree{
		number:                d.Number(),
		halfTonesFromPrime:    d.halfTonesFromPrime,
		previous:              d.previous,
		next:                  d.next,
		note:                  d.note.Copy(),
		modalCharacteristics:  d.modalCharacteristics.Copy(),
		absoluteModalPosition: d.absoluteModalPosition.Copy(),
	}
}

// CopyCut creates copy of current degree without links to next and previous degrees.
func (d *Degree) CopyCut() *Degree {
	return &Degree{
		number:                d.Number(),
		halfTonesFromPrime:    d.halfTonesFromPrime,
		previous:              nil,
		next:                  nil,
		note:                  d.note.Copy(),
		modalCharacteristics:  d.modalCharacteristics.Copy(),
		absoluteModalPosition: d.absoluteModalPosition.Copy(),
	}
}

// InsertBetween inserts the degree between given degrees, e.g. attaches current degree after the first and before the second degree.
func (d *Degree) InsertBetween(degree1, degree2 *Degree) {
	degree1.AttachNext(d)
	degree2.AttachPrevious(d)
}

// AttachNext adds degree as next to current degree with mutual reference.
func (d *Degree) AttachNext(degree *Degree) {
	d.SetNext(degree)
	degree.SetPrevious(d)
}

// InsertNext inserts the specified degree between the current one and the next one, if it exists.
func (d *Degree) InsertNext(degree *Degree) {
	if d.NextExists() {
		d.GetNext().AttachPrevious(degree)
	}

	d.AttachNext(degree)
}

// AttachPrevious adds degree as previous to current degree with mutual reference.
func (d *Degree) AttachPrevious(degree *Degree) {
	d.SetPrevious(degree)
	degree.SetNext(d)
}

// InsertPrevious inserts the specified degree between the current one and the previous one, if it exists.
func (d *Degree) InsertPrevious(degree *Degree) {
	if d.PreviousExists() {
		d.GetPrevious().AttachNext(degree)
	}

	d.AttachPrevious(degree)
}

// NextExists checks for the existence of the next degree.
func (d *Degree) NextExists() bool {
	return d.next != nil
}

// PreviousExists checks for the existence of the previous degree.
func (d *Degree) PreviousExists() bool {
	return d.previous != nil
}

// GetLast iterates to the specified end (left or right) and returns the last degree.
// If degrees chain is cycled, the method returns the last degree before the current or next to the current degree depending on specified argument.
// Otherwise, in returns the last degree in the chain.
func (d *Degree) GetLast(left bool) *Degree {
	if d == nil {
		return nil
	}

	var next func(degree *Degree) *Degree

	switch left {
	case true:
		if !d.PreviousExists() {
			return d
		}
		next = func(degree *Degree) *Degree {
			return degree.GetPrevious()
		}
	default:
		if !d.NextExists() {
			return d
		}
		next = func(degree *Degree) *Degree {
			return degree.GetNext()
		}
	}

	firstDegree := d
	var currentDegree *Degree

	currentDegree = next(firstDegree)
	for next(currentDegree) != nil && unsafe.Pointer(next(currentDegree)) != unsafe.Pointer(firstDegree) {
		currentDegree = next(currentDegree)
	}

	return currentDegree
}

// AttachToTheEnd attaches the degree to the last degree in the chain.
// If degrees chain is cycled, the method will attach given degree
// before the current or next to the current degree depending on specified argument.
// Otherwise, the degree will be attached to the last degree in the chain by the specified side.
func (d *Degree) AttachToTheEnd(degree *Degree, left bool) {
	lastDegree := d.GetLast(left)
	switch left {
	case true:
		lastDegree.AttachPrevious(degree)
	default:
		lastDegree.AttachNext(degree)
	}
}

// ReverseSequence creates new reversed sequence of degrees and returns the first degree.
func (d *Degree) ReverseSequence() *Degree {
	var firstDegreeResult *Degree
	lastDegree := d.GetLast(false)

	var iterateFrom *Degree
	// in case of cycled sequence
	if unsafe.Pointer(d) == unsafe.Pointer(lastDegree.GetNext()) &&
		unsafe.Pointer(d.GetPrevious()) == unsafe.Pointer(lastDegree) {
		iterateFrom = d.GetPrevious()
		// in case of the sequence is not cyclic
	} else if lastDegree.GetNext() == nil && d.GetPrevious() == nil {
		iterateFrom = lastDegree
	}

	isFirst := true
	var lastDegreeResult *Degree
	for degree := range iterateFrom.IterateOneRound(true) {
		if isFirst {
			firstDegreeResult = degree.CopyCut()
			isFirst = false
			lastDegreeResult = firstDegreeResult

			continue
		}
		lastDegreeResult.AttachNext(degree.CopyCut())
		lastDegreeResult = lastDegreeResult.GetNext()
	}
	firstDegreeResult.AttachPrevious(lastDegreeResult)

	return firstDegreeResult
}
