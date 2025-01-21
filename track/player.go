package track

import (
	"time"
)

// Player is iterable ordered by time set of events of the Track.
type Player <-chan *PlayEvent

type playEvents []*PlayEvent

// PlayEventType is the event type - start or end of the Event.
type PlayEventType string

const (
	EventTypeStart = PlayEventType("start")
	EventTypeEnd   = PlayEventType("end")
)

// Add adds an event to the slice while maintaining the sort order by event time.
func (pes *playEvents) Add(event *Event, eventType PlayEventType, time time.Duration) playEvents {
	playEvent := &PlayEvent{
		Event:     event,
		eventType: eventType,
		time:      time,
	}

	for i := 0; i <= len(*pes)-1; i++ {
		if playEvent.time < (*pes)[i].time {
			*pes = append(*pes, nil)
			copy((*pes)[i+1:], (*pes)[i:])
			(*pes)[i] = playEvent

			return *pes
		}
	}

	*pes = append(*pes, playEvent)

	return *pes
}

// PlayEvent signifies the start or end of an Event at the moment of time.
type PlayEvent struct {
	*Event
	eventType PlayEventType
	time      time.Duration
}

// Time returns duratoin al time of occurrence of the play event.
func (p *PlayEvent) Time() time.Duration {
	if p == nil {
		return 0
	}

	return p.time
}

// EventType returns type of the event.
func (p *PlayEvent) EventType() PlayEventType {
	if p == nil {
		return ""
	}

	return p.eventType
}

// Player returns channel with events ordered by time.
func (t *Track) Player() Player {
	const startPlusEndPlayerEvents = 2
	pes := make(playEvents, 0, len(t.events)*startPlusEndPlayerEvents)

	for _, event := range t.events {
		pes.Add(event, EventTypeStart, event.startTime)
		pes.Add(event, EventTypeEnd, t.GetEnd(event))
	}

	c := make(chan *PlayEvent)

	go func(ch chan *PlayEvent, pes playEvents) {
		defer close(ch)

		for _, playEvent := range pes {
			ch <- playEvent
		}
	}(c, pes)

	return c
}
