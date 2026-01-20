package temperament

import (
	"math"
	"testing"
)

func TestNewEqual(t *testing.T) {
	eq := NewEqual()
	if eq == nil {
		t.Fatal("NewEqual() returned nil")
	}
}

func TestEqual_Name(t *testing.T) {
	eq := NewEqual()
	want := "Equal Temperament"
	if got := eq.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestEqual_Frequency_ZeroStepsPerOctave(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0
	got := eq.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestEqual_Frequency_ReferenceNote(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0
	// 0 steps from reference should return reference frequency
	got := eq.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestEqual_Frequency_OctaveUp(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0
	// 12 steps up = 1 octave = double frequency
	got := eq.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestEqual_Frequency_OctaveDown(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0
	// -12 steps = 1 octave down = half frequency
	got := eq.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}

func TestEqual_Frequency_PerfectFifth(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0
	// 7 semitones up = perfect fifth
	got := eq.Frequency(refFreq, 7, 12)
	// In 12-TET, perfect fifth ratio is 2^(7/12) ≈ 1.498
	want := refFreq * math.Pow(2, 7.0/12.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 7, 12) = %v, want %v", got, want)
	}
}

func TestEqual_Frequency_24TET(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0
	// 24 steps in 24-TET = 1 octave
	got := eq.Frequency(refFreq, 24, 24)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 24) = %v, want %v", got, want)
	}
}

func TestEqual_Frequency_QuarterTone(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0
	// 1 step in 24-TET = quarter tone
	got := eq.Frequency(refFreq, 1, 24)
	want := refFreq * math.Pow(2, 1.0/24.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 1, 24) = %v, want %v", got, want)
	}
}

func TestEqual_Frequency_VariousIntervals(t *testing.T) {
	eq := NewEqual()
	refFreq := 440.0

	tests := []struct {
		name     string
		steps    int
		wantFreq float64
	}{
		{"unison", 0, 440.0},
		{"minor second", 1, refFreq * math.Pow(2, 1.0/12.0)},
		{"major second", 2, refFreq * math.Pow(2, 2.0/12.0)},
		{"minor third", 3, refFreq * math.Pow(2, 3.0/12.0)},
		{"major third", 4, refFreq * math.Pow(2, 4.0/12.0)},
		{"perfect fourth", 5, refFreq * math.Pow(2, 5.0/12.0)},
		{"tritone", 6, refFreq * math.Pow(2, 6.0/12.0)},
		{"perfect fifth", 7, refFreq * math.Pow(2, 7.0/12.0)},
		{"octave", 12, 880.0},
		{"two octaves", 24, 1760.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := eq.Frequency(refFreq, tt.steps, 12)
			if math.Abs(got-tt.wantFreq) > 0.001 {
				t.Errorf("Frequency(440, %d, 12) = %v, want %v", tt.steps, got, tt.wantFreq)
			}
		})
	}
}
