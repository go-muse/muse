package muse

type IntervalsStore map[IntervalName]func() *ChromaticInterval

func (is IntervalsStore) GetInterval(intervalName IntervalName) *ChromaticInterval {
	return is[intervalName]()
}

func (is IntervalsStore) AddInterval(interval *ChromaticInterval) {
	is[interval.Name()] = func() *ChromaticInterval { return interval }
	is[interval.ShortName()] = func() *ChromaticInterval { return interval }
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

	intervalsStore[IntervalNameModifiedDiminishedSecond] = IntervalDiminishedSecond
	intervalsStore[IntervalNameModifiedAugmentedUnison] = IntervalAugmentedUnison
	intervalsStore[IntervalNameModifiedDiminishedThird] = IntervalDiminishedThird
	intervalsStore[IntervalNameModifiedAugmentedSecond] = IntervalAugmentedSecond
	intervalsStore[IntervalNameModifiedDiminishedFourth] = IntervalDiminishedFourth
	intervalsStore[IntervalNameModifiedAugmentedThird] = IntervalAugmentedThird
	intervalsStore[IntervalNameModifiedDiminishedFifth] = IntervalDiminishedFifth
	intervalsStore[IntervalNameModifiedAugmentedFourth] = IntervalAugmentedFourth
	intervalsStore[IntervalNameModifiedDiminishedSixth] = IntervalDiminishedSixth
	intervalsStore[IntervalNameModifiedAugmentedFifth] = IntervalAugmentedFifth
	intervalsStore[IntervalNameModifiedDiminishedSeventh] = IntervalDiminishedSeventh
	intervalsStore[IntervalNameModifiedAugmentedSixth] = IntervalAugmentedSixth
	intervalsStore[IntervalNameModifiedDiminishedOctave] = IntervalDiminishedOctave
	intervalsStore[IntervalNameModifiedAugmentedSeventh] = IntervalAugmentedSeventh

	intervalsStore[IntervalNameModifiedDiminishedSecondShort] = IntervalDiminishedSecond
	intervalsStore[IntervalNameModifiedAugmentedUnisonShort] = IntervalAugmentedUnison
	intervalsStore[IntervalNameModifiedDiminishedThirdShort] = IntervalDiminishedThird
	intervalsStore[IntervalNameModifiedAugmentedSecondShort] = IntervalAugmentedSecond
	intervalsStore[IntervalNameModifiedDiminishedFourthShort] = IntervalDiminishedFourth
	intervalsStore[IntervalNameModifiedAugmentedThirdShort] = IntervalAugmentedThird
	intervalsStore[IntervalNameModifiedDiminishedFifthShort] = IntervalDiminishedFifth
	intervalsStore[IntervalNameModifiedAugmentedFourthShort] = IntervalAugmentedFourth
	intervalsStore[IntervalNameModifiedDiminishedSixthShort] = IntervalDiminishedSixth
	intervalsStore[IntervalNameModifiedAugmentedFifthShort] = IntervalAugmentedFifth
	intervalsStore[IntervalNameModifiedDiminishedSeventhShort] = IntervalDiminishedSeventh
	intervalsStore[IntervalNameModifiedAugmentedSixthShort] = IntervalAugmentedSixth
	intervalsStore[IntervalNameModifiedDiminishedOctaveShort] = IntervalDiminishedOctave
	intervalsStore[IntervalNameModifiedAugmentedSeventhShort] = IntervalAugmentedSeventh

	return intervalsStore
}
