package temperament

import (
	"math"
)

// Young implements Young's well temperament (1799),
// designed by Thomas Young, the English polymath.
//
// Young's temperament is based on:
// - Six fifths tempered by 3/16 syntonic comma
// - Six pure fifths
//
// The distribution creates a very practical temperament that:
// - Provides good thirds in common keys
// - Maintains usable thirds in all keys
// - Has a gentle gradation of key colors
//
// Young's temperament is very close to Vallotti and is sometimes
// considered an improvement upon it. It's particularly suitable
// for late Baroque and Classical music.
type Young struct{}

// NewYoung creates a new Young's well temperament instance.
func NewYoung() *Young {
	return &Young{}
}

// youngRatios contains the frequency ratios for Young's well temperament.
// Based on tempering six fifths by 3/16 syntonic comma.
var youngRatios = []float64{
	1.0,        // 0: C (Unison)
	1.05350,    // 1: C#/Db
	1.11916,    // 2: D
	1.18519,    // 3: D#/Eb
	1.25424,    // 4: E
	1.33333,    // 5: F (pure 4/3)
	1.40625,    // 6: F#/Gb
	1.49662,    // 7: G
	1.58025,    // 8: G#/Ab
	1.67411,    // 9: A
	1.77778,    // 10: A#/Bb (pure 16/9)
	1.88145,    // 11: B
}

// Frequency calculates the frequency using Young's well temperament.
func (y *Young) Frequency(referenceFreq float64, stepsFromReference int, stepsPerOctave int) float64 {
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
	if stepsPerOctave == standardTwelveTone && positionInOctave < len(youngRatios) {
		ratio = youngRatios[positionInOctave]
	} else {
		exponent := float64(positionInOctave) / float64(stepsPerOctave)
		ratio = math.Pow(2, exponent)
	}

	octaveMultiplier := math.Pow(2, float64(octaves))
	return referenceFreq * ratio * octaveMultiplier
}

// Name returns the name of the temperament.
func (y *Young) Name() string {
	return "Young's Well Temperament"
}
