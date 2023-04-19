package muse

type AlterationSymbol string

const (
	AlterSymbolFlat  = AlterationSymbol("b")
	AlterSymbolSharp = AlterationSymbol("#")
)

func (as AlterationSymbol) String() string {
	return string(as)
}
