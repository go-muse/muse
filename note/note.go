package note

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/octave"
)

// Note is the representation of a musical sound.
// Each note has a name (i.e., pitch class) and is characterized by octave and duration.
type Note struct {
	name        Name
	octave      *octave.Octave
	durationAbs time.Duration
	durationRel *duration.Relative
}

// New creates new note with a given name and octave number.
func New(noteName Name) (*Note, error) {
	if err := noteName.Validate(); err != nil {
		return nil, err
	}

	return newNote(noteName), nil
}

// NewNoteWithOctave creates new note with a given name and octave number.
func NewNoteWithOctave(noteName Name, octaveNumber octave.Number) (*Note, error) {
	if err := noteName.Validate(); err != nil {
		return nil, err
	}

	oct, err := octave.NewByNumber(octaveNumber)
	if err != nil {
		return nil, fmt.Errorf("create octave with octave number: '%d': %w", octaveNumber, err)
	}

	return newNoteWithOctave(noteName, oct), nil
}

// MustNewNoteWithOctave creates new note with panic in case of invalid note name or octave.
func MustNewNoteWithOctave(noteName Name, octaveNumber octave.Number) *Note {
	note, err := NewNoteWithOctave(noteName, octaveNumber)
	if err != nil {
		panic(err)
	}

	return note
}

// MustNewNote creates new note with panic in case of invalid note name.
func MustNewNote(noteName Name) *Note {
	if err := noteName.Validate(); err != nil {
		panic(err)
	}

	return newNote(noteName)
}

// NewNoteFromString creates a new note from the given string.
func NewNoteFromString(s string) (*Note, error) {
	return New(Name(s))
}

// MustNewNotesFromNoteNames creates a slice of Notes.
func MustNewNotesFromNoteNames(noteNames ...Name) Notes {
	notes := make(Notes, 0, len(noteNames))
	for _, noteName := range noteNames {
		notes = append(notes, MustNewNote(noteName))
	}

	return notes
}

// Name returns name of the note.
func (n *Note) Name() Name {
	if n != nil {
		return n.name
	}

	return ""
}

// Octave returns octave of the note.
func (n *Note) Octave() *octave.Octave {
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

// IsEqualByOctave compares notes by name.
func (n *Note) IsEqualByOctave(note *Note) bool {
	if n == nil || note == nil {
		return false
	}

	return note.Octave().IsEqual(n.octave)
}

// IsEqual compares notes by all parameters.
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

	return &Note{name: n.Name(), octave: n.octave, durationAbs: n.durationAbs, durationRel: n.durationRel}
}

// AlterUp alters the note upwards.
func (n *Note) AlterUp() *Note {
	if n == nil {
		return nil
	}

	if len(n.name) > 1 && strings.HasSuffix(n.name.String(), string(AccidentalFlat)) {
		n.name = n.name[:len(n.name)-len(AccidentalFlat)]

		return n
	}

	n.name = Name(fmt.Sprintf("%s%s", n.name, AccidentalSharp))

	return n
}

// AlterDown alters the note downwards.
func (n *Note) AlterDown() *Note {
	if n == nil {
		return nil
	}

	if len(n.name) > 1 && strings.HasSuffix(n.name.String(), string(AccidentalSharp)) {
		n.name = n.name[:len(n.name)-len(AccidentalSharp)]

		return n
	}

	n.name = Name(fmt.Sprintf("%s%s", n.name, AccidentalFlat))

	return n
}

// AlterUpBy alters the note up by the specified number of times.
func (n *Note) AlterUpBy(i uint8) *Note {
	if n == nil {
		return nil
	}

	for ; i > 0; i-- {
		n.AlterUp()
	}

	return n
}

// AlterDownBy alters the note down by the specified number of times.
func (n *Note) AlterDownBy(i uint8) *Note {
	if n == nil {
		return nil
	}

	for ; i > 0; i-- {
		n.AlterDown()
	}

	return n
}

// BaseName returns note name without accidentals.
func (n *Note) BaseName() Name {
	return n.name[0:1]
}

// GetAlterationShift returns information about alteration of the note (up or down). Sign means direction of alteration.
func (n *Note) GetAlterationShift() int8 {
	var shift int8

	if len(n.name) <= 1 {
		return 0
	}

	if strings.HasSuffix(n.name.String(), string(AccidentalSharp)) {
		for i := len(n.name[1:]); i > 0; i-- {
			shift++
		}

		return shift
	}

	if strings.HasSuffix(n.name.String(), string(AccidentalFlat)) {
		for i := len(n.name[1:]); i > 0; i-- {
			shift--
		}

		return shift
	}

	return shift
}

// SetOctave sets the specified octave to the note and returns the note.
func (n *Note) SetOctave(octave *octave.Octave) *Note {
	n.octave = octave

	return n
}

// SetDurationAbs sets absolute duration to the note and returns the note.
func (n *Note) SetDurationAbs(duration time.Duration) *Note {
	if n == nil {
		return nil
	}

	n.durationAbs = duration

	return n
}

// SetDurationRel sets relative duration to the note and returns the note.
func (n *Note) SetDurationRel(duration *duration.Relative) *Note {
	if n == nil {
		return nil
	}

	n.durationRel = duration

	return n
}

// DurationAbs returns absolute duration of the note.
func (n *Note) DurationAbs() time.Duration {
	if n == nil {
		return 0
	}

	return n.durationAbs
}

// DurationRel returns relative duration of the note.
func (n *Note) DurationRel() *duration.Relative {
	if n == nil {
		return nil
	}

	return n.durationRel
}

// GetTimeDuration calculates and returns time.Duration of the note based on bpm rate, unit and time signature.
func (n *Note) GetTimeDuration(amountOfBars decimal.Decimal) time.Duration {
	if n == nil || n.durationRel == nil {
		return 0
	}

	return n.durationRel.GetTimeDuration(amountOfBars)
}

// GetPartOfBarByRel calculates which part of the bar is occupied by a note with relative duration.
func (n *Note) GetPartOfBarByRel(timeSignature *fraction.Fraction) decimal.Decimal {
	if n == nil || n.durationRel == nil {
		return decimal.Zero
	}

	return n.durationRel.GetPartOfBar(timeSignature)
}

// newNoteWithOctave creates new note with a given name and octave without any restrictions.
func newNoteWithOctave(name Name, octave *octave.Octave) *Note {
	return &Note{name: name, octave: octave}
}

// newNote creates new note with a given name without any restrictions.
func newNote(name Name) *Note {
	return &Note{name: name, octave: nil, durationAbs: 0, durationRel: nil}
}
