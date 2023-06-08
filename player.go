package muse

import (
	"time"
)

type Player <-chan *PlayEvent

type PlayEvents []*PlayEvent

type EventType string

const (
	EventTypeStart = EventType("start")
	EventTypeEnd   = EventType("end")
)

func (pes *PlayEvents) Add(event *Event, eventType EventType, time time.Duration) PlayEvents {
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

type PlayEvent struct {
	*Event
	eventType EventType
	time      time.Duration
}

func (p *PlayEvent) Time() time.Duration {
	if p == nil {
		return 0
	}

	return p.time
}

func (p *PlayEvent) EventType() EventType {
	if p == nil {
		return ""
	}

	return p.eventType
}

func (t *Track) Player() Player {
	playEvents := make(PlayEvents, 0, len(t.events)*2) //nolint:gomnd // 2 means start+end player events

	for _, event := range t.events {
		playEvents.Add(event, EventTypeStart, event.startTime)
		playEvents.Add(event, EventTypeEnd, t.GetEnd(event))
	}

	c := make(chan *PlayEvent)

	go func(ch chan *PlayEvent, playEvents PlayEvents) {
		defer close(ch)

		for _, playEvent := range playEvents {
			ch <- playEvent
		}
	}(c, playEvents)

	return c
}
