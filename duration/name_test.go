package duration

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestName_GetValue(t *testing.T) {
	type testCase struct {
		Name Name
		pow  decimal.Decimal
	}

	testCases := []testCase{
		{NameLarge, decimal.NewFromInt(3)},
		{NameLong, decimal.NewFromInt(2)},
		{NameDoubleWhole, decimal.NewFromInt(1)},
		{NameWhole, decimal.NewFromInt(0)},
		{NameHalf, decimal.NewFromInt(-1)},
		{NameQuarter, decimal.NewFromInt(-2)},
		{NameEighth, decimal.NewFromInt(-3)},
		{NameSixteenth, decimal.NewFromInt(-4)},
		{NameThirtySecond, decimal.NewFromInt(-5)},
		{NameSixtyFourth, decimal.NewFromInt(-6)},
		{NameHundredTwentyEighth, decimal.NewFromInt(-7)},
		{NameTwoHundredFiftySixth, decimal.NewFromInt(-8)},
		{NameFiveHundredTwelfth, decimal.NewFromInt(-9)},
	}

	for _, testCase := range testCases {
		expected := decimal.NewFromInt(2).Pow(testCase.pow)
		assert.True(t, expected.Equal(testCase.Name.GetValue()), "note name '%s' expected '%s', actual '%s'", testCase.Name, expected, testCase.Name.GetValue())
	}
}
