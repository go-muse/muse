package note

import (
	"errors"
	"testing"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name       string
		letter     Letter
		accidental Accidental
		wantStr    string
	}{
		{"C Natural", LetterC, Natural(), "C"},
		{"D Natural", LetterD, Natural(), "D"},
		{"E Natural", LetterE, Natural(), "E"},
		{"F Natural", LetterF, Natural(), "F"},
		{"G Natural", LetterG, Natural(), "G"},
		{"A Natural", LetterA, Natural(), "A"},
		{"B Natural", LetterB, Natural(), "B"},
		{"C Sharp", LetterC, Sharp(), "C#"},
		{"C Flat", LetterC, Flat(), "Cb"},
		{"C DoubleSharp", LetterC, DoubleSharp(), "C##"},
		{"C DoubleFlat", LetterC, DoubleFlat(), "Cbb"},
		{"F# TripleSharp", LetterF, Accidental{Chromatic: 3}, "F###"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := NewName(tt.letter, tt.accidental)
			if got := name.String(); got != tt.wantStr {
				t.Fatalf("NewName().String(): got %q, want %q", got, tt.wantStr)
			}
			if got := name.Letter(); got != tt.letter {
				t.Fatalf("NewName().Letter(): got %v, want %v", got, tt.letter)
			}
			if !name.Accidental().Equal(tt.accidental) {
				t.Fatalf("NewName().Accidental(): got %+v, want %+v", name.Accidental(), tt.accidental)
			}
		})
	}
}

func TestName_String(t *testing.T) {
	tests := []struct {
		name    Name
		wantStr string
	}{
		{C, "C"},
		{D, "D"},
		{CSHARP, "C#"},
		{DFLAT, "Db"},
		{CSHARP2, "C##"},
		{DFLAT2, "Dbb"},
		{NewName(LetterA, Accidental{Chromatic: 3}), "A###"},
		{NewName(LetterB, Accidental{Chromatic: -3}), "Bbbb"},
	}

	for _, tt := range tests {
		t.Run(tt.wantStr, func(t *testing.T) {
			if got := tt.name.String(); got != tt.wantStr {
				t.Fatalf("String(): got %q, want %q", got, tt.wantStr)
			}
		})
	}
}

func TestName_BaseName(t *testing.T) {
	tests := []struct {
		name     Name
		wantBase string
	}{
		{C, "C"},
		{CSHARP, "C"},
		{CFLAT, "C"},
		{CSHARP2, "C"},
		{CFLAT2, "C"},
		{D, "D"},
		{DSHARP, "D"},
		{DFLAT, "D"},
		{NewName(LetterA, Accidental{Chromatic: 5}), "A"},
	}

	for _, tt := range tests {
		t.Run(tt.name.String(), func(t *testing.T) {
			if got := tt.name.BaseName(); got != tt.wantBase {
				t.Fatalf("BaseName(): got %q, want %q", got, tt.wantBase)
			}
		})
	}
}

func TestName_EqualSpelling(t *testing.T) {
	tests := []struct {
		name  string
		a, b  Name
		equal bool
	}{
		{"SameNatural", C, C, true},
		{"SameSharp", CSHARP, CSHARP, true},
		{"SameFlat", CFLAT, CFLAT, true},
		{"DifferentLetter", C, D, false},
		{"DifferentAccidental", CSHARP, CFLAT, false},
		{"EnharmonicDifferentSpelling", CSHARP, DFLAT, false},
		{"SameLetterDifferentAcc", C, CSHARP, false},
		{"DoubleSharpVsSharp", CSHARP2, CSHARP, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.EqualSpelling(tt.b); got != tt.equal {
				t.Fatalf("EqualSpelling(): got %v, want %v", got, tt.equal)
			}
		})
	}
}

