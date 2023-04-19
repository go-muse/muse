package muse

import "github.com/pkg/errors"

type modeBuilderCommon struct{}

func (cbc *modeBuilderCommon) build(modeName ModeName, modeTemplate ModeTemplate, firstNote *Note) *Mode {
	mode := &Mode{name: modeName}

	// Insert first note in the mode
	mode.InsertNote(firstNote, 0)

	// Build and insert all other notes
	for buildingResult := range coreBuildingCommon(modeTemplate, firstNote) {
		mode.InsertNote(buildingResult())
	}

	// Closing the circle of degrees in the mode by default
	mode.CloseCircleOfDegrees()

	return mode
}



// coreBuildingCommonIteratorResult is the type for sending multiple variables to the iterative channel.
type coreBuildingCommonIteratorResult func() (*Note, HalfTones)

// coreBuildingCommon builds notes for the mode.
func coreBuildingCommon(modeTemplate ModeTemplate, firstNote *Note) <-chan coreBuildingCommonIteratorResult {
	send := func(note *Note, halfTones HalfTones) coreBuildingCommonIteratorResult {
		return func() (*Note, HalfTones) { return note, halfTones }
	}

	f := func(c chan coreBuildingCommonIteratorResult) {
		// Get instance with 12 template notes
		templateNotes := getTemplateNotesCommon()

		// Get the first template note based on the first note of the mode
		templateNote := templateNotes.getTemplateNote(firstNote)
		if templateNote == nil {
			panic(errors.New("cant find first template note"))
		}

		// Save the first note of the mode into the template note
		templateNote.saveResultingNote(firstNote)

		// Iterate through the mode template
		for iteratorResult := range modeTemplate.IterateOneRound(false) {
			_, halfTones, halfTonesFromPrime := iteratorResult()

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

			// send resulting note with distance from the prime note in half tones
			c <- send(newNote, halfTonesFromPrime)
		}

		close(c)
	}

	c := make(chan coreBuildingCommonIteratorResult)
	go f(c)

	return c
}

// templateNoteCommon is a template note for the single template instance.
type templateNoteCommon struct {
	next          *templateNoteCommon
	previous      *templateNoteCommon
	isAltered     bool
	notAltered    *Note
	alteredNotes  []*noteRelation
	resultingNote *noteRelation
}

// templateNotes7degree is the instance with template notes.
type templateNotesCommon struct {
	*templateNoteCommon           // the first note of the linked list in the template instance
	baseNote            *baseNote // last used note without alteration symbol to build the mode
}

// Equal Compares Note and template note by name.
func (tn *templateNoteCommon) equal(note *Note) bool {
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

func (tn *templateNoteCommon) getNext() *templateNoteCommon {
	return tn.next
}

// getTemplateNote determines template note by the given Note.
// When building a mode from a specific note, we will need to move the pointer to the first note
// and from it iterate further and pass it to the final notes builder, which decides on their naming.
func (ti *templateNotesCommon) getTemplateNote(note *Note) *templateNoteCommon {
	if ti.templateNoteCommon == nil {
		return nil
	}

	currentTemplateNote := ti.templateNoteCommon
	if currentTemplateNote.equal(note) {
		return currentTemplateNote
	}

	// the condition that it will not endlessly drive in a circle in search of a non-existent note
	for i := 2; i <= int(HalftonesInOctave); i++ {
		currentTemplateNote = currentTemplateNote.getNext()
		if currentTemplateNote.equal(note) {
			return currentTemplateNote
		}
	}

	return nil
}

// getByHalftones returns a template note that is several half steps away from the current one.
func (tn *templateNoteCommon) getByHalftones(halfTones HalfTones) *templateNoteCommon {
	templateNote := tn
	for i := 0; i < int(halfTones); i++ {
		templateNote = templateNote.getNext()
	}

	return templateNote
}

// saveResultingNote inserts the final note into the template note
// with preserving the information about which note is the base for the saved note.
func (tn *templateNoteCommon) saveResultingNote(note *Note) {
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
func (tn *templateNoteCommon) buildNoteByPrevious() *Note {
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

// baseNote is one of "clean" notes: [C, D, E, F, G, A, B].
// It is involved in calculating the alteration of the constructed notes of the mode.
type baseNote struct {
	prev, next *baseNote
	note       *Note
}

// getTemplateNotesCommon creates 12 template notes for the instance when it is created
// this will be hardcoded as to what each template note contains.
func getTemplateNotesCommon() *templateNotesCommon {
	templateNote12 := &templateNoteCommon{
		next:       nil,
		isAltered:  false,
		notAltered: &Note{name: B},
	}

	templateNote11 := &templateNoteCommon{
		next:         templateNote12,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: A}, &Note{name: ASHARP}}, {&Note{name: B}, &Note{name: BFLAT}}},
	}

	templateNote10 := &templateNoteCommon{
		next:         templateNote11,
		isAltered:    false,
		notAltered:   &Note{name: A},
		alteredNotes: nil,
	}

	templateNote9 := &templateNoteCommon{
		next:         templateNote10,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: G}, &Note{name: GSHARP}}, {&Note{name: A}, &Note{name: AFLAT}}},
	}

	templateNote8 := &templateNoteCommon{
		next:         templateNote9,
		isAltered:    false,
		notAltered:   &Note{name: G},
		alteredNotes: nil,
	}

	templateNote7 := &templateNoteCommon{
		next:         templateNote8,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: F}, &Note{name: FSHARP}}, {&Note{name: G}, &Note{name: GFLAT}}},
	}

	templateNote6 := &templateNoteCommon{
		next:         templateNote7,
		isAltered:    false,
		notAltered:   &Note{name: F},
		alteredNotes: nil,
	}

	templateNote5 := &templateNoteCommon{
		next:         templateNote6,
		isAltered:    false,
		notAltered:   &Note{name: E},
		alteredNotes: nil,
	}

	templateNote4 := &templateNoteCommon{
		next:         templateNote5,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: D}, &Note{name: DSHARP}}, {&Note{name: E}, &Note{name: EFLAT}}},
	}

	templateNote3 := &templateNoteCommon{
		next:         templateNote4,
		isAltered:    false,
		notAltered:   &Note{name: D},
		alteredNotes: nil,
	}

	templateNote2 := &templateNoteCommon{
		next:         templateNote3,
		isAltered:    true,
		notAltered:   nil,
		alteredNotes: []*noteRelation{{&Note{name: C}, &Note{name: CSHARP}}, {&Note{name: D}, &Note{name: DFLAT}}},
	}

	templateNote1 := &templateNoteCommon{
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

	baseNote7 := &baseNote{
		prev: nil,
		next: nil,
		note: newNote(B),
	}
	baseNote6 := &baseNote{
		prev: nil,
		next: baseNote7,
		note: newNote(A),
	}
	baseNote5 := &baseNote{
		prev: nil,
		next: baseNote6,
		note: newNote(G),
	}
	baseNote4 := &baseNote{
		prev: nil,
		next: baseNote5,
		note: newNote(F),
	}
	baseNote3 := &baseNote{
		prev: nil,
		next: baseNote4,
		note: newNote(E),
	}
	baseNote2 := &baseNote{
		prev: nil,
		next: baseNote3,
		note: newNote(D),
	}
	baseNote1 := &baseNote{
		prev: nil,
		next: baseNote2,
		note: newNote(C),
	}

	baseNote7.next = baseNote1

	return &templateNotesCommon{templateNote1, baseNote1}
}
