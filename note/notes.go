package note

import "fmt"

// Notes is slice of Note.
type Notes []*Note

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
