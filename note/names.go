//nolint:gochecknoglobals // Predefined note names are intentionally global for convenient usage
package note

// Names is a common name for set of the names.
type Names []Name

// Length returns amount of names in the Names.
func (ns Names) Length() uint64 {
	return uint64(len(ns))
}

// Predefined note names for convenient note creation.
// Usage: note.C.NewNote(), note.CSHARP.NewNote(), etc.

// Natural notes (no accidentals).
var (
	C = NewName(LetterC, Natural())
	D = NewName(LetterD, Natural())
	E = NewName(LetterE, Natural())
	F = NewName(LetterF, Natural())
	G = NewName(LetterG, Natural())
	A = NewName(LetterA, Natural())
	B = NewName(LetterB, Natural())
)

// Sharp notes (single sharp).
var (
	CSHARP = NewName(LetterC, Sharp())
	DSHARP = NewName(LetterD, Sharp())
	ESHARP = NewName(LetterE, Sharp())
	FSHARP = NewName(LetterF, Sharp())
	GSHARP = NewName(LetterG, Sharp())
	ASHARP = NewName(LetterA, Sharp())
	BSHARP = NewName(LetterB, Sharp())
)

// Flat notes (single flat).
var (
	CFLAT = NewName(LetterC, Flat())
	DFLAT = NewName(LetterD, Flat())
	EFLAT = NewName(LetterE, Flat())
	FFLAT = NewName(LetterF, Flat())
	GFLAT = NewName(LetterG, Flat())
	AFLAT = NewName(LetterA, Flat())
	BFLAT = NewName(LetterB, Flat())
)

// Double-sharp notes.
var (
	CSHARP2 = NewName(LetterC, DoubleSharp())
	DSHARP2 = NewName(LetterD, DoubleSharp())
	ESHARP2 = NewName(LetterE, DoubleSharp())
	FSHARP2 = NewName(LetterF, DoubleSharp())
	GSHARP2 = NewName(LetterG, DoubleSharp())
	ASHARP2 = NewName(LetterA, DoubleSharp())
	BSHARP2 = NewName(LetterB, DoubleSharp())
)

// Double-flat notes.
var (
	CFLAT2 = NewName(LetterC, DoubleFlat())
	DFLAT2 = NewName(LetterD, DoubleFlat())
	EFLAT2 = NewName(LetterE, DoubleFlat())
	FFLAT2 = NewName(LetterF, DoubleFlat())
	GFLAT2 = NewName(LetterG, DoubleFlat())
	AFLAT2 = NewName(LetterA, DoubleFlat())
	BFLAT2 = NewName(LetterB, DoubleFlat())
)

func GetChromaticNamesSharp() []Name {
	return []Name{C, CSHARP, D, DSHARP, E, F, FSHARP, G, GSHARP, A, ASHARP, B}
}

func GetChromaticNamesFlat() []Name {
	return []Name{C, DFLAT, D, EFLAT, E, F, GFLAT, G, AFLAT, A, BFLAT, B}
}
