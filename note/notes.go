package note

import "fmt"

// Notes is slice of Note.
type Notes []Note

// String is stringer for Note.
func (ns Notes) String() string {
	noteNames := make([]Name, len(ns))
	for i, note := range ns {
		noteNames[i] = note.Name()
	}

	return fmt.Sprintf("%v", noteNames)
}

// Uniques returns a slice of unique notes from the current.
func (ns Notes) Uniques() Notes {
	duplicates := make(map[Name]bool)
	uniqueNotes := make(Notes, 0)

	for _, note := range ns {
		if !duplicates[note.Name()] {
			duplicates[note.Name()] = true
			uniqueNotes = append(uniqueNotes, note)
		}
	}

	return uniqueNotes
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
