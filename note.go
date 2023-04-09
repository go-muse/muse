package muse

import (
	"github.com/pkg/errors"
)

type Note struct {
	name NoteName
}

// newNote creates new note with a given name without any restrictions.
func newNote(name NoteName) *Note {
	return &Note{name: name}
}

// ErrUnknownNoteName is еру error that occurs when trying to determine the name of a note if it is not known.
var ErrUnknownNoteName = errors.New("unknown note name")

// NewNote creates new note with a given name validating it.
func NewNote(noteName NoteName) (*Note, error) {
	switch noteName {
	case C:
		return newNote(C), nil
	case CSHARP:
		return newNote(CSHARP), nil
	case DFLAT:
		return newNote(DFLAT), nil
	case D:
		return newNote(D), nil
	case DSHARP:
		return newNote(DSHARP), nil
	case EFLAT:
		return newNote(EFLAT), nil
	case E:
		return newNote(E), nil
	case F:
		return newNote(F), nil
	case FSHARP:
		return newNote(FSHARP), nil
	case GFLAT:
		return newNote(GFLAT), nil
	case G:
		return newNote(G), nil
	case GSHARP:
		return newNote(GSHARP), nil
	case AFLAT:
		return newNote(AFLAT), nil
	case A:
		return newNote(A), nil
	case ASHARP:
		return newNote(ASHARP), nil
	case BFLAT:
		return newNote(BFLAT), nil
	case B:
		return newNote(B), nil
	}

	return nil, errors.Wrapf(ErrUnknownNoteName, "given name: %s", noteName)
}

// MustNewNote creates new note with suppressing error in case of invalid note name.
func MustNewNote(noteName NoteName) *Note {
	note, err := NewNote(noteName)
	if err != nil {
		panic(err)
	}

	return note
}

// Name returns name of the note.
func (n *Note) Name() NoteName {
	return n.name
}

// IsEqualByName compares notes by name.
func (n *Note) IsEqualByName(note *Note) bool {
	if n == nil || note == nil {
		return false
	}

	return note.Name() == n.Name()
}

// Copy creates full copy of the current Note.
// The method returns a pointer to the new Note containing the same attribute values
// as the original Note that the function was called on.
func (n *Note) Copy() *Note {
	if n == nil {
		return nil
	}

	return &Note{name: n.Name()}
}
