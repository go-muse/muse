package muse

// NoteName is a common note for the note.
type NoteName string

// String is stringer for NoteName type.
func (nn NoteName) String() string {
	return string(nn)
}
