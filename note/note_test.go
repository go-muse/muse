package note

import (
	"errors"
	"testing"
	"time"

	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/octave"
)

func TestNew(t *testing.T) {
	tests := []Name{C, CSHARP, DFLAT, E, FSHARP, GFLAT, A, BSHARP, CFLAT2, DSHARP2}

	for _, name := range tests {
		t.Run(name.String(), func(t *testing.T) {
			note := New(name)
			if !note.Name().EqualSpelling(name) {
				t.Fatalf("New(%v).Name(): got %v, want %v", name, note.Name(), name)
			}
			if note.Octave() != nil {
				t.Fatalf("New(%v).Octave(): got %v, want nil", name, note.Octave())
			}
		})
	}
}

func TestNote_String(t *testing.T) {
	tests := []struct {
		note Note
		want string
	}{
		{New(C), "C"},
		{New(CSHARP), "C#"},
		{New(DFLAT), "Db"},
		{New(ESHARP2), "E##"},
		{New(FFLAT2), "Fbb"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.note.String(); got != tt.want {
				t.Fatalf("String(): got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewFromString(t *testing.T) {
	tests := []struct {
		in      string
		want    Name
		wantErr bool
	}{
		{"C", C, false},
		{"C#", CSHARP, false},
		{"Db", DFLAT, false},
		{"E##", ESHARP2, false},
		{"Fbb", FFLAT2, false},
		{"c", C, false},
		{"d#", DSHARP, false},
		{"", Name{}, true},
		{"H", Name{}, true},
		{"C#b", Name{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := NewFromString(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NewFromString(%q): expected error, got nil", tt.in)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewFromString(%q): unexpected error: %v", tt.in, err)
			}
			if !got.Name().EqualSpelling(tt.want) {
				t.Fatalf("NewFromString(%q): got %v, want %v", tt.in, got.Name(), tt.want)
			}
		})
	}
}

func TestMustNewFromString(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		tests := []struct {
			in   string
			want Name
		}{
			{"C", C},
			{"C#", CSHARP},
			{"Db", DFLAT},
		}

		for _, tt := range tests {
			got := MustNewFromString(tt.in)
			if !got.Name().EqualSpelling(tt.want) {
				t.Fatalf("MustNewFromString(%q): got %v, want %v", tt.in, got.Name(), tt.want)
			}
		}
	})

	t.Run("Invalid_Panics", func(t *testing.T) {
		invalidInputs := []string{"", "H", "C#b"}
		for _, in := range invalidInputs {
			func() {
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("MustNewFromString(%q): expected panic, got none", in)
					}
				}()
				_ = MustNewFromString(in)
			}()
		}
	})
}

func TestNewFromNoteNames(t *testing.T) {
	tests := []struct {
		name  string
		names []Name
		want  int
	}{
		{"Empty", []Name{}, 0},
		{"One", []Name{C}, 1},
		{"Three", []Name{C, D, E}, 3},
		{"Seven", []Name{C, D, E, F, G, A, B}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notes := NewFromNoteNames(tt.names...)
			if got := len(notes); got != tt.want {
				t.Fatalf("NewFromNoteNames(): got %d notes, want %d", got, tt.want)
			}
			for i, n := range notes {
				if !n.Name().EqualSpelling(tt.names[i]) {
					t.Fatalf("NewFromNoteNames()[%d]: got %v, want %v", i, n.Name(), tt.names[i])
				}
			}
		})
	}
}

func TestNewWithOctave(t *testing.T) {
	tests := []struct {
		name       string
		noteName   Name
		octaveNum  octave.Number
		wantOctave octave.Number
		wantErr    bool
	}{
		{"C4", C, 4, 4, false},
		{"A-1", A, -1, -1, false},
		{"G9", G, 9, 9, false},
		{"C0", C, 0, 0, false},
		{"InvalidOctave", C, 15, 0, true},
		{"InvalidOctaveNeg", C, -5, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithOctave(tt.noteName, tt.octaveNum)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NewWithOctave(%v, %d): expected error, got nil", tt.noteName, tt.octaveNum)
				}
				if !errors.Is(err, octave.ErrOctaveNumberUnknown) {
					t.Fatalf("NewWithOctave(): expected ErrOctaveNumberUnknown, got %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewWithOctave(%v, %d): unexpected error: %v", tt.noteName, tt.octaveNum, err)
			}
			if !got.Name().EqualSpelling(tt.noteName) {
				t.Fatalf("NewWithOctave().Name(): got %v, want %v", got.Name(), tt.noteName)
			}
			if got.Octave() == nil {
				t.Fatalf("NewWithOctave().Octave(): got nil, want %d", tt.wantOctave)
			}
			if got.Octave().Number() != tt.wantOctave {
				t.Fatalf("NewWithOctave().Octave().Number(): got %d, want %d", got.Octave().Number(), tt.wantOctave)
			}
		})
	}
}

