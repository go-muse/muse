package interval

import (
	"errors"

	"github.com/go-muse/muse/builder"
	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

// Name is string type for interval's name.
type Name string

// nameExtended is a structure to extend name of interval with a short name.
type nameExtended struct {
	name      Name
	shortName Name
}

// Chromatic is the interval defined just by halftone, without degrees.
type Chromatic struct {
	Sonance
	names     *nameExtended
	halfTones halftone.HalfTones
}

// Name returns interval's name.
func (ic *Chromatic) Name() Name {
	if ic == nil {
		return ""
	}

	return ic.names.name
}

// ShortName returns interval's short name.
func (ic *Chromatic) ShortName() Name {
	if ic == nil {
		return ""
	}

	return ic.names.shortName
}

// HalfTones returns amount of halftone in the interval.
func (ic *Chromatic) HalfTones() halftone.HalfTones {
	if ic == nil {
		return 0
	}

	return ic.halfTones
}

// NewChromatic creates interval just by halftone between the notes
// Such interval is known as chromatic  interval or acoustic interval.
func NewChromatic(halfTones halftone.HalfTones) (*Chromatic, error) {
	switch halfTones {
	case HalfTones0:
		return PerfectUnison(), nil
	case HalfTones1:
		return MinorSecond(), nil
	case HalfTones2:
		return MajorSecond(), nil
	case HalfTones3:
		return MinorThird(), nil
	case HalfTones4:
		return MajorThird(), nil
	case HalfTones5:
		return PerfectFourth(), nil
	case HalfTones6:
		return Tritone(), nil
	case HalfTones7:
		return PerfectFifth(), nil
	case HalfTones8:
		return MinorSixth(), nil
	case HalfTones9:
		return MajorSixth(), nil
	case HalfTones10:
		return MinorSeventh(), nil
	case HalfTones11:
		return MajorSeventh(), nil
	case HalfTones12:
		return PerfectOctave(), nil

	case HalfTones13:
		return MinorNinth(), nil
	case HalfTones14:
		return MajorNinth(), nil
	case HalfTones15:
		return MinorTenth(), nil
	case HalfTones16:
		return MajorTenth(), nil
	case HalfTones17:
		return PerfectEleventh(), nil
	case HalfTones18:
		return OctaveWithTritone(), nil
	case HalfTones19:
		return PerfectTwelfth(), nil
	case HalfTones20:
		return MinorThirteenth(), nil
	case HalfTones21:
		return MajorThirteenth(), nil
	case HalfTones22:
		return MinorFourteenth(), nil
	case HalfTones23:
		return MajorFourteenth(), nil
	case HalfTones24:
		return PerfectFifteenth(), nil
	}

	return nil, ErrIntervalUnknown
}

// ErrIntervalUnknown means that it is impossible to determine the interval based on the available data.
var ErrIntervalUnknown = errors.New("unknown interval")

// NewIntervalByHalfTonesAndDegrees creates interval by 1) halftone between degrees 2) amount of degrees between.
//
//nolint:mnd
func NewIntervalByHalfTonesAndDegrees(halfTones halftone.HalfTones, degrees degree.Number) (*Chromatic, error) {
	switch halfTones {
	case HalfTones0:
		switch degrees {
		case 0:
			return NewChromatic(halfTones)
		case 1:
			return DiminishedSecond(), nil
		}
	case HalfTones1:
		switch degrees {
		case 0:
			return AugmentedUnison(), nil
		case 1:
			return NewChromatic(halfTones)
		}
	case HalfTones2:
		switch degrees {
		case 1:
			return NewChromatic(halfTones)
		case 2:
			return DiminishedThird(), nil
		}
	case HalfTones3:
		switch degrees {
		case 1:
			return AugmentedSecond(), nil
		case 2:
			return NewChromatic(halfTones)
		}
	case HalfTones4:
		switch degrees {
		case 2:
			return NewChromatic(halfTones)
		case 3:
			return DiminishedFourth(), nil
		}
	case HalfTones5:
		switch degrees {
		case 2:
			return AugmentedThird(), nil
		case 3:
			return NewChromatic(halfTones)
		}
	case HalfTones6:
		switch degrees {
		case 3:
			return AugmentedFourth(), nil
		case 4:
			return DiminishedFifth(), nil
		}
	case HalfTones7:
		switch degrees {
		case 4:
			return NewChromatic(halfTones)
		case 5:
			return DiminishedSixth(), nil
		}
	case HalfTones8:
		switch degrees {
		case 4:
			return AugmentedFifth(), nil
		case 5:
			return NewChromatic(halfTones)
		}
	case HalfTones9:
		switch degrees {
		case 5:
			return NewChromatic(halfTones)
		case 6:
			return DiminishedSeventh(), nil
		}
	case HalfTones10:
		switch degrees {
		case 5:
			return AugmentedSixth(), nil
		case 6:
			return NewChromatic(halfTones)
		}
	case HalfTones11:
		switch degrees {
		case 6:
			return NewChromatic(halfTones)
		case 7:
			return DiminishedOctave(), nil
		}
	case HalfTones12:
		switch degrees {
		case 6:
			return AugmentedSeventh(), nil
		case 7:
			return NewChromatic(halfTones)
		}
	}

	return nil, ErrIntervalUnknown
}

// MustNewIntervalByHalfTonesAndDegrees creates interval by 1) halftone between degrees 2) amount of degrees between.
// The function panics in case of unknown interval.
func MustNewIntervalByHalfTonesAndDegrees(halfTones halftone.HalfTones, degrees degree.Number) *Chromatic {
	chromaticInterval, err := NewIntervalByHalfTonesAndDegrees(halfTones, degrees)
	if err != nil {
		panic(err)
	}

	return chromaticInterval
}

// NewIntervalByName returns new interval by its name.
func NewIntervalByName(intervalName Name) (*Chromatic, error) {
	switch intervalName {
	// Chromatic intervals
	case NamePerfectUnison, NamePerfectUnisonShort:
		return PerfectUnison(), nil
	case NameMinorSecond, NameMinorSecondShort:
		return MinorSecond(), nil
	case NameMajorSecond, NameMajorSecondShort:
		return MajorSecond(), nil
	case NameMinorThird, NameMinorThirdShort:
		return MinorThird(), nil
	case NameMajorThird, NameMajorThirdShort:
		return MajorThird(), nil
	case NamePerfectFourth, NamePerfectFourthShort:
		return PerfectFourth(), nil
	case NameTritone, NameTritoneShort:
		return Tritone(), nil
	case NamePerfectFifth, NamePerfectFifthShort:
		return PerfectFifth(), nil
	case NameMinorSixth, NameMinorSixthShort:
		return MinorSixth(), nil
	case NameMajorSixth, NameMajorSixthShort:
		return MajorSixth(), nil
	case NameMinorSeventh, NameMinorSeventhShort:
		return MinorSeventh(), nil
	case NameMajorSeventh, NameMajorSeventhShort:
		return MajorSeventh(), nil
	case NamePerfectOctave, NamePerfectOctaveShort:
		return PerfectOctave(), nil

	// Diatonic intervals
	case NameDiminishedSecond, NameDiminishedSecondShort:
		return DiminishedSecond(), nil

	case NameAugmentedUnison, NameAugmentedUnisonShort:
		return DiminishedSecond(), nil

	case NameDiminishedThird, NameDiminishedThirdShort:
		return DiminishedSecond(), nil

	case NameAugmentedSecond, NameAugmentedSecondShort:
		return DiminishedSecond(), nil

	case NameDiminishedFourth, NameDiminishedFourthShort:
		return DiminishedSecond(), nil

	case NameAugmentedThird, NameAugmentedThirdShort:
		return DiminishedSecond(), nil

	case NameDiminishedFifth, NameDiminishedFifthShort:
		return DiminishedSecond(), nil

	case NameAugmentedFourth, NameAugmentedFourthShort:
		return DiminishedSecond(), nil

	case NameDiminishedSixth, NameDiminishedSixthShort:
		return DiminishedSecond(), nil

	case NameAugmentedFifth, NameAugmentedFifthShort:
		return DiminishedSecond(), nil

	case NameDiminishedSeventh, NameDiminishedSeventhShort:
		return DiminishedSecond(), nil

	case NameAugmentedSixth, NameAugmentedSixthShort:
		return DiminishedSecond(), nil

	case NameDiminishedOctave, NameDiminishedOctaveShort:
		return DiminishedSecond(), nil

	case NameAugmentedSeventh, NameAugmentedSeventhShort:
		return DiminishedSecond(), nil
	}

	return nil, ErrIntervalUnknown
}

// MakeNoteByName creates new note by the given interval name.
func MakeNoteByName(firstNote *note.Note, intervalName Name) (*note.Note, error) {
	interval, err := NewIntervalByName(intervalName)
	if err != nil {
		return nil, err
	}

	b := builder.NewBuilderCommon(halftone.Template{interval.HalfTones()}, firstNote)
	newNote, _ := (<-b)()

	return newNote, nil
}

// ErrDegreeEmpty means nil degree was given as a parameter.
var ErrDegreeEmpty = errors.New("empty degree")

// MakeDegreeByName creates new degree by the given interval name.
func MakeDegreeByName(d *degree.Degree, intervalName Name) (*degree.Degree, error) {
	if d == nil {
		return nil, ErrDegreeEmpty
	}

	interval, err := NewIntervalByName(intervalName)
	if err != nil {
		return nil, err
	}

	note, err := MakeNoteByName(d.Note(), intervalName)
	if err != nil {
		return nil, err
	}

	newDegree := degree.New(
		d.Number()+1,
		d.HalfTonesFromPrime()+interval.HalfTones(),
		nil,
		nil,
		note,
		nil,
		nil,
	)

	return newDegree, nil
}
