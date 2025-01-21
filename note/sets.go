package note

// GetSetFullChromatic returns all possible notes of the tonal system within one alteration.
func GetSetFullChromatic() []*Note {
	return []*Note{
		MustNewNote(C),
		MustNewNote(DFLAT),
		MustNewNote(CSHARP),
		MustNewNote(D),
		MustNewNote(EFLAT),
		MustNewNote(DSHARP),
		MustNewNote(E),
		MustNewNote(F),
		MustNewNote(GFLAT),
		MustNewNote(FSHARP),
		MustNewNote(G),
		MustNewNote(AFLAT),
		MustNewNote(GSHARP),
		MustNewNote(A),
		MustNewNote(BFLAT),
		MustNewNote(ASHARP),
		MustNewNote(B),
	}
}

// GetSetFullChromaticDoubleAltered returns all possible notes of the tonal system within two alterations.
func GetSetFullChromaticDoubleAltered() []*Note {
	return []*Note{
		MustNewNote(C),
		MustNewNote(DFLAT2),
		MustNewNote(DFLAT),
		MustNewNote(CSHARP),
		MustNewNote(CSHARP2),
		MustNewNote(D),
		MustNewNote(EFLAT2),
		MustNewNote(EFLAT),
		MustNewNote(DSHARP),
		MustNewNote(DSHARP2),
		MustNewNote(E),
		MustNewNote(F),
		MustNewNote(GFLAT2),
		MustNewNote(GFLAT),
		MustNewNote(FSHARP),
		MustNewNote(FSHARP2),
		MustNewNote(G),
		MustNewNote(AFLAT2),
		MustNewNote(AFLAT),
		MustNewNote(GSHARP),
		MustNewNote(GSHARP2),
		MustNewNote(A),
		MustNewNote(BFLAT2),
		MustNewNote(BFLAT),
		MustNewNote(ASHARP),
		MustNewNote(ASHARP2),
		MustNewNote(B),
	}
}
