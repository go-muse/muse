package degree

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegreesIterator_GetAllNotes(t *testing.T) {
	t.Run("GetAllNotes: get all notes from degrees iterator", func(t *testing.T) {
		firstDegree := generateDegrees(7, true)
		iter := firstDegree.IterateOneRound(false)
		notes := iter.GetAllNotes()

		currentDegree := firstDegree
		for _, note := range notes {
			if !reflect.DeepEqual(*currentDegree.Note(), note) {
				t.Errorf("expected: %+v, result: %+v", *currentDegree.Note(), note)
			}

			currentDegree = currentDegree.GetNext()
		}
	})

	t.Run("GetAllNotes: get all notes from nil degrees iterator", func(t *testing.T) {
		var nilDI Iterator
		assert.Nil(t, nilDI.GetAllNotes())
	})
}
