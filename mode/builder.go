package mode

import (
	"github.com/go-muse/muse/builder"
	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/halftone"
	"github.com/go-muse/muse/note"
)

const (
	DegreesInDiatonic   = degree.Number(7)
	DegreesInHeptatonic = degree.Number(7)
	DegreesInPentatonic = degree.Number(5)
	DegreesInTonality   = degree.Number(17)
)

// modeBuilder is the instance of mode builder that will receive mode name and the first note and choose the method of building the mode.
// The builder will be responsible for using an immutable set of predefined notes.
type modeBuilder struct {
	modeTemplate halftone.Template
	buildingFunc
}

// buildingFunc is a function that builds notes and halftone for the mode.
type buildingFunc func(builder.HalftonesIterator, *note.Note) builder.Builder

func (mbc *modeBuilder) build(modeName Name, firstNote *note.Note) *Mode {
	mode := &Mode{name: modeName}

	// Insert first note in the mode
	mode.InsertNote(firstNote, 0)

	// Build and insert all other notes
	for buildingResult := range mbc.buildingFunc(mbc.modeTemplate, firstNote) {
		mode.InsertNote(buildingResult())
	}

	// Closing the circle of degrees in the mode by default
	mode.CloseCircleOfDegrees()

	// Modal positions make sense not in all modes
	if isModalPositionsActual(Template(mbc.modeTemplate)) {
		// Calculation of relative and absolute modal positions of the mode
		mode.setRelativeModalPositions(Template(mbc.modeTemplate))
		mode.setAbsoluteModalPositions()
	}

	return mode
}

// newModeBuilder is constructor for mode builder.
func newModeBuilder(modeTemplate Template) *modeBuilder {
	switch {
	case modeTemplate.Length() == DegreesInHeptatonic:
		return &modeBuilder{halftone.Template(modeTemplate), builder.NewBuilderHeptatonic}
	case modeTemplate.Length() < DegreesInHeptatonic:
		return &modeBuilder{halftone.Template(modeTemplate), builder.NewBuilderCommon}
	default:
		panic("Modes with eight or more steps are not yet supported.")
	}
}

// isModalPositionsActual checks whether the concept of modal position makes sense for the given mode.
func isModalPositionsActual(mt Template) bool {
	// TODO: another modes?
	return mt.IsHeptatonic()
}
