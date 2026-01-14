package temperament

import (
	"math"
)

// Equal implements equal temperament tuning where all intervals
// of the same type have exactly the same size.
//
// The frequency formula is: f = f_ref * 2^(steps/N)
// where N is the number of divisions per octave.
//
// This is the standard tuning system used in most modern Western music.
// Each semitone in 12-TET is exactly 100 cents (1200 cents / 12 = 100 cents).
type Equal struct{}

// NewEqual creates a new equal temperament instance.
func NewEqual() *Equal {
	return &Equal{}
}

// Frequency calculates the frequency using equal temperament.
// Formula: f = referenceFreq * 2^(stepsFromReference / stepsPerOctave)
func (e *Equal) Frequency(referenceFreq float64, stepsFromReference int, stepsPerOctave int) float64 {
	if stepsPerOctave == 0 {
		return referenceFreq
	}
	exponent := float64(stepsFromReference) / float64(stepsPerOctave)
	return referenceFreq * math.Pow(2, exponent)
}

// Name returns the name of the temperament.
func (e *Equal) Name() string {
	return "Equal Temperament"
}
