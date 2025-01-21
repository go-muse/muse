package builder

import (
	"github.com/go-muse/muse/halftone"
)

// HalftonesIterator produces a function that provides two amounts of halftones: the distance from the previous note and from the first.
// The iterator is expected to produce a finite amount of values (no cyclic generation) and implies forward generation.
type HalftonesIterator interface {
	Iterate() <-chan func() (halftone.HalfTones, halftone.HalfTones)
}
