package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDegree_generateDegreesWithNotes(t *testing.T) {
	firstDegree := &Degree{number: 1, halfTonesFromPrime: 0, note: &Note{name: C}}
	secondDegree := &Degree{number: 2, halfTonesFromPrime: 2, note: &Note{name: D}, previous: firstDegree}
	thirdDegree := &Degree{number: 3, halfTonesFromPrime: 4, note: &Note{name: E}, previous: secondDegree}
	fourthDegree := &Degree{number: 4, halfTonesFromPrime: 5, note: &Note{name: F}, previous: thirdDegree}
	fifthDegree := &Degree{number: 5, halfTonesFromPrime: 7, note: &Note{name: G}, previous: fourthDegree}
	sixthDegree := &Degree{number: 6, halfTonesFromPrime: 9, note: &Note{name: A}, previous: fifthDegree}
	seventhDegree := &Degree{number: 7, halfTonesFromPrime: 11, note: &Note{name: B}, previous: sixthDegree}

	firstDegree.previous = seventhDegree

	firstDegree.next = secondDegree
	secondDegree.next = thirdDegree
	thirdDegree.next = fourthDegree
	fourthDegree.next = fifthDegree
	fifthDegree.next = sixthDegree
	sixthDegree.next = seventhDegree
	seventhDegree.next = firstDegree

	resFirstDegree := generateDegreesWithNotes(true, TemplateNaturalMajor(), &Note{name: C})

	currentDegree := firstDegree
	for degree := range resFirstDegree.IterateOneRound(false) {
		assert.Equal(t, currentDegree, degree, "expected: %+v, actual: %+v", currentDegree, degree)
		currentDegree = currentDegree.GetNext()
	}
}

func Test_generateModeWithNotes(t *testing.T) {
	noteNames := []NoteName{C, D, EFLAT, F, G, AFLAT, BFLAT}
	mt := TemplateAeolian()
	mode := generateModeWithNotes(mt, noteNames)
	var halfTonesFromPrime HalfTones
	for degree := range mode.IterateOneRound(false) {
		if degree.Number() >= 2 {
			halfTonesFromPrime += mt[int(degree.Number())-2]
		}
		assert.Equal(t, noteNames[degree.Number()-1], degree.Note().Name())
		if degree.Number() >= 2 {
			assert.Equal(t, halfTonesFromPrime, degree.HalfTonesFromPrime(), "expected :%d, actual: %d", halfTonesFromPrime, degree.HalfTonesFromPrime())
		}
	}
}
