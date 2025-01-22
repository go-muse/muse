package track

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/chord"
	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/tuplet"
)

func TestNewTrack(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	assert.Equal(t, trackSettings.BPM, track.Settings.BPM, "they should be equal")
	assert.Equal(t, trackSettings.Unit, track.Settings.Unit, "they should be equal")
	assert.Equal(t, trackSettings.TimeSignature, track.Settings.TimeSignature, "they should be equal")
	assert.Empty(t, track.events, "track events should be empty")
}

func TestSettings_GetAmountOfBars(t *testing.T) {
	type (
		args struct {
			bpm           uint64
			unit          *fraction.Fraction
			timeSignature *fraction.Fraction
		}
		want struct {
			amountOfBars decimal.Decimal
		}
	)

	testCases := []struct {
		args args
		want want
	}{
		{
			args: args{
				bpm:           60,                 // 60 bpm
				unit:          fraction.New(1, 1), // unit 1
				timeSignature: fraction.New(1, 1), // time signature 1/1
			},
			want: want{
				amountOfBars: decimal.NewFromUint64(60), // (60bpm*(1/1)/(1/1)) = 60 bars
			},
		},
		{
			args: args{
				bpm:           60,                 // 60 bpm
				unit:          fraction.New(1, 2), // unit 1/2
				timeSignature: fraction.New(3, 4), // time signature 3/4
			},
			want: want{
				amountOfBars: decimal.NewFromUint64(40), // (60bpm*(1/2)/(3/4)) = 40 bars
			},
		},
		{
			args: args{
				bpm:           120,                // 120 bpm
				unit:          fraction.New(1, 2), // unit 1/2
				timeSignature: fraction.New(4, 4), // time signature 4/4
			},
			want: want{
				amountOfBars: decimal.NewFromUint64(60), // (120bpm*(1/2)/(4/4)) = 60 bars
			},
		},
		{
			args: args{
				bpm:           140,                // 140 bpm
				unit:          fraction.New(1, 2), // unit 1/2
				timeSignature: fraction.New(3, 4), // time signature 3/4
			},
			want: want{
				amountOfBars: decimal.NewFromUint64(280).Div(decimal.NewFromUint64(3)), // (140bpm*(1/2)/(3/4)) = 280/3 = 93,(3) bars
			},
		},
	}

	var amountOfBars decimal.Decimal
	var trackSettings *Settings
	for _, testCase := range testCases {
		trackSettings = &Settings{BPM: testCase.args.bpm, Unit: *testCase.args.unit, TimeSignature: *testCase.args.timeSignature}
		amountOfBars = trackSettings.GetAmountOfBars()
		assert.True(t, testCase.want.amountOfBars.Equal(amountOfBars), "expected: %v, actual: %v, args: %d, %+v, %+v", testCase.want.amountOfBars.BigFloat(), amountOfBars.BigFloat(), testCase.args.bpm, testCase.args.unit, testCase.args.timeSignature)
	}
}

func TestTrack_AddNote(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	n := note.MustNewNoteWithOctave(note.C, 4)
	startTime := time.Second
	isAbsolute := true

	track.AddNote(n, startTime, isAbsolute)

	assert.Len(t, track.events, 1, "there should be one event")

	event := track.events[0]

	assert.Equal(t, n, event.note, "they should be equal")
	assert.Equal(t, startTime, event.startTime, "they should be equal")
	assert.Equal(t, isAbsolute, event.isAbsolute, "they should be equal")
}

func TestTrack_AddNotes(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	notes := note.Notes{
		note.MustNewNoteWithOctave(note.C, 4),
		note.MustNewNoteWithOctave(note.D, 4),
	}
	startTime := time.Second
	isAbsolute := true

	track.AddNotes(notes, startTime, isAbsolute)

	assert.Len(t, track.events, 2, "there should be two events")

	for i, event := range track.events {
		assert.Equal(t, notes[i], event.note, "they should be equal")
		assert.Equal(t, startTime, event.startTime, "they should be equal")
		assert.Equal(t, isAbsolute, event.isAbsolute, "they should be equal")
	}
}

func TestTrack_AddChord(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	chordNotes := note.Notes{
		note.MustNewNoteWithOctave(note.C, 4),
		note.MustNewNoteWithOctave(note.D, 4),
	}
	c := chord.NewChord(chordNotes...).SetValue(duration.NewRelative(duration.NameWhole))

	startTime := time.Second
	isAbsolute := true

	track.AddChord(c, startTime, isAbsolute)

	assert.Len(t, track.events, 2, "there should be two events")

	for i, event := range track.events {
		assert.Equal(t, c.Notes()[i], event.note, "notes should be equal")
		assert.Equal(t, startTime, event.startTime, "startTime should be equal")
		assert.Equal(t, isAbsolute, event.isAbsolute, "isAbsolute should be equal")
	}
}

func TestTrack_AddEvent(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	event := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4),
		startTime:  time.Second,
		isAbsolute: true,
	}

	result := track.AddEvent(event)

	assert.Equal(t, track, result, "AddEvent should return the same track")

	assert.Len(t, track.events, 1, "there should be one event")
	assert.Equal(t, event, track.events[0], "they should be equal")
}

