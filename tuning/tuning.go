package tuning

// Standard reference frequencies for A4 in Hz.
const (
	// FreqA444 is used by some orchestras for a brighter sound.
	FreqA444 = 444.0

	// FreqA440 is the modern international standard pitch (ISO 16).
	FreqA440 = 440.0

	// FreqA432 is known as "Verdi's A" or "philosophical pitch",
	// sometimes claimed to have special properties.
	FreqA432 = 432.0

	// FreqA415 is commonly used for Baroque music performance,
	// approximately a semitone lower than A440.
	FreqA415 = 415.0
)

// Tuning combines all parameters needed to calculate note frequencies:
// the reference frequency, the temperament system, and the tone system.
type Tuning struct {
	// ReferenceFreq is the frequency of the reference note (A4) in Hz.
	ReferenceFreq float64

	// Temperament defines how intervals are calculated.
	Temperament Temperament

	// ToneSystem defines the number of equal divisions per octave.
	ToneSystem ToneSystem
}

// New creates a new Tuning with custom parameters.
func New(referenceFreq float64, temperament Temperament, toneSystem ToneSystem) Tuning {
	return Tuning{
		ReferenceFreq: referenceFreq,
		Temperament:   temperament,
		ToneSystem:    toneSystem,
	}
}

// Standard12TET returns the standard modern tuning:
// A4 = 440 Hz, equal temperament, 12 tones per octave.
func Standard12TET() Tuning {
	return Tuning{
		ReferenceFreq: FreqA440,
		Temperament:   NewEqualTemperament(),
		ToneSystem:    TwelveTone,
	}
}

// Baroque12TET returns a Baroque-style tuning:
// A4 = 415 Hz, equal temperament, 12 tones per octave.
func Baroque12TET() Tuning {
	return Tuning{
		ReferenceFreq: FreqA415,
		Temperament:   NewEqualTemperament(),
		ToneSystem:    TwelveTone,
	}
}

// Bright12TET returns a bright orchestral tuning:
// A4 = 444 Hz, equal temperament, 12 tones per octave.
func Bright12TET() Tuning {
	return Tuning{
		ReferenceFreq: FreqA444,
		Temperament:   NewEqualTemperament(),
		ToneSystem:    TwelveTone,
	}
}

// Verdi12TET returns Verdi's tuning:
// A4 = 432 Hz, equal temperament, 12 tones per octave.
func Verdi12TET() Tuning {
	return Tuning{
		ReferenceFreq: FreqA432,
		Temperament:   NewEqualTemperament(),
		ToneSystem:    TwelveTone,
	}
}

// JustIntonation12 returns a just intonation tuning:
// A4 = 440 Hz, just intonation, 12 tones per octave.
func JustIntonation12() Tuning {
	return Tuning{
		ReferenceFreq: FreqA440,
		Temperament:   NewJustIntonation(),
		ToneSystem:    TwelveTone,
	}
}

// Pythagorean12 returns a Pythagorean tuning:
// A4 = 440 Hz, Pythagorean temperament, 12 tones per octave.
func Pythagorean12() Tuning {
	return Tuning{
		ReferenceFreq: FreqA440,
		Temperament:   NewPythagoreanTemperament(),
		ToneSystem:    TwelveTone,
	}
}

// Meantone12 returns a quarter-comma meantone tuning:
// A4 = 440 Hz, meantone temperament, 12 tones per octave.
func Meantone12() Tuning {
	return Tuning{
		ReferenceFreq: FreqA440,
		Temperament:   NewMeantoneTemperament(),
		ToneSystem:    TwelveTone,
	}
}

// QuarterTone24TET returns a quarter-tone tuning:
// A4 = 440 Hz, equal temperament, 24 tones per octave.
func QuarterTone24TET() Tuning {
	return Tuning{
		ReferenceFreq: FreqA440,
		Temperament:   NewEqualTemperament(),
		ToneSystem:    TwentyFourTone,
	}
}

// Frequency calculates the frequency of a note given its steps from the reference note (A4).
func (t Tuning) Frequency(stepsFromReference int) float64 {
	return t.Temperament.Frequency(t.ReferenceFreq, stepsFromReference, t.ToneSystem)
}

// WithReferenceFreq returns a copy of the Tuning with a different reference frequency.
func (t Tuning) WithReferenceFreq(freq float64) Tuning {
	t.ReferenceFreq = freq
	return t
}

// WithTemperament returns a copy of the Tuning with a different temperament.
func (t Tuning) WithTemperament(temperament Temperament) Tuning {
	t.Temperament = temperament
	return t
}

// WithToneSystem returns a copy of the Tuning with a different tone system.
func (t Tuning) WithToneSystem(toneSystem ToneSystem) Tuning {
	t.ToneSystem = toneSystem
	return t
}
