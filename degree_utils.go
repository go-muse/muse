package muse

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
