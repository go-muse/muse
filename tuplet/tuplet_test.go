package tuplet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuplet(t *testing.T) {
	testCases := []struct {
		m, n uint64
		want *Tuplet
	}{
		{
			m: 1, n: 1,
			want: &Tuplet{m: 1, n: 1},
		},
		{
			m: 2, n: 1,
			want: &Tuplet{m: 2, n: 1},
		},
		{
			m: 2, n: 2,
			want: &Tuplet{m: 2, n: 2},
		},
		{
			m: 2, n: 3,
			want: &Tuplet{m: 2, n: 3},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, New(testCase.m, testCase.n))
	}
}

func TestTupletSet(t *testing.T) {
	testCases := []struct {
		*Tuplet
	}{
		{&Tuplet{1, 1}},
		{&Tuplet{2, 1}},
		{&Tuplet{1, 2}},
		{&Tuplet{2, 2}},
		{&Tuplet{3, 3}},
		{&Tuplet{4, 4}},
		{nil},
	}

	n, m := uint64(1), uint64(1)
	expectedTuplet := &Tuplet{m, n}

	for _, testCase := range testCases {
		assert.Equal(t, expectedTuplet, testCase.Set(m, n))
	}
}

func TestTupletSetTriplet(t *testing.T) {
	testCases := []struct {
		*Tuplet
	}{
		{&Tuplet{1, 1}},
		{&Tuplet{2, 1}},
		{&Tuplet{1, 2}},
		{&Tuplet{2, 2}},
		{&Tuplet{3, 3}},
		{&Tuplet{4, 4}},
		{nil},
	}

	expectedTuplet := &Tuplet{3, 2}

	for _, testCase := range testCases {
		assert.Equal(t, expectedTuplet, testCase.SetTriplet())
	}
}

func TestTupletSetDuplet(t *testing.T) {
	testCases := []struct {
		*Tuplet
	}{
		{&Tuplet{1, 1}},
		{&Tuplet{2, 1}},
		{&Tuplet{1, 2}},
		{&Tuplet{2, 2}},
		{&Tuplet{3, 3}},
		{&Tuplet{4, 4}},
		{nil},
	}

	expectedTuplet := &Tuplet{2, 3}

	for _, testCase := range testCases {
		assert.Equal(t, expectedTuplet, testCase.SetDuplet())
	}
}
