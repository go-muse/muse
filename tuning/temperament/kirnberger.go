package temperament

import (
	"math"
)

// KirnbergerIII implements Kirnberger III temperament (1779),
// designed by Johann Philipp Kirnberger, a student of J.S. Bach.
//
// Kirnberger III is notable for having:
// - Pure major thirds on C-E, G-B, and F-A
// - Four fifths tempered by 1/4 syntonic comma each
// - One fifth (D-A) tempered by the schisma
//
// This temperament provides excellent triads in keys near C major
// while maintaining usability in remote keys. It's considered
// one of the most practical well temperaments for Baroque music.
type KirnbergerIII struct{}

// NewKirnbergerIII creates a new Kirnberger III temperament instance.
func NewKirnbergerIII() *KirnbergerIII {
	return &KirnbergerIII{}
}

// kirnbergerIIIRatios contains the frequency ratios for Kirnberger III.
// Based on tempering specific fifths to achieve pure major thirds
// on C-E, G-B, and F-A.
var kirnbergerIIIRatios = []float64{
	1.0,                   // 0: C (Unison)
	256.0 / 243.0,         // 1: C#/Db
	math.Sqrt(2) / math.Pow(3.0/2.0, 0.5) * (9.0 / 8.0), // 2: D
	32.0 / 27.0,           // 3: D#/Eb
	5.0 / 4.0,             // 4: E (pure major third)
	4.0 / 3.0,             // 5: F
	45.0 / 32.0,           // 6: F#/Gb
	math.Pow(5, 0.25),     // 7: G (tempered)
	128.0 / 81.0,          // 8: G#/Ab
	5.0 / 3.0,             // 9: A (pure major sixth from F)
	16.0 / 9.0,            // 10: A#/Bb
	15.0 / 8.0,            // 11: B (pure major third from G)
}

// Frequency calculates the frequency using Kirnberger III temperament.
func (k *KirnbergerIII) Frequency(referenceFreq float64, stepsFromReference int, stepsPerOctave int) float64 {
	if stepsPerOctave == 0 {
		return referenceFreq
	}

	octaves := stepsFromReference / stepsPerOctave
	positionInOctave := stepsFromReference % stepsPerOctave

	if positionInOctave < 0 {
		positionInOctave += stepsPerOctave
		octaves--
	}

	var ratio float64
	if stepsPerOctave == standardTwelveTone && positionInOctave < len(kirnbergerIIIRatios) {
		ratio = kirnbergerIIIRatios[positionInOctave]
	} else {
		exponent := float64(positionInOctave) / float64(stepsPerOctave)
		ratio = math.Pow(2, exponent)
	}

	octaveMultiplier := math.Pow(2, float64(octaves))
	return referenceFreq * ratio * octaveMultiplier
}

// Name returns the name of the temperament.
func (k *KirnbergerIII) Name() string {
	return "Kirnberger III"
}
