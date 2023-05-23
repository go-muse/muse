package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChromaticIntervalHalfTonesMethod(t *testing.T) {
	expectedHalfTones := HalfTones(12)
	// Create a test interval with a known number of half tones
	testInterval := ChromaticInterval{
		halfTones: expectedHalfTones,
	}

	// Ensure that the HalfTones() method returns the expected number of half tones
	if testInterval.HalfTones() != expectedHalfTones {
		t.Errorf("HalfTones() returned %d, expected %d", testInterval.HalfTones(), expectedHalfTones)
	}
}

func TestMakeNoteByIntervalName(t *testing.T) {
	firstNote := MustNewNoteWithoutOctave(C)
	note, err := MakeNoteByIntervalName(firstNote, IntervalNameTritone)
	assert.NoError(t, err)
	assert.NotNil(t, note)
	assert.Equal(t, note.Name(), FSHARP)
}

func TestMakeDegreeByIntervalName(t *testing.T) {
	firstDegree := NewDegree(1, 0, nil, nil, MustNewNoteWithoutOctave(C), nil, nil)
	interval, err := NewIntervalChromatic(6)
	assert.NoError(t, err)
	secondDegree, err := MakeDegreeByIntervalName(firstDegree, interval.Name())
	assert.NoError(t, err)
	assert.NotNil(t, secondDegree)
	assert.Equal(t, secondDegree.Note().Name(), FSHARP)
	assert.Equal(t, secondDegree.Number(), firstDegree.Number()+1)
	assert.Equal(t, secondDegree.HalfTonesFromPrime(), interval.HalfTones())
}
