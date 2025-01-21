package mode

import (
	"github.com/go-muse/muse/degree"
	"github.com/go-muse/muse/scale"
)

// GenerateScale generates an ascending or descending scale.
func (m *Mode) GenerateScale(desc bool) scale.Scale {
	if m == nil || m.degree == nil || m.degree.Note() == nil {
		return nil
	}

	scl := make(scale.Scale, m.Length())

	var fromDegree *degree.Degree
	if desc {
		fromDegree = m.GetLastDegree()
	} else {
		fromDegree = m.GetFirstDegree()
	}

	var i int
	for deg := range fromDegree.IterateOneRound(desc) {
		scl[i] = deg.Note()
		i++
	}

	return scl
}
