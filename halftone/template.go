package halftone

// Template is just a set of halftone.
type Template []HalfTones

// NewTemplate creates a new Template with given values.
func NewTemplate(halftones ...HalfTones) Template {
	template := make(Template, len(halftones))
	copy(template, halftones)

	return template
}

// Length returns amount of halftone in the template.
func (t Template) Length() uint64 {
	return uint64(len(t))
}

// IterateOneRound returns channel for iterating over a halftone' template.
// At each iteration, it returns TemplateIteratorResult containing a set of values.
func (t Template) IterateOneRound(withOctave bool) <-chan func() (HalfTones, HalfTones) {
	send := func(halfTones, halfTonesFromPrime HalfTones) func() (HalfTones, HalfTones) {
		return func() (HalfTones, HalfTones) {
			return halfTones, halfTonesFromPrime
		}
	}

	f := func(c chan func() (HalfTones, HalfTones)) {
		var halfTonesFromPrime HalfTones
		length := t.Length()
		for i, halfTones := range t {
			if i > 0 && uint64(i) == length-1 && !withOctave {
				break
			}
			halfTonesFromPrime += halfTones
			c <- send(halfTones, halfTonesFromPrime)
		}
		close(c)
	}

	c := make(chan func() (HalfTones, HalfTones))
	go f(c)

	return c
}

// Iterate returns channel for iterating over a halftone' template without octave repetition of the first note.
func (t Template) Iterate() <-chan func() (HalfTones, HalfTones) {
	return t.IterateOneRound(false)
}
