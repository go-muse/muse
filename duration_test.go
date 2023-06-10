package muse

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewDuration(t *testing.T) {
	testCases := []struct {
		name DurationName
		want *Duration
	}{
		{name: DurationNameLarge, want: &Duration{0, relativeDuration{name: DurationNameLarge, dots: 0, tuplet: nil}}},
		{name: DurationNameLong, want: &Duration{0, relativeDuration{name: DurationNameLong, dots: 0, tuplet: nil}}},
		{name: DurationNameDoubleWhole, want: &Duration{0, relativeDuration{name: DurationNameDoubleWhole, dots: 0, tuplet: nil}}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, NewDurationWithRelativeValue(testCase.name))
	}
}

func TestDuration_GetAmountOfBars(t *testing.T) {
	type (
		args struct {
			bpm           uint64
			unit          *Fraction
			timeSignature *Fraction
		}
		want struct {
			amountOfBars decimal.Decimal
		}
	)

	testCases := []struct {
		args args
		want want
	}{
		{
			args: args{
				bpm:           60,              // 60 bpm
				unit:          &Fraction{1, 1}, // unit 1
				timeSignature: &Fraction{1, 1}, // time signature 1/1
			},
			want: want{
				amountOfBars: decimal.NewFromInt(60), // (60bpm*(1/1)/(1/1)) = 60 bars
			},
		},
		{
			args: args{
				bpm:           60,              // 60 bpm
				unit:          &Fraction{1, 2}, // unit 1/2
				timeSignature: &Fraction{3, 4}, // time signature 3/4
			},
			want: want{
				amountOfBars: decimal.NewFromInt(40), // (60bpm*(1/2)/(3/4)) = 40 bars
			},
		},
		{
			args: args{
				bpm:           120,             // 120 bpm
				unit:          &Fraction{1, 2}, // unit 1/2
				timeSignature: &Fraction{4, 4}, // time signature 4/4
			},
			want: want{
				amountOfBars: decimal.NewFromInt(60), // (120bpm*(1/2)/(4/4)) = 60 bars
			},
		},
		{
			args: args{
				bpm:           140,             // 140 bpm
				unit:          &Fraction{1, 2}, // unit 1/2
				timeSignature: &Fraction{3, 4}, // time signature 3/4
			},
			want: want{
				amountOfBars: decimal.NewFromInt(280).Div(decimal.NewFromInt(3)), // (140bpm*(1/2)/(3/4)) = 280/3 = 93,(3) bars
			},
		},
	}

	var amountOfBars decimal.Decimal
	for _, testCase := range testCases {
		amountOfBars = GetAmountOfBars(TrackSettings{testCase.args.bpm, *testCase.args.unit, *testCase.args.timeSignature})
		assert.True(t, testCase.want.amountOfBars.Equal(amountOfBars), "expected: %v, actual: %v, args: %d, %+v, %+v", testCase.want.amountOfBars.BigFloat(), amountOfBars.BigFloat(), testCase.args.bpm, testCase.args.unit, testCase.args.timeSignature)
	}
}

func TestDuration_TimeDuration(t *testing.T) {
	type (
		args struct {
			*TrackSettings
			*Duration
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
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 1}, // unit 1
					TimeSignature: Fraction{1, 1}, // time signature 1/1
				},
				NewDurationWithRelativeValue(DurationNameWhole), // whole note without dots
			},
			want: want{
				duration: time.Second, // ( 1 min / ((60*1/1)/(1/1)) ) * 1 = (1 min / 60) * 1 = 1sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 2}, // unit 1/2
					TimeSignature: Fraction{3, 4}, // time signature 3/4
				},
				NewDurationWithRelativeValue(DurationNameHalf), // half note without dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(750), // ( 1 min / ((60*1/2)/(3/4)) ) * 1/2 = (1 min / 40) * 1/2 = 0,75sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           120,            // 120 bpm
					Unit:          Fraction{1, 2}, // unit 1/2
					TimeSignature: Fraction{4, 4}, // time signature 4/4
				},
				NewDurationWithRelativeValue(DurationNameQuarter), // quarter note without dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(250), // ( 1 min / ((120*1/2)/(4/4)) ) * 1/4 = (1 min / 60) * 1/4 = 0,25sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 1}, // unit 1
					TimeSignature: Fraction{1, 1}, // time signature 1
				},
				NewDurationWithRelativeValue(DurationNameWhole).AddDot(), // whole note with dot
			},
			want: want{
				duration: time.Millisecond * time.Duration(1500), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with dot = 1,5 sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 1}, // unit 1
					TimeSignature: Fraction{1, 1}, // time signature 1
				},
				NewDurationWithRelativeValue(DurationNameWhole).SetDots(2), // whole note with two dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(1750), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with double dot = 1,75 sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 1}, // unit 1
					TimeSignature: Fraction{1, 1}, // time signature 1
				}, NewDurationWithRelativeValue(DurationNameWhole).SetDots(4), // whole note  with four dots
			},
			want: want{
				duration: time.Microsecond * time.Duration(1937500), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with 4 dots = 1,9375 sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 1}, // unit 1
					TimeSignature: Fraction{1, 1}, // time signature 1
				}, NewDurationWithRelativeValue(DurationNameWhole).SetTuplet(NewTuplet(2, 3)), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(2) / time.Duration(3), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 1}, // unit 1
					TimeSignature: Fraction{1, 1}, // time signature 1
				},
				NewDurationWithRelativeValue(DurationNameWhole).SetTupletDuplet(), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(2) / time.Duration(3), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           60,             // 60 bpm
					Unit:          Fraction{1, 1}, // unit 1
					TimeSignature: Fraction{1, 1}, // time signature 1
				},
				NewDurationWithRelativeValue(DurationNameWhole).SetTupletTriplet(), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(3) / time.Duration(2), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				&TrackSettings{
					BPM:           140,            // 140 bpm
					Unit:          Fraction{1, 2}, // unit 1/2
					TimeSignature: Fraction{3, 4}, // time signature 3/4
				},
				NewDurationWithRelativeValue(DurationNameWhole).SetTupletTriplet(), // whole note with tuplet 3/2
			},
			want: want{
				duration: time.Second * time.Duration(964285714) / time.Duration(1000000000), // (( 1 min / ((140*1/2)/(3/4)) ) * 1) * 3/2 = ((1 min / 93.(3)) * 1) * 3/2 ≈ 0,642857143 * 3/2 ≈ 0,964285714 sec
			},
		},
	}

	var duration time.Duration
	for _, testCase := range testCases {
		duration = testCase.args.Duration.GetTimeDuration(*testCase.args.TrackSettings)
		assert.Equal(t, testCase.want.duration, duration)
	}
}

