package muse

const (
	// Chromatic intervals.

	IntervalNamePerfectUnison = IntervalName("PerfectUnison")
	IntervalNameMinorSecond   = IntervalName("MinorSecond")
	IntervalNameMajorSecond   = IntervalName("MajorSecond")
	IntervalNameMinorThird    = IntervalName("MinorThird")
	IntervalNameMajorThird    = IntervalName("MajorThird")
	IntervalNamePerfectFourth = IntervalName("PerfectFourth")
	IntervalNameTritone       = IntervalName("Tritone")
	IntervalNamePerfectFifth  = IntervalName("PerfectFifth")
	IntervalNameMinorSixth    = IntervalName("MinorSixth")
	IntervalNameMajorSixth    = IntervalName("MajorSixth")
	IntervalNameMinorSeventh  = IntervalName("MinorSeventh")
	IntervalNameMajorSeventh  = IntervalName("MajorSeventh")
	IntervalNamePerfectOctave = IntervalName("PerfectOctave")

	IntervalNameMinorNinth       = IntervalName("MinorNinth")
	IntervalNameMajorNinth       = IntervalName("MajorNinth")
	IntervalNameMinorTenth       = IntervalName("MinorTenth")
	IntervalNameMajorTenth       = IntervalName("MajorTenth")
	IntervalNamePerfectEleventh  = IntervalName("PerfectEleventh")
	IntervalNamePerfectTwelfth   = IntervalName("PerfectTwelfth") // Tritave
	IntervalNameMinorThirteenth  = IntervalName("MinorThirteenth")
	IntervalNameMajorThirteenth  = IntervalName("MajorThirteenth")
	IntervalNameMinorFourteenth  = IntervalName("MinorFourteenth")
	IntervalNameMajorFourteenth  = IntervalName("MajorFourteenth")
	IntervalNamePerfectFifteenth = IntervalName("PerfectFifteenth ") // Double octave

	IntervalNamePerfectUnisonShort = IntervalName("P1")
	IntervalNameMinorSecondShort   = IntervalName("m2")
	IntervalNameMajorSecondShort   = IntervalName("M2")
	IntervalNameMinorThirdShort    = IntervalName("m3")
	IntervalNameMajorThirdShort    = IntervalName("M3")
	IntervalNamePerfectFourthShort = IntervalName("P4")
	IntervalNameTritoneShort       = IntervalName("TT")
	IntervalNamePerfectFifthShort  = IntervalName("P5")
	IntervalNameMinorSixthShort    = IntervalName("m6")
	IntervalNameMajorSixthShort    = IntervalName("M6")
	IntervalNameMinorSeventhShort  = IntervalName("m7")
	IntervalNameMajorSeventhShort  = IntervalName("M7")
	IntervalNamePerfectOctaveShort = IntervalName("P8")

	IntervalNameMinorNinthShort       = IntervalName("m9")
	IntervalNameMajorNinthShort       = IntervalName("M9")
	IntervalNameMinorTenthShort       = IntervalName("m10")
	IntervalNameMajorTenthShort       = IntervalName("M10")
	IntervalNamePerfectEleventhShort  = IntervalName("P11")
	IntervalNamePerfectTwelfthShort   = IntervalName("P12")
	IntervalNameMinorThirteenthShort  = IntervalName("m13")
	IntervalNameMajorThirteenthShort  = IntervalName("M13")
	IntervalNameMinorFourteenthShort  = IntervalName("m14")
	IntervalNameMajorFourteenthShort  = IntervalName("M14")
	IntervalNamePerfectFifteenthShort = IntervalName("P15")

	// Diatonic intervals.

	IntervalNameModifiedDiminishedSecond  = IntervalName("DiminishedSecond")
	IntervalNameModifiedAugmentedUnison   = IntervalName("AugmentedUnison")
	IntervalNameModifiedDiminishedThird   = IntervalName("DiminishedThird")
	IntervalNameModifiedAugmentedSecond   = IntervalName("AugmentedSecond")
	IntervalNameModifiedDiminishedFourth  = IntervalName("DiminishedFourth")
	IntervalNameModifiedAugmentedThird    = IntervalName("AugmentedThird")
	IntervalNameModifiedDiminishedFifth   = IntervalName("DiminishedFifth")
	IntervalNameModifiedAugmentedFourth   = IntervalName("AugmentedFourth")
	IntervalNameModifiedDiminishedSixth   = IntervalName("DiminishedSixth")
	IntervalNameModifiedAugmentedFifth    = IntervalName("AugmentedFifth")
	IntervalNameModifiedDiminishedSeventh = IntervalName("DiminishedSeventh")
	IntervalNameModifiedAugmentedSixth    = IntervalName("AugmentedSixth")
	IntervalNameModifiedDiminishedOctave  = IntervalName("DiminishedOctave")
	IntervalNameModifiedAugmentedSeventh  = IntervalName("AugmentedSeventh")

	IntervalNameModifiedDiminishedNinth      = IntervalName("DiminishedNinth")
	IntervalNameModifiedAugmentedOctave      = IntervalName("AugmentedOctave")
	IntervalNameModifiedDiminishedTenth      = IntervalName("DiminishedTenth")
	IntervalNameModifiedAugmentedNinth       = IntervalName("AugmentedNinth")
	IntervalNameModifiedDiminishedEleventh   = IntervalName("DiminishedEleventh")
	IntervalNameModifiedAugmentedTenth       = IntervalName("AugmentedTenth")
	IntervalNameModifiedDiminishedTwelfth    = IntervalName("DiminishedTwelfth")
	IntervalNameModifiedAugmentedEleventh    = IntervalName("AugmentedEleventh")
	IntervalNameModifiedDiminishedThirteenth = IntervalName("DiminishedThirteenth")
	IntervalNameModifiedAugmentedTwelfth     = IntervalName("AugmentedTwelfth")
	IntervalNameModifiedDiminishedFourteenth = IntervalName("DiminishedFourteenth")
	IntervalNameModifiedAugmentedThirteenth  = IntervalName("AugmentedThirteenth")
	IntervalNameModifiedDiminishedFifteenth  = IntervalName("DiminishedFifteenth")
	IntervalNameModifiedAugmentedFourteenth  = IntervalName("AugmentedFourteenth") // Double octave
	IntervalNameModifiedAugmentedFifteenth   = IntervalName("AugmentedFifteenth")

	IntervalNameModifiedDiminishedSecondShort  = IntervalName("d2")
	IntervalNameModifiedAugmentedUnisonShort   = IntervalName("A1")
	IntervalNameModifiedDiminishedThirdShort   = IntervalName("d3")
	IntervalNameModifiedAugmentedSecondShort   = IntervalName("A2")
	IntervalNameModifiedDiminishedFourthShort  = IntervalName("d4")
	IntervalNameModifiedAugmentedThirdShort    = IntervalName("A3")
	IntervalNameModifiedDiminishedFifthShort   = IntervalName("d5")
	IntervalNameModifiedAugmentedFourthShort   = IntervalName("A4")
	IntervalNameModifiedDiminishedSixthShort   = IntervalName("d6")
	IntervalNameModifiedAugmentedFifthShort    = IntervalName("A5")
	IntervalNameModifiedDiminishedSeventhShort = IntervalName("d7")
	IntervalNameModifiedAugmentedSixthShort    = IntervalName("A6")
	IntervalNameModifiedDiminishedOctaveShort  = IntervalName("d8")
	IntervalNameModifiedAugmentedSeventhShort  = IntervalName("A7")

	IntervalNameModifiedDiminishedNinthShort      = IntervalName("d9")
	IntervalNameModifiedAugmentedOctaveShort      = IntervalName("A8")
	IntervalNameModifiedDiminishedTenthShort      = IntervalName("d10")
	IntervalNameModifiedAugmentedNinthShort       = IntervalName("A9")
	IntervalNameModifiedDiminishedEleventhShort   = IntervalName("d11")
	IntervalNameModifiedAugmentedTenthShort       = IntervalName("A10")
	IntervalNameModifiedDiminishedTwelfthShort    = IntervalName("d12")
	IntervalNameModifiedAugmentedEleventhShort    = IntervalName("A11")
	IntervalNameModifiedDiminishedThirteenthShort = IntervalName("d13")
	IntervalNameModifiedAugmentedTwelfthShort     = IntervalName("A12")
	IntervalNameModifiedDiminishedFourteenthShort = IntervalName("d14")
	IntervalNameModifiedAugmentedThirteenthShort  = IntervalName("A13")
	IntervalNameModifiedDiminishedFifteenthShort  = IntervalName("d15")
	IntervalNameModifiedAugmentedFourteenthShort  = IntervalName("A14")
	IntervalNameModifiedAugmentedFifteenthShort   = IntervalName("A15")
)

