package muse

import (
	"fmt"
	"unsafe"
)

// Degree number is a position of a note in a mode, from 1 to last.
type DegreeNum uint8

// Degree is a position of a note in a mode that has specific characteristics.
type Degree struct {
	number                DegreeNum
	halfTonesFromPrime    HalfTones
	previous              *Degree
	next                  *Degree
	note                  *Note
	modalCharacteristics  ModalCharacteristics // Modal positions of another degrees
	absoluteModalPosition *ModalPosition       // Absolute modal position of this degree
}

// NewDegree creates a new degree with specified parameters.
func NewDegree(
	number DegreeNum,
	halfTonesFromPrime HalfTones,
	previous, next *Degree,
	note *Note,
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

// GetNext returns degree's number (it's number in a mode).
func (d *Degree) Number() DegreeNum {
	return d.number
}

// HalfTonesFromPrime returns degree's distance from prime in half rones.
func (d *Degree) HalfTonesFromPrime() HalfTones {
	return d.halfTonesFromPrime
}

// GetNext returns next degree.
func (d *Degree) GetNext() *Degree {
	if d == nil {
		return nil
	}

	return d.next
}

// SetNext sets next degree for the current degree.
func (d *Degree) SetNext(nextDegree *Degree) {
	d.next = nextDegree
}

// GetPrevious returns previous degree.
func (d *Degree) GetPrevious() *Degree {
	if d == nil {
		return nil
	}

	return d.previous
}

// SetPrevious sets previous degree for the current degree.
func (d *Degree) SetPrevious(previousDegree *Degree) {
	if d != nil {
		d.previous = previousDegree
	}
}

// Note returns a pointer to the note lying on this degree.
func (d *Degree) Note() *Note {
	if d == nil {
		return nil
	}

	return d.note
}

// Note sets note for the current degree.
func (d *Degree) SetNote(note *Note) {
	if d != nil {
		d.note = note
	}
}

// ModalCharacteristics returns modal characteristics of the note.
func (d *Degree) ModalCharacteristics() ModalCharacteristics {
	return d.modalCharacteristics
}

// ModalCharacteristics returns absolute modal position of the note.
func (d *Degree) AbsoluteModalPosition() *ModalPosition {
	return d.absoluteModalPosition
}

// getDegreeByDegreeNum returns the degree from the chain of degrees by its number, if it exists.
func (d *Degree) getDegreeByDegreeNum(degreeNum DegreeNum) *Degree {
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
func (d *Degree) GetForwardDegreeByDegreeNum(forwardDegrees DegreeNum) *Degree {
	if d == nil {
		return nil
	}

	currentDegree := d
	for ; currentDegree.GetNext() != nil && forwardDegrees != 0; currentDegree = currentDegree.GetNext() {
		forwardDegrees--
	}

	return currentDegree
}

// DegreesIterator is an object that allows iterating through a sequence of degrees
// and also provides additional functionality.
type DegreesIterator <-chan *Degree

// GetAllDegrees iterates through a sequence of degrees
// and returns them as slice.
func (di DegreesIterator) GetAllDegrees() []*Degree {
	if di == nil {
		return nil
	}

	var degrees []*Degree
	for {
		degree, ok := <-di
		if !ok {
			return degrees
		}
		degrees = append(degrees, degree)
	}
}

// GetAllNotes iterates through a sequence of degrees
// and returns their notes as slice.
func (di DegreesIterator) GetAllNotes() Notes {
	if di == nil {
		return nil
	}

	notes := make(Notes, 0)
	for {
		degree, ok := <-di
		if !ok {
			return notes
		}
		if degree.Note() != nil {
			notes = append(notes, *degree.Note())
		}
	}
}

// IterateOneRound iterates through a chain of steps forwards or backwards depending on the argument.
// If the sequence is closed, the last element will be the previous one.
// If it is not closed, it will be the last in the chain.
func (d *Degree) IterateOneRound(left bool) DegreesIterator {
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

// sortByAbsoluteModalPositions sorts the chain of degrees by absolute modal positions.
func (d *Degree) sortByAbsoluteModalPositions() *Degree {
	// validation of degrees chain
	for degree := range d.IterateOneRound(false) {
		if degree.absoluteModalPosition == nil {
			return nil
		}
	}

	firstElement := d.CopyCut()
	isFirst := true
	for degree := range d.IterateOneRound(false) {
		if isFirst {
			isFirst = false

			continue
		}

		unsortedDegree := degree.CopyCut()

		for sortedDegree := range firstElement.IterateOneRound(false) {
			if unsortedDegree.absoluteModalPosition.GetWeight() > sortedDegree.absoluteModalPosition.GetWeight() { //nolint:nestif
				if sortedDegree.NextExists() {
					if unsortedDegree.absoluteModalPosition.GetWeight() > sortedDegree.GetNext().absoluteModalPosition.GetWeight() {
						continue
					}
					sortedDegree.GetNext().AttachPrevious(unsortedDegree)
				}
				sortedDegree.AttachNext(unsortedDegree)

				break
			} else { //nolint:revive
				if sortedDegree.PreviousExists() {
					sortedDegree.GetPrevious().AttachNext(unsortedDegree)
				} else {
					firstElement = unsortedDegree
				}
				sortedDegree.AttachPrevious(unsortedDegree)

				break
			}
		}
	}

	firstElement = firstElement.GetLast(true)
	firstElement.AttachToTheEnd(firstElement, false)

	return firstElement
}

// String is stringer for degree object.
func (d *Degree) String() string {
	if d == nil {
		return "nil degree"
	}

	degreeString := fmt.Sprintf("DegreeNum: %d, HalfTonesFromPrime: %d, previous exist: %t, next exist: %t",
		d.Number(), d.halfTonesFromPrime, d.PreviousExists(), d.NextExists())

	if d.note != nil {
		degreeString = fmt.Sprintf("%s, note: %s", degreeString, d.note.Name())
	}

	if d.absoluteModalPosition != nil {
		degreeString = fmt.Sprintf("%s, absolute modal position: %s (weight:%d)", degreeString, d.absoluteModalPosition.name, d.absoluteModalPosition.GetWeight())
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

// Equal compares degrees by degree number only.
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
// If degrees chain is cycled, the method will attach given degree before the current or next to the current degree depending on specified argument.
// Otherwise, the degre will be attached to the last degree in the chain by the specified side.
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
