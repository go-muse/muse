package tuning

// ToneSystem defines the number of equal divisions of the octave.
// Common systems include 12-TET (standard Western music), 19-TET, 24-TET (quarter-tones),
// and 31-TET. The value represents how many equal steps divide one octave.
type ToneSystem uint8

// CentsPerOctave is the number of cents in one octave.
// This is a fundamental constant in music theory: 1 octave = 1200 cents.
const CentsPerOctave = 1200.0

// SemitonesPerOctave is the number of semitones in one octave
// in the standard 12-tone chromatic scale.
const SemitonesPerOctave = 12

// Standard tone systems (equal divisions of the octave).
const (
	// TwelveTone is the standard Western 12-tone equal temperament (12-TET).
	// Each semitone is 100 cents, and an octave spans 1200 cents.
	TwelveTone ToneSystem = SemitonesPerOctave

	// NineteenTone is 19-TET, historically used and provides better major thirds.
	// Each step is approximately 63.16 cents.
	NineteenTone ToneSystem = 19

	// TwentyFourTone is 24-TET, the quarter-tone system used in some
	// Middle Eastern music and contemporary classical music.
	// Each step is 50 cents (quarter-tone).
	TwentyFourTone ToneSystem = 24

	// ThirtyOneTone is 31-TET, which provides excellent approximations
	// of just intonation intervals. Each step is approximately 38.71 cents.
	ThirtyOneTone ToneSystem = 31
)

// StepsPerOctave returns the number of steps in one octave for this tone system.
func (t ToneSystem) StepsPerOctave() int {
	return int(t)
}

// CentsPerStep returns the size of one step in cents.
// One octave equals 1200 cents.
func (t ToneSystem) CentsPerStep() float64 {
	if t == 0 {
		return 0
	}

	return CentsPerOctave / float64(t)
}
