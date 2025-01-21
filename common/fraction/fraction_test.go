package fraction

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFraction_MustValue(t *testing.T) {
	numerator, denominator := uint64(3), uint64(4)
	fraction := New(
		numerator,
		denominator,
	)

	assert.Equal(t, decimal.NewFromUint64(numerator).Div(decimal.NewFromUint64(denominator)), fraction.MustValue())
}
