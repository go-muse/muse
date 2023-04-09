package muse

import (
	"github.com/pkg/errors"
)

// ModalPosition is a modal position of a relative degree.
type ModalPosition struct {
	name   ModalPositionName
	weight Weight
}

func (mp *ModalPosition) GetWeight() Weight {
	return mp.weight
}

func (mp *ModalPosition) GetName() ModalPositionName {
	return mp.name
}

func (mp *ModalPosition) Copy() *ModalPosition {
	if mp == nil {
		return nil
	}

	return &ModalPosition{
		name:   mp.name,
		weight: mp.GetWeight(),
	}
}

// Weight is relative property of characteristics, for their comparison and ordering.
// This is not a real musical concept.
type Weight int8

type ModalPositionName string

const (
	ModalPositionNameLow     = ModalPositionName("Low")
	ModalPositionNameNeutral = ModalPositionName("Neutral")
	ModalPositionNameHigh    = ModalPositionName("High")
)

func MustNewModalPosition(dcName DegreeCharacteristicName, weight Weight) *ModalPosition {
	nmp, err := NewModalPosition(dcName, weight)
	if err != nil {
		panic(err)
	}

	return nmp
}

var ErrDegreeCharacteristicNameUnknown = errors.New("unknown degree characteristic name")

func NewModalPosition(dcName DegreeCharacteristicName, weight Weight) (*ModalPosition, error) {
	switch dcName {
	case DegreeCharacteristicMinor, DegreeCharacteristicDim, DegreeCharacteristic2xDim, DegreeCharacteristic3xDim:
		return &ModalPosition{ModalPositionNameLow, weight}, nil
	case DegreeCharacteristicClean:
		return &ModalPosition{ModalPositionNameNeutral, weight}, nil
	case DegreeCharacteristicMajor, DegreeCharacteristicAug, DegreeCharacteristic2xAug, DegreeCharacteristic3xAug:
		return &ModalPosition{ModalPositionNameHigh, weight}, nil
	}

	return nil, errors.Wrapf(ErrDegreeCharacteristicNameUnknown, "got degree characteristic name: %s", dcName)
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
