package muse

// // templateNote is a template note for the single template instance.
// type templateNote struct {
// 	isTonic            bool
// 	next               *templateNote
// 	previous           *templateNote
// 	halfTonesFromPrime HalfTones
// 	isAltered          bool
// 	notAltered         *Note
// 	alteredNotes       []*noteRelation
// 	resultingNote      *noteRelation
// }

// func (tn *templateNote) getNext() *templateNote {
// 	return tn.next
// }

// func (tn *templateNote) getPrevious() *templateNote {
// 	return tn.previous
// }

// // getTonic returns the first note of the constructed mode.
// func (tn *templateNote) getTonic() *templateNote {
// 	templateNote := tn
// 	for {
// 		templateNote = templateNote.getNext()
// 		if templateNote.isTonic {
// 			break
// 		}
// 	}

// 	return templateNote
// }

// func (tn *templateNote) getPreviousUsedBaseNote() *Note {
// 	templateNote := tn
// 	for {
// 		templateNote = templateNote.getPrevious()
// 		if templateNote.resultingNote != nil {
// 			break
// 		}
// 	}

// 	return templateNote.resultingNote.base
// }

// // templateInstance is the instance with template notes.
// type templateInstance struct {
// 	*templateNote           // the first note of the linked list in the template instance
// 	baseNote      *baseNote // last used note without alteration symbol to build the mode
// }

// // getTemplateNote determines template note by the given Note.
// // When building a mode from a specific note, we will need to move the pointer to the first note
// // and from it iterate further and pass it to the final notes builder, which decides on their naming.
// func (ti *templateInstance) getTemplateNote(note *Note) *templateNote {
// 	if ti.templateNote == nil {
// 		return nil
// 	}

// 	currentTemplateNote := ti.templateNote
// 	if currentTemplateNote.Equal(note) {
// 		return currentTemplateNote
// 	}

// 	// the condition that it will not endlessly drive in a circle in search of a non-existent note
// 	for i := 2; i <= int(HalftonesInOctave); i++ {
// 		currentTemplateNote = currentTemplateNote.getNext()
// 		if currentTemplateNote.Equal(note) {
// 			return currentTemplateNote
// 		}
// 	}

// 	return nil
// }
