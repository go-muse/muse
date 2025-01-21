package interval

import (
	"fmt"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
)

// Type is string type for interval's name.
type Type string

// Diatonic interval is determined by 1) halftone between degrees 2) amount of degrees between.
// Such interval is known as a diatonic interval, but exists not only in diatonic modes.
type Diatonic struct {
	chromaticInterval *Chromatic
	degree1, degree2  *degree.Degree
}

// NewDiatonic creates mode interval by degrees and halftone between them.
//
// The function automatically creates a chromatic interval within itself.
// The number of halftones from the tonic for each degree must be specified.
func NewDiatonic(degree1, degree2 *degree.Degree) (*Diatonic, error) {
	halfTonesDiff := degree2.HalfTonesFromPrime() - degree1.HalfTonesFromPrime()

	return newIntervalByDegreesAndHalfTones(halfTonesDiff, degree1, degree2)
}

// Chromatic returns chromatic representation of the diatonic interval.
func (di *Diatonic) Chromatic() *Chromatic {
	if di == nil {
		return nil
	}

	return di.chromaticInterval
}

// HalfTones returns the number of halftones between the two degrees.
func (di *Diatonic) HalfTones() halftone.HalfTones {
	if di == nil {
		return 0
	}

	return di.chromaticInterval.HalfTones()
}

// Name returns the name of the diatonic interval.
func (di *Diatonic) Name() Name {
	if di == nil {
		return ""
	}

	return di.chromaticInterval.Name()
}

// ShortName returns the short name of the diatonic interval.
func (di *Diatonic) ShortName() Name {
	if di == nil {
		return ""
	}

	return di.chromaticInterval.ShortName()
}

// Degree1 returns the first degree of the diatonic interval.
func (di *Diatonic) Degree1() *degree.Degree {
	if di == nil {
		return nil
	}

	return di.degree1
}

// Degree2 returns the second degree of the diatonic interval.
func (di *Diatonic) Degree2() *degree.Degree {
	if di == nil {
		return nil
	}

	return di.degree2
}

// String returns a string representation of the diatonic interval.
func (di *Diatonic) String() string {
	return fmt.Sprintf("Half tones: %d, name: %s, short name: %s, Sonance: %d, first degree num: %d, second degree num: %d",
		di.chromaticInterval.HalfTones(),
		di.chromaticInterval.Name(),
		di.chromaticInterval.ShortName(),
		di.chromaticInterval.Sonance,
		di.degree1.Number(),
		di.degree2.Number(),
	)
}

// newIntervalByDegreesAndHalfTones creates a new Diatonic a diatonic interval
// based on the number of halftones between two degrees (degree1 and degree2).
//
// Mode interval is determined by degrees and halftone between them.
// The function is needed in case of lack of information about the halftone distance of the steps from each other.
func newIntervalByDegreesAndHalfTones(halfTones halftone.HalfTones, degree1, degree2 *degree.Degree) (*Diatonic, error) {
	degreesDiff := degree2.Number() - degree1.Number()

	chromaticInterval, err := NewIntervalByHalfTonesAndDegrees(halfTones, degreesDiff)
	if err != nil {
		return nil, ErrIntervalUnknown
	}

	return &Diatonic{
		chromaticInterval: chromaticInterval,
		degree1:           degree1,
		degree2:           degree2,
	}, nil
}
