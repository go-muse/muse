package tuning

import (
	"math"
	"testing"

	"github.com/go-muse/muse/tuning/temperament"
)

func TestNew(t *testing.T) {
	tuning := New(440.0, temperament.NewEqual(), TwelveTone)

	if tuning.ReferenceFreq != 440.0 {
		t.Errorf("ReferenceFreq = %v, want 440.0", tuning.ReferenceFreq)
	}
	if tuning.ToneSystem != TwelveTone {
		t.Errorf("ToneSystem = %v, want %v", tuning.ToneSystem, TwelveTone)
	}
	if tuning.Temperament == nil {
		t.Error("Temperament is nil")
	}
}

func TestStandard12TET(t *testing.T) {
	tuning := Standard12TET()

	if tuning.ReferenceFreq != FreqA440 {
		t.Errorf("ReferenceFreq = %v, want %v", tuning.ReferenceFreq, FreqA440)
	}
	if tuning.ToneSystem != TwelveTone {
		t.Errorf("ToneSystem = %v, want %v", tuning.ToneSystem, TwelveTone)
	}
	if tuning.Temperament.Name() != "Equal Temperament" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Equal Temperament")
	}
}

func TestBaroque12TET(t *testing.T) {
	tuning := Baroque12TET()

	if tuning.ReferenceFreq != FreqA415 {
		t.Errorf("ReferenceFreq = %v, want %v", tuning.ReferenceFreq, FreqA415)
	}
	if tuning.ToneSystem != TwelveTone {
		t.Errorf("ToneSystem = %v, want %v", tuning.ToneSystem, TwelveTone)
	}
}

func TestBright12TET(t *testing.T) {
	tuning := Bright12TET()

	if tuning.ReferenceFreq != FreqA444 {
		t.Errorf("ReferenceFreq = %v, want %v", tuning.ReferenceFreq, FreqA444)
	}
}

func TestVerdi12TET(t *testing.T) {
	tuning := Verdi12TET()

	if tuning.ReferenceFreq != FreqA432 {
		t.Errorf("ReferenceFreq = %v, want %v", tuning.ReferenceFreq, FreqA432)
	}
}

func TestJustIntonation12(t *testing.T) {
	tuning := JustIntonation12()

	if tuning.Temperament.Name() != "Just Intonation" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Just Intonation")
	}
}

func TestPythagorean12(t *testing.T) {
	tuning := Pythagorean12()

	if tuning.Temperament.Name() != "Pythagorean Temperament" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Pythagorean Temperament")
	}
}

func TestMeantone12(t *testing.T) {
	tuning := Meantone12()

	if tuning.Temperament.Name() != "Quarter-Comma Meantone Temperament" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Quarter-Comma Meantone Temperament")
	}
}

func TestWerckmeisterIII12(t *testing.T) {
	tuning := WerckmeisterIII12()

	if tuning.Temperament.Name() != "Werckmeister III" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Werckmeister III")
	}
}

func TestKirnbergerIII12(t *testing.T) {
	tuning := KirnbergerIII12()

	if tuning.Temperament.Name() != "Kirnberger III" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Kirnberger III")
	}
}

func TestVallotti12(t *testing.T) {
	tuning := Vallotti12()

	if tuning.Temperament.Name() != "Vallotti" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Vallotti")
	}
}

func TestYoung12(t *testing.T) {
	tuning := Young12()

	if tuning.Temperament.Name() != "Young's Well Temperament" {
		t.Errorf("Temperament.Name() = %q, want %q", tuning.Temperament.Name(), "Young's Well Temperament")
	}
}

func TestQuarterTone24TET(t *testing.T) {
	tuning := QuarterTone24TET()

	if tuning.ToneSystem != TwentyFourTone {
		t.Errorf("ToneSystem = %v, want %v", tuning.ToneSystem, TwentyFourTone)
	}
}

func TestTuning_Frequency_A4(t *testing.T) {
	tuning := Standard12TET()
	// A4 is 0 steps from reference
	got := tuning.Frequency(0)
	want := 440.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(0) = %v, want %v", got, want)
	}
}

func TestTuning_Frequency_A5(t *testing.T) {
	tuning := Standard12TET()
	// A5 is 12 steps (one octave) up from A4
	got := tuning.Frequency(12)
	want := 880.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(12) = %v, want %v", got, want)
	}
}

func TestTuning_Frequency_A3(t *testing.T) {
	tuning := Standard12TET()
	// A3 is 12 steps down from A4
	got := tuning.Frequency(-12)
	want := 220.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("Frequency(-12) = %v, want %v", got, want)
	}
}

func TestTuning_Frequency_MiddleC(t *testing.T) {
	tuning := Standard12TET()
	// C4 is 9 semitones below A4 (A4 - G#4 - G4 - F#4 - F4 - E4 - D#4 - D4 - C#4 - C4)
	got := tuning.Frequency(-9)
	// C4 in 12-TET with A4=440 is approximately 261.63 Hz
	want := 261.63
	if math.Abs(got-want) > 0.01 {
		t.Errorf("Frequency(-9) = %v, want approximately %v", got, want)
	}
}

func TestTuning_WithReferenceFreq(t *testing.T) {
	original := Standard12TET()
	modified := original.WithReferenceFreq(442.0)

	if modified.ReferenceFreq != 442.0 {
		t.Errorf("WithReferenceFreq: ReferenceFreq = %v, want 442.0", modified.ReferenceFreq)
	}
	// Original should be unchanged
	if original.ReferenceFreq != 440.0 {
		t.Errorf("Original modified: ReferenceFreq = %v, want 440.0", original.ReferenceFreq)
	}
}

func TestTuning_WithTemperament(t *testing.T) {
	original := Standard12TET()
	modified := original.WithTemperament(temperament.NewPythagorean())

	if modified.Temperament.Name() != "Pythagorean Temperament" {
		t.Errorf("WithTemperament: Name = %q, want %q", modified.Temperament.Name(), "Pythagorean Temperament")
	}
	// Original should be unchanged
	if original.Temperament.Name() != "Equal Temperament" {
		t.Errorf("Original modified: Name = %q, want %q", original.Temperament.Name(), "Equal Temperament")
	}
}

func TestTuning_WithToneSystem(t *testing.T) {
	original := Standard12TET()
	modified := original.WithToneSystem(TwentyFourTone)

	if modified.ToneSystem != TwentyFourTone {
		t.Errorf("WithToneSystem: ToneSystem = %v, want %v", modified.ToneSystem, TwentyFourTone)
	}
	// Original should be unchanged
	if original.ToneSystem != TwelveTone {
		t.Errorf("Original modified: ToneSystem = %v, want %v", original.ToneSystem, TwelveTone)
	}
}

func TestTuning_Frequency_DifferentReferences(t *testing.T) {
	tests := []struct {
		name   string
		tuning Tuning
		wantA4 float64
	}{
		{"Standard A440", Standard12TET(), 440.0},
		{"Baroque A415", Baroque12TET(), 415.0},
		{"Bright A444", Bright12TET(), 444.0},
		{"Verdi A432", Verdi12TET(), 432.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tuning.Frequency(0)
			if math.Abs(got-tt.wantA4) > 0.001 {
				t.Errorf("Frequency(0) = %v, want %v", got, tt.wantA4)
			}
		})
	}
}
