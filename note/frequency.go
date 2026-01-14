package note

import (
	"github.com/go-muse/muse/tuning"
)

// referenceOctave is the octave number of the reference pitch A4.
const referenceOctave = 4

// StepsFromA4 returns the number of steps from A4 in the given tone system.
// For 12-tone system, this is the number of semitones.
// For 24-tone system, this is the number of quarter-tones (each semitone = 2 steps).
//
// Requires the note to have an octave set. If octave is nil, returns 0.
func (n Note) StepsFromA4(toneSystem tuning.ToneSystem) int {
	if n.octave == nil {
		return 0
	}

	// Get base semitone position within octave (C=0, D=2, ..., A=9, B=11)
	baseSemitone := int(n.name.Letter().Semitone())

	// Calculate semitones from A (within same octave)
	// A=9, so C=0 means -9 from A, D=2 means -7 from A, etc.
	stepsFromA := baseSemitone - int(LetterA.Semitone())

	// Add alteration (sharps/flats)
	stepsFromA += int(n.AlterationShift())

	// Add octave offset (octave 4 is the reference)
	// Each octave is 12 semitones
	octaveOffset := int(n.octave.Number()) - referenceOctave
	stepsIn12Tone := stepsFromA + octaveOffset*tuning.SemitonesPerOctave

	// Convert to the target tone system
	// For 12-tone: multiply by 1
	// For 24-tone: multiply by 2 (each semitone = 2 quarter-tones)
	if toneSystem == tuning.TwelveTone || toneSystem == 0 {
		return stepsIn12Tone
	}

	// Scale to target tone system
	return stepsIn12Tone * int(toneSystem) / tuning.SemitonesPerOctave
}

// Frequency calculates and returns the frequency of the note using the given tuning.
//
// The tuning specifies:
//   - Reference frequency (e.g., 440 Hz for A4)
//   - Temperament (equal, just intonation, Pythagorean, etc.)
//   - Tone system (12, 19, 24, 31 tones per octave)
//
// Requires the note to have an octave set.
func (n Note) Frequency(t tuning.Tuning) float64 {
	steps := n.StepsFromA4(t.ToneSystem)
	return t.Frequency(steps)
}

// Frequency440 returns the frequency of the note in standard 12-TET tuning (A4 = 440 Hz).
func (n Note) Frequency440() float64 {
	return n.Frequency(tuning.Standard12TET())
}

// Frequency444 returns the frequency of the note in bright orchestral tuning (A4 = 444 Hz).
func (n Note) Frequency444() float64 {
	return n.Frequency(tuning.Bright12TET())
}

// Frequency432 returns the frequency of the note in Verdi tuning (A4 = 432 Hz).
func (n Note) Frequency432() float64 {
	return n.Frequency(tuning.Verdi12TET())
}

// Frequency415 returns the frequency of the note in Baroque tuning (A4 = 415 Hz).
func (n Note) Frequency415() float64 {
	return n.Frequency(tuning.Baroque12TET())
}
