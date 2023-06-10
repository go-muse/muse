package muse

import (
	"fmt"
	"time"
)

type Event struct {
	startTime  time.Duration
	note       *Note
	isAbsolute bool
}

func NewEvent(note *Note, startTime time.Duration, isAbsolute bool) *Event {
	return &Event{
		startTime:  startTime,
		note:       note,
		isAbsolute: isAbsolute,
	}
}

// String is stringer for Event object.
func (e *Event) String() string {
	return fmt.Sprintf("start time: %v, note: %s, is absolute: %t", e.startTime, e.note.name, e.isAbsolute)
}

func (e *Event) Note() *Note {
	if e == nil {
		return nil
	}

	return e.note
}

func (e *Event) SetNote(n *Note) *Event {
	if e == nil {
		return nil
	}

	e.note = n

	return e
}

func (e *Event) StartTime() time.Duration {
	if e == nil {
		return 0
	}

	return e.startTime
}

func (e *Event) SetStartTime(startTime time.Duration) *Event {
	if e == nil {
		return nil
	}

	e.startTime = startTime

	return e
}

func (e *Event) SetIsAbsolute(isAbsolute bool) *Event {
	if e == nil {
		return e
	}

	e.isAbsolute = isAbsolute

	return e
}

func (e *Event) IsAbsolute() bool {
	if e == nil {
		return false
	}

	return e.isAbsolute
}
