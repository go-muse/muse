package track

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/go-muse/muse/chord"
	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/note"
)

// Track is a set of Events. Track also contains settings that allow to define the absolute duration of notes in the Events.
type Track struct {
	events []*Event
	*Settings
}

type Settings struct {
	BPM                 uint64
	Unit, TimeSignature fraction.Fraction
}

// GetAmountOfBars calculates and returns amount of bars within one minute.
func (s *Settings) GetAmountOfBars() decimal.Decimal {
	return decimal.NewFromUint64(s.BPM).Mul(s.Unit.MustValue()).Div(s.TimeSignature.MustValue())
}

// NewTrack creates a new track with specified settings.
func NewTrack(trackSettings *Settings) *Track {
	return &Track{
		events:   []*Event{},
		Settings: trackSettings,
	}
}

// AddNote adds a note to the track with the specified start time.
func (t *Track) AddNote(n *note.Note, startTime time.Duration, isAbsolute bool) *Track {
	if t == nil {
		return nil
	}

	t.events = append(t.events, &Event{
		note:       n,
		startTime:  startTime,
		isAbsolute: isAbsolute,
	})

	return t
}

// AddNotes adds notes to the track with the specified start time.
func (t *Track) AddNotes(notes note.Notes, startTime time.Duration, isAbsolute bool) *Track {
	if t == nil {
		return nil
	}

	for _, note := range notes {
		t.events = append(t.events, &Event{
			note:       note.Copy(),
			startTime:  startTime,
			isAbsolute: isAbsolute,
		})
	}

	return t
}

// AddChord adds a notes from the chord to the track with the specified start time.
func (t *Track) AddChord(c *chord.Chord, startTime time.Duration, isAbsolute bool) *Track {
	if t == nil {
		return nil
	}

	for _, chordNote := range c.Notes() {
		t.events = append(t.events, &Event{
			note:       chordNote.Copy(),
			startTime:  startTime,
			isAbsolute: isAbsolute,
		})
	}

	return t
}

// AddEvent adds an event to the track.
func (t *Track) AddEvent(event *Event) *Track {
	if t == nil {
		return nil
	}

	t.events = append(t.events, event)

	return t
}

// Events returns events of the track.
func (t *Track) Events() []*Event {
	if t == nil {
		return nil
	}

	return t.events
}

// AddNoteToTheEnd adds a note to the absolute end of the track.
func (t *Track) AddNoteToTheEnd(n *note.Note, isAbsolute bool) *Track {
	if t == nil {
		return nil
	}

	t.events = append(t.events, &Event{
		note:       n,
		startTime:  t.FindEnd(),
		isAbsolute: isAbsolute,
	})

	return t
}

// FindLastNotes returns the notes whose ending is the end of the track.
func (t *Track) FindLastNotes() (note.Notes, time.Duration) {
	if t == nil || t.events == nil {
		return nil, 0
	}

	var maxEnd time.Duration
	var notes note.Notes
	for _, event := range t.events {
		end := t.GetEnd(event)
		if end > maxEnd {
			maxEnd = end
			notes = note.Notes{event.note}
		} else if end == maxEnd {
			notes = append(notes, event.note)
		}
	}

	return notes, maxEnd
}

// FindLastEvents returns the events whose ending is the end of the track.
func (t *Track) FindLastEvents() ([]*Event, time.Duration) {
	if t == nil || t.events == nil {
		return nil, 0
	}

	var maxEnd time.Duration
	var events []*Event
	for _, event := range t.events {
		end := t.GetEnd(event)
		if end > maxEnd {
			maxEnd = end
			events = []*Event{event}
		} else if end == maxEnd {
			events = append(events, event)
		}
	}

	return events, maxEnd
}

// FindEnd returns the end time of the track (i.e. the length of its time).
func (t *Track) FindEnd() time.Duration {
	if t == nil || t.events == nil {
		return 0
	}

	var end time.Duration
	for _, event := range t.events {
		enventEnd := t.GetEnd(event)
		if enventEnd > end {
			end = enventEnd
		}
	}

	return end
}

// GetStartAndEnd returns the start and end time of the event.
func (t *Track) GetStartAndEnd(event *Event) (time.Duration, time.Duration) {
	if t == nil || event == nil || event.note == nil {
		return 0, 0
	}

	if event.isAbsolute {
		return event.startTime, event.startTime + event.note.Duration()
	}

	return event.startTime, event.startTime + event.note.GetTimeDuration(t.GetAmountOfBars())
}

// GetEnd returns the end time of the event.
func (t *Track) GetEnd(event *Event) time.Duration {
	if t == nil || event == nil || event.note == nil {
		return 0
	}

	if event.isAbsolute {
		return event.startTime + event.note.Duration()
	}

	return event.startTime + event.note.GetTimeDuration(t.GetAmountOfBars())
}
