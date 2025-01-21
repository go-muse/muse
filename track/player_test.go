package track

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/note"
)

func TestPlayEvent_Time(t *testing.T) {
	timeVal := 1 * time.Second
	playEvent := &PlayEvent{time: timeVal}

	assert.Equal(t, timeVal, playEvent.Time(), "they should be equal")

	playEvent = &PlayEvent{}
	assert.Equal(t, time.Duration(0), playEvent.Time(), "it should be 0")

	playEvent = nil
	assert.Equal(t, time.Duration(0), playEvent.Time(), "it should be 0")
}

func TestPlayEvent_EventType(t *testing.T) {
	eventType := PlayEventType("Test")
	playEvent := &PlayEvent{eventType: eventType}

	assert.Equal(t, eventType, playEvent.EventType(), "they should be equal")

	playEvent = &PlayEvent{}
	assert.Equal(t, PlayEventType(""), playEvent.EventType(), "it should be empty")

	playEvent = nil
	assert.Equal(t, PlayEventType(""), playEvent.EventType(), "it should be empty")
}

func TestPlayer_Add(t *testing.T) {
	type (
		args struct {
			*Event
			eventType PlayEventType
			time      time.Duration
			playEvents
		}

		testCase struct {
			args args
			want playEvents
		}
	)

	testCases := []testCase{
		{
			args: args{
				Event:     &Event{},
				eventType: EventTypeStart,
				time:      time.Duration(5),
				playEvents: []*PlayEvent{
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
				playEvents: []*PlayEvent{
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
				playEvents: []*PlayEvent{
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
				playEvents: []*PlayEvent{},
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
				playEvents: []*PlayEvent{
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
				playEvents: []*PlayEvent{
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
		result := testCase.args.playEvents.Add(testCase.args.Event, testCase.args.eventType, testCase.args.time)
		assert.Equal(t, testCase.want, result)
	}
}

func TestTrack_Play(t *testing.T) {
	track := NewTrack(&Settings{120, *fraction.New(1, 2), *fraction.New(4, 4)})

	track.events = []*Event{
		{
			startTime:  1,
			note:       note.C.MustNewNote().SetDurationAbs(1),
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
				{startTime: 1, note: note.C.MustNewNote().SetDurationAbs(3), isAbsolute: true},
			},
			want: []*PlayEvent{
				{
					Event:     &Event{startTime: 1, note: note.C.MustNewNote().SetDurationAbs(3), isAbsolute: true},
					eventType: EventTypeStart,
					time:      1,
				},
				{
					Event:     &Event{startTime: 1, note: note.C.MustNewNote().SetDurationAbs(3), isAbsolute: true},
					eventType: EventTypeEnd,
					time:      4,
				},
			},
		},
		{
			events: []*Event{
				{startTime: 1, note: note.C.MustNewNote().SetDurationAbs(3), isAbsolute: true},
				{startTime: 2, note: note.D.MustNewNote().SetDurationAbs(1), isAbsolute: true},
			},
			want: []*PlayEvent{
				{
					Event:     &Event{startTime: 1, note: note.C.MustNewNote().SetDurationAbs(3), isAbsolute: true},
					eventType: EventTypeStart,
					time:      1,
				},
				{
					Event:     &Event{startTime: 2, note: note.D.MustNewNote().SetDurationAbs(1), isAbsolute: true},
					eventType: EventTypeStart,
					time:      2,
				},
				{
					Event:     &Event{startTime: 2, note: note.D.MustNewNote().SetDurationAbs(1), isAbsolute: true},
					eventType: EventTypeEnd,
					time:      3,
				},
				{
					Event:     &Event{startTime: 1, note: note.C.MustNewNote().SetDurationAbs(3), isAbsolute: true},
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
