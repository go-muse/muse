package muse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTrack(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 4},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	assert.Equal(t, trackSettings.BPM, track.TrackSettings.BPM, "they should be equal")
	assert.Equal(t, trackSettings.Unit, track.TrackSettings.Unit, "they should be equal")
	assert.Equal(t, trackSettings.TimeSignature, track.TrackSettings.TimeSignature, "they should be equal")
	assert.Empty(t, track.events, "track events should be empty")
}

func TestTrack_AddNote(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 4},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	note := MustNewNoteWithOctave(C, 4)
	startTime := time.Second
	isAbsolute := true

	track.AddNote(note, startTime, isAbsolute)

	assert.Len(t, track.events, 1, "there should be one event")

	event := track.events[0]

	assert.Equal(t, note, event.note, "they should be equal")
	assert.Equal(t, startTime, event.startTime, "they should be equal")
	assert.Equal(t, isAbsolute, event.isAbsolute, "they should be equal")
}

func TestTrack_AddNotes(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 4},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	notes := Notes{
		*MustNewNoteWithOctave(C, 4),
		*MustNewNoteWithOctave(D, 4),
	}
	startTime := time.Second
	isAbsolute := true

	track.AddNotes(notes, startTime, isAbsolute)

	assert.Len(t, track.events, 2, "there should be two events")

	for i, event := range track.events {
		assert.Equal(t, &notes[i], event.note, "they should be equal")
		assert.Equal(t, startTime, event.startTime, "they should be equal")
		assert.Equal(t, isAbsolute, event.isAbsolute, "they should be equal")
	}
}

func TestTrack_AddChord(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 4},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	chordNotes := Notes{
		*MustNewNoteWithOctave(C, 4),
		*MustNewNoteWithOctave(D, 4),
	}
	chord := NewChord(chordNotes...).SetDurationRel(NewDurationRel(DurationNameWhole))

	startTime := time.Second
	isAbsolute := true

	track.AddChord(chord, startTime, isAbsolute)

	assert.Len(t, track.events, 2, "there should be two events")

	for i, event := range track.events {
		assert.Equal(t, &chord.notes[i], event.note, "notes should be equal")
		assert.Equal(t, startTime, event.startTime, "startTime should be equal")
		assert.Equal(t, isAbsolute, event.isAbsolute, "isAbsolute should be equal")
	}
}

func TestTrack_AddEvent(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 4},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	event := &Event{
		note:       MustNewNoteWithOctave(C, 4),
		startTime:  time.Second,
		isAbsolute: true,
	}

	result := track.AddEvent(event)

	assert.Equal(t, track, result, "AddEvent should return the same track")

	assert.Len(t, track.events, 1, "there should be one event")
	assert.Equal(t, event, track.events[0], "they should be equal")
}

func TestTrack_Events(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 4},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	event := &Event{
		note:       MustNewNoteWithOctave(C, 4),
		startTime:  1 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event)

	events := track.Events()

	assert.Len(t, events, 1, "there should be one event")
	assert.Equal(t, event, events[0], "they should be equal")
}

func TestTrack_AddNoteToTheEnd(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 4},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	event1 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationAbs(time.Second),
		startTime:  1 * time.Second,
		isAbsolute: true,
	}

	event2 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole)),
		startTime:  2 * time.Second,
		isAbsolute: false,
	}

	track.AddEvent(event1)
	track.AddEvent(event2)

	noteToEnd := MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole))
	track.AddNoteToTheEnd(noteToEnd, false)

	assert.Equal(t, 3, len(track.Events()), "they should be equal")

	lastEvent := track.Events()[len(track.Events())-1]
	assert.Equal(t, noteToEnd, lastEvent.note, "they should be equal")
	assert.Equal(t, 4*time.Second, lastEvent.startTime, "they should be equal")
	assert.Equal(t, 6*time.Second, track.GetEnd(lastEvent), "they should be equal")
	assert.False(t, lastEvent.isAbsolute, "it should be false")
}

