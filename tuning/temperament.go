package tuning

// Temperament defines how to calculate the frequency of a note
// based on its distance from a reference pitch.
//
// Different temperaments produce different interval sizes and characteristics.
// The most common is equal temperament, but historical and experimental
// temperaments offer different musical qualities.
type Temperament interface {
	// Frequency calculates the frequency of a note.
	//
	// Parameters:
	//   - referenceFreq: frequency of the reference note in Hz (e.g., A4 = 440 Hz)
	//   - stepsFromReference: number of steps from the reference note
	//     (positive = higher, negative = lower)
	//   - toneSystem: the number of equal divisions per octave
	//
	// Returns the calculated frequency in Hz.
	Frequency(referenceFreq float64, stepsFromReference int, toneSystem ToneSystem) float64

	// Name returns the human-readable name of the temperament.
	Name() string
}
