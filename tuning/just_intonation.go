package tuning

import "math"

// JustIntonation implements just intonation (pure intonation) tuning
// based on simple integer frequency ratios derived from the harmonic series.
//
// Just intonation produces pure, beatless intervals for certain keys but
// does not allow free modulation between keys. It is commonly used in
// choral music, barbershop, and early music performance.
//
// The ratios used (for 12-tone system relative to the tonic):
//
//	Unison:      1/1
//	Minor 2nd:   16/15
//	Major 2nd:   9/8
//	Minor 3rd:   6/5
//	Major 3rd:   5/4
//	Perfect 4th: 4/3
//	Tritone:     45/32
//	Perfect 5th: 3/2
//	Minor 6th:   8/5
//	Major 6th:   5/3
//	Minor 7th:   9/5
//	Major 7th:   15/8
type JustIntonation struct{}

// NewJustIntonation creates a new just intonation instance.
func NewJustIntonation() *JustIntonation {
	return &JustIntonation{}
}

// justRatios contains the frequency ratios for each step in 12-tone just intonation.
// These are 5-limit just intonation ratios (using primes 2, 3, and 5).
var justRatios = []float64{
	1.0,          // 0: Unison (1/1)
	16.0 / 15.0,  // 1: Minor 2nd (16/15)
	9.0 / 8.0,    // 2: Major 2nd (9/8)
	6.0 / 5.0,    // 3: Minor 3rd (6/5)
	5.0 / 4.0,    // 4: Major 3rd (5/4)
	4.0 / 3.0,    // 5: Perfect 4th (4/3)
	45.0 / 32.0,  // 6: Tritone (45/32)
	3.0 / 2.0,    // 7: Perfect 5th (3/2)
	8.0 / 5.0,    // 8: Minor 6th (8/5)
	5.0 / 3.0,    // 9: Major 6th (5/3)
	9.0 / 5.0,    // 10: Minor 7th (9/5)
	15.0 / 8.0,   // 11: Major 7th (15/8)
}

// Frequency calculates the frequency using just intonation.
// For 12-tone systems, uses predefined ratios. For other systems,
// falls back to equal temperament.
func (j *JustIntonation) Frequency(referenceFreq float64, stepsFromReference int, toneSystem ToneSystem) float64 {
	if toneSystem == 0 {
		return referenceFreq
	}

	// Calculate octave offset and position within octave
	stepsPerOctave := int(toneSystem)
	octaves := stepsFromReference / stepsPerOctave
	positionInOctave := stepsFromReference % stepsPerOctave

	// Handle negative positions
	if positionInOctave < 0 {
		positionInOctave += stepsPerOctave
		octaves--
	}

	// Get the ratio for the position within the octave
	var ratio float64
	if toneSystem == TwelveTone && positionInOctave < len(justRatios) {
		ratio = justRatios[positionInOctave]
	} else {
		// Fall back to equal temperament for non-12-tone systems
		exponent := float64(positionInOctave) / float64(toneSystem)
		ratio = math.Pow(2, exponent)
	}

	// Apply octave transposition
	octaveMultiplier := math.Pow(2, float64(octaves))

	return referenceFreq * ratio * octaveMultiplier
}

// Name returns the name of the temperament.
func (j *JustIntonation) Name() string {
	return "Just Intonation"
}

// Ensure JustIntonation implements Temperament interface.
var _ Temperament = (*JustIntonation)(nil)
