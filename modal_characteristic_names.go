package muse

import "github.com/pkg/errors"

// DegreeCharacteristicName is a characteristic name for an interval. Each interval has its own characteristic name.
type DegreeCharacteristicName string

const (
	DegreeCharacteristic3xDim = DegreeCharacteristicName("3xDim") // Triple Diminished
	DegreeCharacteristic2xDim = DegreeCharacteristicName("2xDim") // Double Diminished
	DegreeCharacteristicDim   = DegreeCharacteristicName("Dim")   // Diminished
	DegreeCharacteristicMinor = DegreeCharacteristicName("Minor") // Minor
	DegreeCharacteristicClean = DegreeCharacteristicName("Clean") // Clean
	DegreeCharacteristicMajor = DegreeCharacteristicName("Major") // Major
	DegreeCharacteristicAug   = DegreeCharacteristicName("Aug")   // Augmented
	DegreeCharacteristic2xAug = DegreeCharacteristicName("2xAug") // Double Augmented
	DegreeCharacteristic3xAug = DegreeCharacteristicName("3xAug") // Triple Augmented
)

var ErrDegreeCharacteristicUnknown = errors.New("unknown degree characteristic")

func mustGetDegreeCharacteristicName(diff int8, degreeNum DegreeNum) DegreeCharacteristicName {
	dcName, err := getDegreeCharacteristicName(diff, degreeNum)
	if err != nil {
		panic(err)
	}

	return dcName
}

//nolint:gomnd
func getDegreeCharacteristicName(diff int8, degreeNum DegreeNum) (DegreeCharacteristicName, error) {
	switch degreeNum {
	// These degrees can be clean
	case 1, 4, 5, 8:
		switch diff {
		case 0:
			return DegreeCharacteristicClean, nil
		case -1:
			return DegreeCharacteristicDim, nil
		case -2:
			return DegreeCharacteristic2xDim, nil
		case -3:
			return DegreeCharacteristic3xDim, nil
		case 1:
			return DegreeCharacteristicAug, nil
		case 2:
			return DegreeCharacteristic2xAug, nil
		case 3:
			return DegreeCharacteristic3xAug, nil
		}
	// These degrees can't be clean
	case 2, 3, 6, 7:
		switch diff {
		case 0:
			return DegreeCharacteristicMajor, nil
		case -1:
			return DegreeCharacteristicMinor, nil
		case -2:
			return DegreeCharacteristicDim, nil
		case -3:
			return DegreeCharacteristic2xDim, nil
		case 1:
			return DegreeCharacteristicAug, nil
		case 2:
			return DegreeCharacteristic2xAug, nil
		}
	}

	return "", errors.Wrapf(ErrDegreeCharacteristicUnknown, "diff: %d, degreeNum: %d", diff, degreeNum)
}

// getWeightByDCName defines weight by modal position name.
func getWeightByDCName(dcName DegreeCharacteristicName) Weight {
	switch dcName {
	case DegreeCharacteristic3xDim:
		return -4
	case DegreeCharacteristic2xDim:
		return -3
	case DegreeCharacteristicDim:
		return -2
	case DegreeCharacteristicMinor:
		return -1
	case DegreeCharacteristicClean:
		return -0
	case DegreeCharacteristicMajor:
		return +1
	case DegreeCharacteristicAug:
		return +2
	case DegreeCharacteristic2xAug:
		return +3
	case DegreeCharacteristic3xAug:
		return +4
	default:
		panic(dcName)
	}
}