func TestTrack_Events(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	event := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4),
		startTime:  1 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event)

	events := track.Events()

	assert.Len(t, events, 1, "there should be one event")
	assert.Equal(t, event, events[0], "they should be equal")
}

func TestTrack_AddNoteToTheEnd(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	event1 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetDuration(time.Second),
		startTime:  1 * time.Second,
		isAbsolute: true,
	}

	event2 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole)),
		startTime:  2 * time.Second,
		isAbsolute: false,
	}

	track.AddEvent(event1)
	track.AddEvent(event2)

	noteToEnd := note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole))
	track.AddNoteToTheEnd(noteToEnd, false)

	assert.Len(t, track.Events(), 3, "they should be equal")

	lastEvent := track.Events()[len(track.Events())-1]
	assert.Equal(t, noteToEnd, lastEvent.note, "they should be equal")
	assert.Equal(t, 4*time.Second, lastEvent.startTime, "they should be equal")
	assert.Equal(t, 6*time.Second, track.GetEnd(lastEvent), "they should be equal")
	assert.False(t, lastEvent.isAbsolute, "it should be false")
}

func TestTrack_FindLastNotes(t *testing.T) {
	t.Run("Track_FindLastNotes: absolute events, one result", func(t *testing.T) {
		expectedResult := note.Notes{
			note.C.MustNewNote().SetDuration(15 * time.Second),
		}

		track := &Track{
			events: []*Event{
				{startTime: 0 * time.Second, note: note.C.MustNewNote().SetDuration(7 * time.Second), isAbsolute: true}, // 7
				{startTime: 1 * time.Second, note: note.C.MustNewNote().SetDuration(3 * time.Second), isAbsolute: true}, // 3
				{startTime: 2 * time.Second, note: note.C.MustNewNote().SetDuration(3 * time.Second), isAbsolute: true}, // 3
				{startTime: 0 * time.Second, note: note.C.MustNewNote().SetDuration(4 * time.Second), isAbsolute: true}, // 4
				{startTime: 1 * time.Second, note: note.C.MustNewNote().SetDuration(2 * time.Second), isAbsolute: true}, // 2
				{startTime: 3 * time.Second, note: expectedResult[0], isAbsolute: true},                                 // 8 is max
				{startTime: 3 * time.Second, note: note.C.MustNewNote().SetDuration(6 * time.Second), isAbsolute: true}, // 6
				{startTime: 1 * time.Second, note: note.C.MustNewNote().SetDuration(3 * time.Second), isAbsolute: true}, // 3
				{startTime: 3 * time.Second, note: note.C.MustNewNote().SetDuration(7 * time.Second), isAbsolute: true}, // 7
				{startTime: 5 * time.Second, note: note.C.MustNewNote().SetDuration(6 * time.Second), isAbsolute: true}, // 6
				{startTime: 0 * time.Second, note: note.C.MustNewNote().SetDuration(6 * time.Second), isAbsolute: true}, // 6
			},
			Settings: &Settings{
				BPM:           120,
				Unit:          *fraction.New(1, 2),
				TimeSignature: *fraction.New(4, 4),
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, 18*time.Second, endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: absolute events, multiple results", func(t *testing.T) {
		expectedResult := note.Notes{
			note.C.MustNewNote().SetDuration(5 * time.Second),
			note.C.MustNewNote().SetDuration(4 * time.Second),
			note.C.MustNewNote().SetDuration(5 * time.Second),
		}
		track := &Track{
			events: []*Event{
				{startTime: 0 * time.Second, note: note.C.MustNewNote().SetDuration(7 * time.Second), isAbsolute: true}, // 7
				{startTime: 1 * time.Second, note: note.C.MustNewNote().SetDuration(3 * time.Second), isAbsolute: true}, // 4
				{startTime: 2 * time.Second, note: note.C.MustNewNote().SetDuration(3 * time.Second), isAbsolute: true}, // 5
				{startTime: 0 * time.Second, note: note.C.MustNewNote().SetDuration(4 * time.Second), isAbsolute: true}, // 4
				{startTime: 1 * time.Second, note: note.C.MustNewNote().SetDuration(2 * time.Second), isAbsolute: true}, // 3
				{startTime: 3 * time.Second, note: expectedResult[0], isAbsolute: true},                                 // 8 is max
				{startTime: 1 * time.Second, note: note.C.MustNewNote().SetDuration(3 * time.Second), isAbsolute: true}, // 4
				{startTime: 4 * time.Second, note: expectedResult[1], isAbsolute: true},                                 // 8 is max
				{startTime: 3 * time.Second, note: expectedResult[2], isAbsolute: true},                                 // 8 is max
				{startTime: 5 * time.Second, note: note.C.MustNewNote().SetDuration(1 * time.Second), isAbsolute: true}, // 6
				{startTime: 0 * time.Second, note: note.C.MustNewNote().SetDuration(6 * time.Second), isAbsolute: true}, // 6
			},
			Settings: &Settings{
				BPM:           120,
				Unit:          *fraction.New(1, 2),
				TimeSignature: *fraction.New(4, 4),
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, 8*time.Second, endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: relative events, one result", func(t *testing.T) {
		expectedResult := note.Notes{
			note.C.MustMakeNote().SetValue(duration.NewRelative(duration.NameWhole).SetTupletTriplet()),
		}
		track := &Track{
			events: []*Event{
				{startTime: time.Second, note: note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole).SetTupletDuplet())},                                 // 1,(6)s0
				{startTime: time.Millisecond, note: note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole).SetTuplet(tuplet.New(2, 8)))},                  // 251ms
				{startTime: time.Millisecond * 400, note: note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole).SetTuplet(tuplet.New(2, 3)).SetDots(1))}, // 1,4s
				{startTime: time.Second * 3, note: expectedResult[0], isAbsolute: false},                                                                                  // 4,5s max
				{startTime: time.Second, note: note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole).SetTuplet(tuplet.New(2, 3)).SetDots(2))},            // 2.1(6)s
				{startTime: time.Second * 2, note: note.C.MustNewNote(), isAbsolute: false},                                                                               // 3s (1 is default without proper duration name)
			},
			Settings: &Settings{
				BPM:           120,
				Unit:          *fraction.New(1, 2),
				TimeSignature: *fraction.New(4, 4),
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Millisecond*4500, endTime)
		assert.Equal(t, expectedResult, notes)
	})

	t.Run("Track_FindLastNotes: relative events, multiple results", func(t *testing.T) {
		expectedResult := note.Notes{
			note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole).SetTupletTriplet()),
			note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole).SetDots(1)),
			note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameDoubleWhole)),
		}
		track := &Track{
			events: []*Event{
				{startTime: time.Second, note: note.MustNewNote(note.C).SetValue(duration.NewRelative(duration.NameWhole).SetTupletDuplet())},                           // 1,(6)s
				{startTime: time.Millisecond, note: note.MustNewNote(note.C).SetValue(duration.NewRelative(duration.NameWhole).SetTuplet(tuplet.New(2, 8)).SetDots(1))}, // 251ms
				{startTime: time.Millisecond, note: note.MustNewNote(note.C).SetValue(duration.NewRelative(duration.NameWhole).SetTuplet(tuplet.New(2, 3)).SetDots(1))}, // 1,4s
				{startTime: time.Second * 3, note: expectedResult[0], isAbsolute: false},                                                                                // 4,5s max
				{startTime: time.Second, note: note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole).SetDots(2).SetTupletDuplet()), isAbsolute: false}, // 2.1(6)s
				{startTime: time.Second * 3, note: expectedResult[1], isAbsolute: false},                                                                                // 4,5s max
				{startTime: time.Second * 2, note: note.C.MustNewNote().SetValue(duration.NewRelative(duration.NameWhole)), isAbsolute: false},                          // 3s (1 is default without proper duration name)
				{startTime: time.Millisecond * 2500, note: expectedResult[2], isAbsolute: false},
			},
			Settings: &Settings{
				BPM:           120,
				Unit:          *fraction.New(1, 2),
				TimeSignature: *fraction.New(4, 4),
			},
		}

		notes, endTime := track.FindLastNotes()
		assert.Equal(t, time.Millisecond*4500, endTime)
		assert.Equal(t, expectedResult, notes)
	})
}

