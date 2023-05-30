package muse

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
		return 3
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
	return n.getBaseNoteNumberWithinOctave() + n.GetAlterationShift()
}
