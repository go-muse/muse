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
	tuplet *tuplet.Tuplet
}

// NewRelative creates new Duration by the given relative duration name.
func NewRelative(name Name) *Relative {
	return &Relative{
		name:   name,
		dots:   0,
		tuplet: nil,
	}
}

// Name returns the duration's name.
func (dr *Relative) Name() Name {
	if dr == nil {
		return ""
	}

	return dr.name
}

// Dots returns amount of the dots.
func (dr *Relative) Dots() uint8 {
	if dr == nil {
		return 0
	}

	return dr.dots
}

// AddDot increments amount of the dots and returns the duration.
func (dr *Relative) AddDot() *Relative {
	if dr == nil {
		return dr
	}

	dr.dots++

	return dr
}

// SetDots sets amount of the dots and returns the duration.
func (dr *Relative) SetDots(n uint8) *Relative {
	if dr == nil {
		return dr
	}

	dr.dots = n

	return dr
}

// RemoveDot decrements amount of the dots and returns the duration.
func (dr *Relative) RemoveDot() *Relative {
	if dr == nil {
		return dr
	}

	if dr.dots > 0 {
		dr.dots--
	}

	return dr
}

// RemoveDots removes all dots and returns the duration.
func (dr *Relative) RemoveDots() *Relative {
	if dr == nil {
		return dr
	}

	dr.dots = 0

	return dr
}

// Tuplet returns tuplet from the duration.
func (dr *Relative) Tuplet() *tuplet.Tuplet {
	if dr == nil {
		return nil
	}

	return dr.tuplet
}

// SetTuplet sets the given tuplet for the duration and returns the duration.
func (dr *Relative) SetTuplet(t *tuplet.Tuplet) *Relative {
	if dr == nil {
		return dr
	}

	dr.tuplet = t

	return dr
}

// RemoveTuplet sets the tuplet from the duration and returns the duration.
func (dr *Relative) RemoveTuplet() *Relative {
	if dr == nil {
		return dr
	}

	dr.tuplet = nil

	return dr
}

// SetTupletDuplet sets the duplet as tuplet for the duration and returns the duration.
func (dr *Relative) SetTupletDuplet() *Relative {
	if dr == nil {
		return dr
	}

	if dr.tuplet != nil {
		dr.tuplet.SetDuplet()
	} else {
		dr.tuplet = tuplet.NewDuplet()
	}

	return dr
}

// SetTupletTriplet sets the triplet as tuplet for the duration and returns the duration.
func (dr *Relative) SetTupletTriplet() *Relative {
	if dr == nil {
		return dr
	}

	if dr.tuplet != nil {
		dr.tuplet.SetTriplet()
	} else {
		dr.tuplet = tuplet.NewTriplet()
	}

	return dr
}

// GetTimeDuration calculates and returns time.Duration of the current duration.
func (dr *Relative) GetTimeDuration(amountOfBars decimal.Decimal) time.Duration {
	if dr == nil || amountOfBars.LessThanOrEqual(decimal.Zero) {
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

	if dr.tuplet != nil {
		// Multiplying before dividing gives a more accurate result than multiplying by the calculated fraction.
		result = result.Mul(decimal.NewFromUint64(dr.tuplet.M())).Div(decimal.NewFromUint64(dr.tuplet.N()))
	}

	return time.Duration(result.BigInt().Int64())
}

// GetPartOfBar returns duration as part of a bar by relative duration.
func (dr *Relative) GetPartOfBar(timeSignature *fraction.Fraction) decimal.Decimal {
	if dr == nil || timeSignature == nil || !timeSignature.IsValid() {
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

	if dr.tuplet != nil {
		// Multiplying before dividing gives a more accurate result than multiplying by the calculated fraction.
		result = result.Mul(decimal.NewFromUint64(dr.tuplet.N())).Div(decimal.NewFromUint64(dr.tuplet.M()))
	}

	return result
}
