package interval

// Store is a storage of chromatic intervals that can be accessed by their name.
type Store map[Name]func() *Chromatic

// GetInterval returns chromatic interval by its name.
func (s Store) GetInterval(intervalName Name) *Chromatic {
	return s[intervalName]()
}

// AddInterval adds a new chromatic interval to the store.
func (s Store) AddInterval(interval *Chromatic) {
	s[interval.Name()] = func() *Chromatic { return interval }
	s[interval.ShortName()] = func() *Chromatic { return interval }
}

// InitIntervalsStore returns a new initialized intervals store with chromatic intervals.
func InitIntervalsStore() Store {
	intervalsStore := make(Store)

	// Acoustic intervals

	intervalsStore[NamePerfectUnison] = PerfectUnison
	intervalsStore[NameMinorSecond] = MinorSecond
	intervalsStore[NameMajorSecond] = MajorSecond
	intervalsStore[NameMinorThird] = MinorThird
	intervalsStore[NameMajorThird] = MajorThird
	intervalsStore[NamePerfectFourth] = PerfectFourth
	intervalsStore[NameTritone] = Tritone
	intervalsStore[NamePerfectFifth] = PerfectFifth
	intervalsStore[NameMinorSixth] = MinorSixth
	intervalsStore[NameMajorSixth] = MajorSixth
	intervalsStore[NameMinorSeventh] = MinorSeventh
	intervalsStore[NameMajorSeventh] = MajorSeventh
	intervalsStore[NamePerfectOctave] = PerfectOctave

	intervalsStore[NamePerfectUnisonShort] = PerfectUnison
	intervalsStore[NameMinorSecondShort] = MinorSecond
	intervalsStore[NameMajorSecondShort] = MajorSecond
	intervalsStore[NameMinorThirdShort] = MinorThird
	intervalsStore[NameMajorThirdShort] = MajorThird
	intervalsStore[NamePerfectFourthShort] = PerfectFourth
	intervalsStore[NameTritoneShort] = Tritone
	intervalsStore[NamePerfectFifthShort] = PerfectFifth
	intervalsStore[NameMinorSixthShort] = MinorSixth
	intervalsStore[NameMajorSixthShort] = MajorSixth
	intervalsStore[NameMinorSeventhShort] = MinorSeventh
	intervalsStore[NameMajorSeventhShort] = MajorSeventh
	intervalsStore[NamePerfectOctaveShort] = PerfectOctave

	// Mode's intervals

	intervalsStore[NameDiminishedSecond] = DiminishedSecond
	intervalsStore[NameAugmentedUnison] = AugmentedUnison
	intervalsStore[NameDiminishedThird] = DiminishedThird
	intervalsStore[NameAugmentedSecond] = AugmentedSecond
	intervalsStore[NameDiminishedFourth] = DiminishedFourth
	intervalsStore[NameAugmentedThird] = AugmentedThird
	intervalsStore[NameDiminishedFifth] = DiminishedFifth
	intervalsStore[NameAugmentedFourth] = AugmentedFourth
	intervalsStore[NameDiminishedSixth] = DiminishedSixth
	intervalsStore[NameAugmentedFifth] = AugmentedFifth
	intervalsStore[NameDiminishedSeventh] = DiminishedSeventh
	intervalsStore[NameAugmentedSixth] = AugmentedSixth
	intervalsStore[NameDiminishedOctave] = DiminishedOctave
	intervalsStore[NameAugmentedSeventh] = AugmentedSeventh

	intervalsStore[NameDiminishedSecondShort] = DiminishedSecond
	intervalsStore[NameAugmentedUnisonShort] = AugmentedUnison
	intervalsStore[NameDiminishedThirdShort] = DiminishedThird
	intervalsStore[NameAugmentedSecondShort] = AugmentedSecond
	intervalsStore[NameDiminishedFourthShort] = DiminishedFourth
	intervalsStore[NameAugmentedThirdShort] = AugmentedThird
	intervalsStore[NameDiminishedFifthShort] = DiminishedFifth
	intervalsStore[NameAugmentedFourthShort] = AugmentedFourth
	intervalsStore[NameDiminishedSixthShort] = DiminishedSixth
	intervalsStore[NameAugmentedFifthShort] = AugmentedFifth
	intervalsStore[NameDiminishedSeventhShort] = DiminishedSeventh
	intervalsStore[NameAugmentedSixthShort] = AugmentedSixth
	intervalsStore[NameDiminishedOctaveShort] = DiminishedOctave
	intervalsStore[NameAugmentedSeventhShort] = AugmentedSeventh

	return intervalsStore
}
