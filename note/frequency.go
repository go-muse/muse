package note

import (
	"math"
)

// Standards frequencies of A4 in Hz.
const (
	FreqA444 = 444.0 // Certain orchestras for brighter sound
	FreqA440 = 440.0 // Modern music and orchestras
	FreqA432 = 432.0 // Verdi's frequency
	FreqA415 = 415.0 // Performance of Baroque music
)

// getMapOfHalfTones returns map that contains the number of halftones above A4 for each note.
func getMapOfHalfTones() map[Name]int8 {
	return map[Name]int8{
		C:       -9,
		DFLAT2:  -9,
		DFLAT:   -8,
		CSHARP:  -8,
		CSHARP2: -7,
		D:       -7,
		EFLAT2:  -7,
		FFLAT2:  -6,
		EFLAT:   -6,
		DSHARP:  -6,
		DSHARP2: -5,
		E:       -5,
		FFLAT:   -5,
		ESHARP:  -4,
		F:       -4,
		GFLAT2:  -4,
		ESHARP2: -3,
		GFLAT:   -3,
		FSHARP:  -3,
		FSHARP2: -2,
		G:       -2,
		AFLAT2:  -2,
		AFLAT:   -1,
		GSHARP:  -1,
		GSHARP2: 0,
		A:       0,
		BFLAT2:  0,
		BFLAT:   1,
		ASHARP:  1,
		CFLAT2:  1,
		B:       2,
		CFLAT:   2, // Cb = B
		BSHARP:  3, // B# = C
		BSHARP2: 4, // B## = C#
	}
}

// Frequency returns the frequency of the note calculated in the specified standard.
func (n *Note) Frequency(standard float64) float64 {
	if n == nil || n.Octave() == nil {
		return 0
	}

	baseName := n.BaseName()
	accidentals := int8(0)
	for _, char := range n.Name()[1:] {
		if char == rune(AccidentalSharp[0]) {
			accidentals++
		} else if char == rune(AccidentalFlat[0]) {
			accidentals--
		}
	}

	semitoneOffset, exists := getMapOfHalfTones()[baseName]
	if !exists {
		return 0
	}

	totalSemitones := semitoneOffset + accidentals

	// f = 440 * 2^(n/12), where n — halftones from A4 (A in first octave)
	return standard * math.Pow(2, float64(totalSemitones+12*(int8(n.Octave().Number())-4))/12)
}

// FrequencyBy444 returns the frequency of the note calculated in the 444 Hz standard.
func (n *Note) FrequencyBy444() float64 {
	return n.Frequency(FreqA444)
}

// FrequencyBy440 returns the frequency of the note calculated in the 440 Hz standard.
func (n *Note) FrequencyBy440() float64 {
	return n.Frequency(FreqA440)
}

// FrequencyBy432 returns the frequency of the note calculated in the 432 Hz standard.
func (n *Note) FrequencyBy432() float64 {
	return n.Frequency(FreqA432)
}

// FrequencyBy415 returns the frequency of the note calculated in the 415 Hz standard.
func (n *Note) FrequencyBy415() float64 {
	return n.Frequency(FreqA415)
}
