package degree

import (
	"github.com/go-muse/muse/note"
)

// Iterator is a function that yields degree nodes one by one.
// It can be used with range loops (Go 1.23+).
type Iterator func(yield func(*Node) bool)

// GetAllDegrees collects all degree nodes from an iterator into a slice.
func (it Iterator) GetAllDegrees() []*Node {
	if it == nil {
		return nil
	}

	// Pre-allocate with typical capacity for a diatonic scale (7 degrees)
	nodes := make([]*Node, 0, 7)
	for node := range it {
		nodes = append(nodes, node)
	}

	return nodes
}

// GetAllNotes collects all notes from degree nodes in an iterator.
func (it Iterator) GetAllNotes() note.Notes {
	if it == nil {
		return nil
	}

	notes := make(note.Notes, 0)
	for node := range it {
		notes = append(notes, node.Note())
	}

	return notes
}
