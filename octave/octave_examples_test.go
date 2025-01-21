package octave_test

import (
	"fmt"

	"github.com/go-muse/muse/octave"
)

// Creating an octave.
func ExampleOctave() {
	// create octave #3
	octave, err := octave.NewByNumber(octave.Number3)
	if err != nil {
		panic(err)
	}

	fmt.Println(octave.Number(), octave.Name())
	// Output: 3 SmallOctave
}
