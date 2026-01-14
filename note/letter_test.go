package note

import (
	"errors"
	"testing"
)

func TestLetter_String_Valid(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.String(); got != tt.want {
				t.Fatalf("String(): got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLetter_String_Invalid(t *testing.T) {
	var invalid Letter = 255
	if got, want := invalid.String(), "?"; got != want {
		t.Fatalf("String() for invalid: got %q, want %q", got, want)
	}
}

func TestLetter_IsValid(t *testing.T) {
	tests := []struct {
		name string
		in   Letter
		want bool
	}{
		{"LetterC", LetterC, true},
		{"LetterB", LetterB, true},
		{"Invalid255", Letter(255), false},
		{"InvalidAfterB", LetterB + 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.IsValid(); got != tt.want {
				t.Fatalf("IsValid(): got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLetter_Valid(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Letter
	}{
		{"C", "C", LetterC},
		{"D", "D", LetterD},
		{"E", "E", LetterE},
		{"F", "F", LetterF},
		{"G", "G", LetterG},
		{"A", "A", LetterA},
		{"B", "B", LetterB},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLetter(tt.in)
			if err != nil {
				t.Fatalf("ParseLetter(%q): unexpected error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("ParseLetter(%q): got %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseLetter_InvalidLength(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"Empty", ""},
		{"TwoChars", "CD"},
		{"ThreeChars", "ABC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseLetter(tt.in)
			if err == nil {
				t.Fatalf("ParseLetter(%q): expected error, got nil", tt.in)
			}

			if !errors.Is(err, ErrInvalidLetter) {
				t.Fatalf("ParseLetter(%q): expected ErrInvalidLetter, got %v", tt.in, err)
			}
		})
	}
}

func TestParseLetter_InvalidChar(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"Lowercase", "c"},
		{"NonLetter", "#"},
		{"OtherLetter", "H"},
		{"Unicode", "С"}, // cyrillic "С"
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseLetter(tt.in)
			if err == nil {
				t.Fatalf("ParseLetter(%q): expected error, got nil", tt.in)
			}
			if !errors.Is(err, ErrInvalidLetter) {
				t.Fatalf("ParseLetter(%q): expected ErrInvalidLetter, got %v", tt.in, err)
			}
		})
	}
}

func TestMustParseLetter_Valid(t *testing.T) {
	if got, want := MustParseLetter("C"), LetterC; got != want {
		t.Fatalf("MustParseLetter(%q): got %v, want %v", "C", got, want)
	}
}

func TestMustParseLetter_Invalid_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("MustParseLetter(%q): expected panic, got none", "H")
		}
	}()

	_ = MustParseLetter("H")
}
