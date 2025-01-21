package track

import (
	"fmt"
	"time"

	"github.com/go-muse/muse/note"
)

// Event is a single note played at a specific time.
type Event struct {
	startTime  time.Duration
	note       *note.Note
	isAbsolute bool
}

// NewEvent creates a new event with the specified note, start time, and absolute flag.
func NewEvent(note *note.Note, startTime time.Duration, isAbsolute bool) *Event {
	return &Event{
		startTime:  startTime,
		note:       note,
		isAbsolute: isAbsolute,
	}
}

// String is stringer for Event object.
func (e *Event) String() string {
	return fmt.Sprintf("start time: %v, note: %s, is absolute: %t", e.startTime, e.note.Name(), e.isAbsolute)
}

// Note returns the note of the event.
func (e *Event) Note() *note.Note {
	if e == nil {
		return nil
	}

	return e.note
}

// SetNote sets the note of the event and returns the event.
func (e *Event) SetNote(n *note.Note) *Event {
	if e == nil {
		return nil
	}

	e.note = n

	return e
}

// StartTime returns the start time of the event.
func (e *Event) StartTime() time.Duration {
	if e == nil {
		return 0
	}

	return e.startTime
}

// SetStartTime sets the start time of the event and returns the event.
func (e *Event) SetStartTime(startTime time.Duration) *Event {
	if e == nil {
		return nil
	}

	e.startTime = startTime

	return e
}

// SetIsAbsolute sets the absolute flag of the event and returns the event.
func (e *Event) SetIsAbsolute(isAbsolute bool) *Event {
	if e == nil {
		return e
	}

	e.isAbsolute = isAbsolute

	return e
}

// IsAbsolute returns the absolute flag of the event.
func (e *Event) IsAbsolute() bool {
	if e == nil {
		return false
	}

	return e.isAbsolute
}
