package temperament

import (
	"math"
	"testing"
)

func TestNewMeantone(t *testing.T) {
	m := NewMeantone()
	if m == nil {
		t.Fatal("NewMeantone() returned nil")
	}
}

func TestMeantone_Name(t *testing.T) {
	m := NewMeantone()
	want := "Quarter-Comma Meantone Temperament"
	if got := m.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestMeantone_Frequency_ZeroStepsPerOctave(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	got := m.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestMeantone_Frequency_ReferenceNote(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	got := m.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestMeantone_Frequency_MajorThird(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	// Quarter-comma meantone major third is pure 5:4
	got := m.Frequency(refFreq, 4, 12)
	want := refFreq * 1.25
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 4, 12) = %v, want %v (pure major third)", got, want)
	}
}

func TestMeantone_Frequency_PerfectFifth(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	// Quarter-comma meantone fifth is tempered: 5^(1/4) ≈ 1.495349
	got := m.Frequency(refFreq, 7, 12)
	want := refFreq * math.Pow(5, 0.25)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 7, 12) = %v, want %v (tempered fifth)", got, want)
	}
}

func TestMeantone_Frequency_MinorSixth(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	// Quarter-comma meantone minor sixth is pure 8:5
	got := m.Frequency(refFreq, 8, 12)
	want := refFreq * 1.6
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 8, 12) = %v, want %v (pure minor sixth)", got, want)
	}
}

func TestMeantone_Frequency_OctaveUp(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	got := m.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestMeantone_Frequency_OctaveDown(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	got := m.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}

func TestMeantone_Frequency_FallbackToEqual(t *testing.T) {
	m := NewMeantone()
	refFreq := 440.0
	// For non-12-tone systems, falls back to equal temperament
	got := m.Frequency(refFreq, 24, 24)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 24) = %v, want %v (equal temperament fallback)", got, want)
	}
}
