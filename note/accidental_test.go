package note

import (
	"errors"
	"testing"
)

func TestAccidental_Constructors(t *testing.T) {
	tests := []struct {
		name string
		got  Accidental
		want Accidental
	}{
		{"Natural", Natural(), Accidental{Chromatic: 0, Micro: 0}},
		{"Flat", Flat(), Accidental{Chromatic: -1, Micro: 0}},
		{"DoubleFlat", DoubleFlat(), Accidental{Chromatic: -2, Micro: 0}},
		{"Sharp", Sharp(), Accidental{Chromatic: 1, Micro: 0}},
		{"DoubleSharp", DoubleSharp(), Accidental{Chromatic: 2, Micro: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.got.Equal(tt.want) {
				t.Fatalf("got %+v, want %+v", tt.got, tt.want)
			}
		})
	}
}

func TestAccidental_IsNatural(t *testing.T) {
	tests := []struct {
		name string
		acc  Accidental
		want bool
	}{
		{"Natural", Natural(), true},
		{"NaturalZeroValue", Accidental{}, true},
		{"Flat", Flat(), false},
		{"Sharp", Sharp(), false},
		{"DoubleFlat", DoubleFlat(), false},
		{"DoubleSharp", DoubleSharp(), false},
		{"NaturalWithMicro", Accidental{Chromatic: 0, Micro: 1}, false},
		{"NaturalWithNegativeMicro", Accidental{Chromatic: 0, Micro: -1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.acc.IsNatural(); got != tt.want {
				t.Fatalf("IsNatural(): got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccidental_Equal(t *testing.T) {
	tests := []struct {
		name string
		a, b Accidental
		want bool
	}{
		{"IdenticalSharps", Accidental{Chromatic: 1, Micro: 0}, Accidental{Chromatic: 1, Micro: 0}, true},
		{"IdenticalNaturals", Natural(), Natural(), true},
		{"DifferentChromatic", Accidental{Chromatic: 1, Micro: 0}, Accidental{Chromatic: 2, Micro: 0}, false},
		{"DifferentMicro", Accidental{Chromatic: 1, Micro: 0}, Accidental{Chromatic: 1, Micro: 1}, false},
		{"BothDifferent", Accidental{Chromatic: 1, Micro: 0}, Accidental{Chromatic: 2, Micro: 1}, false},
		{"NegativeChromatic", Accidental{Chromatic: -2, Micro: 0}, DoubleFlat(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Fatalf("Equal(): got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccidental_String(t *testing.T) {
	tests := []struct {
		name string
		acc  Accidental
		want string
	}{
		// Standard accidentals
		{"DoubleFlat", DoubleFlat(), "bb"},
		{"Flat", Flat(), "b"},
		{"Natural", Natural(), ""},
		{"Sharp", Sharp(), "#"},
		{"DoubleSharp", DoubleSharp(), "##"},
		// Extended accidentals (triple and beyond)
		{"TripleSharp", Accidental{Chromatic: 3}, "###"},
		{"TripleFlat", Accidental{Chromatic: -3}, "bbb"},
		{"QuadrupleSharp", Accidental{Chromatic: 4}, "####"},
		{"QuadrupleFlat", Accidental{Chromatic: -4}, "bbbb"},
		{"ManySharps", Accidental{Chromatic: 10}, "##########"},
		{"ManyFlats", Accidental{Chromatic: -10}, "bbbbbbbbbb"},
		// Micro is ignored in string representation
		{"SharpWithMicro", Accidental{Chromatic: 1, Micro: 1234}, "#"},
		{"NaturalWithMicro", Accidental{Chromatic: 0, Micro: 500}, ""},
		{"FlatWithMicro", Accidental{Chromatic: -1, Micro: -500}, "b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.acc.String(); got != tt.want {
				t.Fatalf("String(): got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAccidental_AlterUp(t *testing.T) {
	tests := []struct {
		name string
		in   Accidental
		want Accidental
	}{
		{"NaturalToSharp", Natural(), Sharp()},
		{"SharpToDoubleSharp", Sharp(), DoubleSharp()},
		{"FlatToNatural", Flat(), Natural()},
		{"DoubleFlatToFlat", DoubleFlat(), Flat()},
		{"TripleFlatToDoubleFlat", Accidental{Chromatic: -3}, DoubleFlat()},
		{"DoubleSharpToTripleSharp", DoubleSharp(), Accidental{Chromatic: 3}},
		// Micro should be preserved
		{"WithMicro", Accidental{Chromatic: 0, Micro: 100}, Accidental{Chromatic: 1, Micro: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterUp()
			if !got.Equal(tt.want) {
				t.Fatalf("AlterUp(): got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestAccidental_AlterDown(t *testing.T) {
	tests := []struct {
		name string
		in   Accidental
		want Accidental
	}{
		{"NaturalToFlat", Natural(), Flat()},
		{"FlatToDoubleFlat", Flat(), DoubleFlat()},
		{"SharpToNatural", Sharp(), Natural()},
		{"DoubleSharpToSharp", DoubleSharp(), Sharp()},
		{"TripleSharpToDoubleSharp", Accidental{Chromatic: 3}, DoubleSharp()},
		{"DoubleFlatToTripleFlat", DoubleFlat(), Accidental{Chromatic: -3}},
		// Micro should be preserved
		{"WithMicro", Accidental{Chromatic: 0, Micro: 100}, Accidental{Chromatic: -1, Micro: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterDown()
			if !got.Equal(tt.want) {
				t.Fatalf("AlterDown(): got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestAccidental_AlterUpBy(t *testing.T) {
	tests := []struct {
		name  string
		in    Accidental
		steps uint8
		want  Accidental
	}{
		{"NaturalBy0", Natural(), 0, Natural()},
		{"NaturalBy1", Natural(), 1, Sharp()},
		{"NaturalBy2", Natural(), 2, DoubleSharp()},
		{"NaturalBy3", Natural(), 3, Accidental{Chromatic: 3}},
		{"FlatBy1", Flat(), 1, Natural()},
		{"FlatBy2", Flat(), 2, Sharp()},
		{"DoubleFlatBy4", DoubleFlat(), 4, DoubleSharp()},
		// Micro should be preserved
		{"WithMicro", Accidental{Chromatic: 0, Micro: 100}, 2, Accidental{Chromatic: 2, Micro: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterUpBy(tt.steps)
			if !got.Equal(tt.want) {
				t.Fatalf("AlterUpBy(%d): got %+v, want %+v", tt.steps, got, tt.want)
			}
		})
	}
}

func TestAccidental_AlterDownBy(t *testing.T) {
	tests := []struct {
		name  string
		in    Accidental
		steps uint8
		want  Accidental
	}{
		{"NaturalBy0", Natural(), 0, Natural()},
		{"NaturalBy1", Natural(), 1, Flat()},
		{"NaturalBy2", Natural(), 2, DoubleFlat()},
		{"NaturalBy3", Natural(), 3, Accidental{Chromatic: -3}},
		{"SharpBy1", Sharp(), 1, Natural()},
		{"SharpBy2", Sharp(), 2, Flat()},
		{"DoubleSharpBy4", DoubleSharp(), 4, DoubleFlat()},
		// Micro should be preserved
		{"WithMicro", Accidental{Chromatic: 0, Micro: 100}, 2, Accidental{Chromatic: -2, Micro: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterDownBy(tt.steps)
			if !got.Equal(tt.want) {
				t.Fatalf("AlterDownBy(%d): got %+v, want %+v", tt.steps, got, tt.want)
			}
		})
	}
}

func TestParseAccidental(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    Accidental
		wantErr bool
	}{
		// Valid standard accidentals
		{"Natural", "", Natural(), false},
		{"Flat", "b", Flat(), false},
		{"DoubleFlat", "bb", DoubleFlat(), false},
		{"Sharp", "#", Sharp(), false},
		{"DoubleSharp", "##", DoubleSharp(), false},
		// Valid extended accidentals
		{"TripleFlat", "bbb", Accidental{Chromatic: -3}, false},
		{"QuadrupleFlat", "bbbb", Accidental{Chromatic: -4}, false},
		{"TripleSharp", "###", Accidental{Chromatic: 3}, false},
		{"QuadrupleSharp", "####", Accidental{Chromatic: 4}, false},
		{"ManySharps", "##########", Accidental{Chromatic: 10}, false},
		{"ManyFlats", "bbbbbbbbbb", Accidental{Chromatic: -10}, false},
		// Invalid
		{"UnicodeSharp", "♯", Accidental{}, true},
		{"UnicodeFlat", "♭", Accidental{}, true},
		{"UnicodeDoubleSharp", "𝄪", Accidental{}, true},
		{"UnicodeDoubleFlat", "𝄫", Accidental{}, true},
		{"Whitespace", " ", Accidental{}, true},
		{"Random", "x", Accidental{}, true},
		{"Mixed", "b#", Accidental{}, true},
		{"MixedReverse", "#b", Accidental{}, true},
		{"Letter", "A", Accidental{}, true},
		{"Number", "1", Accidental{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAccidental(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseAccidental(%q): expected error, got nil", tt.in)
				}
				if !errors.Is(err, ErrInvalidAccidental) {
					t.Fatalf("ParseAccidental(%q): expected ErrInvalidAccidental, got %v", tt.in, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseAccidental(%q): unexpected error: %v", tt.in, err)
			}
			if !got.Equal(tt.want) {
				t.Fatalf("ParseAccidental(%q): got %+v, want %+v", tt.in, got, tt.want)
			}
		})
	}
}

func TestMustParseAccidental(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		tests := []struct {
			in   string
			want Accidental
		}{
			{"", Natural()},
			{"b", Flat()},
			{"bb", DoubleFlat()},
			{"#", Sharp()},
			{"##", DoubleSharp()},
			{"###", Accidental{Chromatic: 3}},
			{"bbb", Accidental{Chromatic: -3}},
		}

		for _, tt := range tests {
			got := MustParseAccidental(tt.in)
			if !got.Equal(tt.want) {
				t.Fatalf("MustParseAccidental(%q): got %+v, want %+v", tt.in, got, tt.want)
			}
		}
	})

	t.Run("Invalid_Panics", func(t *testing.T) {
		invalidInputs := []string{"x", "♯", "b#", " "}
		for _, in := range invalidInputs {
			func() {
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("MustParseAccidental(%q): expected panic, got none", in)
					}
				}()
				_ = MustParseAccidental(in)
			}()
		}
	})
}

func TestGetNotesWithAlterations(t *testing.T) {
	tests := []struct {
		name        string
		notes       Notes
		alterations uint8
		wantLen     int
	}{
		{
			name:        "NoAlterations",
			notes:       NewFromNoteNames(C, D, E, F, G, A, B),
			alterations: 0,
			wantLen:     0,
		},
		{
			name:        "OneAlteration",
			notes:       NewFromNoteNames(C, D, E, F, G, A, B),
			alterations: 1,
			wantLen:     14, // 7 notes × 2 alterations (up and down)
		},
		{
			name:        "TwoAlterations",
			notes:       NewFromNoteNames(C, D, E, F, G, A, B),
			alterations: 2,
			wantLen:     28, // 7 notes × 4 alterations (up1, down1, up2, down2)
		},
		{
			name:        "EmptyNotes",
			notes:       Notes{},
			alterations: 2,
			wantLen:     0,
		},
		{
			name:        "SingleNote",
			notes:       NewFromNoteNames(C),
			alterations: 3,
			wantLen:     6, // 1 note × 6 alterations
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetNotesWithAlterations(tt.notes, tt.alterations)
			if got := len(result); got != tt.wantLen {
				t.Fatalf("GetNotesWithAlterations(): got %d notes, want %d", got, tt.wantLen)
			}
		})
	}
}

func TestGetNotesWithAlterations_Content(t *testing.T) {
	// Test with C note and 1 alteration
	notes := NewFromNoteNames(C)
	result := GetNotesWithAlterations(notes, 1)

	if len(result) != 2 {
		t.Fatalf("expected 2 notes, got %d", len(result))
	}

	// Should contain C# and Cb
	hasSharp := false
	hasFlat := false
	for _, n := range result {
		if n.Name().EqualSpelling(CSHARP) {
			hasSharp = true
		}
		if n.Name().EqualSpelling(CFLAT) {
			hasFlat = true
		}
	}

	if !hasSharp {
		t.Fatalf("expected C# in results")
	}
	if !hasFlat {
		t.Fatalf("expected Cb in results")
	}
}

func TestGetNotesWithAlterations_TwoAlterations(t *testing.T) {
	// Test with C note and 2 alterations
	notes := NewFromNoteNames(C)
	result := GetNotesWithAlterations(notes, 2)

	if len(result) != 4 {
		t.Fatalf("expected 4 notes, got %d", len(result))
	}

	// Should contain C#, C##, Cb, Cbb
	expected := []Name{CSHARP, CSHARP2, CFLAT, CFLAT2}
	for _, exp := range expected {
		found := false
		for _, n := range result {
			if n.Name().EqualSpelling(exp) {
				found = true

				break
			}
		}
		if !found {
			t.Fatalf("expected %s in results", exp.String())
		}
	}
}
