package degree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllNotes(t *testing.T) {
	t.Run("GetAllNotes: get all notes from nodes iterator", func(t *testing.T) {
		firstNode := generateDegreeNodes(7, true)
		notes := GetAllNotes(firstNode.IterateOneRound(false))

		currentNode := firstNode
		for _, n := range notes {
			if !currentNode.Note().EqualByName(n) {
				t.Errorf("expected: %+v, result: %+v", currentNode.Note(), n)
			}
			currentNode = currentNode.GetNext()
		}
	})

	t.Run("GetAllNotes: get all notes from nil iterator", func(t *testing.T) {
		result := GetAllNotes(nil)
		assert.Nil(t, result)
	})
}