func TestNewNameFromString(t *testing.T) {
	tests := []struct {
		in      string
		want    Name
		wantErr bool
	}{
		// Valid natural notes
		{"C", C, false},
		{"D", D, false},
		{"E", E, false},
		{"F", F, false},
		{"G", G, false},
		{"A", A, false},
		{"B", B, false},
		// Valid sharp notes
		{"C#", CSHARP, false},
		{"D#", DSHARP, false},
		{"F#", FSHARP, false},
		// Valid flat notes
		{"Cb", CFLAT, false},
		{"Db", DFLAT, false},
		{"Eb", EFLAT, false},
		// Valid double accidentals
		{"C##", CSHARP2, false},
		{"Dbb", DFLAT2, false},
		// Valid triple accidentals
		{"C###", NewName(LetterC, Accidental{Chromatic: 3}), false},
		{"Dbbb", NewName(LetterD, Accidental{Chromatic: -3}), false},
		// Lowercase letters
		{"c", C, false},
		{"d#", DSHARP, false},
		{"eb", EFLAT, false},
		// Invalid
		{"", Name{}, true},
		{"H", Name{}, true},
		{"C#b", Name{}, true},
		{"Cb#", Name{}, true},
		{"#C", Name{}, true},
		{"bC", Name{}, true},
		{"CC", Name{}, true},
		{"1", Name{}, true},
		{"#", Name{}, true},
		{"##", Name{}, true},
		// Note: "b" is valid - it's parsed as letter B (lowercase)
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := NewNameFromString(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NewNameFromString(%q): expected error, got nil", tt.in)
				}
				if !errors.Is(err, ErrNameInvalid) {
					t.Fatalf("NewNameFromString(%q): expected ErrNameInvalid, got %v", tt.in, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewNameFromString(%q): unexpected error: %v", tt.in, err)
			}
			if !got.EqualSpelling(tt.want) {
				t.Fatalf("NewNameFromString(%q): got %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestMustNewNameFromString(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		tests := []struct {
			in   string
			want Name
		}{
			{"C", C},
			{"C#", CSHARP},
			{"Db", DFLAT},
			{"E##", ESHARP2},
			{"Fbb", FFLAT2},
		}

		for _, tt := range tests {
			got := MustNewNameFromString(tt.in)
			if !got.EqualSpelling(tt.want) {
				t.Fatalf("MustNewNameFromString(%q): got %v, want %v", tt.in, got, tt.want)
			}
		}
	})

	t.Run("Invalid_Panics", func(t *testing.T) {
		invalidInputs := []string{"", "H", "C#b", "#"}
		for _, in := range invalidInputs {
			func() {
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("MustNewNameFromString(%q): expected panic, got none", in)
					}
				}()
				_ = MustNewNameFromString(in)
			}()
		}
	})
}

func TestName_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		n     Name
		valid bool
	}{
		{"ValidC", C, true},
		{"ValidCSharp", CSHARP, true},
		{"ValidDFlat", DFLAT, true},
		{"ValidWithTripleSharp", NewName(LetterC, Accidental{Chromatic: 3}), true},
		{"InvalidLetter", NewName(Letter(255), Natural()), false},
		{"InvalidLetterWithAcc", NewName(Letter(100), Sharp()), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsValid(); got != tt.valid {
				t.Fatalf("IsValid(): got %v, want %v", got, tt.valid)
			}
		})
	}
}

func TestName_AlterUp(t *testing.T) {
	tests := []struct {
		name string
		in   Name
		want Name
	}{
		{"C->C#", C, CSHARP},
		{"C#->C##", CSHARP, CSHARP2},
		{"Cb->C", CFLAT, C},
		{"Cbb->Cb", CFLAT2, CFLAT},
		{"D->D#", D, DSHARP},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterUp()
			if !got.EqualSpelling(tt.want) {
				t.Fatalf("AlterUp(): got %v, want %v", got, tt.want)
			}
			// Original should be unchanged
			if tt.in.Letter() != got.Letter() {
				t.Fatalf("AlterUp() changed letter")
			}
		})
	}
}

