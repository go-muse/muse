package degree

import (
	"errors"
	"fmt"
)

// Weight is relative property of characteristics, for their comparison and ordering.
// This is not a real musical concept.
type Weight int8

// ModalPosition is a modal position of a relative degree.
type ModalPosition struct {
	name   ModalPositionName
	weight Weight
}

// Weight returns modal position's weight.
func (mp *ModalPosition) Weight() Weight {
	if mp == nil {
		return 0
	}

	return mp.weight
}

// Name returns modal position's name.
func (mp *ModalPosition) Name() ModalPositionName {
	if mp == nil {
		return ""
	}

	return mp.name
}

// Copy makes a copy of the modal position.
func (mp *ModalPosition) Copy() *ModalPosition {
	if mp == nil {
		return nil
	}

	return &ModalPosition{
		name:   mp.name,
		weight: mp.Weight(),
	}
}

func MustNewModalPosition(dcName CharacteristicName, weight Weight) *ModalPosition {
	nmp, err := NewModalPosition(dcName, weight)
	if err != nil {
		panic(err)
	}

	return nmp
}

var ErrDegreeCharacteristicNameUnknown = errors.New("unknown degree characteristic name")

func NewModalPosition(dcName CharacteristicName, weight Weight) (*ModalPosition, error) {
	switch dcName {
	case CharacteristicMinor, CharacteristicDim, Characteristic2xDim, Characteristic3xDim:
		return &ModalPosition{ModalPositionNameLow, weight}, nil
	case CharacteristicClean:
		return &ModalPosition{ModalPositionNameNeutral, weight}, nil
	case CharacteristicMajor, CharacteristicAug, Characteristic2xAug, Characteristic3xAug:
		return &ModalPosition{ModalPositionNameHigh, weight}, nil
	}

	return nil, fmt.Errorf("got degree characteristic name: '%s': %w", dcName, ErrDegreeCharacteristicNameUnknown)
}

func NewModalPositionByWeight(weight Weight) *ModalPosition {
	return &ModalPosition{getModalPositionNameByWeight(weight), weight}
}

func getModalPositionNameByWeight(weight Weight) ModalPositionName {
	var mpn ModalPositionName

	switch {
	case weight < 0:
		mpn = ModalPositionNameHigh
	case weight == 0:
		mpn = ModalPositionNameNeutral
	case weight > 0:
		mpn = ModalPositionNameLow
	}

	return mpn
}
