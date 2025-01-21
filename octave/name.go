package octave

// Name is a name for the octave.
type Name string

// NewOctave creates a new octave by the octave name.
func (n Name) NewOctave() (*Octave, error) {
	return New(n)
}

// MustNewOctave creates a new octave by the octave name with panic on error.
func (n Name) MustNewOctave() *Octave {
	return MustNew(n)
}