func TestMustNewWithOctave(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		tests := []struct {
			name      Name
			octaveNum octave.Number
		}{
			{C, 4},
			{CSHARP, 0},
			{DFLAT, -1},
		}

		for _, tt := range tests {
			note := MustNewWithOctave(tt.name, tt.octaveNum)
			if !note.Name().EqualSpelling(tt.name) {
				t.Fatalf("MustNewWithOctave().Name(): got %v, want %v", note.Name(), tt.name)
			}
			if note.Octave().Number() != tt.octaveNum {
				t.Fatalf("MustNewWithOctave().Octave(): got %d, want %d", note.Octave().Number(), tt.octaveNum)
			}
		}
	})

	t.Run("InvalidOctave_Panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("MustNewWithOctave() with invalid octave: expected panic, got none")
			}
		}()
		_ = MustNewWithOctave(C, 15)
	})
}

func TestNote_EqualByName(t *testing.T) {
	tests := []struct {
		name  string
		a, b  Note
		equal bool
	}{
		{"SameNatural", New(C), New(C), true},
		{"SameSharp", New(CSHARP), New(CSHARP), true},
		{"DifferentLetter", New(C), New(D), false},
		{"DifferentAccidental", New(CSHARP), New(CFLAT), false},
		{"WithOctaveSameName", MustNewWithOctave(C, 4), MustNewWithOctave(C, 5), true},
		{"EnharmonicDifferent", New(CSHARP), New(DFLAT), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.EqualByName(tt.b); got != tt.equal {
				t.Fatalf("EqualByName(): got %v, want %v", got, tt.equal)
			}
		})
	}
}

func TestNote_EqualByOctave(t *testing.T) {
	tests := []struct {
		name  string
		a, b  Note
		equal bool
	}{
		{"BothNilOctave", New(C), New(D), true},
		{"SameOctave", MustNewWithOctave(C, 4), MustNewWithOctave(D, 4), true},
		{"DifferentOctave", MustNewWithOctave(C, 4), MustNewWithOctave(C, 5), false},
		{"OneNilOctave", New(C), MustNewWithOctave(C, 4), false},
		{"OtherNilOctave", MustNewWithOctave(C, 4), New(C), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.EqualByOctave(tt.b); got != tt.equal {
				t.Fatalf("EqualByOctave(): got %v, want %v", got, tt.equal)
			}
		})
	}
}

func TestNote_Equal(t *testing.T) {
	tests := []struct {
		name  string
		a, b  Note
		equal bool
	}{
		{"SameNameNilOctave", New(C), New(C), true},
		{"SameNameSameOctave", MustNewWithOctave(C, 4), MustNewWithOctave(C, 4), true},
		{"DifferentName", New(C), New(D), false},
		{"DifferentOctave", MustNewWithOctave(C, 4), MustNewWithOctave(C, 5), false},
		{"SameNameMixedOctave", New(C), MustNewWithOctave(C, 4), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.equal {
				t.Fatalf("Equal(): got %v, want %v", got, tt.equal)
			}
		})
	}
}

func TestNote_Copy(t *testing.T) {
	t.Run("BasicCopy", func(t *testing.T) {
		original := MustNewWithOctave(CSHARP, 4)
		copy := original.Copy()

		if !copy.Equal(original) {
			t.Fatalf("Copy() should be equal to original")
		}

		// Modify copy's octave shouldn't affect original
		if copy.Octave() == original.Octave() {
			t.Fatalf("Copy() should have different octave pointer")
		}
	})

	t.Run("CopyWithDuration", func(t *testing.T) {
		original := New(C).SetDuration(time.Second).SetValue(duration.NewRelative(duration.NameHalf))
		copy := original.Copy()

		if copy.Duration() != original.Duration() {
			t.Fatalf("Copy().Duration(): got %v, want %v", copy.Duration(), original.Duration())
		}
		if copy.Value() == nil {
			t.Fatalf("Copy().Value(): got nil")
		}
	})

	t.Run("CopyNilOctave", func(t *testing.T) {
		original := New(C)
		copy := original.Copy()

		if !copy.EqualByName(original) {
			t.Fatalf("Copy() with nil octave should preserve name")
		}
		if copy.Octave() != nil {
			t.Fatalf("Copy() with nil octave should have nil octave")
		}
	})
}

