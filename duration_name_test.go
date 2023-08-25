package muse

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDurationName_getValue(t *testing.T) {
	type testCase struct {
		durationName DurationName
		want         decimal.Decimal
	}

	testCases := []testCase{
		{DurationNameLarge, decimal.NewFromInt(3)},
		{DurationNameLong, decimal.NewFromInt(2)},
		{DurationNameDoubleWhole, decimal.NewFromInt(1)},
		{DurationNameWhole, decimal.NewFromInt(0)},
		{DurationNameHalf, decimal.NewFromInt(-1)},
		{DurationNameQuarter, decimal.NewFromInt(-2)},
		{DurationNameEighth, decimal.NewFromInt(-3)},
		{DurationNameSixteenth, decimal.NewFromInt(-4)},
		{DurationNameThirtySecond, decimal.NewFromInt(-5)},
		{DurationNameSixtyFourth, decimal.NewFromInt(-6)},
		{DurationNameHundredTwentyEighth, decimal.NewFromInt(-7)},
		{DurationNameTwoHundredFiftySixth, decimal.NewFromInt(-8)},
		{DurationNameFiveHundredTwelfth, decimal.NewFromInt(-9)},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.durationName.getValue())
	}
}

func TestDurationName_GetValue(t *testing.T) {
	type testCase struct {
		durationName DurationName
		pow          decimal.Decimal
	}

	testCases := []testCase{
		{DurationNameLarge, decimal.NewFromInt(3)},
		{DurationNameLong, decimal.NewFromInt(2)},
		{DurationNameDoubleWhole, decimal.NewFromInt(1)},
		{DurationNameWhole, decimal.NewFromInt(0)},
		{DurationNameHalf, decimal.NewFromInt(-1)},
		{DurationNameQuarter, decimal.NewFromInt(-2)},
		{DurationNameEighth, decimal.NewFromInt(-3)},
		{DurationNameSixteenth, decimal.NewFromInt(-4)},
		{DurationNameThirtySecond, decimal.NewFromInt(-5)},
		{DurationNameSixtyFourth, decimal.NewFromInt(-6)},
		{DurationNameHundredTwentyEighth, decimal.NewFromInt(-7)},
		{DurationNameTwoHundredFiftySixth, decimal.NewFromInt(-8)},
		{DurationNameFiveHundredTwelfth, decimal.NewFromInt(-9)},
	}

	for _, testCase := range testCases {
		assert.Equal(t, decimal.NewFromInt(2).Pow(testCase.pow), testCase.durationName.GetValue())
	}
}
