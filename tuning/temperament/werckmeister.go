package temperament

import (
	"math"
)

// WerckmeisterIII implements Werckmeister III temperament (1691),
// also known as "correct temperament" or "well temperament".
//
// Andreas Werckmeister designed this temperament to allow all keys
// to be usable while giving different keys distinct characters.
// It is one of the most famous well temperaments.
//
// In Werckmeister III:
// - Four fifths (C-G, G-D, D-A, B-F#) are tempered by 1/4 Pythagorean comma
// - All other fifths are pure (3:2)
// - Keys near C major are close to just intonation
// - Remote keys have wider thirds but remain usable
//
// This temperament is often associated with J.S. Bach's "Well-Tempered Clavier",
// though this connection is debated by scholars.
type WerckmeisterIII struct{}

// NewWerckmeisterIII creates a new Werckmeister III temperament instance.
func NewWerckmeisterIII() *WerckmeisterIII {
	return &WerckmeisterIII{}
}

// werckmeisterIIIRatios contains the frequency ratios for Werckmeister III.
// Calculated based on the tempering of specific fifths by 1/4 Pythagorean comma.
// The Pythagorean comma is (3^12)/(2^19) ≈ 1.01364.
// 1/4 comma = comma^(1/4) ≈ 1.00339.
var werckmeisterIIIRatios = []float64{
	1.0,                // 0: C (Unison)
	256.0 / 243.0,      // 1: C#/Db (Pythagorean limma)
	1.1174033085417763, // 2: D (tempered)
	32.0 / 27.0,        // 3: D#/Eb (Pythagorean minor 3rd)
	1.2528272887537059, // 4: E (tempered)
	4.0 / 3.0,          // 5: F (pure fourth)
	1024.0 / 729.0,     // 6: F#/Gb (Pythagorean tritone)
	1.4949569988157949, // 7: G (tempered fifth)
	128.0 / 81.0,       // 8: G#/Ab (Pythagorean minor 6th)
	1.6704363316362592, // 9: A (tempered)
	16.0 / 9.0,         // 10: A#/Bb (Pythagorean minor 7th)
	1.8792409618044082, // 11: B (tempered)
}

// Frequency calculates the frequency using Werckmeister III temperament.
func (w *WerckmeisterIII) Frequency(referenceFreq float64, stepsFromReference int, stepsPerOctave int) float64 {
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
	if stepsPerOctave == standardTwelveTone && positionInOctave < len(werckmeisterIIIRatios) {
		ratio = werckmeisterIIIRatios[positionInOctave]
	} else {
		exponent := float64(positionInOctave) / float64(stepsPerOctave)
		ratio = math.Pow(2, exponent)
	}

	octaveMultiplier := math.Pow(2, float64(octaves))
	return referenceFreq * ratio * octaveMultiplier
}

// Name returns the name of the temperament.
func (w *WerckmeisterIII) Name() string {
	return "Werckmeister III"
}
