package note

import (
	"math"
	"testing"

	"github.com/go-muse/muse/octave"
	"github.com/go-muse/muse/tuning"
)

func TestNote_StepsFromA4(t *testing.T) {
	tests := []struct {
		name      string
		note      Note
		wantSteps int
	}{
		// A4 is the reference (0 steps)
		{"A4", MustNewWithOctave(A, 4), 0},

		// Notes above A4
		{"A#4", MustNewWithOctave(ASHARP, 4), 1},
		{"B4", MustNewWithOctave(B, 4), 2},
		{"C5", MustNewWithOctave(C, 5), 3},
		{"A5", MustNewWithOctave(A, 5), 12},

		// Notes below A4
		{"G#4", MustNewWithOctave(GSHARP, 4), -1},
		{"G4", MustNewWithOctave(G, 4), -2},
		{"C4", MustNewWithOctave(C, 4), -9},
		{"A3", MustNewWithOctave(A, 3), -12},
		{"A2", MustNewWithOctave(A, 2), -24},
		{"A0", MustNewWithOctave(A, 0), -48},
		{"A-1", MustNewWithOctave(A, -1), -60},

		// Enharmonic equivalents
		{"Bb4", MustNewWithOctave(BFLAT, 4), 1},
		{"Ab4", MustNewWithOctave(AFLAT, 4), -1},

		// Double alterations
		{"A##4", MustNewWithOctave(ASHARP2, 4), 2},
		{"Abb4", MustNewWithOctave(AFLAT2, 4), -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.StepsFromA4(tuning.TwelveTone)
			if got != tt.wantSteps {
				t.Fatalf("StepsFromA4(): got %d, want %d", got, tt.wantSteps)
			}
		})
	}
}

func TestNote_StepsFromA4_NilOctave(t *testing.T) {
	note := New(A)
	if got := note.StepsFromA4(tuning.TwelveTone); got != 0 {
		t.Fatalf("StepsFromA4() with nil octave: got %d, want 0", got)
	}
}

func TestNote_StepsFromA4_ToneSystems(t *testing.T) {
	// For 24-tone system, each semitone = 2 quarter-tones
	tests := []struct {
		name       string
		note       Note
		toneSystem tuning.ToneSystem
		wantSteps  int
	}{
		{"A4_12tone", MustNewWithOctave(A, 4), tuning.TwelveTone, 0},
		{"A4_24tone", MustNewWithOctave(A, 4), tuning.TwentyFourTone, 0},
		{"B4_12tone", MustNewWithOctave(B, 4), tuning.TwelveTone, 2},
		{"B4_24tone", MustNewWithOctave(B, 4), tuning.TwentyFourTone, 4},
		{"A5_12tone", MustNewWithOctave(A, 5), tuning.TwelveTone, 12},
		{"A5_24tone", MustNewWithOctave(A, 5), tuning.TwentyFourTone, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.StepsFromA4(tt.toneSystem)
			if got != tt.wantSteps {
				t.Fatalf("StepsFromA4(%v): got %d, want %d", tt.toneSystem, got, tt.wantSteps)
			}
		})
	}
}

func TestNote_Frequency440(t *testing.T) {
	tests := []struct {
		name     string
		note     Note
		expected float64
		epsilon  float64
	}{
		// Standard reference
		{"A4", MustNewWithOctave(A, 4), 440.0, 0.01},

		// One octave relationships
		{"A3", MustNewWithOctave(A, 3), 220.0, 0.01},
		{"A5", MustNewWithOctave(A, 5), 880.0, 0.01},
		{"A2", MustNewWithOctave(A, 2), 110.0, 0.01},

		// Common notes
		{"C4", MustNewWithOctave(C, 4), 261.63, 0.01},
		{"G#4", MustNewWithOctave(GSHARP, 4), 415.30, 0.01},
		{"Bb3", MustNewWithOctave(BFLAT, 3), 233.08, 0.01},
		{"F#5", MustNewWithOctave(FSHARP, 5), 739.99, 0.01},

		// Extreme octaves
		{"C-1", MustNewWithOctave(C, -1), 8.18, 0.02},
		{"C0", MustNewWithOctave(C, 0), 16.35, 0.01},
		{"B8", MustNewWithOctave(B, 8), 7902.13, 1.0},
		{"B9", MustNewWithOctave(B, 9), 15804.27, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.Frequency440()
			if !almostEqual(got, tt.expected, tt.epsilon) {
				t.Fatalf("Frequency440(): got %f, want %f (epsilon %f)", got, tt.expected, tt.epsilon)
			}
		})
	}
}

func TestNote_Frequency444(t *testing.T) {
	// Bright orchestral tuning (A4 = 444 Hz)
	tests := []struct {
		name     string
		note     Note
		expected float64
		epsilon  float64
	}{
		{"A4", MustNewWithOctave(A, 4), 444.0, 0.01},
		{"A3", MustNewWithOctave(A, 3), 222.0, 0.01},
		{"A5", MustNewWithOctave(A, 5), 888.0, 0.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.Frequency444()
			if !almostEqual(got, tt.expected, tt.epsilon) {
				t.Fatalf("Frequency444(): got %f, want %f", got, tt.expected)
			}
		})
	}
}

