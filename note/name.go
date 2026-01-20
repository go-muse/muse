package note

import (
	"errors"
	"fmt"
)

// Name represents a spelled note name (letter + accidental),
// independent of tuning or octave.
type Name struct {
	letter     Letter
	accidental Accidental
}

// ErrNameInvalid is returned when a note name is syntactically invalid
// or cannot be parsed into a valid Letter + Accidental combination.
var ErrNameInvalid = errors.New("invalid note name")

// NewName constructs a Name from a diatonic letter and an accidental.
func NewName(letter Letter, accidental Accidental) Name {
	// accidental currently has no IsValid; we accept any int8 Chromatic and any Micro.
	return Name{letter: letter, accidental: accidental}
}

// String returns the ASCII spelling of the note name, e.g. "C", "Db", "F##".
func (n Name) String() string {
	return n.letter.String() + n.accidental.String()
}

// Copy returns a copy of the note name.
func (n Name) Copy() Name {
	return n
}

// BaseName returns the base diatonic letter name (C, D, E, F, G, A, B).
func (n Name) BaseName() string {
	return n.letter.String()
}

// EqualSpelling reports whether two note names have exactly the same spelling.
func (n Name) EqualSpelling(other Name) bool {
	return n.letter == other.letter && n.accidental.Equal(other.accidental)
}

// NewNameFromString parses an ASCII note name like "C", "Db", "F##".
// Supported accidentals: "", "b", "bb", "#", "##".
//
// Microtonal notation is intentionally not supported here yet.
func NewNameFromString(s string) (Name, error) {
	if len(s) < 1 {
		return Name{}, fmt.Errorf("%w: %q", ErrNameInvalid, s)
	}

	letter, err := ParseLetter(s[:1])
	if err != nil {
		return Name{}, fmt.Errorf("%w: %w", ErrNameInvalid, err)
	}

	accidental, err := ParseAccidental(s[1:])
	if err != nil {
		return Name{}, fmt.Errorf("%w: %w", ErrNameInvalid, err)
	}

	return NewName(letter, accidental), nil
}

// MustNewNameFromString parses a note name and panics on error.
// Intended for tests and internal tables.
func MustNewNameFromString(s string) Name {
	n, err := NewNameFromString(s)
	if err != nil {
		panic(err)
	}

	return n
}

// NewNote makes note with the current note name.
func (n Name) NewNote() Note {
	return New(n)
}

// IsValid reports whether Name satisfies basic invariants.
// Currently, this means the diatonic letter is valid.
func (n Name) IsValid() bool {
	return n.letter.IsValid()
}

// AlterUp returns a new Name with its accidental altered up by one step,
// while keeping the same diatonic letter.
//
// Example: C  -> C#
//
//	Eb -> E
//	F# -> F##
func (n Name) AlterUp() Name {
	n.accidental = n.accidental.AlterUp()
	return n
}

// AlterDown returns a new Name with its accidental altered down by one step,
// while keeping the same diatonic letter.
//
// Example: C# -> C
//
//	E  -> Eb
//	F  -> Fb
func (n Name) AlterDown() Name {
	n.accidental = n.accidental.AlterDown()
	return n
}

// AlterUpBy returns a new Name with its accidental altered up
// by the given number of steps, while keeping the same diatonic letter.
//
// Each step corresponds to one sharp (#).
func (n Name) AlterUpBy(steps uint8) Name {
	n.accidental = n.accidental.AlterUpBy(steps)
	return n
}

// AlterDownBy returns a new Name with its accidental altered down
// by the given number of steps, while keeping the same diatonic letter.
//
// Each step corresponds to one flat (b).
func (n Name) AlterDownBy(steps uint8) Name {
	n.accidental = n.accidental.AlterDownBy(steps)
	return n
}

// AlterationShift returns the chromatic alteration (sharps/flats count).
func (n Name) AlterationShift() int8 {
	return n.accidental.Chromatic
}

// Letter returns the diatonic letter of this note name.
func (n Name) Letter() Letter {
	return n.letter
}

// Accidental returns the accidental of this note name.
func (n Name) Accidental() Accidental {
	return n.accidental
}

// IsBaseNameHigherThan reports whether this note name's diatonic letter
// represents a higher position than the other note name's letter within
// a single octave, using the natural (unaltered) semitone positions:
// C=0, D=2, E=4, F=5, G=7, A=9, B=11.
//
// This comparison ignores accidentals and compares only the base letter names.
// For pitch-based comparison including accidentals, use note.Note methods instead.
//
// Example:
//
//	C.IsBaseNameHigherThan(B)      // false (C=0 < B=11)
//	D.IsBaseNameHigherThan(C)      // true (D=2 > C=0)
//	CSHARP.IsBaseNameHigherThan(C) // false (both are C letter)
//	DFLAT.IsBaseNameHigherThan(C)  // true (D letter > C letter)
//	E.IsBaseNameHigherThan(E)      // false (equal)
func (n Name) IsBaseNameHigherThan(other Name) bool {
	return n.letter.IsHigherThan(other.letter)
}

// IsBaseNameLowerThan reports whether this note name's diatonic letter
// represents a lower position than the other note name's letter within
// a single octave, using the natural (unaltered) semitone positions:
// C=0, D=2, E=4, F=5, G=7, A=9, B=11.
//
// This comparison ignores accidentals and compares only the base letter names.
// For pitch-based comparison including accidentals, use note.Note methods instead.
//
// Example:
//
//	B.IsBaseNameLowerThan(C)      // false (B=11 > C=0)
//	C.IsBaseNameLowerThan(D)      // true (C=0 < D=2)
//	C.IsBaseNameLowerThan(CSHARP) // false (both are C letter)
//	C.IsBaseNameLowerThan(DFLAT)  // true (C letter < D letter)
//	E.IsBaseNameLowerThan(E)      // false (equal)
func (n Name) IsBaseNameLowerThan(other Name) bool {
	return n.letter.IsLowerThan(other.letter)
}
