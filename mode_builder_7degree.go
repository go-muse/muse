package muse

import (
	"unsafe"

	"github.com/pkg/errors"
)

type modeBuilder7degree struct{}

// build7DegreeMode builds seven-degree modes.
func (mb *modeBuilder7degree) build(modeName ModeName, modeTemplate ModeTemplate, firstNote *Note) *Mode {
	mode := &Mode{name: modeName}

	// Insert first note in the mode
	mode.InsertNote(firstNote, 0)

	// Build and insert all other notes
	for buildingResult := range coreBuilding7degree(modeTemplate, firstNote) {
		mode.InsertNote(buildingResult())
	}

	// Closing the circle of degrees in the mode by default
	mode.CloseCircleOfDegrees()

	// Modal positions make sense not in all modes
	if isModalPositionsActual(modeTemplate) {
		// Calculation of relative and absolute modal positions of the mode
		mode.setRelativeModalPositions(modeTemplate)
		mode.setAbsoluteModalPositions()
	}

	return mode
}

type coreBuilding7degreeIteratorResult func() (*Note, HalfTones)

// coreBuilding7degree returns sequence of notes and halftones from prime note
// based on a given mode template and a first note.
func coreBuilding7degree(modeTemplate ModeTemplate, firstNote *Note) <-chan coreBuilding7degreeIteratorResult {
	send := func(note *Note, halfTones HalfTones) coreBuilding7degreeIteratorResult {
		return func() (*Note, HalfTones) { return note, halfTones }
	}

	f := func(c chan coreBuilding7degreeIteratorResult) {
		// Get instance with 12 template notes
		templateNotes := getTemplateNotes7degree()

		// Get the first template note based on the first note of the mode
		templateNote := templateNotes.getTemplateNote(firstNote)
		if templateNote == nil {
			panic(errors.New("cant find first template note"))
		}

		// Calculate half tones from the prime note for all the degrees
		templateNote.calculateHalftonesFromPrime()

		// Set last used "base" note (clean note without alteration symbols)
		templateNotes.setLastUsedBaseNote(firstNote)

		// Iterate through the mode template
		for iteratorResult := range modeTemplate.IterateOneRound(false) {
			_, halfTones, halfTonesFromPrime := iteratorResult()

			// Get next template note based on mode template's step
			nextTemplateNote := templateNote.getByHalftones(halfTones)

			// Get next base note (note after last used base note)
			nextBaseNote := templateNotes.nextBaseNote()

			// Get next template note that is base note in the template notes chain
			nextTemplateNoteByBase := templateNotes.getTemplateNote(nextBaseNote)

			// Difference between their distance from tonic gives us understanding what to do with the next "clean" note
			diff := int8(nextTemplateNoteByBase.halfTonesFromPrime) - int8(nextTemplateNote.halfTonesFromPrime)

			// Alteration of the clean note by it's distance from the clean note with the same name
			for diff != 0 {
				switch {
				case diff > 0:
					nextBaseNote.AlterDown()
					diff--
				case diff < 0:
					nextBaseNote.AlterUp()
					diff++
				}
			}

			// Insert the next note into the current note variable, to use it at the next iteration
			templateNote = nextTemplateNote

			// send resulting note with distance from the prime note in half tones
			c <- send(nextBaseNote, halfTonesFromPrime)
		}

		close(c)
	}

	c := make(chan coreBuilding7degreeIteratorResult)
	go f(c)

	return c
}

// templateNote7degree is a template note for the single template instance.
type templateNote7degree struct {
	next               *templateNote7degree
	halfTonesFromPrime HalfTones
	allNotes           []*noteRelation
}

// templateNotes7degree is the instance with template notes.
type templateNotes7degree struct {
	*templateNote7degree                  // the first note of the linked list in the template instance
	baseNote             *baseNote7degree // last used note without alteration symbol to build the mode
}