func TestTrack_FindLastNotes(t *testing.T) {
	t.Run("Track_FindLastNotes: absolute events, one result", func(t *testing.T) {
		expectedResult := []*Note{
			{durationAbs: 5},
		}

		track := &Track{
			events: []*Event{
				{startTime: 0, note: &Note{durationAbs: 7}, isAbsolute: true}, // 7
				{startTime: 1, note: &Note{durationAbs: 2}, isAbsolute: true}, // 3
				{startTime: 2, note: &Note{durationAbs: 1}, isAbsolute: true}, // 3
				{startTime: 0, note: &Note{durationAbs: 4}, isAbsolute: true}, // 4
				{startTime: 1, note: &Note{durationAbs: 1}, isAbsolute: true}, // 2
				{startTime: 3, note: expectedResult[0], isAbsolute: true},     // 8 is max
				{startTime: 3, note: &Note{durationAbs: 3}, isAbsolute: true}, // 6
				{startTime: 1, note: &Note{durationAbs: 2}, isAbsolute: true}, // 3
				{startTime: 3, note: &Note{durationAbs: 4}, isAbsolute: true}, // 7
				{startTime: 5, note: &Note{durationAbs: 1}, isAbsolute: true}, // 6
				{startTime: 0, note: &Note{durationAbs: 0}, isAbsolute: true}, // 6
			},
			TrackSettings: TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 2},
				TimeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Duration(8), endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: absolute events, multiple results", func(t *testing.T) {
		expectedResult := []*Note{
			{durationAbs: 5},
			{durationAbs: 4},
			{durationAbs: 5},
		}
		track := &Track{
			events: []*Event{
				{startTime: 0, note: &Note{durationAbs: 7}, isAbsolute: true}, // 7
				{startTime: 1, note: &Note{durationAbs: 2}, isAbsolute: true}, // 3
				{startTime: 2, note: &Note{durationAbs: 1}, isAbsolute: true}, // 3
				{startTime: 0, note: &Note{durationAbs: 4}, isAbsolute: true}, // 4
				{startTime: 1, note: &Note{durationAbs: 1}, isAbsolute: true}, // 2
				{startTime: 3, note: expectedResult[0], isAbsolute: true},     // 8 is max
				{startTime: 1, note: &Note{durationAbs: 2}, isAbsolute: true}, // 3
				{startTime: 4, note: expectedResult[1], isAbsolute: true},     // 8 is max
				{startTime: 3, note: expectedResult[2], isAbsolute: true},     // 8 is max
				{startTime: 5, note: &Note{durationAbs: 1}, isAbsolute: true}, // 6
				{startTime: 0, note: &Note{durationAbs: 0}, isAbsolute: true}, // 6
			},
			TrackSettings: TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 2},
				TimeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Duration(8), endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: relative events, one result", func(t *testing.T) {
		expectedResult := []*Note{
			{durationRel: &DurationRel{dots: 0, tuplet: &Tuplet{n: 3, m: 2}}},
		}
		track := &Track{
			events: []*Event{
				{startTime: time.Second, note: &Note{durationRel: &DurationRel{dots: 0, tuplet: &Tuplet{n: 2, m: 3}}}, isAbsolute: false},            // 1,(6)s
				{startTime: time.Millisecond, note: &Note{durationRel: &DurationRel{dots: 0, tuplet: &Tuplet{n: 2, m: 8}}}, isAbsolute: false},       // 251ms
				{startTime: time.Millisecond * 400, note: &Note{durationRel: &DurationRel{dots: 1, tuplet: &Tuplet{n: 2, m: 3}}}, isAbsolute: false}, // 1,4s
				{startTime: time.Second * 3, note: expectedResult[0], isAbsolute: false},                                                             // 4,5s max
				{startTime: time.Second, note: &Note{durationRel: &DurationRel{dots: 2, tuplet: &Tuplet{n: 2, m: 3}}}, isAbsolute: false},            // 2.1(6)s
				{startTime: time.Second * 2, note: &Note{durationRel: &DurationRel{dots: 0, tuplet: nil}}, isAbsolute: false},                        // 3s (1 is default without proper duration name)
			},
			TrackSettings: TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 2},
				TimeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Millisecond*4500, endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: relative events, multiple results", func(t *testing.T) {
		expectedResult := []*Note{
			{durationRel: &DurationRel{dots: 0, tuplet: &Tuplet{n: 3, m: 2}}},
			{durationRel: &DurationRel{dots: 1, tuplet: nil}},
			{durationRel: &DurationRel{name: DurationNameDoubleWhole, dots: 0, tuplet: nil}},
		}
		track := &Track{
			events: []*Event{
				{startTime: time.Second, note: &Note{durationRel: &DurationRel{dots: 0, tuplet: &Tuplet{n: 2, m: 3}}}, isAbsolute: false},            // 1,(6)s
				{startTime: time.Millisecond, note: &Note{durationRel: &DurationRel{dots: 0, tuplet: &Tuplet{n: 2, m: 8}}}, isAbsolute: false},       // 251ms
				{startTime: time.Millisecond * 400, note: &Note{durationRel: &DurationRel{dots: 1, tuplet: &Tuplet{n: 2, m: 3}}}, isAbsolute: false}, // 1,4s
				{startTime: time.Second * 3, note: expectedResult[0], isAbsolute: false},                                                             // 4,5s max
				{startTime: time.Second, note: &Note{durationRel: &DurationRel{dots: 2, tuplet: &Tuplet{n: 2, m: 3}}}, isAbsolute: false},            // 2.1(6)s
				{startTime: time.Second * 3, note: expectedResult[1], isAbsolute: false},                                                             // 4,5s max
				{startTime: time.Second * 2, note: &Note{durationRel: &DurationRel{dots: 0, tuplet: nil}}, isAbsolute: false},                        // 3s (1 is default without proper duration name)
				{startTime: time.Millisecond * 2500, note: expectedResult[2], isAbsolute: false},
			},
			TrackSettings: TrackSettings{
				BPM:           120,
				Unit:          Fraction{1, 2},
				TimeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Millisecond*4500, endTime)
		assert.Equal(t, expectedResult, notes)
	})
}

func TestFindLastEvents(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 2},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	event1 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole)),
		startTime:  1 * time.Second,
		isAbsolute: false,
	}

	event2 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole)),
		startTime:  2 * time.Second,
		isAbsolute: false,
	}

	event3 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole)),
		startTime:  4 * time.Second,
		isAbsolute: false,
	}

	event4 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationAbs(1 * time.Second),
		startTime:  4 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event1)
	track.AddEvent(event2)
	track.AddEvent(event3)
	track.AddEvent(event4)

	events, maxEnd := track.FindLastEvents()

	assert.Equal(t, []*Event{event3, event4}, events, "they should be equal")
	assert.Equal(t, 5*time.Second, maxEnd, "they should be equal")
}