func TestName_AlterDown(t *testing.T) {
	tests := []struct {
		name string
		in   Name
		want Name
	}{
		{"C->Cb", C, CFLAT},
		{"Cb->Cbb", CFLAT, CFLAT2},
		{"C#->C", CSHARP, C},
		{"C##->C#", CSHARP2, CSHARP},
		{"D->Db", D, DFLAT},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterDown()
			if !got.EqualSpelling(tt.want) {
				t.Fatalf("AlterDown(): got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_AlterUpBy(t *testing.T) {
	tests := []struct {
		name  string
		in    Name
		steps uint8
		want  Name
	}{
		{"CBy0", C, 0, C},
		{"CBy1", C, 1, CSHARP},
		{"CBy2", C, 2, CSHARP2},
		{"CbBy1", CFLAT, 1, C},
		{"CbBy2", CFLAT, 2, CSHARP},
		{"CbbBy4", CFLAT2, 4, CSHARP2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterUpBy(tt.steps)
			if !got.EqualSpelling(tt.want) {
				t.Fatalf("AlterUpBy(%d): got %v, want %v", tt.steps, got, tt.want)
			}
		})
	}
}

func TestName_AlterDownBy(t *testing.T) {
	tests := []struct {
		name  string
		in    Name
		steps uint8
		want  Name
	}{
		{"CBy0", C, 0, C},
		{"CBy1", C, 1, CFLAT},
		{"CBy2", C, 2, CFLAT2},
		{"C#By1", CSHARP, 1, C},
		{"C#By2", CSHARP, 2, CFLAT},
		{"C##By4", CSHARP2, 4, CFLAT2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterDownBy(tt.steps)
			if !got.EqualSpelling(tt.want) {
				t.Fatalf("AlterDownBy(%d): got %v, want %v", tt.steps, got, tt.want)
			}
		})
	}
}

func TestName_AlterationShift(t *testing.T) {
	tests := []struct {
		name  string
		n     Name
		shift int8
	}{
		{"Natural", C, 0},
		{"Sharp", CSHARP, 1},
		{"DoubleSharp", CSHARP2, 2},
		{"Flat", CFLAT, -1},
		{"DoubleFlat", CFLAT2, -2},
		{"TripleSharp", NewName(LetterC, Accidental{Chromatic: 3}), 3},
		{"TripleFlat", NewName(LetterC, Accidental{Chromatic: -3}), -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.AlterationShift(); got != tt.shift {
				t.Fatalf("AlterationShift(): got %d, want %d", got, tt.shift)
			}
		})
	}
}

