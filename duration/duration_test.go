package duration

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/tuplet"
)

func TestNewDuration(t *testing.T) {
	testCases := []struct {
		name Name
		want *Relative
	}{
		{name: NameLarge, want: &Relative{name: NameLarge, dots: 0, tuplet: nil}},
		{name: NameLong, want: &Relative{name: NameLong, dots: 0, tuplet: nil}},
		{name: NameDoubleWhole, want: &Relative{name: NameDoubleWhole, dots: 0, tuplet: nil}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, NewRelative(testCase.name))
	}
}

func TestDuration_Name(t *testing.T) {
	validName := NameEighth
	duration1 := &Relative{name: validName}
	assert.Equal(t, validName, duration1.Name())

	var duration2 *Relative
	assert.Equal(t, Name(""), duration2.Name())
}

func TestDuration_GetTimeDuration(t *testing.T) {
	type (
		args struct {
			amountOfBars decimal.Decimal
			*Relative
		}
		want struct {
			duration time.Duration
		}
	)

	testCases := []struct {
		args args
		want want
	}{
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameWhole), // whole note without dots
			},
			want: want{
				duration: time.Second, // ( 1 min / ((60*1/1)/(1/1)) ) * 1 = (1 min / 60) * 1 = 1sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(40),
				NewRelative(NameHalf), // half note without dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(750), // ( 1 min / ((60*1/2)/(3/4)) ) * 1/2 = (1 min / 40) * 1/2 = 0,75sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameQuarter), // quarter note without dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(250), // ( 1 min / ((120*1/2)/(4/4)) ) * 1/4 = (1 min / 60) * 1/4 = 0,25sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameWhole).AddDot(), // whole note with dot
			},
			want: want{
				duration: time.Millisecond * time.Duration(1500), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with dot = 1,5 sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameWhole).SetDots(2), // whole note with two dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(1750), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with double dot = 1,75 sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameWhole).SetDots(4), // whole note  with four dots
			},
			want: want{
				duration: time.Microsecond * time.Duration(1937500), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with 4 dots = 1,9375 sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameWhole).SetTuplet(tuplet.New(2, 3)), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(2) / time.Duration(3), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameWhole).SetTupletDuplet(), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(2) / time.Duration(3), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(60),
				NewRelative(NameWhole).SetTupletTriplet(), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(3) / time.Duration(2), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				decimal.NewFromFloat(93.33333333),
				NewRelative(NameWhole).SetTupletTriplet(), // whole note with tuplet 3/2
			},
			want: want{
				duration: time.Second * time.Duration(964285714) / time.Duration(1000000000), // (( 1 min / ((140*1/2)/(3/4)) ) * 1) * 3/2 = ((1 min / 93.(3)) * 1) * 3/2 ≈ 0,642857143 * 3/2 ≈ 0,964285714 sec
			},
		},
	}

	var duration time.Duration
	for _, testCase := range testCases {
		duration = testCase.args.Relative.GetTimeDuration(testCase.args.amountOfBars)
		assert.Equal(t, testCase.want.duration, duration)
	}
}

func TestDuration_RemoveTuplet(t *testing.T) {
	duration := &Relative{
		name:   "",
		dots:   0,
		tuplet: tuplet.New(2, 3),
	}

	duration.RemoveTuplet()

	assert.Nil(t, duration.Tuplet())

	var duration2 *Relative
	assert.Nil(t, duration2.Tuplet())
}

func TestDuration_RemoveDot(t *testing.T) {
	dots := uint8(3)
	duration := &Relative{
		dots: dots,
	}

	assert.Equal(t, dots-1, duration.RemoveDot().Dots())

	var duration2 *Relative
	assert.Zero(t, duration2.RemoveDot().Dots())
}

func TestDuration_RemoveDots(t *testing.T) {
	duration := &Relative{
		dots: 3,
	}

	assert.Zero(t, duration.RemoveDots().Dots())

	var duration2 *Relative
	assert.Zero(t, duration2.RemoveDots().Dots())
}

