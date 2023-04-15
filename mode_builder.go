package muse

import (
	"github.com/pkg/errors"
)

const (
	DegreesInDiatonic = DegreeNum(7)
	DegreesInTonality = DegreeNum(17)
)

// modeBuilder is the instance of mode builder that will receive ready mode and the first note (and possibly the method of building the mode)
// and will contain different algorithms for different types of modes (and/or methods of building the chord).
// The builder will be responsible for using an immutable set of predefined notes.
type modeBuilder struct{}

// newModeBuilder is constructor for mode builder.
func newModeBuilder() *modeBuilder {
	return &modeBuilder{}
}

// build7DegreeMode builds seven-degree modes.
func (mb *modeBuilder) build7DegreeMode(modeName ModeName, modeTemplate ModeTemplate, firstNote *Note) *Mode {
	mode := &Mode{name: modeName}

	// Insert first note in the mode
	mode.InsertNote(firstNote, 0)

	// Build and insert all other notes
	for buildingResult := range coreBuilding(modeTemplate, firstNote) {
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

type coreBuildingIteratorResult func() (*Note, HalfTones)

// coreBuilding returns sequence of notes and halftones from prime note.
// based on a given mode template and a first note.
func coreBuilding(modeTemplate ModeTemplate, firstNote *Note) <-chan coreBuildingIteratorResult {
	send := func(note *Note, halfTones HalfTones) coreBuildingIteratorResult {
		return func() (*Note, HalfTones) { return note, halfTones }
	}

	f := func(c chan coreBuildingIteratorResult) {
		// Get instance with 12 template notes
		templateNotes := getTemplateNotesInstance()

		// Get the first template note based on the first note of the mode
		templateNote := templateNotes.getTemplateNote(firstNote)
		if templateNote == nil {
			panic(errors.New("cant find first template note"))
		}

		// Save the first note of the mode into the template note
		templateNote.saveTonic(firstNote)

		// Set last used "base" note (clean note without alteration symbols)
		templateNotes.SetLastUsedBaseNote(firstNote)

		// Iterate through the mode template
		for iteratorResult := range modeTemplate.IterateOneRound(false) {
			_, halfTones, halfTonesFromPrime := iteratorResult()

			// Get next template note based on mode template's step
			nextTemplateNote := templateNote.getByHalftones(halfTones)

			// Get nex base note (note after last used base note)
			nextBaseNote := templateNotes.NextBaseNote()

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

	c := make(chan coreBuildingIteratorResult)
	go f(c)

	return c
}

// isModalPositionsActual checks whether the concept of modal position makes sense for the given mode.
func isModalPositionsActual(mt ModeTemplate) bool {
	// TODO: another modes?
	return mt.IsDiatonic()
}

func coreBuildingIntervals(modeTemplate ModeTemplate, firstNote *Note) <-chan coreBuildingIteratorResult {
	send := func(note *Note, halfTones HalfTones) coreBuildingIteratorResult {
		return func() (*Note, HalfTones) { return note, halfTones }
	}

	f := func(c chan coreBuildingIteratorResult) {
		// Get instance with 12 template notes
		templateNotes := getTemplateNotesInstance()

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

	c := make(chan coreBuildingIteratorResult)
	go f(c)

	return c
}