func TestName_Letter(t *testing.T) {
	tests := []struct {
		n    Name
		want Letter
	}{
		{C, LetterC},
		{CSHARP, LetterC},
		{CFLAT, LetterC},
		{D, LetterD},
		{DSHARP, LetterD},
		{DFLAT, LetterD},
		{E, LetterE},
		{F, LetterF},
		{G, LetterG},
		{A, LetterA},
		{B, LetterB},
	}

	for _, tt := range tests {
		t.Run(tt.n.String(), func(t *testing.T) {
			if got := tt.n.Letter(); got != tt.want {
				t.Fatalf("Letter(): got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_Accidental(t *testing.T) {
	tests := []struct {
		name string
		n    Name
		want Accidental
	}{
		{"Natural", C, Natural()},
		{"Sharp", CSHARP, Sharp()},
		{"Flat", CFLAT, Flat()},
		{"DoubleSharp", CSHARP2, DoubleSharp()},
		{"DoubleFlat", CFLAT2, DoubleFlat()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Accidental(); !got.Equal(tt.want) {
				t.Fatalf("Accidental(): got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestName_NewNote(t *testing.T) {
	tests := []Name{C, CSHARP, DFLAT, E, FSHARP, GFLAT, A, BSHARP, CFLAT2, DSHARP2}

	for _, name := range tests {
		t.Run(name.String(), func(t *testing.T) {
			note := name.NewNote()
			if !note.Name().EqualSpelling(name) {
				t.Fatalf("NewNote().Name(): got %v, want %v", note.Name(), name)
			}
		})
	}
}

func TestNames_Length(t *testing.T) {
	tests := []struct {
		name    string
		names   Names
		wantLen uint64
	}{
		{"Empty", Names{}, 0},
		{"One", Names{C}, 1},
		{"Seven", Names{C, D, E, F, G, A, B}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.names.Length(); got != tt.wantLen {
				t.Fatalf("Length(): got %d, want %d", got, tt.wantLen)
			}
		})
	}
}

func TestPredefinedNames(t *testing.T) {
	// Test that all predefined names are valid and have correct properties
	naturals := []struct {
		name   Name
		letter Letter
		str    string
	}{
		{C, LetterC, "C"},
		{D, LetterD, "D"},
		{E, LetterE, "E"},
		{F, LetterF, "F"},
		{G, LetterG, "G"},
		{A, LetterA, "A"},
		{B, LetterB, "B"},
	}

	for _, tt := range naturals {
		t.Run(tt.str, func(t *testing.T) {
			if !tt.name.IsValid() {
				t.Fatalf("%s should be valid", tt.str)
			}
			if tt.name.Letter() != tt.letter {
				t.Fatalf("%s.Letter(): got %v, want %v", tt.str, tt.name.Letter(), tt.letter)
			}
			if tt.name.String() != tt.str {
				t.Fatalf("%s.String(): got %q, want %q", tt.str, tt.name.String(), tt.str)
			}
			if !tt.name.Accidental().IsNatural() {
				t.Fatalf("%s.Accidental() should be natural", tt.str)
			}
		})
	}

	// Test sharps
	sharps := []struct {
		name   Name
		letter Letter
		str    string
	}{
		{CSHARP, LetterC, "C#"},
		{DSHARP, LetterD, "D#"},
		{ESHARP, LetterE, "E#"},
		{FSHARP, LetterF, "F#"},
		{GSHARP, LetterG, "G#"},
		{ASHARP, LetterA, "A#"},
		{BSHARP, LetterB, "B#"},
	}

	for _, tt := range sharps {
		t.Run(tt.str, func(t *testing.T) {
			if !tt.name.IsValid() {
				t.Fatalf("%s should be valid", tt.str)
			}
			if tt.name.Letter() != tt.letter {
				t.Fatalf("%s.Letter(): got %v, want %v", tt.str, tt.name.Letter(), tt.letter)
			}
			if tt.name.String() != tt.str {
				t.Fatalf("%s.String(): got %q, want %q", tt.str, tt.name.String(), tt.str)
			}
			if !tt.name.Accidental().Equal(Sharp()) {
				t.Fatalf("%s.Accidental() should be Sharp", tt.str)
			}
		})
	}

	// Test flats
	flats := []struct {
		name   Name
		letter Letter
		str    string
	}{
		{CFLAT, LetterC, "Cb"},
		{DFLAT, LetterD, "Db"},
		{EFLAT, LetterE, "Eb"},
		{FFLAT, LetterF, "Fb"},
		{GFLAT, LetterG, "Gb"},
		{AFLAT, LetterA, "Ab"},
		{BFLAT, LetterB, "Bb"},
	}

	for _, tt := range flats {
		t.Run(tt.str, func(t *testing.T) {
			if !tt.name.IsValid() {
				t.Fatalf("%s should be valid", tt.str)
			}
			if tt.name.Letter() != tt.letter {
				t.Fatalf("%s.Letter(): got %v, want %v", tt.str, tt.name.Letter(), tt.letter)
			}
			if tt.name.String() != tt.str {
				t.Fatalf("%s.String(): got %q, want %q", tt.str, tt.name.String(), tt.str)
			}
			if !tt.name.Accidental().Equal(Flat()) {
				t.Fatalf("%s.Accidental() should be Flat", tt.str)
			}
		})
	}

	// Test double sharps
	doubleSharps := []Name{CSHARP2, DSHARP2, ESHARP2, FSHARP2, GSHARP2, ASHARP2, BSHARP2}
	for _, name := range doubleSharps {
		t.Run(name.String(), func(t *testing.T) {
			if !name.IsValid() {
				t.Fatalf("%s should be valid", name.String())
			}
			if !name.Accidental().Equal(DoubleSharp()) {
				t.Fatalf("%s.Accidental() should be DoubleSharp", name.String())
			}
		})
	}

	// Test double flats
	doubleFlats := []Name{CFLAT2, DFLAT2, EFLAT2, FFLAT2, GFLAT2, AFLAT2, BFLAT2}
	for _, name := range doubleFlats {
		t.Run(name.String(), func(t *testing.T) {
			if !name.IsValid() {
				t.Fatalf("%s should be valid", name.String())
			}
			if !name.Accidental().Equal(DoubleFlat()) {
				t.Fatalf("%s.Accidental() should be DoubleFlat", name.String())
			}
		})
	}
}
