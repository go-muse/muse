package muse

// utils is a testing helpers functions

func generateDegrees(degreesAmount uint, isCycled bool) *Degree {
	if degreesAmount < 1 {
		return nil
	}
	degreeNum := DegreeNum(1)
	firstDegree := &Degree{number: degreeNum}
	currentDegree := firstDegree
	for degreeNum++; degreeNum <= DegreeNum(degreesAmount); degreeNum++ {
		newDegree := &Degree{number: degreeNum}
		currentDegree.AttachNext(newDegree)
		currentDegree = newDegree
	}

	if isCycled {
		firstDegree.AttachPrevious(currentDegree)
	}

	return firstDegree
}

func generateDegreesWithNotes(isCycled bool, modeTemplate ModeTemplate, firstNote *Note) *Degree {
	if modeTemplate.Length() < 1 {
		return nil
	}

	firstDegree := &Degree{number: 1, note: firstNote, halfTonesFromPrime: 0}
	currentDegree := firstDegree
	for res := range coreBuilding(modeTemplate, firstNote) {
		note, halfTonesFromPrime := res()
		newDegree := &Degree{number: currentDegree.number + 1, note: note, halfTonesFromPrime: halfTonesFromPrime}
		currentDegree.AttachNext(newDegree)
		currentDegree = newDegree
	}

	if isCycled {
		firstDegree.AttachPrevious(currentDegree)
	}

	return firstDegree
}

func generateModeWithNotes(mt ModeTemplate, noteNames []NoteName) *Mode {
	if len(mt) != len(noteNames) {
		panic("length incompatible")
	}

	if mt.Validate() != nil {
		panic(mt.Validate())
	}

	firstDegree := NewDegree(1, 0, nil, nil, MustNewNote(noteNames[0]), nil, nil)
	currentDegree := firstDegree
	var halfTonesFromPrime HalfTones
	const degreeIndexShift = 2
	const noteNameShift = 1
	for i, halfTone := range mt {
		if i == len(mt)-1 {
			break
		}

		degree := NewDegree(DegreeNum(i+degreeIndexShift), 0, nil, nil, newNote(noteNames[i+noteNameShift]), nil, nil)

		halfTonesFromPrime += halfTone
		degree.halfTonesFromPrime = halfTonesFromPrime

		currentDegree.AttachNext(degree)
		currentDegree = degree
	}

	firstDegree.AttachPrevious(currentDegree)

	return &Mode{degree: firstDegree}
}

func generateAbstractPatterns(modeTemplate []HalfTones) [][]HalfTones {
	var res [][]HalfTones
	for j := len(modeTemplate); j > 0; j-- {
		k := j
		for i := 0; i+j <= len(modeTemplate); i++ {
			res = append(res, modeTemplate[i:k])
			k++
		}
	}

	return res
}