func TestNote_AlterUp(t *testing.T) {
	tests := []struct {
		in   Note
		want Name
	}{
		{New(C), CSHARP},
		{New(CSHARP), CSHARP2},
		{New(CFLAT), C},
		{New(CFLAT2), CFLAT},
		{New(B), BSHARP},
	}

	for _, tt := range tests {
		t.Run(tt.in.String()+"->"+tt.want.String(), func(t *testing.T) {
			got := tt.in.AlterUp()
			if !got.Name().EqualSpelling(tt.want) {
				t.Fatalf("AlterUp(): got %v, want %v", got.Name(), tt.want)
			}
		})
	}
}

func TestNote_AlterDown(t *testing.T) {
	tests := []struct {
		in   Note
		want Name
	}{
		{New(C), CFLAT},
		{New(CFLAT), CFLAT2},
		{New(CSHARP), C},
		{New(CSHARP2), CSHARP},
		{New(B), BFLAT},
	}

	for _, tt := range tests {
		t.Run(tt.in.String()+"->"+tt.want.String(), func(t *testing.T) {
			got := tt.in.AlterDown()
			if !got.Name().EqualSpelling(tt.want) {
				t.Fatalf("AlterDown(): got %v, want %v", got.Name(), tt.want)
			}
		})
	}
}

func TestNote_AlterUpBy(t *testing.T) {
	tests := []struct {
		name  string
		in    Note
		steps uint8
		want  Name
	}{
		{"CBy0", New(C), 0, C},
		{"CBy1", New(C), 1, CSHARP},
		{"CBy2", New(C), 2, CSHARP2},
		{"CbbBy4", New(CFLAT2), 4, CSHARP2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterUpBy(tt.steps)
			if !got.Name().EqualSpelling(tt.want) {
				t.Fatalf("AlterUpBy(%d): got %v, want %v", tt.steps, got.Name(), tt.want)
			}
		})
	}
}

func TestNote_AlterDownBy(t *testing.T) {
	tests := []struct {
		name  string
		in    Note
		steps uint8
		want  Name
	}{
		{"CBy0", New(C), 0, C},
		{"CBy1", New(C), 1, CFLAT},
		{"CBy2", New(C), 2, CFLAT2},
		{"C##By4", New(CSHARP2), 4, CFLAT2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AlterDownBy(tt.steps)
			if !got.Name().EqualSpelling(tt.want) {
				t.Fatalf("AlterDownBy(%d): got %v, want %v", tt.steps, got.Name(), tt.want)
			}
		})
	}
}

