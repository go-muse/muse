package muse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTrack(t *testing.T) {
	bpm := uint64(120)
	unit := Fraction{1, 4}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	assert.Equal(t, bpm, track.trackSettings.bpm, "they should be equal")
	assert.Equal(t, unit, track.trackSettings.unit, "they should be equal")
	assert.Equal(t, timeSignature, track.trackSettings.timeSignature, "they should be equal")
	assert.Empty(t, track.events, "track events should be empty")
}

func TestTrack_AddNote(t *testing.T) {
	bpm := uint64(120)
	unit := Fraction{1, 4}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	note := MustNewNote(C, 4)
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
	bpm := uint64(120)
	unit := Fraction{1, 4}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	notes := Notes{
		*MustNewNote(C, 4),
		*MustNewNote(D, 4),
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
	bpm := uint64(120)
	unit := Fraction{1, 4}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	chordNotes := Notes{
		*MustNewNote(C, 4),
		*MustNewNote(D, 4),
	}
	chord := NewChord(*NewDuration(DurationNameWhole), chordNotes...)

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
	bpm := uint64(120)
	unit := Fraction{1, 4}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	event := &Event{
		note:       MustNewNote(C, 4),
		startTime:  time.Second,
		isAbsolute: true,
	}

	result := track.AddEvent(event)

	assert.Equal(t, track, result, "AddEvent should return the same track")

	assert.Len(t, track.events, 1, "there should be one event")
	assert.Equal(t, event, track.events[0], "they should be equal")
}

func TestTrack_Events(t *testing.T) {
	bpm := uint64(120)
	unit := Fraction{1, 4}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	event := &Event{
		note:       MustNewNote(C, 4),
		startTime:  1 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event)

	events := track.Events()

	assert.Len(t, events, 1, "there should be one event")
	assert.Equal(t, event, events[0], "they should be equal")
}

func TestTrack_AddNoteToTheEnd(t *testing.T) {
	bpm := uint64(120)
	unit := Fraction{1, 4}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	event1 := &Event{
		note:       MustNewNote(C, 4).SetAbsoluteDuration(time.Second),
		startTime:  1 * time.Second,
		isAbsolute: false,
	}

	event2 := &Event{
		note:       MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole)),
		startTime:  2 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event1)
	track.AddEvent(event2)

	noteToEnd := MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole))
	track.AddNoteToTheEnd(noteToEnd, true)

	assert.Equal(t, 3, len(track.Events()), "they should be equal")

	lastEvent := track.Events()[len(track.Events())-1]
	assert.Equal(t, noteToEnd, lastEvent.note, "they should be equal")
	assert.Equal(t, 3*time.Second, lastEvent.startTime, "they should be equal")
	assert.True(t, lastEvent.isAbsolute, "it should be true")
}

