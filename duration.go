package muse

import (
	"time"

	"github.com/shopspring/decimal"
)

// Duration is a set of characteristics determining how long a note sounds.
type Duration struct {
	name   DurationName
	dots   uint8
	tuplet *Tuplet
}

// Fraction is the helper to describe unit and time signature.
type Fraction struct {
	Numerator   uint64
	Denominator uint64
}

// minute is amount of nanoseconds in minute.
const minute = int64(60000000000)

// NewDuration creates new Duration by the given duration name.
func NewDuration(durationName DurationName) *Duration {
	return &Duration{
		name:   durationName,
		dots:   0,
		tuplet: nil,
	}
}

// Name returns the duration's name.
func (d *Duration) Name() DurationName {
	return d.name
}

// Name returns amount od the dots.
func (d *Duration) Dots() uint8 {
	return d.dots
}

// AddDot increments amount of the dots and returns the duration.
func (d *Duration) AddDot() *Duration {
	d.dots++

	return d
}

// SetDots sets amount of the dots and returns the duration.
func (d *Duration) SetDots(n uint8) *Duration {
	d.dots = n

	return d
}

// RemoveDot decrements amount of the dots and returns the duration.
func (d *Duration) RemoveDot() *Duration {
	d.dots--

	return d
}

// RemoveDots removes all dots and returns the duration.
func (d *Duration) RemoveDots() *Duration {
	d.dots = 0

	return d
}

// SetTuplet sets the given tuplet for the duration and returns the duration.
func (d *Duration) SetTuplet(t *Tuplet) *Duration {
	d.tuplet = t

	return d
}

// SetTuplet sets the tuplet from the duration and returns the duration.
func (d *Duration) RemoveTuplet() *Duration {
	d.tuplet = nil

	return d
}

// SetTuplet sets the duplet as tuplet for the duration and returns the duration.
//
//nolint:gomnd
func (d *Duration) SetTupletDuplet() *Duration {
	if d.tuplet != nil {
		d.tuplet.SetDuplet()
	} else {
		d.tuplet = NewTuplet(2, 3)
	}

	return d
}

// SetTuplet sets the triplet as tuplet for the duration and returns the duration.
//
//nolint:gomnd
func (d *Duration) SetTupletTriplet() *Duration {
	if d.tuplet != nil {
		d.tuplet.SetTriplet()
	} else {
		d.tuplet = NewTuplet(3, 2)
	}

	return d
}

// TimeDuration calculates and returns time.Duration of the current duration.
func (d *Duration) TimeDuration(bpm uint64, unit, timeSignature *Fraction) time.Duration {
	const baseValue = 2
	noteDurationDecimal := decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(int64(d.name.getValue())))
	bpmDecimal := decimal.NewFromInt(int64(bpm))
	unitDecimal := decimal.NewFromInt(int64(unit.Numerator)).Div(decimal.NewFromInt(int64(unit.Denominator)))
	timeSignatureDecimal := decimal.NewFromInt(int64(timeSignature.Numerator)).Div(decimal.NewFromInt(int64(timeSignature.Denominator)))
	minuteDecimal := decimal.NewFromInt(minute)

	amountOfBars := bpmDecimal.Mul(unitDecimal).Div(timeSignatureDecimal)

	result := minuteDecimal.Div(amountOfBars).Mul(noteDurationDecimal)

	base := result
	for i := int64(1); i <= int64(d.dots); i++ {
		add := base.Div(decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(i)))
		result = result.Add(add)
	}

	if d.tuplet != nil {
		result = result.Mul(decimal.NewFromInt(int64(d.tuplet.n))).Div(decimal.NewFromInt(int64(d.tuplet.m)))
	}

	return time.Duration(result.BigInt().Int64())
}
