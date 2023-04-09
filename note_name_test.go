package muse

import (
	"reflect"
	"testing"
)

func TestNoteNameString(t *testing.T) {
	expectedNoteName := "C"
	var nn = NoteName(expectedNoteName)
	result := nn.String()
	if _, ok := interface{}(result).(string); !ok {
		t.Errorf("note name Stringer return value: %s is not string: %s", result, reflect.TypeOf(result))
	}

	if expectedNoteName != result {
		t.Errorf("note name Stringer expected: %s actual: %s", expectedNoteName, result)
	}
}
