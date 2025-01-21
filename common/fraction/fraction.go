package fraction

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// Fraction is the helper to describe unit and time signature.
type Fraction struct {
	Numerator   uint64
	Denominator uint64
}

// New creates new fraction with the given values.
func New(numerator, denominator uint64) *Fraction {
	return &Fraction{Numerator: numerator, Denominator: denominator}
}

// ErrInvalidFraction is returned when denominator is zero.
var ErrInvalidFraction = errors.New("invalid fraction")

// Validate checks that numerator amd denominator are not zeroes.
func (f *Fraction) Validate() error {
	if f.Denominator == 0 {
		return fmt.Errorf("fraction denominator cannot be zero. Numerator '%d', Denominator '%d': %w", f.Numerator, f.Denominator, ErrInvalidFraction)
	}

	return nil
}

// IzNotZero checks that numerator amd denominator are not zeroes.
func (f *Fraction) IzNotZero() bool {
	if f.Numerator == 0 || f.Denominator == 0 {
		return false
	}

	return true
}

// IsValid checks that denominator is not zero.
func (f *Fraction) IsValid() bool {
	return f.Denominator != 0
}

// IsReciprocalValid checks that numerator is not zero (i.e. reciprocal fraction will be valid).
func (f *Fraction) IsReciprocalValid() bool {
	return f.Numerator != 0
}

// Value returns decimal Value of the fraction.
func (f *Fraction) Value() (decimal.Decimal, error) {
	if f.Denominator == 0 {
		return decimal.Decimal{}, fmt.Errorf("cannot calculate value: denominator is zero: %w", ErrInvalidFraction)
	}

	return decimal.NewFromUint64(f.Numerator).Div(decimal.NewFromUint64(f.Denominator)), nil
}

// MustValue returns decimal Value of the fraction without error value. Panics in case of error.
func (f *Fraction) MustValue() decimal.Decimal {
	if f.Denominator == 0 {
		panic("cannot calculate value: denominator is zero")
	}

	return decimal.NewFromUint64(f.Numerator).Div(decimal.NewFromUint64(f.Denominator))
}

// ValueReciprocal returns decimal Value of the reciprocal fraction.
func (f *Fraction) ValueReciprocal() (decimal.Decimal, error) {
	if f.Numerator == 0 {
		return decimal.Decimal{}, fmt.Errorf("cannot calculate reciprocal value: numerator is zero: %w", ErrInvalidFraction)
	}

	return decimal.NewFromUint64(f.Denominator).Div(decimal.NewFromUint64(f.Numerator)), nil
}

// MustValueReciprocal returns decimal Value of the reciprocal fraction without error value. Panics in case of error.
func (f *Fraction) MustValueReciprocal() decimal.Decimal {
	if f.Numerator == 0 {
		panic("cannot calculate reciprocal value: numerator is zero")
	}

	return decimal.NewFromUint64(f.Denominator).Div(decimal.NewFromUint64(f.Numerator))
}
