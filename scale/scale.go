package scale

import (
	"fmt"

	"github.com/go-muse/muse/note"
)

// Scale is a set of notes.
type Scale note.Notes

// NewScaleFromNotes creates a scale from a given notes.
func NewScaleFromNotes(notes ...note.Note) Scale {
	scale := make(Scale, 0, len(notes))
	scale = append(scale, notes...)

	return scale
}

// NewScaleFromNoteNames creates notes from the given names and then creates a scale from them.
func NewScaleFromNoteNames(noteNames ...note.Name) (Scale, error) {
	scale := make(Scale, 0, len(noteNames))
	for _, noteName := range noteNames {
		scale = append(scale, note.New(noteName))
	}

	return scale, nil
}

// MustNewScaleFromNoteNames creates notes from the given names and then creates a scale from them. Panics in case of errors.
func MustNewScaleFromNoteNames(noteNames ...note.Name) Scale {
	scale := make(Scale, 0, len(noteNames))
	for _, noteName := range noteNames {
		scale = append(scale, noteName.MustMakeNote())
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
		note.C.NewNote(),
		note.DFLAT.NewNote(),
		note.CSHARP.NewNote(),
		note.D.NewNote(),
		note.EFLAT.NewNote(),
		note.DSHARP.NewNote(),
		note.E.NewNote(),
		note.F.NewNote(),
		note.GFLAT.NewNote(),
		note.FSHARP.NewNote(),
		note.G.NewNote(),
		note.AFLAT.NewNote(),
		note.GSHARP.NewNote(),
		note.A.NewNote(),
		note.BFLAT.NewNote(),
		note.ASHARP.NewNote(),
		note.B.NewNote(),
	}
}
