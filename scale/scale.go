package scale

import (
	"fmt"

	"github.com/go-muse/muse/note"
)

// Scale is a set of notes.
type Scale note.Notes

// NewScaleFromNotes creates a scale from a given notes.
func NewScaleFromNotes(notes ...*note.Note) Scale {
	scale := make(Scale, 0, len(notes))
	scale = append(scale, notes...)

	return scale
}

// NewScaleFromNoteNames creates notes from the given names and then creates a scale from them.
func NewScaleFromNoteNames(noteNames ...note.Name) (Scale, error) {
	scale := make(Scale, 0, len(noteNames))
	for _, noteName := range noteNames {
		n, err := note.New(noteName)
		if err != nil {
			return nil, fmt.Errorf("create note by name '%s': %w", noteName, err)
		}

		scale = append(scale, n)
	}

	return scale, nil
}

// MustNewScaleFromNoteNames creates notes from the given names and then creates a scale from them. Panics in case of errors.
func MustNewScaleFromNoteNames(noteNames ...note.Name) Scale {
	scale := make(Scale, 0, len(noteNames))
	for _, noteName := range noteNames {
		scale = append(scale, note.MustNewNote(noteName))
	}

	return scale
}

// String is stringer for Scale.
func (s Scale) String() string {
	noteNames := make([]note.Name, len(s))
	for i, note := range s {
		noteNames[i] = note.Name()
	}

	return fmt.Sprintf("%v", noteNames)
}

// GetFullChromaticScale returns all notes of the tonal system.
func GetFullChromaticScale() Scale {
	return Scale{
		note.MustNewNote(note.C),
		note.MustNewNote(note.DFLAT),
		note.MustNewNote(note.CSHARP),
		note.MustNewNote(note.D),
		note.MustNewNote(note.EFLAT),
		note.MustNewNote(note.DSHARP),
		note.MustNewNote(note.E),
		note.MustNewNote(note.F),
		note.MustNewNote(note.GFLAT),
		note.MustNewNote(note.FSHARP),
		note.MustNewNote(note.G),
		note.MustNewNote(note.AFLAT),
		note.MustNewNote(note.GSHARP),
		note.MustNewNote(note.A),
		note.MustNewNote(note.BFLAT),
		note.MustNewNote(note.ASHARP),
		note.MustNewNote(note.B),
	}
}
