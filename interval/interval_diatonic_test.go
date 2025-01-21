package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
)

func TestNewIntervalByDegrees(t *testing.T) {
	type testCase struct {
		degree1, degree2 *degree.Degree
		want             *Diatonic
		err              error
	}

	makeTestCase := func(degree1, degree2 *degree.Degree, chromaticInterval *Chromatic, err error) testCase {
		return testCase{
			degree1: degree1,
			degree2: degree2,
			want: &Diatonic{
				chromaticInterval: chromaticInterval,
				degree1:           degree1,
				degree2:           degree2,
			},
			err: err,
		}
	}

	// TODO: make more cases
	testCases := []testCase{
		makeTestCase(newDegreeWithNumAndHalfTones(1, 0), newDegreeWithNumAndHalfTones(2, 1),
			&Chromatic{
				Sonance: SonanceMinorSecond,
				names: &nameExtended{
					name:      NameMinorSecond,
					shortName: NameMinorSecondShort,
				},
				halfTones: 1,
			}, nil),
		makeTestCase(newDegreeWithNumAndHalfTones(1, 0), newDegreeWithNumAndHalfTones(2, 2),
			&Chromatic{
				Sonance: SonanceMajorSecond,
				names: &nameExtended{
					name:      NameMajorSecond,
					shortName: NameMajorSecondShort,
				},
				halfTones: 2,
			}, nil),
		makeTestCase(newDegreeWithNumAndHalfTones(1, 0), newDegreeWithNumAndHalfTones(3, 2),
			&Chromatic{
				Sonance: SonanceMajorSecond,
				names: &nameExtended{
					name:      NameDiminishedThird,
					shortName: NameDiminishedThirdShort,
				},
				halfTones: 2,
			}, nil),
		makeTestCase(newDegreeWithNumAndHalfTones(1, 2), newDegreeWithNumAndHalfTones(2, 4),
			&Chromatic{
				Sonance: SonanceMajorSecond,
				names: &nameExtended{
					name:      NameMajorSecond,
					shortName: NameMajorSecondShort,
				},
				halfTones: 2,
			}, nil),
		makeTestCase(newDegreeWithNumAndHalfTones(3, 2), newDegreeWithNumAndHalfTones(2, 4),
			&Chromatic{
				Sonance: SonanceMajorSecond,
				names: &nameExtended{
					name:      NameMajorSecond,
					shortName: NameMajorSecondShort,
				},
				halfTones: 2,
			}, ErrIntervalUnknown),
	}

	var err error
	var interval *Diatonic
	for _, testCase := range testCases {
		interval, err = NewDiatonic(testCase.degree1, testCase.degree2)
		if err != nil {
			require.ErrorIs(t, testCase.err, err)
		} else {
			assert.Equal(t, testCase.want, interval)
		}
	}
}

func newDegreeWithNumAndHalfTones(
	number degree.Number,
	halfTonesFromPrime halftone.HalfTones,
) *degree.Degree {
	return degree.New(
		number,
		halfTonesFromPrime,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}
