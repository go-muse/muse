package note

import (
	"errors"
	"testing"
)

func TestLetter_String(t *testing.T) {
	tests := []struct {
		name string
		in   Letter
		want string
	}{
		{"C", LetterC, "C"},
		{"D", LetterD, "D"},
		{"E", LetterE, "E"},
		{"F", LetterF, "F"},
		{"G", LetterG, "G"},
		{"A", LetterA, "A"},
		{"B", LetterB, "B"},
		{"Invalid", Letter(255), "?"},
		{"InvalidAfterB", LetterB + 1, "?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.String(); got != tt.want {
				t.Fatalf("String(): got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLetter_IsValid(t *testing.T) {
	tests := []struct {
		name string
		in   Letter
		want bool
	}{
		{"LetterC", LetterC, true},
		{"LetterD", LetterD, true},
		{"LetterE", LetterE, true},
		{"LetterF", LetterF, true},
		{"LetterG", LetterG, true},
		{"LetterA", LetterA, true},
		{"LetterB", LetterB, true},
		{"Invalid255", Letter(255), false},
		{"InvalidAfterB", LetterB + 1, false},
		{"Invalid100", Letter(100), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.IsValid(); got != tt.want {
				t.Fatalf("IsValid(): got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLetter_Semitone(t *testing.T) {
	tests := []struct {
		name string
		in   Letter
		want uint8
	}{
		{"C", LetterC, 0},
		{"D", LetterD, 2},
		{"E", LetterE, 4},
		{"F", LetterF, 5},
		{"G", LetterG, 7},
		{"A", LetterA, 9},
		{"B", LetterB, 11},
		{"Invalid", Letter(255), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.Semitone(); got != tt.want {
				t.Fatalf("Semitone(): got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestParseLetter(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    Letter
		wantErr bool
	}{
		// Uppercase valid
		{"UpperC", "C", LetterC, false},
		{"UpperD", "D", LetterD, false},
		{"UpperE", "E", LetterE, false},
		{"UpperF", "F", LetterF, false},
		{"UpperG", "G", LetterG, false},
		{"UpperA", "A", LetterA, false},
		{"UpperB", "B", LetterB, false},
		// Lowercase valid
		{"LowerC", "c", LetterC, false},
		{"LowerD", "d", LetterD, false},
		{"LowerE", "e", LetterE, false},
		{"LowerF", "f", LetterF, false},
		{"LowerG", "g", LetterG, false},
		{"LowerA", "a", LetterA, false},
		{"LowerB", "b", LetterB, false},
		// Invalid
		{"Empty", "", 0, true},
		{"TwoChars", "CD", 0, true},
		{"ThreeChars", "ABC", 0, true},
		{"NonLetter", "#", 0, true},
		{"OtherLetter", "H", 0, true},
		{"OtherLetterLower", "h", 0, true},
		{"Unicode", "С", 0, true}, // Cyrillic "С"
		{"Number", "1", 0, true},
		{"Space", " ", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLetter(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseLetter(%q): expected error, got nil", tt.in)
				}
				if !errors.Is(err, ErrLetterInvalid) {
					t.Fatalf("ParseLetter(%q): expected ErrLetterInvalid, got %v", tt.in, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseLetter(%q): unexpected error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("ParseLetter(%q): got %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestMustParseLetter(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		tests := []struct {
			in   string
			want Letter
		}{
			{"C", LetterC},
			{"d", LetterD},
			{"E", LetterE},
			{"f", LetterF},
			{"G", LetterG},
			{"a", LetterA},
			{"B", LetterB},
		}

		for _, tt := range tests {
			if got := MustParseLetter(tt.in); got != tt.want {
				t.Fatalf("MustParseLetter(%q): got %v, want %v", tt.in, got, tt.want)
			}
		}
	})

	t.Run("Invalid_Panics", func(t *testing.T) {
		invalidInputs := []string{"", "H", "AB", "#", "1"}
		for _, in := range invalidInputs {
			func() {
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("MustParseLetter(%q): expected panic, got none", in)
					}
				}()
				_ = MustParseLetter(in)
			}()
		}
	})
}

func TestLetter_Natural(t *testing.T) {
	tests := []struct {
		letter  Letter
		wantStr string
		wantAcc Accidental
	}{
		{LetterC, "C", Natural()},
		{LetterD, "D", Natural()},
		{LetterE, "E", Natural()},
		{LetterF, "F", Natural()},
		{LetterG, "G", Natural()},
		{LetterA, "A", Natural()},
		{LetterB, "B", Natural()},
	}

	for _, tt := range tests {
		t.Run(tt.wantStr, func(t *testing.T) {
			name := tt.letter.Natural()
			if got := name.String(); got != tt.wantStr {
				t.Fatalf("Natural().String(): got %q, want %q", got, tt.wantStr)
			}
			if got := name.Letter(); got != tt.letter {
				t.Fatalf("Natural().Letter(): got %v, want %v", got, tt.letter)
			}
			if !name.Accidental().Equal(tt.wantAcc) {
				t.Fatalf("Natural().Accidental(): got %+v, want %+v", name.Accidental(), tt.wantAcc)
			}
		})
	}
}

func TestLetter_Sharp(t *testing.T) {
	tests := []struct {
		letter  Letter
		wantStr string
	}{
		{LetterC, "C#"},
		{LetterD, "D#"},
		{LetterE, "E#"},
		{LetterF, "F#"},
		{LetterG, "G#"},
		{LetterA, "A#"},
		{LetterB, "B#"},
	}

	for _, tt := range tests {
		t.Run(tt.wantStr, func(t *testing.T) {
			name := tt.letter.Sharp()
			if got := name.String(); got != tt.wantStr {
				t.Fatalf("Sharp().String(): got %q, want %q", got, tt.wantStr)
			}
			if got := name.Letter(); got != tt.letter {
				t.Fatalf("Sharp().Letter(): got %v, want %v", got, tt.letter)
			}
			if !name.Accidental().Equal(Sharp()) {
				t.Fatalf("Sharp().Accidental(): got %+v, want %+v", name.Accidental(), Sharp())
			}
		})
	}
}

func TestLetter_Flat(t *testing.T) {
	tests := []struct {
		letter  Letter
		wantStr string
	}{
		{LetterC, "Cb"},
		{LetterD, "Db"},
		{LetterE, "Eb"},
		{LetterF, "Fb"},
		{LetterG, "Gb"},
		{LetterA, "Ab"},
		{LetterB, "Bb"},
	}

	for _, tt := range tests {
		t.Run(tt.wantStr, func(t *testing.T) {
			name := tt.letter.Flat()
			if got := name.String(); got != tt.wantStr {
				t.Fatalf("Flat().String(): got %q, want %q", got, tt.wantStr)
			}
			if got := name.Letter(); got != tt.letter {
				t.Fatalf("Flat().Letter(): got %v, want %v", got, tt.letter)
			}
			if !name.Accidental().Equal(Flat()) {
				t.Fatalf("Flat().Accidental(): got %+v, want %+v", name.Accidental(), Flat())
			}
		})
	}
}

func TestLetter_DoubleSharp(t *testing.T) {
	tests := []struct {
		letter  Letter
		wantStr string
	}{
		{LetterC, "C##"},
		{LetterD, "D##"},
		{LetterE, "E##"},
		{LetterF, "F##"},
		{LetterG, "G##"},
		{LetterA, "A##"},
		{LetterB, "B##"},
	}

	for _, tt := range tests {
		t.Run(tt.wantStr, func(t *testing.T) {
			name := tt.letter.DoubleSharp()
			if got := name.String(); got != tt.wantStr {
				t.Fatalf("DoubleSharp().String(): got %q, want %q", got, tt.wantStr)
			}
			if got := name.Letter(); got != tt.letter {
				t.Fatalf("DoubleSharp().Letter(): got %v, want %v", got, tt.letter)
			}
			if !name.Accidental().Equal(DoubleSharp()) {
				t.Fatalf("DoubleSharp().Accidental(): got %+v, want %+v", name.Accidental(), DoubleSharp())
			}
		})
	}
}

func TestLetter_DoubleFlat(t *testing.T) {
	tests := []struct {
		letter  Letter
		wantStr string
	}{
		{LetterC, "Cbb"},
		{LetterD, "Dbb"},
		{LetterE, "Ebb"},
		{LetterF, "Fbb"},
		{LetterG, "Gbb"},
		{LetterA, "Abb"},
		{LetterB, "Bbb"},
	}

	for _, tt := range tests {
		t.Run(tt.wantStr, func(t *testing.T) {
			name := tt.letter.DoubleFlat()
			if got := name.String(); got != tt.wantStr {
				t.Fatalf("DoubleFlat().String(): got %q, want %q", got, tt.wantStr)
			}
			if got := name.Letter(); got != tt.letter {
				t.Fatalf("DoubleFlat().Letter(): got %v, want %v", got, tt.letter)
			}
			if !name.Accidental().Equal(DoubleFlat()) {
				t.Fatalf("DoubleFlat().Accidental(): got %+v, want %+v", name.Accidental(), DoubleFlat())
			}
		})
	}
}

func TestLetter_NameMethods_CreateValidNotes(t *testing.T) {
	// Test that all Letter.Name methods create notes that can be used
	letters := []Letter{LetterC, LetterD, LetterE, LetterF, LetterG, LetterA, LetterB}

	for _, letter := range letters {
		t.Run(letter.String(), func(t *testing.T) {
			// Natural
			n := letter.Natural().NewNote()
			if got := n.Name().Letter(); got != letter {
				t.Fatalf("Natural().NewNote().Name().Letter(): got %v, want %v", got, letter)
			}

			// Sharp
			n = letter.Sharp().NewNote()
			if got := n.Name().Letter(); got != letter {
				t.Fatalf("Sharp().NewNote().Name().Letter(): got %v, want %v", got, letter)
			}

			// Flat
			n = letter.Flat().NewNote()
			if got := n.Name().Letter(); got != letter {
				t.Fatalf("Flat().NewNote().Name().Letter(): got %v, want %v", got, letter)
			}

			// DoubleSharp
			n = letter.DoubleSharp().NewNote()
			if got := n.Name().Letter(); got != letter {
				t.Fatalf("DoubleSharp().NewNote().Name().Letter(): got %v, want %v", got, letter)
			}

			// DoubleFlat
			n = letter.DoubleFlat().NewNote()
			if got := n.Name().Letter(); got != letter {
				t.Fatalf("DoubleFlat().NewNote().Name().Letter(): got %v, want %v", got, letter)
			}
		})
	}
}
