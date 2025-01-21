package tuplet

import (
	"github.com/shopspring/decimal"
)

const (
	dupletNumerator   = 2
	dupletDenominator = 3

	tripletNumerator   = 3
	tripletDenominator = 2
)

// Tuplet is a characteristic of the duration of a given note,
// indicating how many notes of such a tuplet are equal in duration to how many notes outside the tuplet.
type Tuplet struct {
	m uint64
	n uint64
}

// New creates new tuplet with the given values as m/n.
// It means that "m" of such notes will last as "n" notes of the same duration outside the tuplet.
func New(m, n uint64) *Tuplet {
	return &Tuplet{
		m: m,
		n: n,
	}
}

// NewDuplet creates new duplet with values as 2/3.
func NewDuplet() *Tuplet {
	return &Tuplet{
		m: dupletNumerator,
		n: dupletDenominator,
	}
}

// NewTriplet creates new duplet with values as 3/2.
func NewTriplet() *Tuplet {
	return &Tuplet{
		m: tripletNumerator,
		n: tripletDenominator,
	}
}

// Value returns the tuplet ratio m:n as a decimal number.
func (t *Tuplet) Value() decimal.Decimal {
	return decimal.NewFromUint64(t.m).Div(decimal.NewFromUint64(t.n))
}

// ValueReciprocal returns the tuplet ratio n:m as a decimal number.
func (t *Tuplet) ValueReciprocal() decimal.Decimal {
	return decimal.NewFromUint64(t.n).Div(decimal.NewFromUint64(t.m))
}

// Set sets the setting that "m" of such notes will last as "n" notes of the same duration outside the tuplet.
func (t *Tuplet) Set(m, n uint64) *Tuplet {
	if t == nil {
		return New(m, n)
	}

	t.m = m
	t.n = n

	return t
}

// SetTriplet sets the setting that three such notes will last as long as two notes of the same duration outside the tuplet.
func (t *Tuplet) SetTriplet() *Tuplet {
	if t == nil {
		return New(tripletNumerator, tripletDenominator)
	}

	t.m = tripletNumerator
	t.n = tripletDenominator

	return t
}

// SetDuplet sets the setting that two such notes will last as long as three notes of the same duration outside the tuplet.
func (t *Tuplet) SetDuplet() *Tuplet {
	if t == nil {
		return New(dupletNumerator, dupletDenominator)
	}

	t.m = dupletNumerator
	t.n = dupletDenominator

	return t
}

// M is getter for m.
func (t *Tuplet) M() uint64 { return t.m }

// N is getter for n.
func (t *Tuplet) N() uint64 { return t.n }