func TestTrack_FindLastNotes(t *testing.T) {
	t.Run("Track_FindLastNotes: absolute events, one result", func(t *testing.T) {
		expectedResult := []*Note{
			{duration: &Duration{absoluteDuration: 5}},
		}
		track := &Track{
			events: []*Event{
				{startTime: 0, note: &Note{duration: &Duration{absoluteDuration: 7}}, isAbsolute: true}, // 7
				{startTime: 1, note: &Note{duration: &Duration{absoluteDuration: 2}}, isAbsolute: true}, // 3
				{startTime: 2, note: &Note{duration: &Duration{absoluteDuration: 1}}, isAbsolute: true}, // 3
				{startTime: 0, note: &Note{duration: &Duration{absoluteDuration: 4}}, isAbsolute: true}, // 4
				{startTime: 1, note: &Note{duration: &Duration{absoluteDuration: 1}}, isAbsolute: true}, // 2
				{startTime: 3, note: expectedResult[0], isAbsolute: true},                               // 8 is max
				{startTime: 3, note: &Note{duration: &Duration{absoluteDuration: 3}}, isAbsolute: true}, // 6
				{startTime: 1, note: &Note{duration: &Duration{absoluteDuration: 2}}, isAbsolute: true}, // 3
				{startTime: 3, note: &Note{duration: &Duration{absoluteDuration: 4}}, isAbsolute: true}, // 7
				{startTime: 5, note: &Note{duration: &Duration{absoluteDuration: 1}}, isAbsolute: true}, // 6
				{startTime: 0, note: &Note{duration: &Duration{absoluteDuration: 0}}, isAbsolute: true}, // 6
			},
			trackSettings: trackSettings{
				bpm:           120,
				unit:          Fraction{1, 2},
				timeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Duration(8), endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: absolute events, multiple results", func(t *testing.T) {
		expectedResult := []*Note{
			{duration: &Duration{absoluteDuration: 5}},
			{duration: &Duration{absoluteDuration: 4}},
			{duration: &Duration{absoluteDuration: 5}},
		}
		track := &Track{
			events: []*Event{
				{startTime: 0, note: &Note{duration: &Duration{absoluteDuration: 7}}, isAbsolute: true}, // 7
				{startTime: 1, note: &Note{duration: &Duration{absoluteDuration: 2}}, isAbsolute: true}, // 3
				{startTime: 2, note: &Note{duration: &Duration{absoluteDuration: 1}}, isAbsolute: true}, // 3
				{startTime: 0, note: &Note{duration: &Duration{absoluteDuration: 4}}, isAbsolute: true}, // 4
				{startTime: 1, note: &Note{duration: &Duration{absoluteDuration: 1}}, isAbsolute: true}, // 2
				{startTime: 3, note: expectedResult[0], isAbsolute: true},                               // 8 is max
				{startTime: 1, note: &Note{duration: &Duration{absoluteDuration: 2}}, isAbsolute: true}, // 3
				{startTime: 4, note: expectedResult[1], isAbsolute: true},                               // 8 is max
				{startTime: 3, note: expectedResult[2], isAbsolute: true},                               // 8 is max
				{startTime: 5, note: &Note{duration: &Duration{absoluteDuration: 1}}, isAbsolute: true}, // 6
				{startTime: 0, note: &Note{duration: &Duration{absoluteDuration: 0}}, isAbsolute: true}, // 6
			},
			trackSettings: trackSettings{
				bpm:           120,
				unit:          Fraction{1, 2},
				timeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Duration(8), endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: relative events, one result", func(t *testing.T) {
		expectedResult := []*Note{
			{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: &Tuplet{n: 3, m: 2}}}},
		}
		track := &Track{
			events: []*Event{
				{startTime: time.Second, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: &Tuplet{n: 2, m: 3}}}}, isAbsolute: false},            // 1,(6)s
				{startTime: time.Millisecond, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: &Tuplet{n: 2, m: 8}}}}, isAbsolute: false},       // 251ms
				{startTime: time.Millisecond * 400, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 1, tuplet: &Tuplet{n: 2, m: 3}}}}, isAbsolute: false}, // 1,4s
				{startTime: time.Second * 3, note: expectedResult[0], isAbsolute: false},                                                                                           // 4,5s max
				{startTime: time.Second, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 2, tuplet: &Tuplet{n: 2, m: 3}}}}, isAbsolute: false},            // 2.1(6)s
				{startTime: time.Second * 2, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: nil}}}, isAbsolute: false},                        // 3s (1 is default without proper duration name)
			},
			trackSettings: trackSettings{
				bpm:           120,
				unit:          Fraction{1, 2},
				timeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Millisecond*4500, endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: relative events, multiple results", func(t *testing.T) {
		expectedResult := []*Note{
			{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: &Tuplet{n: 3, m: 2}}}},
			{duration: &Duration{relativeDuration: relativeDuration{dots: 1, tuplet: nil}}},
			{duration: &Duration{relativeDuration: relativeDuration{name: DurationNameDoubleWhole, dots: 0, tuplet: nil}}},
		}
		track := &Track{
			events: []*Event{
				{startTime: time.Second, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: &Tuplet{n: 2, m: 3}}}}, isAbsolute: false},            // 1,(6)s
				{startTime: time.Millisecond, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: &Tuplet{n: 2, m: 8}}}}, isAbsolute: false},       // 251ms
				{startTime: time.Millisecond * 400, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 1, tuplet: &Tuplet{n: 2, m: 3}}}}, isAbsolute: false}, // 1,4s
				{startTime: time.Second * 3, note: expectedResult[0], isAbsolute: false},                                                                                           // 4,5s max
				{startTime: time.Second, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 2, tuplet: &Tuplet{n: 2, m: 3}}}}, isAbsolute: false},            // 2.1(6)s
				{startTime: time.Second * 3, note: expectedResult[1], isAbsolute: false},                                                                                           // 4,5s max
				{startTime: time.Second * 2, note: &Note{duration: &Duration{relativeDuration: relativeDuration{dots: 0, tuplet: nil}}}, isAbsolute: false},                        // 3s (1 is default without proper duration name)
				{startTime: time.Millisecond * 2500, note: expectedResult[2], isAbsolute: false},                                                                                   // 4,5s max
			},
			trackSettings: trackSettings{
				bpm:           120,
				unit:          Fraction{1, 2},
				timeSignature: Fraction{4, 4},
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Millisecond*4500, endTime)
		assert.Equal(t, expectedResult, notes)
	})
}

func TestFindLastEvents(t *testing.T) {
	bpm := uint64(120)
	unit := Fraction{1, 2}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	event1 := &Event{
		note:       MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole)),
		startTime:  1 * time.Second,
		isAbsolute: false,
	}

	event2 := &Event{
		note:       MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole)),
		startTime:  2 * time.Second,
		isAbsolute: false,
	}

	event3 := &Event{
		note:       MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole)),
		startTime:  4 * time.Second,
		isAbsolute: false,
	}

	event4 := &Event{
		note:       MustNewNote(C, 4).SetAbsoluteDuration(time.Second),
		startTime:  4 * time.Second,
		isAbsolute: false,
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
	bpm := uint64(120)
	unit := Fraction{1, 2}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	event1 := &Event{
		note:       MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole)),
		startTime:  time.Second,
		isAbsolute: false,
	}

	event2 := &Event{
		note:       MustNewNote(C, 4).SetAbsoluteDuration(time.Second),
		startTime:  2 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event1)
	track.AddEvent(event2)

	end := track.FindEnd()

	assert.Equal(t, 3*time.Second, end, "they should be equal")
}

func TestTrack_GetStartAndEnd(t *testing.T) {
	bpm := uint64(120)
	unit := Fraction{1, 2}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	type want struct {
		start, end time.Duration
	}

	testCases := []struct {
		event *Event
		want  *want
	}{
		{
			event: &Event{
				note:       MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole)),
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
				note:       MustNewNote(C, 4).SetAbsoluteDuration(time.Millisecond * 500),
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
	bpm := uint64(120)
	unit := Fraction{1, 2}
	timeSignature := Fraction{4, 4}

	track := NewTrack(bpm, unit, timeSignature)

	testCases := []struct {
		event *Event
		want  time.Duration
	}{
		{
			event: &Event{
				note:       MustNewNote(C, 4).SetDuration(*NewDuration(DurationNameWhole)),
				startTime:  time.Second,
				isAbsolute: false,
			},
			want: 2 * time.Second,
		},
		{
			event: &Event{
				note:       MustNewNote(C, 4).SetAbsoluteDuration(time.Millisecond * 500),
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
