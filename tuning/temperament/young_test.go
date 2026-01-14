package temperament

import (
	"math"
	"testing"
)

func TestNewYoung(t *testing.T) {
	y := NewYoung()
	if y == nil {
		t.Fatal("NewYoung() returned nil")
	}
}

func TestYoung_Name(t *testing.T) {
	y := NewYoung()
	want := "Young's Well Temperament"
	if got := y.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestYoung_Frequency_ZeroStepsPerOctave(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	got := y.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestYoung_Frequency_ReferenceNote(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	got := y.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestYoung_Frequency_PerfectFourth(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	// Young F is pure 4:3
	got := y.Frequency(refFreq, 5, 12)
	want := refFreq * (4.0 / 3.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 5, 12) = %v, want %v (pure fourth)", got, want)
	}
}

func TestYoung_Frequency_MinorSeventh(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	// Young A#/Bb is pure 16:9
	got := y.Frequency(refFreq, 10, 12)
	want := refFreq * (16.0 / 9.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 10, 12) = %v, want %v (pure 16:9)", got, want)
	}
}

func TestYoung_Frequency_OctaveUp(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	got := y.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestYoung_Frequency_OctaveDown(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	got := y.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}

func TestYoung_Frequency_FallbackToEqual(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	got := y.Frequency(refFreq, 24, 24)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 24) = %v, want %v (equal temperament fallback)", got, want)
	}
}

func TestYoung_Frequency_NegativeSteps(t *testing.T) {
	y := NewYoung()
	refFreq := 440.0
	// -5 steps should be 7 steps in lower octave
	got := y.Frequency(refFreq, -5, 12)
	expectedRatio := youngRatios[7] / 2.0
	want := refFreq * expectedRatio
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -5, 12) = %v, want %v", got, want)
	}
}