func TestFindEnd(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 2},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	event1 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole)),
		startTime:  time.Second,
		isAbsolute: false,
	}

	event2 := &Event{
		note:       MustNewNoteWithOctave(C, 4).SetDurationAbs(time.Second),
		startTime:  2 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event1)
	track.AddEvent(event2)

	end := track.FindEnd()

	assert.Equal(t, 3*time.Second, end, "they should be equal")
}

func TestTrack_GetStartAndEnd(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 2},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	type want struct {
		start, end time.Duration
	}

	testCases := []struct {
		event *Event
		want  *want
	}{
		{
			event: &Event{
				note:       MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole)),
				startTime:  time.Second,
				isAbsolute: false,
			},
			want: &want{
				start: time.Second,
				end:   2 * time.Second,
			},
		},
		{
			event: &Event{
				note:       MustNewNoteWithOctave(C, 4).SetDurationAbs(time.Millisecond * 500),
				startTime:  time.Millisecond * 500,
				isAbsolute: true,
			},
			want: &want{
				start: time.Millisecond * 500,
				end:   time.Second,
			},
		},
	}

	var start, end time.Duration
	for _, testCase := range testCases {
		start, end = track.GetStartAndEnd(testCase.event)

		assert.Equal(t, testCase.want.start, start, "they should be equal")
		assert.Equal(t, testCase.want.end, end, "they should be equal")
	}
}

func TestTrack_GetEnd(t *testing.T) {
	trackSettings := &TrackSettings{
		BPM:           uint64(120),
		Unit:          Fraction{1, 2},
		TimeSignature: Fraction{4, 4},
	}

	track := NewTrack(*trackSettings)

	testCases := []struct {
		event *Event
		want  time.Duration
	}{
		{
			event: &Event{
				note:       MustNewNoteWithOctave(C, 4).SetDurationRel(NewDurationRel(DurationNameWhole)),
				startTime:  time.Second,
				isAbsolute: false,
			},
			want: 2 * time.Second,
		},
		{
			event: &Event{
				note:       MustNewNoteWithOctave(C, 4).SetDurationAbs(time.Millisecond * 500),
				startTime:  time.Millisecond * 500,
				isAbsolute: true,
			},
			want: time.Second,
		},
	}

	var end time.Duration
	for _, testCase := range testCases {
		end = track.GetEnd(testCase.event)

		assert.Equal(t, testCase.want, end, "they should be equal")
	}
}
