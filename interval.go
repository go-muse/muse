package muse

import (
	"fmt"

	"github.com/pkg/errors"
)

// IntervalName is string type for interval's name.
type IntervalName string

// IntervalType is string type for interval's name.
type IntervalType string

type intervalNameExtended struct {
	name      IntervalName
	shortName IntervalName
}

// Name returns interval's name.
func (i *ChromaticInterval) Name() IntervalName {
	return i.names.name
}

// ShortName returns interval's short name.
func (i *ChromaticInterval) ShortName() IntervalName {
	return i.names.shortName
}

// HalfTones returns amount of half tones in the interval.
func (i *ChromaticInterval) HalfTones() HalfTones {
	return i.halfTones
}

// ChromaticInterval is the interval defined just by halftones, without degrees.
type ChromaticInterval struct {
	Sonance
	names     *intervalNameExtended
	halfTones HalfTones
}

// Interval is determined by 1) half tones between degrees 2) amount of degrees between.
// Such interval is known as a diatonic interval, but exists not only in diatonic modes.
type Interval struct {
	*ChromaticInterval
	degree1, degree2 *Degree
}

func (iwd *Interval) GetInterval() *ChromaticInterval {
	return iwd.ChromaticInterval
}

func (iwd *Interval) GetDegree1() *Degree {
	return iwd.degree1
}

func (iwd *Interval) GetDegree2() *Degree {
	return iwd.degree2
}

func (iwd *Interval) String() string {
	return fmt.Sprintf("Half tones: %d, name: %s, short name: %s, Sonance: %d, first degree num: %d, second degree num: %d",
		iwd.ChromaticInterval.HalfTones(),
		iwd.ChromaticInterval.Name(),
		iwd.ChromaticInterval.ShortName(),
		iwd.ChromaticInterval.Sonance,
		iwd.degree1.Number(),
		iwd.degree2.Number(),
	)
}

// NewIntervalChromatic creates interval just by half tones between the notes
// Such interval is known as chromatic  interval or acoustic interval.
func NewIntervalChromatic(halfTones HalfTones) (*ChromaticInterval, error) {
	switch halfTones {
	case IntervalHalfTones0:
		return IntervalPerfectUnison(), nil
	case IntervalHalfTones1:
		return IntervalMinorSecond(), nil
	case IntervalHalfTones2:
		return IntervalMajorSecond(), nil
	case IntervalHalfTones3:
		return IntervalMinorThird(), nil
	case IntervalHalfTones4:
		return IntervalMajorThird(), nil
	case IntervalHalfTones5:
		return IntervalPerfectFourth(), nil
	case IntervalHalfTones6:
		return IntervalTritone(), nil
	case IntervalHalfTones7:
		return IntervalPerfectFifth(), nil
	case IntervalHalfTones8:
		return IntervalMinorSixth(), nil
	case IntervalHalfTones9:
		return IntervalMajorSixth(), nil
	case IntervalHalfTones10:
		return IntervalMinorSeventh(), nil
	case IntervalHalfTones11:
		return IntervalMajorSeventh(), nil
	case IntervalHalfTones12:
		return IntervalPerfectOctave(), nil

	case IntervalHalfTones13:
		return IntervalMinorNinth(), nil
	case IntervalHalfTones14:
		return IntervalMajorNinth(), nil
	case IntervalHalfTones15:
		return IntervalMinorTenth(), nil
	case IntervalHalfTones16:
		return IntervalMajorTenth(), nil
	case IntervalHalfTones17:
		return IntervalPerfectEleventh(), nil
	case IntervalHalfTones18:
		// no name for tritone after octave
	case IntervalHalfTones19:
		return IntervalPerfectTwelfth(), nil
	case IntervalHalfTones20:
		return IntervalMinorThirteenth(), nil
	case IntervalHalfTones21:
		return IntervalMajorThirteenth(), nil
	case IntervalHalfTones22:
		return IntervalMinorFourteenth(), nil
	case IntervalHalfTones23:
		return IntervalMajorFourteenth(), nil
	case IntervalHalfTones24:
		return IntervalPerfectFifteenth(), nil
	}

	return nil, ErrUnknownInterval
}

