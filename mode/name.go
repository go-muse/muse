package mode

import "github.com/go-muse/muse/note"

// Name is just a name for a mode.
type Name string

// MakeNewMode makes a new mode with the mode name and given note name.
func (n Name) MakeNewMode(noteName note.Name) (*Mode, error) {
	return MakeNewMode(n, noteName)
}

// MustMakeNewMode makes a new mode with the mode name and given note name. If an error occurs, it panics.
func (n Name) MustMakeNewMode(noteName note.Name) *Mode {
	mode, err := MakeNewMode(n, noteName)
	if err != nil {
		panic(err)
	}

	return mode
}