func TestFindLastEvents(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 2),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	event1 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole)),
		startTime:  1 * time.Second,
		isAbsolute: false,
	}

	event2 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole)),
		startTime:  2 * time.Second,
		isAbsolute: false,
	}

	event3 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole)),
		startTime:  4 * time.Second,
		isAbsolute: false,
	}

	event4 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetDuration(1 * time.Second),
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
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 2),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	event1 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole)),
		startTime:  time.Second,
		isAbsolute: false,
	}

	event2 := &Event{
		note:       note.MustNewNoteWithOctave(note.C, 4).SetDuration(time.Second),
		startTime:  2 * time.Second,
		isAbsolute: true,
	}

	track.AddEvent(event1)
	track.AddEvent(event2)

	end := track.FindEnd()

	assert.Equal(t, 3*time.Second, end, "they should be equal")
}

func TestTrack_GetStartAndEnd(t *testing.T) {
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 2),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	type want struct {
		start, end time.Duration
	}

	testCases := []struct {
		event *Event
		want  *want
	}{
		{
			event: &Event{
				note:       note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole)),
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
				note:       note.MustNewNoteWithOctave(note.C, 4).SetDuration(time.Millisecond * 500),
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
	trackSettings := &Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 2),
		TimeSignature: *fraction.New(4, 4),
	}

	track := NewTrack(trackSettings)

	testCases := []struct {
		event *Event
		want  time.Duration
	}{
		{
			event: &Event{
				note:       note.MustNewNoteWithOctave(note.C, 4).SetValue(duration.NewRelative(duration.NameWhole)),
				startTime:  time.Second,
				isAbsolute: false,
			},
			want: 2 * time.Second,
		},
		{
			event: &Event{
				note:       note.MustNewNoteWithOctave(note.C, 4).SetDuration(time.Millisecond * 500),
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
