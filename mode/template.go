package mode

import (
	"errors"
	"fmt"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
)

// Template is a set of halftones between degrees.
type Template halftone.Template

// Length returns amount of halftone in the template.
// It will be equal to mode length (amount of degrees) in a mode built from the template.
func (t Template) Length() degree.Number {
	return degree.Number(len(t)) //nolint:gosec // 12 is a maximum for usual mode, 24 is for microtonal modes, if you get a panic, you're doing something wrong.
}

// GetHalftonesByDegreeNum returns the number of halftone from the first step to the specified one.
func (t Template) GetHalftonesByDegreeNum(degreeNumber degree.Number) halftone.HalfTones {
	if degreeNumber == 0 || degreeNumber > t.Length() {
		return 0
	}

	var halfTonesFromPrime halftone.HalfTones
	for i := degree.Number(0); i < degreeNumber; i++ {
		halfTonesFromPrime += t[i]
	}

	return halfTonesFromPrime
}

// TemplateIteratorResult is the result of iterating over a mode template.
// It returns the step number,
// the number of halftone from the last step to the current one,
// and the number of halftone from the first step to the current one.
type TemplateIteratorResult func() (degree.Number, halftone.HalfTones, halftone.HalfTones)

// IterateOneRound returns channel for iterating over a mode template.
// At each iteration, it returns TemplateIteratorResult containing a set of values.
func (t Template) IterateOneRound(withOctave bool) <-chan TemplateIteratorResult {
	send := func(degreeNum degree.Number, halfTones, halfTonesFromPrime halftone.HalfTones) TemplateIteratorResult {
		return func() (degree.Number, halftone.HalfTones, halftone.HalfTones) {
			return degreeNum, halfTones, halfTonesFromPrime
		}
	}

	f := func(c chan TemplateIteratorResult) {
		var halfTonesFromPrime halftone.HalfTones
		var degreeNum degree.Number
		const startingDegree = 2 // 2 means we are sending data starting from the second degree
		length := t.Length()
		for i := degree.Number(0); i < length; i++ {
			if i > 0 && i == length-1 && !withOctave {
				break
			}
			halfTones := t[i]
			halfTonesFromPrime += halfTones
			degreeNum = i + startingDegree
			c <- send(degreeNum, halfTones, halfTonesFromPrime)
		}
		close(c)
	}

	c := make(chan TemplateIteratorResult)
	go f(c)

	return c
}

// ErrInvalidModeTemplate is returned when mode template is invalid.
var ErrInvalidModeTemplate = errors.New("invalid mode template")

// Validate checks mode template for length, halftone sum in octave, zero intervals.
func (t Template) Validate() error {
	if t.Length() < 1 {
		return fmt.Errorf("zero length: %w", ErrInvalidModeTemplate)
	}

	var halfTones halftone.HalfTones
	for _, interval := range t {
		if interval == 0 {
			return fmt.Errorf("zero interval: %w", ErrInvalidModeTemplate)
		}

		halfTones += interval
	}

	if halfTones != halftone.HalfTonesInOctave {
		return fmt.Errorf("halftones overflows octave: %w", ErrInvalidModeTemplate)
	}

	return nil
}

// IsDiatonic checks if the mode is diatonic or not.
// Strictly speaking, diatonic scales are characterized by the ability to be decomposed into fifths.
// In this case, we mean that diatonic modes have seven degrees and do not have augmented seconds.
func (t Template) IsDiatonic() bool {
	if t.Length() != DegreesInDiatonic {
		return false
	}

	const augSecond = halftone.HalfTones(3)
	for _, interval := range t {
		if interval >= augSecond {
			return false
		}
	}

	return true
}

// IsHeptatonic checks if the mode is heptatonic or not.
// In this case, we mean that heptatonic modes have seven degrees.
func (t Template) IsHeptatonic() bool {
	return t.Length() == DegreesInHeptatonic
}

// RearrangeFromDegree rebuilds mode template from the specified degree number.
func (t Template) RearrangeFromDegree(degree degree.Number) Template {
	ltArranged := make(Template, 0, t.Length())
	ltArranged = append(ltArranged, t[degree-1:]...)
	ltArranged = append(ltArranged, t[:degree-1]...)

	return ltArranged
}
