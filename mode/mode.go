package mode

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/interval"
	"github.com/go-muse/muse/note"
)

// Mode is a set of degrees located at certain intervals from each other within octave.
// Degrees are circular linked list.
type Mode struct {
	name   Name           // Mode Name.
	degree *degree.Degree // The first degree in the mode (tonal center). The mode always points to the first degree.
}

// Name returns mode's name.
func (m *Mode) Name() Name {
	return m.name
}

// MakeNewMode creates mode based on a mode template and the first note.
func MakeNewMode(modeName Name, firstNoteName note.Name) (*Mode, error) {
	firstNote, err := note.New(firstNoteName)
	if err != nil {
		return nil, fmt.Errorf("create first note by name = '%s': %w", firstNoteName, err)
	}

	modeTemplate, err := GetTemplateByName(modeName)
	if err != nil {
		return nil, fmt.Errorf("get template to create mode by mode name '%s': %w", modeName, err)
	}

	err = modeTemplate.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate mode template to create mode: %w", err)
	}

	// Mode building
	mode := newModeBuilder(modeTemplate).
		build(modeName, firstNote)

	return mode, nil
}

// MustMakeNewMode creates a mode as MakeNewMode does, if you are confident in the correctness of the mode name and notes.
func MustMakeNewMode(modeName Name, firstNoteName note.Name) *Mode {
	mode, err := MakeNewMode(modeName, firstNoteName)
	if err != nil {
		panic(err)
	}

	return mode
}

// MakeNewCustomMode makes custom mode with intervals described in mode template.
func MakeNewCustomMode(modeTemplate Template, firstNoteName string, modeName Name) (*Mode, error) {
	if err := modeTemplate.Validate(); err != nil {
		return nil, fmt.Errorf("validate mode template to create custom mode: %w", err)
	}

	firstNote, err := note.NewNoteFromString(firstNoteName)
	if err != nil {
		return nil, fmt.Errorf("make first note to create custom mode by firstNoteName '%s': %w", firstNoteName, err)
	}

	mode := newModeBuilder(modeTemplate).
		build(modeName, firstNote)

	return mode, nil
}

var ErrDegreeNumberInvalid = errors.New("invalid degree number")

// MakeNewCustomModeWithDegree makes custom mode with given degrees chain.
func MakeNewCustomModeWithDegree(modeName Name, firstDegree *degree.Degree) (*Mode, error) {
	if firstDegree.Number() != degree.Number(1) {
		return nil, fmt.Errorf("create mode with first degree number '%d'. First degree must be '1': %w", firstDegree.Number(), ErrDegreeNumberInvalid)
	}

	return &Mode{modeName, firstDegree}, nil
}

// MustMakeNewCustomModeWithDegree makes custom mode with given degrees chain, panics in case of error.
func MustMakeNewCustomModeWithDegree(modeName Name, firstDegree *degree.Degree) *Mode {
	mode, err := MakeNewCustomModeWithDegree(modeName, firstDegree)
	if err != nil {
		panic(err)
	}

	return mode
}

// Length returns length of the mode (amount of degrees).
func (m *Mode) Length() degree.Number {
	var length degree.Number

	for range m.GetFirstDegree().IterateOneRound(false) {
		length++
	}

	return length
}

