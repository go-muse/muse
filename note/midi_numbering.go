package note

import (
	"errors"
	"fmt"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/octave"
)

const (
	minMIDINumber = uint8(0)
	maxMIDINumber = uint8(127)
)

// chromaticNames contains the 12 chromatic note names starting from C.
// Used for converting MIDI numbers to note names.
var chromaticNames = []Name{C, CSHARP, D, DSHARP, E, F, FSHARP, G, GSHARP, A, ASHARP, B}

// MIDINumber returns note number coded by unsigned integer in range [0; 127] according to RFC 6295.
func (n Note) MIDINumber() uint8 {
	if n.octave == nil {
		return minMIDINumber
	}

	// MIDI octave -1 starts at 0, so we add 1 to convert from musical octave numbering
	// Then multiply by 12 (notes per octave) and add the note position
	result := int(n.mustGetNoteNumberWithinOctave()) + (int(n.octave.Number())+1)*int(octave.NotesInOctave)

	if result < int(minMIDINumber) {
		return minMIDINumber
	}
	if result > int(maxMIDINumber) {
		return maxMIDINumber
	}

	return uint8(result)
}

// mustGetNoteNumberWithinOctave returns note number in range [0; 11]
// according to the note's base name and alteration shift.
// Uses modular arithmetic to handle edge cases like Cb (returns 11) and B# (returns 0).
func (n Note) mustGetNoteNumberWithinOctave() uint8 {
	semitone := int(n.name.Letter().Semitone()) + int(n.AlterationShift())
	// Wrap around using modular arithmetic: ((x % 12) + 12) % 12
	// This handles both negative values (Cb -> 11) and overflow (B# -> 0)
	semitone = ((semitone % 12) + 12) % 12
	return uint8(semitone)
}

// ErrMIDINumberUnknown appears when midi number is outside [0; 127].
var ErrMIDINumberUnknown = errors.New("unknown midi number")

// NewNoteFromMIDINumber creates note from midi number. Altered notes will be sharpened not flatted.
// Note will contain proper octave.
func NewNoteFromMIDINumber(midiNumber uint8) (Note, error) {
	if midiNumber > maxMIDINumber {
		return Note{}, fmt.Errorf("midi number: '%d'. must be in [0; 127]: %w", midiNumber, ErrMIDINumberUnknown)
	}

	name := chromaticNames[midiNumber%uint8(halftone.HalfTonesInOctave)]
	octaveNumber := int8(midiNumber/uint8(halftone.HalfTonesInOctave)) - 1

	oct, err := octave.NewByNumber(octave.Number(octaveNumber))
	if err != nil {
		return Note{}, fmt.Errorf("create octave with octave number '%d': %w", octaveNumber, err)
	}

	return New(name).SetOctave(oct), nil
}
