package muse

import (
	"fmt"

	"github.com/pkg/errors"
)

// ModalCharacteristic is a set of information about calculated relative modal position.
type ModalCharacteristic struct {
	name                  DegreeCharacteristicName
	degree                *Degree
	interval              *Interval
	relativeModalPosition *ModalPosition
}

type ModalCharacteristics []ModalCharacteristic

func (mcs *ModalCharacteristics) Add(mc ModalCharacteristic) {
	*mcs = append(*mcs, mc)
}

func (mc *ModalCharacteristic) GetName() DegreeCharacteristicName {
	return mc.name
}

func (mcs ModalCharacteristics) Copy() ModalCharacteristics {
	var copyMcs ModalCharacteristics
	if mcs == nil {
		copyMcs = nil
	} else {
		copyMcs = make(ModalCharacteristics, len(mcs))
		copy(copyMcs, mcs)
	}

	return copyMcs
}

func newModalCharacteristic(dcName DegreeCharacteristicName, degree *Degree, interval *Interval) *ModalCharacteristic {
	mc := &ModalCharacteristic{name: dcName, degree: degree, interval: interval}
	mc.relativeModalPosition = MustNewModalPosition(dcName, getWeightByDCName(dcName))

	return mc
}

// String is stringer for ModalCharacteristic.
func (mc *ModalCharacteristic) String() string {
	return fmt.Sprintf("Name: %s, weight: %d, note name:%s, RMP: %s, AMP: %s \n", mc.name, mc.relativeModalPosition.GetWeight(), mc.degree.note.Name(), mc.relativeModalPosition.name, mc.degree.absoluteModalPosition.name)
}

// setRelativeMCsOfDegree calculates relative modal characteristics of the degree and stores them in it.
func (degree *Degree) setRelativeMCsOfDegree(modeTemplate ModeTemplate) {
	mcs := make(ModalCharacteristics, 0, modeTemplate.Length())

	// Rebuild mode template from the given degree and iterate through it
	for iteratorResult := range modeTemplate.RearrangeFromDegree(degree.Number()).IterateOneRound(true) {
		degreeNum, _, halfTonesFromPrime := iteratorResult()
		nextDegree := degree.GetForwardDegreeByDegreeNum(degreeNum - 1)

		interval := MustNewIntervalByHalfTonesAndDegrees(halfTonesFromPrime, degreeNum-1) // -1 means we always get interval from 1st degree
		intervalWithDegrees := &Interval{interval, degree, nextDegree}

		// Calculation of relative modal characteristics of the degree
		mc := calculateRelativeMC(degreeNum, nextDegree, halfTonesFromPrime, intervalWithDegrees)
		mcs.Add(*mc)
	}

	degree.modalCharacteristics = mcs
}

// calculateRelativeMC defines modal characteristic by degree's number.
func calculateRelativeMC(degreeNum DegreeNum, nextDegree *Degree, halfTonesFromPrime HalfTones, interval *Interval) *ModalCharacteristic {
	originalPosition := mustGetOriginalPositionOfDegree(degreeNum)

	diff := int8(halfTonesFromPrime - originalPosition)
	dcName := mustGetDegreeCharacteristicName(diff, degreeNum)

	return newModalCharacteristic(dcName, nextDegree, interval)
}

func mustGetOriginalPositionOfDegree(degreeNum DegreeNum) HalfTones {
	op, err := getOriginalPositionOfDegree(degreeNum)
	if err != nil {
		panic(err)
	}

	return op
}

var ErrDegreePositionUnknown = errors.New("unknown degree position")

//nolint:gomnd
func getOriginalPositionOfDegree(degreeNum DegreeNum) (HalfTones, error) {
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

	return 0, errors.Wrapf(ErrDegreePositionUnknown, "unknown original position of degree by degree number: %d", degreeNum)
}

// setRelativeMCsOfDegree calculates absolute modal positions of all the degrees in the mode and stores them.
func (m *Mode) setAbsoluteModalPositions() {
	if m == nil {
		return
	}
	for degree := range m.GetFirstDegree().IterateOneRound(false) {
		var weightSum Weight
		for _, mc := range degree.modalCharacteristics {
			weightSum += mc.relativeModalPosition.GetWeight()
		}
		degree.absoluteModalPosition = NewModalPositionByWeight(weightSum)
	}
}

func (m *Mode) setRelativeModalPositions(modeTemplate ModeTemplate) {
	for degree := range m.GetFirstDegree().IterateOneRound(false) {
		degree.setRelativeMCsOfDegree(modeTemplate)
	}
}
