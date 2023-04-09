package muse

import "fmt"

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
func NewIntervalChromatic(halfTones HalfTones) *ChromaticInterval {
	switch halfTones {
	case IntervalHalfTones0:
		return IntervalPerfectUnison()
	case IntervalHalfTones1:
		return IntervalMinorSecond()
	case IntervalHalfTones2:
		return IntervalMajorSecond()
	case IntervalHalfTones3:
		return IntervalMinorThird()
	case IntervalHalfTones4:
		return IntervalMajorThird()
	case IntervalHalfTones5:
		return IntervalPerfectFourth()
	case IntervalHalfTones6:
		return IntervalTritone()
	case IntervalHalfTones7:
		return IntervalPerfectFifth()
	case IntervalHalfTones8:
		return IntervalMinorSixth()
	case IntervalHalfTones9:
		return IntervalMajorSixth()
	case IntervalHalfTones10:
		return IntervalMinorSeventh()
	case IntervalHalfTones11:
		return IntervalMajorSeventh()
	case IntervalHalfTones12:
		return IntervalPerfectOctave()

	case IntervalHalfTones13:
		return IntervalMinorNinth()
	case IntervalHalfTones14:
		return IntervalMajorNinth()
	case IntervalHalfTones15:
		return IntervalMinorTenth()
	case IntervalHalfTones16:
		return IntervalMajorTenth()
	case IntervalHalfTones17:
		return IntervalPerfectEleventh()
	case IntervalHalfTones18:
		// no name for tritone after octave
	case IntervalHalfTones19:
		return IntervalPerfectTwelfth()
	case IntervalHalfTones20:
		return IntervalMinorThirteenth()
	case IntervalHalfTones21:
		return IntervalMajorThirteenth()
	case IntervalHalfTones22:
		return IntervalMinorFourteenth()
	case IntervalHalfTones23:
		return IntervalMajorFourteenth()
	case IntervalHalfTones24:
		return IntervalPerfectFifteenth()
	}

	return nil
}

