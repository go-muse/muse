package builder

import (
	"unsafe"

	"github.com/go-muse/muse/common/convert"
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

// NewBuilderHeptatonic builds sequence of notes and halftone for heptatonic mode
// from prime note based on a given mode template and a first note.
func NewBuilderHeptatonic(modeTemplate HalftonesIterator, firstNote *note.Note) Builder {
	send := func(n *note.Note, halfTones halftone.HalfTones) func() (*note.Note, halftone.HalfTones) {
		return func() (*note.Note, halftone.HalfTones) { return n, halfTones }
	}

	f := func(c chan func() (*note.Note, halftone.HalfTones)) {
		// Get instance with 12 template notes
		templateNotes := getTemplateNotesHeptatonic()

		// Get the first template note based on the first note of the mode
		templateNote := templateNotes.getTemplateNote(firstNote)
		if templateNote == nil {
			panic(errInvalidFirstTemplateNote)
		}

		// Calculate halftone from the prime note for all the degrees
		templateNote.calculateHalftonesFromPrime()

		// Set last used "base" note (clean note without accidentals)
		templateNotes.setLastUsedBaseNote(firstNote)

		// Iterate through the mode template
		for iteratorResult := range modeTemplate.Iterate() {
			halfTones, halfTonesFromPrime := iteratorResult()

			// Get next template note based on mode template's step
			nextTemplateNote := templateNote.getByHalftones(halfTones)

			// Get next base note (note after last used base note)
			nextBaseNote := templateNotes.nextBaseNote()

			// Get next template note that is base note in the template notes chain
			nextTemplateNoteByBase := templateNotes.getTemplateNote(nextBaseNote)

			// Difference between their distance from tonic gives us understanding what to do with the next "clean" note
			diff := convert.SubUint8Uint8(uint8(nextTemplateNoteByBase.halfTonesFromPrime), uint8(nextTemplateNote.halfTonesFromPrime))

			// Alteration of the clean note by its distance from the clean note with the same name
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

			// send resulting note with distance from the prime note in halftone
			c <- send(nextBaseNote, halfTonesFromPrime)
		}

		close(c)
	}

	c := make(chan func() (*note.Note, halftone.HalfTones))
	go f(c)

	return c
}

// templateNoteHeptatonic is a template note for the single template instance.
type templateNoteHeptatonic struct {
	next               *templateNoteHeptatonic
	halfTonesFromPrime halftone.HalfTones
	allNotes           []*noteRelation
}

// templateNotesHeptatonic is the instance with template notes.
type templateNotesHeptatonic struct {
	*templateNoteHeptatonic                     // the first note of the linked list in the template instance
	baseNote                *baseNoteHeptatonic // last used note without accidental to build the mode
}

// getTemplateNote determines template note by the given Note.
// When building a mode from a specific note, we will need to move the pointer to the first note
// and from it iterate further and pass it to the final notes builder, which decides on their naming.
func (ti *templateNotesHeptatonic) getTemplateNote(note *note.Note) *templateNoteHeptatonic {
	if ti.templateNoteHeptatonic == nil {
		return nil
	}

	currentTemplateNote := ti.templateNoteHeptatonic
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

// equal checks whether the specified note is equal to one of the notes contained in the current template note.
func (tn *templateNoteHeptatonic) equal(note *note.Note) bool {
	for _, ar := range tn.allNotes {
		if ar.realNote().IsEqualByName(note) {
			return true
		}
	}

	return false
}

// getNext returns the text template note.
func (tn *templateNoteHeptatonic) getNext() *templateNoteHeptatonic {
	return tn.next
}

// nextBaseNote sets next base note as current and returns it.
func (ti *templateNotesHeptatonic) nextBaseNote() *note.Note {
	ti.baseNote = ti.baseNote.next

	return ti.baseNote.note
}

// calculateHalftonesFromPrime saves the first note of the constructed mode.
func (tn *templateNoteHeptatonic) calculateHalftonesFromPrime() {
	current := tn
	halfTonesFromPrime := current.halfTonesFromPrime
	for unsafe.Pointer(current.next) != unsafe.Pointer(tn) {
		current = current.getNext()
		halfTonesFromPrime++
		current.halfTonesFromPrime = halfTonesFromPrime
	}
}

// setLastUsedBaseNote sets last used base note as current.
func (ti *templateNotesHeptatonic) setLastUsedBaseNote(note *note.Note) {
	currentNote := ti.baseNote
	for unsafe.Pointer(currentNote.next) != unsafe.Pointer(ti.baseNote) {
		currentNote = currentNote.next
		if newBaseNote(note).IsEqualByName(currentNote.note) {
			ti.baseNote = currentNote

			break
		}
	}
}

// getByHalftones returns a template note that is several half steps away from the current one.
func (tn *templateNoteHeptatonic) getByHalftones(halfTones halftone.HalfTones) *templateNoteHeptatonic {
	templateNote := tn
	for range halfTones {
		templateNote = templateNote.getNext()
	}

	return templateNote
}

// baseNote is one of "clean" notes: [C, D, E, F, G, A, B].
// It is involved in calculating the alteration of the constructed notes of the mode.
type baseNoteHeptatonic struct {
	next *baseNoteHeptatonic
	note *note.Note
}

// getTemplateNotesHeptatonic creates 12 template notes for the instance when it is created
// this will be hardcoded as to what each template note contains.
func getTemplateNotesHeptatonic() *templateNotesHeptatonic {
	templateNote12 := &templateNoteHeptatonic{
		next:     nil,
		allNotes: []*noteRelation{{note.B.MustNewNote(), note.B.MustNewNote()}, {note.A.MustNewNote(), note.ASHARP2.MustNewNote()}, {note.C.MustNewNote(), note.CFLAT.MustNewNote()}},
	}

	templateNote11 := &templateNoteHeptatonic{
		next:     templateNote12,
		allNotes: []*noteRelation{{note.A.MustNewNote(), note.ASHARP.MustNewNote()}, {note.B.MustNewNote(), note.BFLAT.MustNewNote()}, {note.C.MustNewNote(), note.CFLAT2.MustNewNote()}},
	}

	templateNote10 := &templateNoteHeptatonic{
		next:     templateNote11,
		allNotes: []*noteRelation{{note.A.MustNewNote(), note.A.MustNewNote()}, {note.B.MustNewNote(), note.BFLAT2.MustNewNote()}, {note.G.MustNewNote(), note.GSHARP2.MustNewNote()}},
	}

	templateNote9 := &templateNoteHeptatonic{
		next:     templateNote10,
		allNotes: []*noteRelation{{note.G.MustNewNote(), note.GSHARP.MustNewNote()}, {note.A.MustNewNote(), note.AFLAT.MustNewNote()}},
	}

	templateNote8 := &templateNoteHeptatonic{
		next:     templateNote9,
		allNotes: []*noteRelation{{note.G.MustNewNote(), note.G.MustNewNote()}, {note.A.MustNewNote(), note.AFLAT2.MustNewNote()}, {note.F.MustNewNote(), note.FSHARP2.MustNewNote()}},
	}

	templateNote7 := &templateNoteHeptatonic{
		next:     templateNote8,
		allNotes: []*noteRelation{{note.G.MustNewNote(), note.GFLAT.MustNewNote()}, {note.F.MustNewNote(), note.FSHARP.MustNewNote()}, {note.E.MustNewNote(), note.ESHARP2.MustNewNote()}},
	}

	templateNote6 := &templateNoteHeptatonic{
		next:     templateNote7,
		allNotes: []*noteRelation{{note.F.MustNewNote(), note.F.MustNewNote()}, {note.E.MustNewNote(), note.ESHARP.MustNewNote()}, {note.G.MustNewNote(), note.GFLAT2.MustNewNote()}},
	}

	templateNote5 := &templateNoteHeptatonic{
		next:     templateNote6,
		allNotes: []*noteRelation{{note.E.MustNewNote(), note.E.MustNewNote()}, {note.F.MustNewNote(), note.FFLAT.MustNewNote()}, {note.D.MustNewNote(), note.DSHARP2.MustNewNote()}},
	}

	templateNote4 := &templateNoteHeptatonic{
		next:     templateNote5,
		allNotes: []*noteRelation{{note.E.MustNewNote(), note.EFLAT.MustNewNote()}, {note.D.MustNewNote(), note.DSHARP.MustNewNote()}, {note.F.MustNewNote(), note.FFLAT2.MustNewNote()}},
	}

	templateNote3 := &templateNoteHeptatonic{
		next:     templateNote4,
		allNotes: []*noteRelation{{note.D.MustNewNote(), note.D.MustNewNote()}, {note.C.MustNewNote(), note.CSHARP2.MustNewNote()}, {note.E.MustNewNote(), note.EFLAT2.MustNewNote()}},
	}

	templateNote2 := &templateNoteHeptatonic{
		next:     templateNote3,
		allNotes: []*noteRelation{{note.D.MustNewNote(), note.DFLAT.MustNewNote()}, {note.C.MustNewNote(), note.CSHARP.MustNewNote()}, {note.B.MustNewNote(), note.BSHARP2.MustNewNote()}},
	}

	templateNote1 := &templateNoteHeptatonic{
		next:     templateNote2,
		allNotes: []*noteRelation{{note.C.MustNewNote(), note.C.MustNewNote()}, {note.B.MustNewNote(), note.BSHARP.MustNewNote()}, {note.D.MustNewNote(), note.DFLAT2.MustNewNote()}},
	}

	// There must be cycling,
	// this will allow iteration within an octave from any note to any number of steps
	// for example, from E to E
	templateNote12.next = templateNote1

	baseNote7 := &baseNoteHeptatonic{
		next: nil,
		note: note.B.MustNewNote(),
	}
	baseNote6 := &baseNoteHeptatonic{
		next: baseNote7,
		note: note.A.MustNewNote(),
	}
	baseNote5 := &baseNoteHeptatonic{
		next: baseNote6,
		note: note.G.MustNewNote(),
	}
	baseNote4 := &baseNoteHeptatonic{
		next: baseNote5,
		note: note.F.MustNewNote(),
	}
	baseNote3 := &baseNoteHeptatonic{
		next: baseNote4,
		note: note.E.MustNewNote(),
	}
	baseNote2 := &baseNoteHeptatonic{
		next: baseNote3,
		note: note.D.MustNewNote(),
	}
	baseNote1 := &baseNoteHeptatonic{
		next: baseNote2,
		note: note.C.MustNewNote(),
	}

	baseNote7.next = baseNote1

	return &templateNotesHeptatonic{templateNote1, baseNote1}
}