// getTemplateNote determines template note by the given Note.
// When building a mode from a specific note, we will need to move the pointer to the first note
// and from it iterate further and pass it to the final notes builder, which decides on their naming.
func (ti *templateNotes7degree) getTemplateNote(note *Note) *templateNote7degree {
	if ti.templateNote7degree == nil {
		return nil
	}

	currentTemplateNote := ti.templateNote7degree
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

// equal checks whether the specified note is equal to one of the notes contained in the current template note.
func (tn *templateNote7degree) equal(note *Note) bool {
	for _, ar := range tn.allNotes {
		if ar.realNote().IsEqualByName(note) {
			return true
		}
	}

	return false
}

// getNext returns the text template note.
func (tn *templateNote7degree) getNext() *templateNote7degree {
	return tn.next
}

// nextBaseNote sets next base note as current and returns it.
func (ti *templateNotes7degree) nextBaseNote() *Note {
	ti.baseNote = ti.baseNote.next

	return ti.baseNote.note
}

// calculateHalftonesFromPrime saves the first note of the constructed mode.
func (tn *templateNote7degree) calculateHalftonesFromPrime() {
	current := tn
	halfTonesFromPrime := current.halfTonesFromPrime
	for unsafe.Pointer(current.next) != unsafe.Pointer(tn) {
		current = current.getNext()
		halfTonesFromPrime++
		current.halfTonesFromPrime = halfTonesFromPrime
	}
}

// setLastUsedBaseNote sets last used base note as current.
func (ti *templateNotes7degree) setLastUsedBaseNote(note *Note) {
	currentNote := ti.baseNote
	for unsafe.Pointer(currentNote.next) != unsafe.Pointer(ti.baseNote) {
		currentNote = currentNote.next
		if note.baseNote().IsEqualByName(currentNote.note) {
			ti.baseNote = currentNote

			break
		}
	}
}

// getByHalftones returns a template note that is several half steps away from the current one.
func (tn *templateNote7degree) getByHalftones(halfTones HalfTones) *templateNote7degree {
	templateNote := tn
	for i := 0; i < int(halfTones); i++ {
		templateNote = templateNote.getNext()
	}

	return templateNote
}

// baseNote is one of "clean" notes: [C, D, E, F, G, A, B].
// It is involved in calculating the alteration of the constructed notes of the mode.
type baseNote7degree struct {
	next *baseNote7degree
	note *Note
}

// getTemplateNotes7degree creates 12 template notes for the instance when it is created
// this will be hardcoded as to what each template note contains.
func getTemplateNotes7degree() *templateNotes7degree {
	templateNote12 := &templateNote7degree{
		next:     nil,
		allNotes: []*noteRelation{{newNote(B), newNote(B)}, {newNote(A), newNote(ASHARP2)}, {newNote(C), newNote(CFLAT)}},
	}

	templateNote11 := &templateNote7degree{
		next:     templateNote12,
		allNotes: []*noteRelation{{&Note{name: A}, &Note{name: ASHARP}}, {&Note{name: B}, &Note{name: BFLAT}}, {newNote(C), newNote(CFLAT2)}},
	}

	templateNote10 := &templateNote7degree{
		next:     templateNote11,
		allNotes: []*noteRelation{{newNote(A), newNote(A)}, {newNote(B), newNote(BFLAT2)}, {newNote(G), newNote(GSHARP2)}},
	}

	templateNote9 := &templateNote7degree{
		next:     templateNote10,
		allNotes: []*noteRelation{{newNote(G), newNote(GSHARP)}, {newNote(A), newNote(AFLAT)}},
	}

	templateNote8 := &templateNote7degree{
		next:     templateNote9,
		allNotes: []*noteRelation{{newNote(G), newNote(G)}, {newNote(A), newNote(AFLAT2)}, {newNote(F), newNote(FSHARP2)}},
	}

	templateNote7 := &templateNote7degree{
		next:     templateNote8,
		allNotes: []*noteRelation{{newNote(G), newNote(GFLAT)}, {newNote(F), newNote(FSHARP)}, {newNote(E), newNote(ESHARP2)}},
	}

	templateNote6 := &templateNote7degree{
		next:     templateNote7,
		allNotes: []*noteRelation{{newNote(F), newNote(F)}, {newNote(E), newNote(ESHARP)}, {newNote(G), newNote(GFLAT2)}},
	}

	templateNote5 := &templateNote7degree{
		next:     templateNote6,
		allNotes: []*noteRelation{{newNote(E), newNote(E)}, {newNote(F), newNote(FFLAT)}, {newNote(D), newNote(DSHARP2)}},
	}

	templateNote4 := &templateNote7degree{
		next:     templateNote5,
		allNotes: []*noteRelation{{newNote(E), newNote(EFLAT)}, {newNote(D), newNote(DSHARP)}, {newNote(F), newNote(FFLAT2)}},
	}

	templateNote3 := &templateNote7degree{
		next:     templateNote4,
		allNotes: []*noteRelation{{newNote(D), newNote(D)}, {newNote(C), newNote(CSHARP2)}, {newNote(E), newNote(EFLAT2)}},
	}

	templateNote2 := &templateNote7degree{
		next:     templateNote3,
		allNotes: []*noteRelation{{newNote(D), newNote(DFLAT)}, {newNote(C), newNote(CSHARP)}, {newNote(B), newNote(BSHARP2)}},
	}

	templateNote1 := &templateNote7degree{
		next:     templateNote2,
		allNotes: []*noteRelation{{newNote(C), newNote(C)}, {newNote(B), newNote(BSHARP)}, {newNote(D), newNote(DFLAT2)}},
	}

	// There must be cycling,
	// this will allow iteration within an octave from any note to any number of steps
	// for example, from E to E
	templateNote12.next = templateNote1

	baseNote7 := &baseNote7degree{
		next: nil,
		note: newNote(B),
	}
	baseNote6 := &baseNote7degree{
		next: baseNote7,
		note: newNote(A),
	}
	baseNote5 := &baseNote7degree{
		next: baseNote6,
		note: newNote(G),
	}
	baseNote4 := &baseNote7degree{
		next: baseNote5,
		note: newNote(F),
	}
	baseNote3 := &baseNote7degree{
		next: baseNote4,
		note: newNote(E),
	}
	baseNote2 := &baseNote7degree{
		next: baseNote3,
		note: newNote(D),
	}
	baseNote1 := &baseNote7degree{
		next: baseNote2,
		note: newNote(C),
	}

	baseNote7.next = baseNote1

	return &templateNotes7degree{templateNote1, baseNote1}
}
