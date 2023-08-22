package muse

import (
	"fmt"
)

// IntervalType is string type for interval's name.
type IntervalType string

// Interval is determined by 1) half tones between degrees 2) amount of degrees between.
// Such interval is known as a diatonic interval, but exists not only in diatonic modes.
type Interval struct {
	chromaticInterval *IntervalChromatic
	degree1, degree2  *Degree
}

func (iwd *Interval) IntervalChromatic() *IntervalChromatic {
	if iwd == nil {
		return nil
	}

	return iwd.chromaticInterval
}

func (iwd *Interval) HalfTones() HalfTones {
	if iwd == nil {
		return 0
	}

	return iwd.chromaticInterval.HalfTones()
}

func (iwd *Interval) Name() IntervalName {
	if iwd == nil {
		return ""
	}

	return iwd.chromaticInterval.Name()
}

func (iwd *Interval) ShortName() IntervalName {
	if iwd == nil {
		return ""
	}

	return iwd.chromaticInterval.ShortName()
}

func (iwd *Interval) Degree1() *Degree {
	if iwd == nil {
		return nil
	}

	return iwd.degree1
}

func (iwd *Interval) Degree2() *Degree {
	if iwd == nil {
		return nil
	}

	return iwd.degree2
}

func (iwd *Interval) String() string {
	return fmt.Sprintf("Half tones: %d, name: %s, short name: %s, Sonance: %d, first degree num: %d, second degree num: %d",
		iwd.chromaticInterval.HalfTones(),
		iwd.chromaticInterval.Name(),
		iwd.chromaticInterval.ShortName(),
		iwd.chromaticInterval.Sonance,
		iwd.degree1.Number(),
		iwd.degree2.Number(),
	)
}

// Mode interval is determined by degrees and half tones between them.
// The function is needed in case of lack of information about the halftone distance of the steps from each other.
func newIntervalByDegreesAndHalfTones(halfTones HalfTones, degree1, degree2 *Degree) (*Interval, error) {
	degreesDiff := degree2.Number() - degree1.Number()

	chromaticInterval, err := NewIntervalByHalfTonesAndDegrees(halfTones, degreesDiff)
	if err != nil {
		return nil, ErrIntervalUnknown
	}

	return &Interval{
		chromaticInterval: chromaticInterval,
		degree1:           degree1,
		degree2:           degree2,
	}, nil
}

// NewIntervalByDegrees creates mode interval by degrees and half tones between them.
func NewIntervalByDegrees(degree1, degree2 *Degree) (*Interval, error) {
	halfTonesDiff := degree2.halfTonesFromPrime - degree1.halfTonesFromPrime

	return newIntervalByDegreesAndHalfTones(halfTonesDiff, degree1, degree2)
}
