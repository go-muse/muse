package builder

import (
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

// Builder is entity that encapsulates functionality to generate sequence of notes and halftones.
type Builder <-chan func() (*note.Note, halftone.HalfTones)
