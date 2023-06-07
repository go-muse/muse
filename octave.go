package muse

import (
	"github.com/pkg/errors"
)

const (
	NotesInOctave = 12
)

// Octave is pair of number and name.
type Octave struct {
	number OctaveNumber
	name   OctaveName
}

const (
	maxOctaveNumber = 9
	minOctaveNumber = -1
)

// NewOctave Creates a new octave by the given octave number.
func NewOctave(octaveNumber OctaveNumber) (*Octave, error) {
	octaveName, err := GetOctaveNameByNumber(octaveNumber)
	if err != nil {
		return nil, err
	}

	return &Octave{
		number: octaveNumber,
		name:   octaveName,
	}, nil
}

// Validate checks octave number.
func (o Octave) Validate() error {
	if o.number < minOctaveNumber || o.number > maxOctaveNumber {
		return errors.Wrapf(ErrOctaveNumberUnknown, "invalid octave number: %d", o.number)
	}

	return nil
}

// MustNewOctave makes octave by the given number with panic on validation.
func MustNewOctave(octaveNumber OctaveNumber) *Octave {
	octave, err := NewOctave(octaveNumber)
	if err != nil {
		panic(err)
	}

	return octave
}

// IsEqual compares the octaves by their numbers.
func (o *Octave) IsEqual(octave *Octave) bool {
	if o == nil || octave == nil {
		return false
	}

	if o.number == octave.number {
		return true
	}

	return false
}

// Name returns the name of the octave.
func (o *Octave) Name() OctaveName {
	if o != nil {
		return o.name
	}

	return ""
}

// Number returns the number of the octave.
func (o *Octave) Number() OctaveNumber {
	return o.number
}

// SetToNote sets the octave to the specified note and returns it.
func (o *Octave) SetToNote(note *Note) *Note {
	if note != nil {
		note.octave = o
	}

	return note
}
