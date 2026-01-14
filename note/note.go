package note

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/octave"
)

// Note is the representation of a musical sound.
// Each note has a name (i.e., pitch class) and is characterized by octave and duration.
type Note struct {
	name     Name
	octave   *octave.Octave
	duration *time.Duration
	value    *duration.Relative
}

// New constructs a Note from a Name.
func New(name Name) Note {
	return Note{name: name}
}

// Name returns the note's spelled name.
func (n Note) Name() Name {
	return n.name
}

// String returns the note's ASCII spelling.
func (n Note) String() string {
	return n.name.String()
}

// NewFromString creates a new note from the given string.
// The string should contain a note name (e.g., "C", "Db", "F##").
// Returns an error if the string format is invalid.
func NewFromString(s string) (Note, error) {
	name, err := NewNameFromString(s)
	if err != nil {
		return Note{}, err
	}

	return New(name), nil
}

// MustNewFromString creates a new note from the given string and panics on error.
// The string should contain a note name (e.g., "C", "Db", "F##").
// Use this function when you are certain the input is valid.
func MustNewFromString(s string) Note {
	n, err := NewFromString(s)
	if err != nil {
		panic(err)
	}

	return n
}

// NewFromNoteNames creates a slice of Notes.
func NewFromNoteNames(names ...Name) Notes {
	notes := make(Notes, 0, len(names))
	for _, name := range names {
		notes = append(notes, New(name))
	}

	return notes
}

// NewWithOctave creates new note with a given name and octave number.
func NewWithOctave(name Name, octaveNumber octave.Number) (Note, error) {
	oct, err := octave.NewByNumber(octaveNumber)
	if err != nil {
		return Note{}, fmt.Errorf("create octave with octave number: '%d': %w", octaveNumber, err)
	}

	return Note{name: name, octave: oct}, nil
}

// MustNewWithOctave creates new note with panic in case of invalid note name or octave.
func MustNewWithOctave(name Name, octaveNumber octave.Number) Note {
	n, err := NewWithOctave(name, octaveNumber)
	if err != nil {
		panic(err)
	}

	return n
}

// Octave returns octave of the note.
func (n Note) Octave() *octave.Octave {
	return n.octave
}

// EqualByName compares notes by name.
func (n Note) EqualByName(other Note) bool {
	return n.name.EqualSpelling(other.name)
}

// EqualByOctave compares notes by octave.
// If both octaves are nil, they are considered equal.
func (n Note) EqualByOctave(other Note) bool {
	if n.octave == nil && other.octave == nil {
		return true
	}

	if n.octave == nil || other.octave == nil {
		return false
	}

	return n.octave.IsEqual(other.octave)
}

// Equal compares notes by all parameters.
func (n Note) Equal(other Note) bool {
	return n.EqualByName(other) && n.EqualByOctave(other)
}

// Copy creates a deep copy of Note with respect to direct pointer fields.
// It clones octave, duration and value objects (if present) by value-copying the pointed structs.
func (n Note) Copy() Note {
	out := n

	if n.octave != nil {
		o := *n.octave
		out.octave = &o
	}

	if n.duration != nil {
		d := *n.duration
		out.duration = &d
	}

	if n.value != nil {
		v := *n.value
		out.value = &v
	}

	return out
}

// AlterUp alters the note upwards.
func (n Note) AlterUp() Note {
	n.name = n.name.AlterUp()
	return n
}

// AlterDown alters the note downwards.
func (n Note) AlterDown() Note {
	n.name = n.name.AlterDown()
	return n
}

// AlterUpBy alters the note up by the specified number of times.
func (n Note) AlterUpBy(i uint8) Note {
	n.name = n.name.AlterUpBy(i)
	return n
}

// AlterDownBy alters the note down by the specified number of times.
func (n Note) AlterDownBy(i uint8) Note {
	n.name = n.name.AlterDownBy(i)
	return n
}

// BaseName returns note name without accidentals.
func (n Note) BaseName() string {
	return n.name.BaseName()
}

// AlterationShift returns information about alteration of the note (up or down). Sign means direction of alteration.
func (n Note) AlterationShift() int8 {
	return n.name.AlterationShift()
}

// SetOctave sets the specified octave to the note and returns the note.
func (n Note) SetOctave(octave *octave.Octave) Note {
	n.octave = octave
	return n
}

// SetDuration sets absolute duration to the note and returns the note.
func (n Note) SetDuration(d time.Duration) Note {
	n.duration = &d
	return n
}

// SetValue sets relative duration to the note and returns the note.
func (n Note) SetValue(v *duration.Relative) Note {
	n.value = v
	return n
}

// Duration returns absolute duration of the note.
// Returns 0 if duration is not set.
func (n Note) Duration() time.Duration {
	if n.duration == nil {
		return 0
	}
	return *n.duration
}

// Value returns relative duration of the note.
func (n Note) Value() *duration.Relative {
	return n.value
}

// GetTimeDuration calculates and returns time.Duration of the note based on bpm rate, unit and time signature.
func (n Note) GetTimeDuration(amountOfBars decimal.Decimal) time.Duration {
	return n.value.GetTimeDuration(amountOfBars)
}

// GetPartOfBarByValue calculates which part of the bar is occupied by a note with its value (relative duration).
func (n Note) GetPartOfBarByValue(timeSignature *fraction.Fraction) decimal.Decimal {
	return n.value.GetPartOfBar(timeSignature)
}
