package muse

import (
	"unsafe"

	"github.com/pkg/errors"
)

// Mode is a set of degrees located at certain intervals from each other within octave.
// Degrees are circular linked list.
type Mode struct {
	name   ModeName // Mode Name.
	degree *Degree  // The first degree in the mode (tonal center). The mode always points to the first degree.
}

// Name returns mode's name.
func (m *Mode) Name() ModeName {
	return m.name
}

// A mode is usually created based on a mode template and the first note.
func MakeNewMode(modeName ModeName, firstNoteName NoteName) (*Mode, error) {
	firstNote, err := NewNote(firstNoteName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to make first note to create mode by firstNoteName = %s", firstNoteName)
	}

	modeTemplate, err := GetTemplateByModeName(modeName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get template to create mode by mode name = %s", modeName)
	}

	err = modeTemplate.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate mode template to create mode")
	}

	// Mode building
	mode := newModeBuilder(modeTemplate).build(modeName, modeTemplate, firstNote)

	return mode, nil
}

// MustMakeNewMode creates a mode as MakeNewMode does, if you are confident in the correctness of the mode name and notes.
func MustMakeNewMode(modeName ModeName, firstNoteName NoteName) *Mode {
	mode, err := MakeNewMode(modeName, firstNoteName)
	if err != nil {
		panic(err)
	}

	return mode
}

// MakeNewCustomMode makes custom mode with intervals described in mode template.
func MakeNewCustomMode(modeTemplate ModeTemplate, firstNoteName string, modeName ModeName) (*Mode, error) {
	if err := modeTemplate.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate mode template to create custom mode")
	}

	firstNote, err := NewNoteFromString(firstNoteName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to make first note to create custom mode by firstNoteName = %s", firstNoteName)
	}

	mode := newModeBuilder(modeTemplate).build(modeName, modeTemplate, firstNote)

	return mode, nil
}

var ErrDegreeNumberInvalid = errors.New("invalid degree number")

// MakeNewCustomModeWithDegree makes custom mode with given degrees chain.
func MakeNewCustomModeWithDegree(modeName ModeName, firstDegree *Degree) (*Mode, error) {
	if firstDegree.Number() != DegreeNum(1) {
		return nil, errors.Wrapf(ErrDegreeNumberInvalid, "failed to create mode with first degree number: %d. First degree must be 1", firstDegree.Number())
	}

	return &Mode{modeName, firstDegree}, nil
}

// Length returns length of the mode (amount of degrees).
func (m *Mode) Length() DegreeNum {
	var length DegreeNum

	for range m.GetFirstDegree().IterateOneRound(false) {
		length++
	}

	return length
}

// GetFirstDegree returns pointer to the first degree in the mode.
func (m *Mode) GetFirstDegree() *Degree {
	if m == nil || m.degree == nil {
		return nil
	}

	if m.degree.Number() == 1 {
		return m.degree
	}

	// TODO: error?

	return nil
}

// GetLastDegree returns pointer to the last degree in the mode.
func (m *Mode) GetLastDegree() *Degree {
	if m == nil || m.degree == nil {
		return nil
	}

	// This will work for both circular and not circular chain of degrees
	currentDegree := m.degree
	firstDegree := m.degree
	for currentDegree.next != nil && unsafe.Pointer(currentDegree.next) != unsafe.Pointer(firstDegree) {
		currentDegree = currentDegree.next
	}

	return currentDegree
}

// GetDegreeByDegreeNum returns degree by degree number.
func (m *Mode) GetDegreeByDegreeNum(degreeNum DegreeNum) *Degree {
	if m == nil || m.degree == nil {
		return nil // or error?
	}

	if degreeNum > m.Length() {
		return nil // or error?
	}

	return m.degree.GetDegreeByDegreeNum(degreeNum)
}

// GetNoteByDegreeNum returns note by degree number.
func (m *Mode) GetNoteByDegreeNum(degreeNum DegreeNum) *Note {
	degree := m.GetDegreeByDegreeNum(degreeNum)
	if degree != nil {
		return degree.note
	}

	return nil
}

