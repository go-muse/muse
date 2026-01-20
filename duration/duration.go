package duration

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/tuplet"
)

const (
	// nanosecondsInMinute is amount of nanoseconds in minute.
	nanosecondsInMinute = uint64(time.Minute / time.Nanosecond)
)

// Relative (relative duration) is a set of characteristics determining how long a note sounds.
type Relative struct {
	name   Name
	dots   uint8
	tuplet tuplet.Tuplet
}

// NewRelative creates new Duration by the given relative duration name.
func NewRelative(name Name) Relative {
	return Relative{
		name:   name,
		dots:   0,
		tuplet: tuplet.Tuplet{},
	}
}

// Name returns the duration's name.
func (dr Relative) Name() Name {
	return dr.name
}

// Dots returns amount of the dots.
func (dr Relative) Dots() uint8 {
	return dr.dots
}

// Equal reports whether two relative durations are exactly equal,
// comparing name, dots, and tuplet.
func (dr Relative) Equal(other Relative) bool {
	return dr.name == other.name &&
		dr.dots == other.dots &&
		dr.tuplet.Equal(other.tuplet)
}

// AddDot increments amount of the dots and returns the duration.
func (dr Relative) AddDot() Relative {
	dr.dots++
	return dr
}

// SetDots sets amount of the dots and returns the duration.
func (dr Relative) SetDots(n uint8) Relative {
	dr.dots = n
	return dr
}

// RemoveDot decrements amount of the dots and returns the duration.
func (dr Relative) RemoveDot() Relative {
	if dr.dots > 0 {
		dr.dots--
	}

	return dr
}

// RemoveDots removes all dots and returns the duration.
func (dr Relative) RemoveDots() Relative {
	dr.dots = 0
	return dr
}

// Tuplet returns tuplet from the duration.
func (dr Relative) Tuplet() tuplet.Tuplet {
	return dr.tuplet
}

// SetTuplet sets the given tuplet for the duration and returns the duration.
func (dr Relative) SetTuplet(t tuplet.Tuplet) Relative {
	dr.tuplet = t
	return dr
}

// RemoveTuplet sets the tuplet from the duration and returns the duration.
func (dr Relative) RemoveTuplet() Relative {
	dr.tuplet = tuplet.Tuplet{}
	return dr
}

// SetTupletDuplet sets the duplet as tuplet for the duration and returns the duration.
func (dr Relative) SetTupletDuplet() Relative {
	dr.tuplet = dr.tuplet.SetDuplet()
	return dr
}

// SetTupletTriplet sets the triplet as tuplet for the duration and returns the duration.
func (dr Relative) SetTupletTriplet() Relative {
	dr.tuplet = dr.tuplet.SetTriplet()
	return dr
}

// GetTimeDuration calculates and returns time.Duration of the current duration.
func (dr Relative) GetTimeDuration(amountOfBars decimal.Decimal) time.Duration {
	if amountOfBars.LessThanOrEqual(decimal.Zero) {
		return 0
	}

	const baseValue = uint64(2)
	baseValueDecimal := decimal.NewFromUint64(baseValue)
	noteDurationDecimal := dr.Name().GetValue()
	minuteDecimal := decimal.NewFromUint64(nanosecondsInMinute)

	result := minuteDecimal.Mul(noteDurationDecimal).Div(amountOfBars)

	base := result
	for i := uint64(1); i <= uint64(dr.dots); i++ {
		add := base.Div(baseValueDecimal.Pow(decimal.NewFromUint64(i)))
		result = result.Add(add)
	}

	if dr.tuplet.IsSet() {
		// Multiplying before dividing gives a more accurate result than multiplying by the calculated fraction.
		result = result.Mul(decimal.NewFromUint64(dr.tuplet.M())).Div(decimal.NewFromUint64(dr.tuplet.N()))
	}

	return time.Duration(result.BigInt().Int64())
}

// GetPartOfBar returns duration as part of a bar by relative duration.
func (dr Relative) GetPartOfBar(timeSignature *fraction.Fraction) decimal.Decimal {
	if timeSignature == nil || !timeSignature.IsValid() {
		return decimal.Zero
	}

	const baseValue = uint64(2)
	baseValueDecimal := decimal.NewFromUint64(baseValue)
	noteDurationDecimal := dr.Name().GetValue()

	result := timeSignature.MustValue().Div(noteDurationDecimal)

	base := result
	for i := uint64(1); i <= uint64(dr.dots); i++ {
		add := base.Div(baseValueDecimal.Pow(decimal.NewFromUint64(i)))
		result = result.Add(add)
	}

	if dr.tuplet.IsSet() {
		// Multiplying before dividing gives a more accurate result than multiplying by the calculated fraction.
		result = result.Mul(decimal.NewFromUint64(dr.tuplet.N())).Div(decimal.NewFromUint64(dr.tuplet.M()))
	}

	return result
}
