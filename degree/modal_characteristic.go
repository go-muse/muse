package degree

import (
	"errors"
	"fmt"

	"github.com/go-muse/muse/common/convert"
	"github.com/go-muse/muse/halftone"
)

// ModalCharacteristic is a set of information about calculated relative modal position.
type ModalCharacteristic struct {
	name                  CharacteristicName
	degree                *Degree
	relativeModalPosition *ModalPosition
}

// Name returns modal characteristic's name.
func (mc *ModalCharacteristic) Name() CharacteristicName {
	if mc == nil {
		return ""
	}

	return mc.name
}

// RelativeModalPosition returns relative modal position as part of modal characteristic of the deg.
func (mc *ModalCharacteristic) RelativeModalPosition() *ModalPosition {
	if mc == nil {
		return nil
	}

	return mc.relativeModalPosition
}

// Degree returns the degree which, when compared, yields this modal characteristic.
func (mc *ModalCharacteristic) Degree() *Degree {
	if mc == nil {
		return nil
	}

	return mc.degree
}

type ModalCharacteristics []ModalCharacteristic

// Add appends a new ModalCharacteristic to ModalCharacteristics and returns the updated slice.
func (mcs ModalCharacteristics) Add(mc ModalCharacteristic) ModalCharacteristics {
	return append(mcs, mc)
}

// Copy creates and returns a copy of ModalCharacteristics.
func (mcs ModalCharacteristics) Copy() ModalCharacteristics {
	if mcs == nil {
		return nil
	}

	copyMcs := make(ModalCharacteristics, len(mcs))
	copy(copyMcs, mcs)

	return copyMcs
}

func newModalCharacteristic(dcName CharacteristicName, degree *Degree) *ModalCharacteristic {
	mc := &ModalCharacteristic{name: dcName, degree: degree}
	mc.relativeModalPosition = MustNewModalPosition(dcName, getWeightByDCName(dcName))

	return mc
}

// String is stringer for ModalCharacteristic.
func (mc *ModalCharacteristic) String() string {
	return fmt.Sprintf("Name: '%s', Weight: '%d', note name:'%s', RMP: '%s', AMP: '%s'\n", mc.Name(), mc.RelativeModalPosition().Weight(), mc.Degree().Note().Name(), mc.RelativeModalPosition().Name(), mc.Degree().AbsoluteModalPosition().Name())
}

// CalculateRelativeMC defines modal characteristic by degree's number.
func CalculateRelativeMC(degreeNum Number, nextDegree *Degree, halfTonesFromPrime halftone.HalfTones) (*ModalCharacteristic, error) {
	originalPosition, err := getOriginalPositionOfDegree(degreeNum)
	if err != nil {
		return nil, fmt.Errorf("get original position of degree '%d': %w", degreeNum, err)
	}

	diff := convert.SubUint8Uint8(uint8(halfTonesFromPrime), uint8(originalPosition))
	dcName, err := getDegreeCharacteristicName(diff, degreeNum)
	if err != nil {
		return nil, fmt.Errorf("get degree characteristic name of degree '%d' with diff '%d': %w", degreeNum, diff, err)
	}

	return newModalCharacteristic(dcName, nextDegree), nil
}

// MustCalculateRelativeMC defines modal characteristic by degree's number. Panics in case of error.
func MustCalculateRelativeMC(degreeNum Number, nextDegree *Degree, halfTonesFromPrime halftone.HalfTones) *ModalCharacteristic {
	originalPosition := mustGetOriginalPositionOfDegree(degreeNum)

	diff := convert.SubUint8Uint8(uint8(halfTonesFromPrime), uint8(originalPosition))
	dcName := mustGetDegreeCharacteristicName(diff, degreeNum)

	return newModalCharacteristic(dcName, nextDegree)
}

func mustGetOriginalPositionOfDegree(degreeNum Number) halftone.HalfTones {
	op, err := getOriginalPositionOfDegree(degreeNum)
	if err != nil {
		panic(err)
	}

	return op
}

var ErrDegreePositionUnknown = errors.New("unknown degree position")

//nolint:mnd
func getOriginalPositionOfDegree(degreeNum Number) (halftone.HalfTones, error) {
	switch degreeNum {
	case 1:
		return 0, nil
	case 2:
		return 2, nil
	case 3:
		return 4, nil
	case 4:
		return 5, nil
	case 5:
		return 7, nil
	case 6:
		return 9, nil
	case 7:
		return 11, nil
	case 8:
		return 12, nil
	}

	return 0, fmt.Errorf("unknown original position of degree by degree number '%d': %w", degreeNum, ErrDegreePositionUnknown)
}
