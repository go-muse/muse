package temperament

import (
	"math"
	"testing"
)

func TestNewJust(t *testing.T) {
	j := NewJust()
	if j == nil {
		t.Fatal("NewJust() returned nil")
	}
}

func TestJust_Name(t *testing.T) {
	j := NewJust()
	want := "Just Intonation"
	if got := j.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestJust_Frequency_ZeroStepsPerOctave(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	got := j.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestJust_Frequency_ReferenceNote(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	got := j.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestJust_Frequency_PerfectFifth(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	// Just intonation perfect fifth is pure 3:2
	got := j.Frequency(refFreq, 7, 12)
	want := refFreq * 1.5
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 7, 12) = %v, want %v (pure fifth)", got, want)
	}
}

func TestJust_Frequency_MajorThird(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	// Just intonation major third is pure 5:4
	got := j.Frequency(refFreq, 4, 12)
	want := refFreq * 1.25
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 4, 12) = %v, want %v (pure major third)", got, want)
	}
}

func TestJust_Frequency_MinorThird(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	// Just intonation minor third is 6:5
	got := j.Frequency(refFreq, 3, 12)
	want := refFreq * 1.2
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 3, 12) = %v, want %v (pure minor third)", got, want)
	}
}

func TestJust_Frequency_AllRatios(t *testing.T) {
	j := NewJust()
	refFreq := 440.0

	expectedRatios := []float64{
		1.0,         // 0: Unison
		16.0 / 15.0, // 1: Minor 2nd
		9.0 / 8.0,   // 2: Major 2nd
		6.0 / 5.0,   // 3: Minor 3rd
		5.0 / 4.0,   // 4: Major 3rd
		4.0 / 3.0,   // 5: Perfect 4th
		45.0 / 32.0, // 6: Tritone
		3.0 / 2.0,   // 7: Perfect 5th
		8.0 / 5.0,   // 8: Minor 6th
		5.0 / 3.0,   // 9: Major 6th
		9.0 / 5.0,   // 10: Minor 7th
		15.0 / 8.0,  // 11: Major 7th
	}

	for i, ratio := range expectedRatios {
		t.Run(intervalName(i), func(t *testing.T) {
			got := j.Frequency(refFreq, i, 12)
			want := refFreq * ratio
			if math.Abs(got-want) > 0.001 {
				t.Errorf("Frequency(440, %d, 12) = %v, want %v (ratio %v)", i, got, want, ratio)
			}
		})
	}
}

func TestJust_Frequency_OctaveUp(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	got := j.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestJust_Frequency_TwoOctavesUp(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	got := j.Frequency(refFreq, 24, 12)
	want := 1760.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 12) = %v, want %v", got, want)
	}
}

func TestJust_Frequency_OctaveDown(t *testing.T) {
	j := NewJust()
	refFreq := 440.0
	got := j.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}
