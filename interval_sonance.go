package muse

// Hindemith Grading of Sonances
// The most consonant intervals have more "weight"

type Sonance uint8

const (
	IntervalSonanceMinorSecond = Sonance(iota + 1)
	IntervalSonanceMajorSeventh
	IntervalSonanceTritone
	IntervalSonanceMajorSecond
	IntervalSonanceMinorSeventh
	IntervalSonanceMinorThird
	IntervalSonanceMajorSixth
	IntervalSonanceMajorThird
	IntervalSonanceMinorSixth
	IntervalSonancePerfectFourth
	IntervalSonancePerfectFifth
	IntervalSonancePerfectOctave
	IntervalSonancePerfectUnison
)

// TODO: выходящие за пределы октавы интервалы
