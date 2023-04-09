package muse

// Scale is a set of notes.
type Scale []*Note

// GenerateScale generates  an ascending or descending scale.
func (m *Mode) GenerateScale(desc bool) Scale {
	if m == nil || m.degree == nil || m.degree.note == nil {
		return nil
	}

	scale := make(Scale, m.Length())

	var fromDegree *Degree
	if desc {
		fromDegree = m.GetLastDegree()
	} else {
		fromDegree = m.GetFirstDegree()
	}

	var i int
	for degree := range fromDegree.IterateOneRound(desc) {
		scale[i] = degree.Note()
		i++
	}

	return scale
}

// GetFullChromaticScale returns all notes of the tonal system.
func GetFullChromaticScale() Scale {
	return Scale{
		newNote(C),
		newNote(DFLAT),
		newNote(CSHARP),
		newNote(D),
		newNote(EFLAT),
		newNote(DSHARP),
		newNote(E),
		newNote(F),
		newNote(GFLAT),
		newNote(FSHARP),
		newNote(G),
		newNote(AFLAT),
		newNote(GSHARP),
		newNote(A),
		newNote(BFLAT),
		newNote(ASHARP),
		newNote(B),
	}
}