const (
	IntervalHalfTones0  = HalfTones(0)
	IntervalHalfTones1  = HalfTones(1)
	IntervalHalfTones2  = HalfTones(2)
	IntervalHalfTones3  = HalfTones(3)
	IntervalHalfTones4  = HalfTones(4)
	IntervalHalfTones5  = HalfTones(5)
	IntervalHalfTones6  = HalfTones(6)
	IntervalHalfTones7  = HalfTones(7)
	IntervalHalfTones8  = HalfTones(8)
	IntervalHalfTones9  = HalfTones(9)
	IntervalHalfTones10 = HalfTones(10)
	IntervalHalfTones11 = HalfTones(11)
	IntervalHalfTones12 = HalfTones(12)
	IntervalHalfTones13 = HalfTones(13)
	IntervalHalfTones14 = HalfTones(14)
	IntervalHalfTones15 = HalfTones(15)
	IntervalHalfTones16 = HalfTones(16)
	IntervalHalfTones17 = HalfTones(17)
	IntervalHalfTones18 = HalfTones(18)
	IntervalHalfTones19 = HalfTones(19)
	IntervalHalfTones20 = HalfTones(20)
	IntervalHalfTones21 = HalfTones(21)
	IntervalHalfTones22 = HalfTones(22)
	IntervalHalfTones23 = HalfTones(23)
	IntervalHalfTones24 = HalfTones(24)
)

