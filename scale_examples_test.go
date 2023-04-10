package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// Generate scale from mode.
func ExampleMode_GenerateScale() {
	// If we set mode's name and note from muse, we can be sure that it won't return error.
	mode := muse.MustMakeNewMode(muse.ModeNameLydianDominant, muse.EFLAT)

	scale := mode.GenerateScale(false)

	fmt.Println(scale)
	// Output: [{Eb} {F} {G} {A} {Bb} {C} {Db}]
}