func TestDuration_RemoveTuplet(t *testing.T) {
	duration := &Duration{
		absoluteDuration: 0,
		relativeDuration: relativeDuration{
			name: "",
			dots: 0,
			tuplet: &Tuplet{
				n: 2,
				m: 3,
			},
		},
	}

	duration.RemoveTuplet()

	assert.Nil(t, duration.Tuplet())

	var duration2 *Duration
	assert.Nil(t, duration2.Tuplet())
}

func TestDuration_RemoveDot(t *testing.T) {
	dots := uint8(3)
	duration := &Duration{
		relativeDuration: relativeDuration{
			dots: dots,
		},
	}

	assert.Equal(t, dots-1, duration.RemoveDot().Dots())

	var duration2 *Duration
	assert.Zero(t, duration2.RemoveDot().Dots())
}

func TestDuration_RemoveDots(t *testing.T) {
	duration := &Duration{
		relativeDuration: relativeDuration{
			dots: 3,
		},
	}

	assert.Zero(t, duration.RemoveDots().Dots())

	var duration2 *Duration
	assert.Zero(t, duration2.RemoveDots().Dots())
}

func TestDuration_GetPartOfBarByRelative(t *testing.T) {
	testCases := []struct {
		duration      *Duration
		timeSignature *Fraction
		want          decimal.Decimal
	}{
		{
			duration:      NewDurationWithRelativeValue(DurationNameWhole),
			timeSignature: &Fraction{1, 1},
			want:          decimal.NewFromInt(1),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameHalf),
			timeSignature: &Fraction{1, 2},
			want:          decimal.NewFromInt(1),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameWhole),
			timeSignature: &Fraction{1, 2},
			want:          decimal.NewFromFloat(0.5),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameHalf),
			timeSignature: &Fraction{1, 1},
			want:          decimal.NewFromInt(2),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameWhole).SetDots(1),
			timeSignature: &Fraction{1, 1},
			want:          decimal.NewFromFloat(1.5),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameWhole).SetDots(2),
			timeSignature: &Fraction{1, 1},
			want:          decimal.NewFromFloat(1.75),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameWhole).SetTupletDuplet(),
			timeSignature: &Fraction{1, 1},
			want:          decimal.NewFromFloat(0.6666666666666667),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameWhole).SetTupletTriplet(),
			timeSignature: &Fraction{1, 1},
			want:          decimal.NewFromFloat(1.5),
		},
		{
			duration:      NewDurationWithRelativeValue(DurationNameWhole).SetTupletTriplet().AddDot(),
			timeSignature: &Fraction{1, 1},
			want:          decimal.NewFromFloat(2.25),
		},
	}

	for _, testCase := range testCases {
		assert.True(t, testCase.want.Equal(testCase.duration.GetPartOfBarByRelative(testCase.timeSignature)), "expected: %+v, actual: %+v", testCase.want, testCase.duration.GetPartOfBarByRelative(testCase.timeSignature))
	}
}

func TestDuration_GetPartOfBarByAbsolute(t *testing.T) {
	testCases := []struct {
		*TrackSettings
		duration *Duration
		want     decimal.Decimal
	}{
		{
			&TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 1},
				TimeSignature: Fraction{1, 1},
			},
			NewDurationWithAbsoluteValue(time.Second),
			decimal.NewFromFloat(0.5),
		},
		{
			&TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 2},
				TimeSignature: Fraction{1, 1},
			},
			NewDurationWithAbsoluteValue(time.Second),
			decimal.NewFromInt(1),
		},
		{
			&TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 4},
				TimeSignature: Fraction{4, 4},
			},
			NewDurationWithAbsoluteValue(2*time.Second),
			decimal.NewFromInt(1),
		},
		{
			&TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 4},
				TimeSignature: Fraction{4, 4},
			},
			NewDurationWithAbsoluteValue(2*time.Second).SetDots(2),  // dots doesn't affect
			decimal.NewFromInt(1),
		},
		{
			&TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 4},
				TimeSignature: Fraction{4, 4},
			},
			NewDurationWithAbsoluteValue(2*time.Second).SetTupletDuplet(),  // tuplets doesn't affect
			decimal.NewFromInt(1),
		},
	}

	for _, testCase := range testCases {
		assert.True(t, testCase.want.Equal(testCase.duration.GetPartOfBarByAbsolute(*testCase.TrackSettings)), "expected: %+v, actual: %+v, settings: %+v", testCase.want, testCase.duration.GetPartOfBarByAbsolute(*testCase.TrackSettings), testCase.TrackSettings)
	}
}