func TestNote_Frequency432(t *testing.T) {
	// Verdi tuning (A4 = 432 Hz)
	tests := []struct {
		name     string
		note     Note
		expected float64
		epsilon  float64
	}{
		{"A4", MustNewWithOctave(A, 4), 432.0, 0.01},
		{"A3", MustNewWithOctave(A, 3), 216.0, 0.01},
		{"A5", MustNewWithOctave(A, 5), 864.0, 0.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.Frequency432()
			if !almostEqual(got, tt.expected, tt.epsilon) {
				t.Fatalf("Frequency432(): got %f, want %f", got, tt.expected)
			}
		})
	}
}

func TestNote_Frequency415(t *testing.T) {
	// Baroque tuning (A4 = 415 Hz)
	tests := []struct {
		name     string
		note     Note
		expected float64
		epsilon  float64
	}{
		{"A4", MustNewWithOctave(A, 4), 415.0, 0.01},
		{"A3", MustNewWithOctave(A, 3), 207.5, 0.01},
		{"A5", MustNewWithOctave(A, 5), 830.0, 0.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.note.Frequency415()
			if !almostEqual(got, tt.expected, tt.epsilon) {
				t.Fatalf("Frequency415(): got %f, want %f", got, tt.expected)
			}
		})
	}
}

func TestNote_Frequency_CustomTuning(t *testing.T) {
	// Test with custom tuning using WithReferenceFreq
	customTuning := tuning.Standard12TET().WithReferenceFreq(442.0)
	note := MustNewWithOctave(A, 4)

	got := note.Frequency(customTuning)
	if !almostEqual(got, 442.0, 0.01) {
		t.Fatalf("Frequency() with 442 Hz tuning: got %f, want 442.0", got)
	}
}

func TestNote_Frequency_OctaveRelationships(t *testing.T) {
	// Verify that frequency doubles with each octave
	baseNote := MustNewWithOctave(A, 0)
	baseFreq := baseNote.Frequency440()

	for oct := octave.Number(1); oct <= 8; oct++ {
		note := MustNewWithOctave(A, oct)
		expectedFreq := baseFreq * math.Pow(2, float64(oct))
		got := note.Frequency440()

		if !almostEqual(got, expectedFreq, expectedFreq*0.001) {
			t.Fatalf("A%d frequency: got %f, want %f", oct, got, expectedFreq)
		}
	}
}

func TestNote_Frequency_SemitoneRatios(t *testing.T) {
	// In 12-TET, each semitone is a ratio of 2^(1/12) ≈ 1.05946
	semitoneRatio := math.Pow(2, 1.0/12.0)
	a4 := MustNewWithOctave(A, 4)
	a4Freq := a4.Frequency440()

	// Test A# (one semitone up)
	aSharp4 := MustNewWithOctave(ASHARP, 4)
	expectedASharp := a4Freq * semitoneRatio
	gotASharp := aSharp4.Frequency440()

	if !almostEqual(gotASharp, expectedASharp, 0.01) {
		t.Fatalf("A#4 frequency: got %f, want %f", gotASharp, expectedASharp)
	}

	// Test G# (one semitone down)
	gSharp4 := MustNewWithOctave(GSHARP, 4)
	expectedGSharp := a4Freq / semitoneRatio
	gotGSharp := gSharp4.Frequency440()

	if !almostEqual(gotGSharp, expectedGSharp, 0.01) {
		t.Fatalf("G#4 frequency: got %f, want %f", gotGSharp, expectedGSharp)
	}
}

func TestNote_Frequency_EnharmonicEquivalence(t *testing.T) {
	// Enharmonic notes should have the same frequency in 12-TET
	tests := []struct {
		note1, note2 Note
	}{
		{MustNewWithOctave(CSHARP, 4), MustNewWithOctave(DFLAT, 4)},
		{MustNewWithOctave(DSHARP, 4), MustNewWithOctave(EFLAT, 4)},
		{MustNewWithOctave(FSHARP, 4), MustNewWithOctave(GFLAT, 4)},
		{MustNewWithOctave(GSHARP, 4), MustNewWithOctave(AFLAT, 4)},
		{MustNewWithOctave(ASHARP, 4), MustNewWithOctave(BFLAT, 4)},
	}

	for _, tt := range tests {
		t.Run(tt.note1.String()+"="+tt.note2.String(), func(t *testing.T) {
			freq1 := tt.note1.Frequency440()
			freq2 := tt.note2.Frequency440()

			if !almostEqual(freq1, freq2, 0.01) {
				t.Fatalf("Enharmonic notes should have equal frequency: %f vs %f", freq1, freq2)
			}
		})
	}
}

// almostEqual checks if two floats are approximately equal within epsilon.
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}
