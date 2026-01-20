package degree

import (
	"iter"

	"github.com/go-muse/muse/note"
)

// Iterator is a function that yields degree nodes one by one.
// It implements iter.Seq[*Node] for use with range loops.
type Iterator = iter.Seq[*Node]

// GetAllDegrees collects all degree nodes from an iterator into a slice.
func GetAllDegrees(it Iterator) []*Node {
	if it == nil {
		return nil
	}

	var nodes []*Node
	for node := range it {
		nodes = append(nodes, node)
	}

	return nodes
}

// GetAllNotes collects all notes from degree nodes in an iterator.
func GetAllNotes(it Iterator) note.Notes {
	if it == nil {
		return nil
	}

	notes := make(note.Notes, 0)
	for node := range it {
		notes = append(notes, node.Note())
	}

	return notes
}
