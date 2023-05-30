package muse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDuration(t *testing.T) {
	testCases := []struct {
		name DurationName
		want *Duration
	}{
		{name: DurationNameLarge, want: &Duration{name: DurationNameLarge, dots: 0, tuplet: nil}},
		{name: DurationNameLong, want: &Duration{name: DurationNameLong, dots: 0, tuplet: nil}},
		{name: DurationNameDoubleWhole, want: &Duration{name: DurationNameDoubleWhole, dots: 0, tuplet: nil}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, NewDuration(testCase.name))
	}
}

func TestDurationTimeDuration(t *testing.T) {
	type (
		args struct {
			bpm           uint64
			unit          *Fraction
			timeSignature *Fraction
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
				bpm:           60,                             // 60 bpm
				unit:          &Fraction{1, 1},                // unit 1
				timeSignature: &Fraction{1, 1},                // time signature 1
				Duration:      NewDuration(DurationNameWhole), // whole note without dots
			},
			want: want{
				duration: time.Second, // ( 1 min / ((60*1/1)/(1/1)) ) * 1 = (1 min / 60) * 1 = 1sec
			},
		},
		{
			args: args{
				bpm:           60,                            // 60 bpm
				unit:          &Fraction{1, 2},               // unit 1/2
				timeSignature: &Fraction{3, 4},               // time signature 3/4
				Duration:      NewDuration(DurationNameHalf), // half note without dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(750), // ( 1 min / ((60*1/2)/(3/4)) ) * 1/2 = (1 min / 40) * 1/2 = 0,75sec
			},
		},
		{
			args: args{
				bpm:           120,                              // 120 bpm
				unit:          &Fraction{1, 2},                  // unit 1/2
				timeSignature: &Fraction{4, 4},                  // time signature 4/4
				Duration:      NewDuration(DurationNameQuarter), // quarter note without dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(250), // ( 1 min / ((120*1/2)/(4/4)) ) * 1/4 = (1 min / 60) * 1/4 = 0,25sec
			},
		},
		{
			args: args{
				bpm:           60,                                      // 60 bpm
				unit:          &Fraction{1, 1},                         // unit 1
				timeSignature: &Fraction{1, 1},                         // time signature 1
				Duration:      NewDuration(DurationNameWhole).AddDot(), // whole note with dot
			},
			want: want{
				duration: time.Millisecond * time.Duration(1500), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with dot = 1,5 sec
			},
		},
		{
			args: args{
				bpm:           60,                                        // 60 bpm
				unit:          &Fraction{1, 1},                           // unit 1
				timeSignature: &Fraction{1, 1},                           // time signature 1
				Duration:      NewDuration(DurationNameWhole).SetDots(2), // whole note with two dots
			},
			want: want{
				duration: time.Millisecond * time.Duration(1750), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with double dot = 1,75 sec
			},
		},
		{
			args: args{
				bpm:           60,                                        // 60 bpm
				unit:          &Fraction{1, 1},                           // unit 1
				timeSignature: &Fraction{1, 1},                           // time signature 1
				Duration:      NewDuration(DurationNameWhole).SetDots(4), // whole note  with four dots
			},
			want: want{
				duration: time.Microsecond * time.Duration(1937500), // ( 1 min / ((60*1/1)/(1/1)) ) * 1) = (1 min / 60) * 1 = 1sec with 4 dots = 1,9375 sec
			},
		},
		{
			args: args{
				bpm:           60,                                                        // 60 bpm
				unit:          &Fraction{1, 1},                                           // unit 1
				timeSignature: &Fraction{1, 1},                                           // time signature 1
				Duration:      NewDuration(DurationNameWhole).SetTuplet(NewTuplet(2, 3)), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(2) / time.Duration(3), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				bpm:           60,                                               // 60 bpm
				unit:          &Fraction{1, 1},                                  // unit 1
				timeSignature: &Fraction{1, 1},                                  // time signature 1
				Duration:      NewDuration(DurationNameWhole).SetTupletDuplet(), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(2) / time.Duration(3), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				bpm:           60,                                                // 60 bpm
				unit:          &Fraction{1, 1},                                   // unit 1
				timeSignature: &Fraction{1, 1},                                   // time signature 1
				Duration:      NewDuration(DurationNameWhole).SetTupletTriplet(), // whole note with tuplet 2/3
			},
			want: want{
				duration: time.Second * time.Duration(3) / time.Duration(2), // (( 1 min / ((60*1/1)/(1/1)) ) * 1) * 2/3 = ((1 min / 60) * 1) * 2/3 = 1sec * 2/3 = 2/3 sec
			},
		},
		{
			args: args{
				bpm:           140,                                               // 140 bpm
				unit:          &Fraction{1, 2},                                   // unit 1/2
				timeSignature: &Fraction{3, 4},                                   // time signature 3/4
				Duration:      NewDuration(DurationNameWhole).SetTupletTriplet(), // whole note with tuplet 3/2
			},
			want: want{
				duration: time.Second * time.Duration(964285714) / time.Duration(1000000000), // (( 1 min / ((140*1/2)/(3/4)) ) * 1) * 3/2 = ((1 min / 93.(3)) * 1) * 3/2 ≈ 0,642857143 * 3/2 ≈ 0,964285714 sec
			},
		},
	}

	var duration time.Duration
	for _, testCase := range testCases {
		duration = testCase.args.Duration.TimeDuration(testCase.args.bpm, testCase.args.unit, testCase.args.timeSignature)
		assert.Equal(t, testCase.want.duration, duration)
	}
}