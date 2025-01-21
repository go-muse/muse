package interval

// Hindemith Grading of Sonances
// The most consonant intervals have more "weight"

type Sonance uint8

const (
	SonanceMinorSecond = Sonance(iota + 1)
	SonanceMajorSeventh
	SonanceTritone
	SonanceMajorSecond
	SonanceMinorSeventh
	SonanceMinorThird
	SonanceMajorSixth
	SonanceMajorThird
	SonanceMinorSixth
	SonancePerfectFourth
	SonancePerfectFifth
	SonancePerfectOctave
	SonancePerfectUnison
)

// TODO: выходящие за пределы октавы интервалы
