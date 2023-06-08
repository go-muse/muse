package muse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_Add(t *testing.T) {
	type (
		args struct {
			*Event
			eventType EventType
			time      time.Duration
			PlayEvents
		}

		testCase struct {
			args args
			want PlayEvents
		}
	)

	testCases := []testCase{
		{
			args: args{
				Event:     &Event{},
				eventType: EventTypeStart,
				time:      time.Duration(5),
				PlayEvents: []*PlayEvent{
					{Event: &Event{}, time: 3},
					{Event: &Event{}, time: 4},
					{Event: &Event{}, time: 6},
				},
			},
			want: []*PlayEvent{
				{Event: &Event{}, time: 3},
				{Event: &Event{}, time: 4},
				{Event: &Event{}, time: 5, eventType: EventTypeStart},
				{Event: &Event{}, time: 6},
			},
		},
		{
			args: args{
				Event:     &Event{},
				eventType: EventTypeStart,
				time:      time.Duration(5),
				PlayEvents: []*PlayEvent{
					{Event: &Event{}, time: 1},
					{Event: &Event{}, time: 2},
					{Event: &Event{}, time: 4},
				},
			},
			want: []*PlayEvent{
				{Event: &Event{}, time: 1},
				{Event: &Event{}, time: 2},
				{Event: &Event{}, time: 4},
				{Event: &Event{}, time: 5, eventType: EventTypeStart},
			},
		},
		{
			args: args{
				Event:     &Event{},
				eventType: EventTypeStart,
				time:      time.Duration(1),
				PlayEvents: []*PlayEvent{
					{Event: &Event{}, time: 3},
					{Event: &Event{}, time: 4},
					{Event: &Event{}, time: 6},
				},
			},
			want: []*PlayEvent{
				{Event: &Event{}, time: 1, eventType: EventTypeStart},
				{Event: &Event{}, time: 3},
				{Event: &Event{}, time: 4},
				{Event: &Event{}, time: 6},
			},
		},
		{
			args: args{
				Event:      &Event{},
				eventType:  EventTypeStart,
				time:       time.Duration(1),
				PlayEvents: []*PlayEvent{},
			},
			want: []*PlayEvent{
				{Event: &Event{}, time: 1, eventType: EventTypeStart},
			},
		},
		{
			args: args{
				Event:     &Event{},
				eventType: EventTypeStart,
				time:      time.Duration(1),
				PlayEvents: []*PlayEvent{
					{Event: &Event{}, time: 1, eventType: EventTypeEnd},
					{Event: &Event{}, time: 1, eventType: EventTypeEnd},
				},
			},
			want: []*PlayEvent{
				{Event: &Event{}, time: 1, eventType: EventTypeEnd},
				{Event: &Event{}, time: 1, eventType: EventTypeEnd},
				{Event: &Event{}, time: 1, eventType: EventTypeStart},
			},
		},
		{
			args: args{
				Event:     &Event{},
				eventType: EventTypeStart,
				time:      time.Duration(1),
				PlayEvents: []*PlayEvent{
					{Event: &Event{}, time: 1, eventType: EventTypeEnd},
				},
			},
			want: []*PlayEvent{
				{Event: &Event{}, time: 1, eventType: EventTypeEnd},
				{Event: &Event{}, time: 1, eventType: EventTypeStart},
			},
		},
	}

	for _, testCase := range testCases {
		result := testCase.args.PlayEvents.Add(testCase.args.Event, testCase.args.eventType, testCase.args.time)
		assert.Equal(t, testCase.want, result)
	}
}

func TestTrack_Play(t *testing.T) {
	track := NewTrack(120, Fraction{1, 2}, Fraction{4, 4})

	track.events = []*Event{
		{
			startTime:  1,
			note:       newNote(C).SetAbsoluteDuration(1),
			isAbsolute: true,
		},
	}

	type testCase struct {
		events []*Event
		want   []*PlayEvent
	}

	testCases := []testCase{
		{
			events: []*Event{
				{startTime: 1, note: newNote(C).SetAbsoluteDuration(3), isAbsolute: true},
			},
			want: []*PlayEvent{
				{
					Event:     &Event{startTime: 1, note: newNote(C).SetAbsoluteDuration(3), isAbsolute: true},
					eventType: EventTypeStart,
					time:      1,
				},
				{
					Event:     &Event{startTime: 1, note: newNote(C).SetAbsoluteDuration(3), isAbsolute: true},
					eventType: EventTypeEnd,
					time:      4,
				},
			},
		},
		{
			events: []*Event{
				{startTime: 1, note: newNote(C).SetAbsoluteDuration(3), isAbsolute: true},
				{startTime: 2, note: newNote(D).SetAbsoluteDuration(1), isAbsolute: true},
			},
			want: []*PlayEvent{
				{
					Event:     &Event{startTime: 1, note: newNote(C).SetAbsoluteDuration(3), isAbsolute: true},
					eventType: EventTypeStart,
					time:      1,
				},
				{
					Event:     &Event{startTime: 2, note: newNote(D).SetAbsoluteDuration(1), isAbsolute: true},
					eventType: EventTypeStart,
					time:      2,
				},
				{
					Event:     &Event{startTime: 2, note: newNote(D).SetAbsoluteDuration(1), isAbsolute: true},
					eventType: EventTypeEnd,
					time:      3,
				},
				{
					Event:     &Event{startTime: 1, note: newNote(C).SetAbsoluteDuration(3), isAbsolute: true},
					eventType: EventTypeEnd,
					time:      4,
				},
			},
		},
	}

	for _, testCase := range testCases {
		track.events = testCase.events

		var result []*PlayEvent
		for event := range track.Player() {
			result = append(result, event)
		}

		assert.Equal(t, testCase.want, result)
	}
}
