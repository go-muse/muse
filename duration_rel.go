package muse

import (
	"time"

	"github.com/shopspring/decimal"
)

// DurationRel (relative duration) is a set of characteristics determining how long a note sounds.
type DurationRel struct {
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
const (
	secondsInMinute = int64(60)
	minute          = int64(1000000000) * secondsInMinute
)

// NewDurationRel creates new Duration by the given relative duration name.
func NewDurationRel(durationName DurationName) *DurationRel {
	return &DurationRel{
		name:   durationName,
		dots:   0,
		tuplet: nil,
	}
}

// Name returns the duration's name.
func (dr *DurationRel) Name() DurationName {
	if dr == nil {
		return ""
	}

	return dr.name
}

// Dots returns amount of the dots.
func (dr *DurationRel) Dots() uint8 {
	if dr == nil {
		return 0
	}

	return dr.dots
}

// AddDot increments amount of the dots and returns the duration.
func (dr *DurationRel) AddDot() *DurationRel {
	if dr == nil {
		return dr
	}

	dr.dots++

	return dr
}

// SetDots sets amount of the dots and returns the duration.
func (dr *DurationRel) SetDots(n uint8) *DurationRel {
	if dr == nil {
		return dr
	}

	dr.dots = n

	return dr
}

// RemoveDot decrements amount of the dots and returns the duration.
func (dr *DurationRel) RemoveDot() *DurationRel {
	if dr == nil {
		return dr
	}

	if dr.dots > 0 {
		dr.dots--
	}

	return dr
}

// RemoveDots removes all dots and returns the duration.
func (dr *DurationRel) RemoveDots() *DurationRel {
	if dr == nil {
		return dr
	}

	dr.dots = 0

	return dr
}

// Tuplet returns tuplet from the duration.
func (dr *DurationRel) Tuplet() *Tuplet {
	if dr == nil {
		return nil
	}

	return dr.tuplet
}

// SetTuplet sets the given tuplet for the duration and returns the duration.
func (dr *DurationRel) SetTuplet(t *Tuplet) *DurationRel {
	if dr == nil {
		return dr
	}

	dr.tuplet = t

	return dr
}

// SetTuplet sets the tuplet from the duration and returns the duration.
func (dr *DurationRel) RemoveTuplet() *DurationRel {
	if dr == nil {
		return dr
	}

	dr.tuplet = nil

	return dr
}

// SetTuplet sets the duplet as tuplet for the duration and returns the duration.
//
//nolint:gomnd
func (dr *DurationRel) SetTupletDuplet() *DurationRel {
	if dr == nil {
		return dr
	}

	if dr.tuplet != nil {
		dr.tuplet.SetDuplet()
	} else {
		dr.tuplet = NewTuplet(2, 3)
	}

	return dr
}

// SetTuplet sets the triplet as tuplet for the duration and returns the duration.
//
//nolint:gomnd
func (dr *DurationRel) SetTupletTriplet() *DurationRel {
	if dr == nil {
		return dr
	}

	if dr.tuplet != nil {
		dr.tuplet.SetTriplet()
	} else {
		dr.tuplet = NewTuplet(3, 2)
	}

	return dr
}

// GetAmountOfBars calculates and returns amount of bars within one minute.
func GetAmountOfBars(trackSettings TrackSettings) decimal.Decimal {
	bpmDecimal := decimal.NewFromInt(int64(trackSettings.BPM))
	unitDecimal := decimal.NewFromInt(int64(trackSettings.Unit.Numerator)).Div(decimal.NewFromInt(int64(trackSettings.Unit.Denominator)))
	timeSignatureDecimal := decimal.NewFromInt(int64(trackSettings.TimeSignature.Numerator)).Div(decimal.NewFromInt(int64(trackSettings.TimeSignature.Denominator)))

	return bpmDecimal.Mul(unitDecimal).Div(timeSignatureDecimal)
}

// GetTimeDuration calculates and returns time.Duration of the current duration.
func (dr *DurationRel) GetTimeDuration(trackSettings TrackSettings) time.Duration {
	if dr == nil {
		return 0
	}

	amountOfBars := GetAmountOfBars(trackSettings)

	const baseValue = 2
	noteDurationDecimal := decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(int64(dr.name.getValue())))
	minuteDecimal := decimal.NewFromInt(minute)

	result := minuteDecimal.Div(amountOfBars).Mul(noteDurationDecimal)

	base := result
	for i := int64(1); i <= int64(dr.dots); i++ {
		add := base.Div(decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(i)))
		result = result.Add(add)
	}

	if dr.tuplet != nil {
		result = result.Mul(decimal.NewFromInt(int64(dr.tuplet.n))).Div(decimal.NewFromInt(int64(dr.tuplet.m)))
	}

	return time.Duration(result.BigInt().Int64())
}

// GetPartOfBar returns duration as part of a bar by relative duration.
func (dr *DurationRel) GetPartOfBar(timeSignature *Fraction) decimal.Decimal {
	if dr == nil {
		return decimal.Zero
	}

	const baseValue = 2
	noteDurationDecimal := decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(int64(dr.name.getValue())))
	timeSignatureDecimal := decimal.NewFromInt(int64(timeSignature.Numerator)).Div(decimal.NewFromInt(int64(timeSignature.Denominator)))

	result := timeSignatureDecimal.Div(noteDurationDecimal)

	base := result
	for i := int64(1); i <= int64(dr.dots); i++ {
		add := base.Div(decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(i)))
		result = result.Add(add)
	}

	if dr.tuplet != nil {
		result = result.Mul(decimal.NewFromInt(int64(dr.tuplet.n))).Div(decimal.NewFromInt(int64(dr.tuplet.m)))
	}

	return result
}