func TestDuration_GetPartOfBar(t *testing.T) {
	testCases := []struct {
		duration      *Relative
		timeSignature *fraction.Fraction
		want          decimal.Decimal
	}{
		{
			duration:      NewRelative(NameWhole),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromUint64(1),
		},
		{
			duration:      NewRelative(NameHalf),
			timeSignature: fraction.New(1, 2),
			want:          decimal.NewFromUint64(1),
		},
		{
			duration:      NewRelative(NameWhole),
			timeSignature: fraction.New(1, 2),
			want:          decimal.NewFromFloat(0.5),
		},
		{
			duration:      NewRelative(NameHalf),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromUint64(2),
		},
		{
			duration:      NewRelative(NameWhole).SetDots(1),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1.5),
		},
		{
			duration:      NewRelative(NameWhole).SetDots(2),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1.75),
		},
		{
			duration:      NewRelative(NameWhole).SetTupletDuplet(),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1.5),
		},
		{
			duration:      NewRelative(NameWhole).SetTupletDuplet().AddDot(),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(2.25),
		},
		{
			duration:      NewRelative(NameWhole).SetTupletTriplet(),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(0.6666666666666667),
		},
		{
			duration:      NewRelative(NameWhole).SetTupletTriplet().AddDot(),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1),
		},
	}

	for _, testCase := range testCases {
		assert.True(t, testCase.want.Equal(testCase.duration.GetPartOfBar(testCase.timeSignature)), "expected: %+v, actual: %+v", testCase.want, testCase.duration.GetPartOfBar(testCase.timeSignature))
	}
}

func TestDuration_SetTuplet(t *testing.T) {
	t.Run("Duration_SetTuplet: positive cases", func(t *testing.T) {
		testCases := []struct {
			duration *Relative
			tuplet   *tuplet.Tuplet
		}{
			{
				duration: &Relative{
					tuplet: nil,
				},
				tuplet: tuplet.New(2, 3),
			},
			{
				duration: &Relative{
					tuplet: tuplet.New(3, 2),
				},
				tuplet: tuplet.New(2, 3),
			},
			{
				duration: &Relative{
					tuplet: tuplet.New(2, 3),
				},
				tuplet: tuplet.New(2, 3),
			},
		}

		for _, testCase := range testCases {
			assert.Equal(t, testCase.tuplet, testCase.duration.SetTuplet(testCase.tuplet).tuplet)
		}
	})

	t.Run("Duration_SetTuplet: negative cases", func(t *testing.T) {
		var duration *Relative
		assert.Nil(t, duration.SetTuplet(tuplet.New(2, 3)))
	})
}

func TestDuration_SetTupletDuplet(t *testing.T) {
	t.Run("SetTupletDuplet: positive cases", func(t *testing.T) {
		testCases := []struct {
			duration *Relative
		}{
			{
				duration: &Relative{
					tuplet: nil,
				},
			},
			{
				duration: &Relative{
					tuplet: tuplet.New(3, 2),
				},
			},
			{
				duration: &Relative{
					tuplet: tuplet.New(2, 3),
				},
			},
		}

		for _, testCase := range testCases {
			assert.Equal(t, tuplet.New(2, 3), testCase.duration.SetTupletDuplet().tuplet)
		}
	})

	t.Run("SetTupletDuplet: negative cases", func(t *testing.T) {
		var duration *Relative
		assert.Nil(t, duration.SetTupletDuplet())
	})
}

func TestDuration_SetTupletTriplet(t *testing.T) {
	t.Run("SetTupletTriplet: positive cases", func(t *testing.T) {
		testCases := []struct {
			duration *Relative
		}{
			{
				duration: &Relative{
					tuplet: nil,
				},
			},
			{
				duration: &Relative{
					tuplet: tuplet.New(3, 2),
				},
			},
			{
				duration: &Relative{
					tuplet: tuplet.New(2, 3),
				},
			},
		}

		for _, testCase := range testCases {
			assert.Equal(t, tuplet.New(3, 2), testCase.duration.SetTupletTriplet().tuplet)
		}
	})

	t.Run("SetTupletTriplet: negative cases", func(t *testing.T) {
		var duration *Relative
		assert.Nil(t, duration.SetTupletTriplet())
	})
}
