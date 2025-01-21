package octave

import (
	"errors"
)

// Number is a number of octave.
type Number int8

const (
	NumberMinus1 = Number(-1) // Octave name is SubSubContraOctave
	Number0      = Number(0)  // Octave name is SubContraOctave
	Number1      = Number(1)  // Octave name is ContraOctave
	Number2      = Number(2)  // Octave name is GreatOctave
	Number3      = Number(3)  // Octave name is SmallOctave
	Number4      = Number(4)  // Octave name is FirstOctave
	Number5      = Number(5)  // Octave name is SecondOctave
	Number6      = Number(6)  // Octave name is ThirdOctave
	Number7      = Number(7)  // Octave name is FourthOctave
	Number8      = Number(8)  // Octave name is FifthOctave
	Number9      = Number(9)  // Octave name is SixthOctave

	NumberDefault = Number(0) // Octave name is SubContraOctave
)

// NewOctave creates a new octave by the octave number.
func (n Number) NewOctave() (*Octave, error) {
	return NewByNumber(n)
}

// MustNewOctave creates a new octave by the octave number with panic on error.
func (n Number) MustNewOctave() *Octave {
	return MustNewByNumber(n)
}

// ErrOctaveNumberUnknown is the error that occurs when octave number is outside [-1; 9].
var ErrOctaveNumberUnknown = errors.New("unknown octave number")

// GetOctaveNameByNumber returns the name of the octave by the given octave number.
func GetOctaveNameByNumber(octaveNumber Number) (Name, error) {
	switch octaveNumber {
	case NumberMinus1:
		return NameSubSubContraOctave, nil
	case Number0:
		return NameSubContraOctave, nil
	case Number1:
		return NameContraOctave, nil
	case Number2:
		return NameGreatOctave, nil
	case Number3:
		return SmallOctave, nil
	case Number4:
		return NameFirstOctave, nil
	case Number5:
		return NameSecondOctave, nil
	case Number6:
		return NameThirdOctave, nil
	case Number7:
		return NameFourthOctave, nil
	case Number8:
		return NameFifthOctave, nil
	case Number9:
		return NameSixthOctave, nil
	}

	return "", ErrOctaveNumberUnknown
}

// MustGetOctaveNameByNumber returns the name of the octave by the given octave number with panic in case of incorrect octave number.
func MustGetOctaveNameByNumber(octaveNumber Number) Name {
	octaveName, err := GetOctaveNameByNumber(octaveNumber)
	if err != nil {
		panic(err)
	}

	return octaveName
}

// ErrOctaveNameUnknown is the error that occurs when octave name is unknown.
var ErrOctaveNameUnknown = errors.New("unknown octave name")

// GetOctaveNumberByName returns the name of the octave by the given octave number.
func GetOctaveNumberByName(octaveName Name) (Number, error) {
	switch octaveName {
	case NameSubSubContraOctave:
		return NumberMinus1, nil
	case NameSubContraOctave:
		return Number0, nil
	case NameContraOctave:
		return Number1, nil
	case NameGreatOctave:
		return Number2, nil
	case SmallOctave:
		return Number3, nil
	case NameFirstOctave:
		return Number4, nil
	case NameSecondOctave:
		return Number5, nil
	case NameThirdOctave:
		return Number6, nil
	case NameFourthOctave:
		return Number7, nil
	case NameFifthOctave:
		return Number8, nil
	case NameSixthOctave:
		return Number9, nil
	}

	return -100, ErrOctaveNameUnknown
}

// MustGetOctaveNumberByName returns the number of the octave by the given octave name with panic in case of incorrect octave number.
func MustGetOctaveNumberByName(octaveName Name) Number {
	octaveNumber, err := GetOctaveNumberByName(octaveName)
	if err != nil {
		panic(err)
	}

	return octaveNumber
}
