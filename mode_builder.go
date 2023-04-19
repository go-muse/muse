package muse

const (
	DegreesInDiatonic   = DegreeNum(7)
	DegreesInHeptatonic = DegreeNum(7)
	DegreesInTonality   = DegreeNum(17)
)

// modeBuilder is the instance of mode builder that will receive ready mode and the first note (and possibly the method of building the mode)
// and will contain different algorithms for different types of modes (and/or methods of building the chord).
// The builder will be responsible for using an immutable set of predefined notes.
type modeBuilder interface {
	build(ModeName, ModeTemplate, *Note) *Mode
}

// newModeBuilder is constructor for mode builder.
func newModeBuilder(modeTemplate ModeTemplate) modeBuilder {
	switch modeTemplate.Length() {
	case DegreesInHeptatonic:
		return &modeBuilder7degree{}
	default:
		return &modeBuilderCommon{}
	}
}

// isModalPositionsActual checks whether the concept of modal position makes sense for the given mode.
func isModalPositionsActual(mt ModeTemplate) bool {
	// TODO: another modes?
	return mt.IsHeptatonic()
}
