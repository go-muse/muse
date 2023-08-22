package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntervalByDegrees(t *testing.T) {
	type testCase struct {
		degree1, degree2 *Degree
		want             *Interval
		err              error
	}

	makeTestCase := func(degree1, degree2 *Degree, chromaticInterval *IntervalChromatic, err error) testCase {
		return testCase{
			degree1: degree1,
			degree2: degree2,
			want: &Interval{
				chromaticInterval: chromaticInterval,
				degree1:           degree1,
				degree2:           degree2,
			},
			err: err,
		}
	}

	// TODO: make more cases
	testCases := []testCase{
		makeTestCase(&Degree{number: 1, halfTonesFromPrime: 0}, &Degree{number: 2, halfTonesFromPrime: 1},
			&IntervalChromatic{
				Sonance: IntervalSonanceMinorSecond,
				names: &intervalNameExtended{
					name:      IntervalNameMinorSecond,
					shortName: IntervalNameMinorSecondShort,
				},
				halfTones: 1,
			}, nil),
		makeTestCase(&Degree{number: 1, halfTonesFromPrime: 0}, &Degree{number: 2, halfTonesFromPrime: 2},
			&IntervalChromatic{
				Sonance: IntervalSonanceMajorSecond,
				names: &intervalNameExtended{
					name:      IntervalNameMajorSecond,
					shortName: IntervalNameMajorSecondShort,
				},
				halfTones: 2,
			}, nil),
		makeTestCase(&Degree{number: 1, halfTonesFromPrime: 0}, &Degree{number: 3, halfTonesFromPrime: 2},
			&IntervalChromatic{
				Sonance: IntervalSonanceMajorSecond,
				names: &intervalNameExtended{
					name:      IntervalNameDiminishedThird,
					shortName: IntervalNameDiminishedThirdShort,
				},
				halfTones: 2,
			}, nil),
		makeTestCase(&Degree{number: 1, halfTonesFromPrime: 2}, &Degree{number: 2, halfTonesFromPrime: 4},
			&IntervalChromatic{
				Sonance: IntervalSonanceMajorSecond,
				names: &intervalNameExtended{
					name:      IntervalNameMajorSecond,
					shortName: IntervalNameMajorSecondShort,
				},
				halfTones: 2,
			}, nil),
		makeTestCase(&Degree{number: 3, halfTonesFromPrime: 2}, &Degree{number: 2, halfTonesFromPrime: 4},
			&IntervalChromatic{
				Sonance: IntervalSonanceMajorSecond,
				names: &intervalNameExtended{
					name:      IntervalNameMajorSecond,
					shortName: IntervalNameMajorSecondShort,
				},
				halfTones: 2,
			}, ErrIntervalUnknown),
	}

	var err error
	var interval *Interval
	for _, testCase := range testCases {
		interval, err = NewIntervalByDegrees(testCase.degree1, testCase.degree2)
		if err != nil {
			assert.ErrorIs(t, testCase.err, err)
		} else {
			assert.Equal(t, testCase.want, interval)
		}
	}
}
