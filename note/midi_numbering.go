package note

import (
	"errors"
	"fmt"

	"github.com/go-muse/muse/common/convert"
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/octave"
)

const (
	minMIDINumber = uint8(0)
	maxMIDINumber = uint8(127)
)

// MIDINumber returns note number coded by unsigned integer in range [0; 127] according to RFC 6295.
func (n *Note) MIDINumber() uint8 {
	if n == nil || n.octave == nil {
		return minMIDINumber
	}
	result := n.mustGetNoteNumberWithinOctave() + (convert.AddUint8Int8(1, int8(n.octave.Number())))*octave.NotesInOctave

	if result < minMIDINumber {
		return minMIDINumber
	} else if result > maxMIDINumber {
		return maxMIDINumber
	}

	return result
}

// getBaseNoteNumberWithinOctave returns base note number in range [0; 11] according to the note's name.
//
//nolint:mnd
func (n *Note) getBaseNoteNumberWithinOctave() uint8 {
	switch n.BaseName() {
	case C:
		return 0
	case D:
		return 2
	case E:
		return 4
	case F:
		return 5
	case G:
		return 7
	case A:
		return 9
	case B:
		return 11
	}

	return 0
}

// mustGetNoteNumberWithinOctave returns note number in range [0; 11]
// according to the note's base name and alteration shift.
func (n *Note) mustGetNoteNumberWithinOctave() uint8 {
	if n != nil {
		diff := int16(n.getBaseNoteNumberWithinOctave()) + int16(n.GetAlterationShift())
		// TODO: refactor this
		if diff > 0 && diff < int16(octave.NotesInOctave) {
			return convert.AddUint8Int8(n.getBaseNoteNumberWithinOctave(), n.GetAlterationShift())
		}

		if diff > int16(octave.NotesInOctave) {
			return octave.NotesInOctave
		}
	}

	return 0
}

// ErrMIDINumberUnknown appears when midi number is outside [0; 127].
var ErrMIDINumberUnknown = errors.New("unknown midi number")

// NewNoteFromMIDINumber creates note from midi number. Altered notes will be sharpened not flatted.
// Note will contain proper octave.
func NewNoteFromMIDINumber(midiNumber uint8) (*Note, error) {
	if midiNumber > maxMIDINumber || midiNumber < minMIDINumber {
		return nil, fmt.Errorf("midi number: '%d'. must me in [0; 127]: %w", midiNumber, ErrMIDINumberUnknown)
	}

	noteNames := []Name{C, CSHARP, D, DSHARP, E, F, FSHARP, G, GSHARP, A, ASHARP, B}

	noteName := noteNames[midiNumber%uint8(halftone.HalfTonesInOctave)]
	octaveNumber := convert.SubUint8Uint8(midiNumber/uint8(halftone.HalfTonesInOctave), 1)

	oct, err := octave.NewByNumber(octave.Number(octaveNumber))
	if err != nil {
		return nil, fmt.Errorf("create octave with octave number '%d': %w", octaveNumber, err)
	}

	return newNoteWithOctave(noteName, oct), nil
}
