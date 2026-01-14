package note

import (
	"errors"
)

// Letter represents a diatonic note letter (C, D, E, F, G, A, B).
// It is the base spelling component of a note, independent of tuning.
type Letter uint8

const (
	LetterC Letter = iota
	LetterD
	LetterE
	LetterF
	LetterG
	LetterA
	LetterB
)

// String returns the string representation of the Letter.
// Returns "?" for invalid letter values.
func (l Letter) String() string {
	switch l {
	case LetterC:
		return "C"
	case LetterD:
		return "D"
	case LetterE:
		return "E"
	case LetterF:
		return "F"
	case LetterG:
		return "G"
	case LetterA:
		return "A"
	case LetterB:
		return "B"
	default:
		return "?"
	}
}

func (l Letter) IsValid() bool {
	return l <= LetterB
}

// Semitone returns the semitone position of this letter within an octave.
// C=0, D=2, E=4, F=5, G=7, A=9, B=11.
// This is the chromatic pitch class of the natural (unaltered) note.
func (l Letter) Semitone() uint8 {
	switch l {
	case LetterC:
		return 0
	case LetterD:
		return 2
	case LetterE:
		return 4
	case LetterF:
		return 5
	case LetterG:
		return 7
	case LetterA:
		return 9
	case LetterB:
		return 11
	default:
		return 0
	}
}

// ErrLetterInvalid indicates that a value does not correspond to any valid
// diatonic note letter in the set {C, D, E, F, G, A, B}.
var ErrLetterInvalid = errors.New("invalid diatonic note letter")

// ParseLetter parses a string and returns the corresponding Letter.
// The input string must be exactly one character long and must be one of: C, D, E, F, G, A, or B.
// Both uppercase and lowercase letters are accepted.
// Returns an error if the input is invalid.
func ParseLetter(s string) (Letter, error) {
	if len(s) != 1 {
		return 0, ErrLetterInvalid
	}

	switch s[0] {
	case 'C', 'c':
		return LetterC, nil
	case 'D', 'd':
		return LetterD, nil
	case 'E', 'e':
		return LetterE, nil
	case 'F', 'f':
		return LetterF, nil
	case 'G', 'g':
		return LetterG, nil
	case 'A', 'a':
		return LetterA, nil
	case 'B', 'b':
		return LetterB, nil
	default:
		return 0, ErrLetterInvalid
	}
}

// MustParseLetter parses a string and returns the corresponding Letter.
// Panics if the input string is invalid.
// Use this function when you are certain the input is valid.
func MustParseLetter(s string) Letter {
	l, err := ParseLetter(s)
	if err != nil {
		panic(err)
	}

	return l
}

// Name creation methods for Letter.
// These methods allow creating note names directly from a letter:
//
//	note.LetterC.Natural().NewNote()  // C
//	note.LetterC.Sharp().NewNote()    // C#
//	note.LetterD.Flat().NewNote()     // Db

// Natural returns a Name with this letter and no accidental.
func (l Letter) Natural() Name {
	return NewName(l, Natural())
}

// Sharp returns a Name with this letter and a sharp accidental.
func (l Letter) Sharp() Name {
	return NewName(l, Sharp())
}

// Flat returns a Name with this letter and a flat accidental.
func (l Letter) Flat() Name {
	return NewName(l, Flat())
}

// DoubleSharp returns a Name with this letter and a double-sharp accidental.
func (l Letter) DoubleSharp() Name {
	return NewName(l, DoubleSharp())
}

// DoubleFlat returns a Name with this letter and a double-flat accidental.
func (l Letter) DoubleFlat() Name {
	return NewName(l, DoubleFlat())
}
