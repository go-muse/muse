package muse

import (
	"github.com/shopspring/decimal"
)

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

// getValue returns the power to which the number 2 must be raised to obtain the part of the bar occupied by the note.
//
//nolint:gomnd
func (dn DurationName) getValue() decimal.Decimal {
	switch dn {
	case DurationNameLarge:
		return decimal.NewFromInt(3)
	case DurationNameLong:
		return decimal.NewFromInt(2)
	case DurationNameDoubleWhole:
		return decimal.NewFromInt(1)
	case DurationNameWhole:
		return decimal.NewFromInt(0)
	case DurationNameHalf:
		return decimal.NewFromInt(-1)
	case DurationNameQuarter:
		return decimal.NewFromInt(-2)
	case DurationNameEighth:
		return decimal.NewFromInt(-3)
	case DurationNameSixteenth:
		return decimal.NewFromInt(-4)
	case DurationNameThirtySecond:
		return decimal.NewFromInt(-5)
	case DurationNameSixtyFourth:
		return decimal.NewFromInt(-6)
	case DurationNameHundredTwentyEighth:
		return decimal.NewFromInt(-7)
	case DurationNameTwoHundredFiftySixth:
		return decimal.NewFromInt(-8)
	case DurationNameFiveHundredTwelfth:
		return decimal.NewFromInt(-9)
	default:
		return decimal.NewFromInt(0)
	}
}

// GetValue returns value that represents note's duration as part of a bar.
//
//nolint:gomnd
func (dn DurationName) GetValue() decimal.Decimal {
	const baseValue = 2
	return decimal.NewFromInt(baseValue).Pow(dn.getValue())
}
