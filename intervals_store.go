package muse

type IntervalsStore map[IntervalName]func() *IntervalChromatic

func (is IntervalsStore) GetInterval(intervalName IntervalName) *IntervalChromatic {
	return is[intervalName]()
}

func (is IntervalsStore) AddInterval(interval *IntervalChromatic) {
	is[interval.Name()] = func() *IntervalChromatic { return interval }
	is[interval.ShortName()] = func() *IntervalChromatic { return interval }
}

func InitIntervalsStore() IntervalsStore {
	intervalsStore := make(IntervalsStore)

	// Acoustic intervals

	intervalsStore[IntervalNamePerfectUnison] = IntervalPerfectUnison
	intervalsStore[IntervalNameMinorSecond] = IntervalMinorSecond
	intervalsStore[IntervalNameMajorSecond] = IntervalMajorSecond
	intervalsStore[IntervalNameMinorThird] = IntervalMinorThird
	intervalsStore[IntervalNameMajorThird] = IntervalMajorThird
	intervalsStore[IntervalNamePerfectFourth] = IntervalPerfectFourth
	intervalsStore[IntervalNameTritone] = IntervalTritone
	intervalsStore[IntervalNamePerfectFifth] = IntervalPerfectFifth
	intervalsStore[IntervalNameMinorSixth] = IntervalMinorSixth
	intervalsStore[IntervalNameMajorSixth] = IntervalMajorSixth
	intervalsStore[IntervalNameMinorSeventh] = IntervalMinorSeventh
	intervalsStore[IntervalNameMajorSeventh] = IntervalMajorSeventh
	intervalsStore[IntervalNamePerfectOctave] = IntervalPerfectOctave

	intervalsStore[IntervalNamePerfectUnisonShort] = IntervalPerfectUnison
	intervalsStore[IntervalNameMinorSecondShort] = IntervalMinorSecond
	intervalsStore[IntervalNameMajorSecondShort] = IntervalMajorSecond
	intervalsStore[IntervalNameMinorThirdShort] = IntervalMinorThird
	intervalsStore[IntervalNameMajorThirdShort] = IntervalMajorThird
	intervalsStore[IntervalNamePerfectFourthShort] = IntervalPerfectFourth
	intervalsStore[IntervalNameTritoneShort] = IntervalTritone
	intervalsStore[IntervalNamePerfectFifthShort] = IntervalPerfectFifth
	intervalsStore[IntervalNameMinorSixthShort] = IntervalMinorSixth
	intervalsStore[IntervalNameMajorSixthShort] = IntervalMajorSixth
	intervalsStore[IntervalNameMinorSeventhShort] = IntervalMinorSeventh
	intervalsStore[IntervalNameMajorSeventhShort] = IntervalMajorSeventh
	intervalsStore[IntervalNamePerfectOctaveShort] = IntervalPerfectOctave

	// Mode's intervals

	intervalsStore[IntervalNameDiminishedSecond] = IntervalDiminishedSecond
	intervalsStore[IntervalNameAugmentedUnison] = IntervalAugmentedUnison
	intervalsStore[IntervalNameDiminishedThird] = IntervalDiminishedThird
	intervalsStore[IntervalNameAugmentedSecond] = IntervalAugmentedSecond
	intervalsStore[IntervalNameDiminishedFourth] = IntervalDiminishedFourth
	intervalsStore[IntervalNameAugmentedThird] = IntervalAugmentedThird
	intervalsStore[IntervalNameDiminishedFifth] = IntervalDiminishedFifth
	intervalsStore[IntervalNameAugmentedFourth] = IntervalAugmentedFourth
	intervalsStore[IntervalNameDiminishedSixth] = IntervalDiminishedSixth
	intervalsStore[IntervalNameAugmentedFifth] = IntervalAugmentedFifth
	intervalsStore[IntervalNameDiminishedSeventh] = IntervalDiminishedSeventh
	intervalsStore[IntervalNameAugmentedSixth] = IntervalAugmentedSixth
	intervalsStore[IntervalNameDiminishedOctave] = IntervalDiminishedOctave
	intervalsStore[IntervalNameAugmentedSeventh] = IntervalAugmentedSeventh

	intervalsStore[IntervalNameDiminishedSecondShort] = IntervalDiminishedSecond
	intervalsStore[IntervalNameAugmentedUnisonShort] = IntervalAugmentedUnison
	intervalsStore[IntervalNameDiminishedThirdShort] = IntervalDiminishedThird
	intervalsStore[IntervalNameAugmentedSecondShort] = IntervalAugmentedSecond
	intervalsStore[IntervalNameDiminishedFourthShort] = IntervalDiminishedFourth
	intervalsStore[IntervalNameAugmentedThirdShort] = IntervalAugmentedThird
	intervalsStore[IntervalNameDiminishedFifthShort] = IntervalDiminishedFifth
	intervalsStore[IntervalNameAugmentedFourthShort] = IntervalAugmentedFourth
	intervalsStore[IntervalNameDiminishedSixthShort] = IntervalDiminishedSixth
	intervalsStore[IntervalNameAugmentedFifthShort] = IntervalAugmentedFifth
	intervalsStore[IntervalNameDiminishedSeventhShort] = IntervalDiminishedSeventh
	intervalsStore[IntervalNameAugmentedSixthShort] = IntervalAugmentedSixth
	intervalsStore[IntervalNameDiminishedOctaveShort] = IntervalDiminishedOctave
	intervalsStore[IntervalNameAugmentedSeventhShort] = IntervalAugmentedSeventh

	return intervalsStore
}
