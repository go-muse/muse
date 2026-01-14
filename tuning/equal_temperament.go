package tuning

import "math"

// EqualTemperament implements equal temperament tuning where all intervals
// of the same type have exactly the same size.
//
// The frequency formula is: f = f_ref * 2^(steps/N)
// where N is the number of divisions per octave.
//
// This is the standard tuning system used in most modern Western music.
// Each semitone in 12-TET is exactly 100 cents (1200 cents / 12 = 100 cents).
type EqualTemperament struct{}

// NewEqualTemperament creates a new equal temperament instance.
func NewEqualTemperament() *EqualTemperament {
	return &EqualTemperament{}
}

// Frequency calculates the frequency using equal temperament.
// Formula: f = referenceFreq * 2^(stepsFromReference / tonesPerOctave)
func (e *EqualTemperament) Frequency(referenceFreq float64, stepsFromReference int, toneSystem ToneSystem) float64 {
	if toneSystem == 0 {
		return referenceFreq
	}
	exponent := float64(stepsFromReference) / float64(toneSystem)
	return referenceFreq * math.Pow(2, exponent)
}

// Name returns the name of the temperament.
func (e *EqualTemperament) Name() string {
	return "Equal Temperament"
}

// Ensure EqualTemperament implements Temperament interface.
var _ Temperament = (*EqualTemperament)(nil)
