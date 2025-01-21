package duration

import (
	"github.com/shopspring/decimal"

	"github.com/go-muse/muse/tuplet"
)

// Name is a type for duration names.
type Name string

// Names is a type for set of duration names.
type Names []string

// GetTuplet returns the tuplet ratio m:n as a Tuplet by duration's name.
//
//nolint:mnd
func (n Name) GetTuplet() *tuplet.Tuplet {
	switch n {
	case NameLarge:
		return tuplet.New(8, 1)
	case NameLong:
		return tuplet.New(4, 1)
	case NameDoubleWhole:
		return tuplet.New(2, 1)
	case NameWhole:
		return tuplet.New(1, 1)
	case NameHalf:
		return tuplet.New(1, 2)
	case NameQuarter:
		return tuplet.New(1, 4)
	case NameEighth:
		return tuplet.New(1, 8)
	case NameSixteenth:
		return tuplet.New(1, 16)
	case NameThirtySecond:
		return tuplet.New(1, 32)
	case NameSixtyFourth:
		return tuplet.New(1, 64)
	case NameHundredTwentyEighth:
		return tuplet.New(1, 128)
	case NameTwoHundredFiftySixth:
		return tuplet.New(1, 256)
	case NameFiveHundredTwelfth:
		return tuplet.New(1, 512)
	default:
		return nil
	}
}

// GetValue returns Value that represents note's duration as part of a bar.
func (n Name) GetValue() decimal.Decimal {
	return n.GetTuplet().Value()
}

// NewDuration makes new relative duration by the given duration name.
func (n Name) NewDuration() *Relative {
	return NewRelative(n)
}
