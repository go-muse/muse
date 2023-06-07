package muse

import (
	"fmt"
	"time"
)

// Chord is a set of notes of the same duration.
// When a note is added to a chord, it is assigned the duration of the chord.
// When the duration of the chord changes, all notes in the chord are assigned the duration of the chord.
type Chord struct {
	notes    Notes
	duration *Duration
}

// NewChord creates a new chord with the specified notes and duration, assigning the specified duration to all notes.
func NewChord(duration Duration, notes ...Note) *Chord {
	chord := &Chord{duration: &duration}

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

	return fmt.Sprintf("notes: %+v, duration name: %+v, custom duration: %+v", c.notes, c.duration.Name(), c.duration.absoluteDuration)
}

// AddNote adds a note to the chord, replacing it in case of a match.
func (c *Chord) AddNote(n Note) *Chord {
	if c == nil {
		return c
	}

	n.duration = c.duration

	for _, chordNote := range c.notes {
		if chordNote.IsEqual(&n) {
			return c
		}
	}

	c.notes = append(c.notes, n)

	return c
}

// AddNote adds notes to the chord, replacing them in case of a match.
func (c *Chord) AddNotes(notes ...Note) *Chord {
	if c == nil {
		return c
	}

	var additionalNotes Notes
	for _, note := range notes {
		for _, chordNote := range c.notes {
			if chordNote.IsEqual(&note) {
				goto NEXT
			}
		}

		note.duration = c.duration
		c.notes = append(c.notes, note)

	NEXT:
	}

	c.notes = append(c.notes, additionalNotes...)

	return c
}

// GetNotes returns all notes of the chord.
func (c *Chord) GetNotes() Notes {
	if c == nil {
		return nil
	}

	return c.notes
}

// SetAbsoluteDuration sets custom duration to the chord and returns the chord.
func (c *Chord) SetAbsoluteDuration(d time.Duration) *Chord {
	if c == nil {
		return c
	}

	if c.duration == nil {
		c.duration = &Duration{absoluteDuration: 0}
	}

	c.duration.absoluteDuration = d

	for i := range c.notes {
		if c.notes[i].duration == nil {
			c.notes[i].duration = c.duration

			continue
		}

		c.notes[i].duration.absoluteDuration = d
	}

	return c
}

// GetAbsoluteDuration returns custom duration of the chord.
func (c *Chord) GetAbsoluteDuration() time.Duration {
	if c == nil || c.duration == nil {
		return 0
	}

	return c.duration.absoluteDuration
}

// SetDuration sets duration to the chord and returns the chord.
func (c *Chord) SetDuration(duration Duration) *Chord {
	if c == nil {
		return c
	}

	c.duration = &duration

	for i := range c.notes {
		c.notes[i].duration = &duration
	}

	return c
}

// GetDuration returns duration of the chord.
func (c *Chord) GetDuration() *Duration {
	if c == nil {
		return nil
	}

	return c.duration
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
func (c *Chord) RemoveNote(note Note) *Chord {
	if c == nil {
		return nil
	}

	for i, chordNote := range c.notes {
		if chordNote.IsEqual(&note) {
			c.notes = append(c.notes[:i], c.notes[i+1:]...)
		}
	}

	return c
}

// RemoveNotes removes notes from the chord that are similar to the specified by it's name and octave.
func (c *Chord) RemoveNotes(notes Notes) *Chord {
	if c == nil {
		return nil
	}

	for _, note := range notes {
		for i, chordNote := range c.notes {
			if chordNote.IsEqual(&note) {
				c.notes = append(c.notes[:i], c.notes[i+1:]...)
				break
			}
		}
	}

	return c
}

// Exists checks if a note exists in the chord.
func (c *Chord) Exists(note Note) bool {
	if c == nil {
		return false
	}

	for _, chordNote := range c.notes {
		if chordNote.IsEqual(&note) {
			return true
		}
	}

	return false
}
