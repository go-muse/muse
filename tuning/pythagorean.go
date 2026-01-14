package tuning

import "math"

// PythagoreanTemperament implements Pythagorean tuning based on
// a chain of pure perfect fifths (ratio 3:2).
//
// In Pythagorean tuning, all intervals are derived from stacking
// perfect fifths. This produces pure fifths and fourths, but
// the major thirds are wider than pure (81:64 instead of 5:4).
//
// The Pythagorean comma (about 23.46 cents) is the difference between
// 12 pure fifths and 7 octaves, which creates a "wolf fifth" in closed systems.
//
// Pythagorean ratios (for 12-tone system):
//
//	Unison:      1/1
//	Minor 2nd:   256/243 (limma)
//	Major 2nd:   9/8
//	Minor 3rd:   32/27
//	Major 3rd:   81/64 (ditone)
//	Perfect 4th: 4/3
//	Tritone:     729/512
//	Perfect 5th: 3/2
//	Minor 6th:   128/81
//	Major 6th:   27/16
//	Minor 7th:   16/9
//	Major 7th:   243/128
type PythagoreanTemperament struct{}

// NewPythagoreanTemperament creates a new Pythagorean temperament instance.
func NewPythagoreanTemperament() *PythagoreanTemperament {
	return &PythagoreanTemperament{}
}

// pythagoreanRatios contains the frequency ratios for each step in 12-tone Pythagorean tuning.
// All intervals are derived from the ratio 3:2 (pure fifth).
var pythagoreanRatios = []float64{
	1.0,             // 0: Unison (1/1)
	256.0 / 243.0,   // 1: Minor 2nd - limma (256/243)
	9.0 / 8.0,       // 2: Major 2nd (9/8)
	32.0 / 27.0,     // 3: Minor 3rd (32/27)
	81.0 / 64.0,     // 4: Major 3rd - ditone (81/64)
	4.0 / 3.0,       // 5: Perfect 4th (4/3)
	729.0 / 512.0,   // 6: Tritone (729/512)
	3.0 / 2.0,       // 7: Perfect 5th (3/2)
	128.0 / 81.0,    // 8: Minor 6th (128/81)
	27.0 / 16.0,     // 9: Major 6th (27/16)
	16.0 / 9.0,      // 10: Minor 7th (16/9)
	243.0 / 128.0,   // 11: Major 7th (243/128)
}

// Frequency calculates the frequency using Pythagorean tuning.
// For 12-tone systems, uses predefined ratios. For other systems,
// falls back to equal temperament.
func (p *PythagoreanTemperament) Frequency(referenceFreq float64, stepsFromReference int, toneSystem ToneSystem) float64 {
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
	if toneSystem == TwelveTone && positionInOctave < len(pythagoreanRatios) {
		ratio = pythagoreanRatios[positionInOctave]
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
func (p *PythagoreanTemperament) Name() string {
	return "Pythagorean Temperament"
}

// Ensure PythagoreanTemperament implements Temperament interface.
var _ Temperament = (*PythagoreanTemperament)(nil)
