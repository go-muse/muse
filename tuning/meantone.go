package tuning

import "math"

// MeantoneTemperament implements quarter-comma meantone temperament,
// the most common form of meantone tuning historically used from
// the late 15th to the late 18th century.
//
// In quarter-comma meantone, the fifths are narrowed by 1/4 of a
// syntonic comma (about 5.38 cents) to produce pure major thirds (5:4).
// This makes the temperament particularly suitable for music with
// many thirds and sixths.
//
// The trade-off is that the fifths are slightly flat, and one "wolf fifth"
// (usually G#-Eb) is very out of tune. Keys remote from the central keys
// become unusable.
//
// Quarter-comma meantone ratios (for 12-tone system):
// The fifth ratio is 5^(1/4) ≈ 1.495349 (vs 1.5 for pure)
type MeantoneTemperament struct{}

// NewMeantoneTemperament creates a new quarter-comma meantone temperament instance.
func NewMeantoneTemperament() *MeantoneTemperament {
	return &MeantoneTemperament{}
}

// meantoneRatios contains the frequency ratios for each step in 12-tone quarter-comma meantone.
// These are calculated based on tempered fifths of ratio 5^(1/4).
var meantoneRatios = []float64{
	1.0,                                  // 0: Unison
	8.0 / (5.0 * math.Pow(5, 0.25)),      // 1: Minor 2nd (diatonic semitone)
	math.Sqrt(5) / 2.0,                   // 2: Major 2nd
	4.0 / math.Pow(5, 0.75),              // 3: Minor 3rd
	5.0 / 4.0,                            // 4: Major 3rd (pure)
	2.0 * math.Pow(5, 0.25) / math.Sqrt(5), // 5: Perfect 4th
	math.Pow(5, 0.25) * math.Sqrt(5) / 2.0, // 6: Augmented 4th
	math.Pow(5, 0.25),                    // 7: Perfect 5th (tempered)
	8.0 / 5.0,                            // 8: Minor 6th (pure)
	math.Pow(5, 0.5) * math.Pow(5, 0.25) / 2.0, // 9: Major 6th
	4.0 * math.Pow(5, 0.25) / 5.0,        // 10: Minor 7th
	5.0 * math.Pow(5, 0.25) / 4.0,        // 11: Major 7th
}

// Frequency calculates the frequency using quarter-comma meantone temperament.
// For 12-tone systems, uses predefined ratios. For other systems,
// falls back to equal temperament.
func (m *MeantoneTemperament) Frequency(referenceFreq float64, stepsFromReference int, toneSystem ToneSystem) float64 {
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
	if toneSystem == TwelveTone && positionInOctave < len(meantoneRatios) {
		ratio = meantoneRatios[positionInOctave]
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
func (m *MeantoneTemperament) Name() string {
	return "Quarter-Comma Meantone Temperament"
}

// Ensure MeantoneTemperament implements Temperament interface.
var _ Temperament = (*MeantoneTemperament)(nil)