// Chromatic intervals.

func IntervalPerfectUnison() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectUnison,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectUnison,
			shortName: IntervalNamePerfectUnisonShort,
		},
		halfTones: IntervalHalfTones0,
	}
}

func IntervalMinorSecond() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameMinorSecond,
			shortName: IntervalNameMinorSecondShort,
		},
		halfTones: IntervalHalfTones1,
	}
}

func IntervalMajorSecond() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameMajorSecond,
			shortName: IntervalNameMajorSecondShort,
		},
		halfTones: IntervalHalfTones2,
	}
}

func IntervalMinorThird() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorThird,
		names: &intervalNameExtended{
			name:      IntervalNameMinorThird,
			shortName: IntervalNameMinorThirdShort,
		},
		halfTones: IntervalHalfTones3,
	}
}

func IntervalMajorThird() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorThird,
		names: &intervalNameExtended{
			name:      IntervalNameMajorThird,
			shortName: IntervalNameMajorThirdShort,
		},
		halfTones: IntervalHalfTones4,
	}
}

func IntervalPerfectFourth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectFourth,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectFourth,
			shortName: IntervalNamePerfectFourthShort,
		},
		halfTones: IntervalHalfTones5,
	}
}

func IntervalTritone() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceTritone,
		names: &intervalNameExtended{
			name:      IntervalNameTritone,
			shortName: IntervalNameTritoneShort,
		},
		halfTones: IntervalHalfTones6,
	}
}

func IntervalPerfectFifth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectFifth,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectFifth,
			shortName: IntervalNamePerfectFifthShort,
		},
		halfTones: IntervalHalfTones7,
	}
}

func IntervalMinorSixth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameMinorSixth,
			shortName: IntervalNameMinorSixthShort,
		},
		halfTones: IntervalHalfTones8,
	}
}

func IntervalMajorSixth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameMajorSixth,
			shortName: IntervalNameMajorSixthShort,
		},
		halfTones: IntervalHalfTones9,
	}
}

func IntervalMinorSeventh() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameMinorSeventh,
			shortName: IntervalNameMinorSeventhShort,
		},
		halfTones: IntervalHalfTones10,
	}
}

func IntervalMajorSeventh() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameMajorSeventh,
			shortName: IntervalNameMajorSeventhShort,
		},
		halfTones: IntervalHalfTones11,
	}
}

func IntervalPerfectOctave() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectOctave,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectOctave,
			shortName: IntervalNamePerfectOctaveShort,
		},
		halfTones: IntervalHalfTones12,
	}
}

func IntervalMinorNinth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorNinth,
			shortName: IntervalNameMinorNinthShort,
		},
		halfTones: IntervalHalfTones13,
	}
}

func IntervalMajorNinth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorNinth,
			shortName: IntervalNameMajorNinthShort,
		},
		halfTones: IntervalHalfTones14,
	}
}

func IntervalMinorTenth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorTenth,
			shortName: IntervalNameMinorTenthShort,
		},
		halfTones: IntervalHalfTones15,
	}
}

func IntervalMajorTenth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorTenth,
			shortName: IntervalNameMajorTenthShort,
		},
		halfTones: IntervalHalfTones16,
	}
}

