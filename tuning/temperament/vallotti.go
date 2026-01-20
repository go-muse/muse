package temperament

import (
	"math"
)

// Vallotti implements Vallotti temperament (c. 1730),
// designed by Francesco Antonio Vallotti.
//
// Vallotti is a well temperament where:
// - Six fifths (F-C-G-D-A-E-B) are each tempered by 1/6 Pythagorean comma
// - Six fifths (B-F#-C#-G#-D#-A#-F) are pure
//
// This creates a symmetric pattern and provides a gentle gradation
// of key colors. Keys with fewer sharps/flats have purer thirds,
// while remote keys have wider thirds but remain usable.
//
// Vallotti is popular for Classical and early Romantic music
// and is favored by many harpsichordists and organists.
type Vallotti struct{}

// NewVallotti creates a new Vallotti temperament instance.
func NewVallotti() *Vallotti {
	return &Vallotti{}
}

// vallottiRatios contains the frequency ratios for Vallotti temperament.
// The Pythagorean comma divided by 6 gives the tempering for each of six fifths.
// Pythagorean comma ≈ 23.46 cents, so each tempered fifth is narrowed by ~3.91 cents.
var vallottiRatios = []float64{
	1.0,        // 0: C (Unison)
	1.05350,    // 1: C#/Db
	1.11740,    // 2: D
	1.18518,    // 3: D#/Eb
	1.25283,    // 4: E
	4.0 / 3.0,  // 5: F (pure 4/3)
	1.40625,    // 6: F#/Gb
	1.49493,    // 7: G
	1.58025,    // 8: G#/Ab
	1.67044,    // 9: A
	16.0 / 9.0, // 10: A#/Bb (pure 16/9)
	1.87924,    // 11: B
}

// Frequency calculates the frequency using Vallotti temperament.
func (v *Vallotti) Frequency(referenceFreq float64, stepsFromReference int, stepsPerOctave int) float64 {
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
	if stepsPerOctave == standardTwelveTone && positionInOctave < len(vallottiRatios) {
		ratio = vallottiRatios[positionInOctave]
	} else {
		exponent := float64(positionInOctave) / float64(stepsPerOctave)
		ratio = math.Pow(2, exponent)
	}

	octaveMultiplier := math.Pow(2, float64(octaves))
	return referenceFreq * ratio * octaveMultiplier
}

// Name returns the name of the temperament.
func (v *Vallotti) Name() string {
	return "Vallotti"
}
