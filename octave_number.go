package muse

import "github.com/pkg/errors"

// OctaveNumber is a number of octave.
type OctaveNumber int8

const (
	OctaveNumberMinus1 = OctaveNumber(-1) // Octave name is SubSubContraOctave
	OctaveNumber0      = OctaveNumber(0)  // Octave name is SubContraOctave
	OctaveNumber1      = OctaveNumber(1)  // Octave name is ContraOctave
	OctaveNumber2      = OctaveNumber(2)  // Octave name is GreatOctave
	OctaveNumber3      = OctaveNumber(3)  // Octave name is SmallOctave
	OctaveNumber4      = OctaveNumber(4)  // Octave name is FirstOctave
	OctaveNumber5      = OctaveNumber(5)  // Octave name is SecondOctave
	OctaveNumber6      = OctaveNumber(6)  // Octave name is ThirdOctave
	OctaveNumber7      = OctaveNumber(7)  // Octave name is FourthOctave
	OctaveNumber8      = OctaveNumber(8)  // Octave name is FifthOctave
	OctaveNumber9      = OctaveNumber(9)  // Octave name is SixthOctave

	OctaveNumberDefault = OctaveNumber(0) // Octave name is SubContraOctave
)

// ErrOctaveNumberUnknown is the error that occurs when octave number is outside of [-1; 9].
var ErrOctaveNumberUnknown = errors.New("unknown octave number")

// GetOctaveNameByNumber returns the name of the octave by the given octave number.
func GetOctaveNameByNumber(octaveNumber OctaveNumber) (OctaveName, error) {
	switch octaveNumber {
	case OctaveNumberMinus1:
		return OctaveNameSubSubContraOctave, nil
	case OctaveNumber0:
		return OctaveNameSubContraOctave, nil
	case OctaveNumber1:
		return OctaveNameContraOctave, nil
	case OctaveNumber2:
		return OctaveNameGreatOctave, nil
	case OctaveNumber3:
		return OctaveNameSmallOctave, nil
	case OctaveNumber4:
		return OctaveNameFirstOctave, nil
	case OctaveNumber5:
		return OctaveNameSecondOctave, nil
	case OctaveNumber6:
		return OctaveNameThirdOctave, nil
	case OctaveNumber7:
		return OctaveNameFourthOctave, nil
	case OctaveNumber8:
		return OctaveNameFifthOctave, nil
	case OctaveNumber9:
		return OctaveNameSixthOctave, nil
	}

	return "", ErrOctaveNumberUnknown
}

// MustNewOctave returns the name of the octave by the given octave number with panic in case of incorrect octave number.
func MustGetOctaveNameByNumber(octaveNumber OctaveNumber) OctaveName {
	octaveName, err := GetOctaveNameByNumber(octaveNumber)
	if err != nil {
		panic(err)
	}

	return octaveName
}