// NewIntervalByHalfTonesAndDegrees creates interval by 1) half tones between degrees 2) amount of degrees between.
//
//nolint:gomnd
func NewIntervalByHalfTonesAndDegrees(halfTones HalfTones, degrees DegreeNum) *ChromaticInterval {
	switch halfTones {
	case IntervalHalfTones0:
		switch degrees {
		case 0:
			return NewIntervalChromatic(halfTones)
		case 1:
			return IntervalDiminishedSecond()
		}
	case IntervalHalfTones1:
		switch degrees {
		case 0:
			return IntervalAugmentedUnison()
		case 1:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones2:
		switch degrees {
		case 1:
			return NewIntervalChromatic(halfTones)
		case 2:
			return IntervalDiminishedThird()
		}
	case IntervalHalfTones3:
		switch degrees {
		case 1:
			return IntervalAugmentedSecond()
		case 2:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones4:
		switch degrees {
		case 2:
			return NewIntervalChromatic(halfTones)
		case 3:
			return IntervalDiminishedFourth()
		}
	case IntervalHalfTones5:
		switch degrees {
		case 2:
			return IntervalAugmentedThird()
		case 3:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones6:
		switch degrees {
		case 3:
			return IntervalAugmentedFourth()
		case 4:
			return IntervalDiminishedFifth()
		}
	case IntervalHalfTones7:
		switch degrees {
		case 4:
			return NewIntervalChromatic(halfTones)
		case 5:
			return IntervalDiminishedSixth()
		}
	case IntervalHalfTones8:
		switch degrees {
		case 4:
			return IntervalAugmentedFifth()
		case 5:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones9:
		switch degrees {
		case 5:
			return NewIntervalChromatic(halfTones)
		case 6:
			return IntervalDiminishedSeventh()
		}
	case IntervalHalfTones10:
		switch degrees {
		case 5:
			return IntervalAugmentedSixth()
		case 6:
			return NewIntervalChromatic(halfTones)
		}
	case IntervalHalfTones11:
		switch degrees {
		case 6:
			return NewIntervalChromatic(halfTones)
		case 7:
			return IntervalDiminishedOctave()
		}
	case IntervalHalfTones12:
		switch degrees {
		case 6:
			return IntervalAugmentedSeventh()
		case 7:
			return NewIntervalChromatic(halfTones)
		}
	}

	return nil
}

// Mode interval is determined by degrees and half tones between them.
// The function is needed in case of lack of information about the halftone distance of the steps from each other.
func newIntervalByDegreesAndHalfTones(halfTones HalfTones, degree1, degree2 *Degree) *Interval {
	degreesDiff := degree2.Number() - degree1.Number()

	return &Interval{
		ChromaticInterval: NewIntervalByHalfTonesAndDegrees(halfTones, degreesDiff),
		degree1:           degree1,
		degree2:           degree2,
	}
}

// NewIntervalByDegrees creates mode interval by degrees and half tones between them.
func NewIntervalByDegrees(degree1, degree2 *Degree) *Interval {
	halfTonesDiff := degree2.halfTonesFromPrime - degree1.halfTonesFromPrime

	return newIntervalByDegreesAndHalfTones(halfTonesDiff, degree1, degree2)
}

// NewIntervalByName returns new interval by it's name.
func NewIntervalByName(intervalName IntervalName) *ChromaticInterval {
	switch intervalName {
	// Chromatic intervals
	case IntervalNamePerfectUnison, IntervalNamePerfectUnisonShort:
		return IntervalPerfectUnison()
	case IntervalNameMinorSecond, IntervalNameMinorSecondShort:
		return IntervalMinorSecond()
	case IntervalNameMajorSecond, IntervalNameMajorSecondShort:
		return IntervalMajorSecond()
	case IntervalNameMinorThird, IntervalNameMinorThirdShort:
		return IntervalMinorThird()
	case IntervalNameMajorThird, IntervalNameMajorThirdShort:
		return IntervalMajorThird()
	case IntervalNamePerfectFourth, IntervalNamePerfectFourthShort:
		return IntervalPerfectFourth()
	case IntervalNameTritone, IntervalNameTritoneShort:
		return IntervalTritone()
	case IntervalNamePerfectFifth, IntervalNamePerfectFifthShort:
		return IntervalPerfectFifth()
	case IntervalNameMinorSixth, IntervalNameMinorSixthShort:
		return IntervalMinorSixth()
	case IntervalNameMajorSixth, IntervalNameMajorSixthShort:
		return IntervalMajorSixth()
	case IntervalNameMinorSeventh, IntervalNameMinorSeventhShort:
		return IntervalMinorSeventh()
	case IntervalNameMajorSeventh, IntervalNameMajorSeventhShort:
		return IntervalMajorSeventh()
	case IntervalNamePerfectOctave, IntervalNamePerfectOctaveShort:
		return IntervalPerfectOctave()

	// Diatonic intervals
	case IntervalNameModifiedDiminishedSecond, IntervalNameModifiedDiminishedSecondShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedAugmentedUnison, IntervalNameModifiedAugmentedUnisonShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedDiminishedThird, IntervalNameModifiedDiminishedThirdShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedAugmentedSecond, IntervalNameModifiedAugmentedSecondShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedDiminishedFourth, IntervalNameModifiedDiminishedFourthShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedAugmentedThird, IntervalNameModifiedAugmentedThirdShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedDiminishedFifth, IntervalNameModifiedDiminishedFifthShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedAugmentedFourth, IntervalNameModifiedAugmentedFourthShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedDiminishedSixth, IntervalNameModifiedDiminishedSixthShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedAugmentedFifth, IntervalNameModifiedAugmentedFifthShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedDiminishedSeventh, IntervalNameModifiedDiminishedSeventhShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedAugmentedSixth, IntervalNameModifiedAugmentedSixthShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedDiminishedOctave, IntervalNameModifiedDiminishedOctaveShort:
		return IntervalDiminishedSecond()

	case IntervalNameModifiedAugmentedSeventh, IntervalNameModifiedAugmentedSeventhShort:
		return IntervalDiminishedSecond()
	}

	return nil
}

// MakeNoteByInterval creates new note by the given interval.
func MakeNoteByInterval(firstNote *Note, intervalName IntervalName) *Note {
	// TODO: check interval for nil or error
	interval := NewIntervalByName(intervalName)

	newNote, _ := (<-coreBuilding(ModeTemplate{interval.HalfTones()}, firstNote))()

	return newNote
}

// MakeDegreeByInterval creates new degree by the given interval.
func MakeDegreeByInterval(degree *Degree, intervalName IntervalName) *Degree {
	if degree == nil {
		return nil
	}
	// TODO: check interval for nil or error
	interval := NewIntervalByName(intervalName)

	newDegree := &Degree{
		number:                degree.Number() + 1,
		halfTonesFromPrime:    degree.halfTonesFromPrime + interval.HalfTones(),
		previous:              degree,
		next:                  nil,
		note:                  MakeNoteByInterval(degree.note, intervalName),
		modalCharacteristics:  nil,
		absoluteModalPosition: nil,
	}

	return newDegree
}
