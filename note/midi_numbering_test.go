package note

import (
	"errors"
	"testing"

	"github.com/go-muse/muse/octave"
)

func TestNote_MIDINumber(t *testing.T) {
	tests := []struct {
		name string
		note Note
		want uint8
	}{
		// Octave -1 (MIDI 0-11)
		{"C-1", MustNewWithOctave(C, -1), 0},
		{"C#-1", MustNewWithOctave(CSHARP, -1), 1},
		{"Db-1", MustNewWithOctave(DFLAT, -1), 1},
		{"D-1", MustNewWithOctave(D, -1), 2},
		{"D#-1", MustNewWithOctave(DSHARP, -1), 3},
		{"E-1", MustNewWithOctave(E, -1), 4},
		{"F-1", MustNewWithOctave(F, -1), 5},
		{"F#-1", MustNewWithOctave(FSHARP, -1), 6},
		{"G-1", MustNewWithOctave(G, -1), 7},
		{"G#-1", MustNewWithOctave(GSHARP, -1), 8},
		{"A-1", MustNewWithOctave(A, -1), 9},
		{"A#-1", MustNewWithOctave(ASHARP, -1), 10},
		{"B-1", MustNewWithOctave(B, -1), 11},

		// Octave 0 (MIDI 12-23)
		{"C0", MustNewWithOctave(C, 0), 12},
		{"C#0", MustNewWithOctave(CSHARP, 0), 13},

		// Octave 1 (MIDI 24-35)
		{"C1", MustNewWithOctave(C, 1), 24},

		// Octave 2 (MIDI 36-47)
		{"C2", MustNewWithOctave(C, 2), 36},

		// Octave 4 - Middle C and standard reference
		{"C4", MustNewWithOctave(C, 4), 60},
		{"C#4", MustNewWithOctave(CSHARP, 4), 61},
		{"D4", MustNewWithOctave(D, 4), 62},
		{"D#4", MustNewWithOctave(DSHARP, 4), 63},
		{"E4", MustNewWithOctave(E, 4), 64},
		{"F4", MustNewWithOctave(F, 4), 65},
		{"F#4", MustNewWithOctave(FSHARP, 4), 66},
		{"G4", MustNewWithOctave(G, 4), 67},
		{"G#4", MustNewWithOctave(GSHARP, 4), 68},
		{"A4", MustNewWithOctave(A, 4), 69},
		{"A#4", MustNewWithOctave(ASHARP, 4), 70},
		{"B4", MustNewWithOctave(B, 4), 71},

		// Octave 9 - highest octave (MIDI 120-127)
		{"C9", MustNewWithOctave(C, 9), 120},
		{"G9", MustNewWithOctave(G, 9), 127},

		// Enharmonic equivalents
		{"Db4=C#4", MustNewWithOctave(DFLAT, 4), 61},
		{"Eb4=D#4", MustNewWithOctave(EFLAT, 4), 63},
		{"Gb4=F#4", MustNewWithOctave(GFLAT, 4), 66},
		{"Ab4=G#4", MustNewWithOctave(AFLAT, 4), 68},
		{"Bb4=A#4", MustNewWithOctave(BFLAT, 4), 70},

		// Edge cases with alterations crossing octave boundaries
		// Note: MIDI calculation doesn't correct octave for enharmonic equivalents
		{"B#-1", MustNewWithOctave(BSHARP, -1), 0}, // B# in octave -1 wraps to 0
		{"Cb0", MustNewWithOctave(CFLAT, 0), 23},   // Cb in octave 0 = C0(12) - 1 + 12 = 23

		// Double alterations
		{"C##4", MustNewWithOctave(CSHARP2, 4), 62},
		{"Dbb4", MustNewWithOctave(DFLAT2, 4), 60},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.note.MIDINumber(); got != tt.want {
				t.Fatalf("MIDINumber(): got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNote_MIDINumber_NilOctave(t *testing.T) {
	// Note without octave should return 0
	note := New(C)
	if got := note.MIDINumber(); got != 0 {
		t.Fatalf("MIDINumber() with nil octave: got %d, want 0", got)
	}
}

func TestNote_MIDINumber_Clamping(t *testing.T) {
	tests := []struct {
		name string
		note Note
		want uint8
	}{
		// Notes with extreme alterations that would go below 0 (clamped to 0)
		// C-1 altered down by 12 would mathematically be -12, clamped to 0
		{"C-1_AlterDownBy12", MustNewWithOctave(C, -1).AlterDownBy(12), 0},
		{"C-1_AlterDownBy24", MustNewWithOctave(C, -1).AlterDownBy(24), 0},

		// Above maximum (should clamp to 127)
		{"A9", MustNewWithOctave(A, 9), 127},
		{"B9", MustNewWithOctave(B, 9), 127},
		{"G#9_AlterUp", MustNewWithOctave(GSHARP, 9).AlterUp(), 127},
		{"G9_AlterUpBy2", MustNewWithOctave(G, 9).AlterUpBy(2), 127},
		{"A9_AlterUpBy10", MustNewWithOctave(A, 9).AlterUpBy(10), 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.note.MIDINumber(); got != tt.want {
				t.Fatalf("MIDINumber(): got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNote_mustGetNoteNumberWithinOctave(t *testing.T) {
	tests := []struct {
		name Name
		want uint8
	}{
		{C, 0},
		{CSHARP, 1},
		{D, 2},
		{DSHARP, 3},
		{E, 4},
		{F, 5},
		{FSHARP, 6},
		{G, 7},
		{GSHARP, 8},
		{A, 9},
		{ASHARP, 10},
		{B, 11},
		// Flats
		{DFLAT, 1},
		{EFLAT, 3},
		{GFLAT, 6},
		{AFLAT, 8},
		{BFLAT, 10},
		// Double alterations
		{CSHARP2, 2},
		{DFLAT2, 0},
		// Wrap-around cases
		{CFLAT, 11}, // Cb = B
		{BSHARP, 0}, // B# = C
	}

	for _, tt := range tests {
		t.Run(tt.name.String(), func(t *testing.T) {
			note := New(tt.name)
			if got := note.mustGetNoteNumberWithinOctave(); got != tt.want {
				t.Fatalf("mustGetNoteNumberWithinOctave(): got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNewNoteFromMIDINumber(t *testing.T) {
	tests := []struct {
		midi       uint8
		wantName   Name
		wantOctave octave.Number
	}{
		{0, C, -1},
		{1, CSHARP, -1},
		{2, D, -1},
		{3, DSHARP, -1},
		{4, E, -1},
		{5, F, -1},
		{6, FSHARP, -1},
		{7, G, -1},
		{8, GSHARP, -1},
		{9, A, -1},
		{10, ASHARP, -1},
		{11, B, -1},
		{12, C, 0},
		{24, C, 1},
		{36, C, 2},
		{48, C, 3},
		{60, C, 4}, // Middle C
		{69, A, 4}, // A440
		{72, C, 5},
		{84, C, 6},
		{96, C, 7},
		{108, C, 8},
		{120, C, 9},
		{126, FSHARP, 9},
		{127, G, 9},
	}

	for _, tt := range tests {
		t.Run(tt.wantName.String(), func(t *testing.T) {
			note, err := NewNoteFromMIDINumber(tt.midi)
			if err != nil {
				t.Fatalf("NewNoteFromMIDINumber(%d): unexpected error: %v", tt.midi, err)
			}
			if !note.Name().EqualSpelling(tt.wantName) {
				t.Fatalf("NewNoteFromMIDINumber(%d).Name(): got %v, want %v", tt.midi, note.Name(), tt.wantName)
			}
			if note.Octave().Number() != tt.wantOctave {
				t.Fatalf("NewNoteFromMIDINumber(%d).Octave(): got %d, want %d", tt.midi, note.Octave().Number(), tt.wantOctave)
			}
		})
	}
}

func TestNewNoteFromMIDINumber_Invalid(t *testing.T) {
	invalidMIDI := []uint8{128, 129, 200, 255}

	for _, midi := range invalidMIDI {
		t.Run("", func(t *testing.T) {
			_, err := NewNoteFromMIDINumber(midi)
			if err == nil {
				t.Fatalf("NewNoteFromMIDINumber(%d): expected error, got nil", midi)
			}
			if !errors.Is(err, ErrMIDINumberUnknown) {
				t.Fatalf("NewNoteFromMIDINumber(%d): expected ErrMIDINumberUnknown, got %v", midi, err)
			}
		})
	}
}

func TestNewNoteFromMIDINumber_RoundTrip(t *testing.T) {
	// Test that MIDI -> Note -> MIDI gives the same result
	// (for notes that use sharps, since NewNoteFromMIDINumber uses sharps)
	for midi := uint8(0); midi <= 127; midi++ {
		note, err := NewNoteFromMIDINumber(midi)
		if err != nil {
			t.Fatalf("NewNoteFromMIDINumber(%d): unexpected error: %v", midi, err)
		}

		gotMIDI := note.MIDINumber()
		if gotMIDI != midi {
			t.Fatalf("Round trip MIDI %d -> Note -> MIDI: got %d", midi, gotMIDI)
		}
	}
}

func TestMIDINumber_AllOctaves(t *testing.T) {
	// Test C note across all valid octaves
	expectedMIDI := []uint8{0, 12, 24, 36, 48, 60, 72, 84, 96, 108, 120}

	for i, octNum := range []octave.Number{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		note := MustNewWithOctave(C, octNum)
		if got := note.MIDINumber(); got != expectedMIDI[i] {
			t.Fatalf("C%d.MIDINumber(): got %d, want %d", octNum, got, expectedMIDI[i])
		}
	}
}

func TestMIDINumber_ChromaticScale(t *testing.T) {
	// Test chromatic scale in octave 4 (MIDI 60-71)
	chromatic := []struct {
		name Name
		midi uint8
	}{
		{C, 60},
		{CSHARP, 61},
		{D, 62},
		{DSHARP, 63},
		{E, 64},
		{F, 65},
		{FSHARP, 66},
		{G, 67},
		{GSHARP, 68},
		{A, 69},
		{ASHARP, 70},
		{B, 71},
	}

	for _, tt := range chromatic {
		note := MustNewWithOctave(tt.name, 4)
		if got := note.MIDINumber(); got != tt.midi {
			t.Fatalf("%s4.MIDINumber(): got %d, want %d", tt.name.String(), got, tt.midi)
		}
	}
}
