package degree

import (
	"errors"
	"fmt"
)

// CharacteristicName is a characteristic name for an interval. Each interval has its own characteristic name.
type CharacteristicName string

const (
	Characteristic3xDim = CharacteristicName("3xDim") // Triple Diminished
	Characteristic2xDim = CharacteristicName("2xDim") // Double Diminished
	CharacteristicDim   = CharacteristicName("Dim")   // Diminished
	CharacteristicMinor = CharacteristicName("Minor") // Minor
	CharacteristicClean = CharacteristicName("Clean") // Clean
	CharacteristicMajor = CharacteristicName("Major") // Major
	CharacteristicAug   = CharacteristicName("Aug")   // Augmented
	Characteristic2xAug = CharacteristicName("2xAug") // Double Augmented
	Characteristic3xAug = CharacteristicName("3xAug") // Triple Augmented
)

var ErrDegreeCharacteristicUnknown = errors.New("unknown degree characteristic")

func mustGetDegreeCharacteristicName(diff int8, degreeNum Number) CharacteristicName {
	dcName, err := getDegreeCharacteristicName(diff, degreeNum)
	if err != nil {
		panic(err)
	}

	return dcName
}

//nolint:mnd
func getDegreeCharacteristicName(diff int8, degreeNum Number) (CharacteristicName, error) {
	switch degreeNum {
	// These degrees can be clean
	case 1, 4, 5, 8:
		switch diff {
		case 0:
			return CharacteristicClean, nil
		case -1:
			return CharacteristicDim, nil
		case -2:
			return Characteristic2xDim, nil
		case -3:
			return Characteristic3xDim, nil
		case 1:
			return CharacteristicAug, nil
		case 2:
			return Characteristic2xAug, nil
		case 3:
			return Characteristic3xAug, nil
		}
	// These degrees can't be clean
	case 2, 3, 6, 7:
		switch diff {
		case 0:
			return CharacteristicMajor, nil
		case -1:
			return CharacteristicMinor, nil
		case -2:
			return CharacteristicDim, nil
		case -3:
			return Characteristic2xDim, nil
		case 1:
			return CharacteristicAug, nil
		case 2:
			return Characteristic2xAug, nil
		}
	}

	return "", fmt.Errorf("diff: '%d', degreeNum: '%d': %w", diff, degreeNum, ErrDegreeCharacteristicUnknown)
}

// getWeightByDCName defines Weight by modal position name.
func getWeightByDCName(dcName CharacteristicName) Weight {
	switch dcName {
	case Characteristic3xDim:
		return -4
	case Characteristic2xDim:
		return -3
	case CharacteristicDim:
		return -2
	case CharacteristicMinor:
		return -1
	case CharacteristicClean:
		return -0
	case CharacteristicMajor:
		return +1
	case CharacteristicAug:
		return +2
	case Characteristic2xAug:
		return +3
	case Characteristic3xAug:
		return +4
	default:
		panic(dcName)
	}
}
