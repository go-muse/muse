package builder

import (
	"errors"

	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

var errInvalidFirstTemplateNote = errors.New("invalid first template note")

// NewBuilderCommon builds notes for the mode.
func NewBuilderCommon(modeTemplate HalftonesIterator, firstNote *note.Note) Builder {
	send := func(n *note.Note, halfTones halftone.HalfTones) func() (*note.Note, halftone.HalfTones) {
		return func() (*note.Note, halftone.HalfTones) { return n, halfTones }
	}

	f := func(c chan func() (*note.Note, halftone.HalfTones)) {
		// Get instance with 12 template notes
		templateNotes := getTemplateNotesCommon()

		// Get the first template note based on the first note of the mode
		templateNote := templateNotes.getTemplateNote(firstNote)
		if templateNote == nil {
			panic(errInvalidFirstTemplateNote)
		}

		// Save the first note of the mode into the template note
		templateNote.saveResultingNote(firstNote)

		// Iterate through the mode template
		for iteratorResult := range modeTemplate.Iterate() {
			halfTones, halfTonesFromPrime := iteratorResult()

			// To avoid duplicating root notes for one-note modes
			if halfTones == halftone.HalfTonesInOctave {
				break
			}

			// Get next template note based on mode template's step
			nextTemplateNote := templateNote.getByHalftones(halfTones)

			// Attach the resulting template note to the current one
			templateNote.setNextTemplateNote(nextTemplateNote)

			// Build a new note based on the new template note and the previous one
			newNote := nextTemplateNote.buildNoteByPrevious()

			// Save new note into the next template note
			nextTemplateNote.saveResultingNote(newNote)

			// Insert the next note into the current note variable, to use it at the next iteration
			templateNote = nextTemplateNote

			// send resulting note with distance from the prime note in halftone
			c <- send(newNote, halfTonesFromPrime)
		}

		close(c)
	}

	c := make(chan func() (*note.Note, halftone.HalfTones))
	go f(c)

	return c
}

// templateNoteCommon is a template note for the single template instance.
type templateNoteCommon struct {
	next          *templateNoteCommon
	previous      *templateNoteCommon
	isAltered     bool
	notAltered    *note.Note
	alteredNotes  []*noteRelation
	resultingNote *noteRelation
}

// templateNotesCommon is the instance with template notes.
type templateNotesCommon struct {
	*templateNoteCommon           // the first note of the linked list in the template instance
	baseNote            *baseNote // last used note without accidental to build the mode
}

// equal Compares Note and template note by name.
func (tn *templateNoteCommon) equal(note *note.Note) bool {
	if !tn.isAltered {
		if tn.notAltered.IsEqualByName(note) {
			return true
		}
	}

	if tn.alteredNotes != nil {
		for _, ar := range tn.alteredNotes {
			if ar.realNote().IsEqualByName(note) {
				return true
			}
		}
	}

	return false
}

// getNext() returns the next template note.
func (tn *templateNoteCommon) getNext() *templateNoteCommon {
	return tn.next
}

// getTemplateNote determines template note by the given Note.
// When building a mode from a specific note, we will need to move the pointer to the first note
// and from it iterate further and pass it to the final notes builder, which decides on their naming.
func (ti *templateNotesCommon) getTemplateNote(note *note.Note) *templateNoteCommon {
	if ti.templateNoteCommon == nil {
		return nil
	}

	currentTemplateNote := ti.templateNoteCommon
	if currentTemplateNote.equal(note) {
		return currentTemplateNote
	}

	// the condition that it will not endlessly drive in a circle in search of a non-existent note
	for i := 2; i <= int(halftone.HalfTonesInOctave); i++ {
		currentTemplateNote = currentTemplateNote.getNext()
		if currentTemplateNote.equal(note) {
			return currentTemplateNote
		}
	}

	return nil
}

// getByHalftones returns a template note that is several half steps away from the current one.
func (tn *templateNoteCommon) getByHalftones(halfTones halftone.HalfTones) *templateNoteCommon {
	templateNote := tn
	for range halfTones {
		templateNote = templateNote.getNext()
	}

	return templateNote
}

// saveResultingNote inserts the final note into the template note
// with preserving the information about which note is the base for the saved note.
func (tn *templateNoteCommon) saveResultingNote(note *note.Note) {
	if tn.isAltered {
		for _, ar := range tn.alteredNotes {
			if ar.realNote().IsEqualByName(note) {
				tn.resultingNote = &noteRelation{ar.baseNote(), note}
			}
		}
	} else {
		tn.resultingNote = &noteRelation{note, note}
	}
}

// setNextTemplateNote function is necessary to obtain a template chain of the same length as the mode,
// in order to correctly compare the current note with the immediately preceding one,
// without guessing how many skips should be made, from a 12-note octave chain of a template set.
func (tn *templateNoteCommon) setNextTemplateNote(next *templateNoteCommon) {
	tn.next = next
	next.previous = tn
}

// buildNoteByPrevious determines the final note based on a template note using the previous template note.
func (tn *templateNoteCommon) buildNoteByPrevious() *note.Note {
	if tn == nil {
		return nil
	}

	// If it is not an alterable note (C, D, E, F, G, A, B) - we return the note itself.
	if !tn.isAltered {
		// This code will not work for Fb and Cb.
		return tn.notAltered

		// If it is an alterable note, then we compare it with the previous one.
	} else { //nolint
		for _, ar := range tn.alteredNotes {
			if !ar.baseNote().IsEqualByName(tn.previous.resultingNote.baseNote()) {
				return ar.realNote()
			}
		}
	}

	return nil
}

// noteRelation is relation between note, and it's base note.
// It is used sto get base note by altered note or
// to get base note from resulting note.
type noteRelation struct {
	base *note.Note
	real *note.Note
}

// baseNote returns the base note from real note and base note relation.
func (ar *noteRelation) baseNote() *note.Note {
	return ar.base
}

// realNote returns the real note from real note and base note relation.
func (ar *noteRelation) realNote() *note.Note {
	return ar.real
}

// baseNote is one of "clean" notes: [C, D, E, F, G, A, B].
// It is involved in calculating the alteration of the constructed notes of the mode.
type baseNote struct {
	prev, next *baseNote
	note       *note.Note
}

// getTemplateNotesCommon creates 12 template notes for the instance when it is created
// this will be hardcoded as to what each template note contains.
func getTemplateNotesCommon() *templateNotesCommon {
	templateNote12 := &templateNoteCommon{
		next:         nil,
		isAltered:    false,
		notAltered:   note.B.MustNewNote(),
		alteredNotes: []*noteRelation{{note.C.MustNewNote(), note.CFLAT.MustNewNote()}},
	}

	templateNote11 := &templateNoteCommon{
		next:         templateNote12,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{note.A.MustNewNote(), note.ASHARP.MustNewNote()}, {note.B.MustNewNote(), note.BFLAT.MustNewNote()}},
	}

	templateNote10 := &templateNoteCommon{
		next:         templateNote11,
		isAltered:    false,
		notAltered:   note.A.MustNewNote(),
		alteredNotes: nil,
	}

	templateNote9 := &templateNoteCommon{
		next:         templateNote10,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{note.G.MustNewNote(), note.GSHARP.MustNewNote()}, {note.A.MustNewNote(), note.AFLAT.MustNewNote()}},
	}

	templateNote8 := &templateNoteCommon{
		next:         templateNote9,
		isAltered:    false,
		notAltered:   note.G.MustNewNote(),
		alteredNotes: nil,
	}

	templateNote7 := &templateNoteCommon{
		next:         templateNote8,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{note.F.MustNewNote(), note.FSHARP.MustNewNote()}, {note.G.MustNewNote(), note.GFLAT.MustNewNote()}},
	}

	templateNote6 := &templateNoteCommon{
		next:         templateNote7,
		isAltered:    false,
		notAltered:   note.F.MustNewNote(),
		alteredNotes: []*noteRelation{{note.E.MustNewNote(), note.ESHARP.MustNewNote()}},
	}

	templateNote5 := &templateNoteCommon{
		next:         templateNote6,
		isAltered:    false,
		notAltered:   note.E.MustNewNote(),
		alteredNotes: []*noteRelation{{note.F.MustNewNote(), note.FFLAT.MustNewNote()}},
	}

	templateNote4 := &templateNoteCommon{
		next:         templateNote5,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{note.D.MustNewNote(), note.DSHARP.MustNewNote()}, {note.E.MustNewNote(), note.EFLAT.MustNewNote()}},
	}

	templateNote3 := &templateNoteCommon{
		next:         templateNote4,
		isAltered:    false,
		notAltered:   note.D.MustNewNote(),
		alteredNotes: nil,
	}

	templateNote2 := &templateNoteCommon{
		next:         templateNote3,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{note.C.MustNewNote(), note.CSHARP.MustNewNote()}, {note.D.MustNewNote(), note.DFLAT.MustNewNote()}},
	}

	templateNote1 := &templateNoteCommon{
		next:         templateNote2,
		isAltered:    false,
		notAltered:   note.C.MustNewNote(),
		alteredNotes: []*noteRelation{{note.B.MustNewNote(), note.BSHARP.MustNewNote()}},
	}

	// set links to previous template notes
	templateNote2.previous = templateNote1
	templateNote3.previous = templateNote2
	templateNote4.previous = templateNote3
	templateNote5.previous = templateNote4
	templateNote6.previous = templateNote5
	templateNote7.previous = templateNote6
	templateNote8.previous = templateNote7
	templateNote9.previous = templateNote8
	templateNote10.previous = templateNote9
	templateNote11.previous = templateNote10
	templateNote12.previous = templateNote11

	// There must be cycling,
	// this will allow iteration within an octave from any note to any number of steps
	// for example, from E to E
	templateNote12.next = templateNote1
	templateNote1.previous = templateNote12

	baseNote7 := &baseNote{
		prev: nil,
		next: nil,
		note: note.MustNewNote(note.B),
	}
	baseNote6 := &baseNote{
		prev: nil,
		next: baseNote7,
		note: note.MustNewNote(note.A),
	}
	baseNote5 := &baseNote{
		prev: nil,
		next: baseNote6,
		note: note.MustNewNote(note.G),
	}
	baseNote4 := &baseNote{
		prev: nil,
		next: baseNote5,
		note: note.MustNewNote(note.F),
	}
	baseNote3 := &baseNote{
		prev: nil,
		next: baseNote4,
		note: note.MustNewNote(note.E),
	}
	baseNote2 := &baseNote{
		prev: nil,
		next: baseNote3,
		note: note.MustNewNote(note.D),
	}
	baseNote1 := &baseNote{
		prev: nil,
		next: baseNote2,
		note: note.MustNewNote(note.C),
	}

	baseNote7.next = baseNote1

	return &templateNotesCommon{templateNote1, baseNote1}
}

func newBaseNote(n *note.Note) *note.Note {
	return note.MustNewNote(n.Name()[0:1]).SetOctave(n.Octave())
}