func IntervalPerfectEleventh() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectEleventh,
			shortName: IntervalNamePerfectEleventhShort,
		},
		halfTones: IntervalHalfTones17,
	}
}

// No chromatic name for 18 halftones

func IntervalPerfectTwelfth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectTwelfth,
			shortName: IntervalNamePerfectTwelfthShort,
		},
		halfTones: IntervalHalfTones19,
	}
}

func IntervalMinorThirteenth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorThirteenth,
			shortName: IntervalNameMinorThirteenthShort,
		},
		halfTones: IntervalHalfTones20,
	}
}

func IntervalMajorThirteenth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorThirteenth,
			shortName: IntervalNameMajorThirteenthShort,
		},
		halfTones: IntervalHalfTones21,
	}
}

func IntervalMinorFourteenth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorFourteenth,
			shortName: IntervalNameMinorFourteenthShort,
		},
		halfTones: IntervalHalfTones22,
	}
}

func IntervalMajorFourteenth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorFourteenth,
			shortName: IntervalNameMajorFourteenthShort,
		},
		halfTones: IntervalHalfTones23,
	}
}

func IntervalPerfectFifteenth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectFifteenth,
			shortName: IntervalNamePerfectFifteenthShort,
		},
		halfTones: IntervalHalfTones24,
	}
}

// Compound chromatic (diatonic) intervals

func IntervalDiminishedSecond() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectUnison,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedDiminishedSecond,
			shortName: IntervalNameModifiedDiminishedSecondShort,
		},
		halfTones: IntervalHalfTones0,
	}
}

func IntervalAugmentedUnison() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedAugmentedUnison,
			shortName: IntervalNameModifiedAugmentedUnisonShort,
		},
		halfTones: IntervalHalfTones1,
	}
}

func IntervalDiminishedThird() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedDiminishedThird,
			shortName: IntervalNameModifiedDiminishedThirdShort,
		},
		halfTones: IntervalHalfTones2,
	}
}

func IntervalAugmentedSecond() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorThird,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedAugmentedSecond,
			shortName: IntervalNameModifiedAugmentedSecondShort,
		},
		halfTones: IntervalHalfTones3,
	}
}

func IntervalDiminishedFourth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorThird,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedDiminishedFourth,
			shortName: IntervalNameModifiedDiminishedFourthShort,
		},
		halfTones: IntervalHalfTones4,
	}
}

func IntervalAugmentedThird() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectFourth,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedAugmentedThird,
			shortName: IntervalNameModifiedAugmentedThirdShort,
		},
		halfTones: IntervalHalfTones5,
	}
}

func IntervalDiminishedFifth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceTritone,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedDiminishedFifth,
			shortName: IntervalNameModifiedDiminishedFifthShort,
		},
		halfTones: IntervalHalfTones6,
	}
}

func IntervalAugmentedFourth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceTritone,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedAugmentedFourth,
			shortName: IntervalNameModifiedAugmentedFourthShort,
		},
		halfTones: IntervalHalfTones6,
	}
}

func IntervalDiminishedSixth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectFifth,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedDiminishedSixth,
			shortName: IntervalNameModifiedDiminishedSixthShort,
		},
		halfTones: IntervalHalfTones7,
	}
}

func IntervalAugmentedFifth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedAugmentedFifth,
			shortName: IntervalNameModifiedAugmentedFifthShort,
		},
		halfTones: IntervalHalfTones8,
	}
}

func IntervalDiminishedSeventh() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedDiminishedSeventh,
			shortName: IntervalNameModifiedDiminishedSeventhShort,
		},
		halfTones: IntervalHalfTones9,
	}
}

func IntervalAugmentedSixth() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMinorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedAugmentedSixth,
			shortName: IntervalNameModifiedAugmentedSixthShort,
		},
		halfTones: IntervalHalfTones10,
	}
}

func IntervalDiminishedOctave() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonanceMajorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedDiminishedOctave,
			shortName: IntervalNameModifiedDiminishedOctaveShort,
		},
		halfTones: IntervalHalfTones11,
	}
}

func IntervalAugmentedSeventh() *ChromaticInterval {
	return &ChromaticInterval{
		Sonance: IntervalSonancePerfectOctave,
		names: &intervalNameExtended{
			name:      IntervalNameModifiedAugmentedSeventh,
			shortName: IntervalNameModifiedAugmentedSeventhShort,
		},
		halfTones: IntervalHalfTones12,
	}
}
