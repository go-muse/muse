package muse

import "testing"

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
