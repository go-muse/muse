package track

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/note"
)

func TestNewEvent(t *testing.T) {
	n := note.MustNewNoteWithOctave(note.C, 4)
	startTime := 1 * time.Second
	isAbsolute := true

	event := NewEvent(n, startTime, isAbsolute)

	assert.NotNil(t, event, "event should not be nil")
	assert.Equal(t, n, event.Note(), "they should be equal")
	assert.Equal(t, startTime, event.startTime, "they should be equal")
	assert.Equal(t, isAbsolute, event.isAbsolute, "they should be equal")
}

func TestEvent_Note(t *testing.T) {
	n := note.MustNewNoteWithOctave(note.C, 4)
	event := &Event{note: n}

	assert.Equal(t, n, event.Note(), "they should be equal")

	event = &Event{}
	assert.Nil(t, event.Note(), "it should be nil")

	event = nil
	assert.Nil(t, event.Note(), "it should be nil")
}

func TestEvent_SetNote(t *testing.T) {
	note1 := note.MustNewNoteWithOctave(note.C, 4)
	event := &Event{note: note1}

	note2 := note.MustNewNoteWithOctave(note.D, 4)
	event.SetNote(note2)

	assert.Equal(t, note2, event.Note(), "they should be equal")

	event = nil
	assert.Nil(t, event.SetNote(note2), "it should be nil")
}

func TestEvent_StartTime(t *testing.T) {
	startTime := 1 * time.Second
	event := &Event{startTime: startTime}

	assert.Equal(t, startTime, event.StartTime(), "they should be equal")

	event = &Event{}
	assert.Equal(t, time.Duration(0), event.StartTime(), "it should be 0")

	event = nil
	assert.Equal(t, time.Duration(0), event.StartTime(), "it should be 0")
}

func TestEvent_SetStartTime(t *testing.T) {
	startTime := 1 * time.Second
	event := &Event{startTime: 0}

	event.SetStartTime(startTime)

	assert.Equal(t, startTime, event.StartTime(), "they should be equal")

	event = nil
	assert.Nil(t, event.SetStartTime(startTime), "it should be nil")
}

func TestEvent_SetIsAbsolute(t *testing.T) {
	isAbsolute := true
	event := &Event{isAbsolute: false}

	event.SetIsAbsolute(isAbsolute)

	assert.Equal(t, isAbsolute, event.IsAbsolute(), "they should be equal")

	event = nil
	assert.Nil(t, event.SetIsAbsolute(isAbsolute), "it should be nil")
}

func TestEvent_IsAbsolute(t *testing.T) {
	isAbsolute := true
	event := &Event{isAbsolute: isAbsolute}

	assert.Equal(t, isAbsolute, event.IsAbsolute(), "they should be equal")

	event = &Event{}
	assert.False(t, event.IsAbsolute(), "it should be false")

	event = nil
	assert.False(t, event.IsAbsolute(), "it should be false")
}
