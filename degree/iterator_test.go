package degree

import (
	"testing"
)

func TestGetAllNotes(t *testing.T) {
	t.Run("GetAllNotes: get all notes from nodes iterator", func(t *testing.T) {
		firstNode := generateDegreeNodes(7, true)
		notes := firstNode.IterateOneRound(false).GetAllNotes()

		currentNode := firstNode
		for _, n := range notes {
			if !currentNode.Note().EqualByName(n) {
				t.Errorf("expected: %+v, result: %+v", currentNode.Note(), n)
			}
			currentNode = currentNode.GetNext()
		}
	})
}
