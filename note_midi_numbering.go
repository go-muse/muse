package muse

import (
	"github.com/pkg/errors"
)

const (
	minMIDINumber = 0
	maxMIDINumber = 127
)

// MIDINumber returns note number coded by unsigned integer in range [0; 127] according to RFC 6295.
func (n *Note) MIDINumber() uint8 {
	if n == nil || n.octave == nil {
		return minMIDINumber
	}

	result := int16(n.getNoteNumberWithinOctave()) + (int16(n.octave.number)+1)*NotesInOctave

	if result < minMIDINumber {
		return minMIDINumber
	} else if result > maxMIDINumber {
		return maxMIDINumber
	}

	return uint8(result)
}

//nolint:gomnd
func (n *Note) getBaseNoteNumberWithinOctave() int8 {
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

func (n *Note) getNoteNumberWithinOctave() int8 {
	if n != nil {
		return n.getBaseNoteNumberWithinOctave() + n.GetAlterationShift()
	}

	return 0
}

// ErrMIDINumberUnknown appears when midi number is outside [0; 127].
var ErrMIDINumberUnknown = errors.New("unknown midi number")

// NewNoteFromMIDINumber creates note from midi number. Alterated notes will be sharpened not flatted. Note will contain proper octave.
func NewNoteFromMIDINumber(midiNumber uint8) (*Note, error) {
	if midiNumber > maxMIDINumber || midiNumber < minMIDINumber {
		return nil, errors.Wrapf(ErrMIDINumberUnknown, "midi number: %d. must me in [0; 127]", midiNumber)
	}

	noteNames := []NoteName{C, CSHARP, D, DSHARP, E, F, FSHARP, G, GSHARP, A, ASHARP, B}

	noteName := noteNames[midiNumber%12]
	octave := midiNumber/12 - 1 // Subtract 1 because MIDI note numbers start from C-1 (which is note number 12)

	return newNoteWithOctave(noteName, MustNewOctave(OctaveNumber(octave))), nil
}
