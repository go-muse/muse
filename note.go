package muse

type Note struct {
	name NoteName
}

// Name returns name of the note.
func (n *Note) Name() NoteName {
	return n.name
}