// InsertNote adds note as degree in the mode during mode building.
// This will work for both closed and open degrees circle in the mode.
func (m *Mode) InsertNote(note *Note, halfTonesFromPrime HalfTones) {
	if m == nil {
		return
	}

	// The first note in the mode
	if m.degree == nil {
		m.degree = &Degree{
			number:             1,
			halfTonesFromPrime: 0,
			previous:           nil,
			next:               nil,
			note:               note,
		}

		return
	}

	newDegree := &Degree{
		number:             m.GetLastDegree().Number() + 1,
		halfTonesFromPrime: halfTonesFromPrime,
		previous:           m.GetLastDegree(),
		next:               nil,
		note:               note,
	}

	// in case of closed degree's circle
	if m.IsClosedCircleOfDegrees() {
		newDegree.SetNext(m.GetFirstDegree())
		m.GetFirstDegree().SetPrevious(newDegree)
	}

	m.GetLastDegree().SetNext(newDegree)
}

// IterateOneRound returns channel with degrees of the mode to iterate through them.
func (m *Mode) IterateOneRound(left bool) DegreesIterator {
	return m.degree.IterateOneRound(left)
}

// Equal compares modes.
func (m *Mode) Equal(mode *Mode) bool {
	if m == nil || mode == nil {
		return false
	}

	if m.name != mode.name || m.Length() != mode.Length() {
		return false
	}

	d1chan := m.degree.IterateOneRound(false)
	for d2 := range m.degree.IterateOneRound(false) {
		if !(<-d1chan).Equal(d2) {
			return false
		}
	}

	return true
}

// GetIntervals returns all intervals from the mode.
// It is possible to specify filtering options.
func (m *Mode) GetIntervals(opts *IntervalFilteringOptions) []Interval {
	intervalsWithDegrees := make([]Interval, 0)
	for degree := range m.GetFirstDegree().IterateOneRound(false) {
		if !opts.FilterByAbsoluteModalPositionExist(degree.absoluteModalPosition.name) {
			for _, mc := range degree.modalCharacteristics {
				if !opts.FilterByDegreeCharacteristicNameExist(mc.name) && !opts.FilterBySonanceExist(mc.interval.Sonance) {
					intervalWithDegrees := Interval{
						ChromaticInterval: mc.interval.GetInterval(),
						degree1:           degree,
						degree2:           mc.degree,
					}
					intervalsWithDegrees = append(intervalsWithDegrees, intervalWithDegrees)
				}
			}
		}
	}

	return intervalsWithDegrees
}

// CloseCircleOfDegrees makes linked list of degrees circular.
func (m *Mode) CloseCircleOfDegrees() {
	if m.GetFirstDegree().GetPrevious() == nil && m.GetLastDegree().GetNext() == nil {
		m.GetFirstDegree().SetPrevious(m.GetLastDegree())
		m.GetLastDegree().SetNext(m.GetFirstDegree())
	}
}

// OpenTheCircleOfDegrees makes linked list of degrees not circular.
func (m *Mode) OpenCircleOfDegrees() {
	if m.GetFirstDegree().GetPrevious() != nil && m.GetLastDegree().GetNext() != nil {
		m.GetFirstDegree().SetPrevious(nil)
		m.GetLastDegree().SetNext(nil)
	}
}

// IsClosedCircleOfDegrees checks if the circle of degrees is closed or not.
func (m *Mode) IsClosedCircleOfDegrees() bool {
	if m.GetFirstDegree().GetPrevious() != nil &&
		unsafe.Pointer(m.GetFirstDegree().GetPrevious()) == unsafe.Pointer(m.GetLastDegree()) &&
		m.GetLastDegree().GetNext() != nil && unsafe.Pointer(m.GetLastDegree().GetNext()) == unsafe.Pointer(m.GetFirstDegree()) {
		return true
	}

	return false
}

// SortByAbsoluteModalPositions sorts the chain of mode's degrees by absolute modal positions.
// The order of sorting (ascending or descending) can be specified by the argument.
func (m *Mode) SortByAbsoluteModalPositions(asc bool) {
	firstDegree := m.GetFirstDegree()
	if firstDegree == nil {
		return
	}

	degree := firstDegree.sortByAbsoluteModalPositions()
	if degree != nil {
		m.degree = degree
	}

	if asc {
		m.degree = m.degree.ReverseSequence()
	}
}
