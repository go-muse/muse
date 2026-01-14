package temperament

import (
	"math"
	"testing"
)

func TestNewWerckmeisterIII(t *testing.T) {
	w := NewWerckmeisterIII()
	if w == nil {
		t.Fatal("NewWerckmeisterIII() returned nil")
	}
}

func TestWerckmeisterIII_Name(t *testing.T) {
	w := NewWerckmeisterIII()
	want := "Werckmeister III"
	if got := w.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestWerckmeisterIII_Frequency_ZeroStepsPerOctave(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	got := w.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestWerckmeisterIII_Frequency_ReferenceNote(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	got := w.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestWerckmeisterIII_Frequency_OctaveUp(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	got := w.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestWerckmeisterIII_Frequency_OctaveDown(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	got := w.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}

func TestWerckmeisterIII_Frequency_PythagoreanMinorSecond(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	// In Werckmeister III, C#/Db uses Pythagorean limma 256/243
	got := w.Frequency(refFreq, 1, 12)
	want := refFreq * (256.0 / 243.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 1, 12) = %v, want %v (Pythagorean limma)", got, want)
	}
}

func TestWerckmeisterIII_Frequency_PerfectFourth(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	// In Werckmeister III, F is pure 4:3
	got := w.Frequency(refFreq, 5, 12)
	want := refFreq * (4.0 / 3.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 5, 12) = %v, want %v (pure fourth)", got, want)
	}
}

func TestWerckmeisterIII_Frequency_FallbackToEqual(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	got := w.Frequency(refFreq, 24, 24)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 24) = %v, want %v (equal temperament fallback)", got, want)
	}
}

func TestWerckmeisterIII_Frequency_NegativeSteps(t *testing.T) {
	w := NewWerckmeisterIII()
	refFreq := 440.0
	// -5 steps should be 7 steps in lower octave
	got := w.Frequency(refFreq, -5, 12)
	// Position 7 (perfect fifth) in lower octave
	expectedRatio := werckmeisterIIIRatios[7] / 2.0
	want := refFreq * expectedRatio
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -5, 12) = %v, want %v", got, want)
	}
}
