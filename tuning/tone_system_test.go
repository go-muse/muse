package tuning

import (
	"math"
	"testing"
)

func TestToneSystem_StepsPerOctave(t *testing.T) {
	tests := []struct {
		name       string
		toneSystem ToneSystem
		want       int
	}{
		{"TwelveTone", TwelveTone, 12},
		{"NineteenTone", NineteenTone, 19},
		{"TwentyFourTone", TwentyFourTone, 24},
		{"ThirtyOneTone", ThirtyOneTone, 31},
		{"Zero", ToneSystem(0), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.toneSystem.StepsPerOctave(); got != tt.want {
				t.Errorf("StepsPerOctave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToneSystem_CentsPerStep(t *testing.T) {
	tests := []struct {
		name       string
		toneSystem ToneSystem
		want       float64
	}{
		{"TwelveTone", TwelveTone, 100.0},
		{"TwentyFourTone", TwentyFourTone, 50.0},
		{"NineteenTone", NineteenTone, 1200.0 / 19.0},
		{"ThirtyOneTone", ThirtyOneTone, 1200.0 / 31.0},
		{"Zero", ToneSystem(0), 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.toneSystem.CentsPerStep()
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("CentsPerStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToneSystem_Constants(t *testing.T) {
	if TwelveTone != 12 {
		t.Errorf("TwelveTone = %v, want 12", TwelveTone)
	}
	if NineteenTone != 19 {
		t.Errorf("NineteenTone = %v, want 19", NineteenTone)
	}
	if TwentyFourTone != 24 {
		t.Errorf("TwentyFourTone = %v, want 24", TwentyFourTone)
	}
	if ThirtyOneTone != 31 {
		t.Errorf("ThirtyOneTone = %v, want 31", ThirtyOneTone)
	}
}

func TestCentsPerOctave(t *testing.T) {
	if CentsPerOctave != 1200.0 {
		t.Errorf("CentsPerOctave = %v, want 1200.0", CentsPerOctave)
	}
}

func TestSemitonesPerOctave(t *testing.T) {
	if SemitonesPerOctave != 12 {
		t.Errorf("SemitonesPerOctave = %v, want 12", SemitonesPerOctave)
	}
}

func TestReferenceFrequencies(t *testing.T) {
	tests := []struct {
		name string
		freq float64
		want float64
	}{
		{"FreqA440", FreqA440, 440.0},
		{"FreqA444", FreqA444, 444.0},
		{"FreqA432", FreqA432, 432.0},
		{"FreqA415", FreqA415, 415.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.freq != tt.want {
				t.Errorf("%s = %v, want %v", tt.name, tt.freq, tt.want)
			}
		})
	}
}
