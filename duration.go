package muse

import (
	"time"

	"github.com/shopspring/decimal"
)

// Duration is a combination of two concepts - AbsoluteDuration and MusicalDuration.
type Duration struct {
	absoluteDuration time.Duration
	relativeDuration
}

// relativeDuration is a set of characteristics determining how long a note sounds.
type relativeDuration struct {
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

// NewDurationWithRelativeValue creates new Duration by the given relative duration name.
func NewDurationWithRelativeValue(durationName DurationName) *Duration {
	return &Duration{
		0,
		relativeDuration{
			name:   durationName,
			dots:   0,
			tuplet: nil,
		},
	}
}

// NewDurationWithAbsoluteValue creates new Duration by the given time.Duration.
func NewDurationWithAbsoluteValue(absoluteDuration time.Duration) *Duration {
	return &Duration{
		absoluteDuration,
		relativeDuration{},
	}
}

// Name returns the duration's name.
func (d *Duration) Name() DurationName {
	if d == nil {
		return ""
	}

	return d.relativeDuration.name
}

// Dots returns amount of the dots.
func (d *Duration) Dots() uint8 {
	if d == nil {
		return 0
	}

	return d.dots
}

// AddDot increments amount of the dots and returns the duration.
func (d *Duration) AddDot() *Duration {
	if d == nil {
		return d
	}

	d.dots++

	return d
}

// SetDots sets amount of the dots and returns the duration.
func (d *Duration) SetDots(n uint8) *Duration {
	if d == nil {
		return d
	}

	d.dots = n

	return d
}

// RemoveDot decrements amount of the dots and returns the duration.
func (d *Duration) RemoveDot() *Duration {
	if d == nil {
		return d
	}

	d.dots--

	return d
}

// RemoveDots removes all dots and returns the duration.
func (d *Duration) RemoveDots() *Duration {
	if d == nil {
		return d
	}

	d.dots = 0

	return d
}

// SetTuplet sets the given tuplet for the duration and returns the duration.
func (d *Duration) SetTuplet(t *Tuplet) *Duration {
	if d == nil {
		return d
	}

	d.tuplet = t

	return d
}

// SetTuplet sets the tuplet from the duration and returns the duration.
func (d *Duration) RemoveTuplet() *Duration {
	if d == nil {
		return d
	}

	d.tuplet = nil

	return d
}

// SetTuplet sets the duplet as tuplet for the duration and returns the duration.
//
//nolint:gomnd
func (d *Duration) SetTupletDuplet() *Duration {
	if d == nil {
		return d
	}

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
	if d == nil {
		return d
	}

	if d.tuplet != nil {
		d.tuplet.SetTriplet()
	} else {
		d.tuplet = NewTuplet(3, 2)
	}

	return d
}

// GetAmountOfBars calculates and returns amount of bars within one minute.
func GetAmountOfBars(trackSettings TrackSettings) decimal.Decimal {
	bpmDecimal := decimal.NewFromInt(int64(trackSettings.BPM))
	unitDecimal := decimal.NewFromInt(int64(trackSettings.Unit.Numerator)).Div(decimal.NewFromInt(int64(trackSettings.Unit.Denominator)))
	timeSignatureDecimal := decimal.NewFromInt(int64(trackSettings.TimeSignature.Numerator)).Div(decimal.NewFromInt(int64(trackSettings.TimeSignature.Denominator)))

	return bpmDecimal.Mul(unitDecimal).Div(timeSignatureDecimal)
}

// GetTimeDuration calculates and returns time.Duration of the current duration.
func (d *Duration) GetTimeDuration(trackSettings TrackSettings) time.Duration {
	if d == nil {
		return 0
	}

	amountOfBars := GetAmountOfBars(trackSettings)

	const baseValue = 2
	noteDurationDecimal := decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(int64(d.name.getValue())))
	minuteDecimal := decimal.NewFromInt(minute)

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

// GetPartOfBar returns duration as part of a bar by relative duration.
func (d *Duration) GetPartOfBarByRelative(timeSignature *Fraction) decimal.Decimal {
	if d == nil {
		return decimal.Zero
	}

	const baseValue = 2
	noteDurationDecimal := decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(int64(d.name.getValue())))
	timeSignatureDecimal := decimal.NewFromInt(int64(timeSignature.Numerator)).Div(decimal.NewFromInt(int64(timeSignature.Denominator)))

	result := timeSignatureDecimal.Div(noteDurationDecimal)

	base := result
	for i := int64(1); i <= int64(d.dots); i++ {
		add := base.Div(decimal.NewFromInt(baseValue).Pow(decimal.NewFromInt(i)))
		result = result.Add(add)
	}

	if d.tuplet != nil {
		result = result.Mul(decimal.NewFromInt(int64(d.tuplet.n))).Div(decimal.NewFromInt(int64(d.tuplet.m)))
	}

	return result
}

// GetPartOfBarByAbsolute returns duration as part of a bar by absolute duration.
func (d *Duration) GetPartOfBarByAbsolute(trackSettings TrackSettings) decimal.Decimal {
	if d == nil {
		return decimal.Zero
	}

	amountofBars := GetAmountOfBars(trackSettings)
	secondsInBar := decimal.NewFromInt(int64(time.Duration(secondsInMinute) * time.Second)).Div(amountofBars)

	result := secondsInBar.Div(decimal.NewFromFloat(float64(d.absoluteDuration)))

	return result
}

// Tuplet returns tuplet from the duration.
func (d *Duration) Tuplet() *Tuplet {
	if d == nil {
		return nil
	}

	return d.relativeDuration.tuplet
}
