package muse

// NoteName is a common note for the note.
type NoteName string

// String is stringer for NoteName type.
func (nn NoteName) String() string {
	return string(nn)
}

// MustMakeNote makes note with the current note name with panic on validation.
func (nn NoteName) MustMakeNote() *Note {
	if err := nn.Validate(); err != nil {
		panic(err)
	}

	return newNote(nn)
}

// MakeNote makes note with the current note name.
func (nn NoteName) MakeNote() (*Note, error) {
	if err := nn.Validate(); err != nil {
		return nil, err
	}

	return newNote(nn), nil
}

// Validate checks note name.
func (nn NoteName) Validate() error {
	if len(nn) < 1 {
		return ErrNoteNameUnknown
	}

	baseNoteNames := map[NoteName]struct{}{
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
				(firstAlterationSymbol != AlterSymbolFlat.String() && firstAlterationSymbol != AlterSymbolSharp.String()) {
				return ErrNoteNameUnknown
			}
		}
	}

	return nil
}
