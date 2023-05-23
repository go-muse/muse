package muse

import "fmt"

// Scale is a set of notes.
type Scale []Note

// NewScaleFromNotes creates a scale from a given notes.
func NewScaleFromNotes(notes ...Note) Scale {
	scale := make(Scale, 0, len(notes))
	scale = append(scale, notes...)

	return scale
}

// NewScaleFromNoteNames creates notes from the given names and then creates a scale from them.
func NewScaleFromNoteNames(noteNames ...NoteName) Scale {
	scale := make(Scale, 0, len(noteNames))
	for _, noteName := range noteNames {
		scale = append(scale, *newNote(noteName))
	}

	return scale
}

// GenerateScale generates an ascending or descending scale.
func (m *Mode) GenerateScale(desc bool) Scale {
	if m == nil || m.degree == nil || m.degree.note == nil {
		return nil
	}

	scale := make(Scale, m.Length())

	var fromDegree *Degree
	if desc {
		fromDegree = m.GetLastDegree()
	} else {
		fromDegree = m.GetFirstDegree()
	}

	var i int
	for degree := range fromDegree.IterateOneRound(desc) {
		scale[i] = *degree.Note()
		i++
	}

	return scale
}

func (s Scale) String() string {
	noteNames := make([]NoteName, len(s))
	for i, note := range s {
		noteNames[i] = note.Name()
	}

	return fmt.Sprintf("%v", noteNames)
}

// GetFullChromaticScale returns all notes of the tonal system.
func GetFullChromaticScale() Scale {
	return Scale{
		*newNote(C),
		*newNote(DFLAT),
		*newNote(CSHARP),
		*newNote(D),
		*newNote(EFLAT),
		*newNote(DSHARP),
		*newNote(E),
		*newNote(F),
		*newNote(GFLAT),
		*newNote(FSHARP),
		*newNote(G),
		*newNote(AFLAT),
		*newNote(GSHARP),
		*newNote(A),
		*newNote(BFLAT),
		*newNote(ASHARP),
		*newNote(B),
	}
}

func GetAllPossibleNotes(alterations uint8) Scale {
	scale := NewScaleFromNoteNames(C, D, E, F, G, A, B)
	var i uint8
	for _, note := range scale {
		for i = alterations; i > 0; i-- {
			scale = append(scale, *note.Copy().AlterUpBy(i), *note.Copy().AlterDownBy(i))
		}
	}

	return scale
}
