package degree

import "github.com/go-muse/muse/note"

// Iterator is an object that allows iterating through a sequence of degrees
// and also provides additional functionality.
type Iterator <-chan *Degree

// GetAllDegrees iterates through a sequence of degrees
// and returns them as slice.
func (di Iterator) GetAllDegrees() []*Degree {
	if di == nil {
		return nil
	}

	var degrees []*Degree
	for {
		degree, ok := <-di
		if !ok {
			return degrees
		}
		degrees = append(degrees, degree)
	}
}

// GetAllNotes iterates through a sequence of degrees
// and returns their notes as slice.
func (di Iterator) GetAllNotes() note.Notes {
	if di == nil {
		return nil
	}

	notes := make(note.Notes, 0)
	for {
		degree, ok := <-di
		if !ok {
			return notes
		}
		if degree.Note() != nil {
			notes = append(notes, degree.Note())
		}
	}
}
