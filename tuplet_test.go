package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuplet(t *testing.T) {
	testCases := []struct {
		n, m uint8
		want *Tuplet
	}{
		{
			n: 1, m: 1,
			want: &Tuplet{n: 1, m: 1},
		},
		{
			n: 1, m: 2,
			want: &Tuplet{n: 1, m: 2},
		},
		{
			n: 2, m: 2,
			want: &Tuplet{n: 2, m: 2},
		},
		{
			n: 3, m: 2,
			want: &Tuplet{n: 3, m: 2},
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, NewTuplet(testCase.n, testCase.m))
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

	n, m := uint8(1), uint8(1)
	expectedTuplet := &Tuplet{n, m}

	for _, testCase := range testCases {
		assert.Equal(t, expectedTuplet, testCase.Set(n, m))
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

func TestTupletRemoveTuplet(t *testing.T) {
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

	for _, testCase := range testCases {
		assert.Nil(t, testCase.RemoveTuplet())
	}
}
