package duration_test

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/octave"
)

// Creating a new relative duration and assigning it to the note.
func ExampleNewRelative() {
	// half note duration
	dur := duration.NewRelative(duration.NameHalf)

	// creating note and setting duration
	note := note.MustNewNoteWithOctave(note.C, octave.Number3).SetValue(dur)

	fmt.Println(note.Value().Name())
	// Output: Half
}

// Getting time.Duration from duration.
func ExampleRelative_GetTimeDuration() {
	// musical settings
	// you can use track settings and get amount of bars from there
	amountOfBars := decimal.NewFromUint64(60)

	// half note duration
	dur := duration.NewRelative(duration.NameHalf)

	fmt.Println(dur.GetTimeDuration(amountOfBars))
	// Output: 500ms
}