// GetFirstDegree returns pointer to the first degree in the mode.
func (m *Mode) GetFirstDegree() *degree.Degree {
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
func (m *Mode) GetLastDegree() *degree.Degree {
	if m == nil || m.degree == nil {
		return nil
	}

	// This will work for both circular and not circular chain of degrees
	currentDegree := m.degree
	firstDegree := m.degree
	for currentDegree.GetNext() != nil && unsafe.Pointer(currentDegree.GetNext()) != unsafe.Pointer(firstDegree) {
		currentDegree = currentDegree.GetNext()
	}

	return currentDegree
}

// GetDegreeByDegreeNum returns degree by degree number.
func (m *Mode) GetDegreeByDegreeNum(degreeNum degree.Number) *degree.Degree {
	if m == nil || m.degree == nil {
		return nil // or error?
	}

	if degreeNum > m.Length() {
		return nil // or error?
	}

	return m.degree.GetDegreeByDegreeNum(degreeNum)
}

// GetNoteByDegreeNum returns note by degree number.
func (m *Mode) GetNoteByDegreeNum(degreeNum degree.Number) *note.Note {
	degree := m.GetDegreeByDegreeNum(degreeNum)
	if degree != nil {
		return degree.Note()
	}

	return nil
}

// InsertNote adds note as degree in the mode during mode building.
// This will work for both closed and open degrees circle in the mode.
func (m *Mode) InsertNote(note *note.Note, halfTonesFromPrime halftone.HalfTones) {
	if m == nil {
		return
	}

	// The first note in the mode
	if m.degree == nil {
		m.degree = degree.New(
			1,
			halfTonesFromPrime,
			nil,
			nil,
			note,
			nil,
			nil,
		)

		return
	}

	newDegree := degree.New(
		m.GetLastDegree().Number()+1,
		halfTonesFromPrime,
		m.GetLastDegree(),
		nil,
		note,
		nil,
		nil,
	)

	// in case of closed degree's circle
	if m.IsClosedCircleOfDegrees() {
		newDegree.SetNext(m.GetFirstDegree())
		m.GetFirstDegree().SetPrevious(newDegree)
	}

	m.GetLastDegree().SetNext(newDegree)
}

// IterateOneRound returns channel with degrees of the mode to iterate through them.
func (m *Mode) IterateOneRound(left bool) degree.Iterator {
	return m.degree.IterateOneRound(left)
}

// IsEqual compares modes.
func (m *Mode) IsEqual(mode *Mode) bool {
	if m == nil || mode == nil {
		return false
	}

	if m.name != mode.name || m.Length() != mode.Length() {
		return false
	}
	d1chan := m.degree.IterateOneRound(false)
	for d2 := range mode.degree.IterateOneRound(false) {
		if !(<-d1chan).IsEqual(d2) {
			return false
		}
	}

	return true
}

// GetIntervals returns all intervals from the mode.
// It is possible to specify filtering options.
func (m *Mode) GetIntervals(opts *interval.FilteringOptions) []interval.Diatonic {
	intervalsWithDegrees := make([]interval.Diatonic, 0)
	for degree := range m.GetFirstDegree().IterateOneRound(false) {
		if !opts.HasFilterByAbsoluteModalPosition(degree.AbsoluteModalPosition().Name()) {
			for _, mc := range degree.ModalCharacteristics() {
				if !opts.HasFilterByDegreeCharacteristicName(mc.Name()) {
					intervalWithDegrees, _ := interval.NewDiatonic(degree, mc.Degree())
					if !opts.HasFilterBySonance(intervalWithDegrees.Chromatic().Sonance) {
						intervalsWithDegrees = append(intervalsWithDegrees, *intervalWithDegrees)
					}
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

// OpenCircleOfDegrees makes linked list of degrees not circular.
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

	degree := firstDegree.SortByAbsoluteModalPositions(true)
	if degree != nil {
		m.degree = degree
	}

	if asc {
		m.degree = m.degree.ReverseSequence()
	}
}

// Contains checks for the presence of the specified note in the current mode.
func (m *Mode) Contains(note *note.Note) bool {
	for degree := range m.IterateOneRound(false) {
		if degree.Note().IsEqualByName(note) {
			return true
		}
	}

	return false
}

// setAbsoluteModalPositions calculates absolute modal positions of all the degrees in the mode and stores them.
func (m *Mode) setAbsoluteModalPositions() {
	if m == nil {
		return
	}
	for currentDegree := range m.GetFirstDegree().IterateOneRound(false) {
		var weightSum degree.Weight
		for _, mc := range currentDegree.ModalCharacteristics() {
			weightSum += mc.RelativeModalPosition().Weight()
		}

		currentDegree.SetAbsoluteModalPosition(degree.NewModalPositionByWeight(weightSum))
	}
}

func (m *Mode) setRelativeModalPositions(modeTemplate Template) {
	for currentDegree := range m.GetFirstDegree().IterateOneRound(false) {
		setRelativeMCsOfDegree(currentDegree, modeTemplate)
	}
}

// setRelativeMCsOfDegree calculates relative modal characteristics of the degree and stores them in it.
func setRelativeMCsOfDegree(d *degree.Degree, modeTemplate Template) {
	mcs := make(degree.ModalCharacteristics, 0, modeTemplate.Length())

	// Rebuild mode template from the given degree and iterate through it
	for iteratorResult := range modeTemplate.RearrangeFromDegree(d.Number()).IterateOneRound(true) {
		degreeNum, _, halfTonesFromPrime := iteratorResult()
		nextDegree := d.GetForwardDegreeByDegreeNum(degreeNum - 1)

		// Calculation of relative modal characteristics of the degree
		mc := degree.MustCalculateRelativeMC(degreeNum, nextDegree, halfTonesFromPrime)
		mcs = mcs.Add(*mc)
	}

	d.SetModalCharacteristics(mcs)
}
