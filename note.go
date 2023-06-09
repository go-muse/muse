package muse

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// Note is the representation of a musical sound.
// Each note has a name (i.e., pitch class) and is characterized by octave and duration.
type Note struct {
	name     NoteName
	octave   *Octave
	duration *Duration
}

// Notes is slice of Note.
type Notes []Note

// String is stringer for Note.
func (ns Notes) String() string {
	noteNames := make([]NoteName, len(ns))
	for i, note := range ns {
		noteNames[i] = note.Name()
	}

	return fmt.Sprintf("%v", noteNames)
}

// newNote creates new note with a given name without any restrictions.
func newNote(name NoteName) *Note {
	return &Note{name: name, octave: nil, duration: nil}
}

// NewNote creates new note with a given name and octave number.
func NewNote(noteName NoteName) (*Note, error) {
	if err := noteName.Validate(); err != nil {
		return nil, err
	}

	return newNote(noteName), nil
}

// newNote creates new note with a given name and octave without any restrictions.
func newNoteWithOctave(name NoteName, octave *Octave) *Note {
	return &Note{name: name, octave: octave}
}

// NewNoteWithOctave creates new note with a given name and octave number.
func NewNoteWithOctave(noteName NoteName, octaveNumber OctaveNumber) (*Note, error) {
	if err := noteName.Validate(); err != nil {
		return nil, err
	}

	octave, err := NewOctave(octaveNumber)
	if err != nil {
		return nil, err
	}

	return newNoteWithOctave(noteName, octave), nil
}

// MustNewNoteWithOctave creates new note with panic in case of invalid note name or octave.
func MustNewNoteWithOctave(noteName NoteName, octaveNumber OctaveNumber) *Note {
	note, err := NewNoteWithOctave(noteName, octaveNumber)
	if err != nil {
		panic(err)
	}

	return note
}

// MustNewNote creates new note with panic in case of invalid note name.
func MustNewNote(noteName NoteName) *Note {
	if err := noteName.Validate(); err != nil {
		panic(err)
	}

	return newNote(noteName)
}

// NewNoteFromString creates a new note from the given string.
func NewNoteFromString(s string) (*Note, error) {
	return NewNote(NoteName(s))
}

// Name returns name of the note.
func (n *Note) Name() NoteName {
	if n != nil {
		return n.name
	}

	return ""
}

// Name returns octave of the note.
func (n *Note) Octave() *Octave {
	if n != nil {
		return n.octave
	}

	return nil
}

// IsEqualByName compares notes by name.
func (n *Note) IsEqualByName(note *Note) bool {
	if n == nil || note == nil {
		return false
	}

	return note.Name() == n.Name()
}

// IsEqualByName compares notes by name.
func (n *Note) IsEqualByOctave(note *Note) bool {
	if n == nil || note == nil {
		return false
	}

	return note.Octave().IsEqual(n.octave)
}

// IsEqualByName compares notes by all parameters.
func (n *Note) IsEqual(note *Note) bool {
	if n == nil || note == nil {
		return false
	}

	if !note.IsEqualByName(n) || !note.IsEqualByOctave(n) {
		return false
	}

	return true
}

// Copy creates full copy of the current Note.
// The method returns a pointer to the new Note containing the same attribute values
// as the original Note that the function was called on.
func (n *Note) Copy() *Note {
	if n == nil {
		return nil
	}

	return &Note{name: n.Name(), octave: n.octave, duration: n.duration}
}

// AlterUp alterates the note upwards.
func (n *Note) AlterUp() *Note {
	if n == nil {
		return nil
	}

	if len(n.name) > 1 && strings.HasSuffix(n.name.String(), string(AlterSymbolFlat)) {
		n.name = n.name[:len(n.name)-len(AlterSymbolFlat)]

		return n
	}

	n.name = NoteName(fmt.Sprintf("%s%s", n.name, AlterSymbolSharp))

	return n
}

// AlterDown alterates the note downwards.
func (n *Note) AlterDown() *Note {
	if n == nil {
		return nil
	}

	if len(n.name) > 1 && strings.HasSuffix(n.name.String(), string(AlterSymbolSharp)) {
		n.name = n.name[:len(n.name)-len(AlterSymbolSharp)]

		return n
	}

	n.name = NoteName(fmt.Sprintf("%s%s", n.name, AlterSymbolFlat))

	return n
}

// AlterDown alterates the note up by the specified number of times.
func (n *Note) AlterUpBy(i uint8) *Note {
	if n == nil {
		return nil
	}

	for ; i > 0; i-- {
		n.AlterUp()
	}

	return n
}

// AlterDown alterates the note down by the specified number of times.
func (n *Note) AlterDownBy(i uint8) *Note {
	if n == nil {
		return nil
	}

	for ; i > 0; i-- {
		n.AlterDown()
	}

	return n
}

// BaseName returns note name without alteration symbols.
func (n *Note) BaseName() NoteName {
	return n.name[0:1]
}

// GetAlterationShift returns information about alteration of the note (up or down). Sign means direction of alteration.
func (n *Note) GetAlterationShift() int8 {
	var shift int8

	if len(n.name) <= 1 {
		return 0
	}

	if strings.HasSuffix(n.name.String(), string(AlterSymbolSharp)) {
		for i := len(n.name[1:]); i > 0; i-- {
			shift++
		}

		return shift
	}

	if strings.HasSuffix(n.name.String(), string(AlterSymbolFlat)) {
		for i := len(n.name[1:]); i > 0; i-- {
			shift--
		}

		return shift
	}

	return shift
}

func (n *Note) baseNote() *Note {
	return newNoteWithOctave(n.name[0:1], n.octave)
}

// SetOctave sets the specified octave to the note and returns the note.
func (n *Note) SetOctave(octave *Octave) *Note {
	n.octave = octave

	return n
}

// SetDuration sets duration to the note and returns the note.
func (n *Note) SetDuration(duration Duration) *Note {
	if n == nil {
		return nil
	}

	n.duration = &duration

	return n
}

// Duration returns duration of the note.
func (n *Note) Duration() *Duration {
	if n == nil {
		return nil
	}

	return n.duration
}

// TimeDuration returns time.Duration of the note based on bpm rate, unit and time signature.
func (n *Note) TimeDuration(trackSettings TrackSettings) time.Duration {
	if n == nil || n.duration == nil {
		return 0
	}

	return n.duration.GetTimeDuration(trackSettings)
}

// SetAbsoluteDuration sets custom duration to the note and returns the note.
func (n *Note) SetAbsoluteDuration(d time.Duration) *Note {
	if n == nil {
		return n
	}

	if n.duration == nil {
		n.duration = &Duration{
			absoluteDuration: 0,
			relativeDuration: relativeDuration{},
		}
	}

	n.duration.absoluteDuration = d

	return n
}

// GetAbsoluteDuration returns custom duration of the note.
func (n *Note) GetAbsoluteDuration() time.Duration {
	if n == nil || n.duration == nil {
		return 0
	}

	return n.duration.absoluteDuration
}

// GetPartOfBarByRelative returns duration value in decimal.
func (n *Note) GetPartOfBarByRelative(timeSignature *Fraction) decimal.Decimal {
	if n == nil || n.duration == nil {
		return decimal.Zero
	}

	return n.duration.GetPartOfBarByRelative(timeSignature)
}

// GetPartOfBarByAbsolute returns duration value in decimal.
func (n *Note) GetPartOfBarByAbsolute(trackSettings TrackSettings) decimal.Decimal {
	if n == nil || n.duration == nil {
		return decimal.Zero
	}

	return n.duration.GetPartOfBarByAbsolute(trackSettings)
}
