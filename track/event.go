package track

import (
	"fmt"
	"time"

	"github.com/go-muse/muse/note"
)

// Event is a single note played at a specific time.
type Event struct {
	startTime  time.Duration
	note       note.Note
	isAbsolute bool
}

// NewEvent creates a new event with the specified note, start time, and absolute flag.
func NewEvent(note note.Note, startTime time.Duration, isAbsolute bool) *Event {
	return &Event{
		startTime:  startTime,
		note:       note,
		isAbsolute: isAbsolute,
	}
}

// String is stringer for Event object.
func (e Event) String() string {
	return fmt.Sprintf("start time: %v, note: %s, is absolute: %t", e.startTime, e.note.Name(), e.isAbsolute)
}

// Note returns the note of the event.
func (e Event) Note() note.Note {
	return e.note
}

// WithNote sets the note of the event and returns the event.
func (e Event) WithNote(n note.Note) Event {
	e.note = n
	return e
}

// StartTime returns the start time of the event.
func (e Event) StartTime() time.Duration {
	return e.startTime
}

// WithStartTime sets the start time of the event and returns the event.
func (e Event) WithStartTime(startTime time.Duration) Event {
	e.startTime = startTime
	return e
}

// WithIsAbsolute sets the absolute flag of the event and returns the event.
func (e Event) WithIsAbsolute(isAbsolute bool) Event {
	e.isAbsolute = isAbsolute
	return e
}

// IsAbsolute returns the absolute flag of the event.
func (e Event) IsAbsolute() bool {
	return e.isAbsolute
}
