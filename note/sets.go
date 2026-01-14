package note

// GetSetFullChromatic returns all possible notes of the tonal system within one alteration.
func GetSetFullChromatic() []Note {
	return []Note{
		New(C),
		New(DFLAT),
		New(CSHARP),
		New(D),
		New(EFLAT),
		New(DSHARP),
		New(E),
		New(F),
		New(GFLAT),
		New(FSHARP),
		New(G),
		New(AFLAT),
		New(GSHARP),
		New(A),
		New(BFLAT),
		New(ASHARP),
		New(B),
	}
}

// GetSetFullChromaticDoubleAltered returns all possible notes of the tonal system within two alterations.
func GetSetFullChromaticDoubleAltered() []Note {
	return []Note{
		New(C),
		New(DFLAT2),
		New(DFLAT),
		New(CSHARP),
		New(CSHARP2),
		New(D),
		New(EFLAT2),
		New(EFLAT),
		New(DSHARP),
		New(DSHARP2),
		New(E),
		New(F),
		New(GFLAT2),
		New(GFLAT),
		New(FSHARP),
		New(FSHARP2),
		New(G),
		New(AFLAT2),
		New(AFLAT),
		New(GSHARP),
		New(GSHARP2),
		New(A),
		New(BFLAT2),
		New(BFLAT),
		New(ASHARP),
		New(ASHARP2),
		New(B),
	}
}
