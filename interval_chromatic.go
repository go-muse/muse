package muse

import "github.com/pkg/errors"

// IntervalName is string type for interval's name.
type IntervalName string

type intervalNameExtended struct {
	name      IntervalName
	shortName IntervalName
}

// IntervalChromatic is the interval defined just by halftones, without degrees.
type IntervalChromatic struct {
	Sonance
	names     *intervalNameExtended
	halfTones HalfTones
}

// Name returns interval's name.
func (ic *IntervalChromatic) Name() IntervalName {
	if ic == nil {
		return ""
	}

	return ic.names.name
}

// ShortName returns interval's short name.
func (ic *IntervalChromatic) ShortName() IntervalName {
	if ic == nil {
		return ""
	}

	return ic.names.shortName
}

// HalfTones returns amount of half tones in the interval.
func (ic *IntervalChromatic) HalfTones() HalfTones {
	if ic == nil {
		return 0
	}

	return ic.halfTones
}

// NewIntervalChromatic creates interval just by half tones between the notes
// Such interval is known as chromatic  interval or acoustic interval.
func NewIntervalChromatic(halfTones HalfTones) (*IntervalChromatic, error) {
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
		return IntervalOctaveWithTritone(), nil
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

	return nil, ErrIntervalUnknown
}

// ErrIntervalUnknown means that it is impossible to determine the interval based on the available data.
var ErrIntervalUnknown = errors.New("unknown interval")

// NewIntervalByHalfTonesAndDegrees creates interval by 1) half tones between degrees 2) amount of degrees between.
//
//nolint:gomnd
func NewIntervalByHalfTonesAndDegrees(halfTones HalfTones, degrees DegreeNum) (*IntervalChromatic, error) {
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

	return nil, ErrIntervalUnknown
}

// MustNewIntervalByHalfTonesAndDegrees creates interval by 1) half tones between degrees 2) amount of degrees between.
// The function panics in case of unknown interval.
func MustNewIntervalByHalfTonesAndDegrees(halfTones HalfTones, degrees DegreeNum) *IntervalChromatic {
	chromaticInterval, err := NewIntervalByHalfTonesAndDegrees(halfTones, degrees)
	if err != nil {
		panic(err)
	}

	return chromaticInterval
}

// NewIntervalByName returns new interval by it's name.
func NewIntervalByName(intervalName IntervalName) (*IntervalChromatic, error) {
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
	case IntervalNameDiminishedSecond, IntervalNameDiminishedSecondShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameAugmentedUnison, IntervalNameAugmentedUnisonShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameDiminishedThird, IntervalNameDiminishedThirdShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameAugmentedSecond, IntervalNameAugmentedSecondShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameDiminishedFourth, IntervalNameDiminishedFourthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameAugmentedThird, IntervalNameAugmentedThirdShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameDiminishedFifth, IntervalNameDiminishedFifthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameAugmentedFourth, IntervalNameAugmentedFourthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameDiminishedSixth, IntervalNameDiminishedSixthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameAugmentedFifth, IntervalNameAugmentedFifthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameDiminishedSeventh, IntervalNameDiminishedSeventhShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameAugmentedSixth, IntervalNameAugmentedSixthShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameDiminishedOctave, IntervalNameDiminishedOctaveShort:
		return IntervalDiminishedSecond(), nil

	case IntervalNameAugmentedSeventh, IntervalNameAugmentedSeventhShort:
		return IntervalDiminishedSecond(), nil
	}

	return nil, ErrIntervalUnknown
}

// MakeNoteByIntervalName creates new note by the given interval name.
func MakeNoteByIntervalName(firstNote *Note, intervalName IntervalName) (*Note, error) {
	interval, err := NewIntervalByName(intervalName)
	if err != nil {
		return nil, err
	}
	newNote, _ := (<-coreBuildingCommon(ModeTemplate{interval.HalfTones()}, firstNote))()

	return newNote, nil
}

// ErrDegreeEmpty means nil degree was given as a parameter.
var ErrDegreeEmpty = errors.New("empty degree")

// MakeDegreeByIntervalName creates new degree by the given interval name.
func MakeDegreeByIntervalName(degree *Degree, intervalName IntervalName) (*Degree, error) {
	if degree == nil {
		return nil, ErrDegreeEmpty
	}

	interval, err := NewIntervalByName(intervalName)
	if err != nil {
		return nil, err
	}

	note, err := MakeNoteByIntervalName(degree.note, intervalName)
	if err != nil {
		return nil, err
	}

	newDegree := &Degree{
		number:                degree.Number() + 1,
		halfTonesFromPrime:    degree.halfTonesFromPrime + interval.HalfTones(),
		previous:              nil,
		next:                  nil,
		note:                  note,
		modalCharacteristics:  nil,
		absoluteModalPosition: nil,
	}

	return newDegree, nil
}
