package octave

import (
	"fmt"
)

const (
	NotesInOctave = uint8(12)

	MaxOctaveNumber = Number(9)
	MinOctaveNumber = Number(-1)
)

// Octave is a pair of number and name.
type Octave struct {
	number Number
	name   Name
}

// New Creates a new octave by the given octave number.
func New(octaveName Name) (*Octave, error) {
	octaveNumber, err := GetOctaveNumberByName(octaveName)
	if err != nil {
		return nil, err
	}

	return &Octave{
		number: octaveNumber,
		name:   octaveName,
	}, nil
}

// MustNew makes octave by the given name with panic on error.
func MustNew(octaveName Name) *Octave {
	octave, err := New(octaveName)
	if err != nil {
		panic(err)
	}

	return octave
}

// NewByNumber Creates a new octave by the given octave number.
func NewByNumber(octaveNumber Number) (*Octave, error) {
	octaveName, err := GetOctaveNameByNumber(octaveNumber)
	if err != nil {
		return nil, err
	}

	return &Octave{
		number: octaveNumber,
		name:   octaveName,
	}, nil
}

// MustNewByNumber makes octave by the given number with panic on validation.
func MustNewByNumber(octaveNumber Number) *Octave {
	octave, err := NewByNumber(octaveNumber)
	if err != nil {
		panic(err)
	}

	return octave
}

// Validate checks octave number.
func (o *Octave) Validate() error {
	if o.number < MinOctaveNumber || o.number > MaxOctaveNumber {
		return fmt.Errorf("invalid octave number '%d': %w", o.number, ErrOctaveNumberUnknown)
	}

	return nil
}

// IsEqual compares the octaves by their numbers and names.
func (o *Octave) IsEqual(octave *Octave) bool {
	if o == nil || octave == nil {
		return false
	}

	if o.number == octave.number && o.name == octave.name {
		return true
	}

	return false
}

// Name returns the name of the octave.
func (o *Octave) Name() Name {
	if o != nil {
		return o.name
	}

	return ""
}

// Number returns the number of the octave.
func (o *Octave) Number() Number {
	return o.number
}
