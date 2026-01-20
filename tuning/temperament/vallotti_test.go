package temperament

import (
	"math"
	"testing"
)

func TestNewVallotti(t *testing.T) {
	v := NewVallotti()
	if v == nil {
		t.Fatal("NewVallotti() returned nil")
	}
}

func TestVallotti_Name(t *testing.T) {
	v := NewVallotti()
	want := "Vallotti"
	if got := v.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestVallotti_Frequency_ZeroStepsPerOctave(t *testing.T) {
	v := NewVallotti()
	refFreq := 440.0
	got := v.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestVallotti_Frequency_ReferenceNote(t *testing.T) {
	v := NewVallotti()
	refFreq := 440.0
	got := v.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestVallotti_Frequency_PerfectFourth(t *testing.T) {
	v := NewVallotti()
	refFreq := 440.0
	// Vallotti F is pure 4:3
	got := v.Frequency(refFreq, 5, 12)
	want := refFreq * (4.0 / 3.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 5, 12) = %v, want %v (pure fourth)", got, want)
	}
}

func TestVallotti_Frequency_MinorSeventh(t *testing.T) {
	v := NewVallotti()
	refFreq := 440.0
	// Vallotti A#/Bb is pure 16:9
	got := v.Frequency(refFreq, 10, 12)
	want := refFreq * (16.0 / 9.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 10, 12) = %v, want %v (pure 16:9)", got, want)
	}
}

func TestVallotti_Frequency_OctaveUp(t *testing.T) {
	v := NewVallotti()
	refFreq := 440.0
	got := v.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestVallotti_Frequency_OctaveDown(t *testing.T) {
	v := NewVallotti()
	refFreq := 440.0
	got := v.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}

func TestVallotti_Frequency_FallbackToEqual(t *testing.T) {
	v := NewVallotti()
	refFreq := 440.0
	got := v.Frequency(refFreq, 24, 24)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 24) = %v, want %v (equal temperament fallback)", got, want)
	}
}
