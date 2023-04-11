package muse

// noteRelation is relation between note and it's base note.
// It is used sto get base note by altered note or
// to get base note from resultin note.
type noteRelation struct {
	base *Note
	real *Note
}

// baseNote returns the base note from real note and base note relation.
func (ar *noteRelation) baseNote() *Note {
	return ar.base
}

// realNote returns the real note from real note and base note relation.
func (ar *noteRelation) realNote() *Note {
	return ar.real
}

// templateNote is a template note for the single template instance.
type templateNote struct {
	next          *templateNote
	previous      *templateNote
	isAltered     bool
	notAltered    *Note
	alteredNotes  []*noteRelation
	resultingNote *noteRelation
}

func (tn *templateNote) getNext() *templateNote {
	return tn.next
}

// saveResultingNote inserts the final note into the template note
// with preserving the information about which note is the base for the saved note.
func (tn *templateNote) saveResultingNote(note *Note) {
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
func (tn *templateNote) setNextTemplateNote(next *templateNote) {
	tn.next = next
	next.previous = tn
}

// getByHalftones returns a template note that is several half steps away from the current one.
func (tn *templateNote) getByHalftones(halfTones HalfTones) *templateNote {
	templateNote := tn
	for i := 0; i < int(halfTones); i++ {
		templateNote = templateNote.getNext()
	}

	return templateNote
}

// buildNoteByPrevious determines the final note based on a template note using the previous template note.
func (tn *templateNote) buildNoteByPrevious() *Note {
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

// templateInstance is the instance with template notes.
type templateInstance struct {
	*templateNote // the first note of the linked list in the template instance
}

// getTemplateNote determines template note by the given Note.
// When building a mode from a specific note, we will need to move the pointer to the first note
// and from it iterate further and pass it to the final notes builder, which decides on their naming.
func (ti *templateInstance) getTemplateNote(note *Note) *templateNote {
	if ti.templateNote == nil {
		return nil
	}

	currentTemplateNote := ti.templateNote
	if currentTemplateNote.Equal(note) {
		return currentTemplateNote
	}

	// the condition that it will not endlessly drive in a circle in search of a non-existent note
	for i := 2; i <= int(HalftonesInOctave); i++ {
		currentTemplateNote = currentTemplateNote.getNext()
		if currentTemplateNote.Equal(note) {
			return currentTemplateNote
		}
	}

	return nil
}

// Equal Compares Note and template note by name.
func (tn *templateNote) Equal(note *Note) bool {
	if tn.isAltered {
		for _, ar := range tn.alteredNotes {
			if ar.realNote().IsEqualByName(note) {
				return true
			}
		}
	} else { //nolint
		if tn.notAltered.IsEqualByName(note) {
			return true
		}
	}

	return false
}

// getTemplateNotesInstance creates 12 template notes for the instance when it is created
// this will be hardcoded as to what each template note contains.
func getTemplateNotesInstance() *templateInstance {
	templateNote12 := &templateNote{
		next:         nil,
		isAltered:    false,
		notAltered:   &Note{name: B},
		alteredNotes: nil,
	}

	templateNote11 := &templateNote{
		next:         templateNote12,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: A}, &Note{name: ASHARP}}, {&Note{name: B}, &Note{name: BFLAT}}},
	}

	templateNote10 := &templateNote{
		next:         templateNote11,
		isAltered:    false,
		notAltered:   &Note{name: A},
		alteredNotes: nil,
	}

	templateNote9 := &templateNote{
		next:         templateNote10,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: G}, &Note{name: GSHARP}}, {&Note{name: A}, &Note{name: AFLAT}}},
	}

	templateNote8 := &templateNote{
		next:         templateNote9,
		isAltered:    false,
		notAltered:   &Note{name: G},
		alteredNotes: nil,
	}

	templateNote7 := &templateNote{
		next:         templateNote8,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: F}, &Note{name: FSHARP}}, {&Note{name: G}, &Note{name: GFLAT}}},
	}

	templateNote6 := &templateNote{
		next:         templateNote7,
		isAltered:    false,
		notAltered:   &Note{name: F},
		alteredNotes: nil,
	}

	templateNote5 := &templateNote{
		next:         templateNote6,
		isAltered:    false,
		notAltered:   &Note{name: E},
		alteredNotes: nil,
	}

	templateNote4 := &templateNote{
		next:         templateNote5,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: D}, &Note{name: DSHARP}}, {&Note{name: E}, &Note{name: EFLAT}}},
	}

	templateNote3 := &templateNote{
		next:         templateNote4,
		isAltered:    false,
		notAltered:   &Note{name: D},
		alteredNotes: nil,
	}

	templateNote2 := &templateNote{
		next:         templateNote3,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: C}, &Note{name: CSHARP}}, {&Note{name: D}, &Note{name: DFLAT}}},
	}

	templateNote1 := &templateNote{
		next:         templateNote2,
		isAltered:    false,
		notAltered:   &Note{name: C},
		alteredNotes: nil,
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

	return &templateInstance{templateNote1}
}
