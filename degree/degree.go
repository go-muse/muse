package degree

import (
	"fmt"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

// Number is a position of a note in a mode, from 1 to last.
// This is a musical concept representing the scale degree number (I, II, III, IV, V, VI, VII, etc.)
type Number uint8

// Degree is a musical entity representing a scale degree with its characteristics.
// It contains only musical information, without any list/chain logic.
type Degree struct {
	number                Number
	halfTonesFromPrime    halftone.HalfTones
	note                  note.Note
	modalCharacteristics  ModalCharacteristics
	absoluteModalPosition ModalPosition
}

// NewDegree creates a new degree with specified parameters.
func NewDegree(
	number Number,
	halfTonesFromPrime halftone.HalfTones,
	n note.Note,
	modalCharacteristics ModalCharacteristics,
	absoluteModalPosition ModalPosition,
) Degree {
	return Degree{
		number:                number,
		halfTonesFromPrime:    halfTonesFromPrime,
		note:                  n,
		modalCharacteristics:  modalCharacteristics,
		absoluteModalPosition: absoluteModalPosition,
	}
}

// Number returns degree's number (its position in a mode).
func (d Degree) Number() Number {
	return d.number
}

// SetNumber sets degree's number.
func (d *Degree) SetNumber(number Number) {
	d.number = number
}

// HalfTonesFromPrime returns degree's distance from prime in halftones.
func (d Degree) HalfTonesFromPrime() halftone.HalfTones {
	return d.halfTonesFromPrime
}

// SetHalfTonesFromPrime sets halftone from prime value in halftones.
func (d *Degree) SetHalfTonesFromPrime(halfTones halftone.HalfTones) {
	d.halfTonesFromPrime = halfTones
}

// Note returns the note lying on this degree.
func (d Degree) Note() note.Note {
	return d.note
}

// SetNote sets note for the degree.
func (d *Degree) SetNote(n note.Note) {
	d.note = n
}

// ModalCharacteristics returns modal characteristics of the degree.
func (d Degree) ModalCharacteristics() ModalCharacteristics {
	return d.modalCharacteristics
}

// SetModalCharacteristics sets modal characteristics to the degree.
func (d *Degree) SetModalCharacteristics(modalCharacteristics ModalCharacteristics) {
	d.modalCharacteristics = modalCharacteristics
}

// AbsoluteModalPosition returns absolute modal position of the degree.
func (d Degree) AbsoluteModalPosition() ModalPosition {
	return d.absoluteModalPosition
}

// SetAbsoluteModalPosition sets absolute modal position to the degree.
func (d *Degree) SetAbsoluteModalPosition(modalPosition ModalPosition) {
	d.absoluteModalPosition = modalPosition
}

// Copy creates a copy of the degree.
func (d Degree) Copy() Degree {
	return Degree{
		number:                d.number,
		halfTonesFromPrime:    d.halfTonesFromPrime,
		note:                  d.note.Copy(),
		modalCharacteristics:  d.modalCharacteristics.Copy(),
		absoluteModalPosition: d.absoluteModalPosition,
	}
}

// String is stringer for degree object.

// String returns a string representation of the degree.
func (d Degree) String() string {
	if !d.absoluteModalPosition.IsSet() {
		return fmt.Sprintf("Number: %d, HalfTonesFromPrime: %d, note: %s",
			d.number, d.halfTonesFromPrime, d.note.Name())
	}

	return fmt.Sprintf("Number: %d, HalfTonesFromPrime: %d, note: %s, absolute modal position: %s (Weight:%d)",
		d.number, d.halfTonesFromPrime, d.note.Name(), d.absoluteModalPosition.name, d.absoluteModalPosition.Weight())
}
