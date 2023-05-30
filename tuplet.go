package muse

// A tuplet is a characteristic of the duration of a given note,
// indicating how many notes of such a tuplet are equal in duration to how many notes outside the tuplet.
type Tuplet struct {
	n uint8
	m uint8
}

// NewTuplet creates new tuplet with the given values as n/m.
// It means that "n" of such notes will last as "m" notes of the same duration outside the tuplet.
func NewTuplet(n, m uint8) *Tuplet {
	return &Tuplet{
		n: n,
		m: m,
	}
}

// Set sets the setting that "n" of such notes will last as "m" notes of the same duration outside the tuplet.
func (t *Tuplet) Set(n, m uint8) *Tuplet {
	if t != nil {
		t.n = n
		t.m = m
	} else {
		t = &Tuplet{
			n: n,
			m: m,
		}
	}

	return t
}

// SetTriplet sets the setting that three such notes will last as long as two notes of the same duration outside the tuplet.
//
//nolint:gomnd
func (t *Tuplet) SetTriplet() *Tuplet {
	n := uint8(3)
	m := uint8(2)
	if t != nil {
		t.n = n
		t.m = m
	} else {
		t = &Tuplet{
			n: n,
			m: m,
		}
	}

	return t
}

// SetDuplet sets the setting that two such notes will last as long as three notes of the same duration outside the tuplet.
//
//nolint:gomnd
func (t *Tuplet) SetDuplet() *Tuplet {
	n := uint8(2)
	m := uint8(3)
	if t != nil {
		t.n = n
		t.m = m
	} else {
		t = &Tuplet{
			n: n,
			m: m,
		}
	}

	return t
}

// RemoveTuplet removes the tuplet.
func (t *Tuplet) RemoveTuplet() *Tuplet {
	t = nil

	return t
}
