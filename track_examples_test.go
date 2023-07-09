package muse_test

import (
	"fmt"
	"time"

	"github.com/go-muse/muse"
)

// Creating track and adding notes.
func ExampleTrack_AddNote() {
	// track settings
	trackSettings := muse.TrackSettings{
		BPM:           uint64(120),
		Unit:          muse.Fraction{1, 2},
		TimeSignature: muse.Fraction{4, 4},
	}

	// creating a track
	track := muse.NewTrack(trackSettings)

	noteC := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber4).SetAbsoluteDuration(time.Second)
	track.AddNote(
		noteC,
		0,
		true,
	)

	noteD := muse.MustNewNoteWithOctave(muse.C, muse.OctaveNumber4).SetDuration(*muse.NewDurationWithRelativeValue(muse.DurationNameHalf))
	track.AddNote(
		noteD,
		noteC.GetAbsoluteDuration(),
		false,
	)

	fmt.Println(track.Events())
	// Output: [start time: 0s, note: C, is absolute: true start time: 1s, note: C, is absolute: false]
}
