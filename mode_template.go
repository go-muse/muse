package muse

import (
	"github.com/pkg/errors"
)

// ModeTemplate is a set of halftones between degrees.
type ModeTemplate []HalfTones

// Length returns amount of halftones in the template.
// It will be equal to mode length (amount of degrees) in a mode built from the template.
func (mt ModeTemplate) Length() DegreeNum {
	return DegreeNum(len(mt))
}

// GetHalftonesByDegreeNum returns the number of half tones from the first step to the specified one.
func (mt ModeTemplate) GetHalftonesByDegreeNum(degreeNumber DegreeNum) HalfTones {
	if degreeNumber == 0 || degreeNumber > mt.Length() {
		return 0
	}

	var halfTonesFromPrime HalfTones
	for i := DegreeNum(0); i < degreeNumber; i++ {
		halfTonesFromPrime += mt[i]
	}

	return halfTonesFromPrime
}

// ModeTemplateIteratorResult is the result of iterating over a mode template.
// It returns the step number,
// the number of half tones from the last step to the current one,
// and the number of half tones from the first step to the current one.
type ModeTemplateIteratorResult func() (DegreeNum, HalfTones, HalfTones)

// IterateOneRound returns channel for iterating over a mode template.
// At each iteration, it returns ModeTemplateIteratorResult containing a set of values.
func (mt ModeTemplate) IterateOneRound(withOctave bool) <-chan ModeTemplateIteratorResult {
	send := func(degreeNum DegreeNum, halfTones, halfTonesFromPrime HalfTones) ModeTemplateIteratorResult {
		return func() (DegreeNum, HalfTones, HalfTones) { return degreeNum, halfTones, halfTonesFromPrime }
	}

	f := func(c chan ModeTemplateIteratorResult) {
		var halfTonesFromPrime HalfTones
		var degreeNum DegreeNum
		const startingDegree = 2 // 2 means we are sending data starting from the second degree
		for i, halfTones := range mt {
			if i > 0 && i == int(mt.Length()-1) && !withOctave {
				break
			}
			halfTonesFromPrime += halfTones
			degreeNum = DegreeNum(i) + startingDegree
			c <- send(degreeNum, halfTones, halfTonesFromPrime)
		}
		close(c)
	}

	c := make(chan ModeTemplateIteratorResult)
	go f(c)

	return c
}

// Validate checks mode template for length, halftones sum in octave, zero intervals.
func (mt ModeTemplate) Validate() error {
	if mt.Length() < 1 {
		return errors.New("invalid mode template length")
	}

	var halfTones HalfTones
	for _, interval := range mt {
		if interval == 0 {
			return errors.New("invalid mode template interval: zero")
		}

		halfTones += interval
	}

	if halfTones != HalftonesInOctave {
		return errors.New("invalid mode template")
	}

	return nil
}

// IsDiatonic checks if the mode is diatonic or not.
// Strictly speaking, diatonic scales are characterized by the ability to be decomposed into fifths.
// In this case, we mean that diatonic modes have seven degrees and do not have augmented seconds.
func (mt ModeTemplate) IsDiatonic() bool {
	if mt.Length() != DegreesInDiatonic {
		return false
	}

	const augSecond = HalfTones(3)
	for _, interval := range mt {
		if interval >= augSecond {
			return false
		}
	}

	return true
}

// IsHeptatonic checks if the mode is heptatonic or not.
// In this case, we mean that heptatonic modes have seven degrees.
func (mt ModeTemplate) IsHeptatonic() bool {
	return mt.Length() == DegreesInHeptatonic
}

// RearrangeFromDegree rebuilds mode template from the specified degree number.
func (mt ModeTemplate) RearrangeFromDegree(degree DegreeNum) ModeTemplate {
	ltArranged := make(ModeTemplate, 0, mt.Length())
	ltArranged = append(ltArranged, mt[degree-1:]...)
	ltArranged = append(ltArranged, mt[:degree-1]...)

	return ltArranged
}
