package track_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
	"github.com/go-muse/muse/track"
)

// Creating track and adding notes.
func ExampleTrack_AddNote() {
	// track settings
	trackSettings := &track.Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 2),
		TimeSignature: *fraction.New(4, 4),
	}

	// creating a track
	track := track.NewTrack(trackSettings)

	noteC := note.MustNewNoteWithOctave(note.C, octave.Number4).SetDuration(time.Second)
	track.AddNote(
		noteC,
		0,
		true,
	)

	noteD := note.MustNewNoteWithOctave(note.C, octave.Number4).SetValue(duration.NewRelative(duration.NameHalf))
	track.AddNote(
		noteD,
		noteC.Duration(),
		false,
	)

	fmt.Println(track.Events())
	// Output: [start time: 0s, note: C, is absolute: true start time: 1s, note: C, is absolute: false]
}

// Knowing the BPM, unit and the time signature(meter), calculate the number of bars in a nanosecondsInMinute.
func ExampleSettings_GetAmountOfBars() {
	trackSettings := track.Settings{
		BPM:           uint64(120),
		Unit:          *fraction.New(1, 4),
		TimeSignature: *fraction.New(4, 4),
	}

	amountOfBars := trackSettings.GetAmountOfBars()

	fmt.Println(amountOfBars)
	// Output: 30
}
