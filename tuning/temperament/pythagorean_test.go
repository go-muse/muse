package temperament

import (
	"math"
	"testing"
)

func TestNewPythagorean(t *testing.T) {
	p := NewPythagorean()
	if p == nil {
		t.Fatal("NewPythagorean() returned nil")
	}
}

func TestPythagorean_Name(t *testing.T) {
	p := NewPythagorean()
	want := "Pythagorean Temperament"
	if got := p.Name(); got != want {
		t.Errorf("Name() = %q, want %q", got, want)
	}
}

func TestPythagorean_Frequency_ZeroStepsPerOctave(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	got := p.Frequency(refFreq, 0, 0)
	if got != refFreq {
		t.Errorf("Frequency with stepsPerOctave=0: got %v, want %v", got, refFreq)
	}
}

func TestPythagorean_Frequency_ReferenceNote(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	got := p.Frequency(refFreq, 0, 12)
	if got != refFreq {
		t.Errorf("Frequency(440, 0, 12) = %v, want %v", got, refFreq)
	}
}

func TestPythagorean_Frequency_PerfectFifth(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	// Pythagorean perfect fifth is pure 3:2
	got := p.Frequency(refFreq, 7, 12)
	want := refFreq * 1.5 // 3:2 ratio
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 7, 12) = %v, want %v (pure fifth)", got, want)
	}
}

func TestPythagorean_Frequency_PerfectFourth(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	// Pythagorean perfect fourth is pure 4:3
	got := p.Frequency(refFreq, 5, 12)
	want := refFreq * (4.0 / 3.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 5, 12) = %v, want %v (pure fourth)", got, want)
	}
}

func TestPythagorean_Frequency_MajorSecond(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	// Pythagorean major second is 9:8
	got := p.Frequency(refFreq, 2, 12)
	want := refFreq * (9.0 / 8.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 2, 12) = %v, want %v (9:8)", got, want)
	}
}

func TestPythagorean_Frequency_MajorThird(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	// Pythagorean major third (ditone) is 81:64
	got := p.Frequency(refFreq, 4, 12)
	want := refFreq * (81.0 / 64.0)
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 4, 12) = %v, want %v (81:64 ditone)", got, want)
	}
}

func TestPythagorean_Frequency_OctaveUp(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	got := p.Frequency(refFreq, 12, 12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 12, 12) = %v, want %v", got, want)
	}
}

func TestPythagorean_Frequency_OctaveDown(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	got := p.Frequency(refFreq, -12, 12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -12, 12) = %v, want %v", got, want)
	}
}

func TestPythagorean_Frequency_NegativeSteps(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	// -7 steps = perfect fifth down (position 5 in lower octave)
	got := p.Frequency(refFreq, -7, 12)
	// -7 steps means position 5 in lower octave (12-7=5), with octave -1
	// Position 5 = perfect fourth ratio (4/3)
	want := (refFreq / 2) * (4.0 / 3.0) // Fourth in lower octave = 293.33 Hz
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, -7, 12) = %v, want %v", got, want)
	}
}

func TestPythagorean_Frequency_FallbackToEqual(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0
	// For non-12-tone systems, falls back to equal temperament
	got := p.Frequency(refFreq, 24, 24)
	want := 880.0 // One octave up
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(440, 24, 24) = %v, want %v (equal temperament fallback)", got, want)
	}
}

func TestPythagorean_Frequency_AllRatios(t *testing.T) {
	p := NewPythagorean()
	refFreq := 440.0

	expectedRatios := []float64{
		1.0,             // 0: Unison
		256.0 / 243.0,   // 1: Minor 2nd
		9.0 / 8.0,       // 2: Major 2nd
		32.0 / 27.0,     // 3: Minor 3rd
		81.0 / 64.0,     // 4: Major 3rd
		4.0 / 3.0,       // 5: Perfect 4th
		729.0 / 512.0,   // 6: Tritone
		3.0 / 2.0,       // 7: Perfect 5th
		128.0 / 81.0,    // 8: Minor 6th
		27.0 / 16.0,     // 9: Major 6th
		16.0 / 9.0,      // 10: Minor 7th
		243.0 / 128.0,   // 11: Major 7th
	}

	for i, ratio := range expectedRatios {
		t.Run(intervalName(i), func(t *testing.T) {
			got := p.Frequency(refFreq, i, 12)
			want := refFreq * ratio
			if math.Abs(got-want) > 0.001 {
				t.Errorf("Frequency(440, %d, 12) = %v, want %v (ratio %v)", i, got, want, ratio)
			}
		})
	}
}

func intervalName(step int) string {
	names := []string{
		"unison", "minor_2nd", "major_2nd", "minor_3rd",
		"major_3rd", "perfect_4th", "tritone", "perfect_5th",
		"minor_6th", "major_6th", "minor_7th", "major_7th",
	}
	if step >= 0 && step < len(names) {
		return names[step]
	}
	return "unknown"
}
