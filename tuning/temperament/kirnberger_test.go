package temperament

import (
	"math"
	"testing"
)

func TestNewKirnbergerIII(t *testing.T) {
	k := NewKirnbergerIII()
	if k == nil {
		t.Fatal("NewKirnbergerIII() returned nil")
	}
}

func TestKirnbergerIII_Name(t *testing.T) {
	k := NewKirnbergerIII()
	want := "Kirnberger III"
	if got := k.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestKirnbergerIII_Frequency_ZeroStepsPerOctave(t *testing.T) {
	k := NewKirnbergerIII()
	refFreq := 440.0
	got := k.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestKirnbergerIII_Frequency_ReferenceNote(t *testing.T) {
	k := NewKirnbergerIII()
	refFreq := 440.0
	got := k.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestKirnbergerIII_Frequency_MajorThird(t *testing.T) {
	k := NewKirnbergerIII()
	refFreq := 440.0
	// Kirnberger III has pure major third 5:4 on C-E
	got := k.Frequency(refFreq, 4, 12)
	want := refFreq * 1.25
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 4, 12) = %v, want %v (pure major third)", got, want)
	}
}

func TestKirnbergerIII_Frequency_PerfectFourth(t *testing.T) {
	k := NewKirnbergerIII()
	refFreq := 440.0
	// Perfect fourth is pure 4:3
	got := k.Frequency(refFreq, 5, 12)
	want := refFreq * (4.0 / 3.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 5, 12) = %v, want %v (pure fourth)", got, want)
	}
}

func TestKirnbergerIII_Frequency_OctaveUp(t *testing.T) {
	k := NewKirnbergerIII()
	refFreq := 440.0
	got := k.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestKirnbergerIII_Frequency_OctaveDown(t *testing.T) {
	k := NewKirnbergerIII()
	refFreq := 440.0
	got := k.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}

func TestKirnbergerIII_Frequency_FallbackToEqual(t *testing.T) {
	k := NewKirnbergerIII()
	refFreq := 440.0
	got := k.Frequency(refFreq, 24, 24)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 24) = %v, want %v (equal temperament fallback)", got, want)
	}
}