func TestNote_BaseName(t *testing.T) {
	tests := []struct {
		note Note
		want string
	}{
		{New(C), "C"},
		{New(CSHARP), "C"},
		{New(CFLAT), "C"},
		{New(CSHARP2), "C"},
		{New(CFLAT2), "C"},
		{New(D), "D"},
		{New(DSHARP), "D"},
	}

	for _, tt := range tests {
		t.Run(tt.note.String(), func(t *testing.T) {
			if got := tt.note.BaseName(); got != tt.want {
				t.Fatalf("BaseName(): got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNote_AlterationShift(t *testing.T) {
	tests := []struct {
		note  Note
		shift int8
	}{
		{New(C), 0},
		{New(CSHARP), 1},
		{New(CSHARP2), 2},
		{New(CFLAT), -1},
		{New(CFLAT2), -2},
	}

	for _, tt := range tests {
		t.Run(tt.note.String(), func(t *testing.T) {
			if got := tt.note.AlterationShift(); got != tt.shift {
				t.Fatalf("AlterationShift(): got %d, want %d", got, tt.shift)
			}
		})
	}
}

func TestNote_SetOctave(t *testing.T) {
	oct4 := octave.MustNewByNumber(4)
	oct5 := octave.MustNewByNumber(5)

	t.Run("SetToNil", func(t *testing.T) {
		note := New(C)
		note = note.SetOctave(oct4)
		if note.Octave() == nil {
			t.Fatalf("SetOctave(): octave should not be nil")
		}
		if note.Octave().Number() != 4 {
			t.Fatalf("SetOctave(): got %d, want 4", note.Octave().Number())
		}
	})

	t.Run("ReplaceExisting", func(t *testing.T) {
		note := MustNewWithOctave(C, 4)
		note = note.SetOctave(oct5)
		if note.Octave().Number() != 5 {
			t.Fatalf("SetOctave(): got %d, want 5", note.Octave().Number())
		}
	})

	t.Run("SetNil", func(t *testing.T) {
		note := MustNewWithOctave(C, 4)
		note = note.SetOctave(nil)
		if note.Octave() != nil {
			t.Fatalf("SetOctave(nil): octave should be nil")
		}
	})
}

func TestNote_SetDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
	}{
		{"Zero", 0},
		{"Second", time.Second},
		{"Millisecond", time.Millisecond},
		{"Minute", time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note := New(C).SetDuration(tt.duration)
			if got := note.Duration(); got != tt.duration {
				t.Fatalf("SetDuration(): got %v, want %v", got, tt.duration)
			}
		})
	}
}

func TestNote_Duration(t *testing.T) {
	t.Run("DefaultZero", func(t *testing.T) {
		note := New(C)
		if got := note.Duration(); got != 0 {
			t.Fatalf("Duration(): got %v, want 0", got)
		}
	})

	t.Run("AfterSet", func(t *testing.T) {
		note := New(C).SetDuration(time.Second)
		if got := note.Duration(); got != time.Second {
			t.Fatalf("Duration(): got %v, want %v", got, time.Second)
		}
	})
}

func TestNote_SetValue(t *testing.T) {
	values := []*duration.Relative{
		duration.NewRelative(duration.NameWhole),
		duration.NewRelative(duration.NameHalf),
		duration.NewRelative(duration.NameQuarter),
		duration.NewRelative(duration.NameEighth),
	}

	for _, v := range values {
		t.Run(string(v.Name()), func(t *testing.T) {
			note := New(C).SetValue(v)
			if note.Value() != v {
				t.Fatalf("SetValue(): got %v, want %v", note.Value(), v)
			}
		})
	}
}

func TestNote_Value(t *testing.T) {
	t.Run("DefaultNil", func(t *testing.T) {
		note := New(C)
		if got := note.Value(); got != nil {
			t.Fatalf("Value(): got %v, want nil", got)
		}
	})

	t.Run("AfterSet", func(t *testing.T) {
		val := duration.NewRelative(duration.NameHalf)
		note := New(C).SetValue(val)
		if got := note.Value(); got != val {
			t.Fatalf("Value(): got %v, want %v", got, val)
		}
	})
}

func TestNotes_String(t *testing.T) {
	tests := []struct {
		name  string
		notes Notes
		want  string
	}{
		{"Empty", Notes{}, "[]"},
		{"One", Notes{New(C)}, "[C]"},
		{"Three", Notes{New(C), New(D), New(E)}, "[C D E]"},
		{"WithAccidentals", Notes{New(CSHARP), New(DFLAT)}, "[C# Db]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.notes.String(); got != tt.want {
				t.Fatalf("String(): got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNotes_Uniques(t *testing.T) {
	tests := []struct {
		name    string
		notes   Notes
		wantLen int
	}{
		{"Empty", Notes{}, 0},
		{"NoDuplicates", Notes{New(C), New(D), New(E)}, 3},
		{"AllDuplicates", Notes{New(C), New(C), New(C)}, 1},
		{"SomeDuplicates", Notes{New(C), New(D), New(C), New(E), New(D)}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniques := tt.notes.Uniques()
			if got := len(uniques); got != tt.wantLen {
				t.Fatalf("Uniques(): got %d, want %d", got, tt.wantLen)
			}
		})
	}
}

func TestGetSetFullChromatic(t *testing.T) {
	notes := GetSetFullChromatic()

	// Should have 17 notes (C, Db, C#, D, Eb, D#, E, F, Gb, F#, G, Ab, G#, A, Bb, A#, B)
	if got := len(notes); got != 17 {
		t.Fatalf("GetSetFullChromatic(): got %d notes, want 17", got)
	}

	// All should be valid
	for _, n := range notes {
		if !n.Name().IsValid() {
			t.Fatalf("GetSetFullChromatic(): %v is not valid", n.Name())
		}
	}
}

func TestGetSetFullChromaticDoubleAltered(t *testing.T) {
	notes := GetSetFullChromaticDoubleAltered()

	// Should have 27 notes
	if got := len(notes); got != 27 {
		t.Fatalf("GetSetFullChromaticDoubleAltered(): got %d notes, want 27", got)
	}

	// All should be valid
	for _, n := range notes {
		if !n.Name().IsValid() {
			t.Fatalf("GetSetFullChromaticDoubleAltered(): %v is not valid", n.Name())
		}
	}

	// Should include double-altered notes
	hasDoubleSharp := false
	hasDoubleFlat := false
	for _, n := range notes {
		if n.AlterationShift() == 2 {
			hasDoubleSharp = true
		}
		if n.AlterationShift() == -2 {
			hasDoubleFlat = true
		}
	}
	if !hasDoubleSharp {
		t.Fatalf("GetSetFullChromaticDoubleAltered(): should include double-sharp notes")
	}
	if !hasDoubleFlat {
		t.Fatalf("GetSetFullChromaticDoubleAltered(): should include double-flat notes")
	}
}