var ErrUnknownInterval = errors.New("unknown interval")

// NewIntervalByHalfTonesAndDegrees creates interval by 1) half tones between degrees 2) amount of degrees between.
//
//nolint:gomnd
func NewIntervalByHalfTonesAndDegrees(halfTones HalfTones, degrees DegreeNum) (*ChromaticInterval, error) {
	switch halfTones {
	case IntervalHalfTones0:
		switch degrees {
		case 0:
			return NewIntervalChromatic(halfTones)
		case 1:
			return IntervalDiminishedSecond(), nil
		}
	case IntervalHalfTones1:
		switch degrees {
		case 0:
			return IntervalAugmentedUnison(), nil
		case 1:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones2:
		switch degrees {
		case 1:
			return NewIntervalChromatic(halfTones)
		case 2:
			return IntervalDiminishedThird(), nil
		}
	case IntervalHalfTones3:
		switch degrees {
		case 1:
			return IntervalAugmentedSecond(), nil
		case 2:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones4:
		switch degrees {
		case 2:
			return NewIntervalChromatic(halfTones)
		case 3:
			return IntervalDiminishedFourth(), nil
		}
	case IntervalHalfTones5:
		switch degrees {
		case 2:
			return IntervalAugmentedThird(), nil
		case 3:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones6:
		switch degrees {
		case 3:
			return IntervalAugmentedFourth(), nil
		case 4:
			return IntervalDiminishedFifth(), nil
		}
	case IntervalHalfTones7:
		switch degrees {
		case 4:
			return NewIntervalChromatic(halfTones)
		case 5:
			return IntervalDiminishedSixth(), nil
		}
	case IntervalHalfTones8:
		switch degrees {
		case 4:
			return IntervalAugmentedFifth(), nil
		case 5:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones9:
		switch degrees {
		case 5:
			return NewIntervalChromatic(halfTones)
		case 6:
			return IntervalDiminishedSeventh(), nil
		}
	case IntervalHalfTones10:
		switch degrees {
		case 5:
			return IntervalAugmentedSixth(), nil
		case 6:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones11:
		switch degrees {
		case 6:
			return NewIntervalChromatic(halfTones)
		case 7:
			return IntervalDiminishedOctave(), nil
		}
	case IntervalHalfTones12:
		switch degrees {
		case 6:
			return IntervalAugmentedSeventh(), nil
		case 7:
			return NewIntervalChromatic(halfTones)
		}
	}

	return nil, ErrUnknownInterval
}

// MustNewIntervalByHalfTonesAndDegrees creates interval by 1) half tones between degrees 2) amount of degrees between.
// The function panics in case of unknown interval.
func MustNewIntervalByHalfTonesAndDegrees(halfTones HalfTones, degrees DegreeNum) *ChromaticInterval {
	chromaticInterval, err := NewIntervalByHalfTonesAndDegrees(halfTones, degrees)
	if err != nil {
		panic(err)
	}

	return chromaticInterval
}

// Mode interval is determined by degrees and half tones between them.
// The function is needed in case of lack of information about the halftone distance of the steps from each other.
func newIntervalByDegreesAndHalfTones(halfTones HalfTones, degree1, degree2 *Degree) (*Interval, error) {
	degreesDiff := degree2.Number() - degree1.Number()

	chromaticInterval, err := NewIntervalByHalfTonesAndDegrees(halfTones, degreesDiff)
	if err != nil {
		return nil, ErrUnknownInterval
	}

	return &Interval{
		ChromaticInterval: chromaticInterval,
		degree1:           degree1,
		degree2:           degree2,
	}, nil
}

// NewIntervalByDegrees creates mode interval by degrees and half tones between them.
func NewIntervalByDegrees(degree1, degree2 *Degree) (*Interval, error) {
	halfTonesDiff := degree2.halfTonesFromPrime - degree1.halfTonesFromPrime

	return newIntervalByDegreesAndHalfTones(halfTonesDiff, degree1, degree2)
}

// NewIntervalByName returns new interval by it's name.
func NewIntervalByName(intervalName IntervalName) (*ChromaticInterval, error) {
	switch intervalName {
	// Chromatic intervals
	case IntervalNamePerfectUnison, IntervalNamePerfectUnisonShort:
		return IntervalPerfectUnison(), nil
	case IntervalNameMinorSecond, IntervalNameMinorSecondShort:
		return IntervalMinorSecond(), nil
	case IntervalNameMajorSecond, IntervalNameMajorSecondShort:
		return IntervalMajorSecond(), nil
	case IntervalNameMinorThird, IntervalNameMinorThirdShort:
		return IntervalMinorThird(), nil
	case IntervalNameMajorThird, IntervalNameMajorThirdShort:
		return IntervalMajorThird(), nil
	case IntervalNamePerfectFourth, IntervalNamePerfectFourthShort:
		return IntervalPerfectFourth(), nil
	case IntervalNameTritone, IntervalNameTritoneShort:
		return IntervalTritone(), nil
	case IntervalNamePerfectFifth, IntervalNamePerfectFifthShort:
		return IntervalPerfectFifth(), nil
	case IntervalNameMinorSixth, IntervalNameMinorSixthShort:
		return IntervalMinorSixth(), nil
	case IntervalNameMajorSixth, IntervalNameMajorSixthShort:
		return IntervalMajorSixth(), nil
	case IntervalNameMinorSeventh, IntervalNameMinorSeventhShort:
		return IntervalMinorSeventh(), nil
	case IntervalNameMajorSeventh, IntervalNameMajorSeventhShort:
		return IntervalMajorSeventh(), nil
	case IntervalNamePerfectOctave, IntervalNamePerfectOctaveShort:
		return IntervalPerfectOctave(), nil

	// Diatonic intervals
	case IntervalNameModifiedDiminishedSecond, IntervalNameModifiedDiminishedSecondShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedAugmentedUnison, IntervalNameModifiedAugmentedUnisonShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedDiminishedThird, IntervalNameModifiedDiminishedThirdShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedAugmentedSecond, IntervalNameModifiedAugmentedSecondShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedDiminishedFourth, IntervalNameModifiedDiminishedFourthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedAugmentedThird, IntervalNameModifiedAugmentedThirdShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedDiminishedFifth, IntervalNameModifiedDiminishedFifthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedAugmentedFourth, IntervalNameModifiedAugmentedFourthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedDiminishedSixth, IntervalNameModifiedDiminishedSixthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedAugmentedFifth, IntervalNameModifiedAugmentedFifthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedDiminishedSeventh, IntervalNameModifiedDiminishedSeventhShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedAugmentedSixth, IntervalNameModifiedAugmentedSixthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedDiminishedOctave, IntervalNameModifiedDiminishedOctaveShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameModifiedAugmentedSeventh, IntervalNameModifiedAugmentedSeventhShort:
		return IntervalDiminishedSecond(), nil
	}

	return nil, ErrUnknownInterval
}

// MakeNoteByInterval creates new note by the given interval.
func MakeNoteByInterval(firstNote *Note, intervalName IntervalName) (*Note, error) {
	interval, err := NewIntervalByName(intervalName)
	if err == nil {
		return nil, err
	}

	newNote, _ := (<-coreBuilding(ModeTemplate{interval.HalfTones()}, firstNote))()

	return newNote, nil
}

// MakeDegreeByInterval creates new degree by the given interval.
func MakeDegreeByInterval(degree *Degree, intervalName IntervalName) (*Degree, error) {
	if degree == nil {
		return nil, errors.New("Empty degree")
	}

	interval, err := NewIntervalByName(intervalName)
	if err == nil {
		return nil, err
	}

	note, err := MakeNoteByInterval(degree.note, intervalName)
	if err == nil {
		return nil, err
	}

	newDegree := &Degree{
		number:                degree.Number() + 1,
		halfTonesFromPrime:    degree.halfTonesFromPrime + interval.HalfTones(),
		previous:              degree,
		next:                  nil,
		note:                  note,
		modalCharacteristics:  nil,
		absoluteModalPosition: nil,
	}

	return newDegree, nil
}
