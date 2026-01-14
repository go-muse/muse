package note

import (
	"errors"
	"fmt"
)

// MilliCents represents one thousandth of a cent.
// 1200 cents (1_200_000 milli-cents) correspond to one octave.
//
// This unit is used to represent microtonal deviations without relying
// on floating-point arithmetic and without binding the model to a
// specific tuning system.
type MilliCents int32

var (
	// ErrInvalidAccidental is returned when a string or value
	// cannot be interpreted as a valid accidental representation.
	ErrInvalidAccidental = errors.New("invalid accidental")
)

// Accidental represents a musical accidental (alteration).
//
// It is intentionally split into two independent components:
//
//   - Chromatic: the diatonic/chromatic alteration expressed in
//     semitone units (e.g. ♭, ♯, 𝄫, 𝄪).
//
//   - Micro: an additional microtonal offset expressed in milli-cents,
//     allowing future support for microtonal notation and alternative
//     tuning systems.
//
// This design keeps the spelling (notation) of a note separate from
// the tuning system used to interpret its pitch.
type Accidental struct {
	// Chromatic represents the traditional diatonic accidental
	// measured in semitone units:
	//
	//   -2 = double flat (bb)
	//   -1 = flat (b)
	//    0 = natural
	//   +1 = sharp (#)
	//   +2 = double sharp (##)
	//
	// Values outside this range are permitted to allow extended or
	// theoretical spellings (e.g. triple sharps), even if they are
	// uncommon in standard notation.
	Chromatic int8

	// Micro represents an additional microtonal offset in milli-cents.
	//
	// For standard 12-tone equal temperament this value should be zero.
	// In future microtonal systems it may represent quarter-tones,
	// commas, or other fine pitch adjustments.
	Micro MilliCents
}

// Natural returns a natural accidental (no alteration).
func Natural() Accidental {
	return Accidental{Chromatic: 0, Micro: 0}
}

// Flat returns a flat accidental (♭).
func Flat() Accidental {
	return Accidental{Chromatic: -1, Micro: 0}
}

// DoubleFlat returns a double-flat accidental (𝄫).
func DoubleFlat() Accidental {
	return Accidental{Chromatic: -2, Micro: 0}
}

// Sharp returns a sharp accidental (♯).
func Sharp() Accidental {
	return Accidental{Chromatic: 1, Micro: 0}
}

// DoubleSharp returns a double-sharp accidental (𝄪).
func DoubleSharp() Accidental {
	return Accidental{Chromatic: 2, Micro: 0}
}

// AlterUp returns a new Accidental with the chromatic alteration
// increased by one step.
//
// Semantically, this corresponds to adding one sharp (#)
// or removing one flat (b), while keeping the same note letter.
func (a Accidental) AlterUp() Accidental {
	a.Chromatic++
	return a
}

// AlterDown returns a new Accidental with the chromatic alteration
// decreased by one step.
//
// Semantically, this corresponds to adding one flat (b)
// or removing one sharp (#), while keeping the same note letter.
func (a Accidental) AlterDown() Accidental {
	a.Chromatic--
	return a
}

// AlterUpBy returns a new Accidental with the chromatic alteration
// increased by the given number of steps.
//
// Each step corresponds to one additional sharp (#).
// The Micro component is not modified.
func (a Accidental) AlterUpBy(steps uint8) Accidental {
	a.Chromatic += int8(steps)
	return a
}

// AlterDownBy returns a new Accidental with the chromatic alteration
// decreased by the given number of steps.
//
// Each step corresponds to one additional flat (b).
// The Micro component is not modified.
func (a Accidental) AlterDownBy(steps uint8) Accidental {
	a.Chromatic -= int8(steps)
	return a
}

// IsNatural reports whether the accidental represents an unaltered note,
// i.e. no chromatic alteration and no microtonal offset.
func (a Accidental) IsNatural() bool {
	return a.Chromatic == 0 && a.Micro == 0
}

// Equal reports whether two accidentals are exactly equal in spelling,
// including both chromatic and microtonal components.
func (a Accidental) Equal(other Accidental) bool {
	return a.Chromatic == other.Chromatic &&
		a.Micro == other.Micro
}

// String returns the ASCII string representation of the accidental.
//
// The chromatic component is rendered using:
//   - "bb", "b", "", "#", "##", etc.
//
// The microtonal component is intentionally not rendered yet.
// Future versions may add configurable rendering (arrows, cents, Unicode).
func (a Accidental) String() string {
	switch a.Chromatic {
	case -2:
		return "bb"
	case -1:
		return "b"
	case 0:
		return ""
	case 1:
		return "#"
	case 2:
		return "##"
	default:
		// Allow arbitrary numbers of sharps or flats for completeness.
		if a.Chromatic > 0 {
			out := ""
			for i := int8(0); i < a.Chromatic; i++ {
				out += "#"
			}
			return out
		}

		out := ""
		for i := int8(0); i < -a.Chromatic; i++ {
			out += "b"
		}
		return out
	}
}

// ParseAccidental parses an ASCII accidental representation.
//
// Supported spellings:
//
//	""    -> natural
//	"b"   -> flat
//	"bb"  -> double flat
//	"bbb" -> triple flat (and so on)
//	"#"   -> sharp
//	"##"  -> double sharp
//	"###" -> triple sharp (and so on)
//
// Unicode symbols (♭, ♯, 𝄫, 𝄪) are intentionally not supported here
// to keep the core parser simple and unambiguous.
//
// Returns ErrInvalidAccidental if the input cannot be parsed.
func ParseAccidental(s string) (Accidental, error) {
	if s == "" {
		return Natural(), nil
	}

	// Check if all characters are sharps
	if allSameChar(s, '#') {
		return Accidental{Chromatic: int8(len(s))}, nil
	}

	// Check if all characters are flats
	if allSameChar(s, 'b') {
		return Accidental{Chromatic: -int8(len(s))}, nil
	}

	return Accidental{}, fmt.Errorf("%w: %q", ErrInvalidAccidental, s)
}

// allSameChar reports whether all characters in s are equal to c.
func allSameChar(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != c {
			return false
		}
	}
	return true
}

// MustParseAccidental parses an accidental string and panics on error.
//
// This function is intended for use in tests, constants, and internal
// tables where the input is guaranteed to be valid.
func MustParseAccidental(s string) Accidental {
	a, err := ParseAccidental(s)
	if err != nil {
		panic(err)
	}

	return a
}
