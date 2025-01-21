package note

import (
	"errors"
)

// Name is a common name for the note.
type Name string

// Names is a common name for set of the names.
type Names []Name

// Length returns amount of names in the Names.
func (ns Names) Length() uint64 {
	return uint64(len(ns))
}

// MustNewNote creates a new note from the note name.
// Creating a note by its name, taken from the Muse library, guarantees panics-free execution of this method.
func (nn Name) MustNewNote() *Note {
	return MustNewNote(nn)
}

// ErrNoteNameUnknown is the error that occurs when trying to determine the name of a note if it is not known.
var ErrNoteNameUnknown = errors.New("unknown note name")

// String is stringer for Name type.
func (nn Name) String() string {
	return string(nn)
}

// NewNote makes note with the current note name.
func (nn Name) NewNote() (*Note, error) {
	if err := nn.Validate(); err != nil {
		return nil, err
	}

	return newNote(nn), nil
}

// MustMakeNote makes note with the current note name with panic on validation.
func (nn Name) MustMakeNote() *Note {
	if err := nn.Validate(); err != nil {
		panic(err)
	}

	return newNote(nn)
}

// Validate checks note name.
func (nn Name) Validate() error {
	if len(nn) < 1 {
		return ErrNoteNameUnknown
	}

	baseNoteNames := map[Name]struct{}{
		C: {},
		D: {},
		E: {},
		F: {},
		G: {},
		A: {},
		B: {},
	}

	baseNoteName := nn[0:1]
	if _, ok := baseNoteNames[baseNoteName]; !ok {
		return ErrNoteNameUnknown
	}

	if len(nn) > 1 {
		firstAlterationSymbol := string(nn[1:2])
		for _, alterationSymbol := range nn[1:] {
			if string(alterationSymbol) != firstAlterationSymbol ||
				(firstAlterationSymbol != AccidentalFlat.String() && firstAlterationSymbol != AccidentalSharp.String()) {
				return ErrNoteNameUnknown
			}
		}
	}

	return nil
}
