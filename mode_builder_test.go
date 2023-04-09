package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild7DegreeMode(t *testing.T) {
	mb := newModeBuilder()
	firstNote := &Note{name: C}
	mode := mb.build7DegreeMode(ModeNameIonian, TemplateIonian(), firstNote)

	d1 := &Degree{
		number: 1,
		note:   firstNote,
	}
	d2 := &Degree{
		number:   2,
		note:     &Note{name: D},
		previous: d1,
	}
	d3 := &Degree{
		number:   3,
		note:     &Note{name: E},
		previous: d2,
	}
	d4 := &Degree{
		number:   4,
		note:     &Note{name: F},
		previous: d3,
	}
	d5 := &Degree{
		number:   5,
		note:     &Note{name: G},
		previous: d4,
	}
	d6 := &Degree{
		number:   6,
		note:     &Note{name: A},
		previous: d5,
	}
	d7 := &Degree{
		number:   7,
		note:     &Note{name: B},
		previous: d6,
		next:     d1,
	}

	d1.previous = d7
	d1.next = d2
	d2.next = d3
	d3.next = d4
	d4.next = d5
	d5.next = d6
	d6.next = d7

	m := &Mode{
		name:   ModeNameIonian,
		degree: d1,
	}

	d1chan := mode.degree.IterateOneRound(false)
	for d2 := range m.degree.IterateOneRound(false) {
		assert.True(t, (<-d1chan).Equal(d2), "d1: %+v, d2: %+v", d1, d2)
	}
}
