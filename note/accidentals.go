package note

// Accidental is a type for the accidental symbol. It represents a sharp, flat, or natural note.
type Accidental string

const (
	AccidentalFlat    = Accidental("b")
	AccidentalSharp   = Accidental("#")
	AccidentalNatural = Accidental("♮")
)

// String returns the string representation of the accidental.
func (as Accidental) String() string {
	return string(as)
}

// GetNotesWithAlterations returns the provided notes, each altered upward and downward by the specified number of semitones.
//
// Parameters:
//   - notes: The original set of notes to be altered.
//   - alterations: The number of semitone steps to alter each note.
//
// Returns:
//   - A slice of Note objects with the applied alterations.
func GetNotesWithAlterations(notes Notes, alterations uint8) (result Notes) {
	var i uint8
	for _, n := range notes {
		for i = alterations; i > 0; i-- {
			result = append(result, n.Copy().AlterUpBy(i), n.Copy().AlterDownBy(i))
		}
	}

	return
}
