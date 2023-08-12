package muse

import "math"

type DurationName string

const (
	DurationNameLarge                = DurationName("Large")
	DurationNameLong                 = DurationName("Long")
	DurationNameDoubleWhole          = DurationName("DoubleWhole")
	DurationNameWhole                = DurationName("Whole")
	DurationNameHalf                 = DurationName("Half")
	DurationNameQuarter              = DurationName("Quarter")
	DurationNameEighth               = DurationName("Eighth")
	DurationNameSixteenth            = DurationName("Sixteenth")
	DurationNameThirtySecond         = DurationName("ThirtySecond")
	DurationNameSixtyFourth          = DurationName("SixtyFourth")
	DurationNameHundredTwentyEighth  = DurationName("HundredTwentyEighth")
	DurationNameTwoHundredFiftySixth = DurationName("TwoHundredFiftySixth")
	DurationNameFiveHundredTwelfth   = DurationName("FiveHundredTwelfth")
)

//nolint:gomnd
func (dn DurationName) getValue() int8 {
	switch dn {
	case DurationNameLarge:
		return 3
	case DurationNameLong:
		return 2
	case DurationNameDoubleWhole:
		return 1
	case DurationNameWhole:
		return 0
	case DurationNameHalf:
		return -1
	case DurationNameQuarter:
		return -2
	case DurationNameEighth:
		return -3
	case DurationNameSixteenth:
		return -4
	case DurationNameThirtySecond:
		return -5
	case DurationNameSixtyFourth:
		return -6
	case DurationNameHundredTwentyEighth:
		return -7
	case DurationNameTwoHundredFiftySixth:
		return -8
	case DurationNameFiveHundredTwelfth:
		return -9
	default:
		return 0
	}
}

// GetValue returns float value that represents note's duration as part of a bar.
//
//nolint:gomnd
func (dn DurationName) GetValue() float64 {
	return math.Pow(2, float64(dn.getValue())) // 2 means base to get real note value
}
