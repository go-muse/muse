package chord

import (
	"fmt"
	"time"

	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/note"
)

// Chord is a set of notes of the same duration.
// When a note is added to a chord, it is assigned the duration of the chord.
// When the duration of the chord changes, all notes in the chord are assigned the duration of the chord.
type Chord struct {
	notes       note.Notes
	durationAbs time.Duration
	durationRel *duration.Relative
}

// NewChord creates a new chord with the specified notes.
func NewChord(notes ...*note.Note) *Chord {
	chord := &Chord{}

	return chord.AddNotes(notes...)
}

// NewChordEmpty creates a new empty chord.
func NewChordEmpty() *Chord {
	return &Chord{}
}

// String is stringer for Chord object.
func (c *Chord) String() string {
	if c == nil {
		return "nil chord"
	}

	return fmt.Sprintf("notes: %+v, duration name: %+v, custom duration: %+v", c.notes, c.durationRel.Name(), c.durationAbs)
}

// AddNote adds a note to the chord, replacing it in case of a match.
func (c *Chord) AddNote(n *note.Note) *Chord {
	if c == nil {
		return c
	}

	n.SetDurationAbs(c.durationAbs)
	n.SetDurationRel(c.durationRel)

	for _, chordNote := range c.notes {
		if chordNote.IsEqual(n) {
			return c
		}
	}

	c.notes = append(c.notes, n)

	return c
}

// AddNotes adds notes to the chord, replacing them in case of a match.
func (c *Chord) AddNotes(notes ...*note.Note) *Chord {
	if c == nil {
		return c
	}

	var additionalNotes note.Notes
	for _, note := range notes {
		for _, chordNote := range c.notes {
			if chordNote.IsEqual(note) {
				goto NEXT
			}
		}

		note.SetDurationAbs(c.durationAbs)
		note.SetDurationRel(c.durationRel)
		c.notes = append(c.notes, note)

	NEXT:
	}

	c.notes = append(c.notes, additionalNotes...)

	return c
}

// Notes returns all notes of the chord.
func (c *Chord) Notes() note.Notes {
	if c == nil {
		return nil
	}

	return c.notes
}

// SetDurationAbs sets custom duration to the chord and returns the chord.
func (c *Chord) SetDurationAbs(d time.Duration) *Chord {
	if c == nil {
		return c
	}

	c.durationAbs = d

	for i := range c.notes {
		c.notes[i].SetDurationAbs(c.durationAbs)
	}

	return c
}

// GetDurationAbs returns custom duration of the chord.
func (c *Chord) GetDurationAbs() time.Duration {
	if c == nil {
		return 0
	}

	return c.durationAbs
}

// SetDurationRel sets duration to the chord and returns the chord.
func (c *Chord) SetDurationRel(dr *duration.Relative) *Chord {
	if c == nil {
		return c
	}

	c.durationRel = dr

	for i := range c.notes {
		c.notes[i].SetDurationRel(c.durationRel)
	}

	return c
}

// DurationRel returns duration of the chord.
func (c *Chord) DurationRel() *duration.Relative {
	if c == nil {
		return nil
	}

	return c.durationRel
}

// Empty removes all the notes from the chord.
func (c *Chord) Empty() *Chord {
	if c == nil {
		return nil
	}

	c.notes = nil

	return c
}

// RemoveNote removes a note from the chord that is similar to the specified by it's name and octave.
func (c *Chord) RemoveNote(note *note.Note) *Chord {
	if c == nil {
		return nil
	}

	for i, chordNote := range c.notes {
		if chordNote.IsEqual(note) {
			c.notes = append(c.notes[:i], c.notes[i+1:]...)
		}
	}

	return c
}

// RemoveNotes removes notes from the chord that are similar to the specified by it's name and octave.
func (c *Chord) RemoveNotes(notes note.Notes) *Chord {
	if c == nil {
		return nil
	}

	for _, note := range notes {
		c.RemoveNote(note)
	}

	return c
}

// Exists checks if a note exists in the chord.
func (c *Chord) Exists(note *note.Note) bool {
	if c == nil {
		return false
	}

	for _, chordNote := range c.notes {
		if chordNote.IsEqual(note) {
			return true
		}
	}

	return false
}
