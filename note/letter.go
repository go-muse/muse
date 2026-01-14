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

// ErrInvalidLetter indicates that a value does not correspond to any valid
// diatonic note letter in the set {C, D, E, F, G, A, B}.
var ErrInvalidLetter = errors.New("invalid diatonic note letter")

// ParseLetter parses a string and returns the corresponding Letter.
// The input string must be exactly one character long and must be one of: C, D, E, F, G, A, or B.
// Returns an error if the input is invalid.
func ParseLetter(s string) (Letter, error) {
	if len(s) != 1 {
		return 0, ErrInvalidLetter
	}

	switch s[0] {
	case 'C':
		return LetterC, nil
	case 'D':
		return LetterD, nil
	case 'E':
		return LetterE, nil
	case 'F':
		return LetterF, nil
	case 'G':
		return LetterG, nil
	case 'A':
		return LetterA, nil
	case 'B':
		return LetterB, nil
	default:
		return 0, ErrInvalidLetter
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
