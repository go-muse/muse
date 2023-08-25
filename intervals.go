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

	IntervalNameMinorNinth        = IntervalName("MinorNinth")
	IntervalNameMajorNinth        = IntervalName("MajorNinth")
	IntervalNameMinorTenth        = IntervalName("MinorTenth")
	IntervalNameMajorTenth        = IntervalName("MajorTenth")
	IntervalNamePerfectEleventh   = IntervalName("PerfectEleventh")
	IntervalNameOctaveWithTritone = IntervalName("OctaveWithTritone") // no real name for tritone after octave
	IntervalNamePerfectTwelfth    = IntervalName("PerfectTwelfth")    // Tritave
	IntervalNameMinorThirteenth   = IntervalName("MinorThirteenth")
	IntervalNameMajorThirteenth   = IntervalName("MajorThirteenth")
	IntervalNameMinorFourteenth   = IntervalName("MinorFourteenth")
	IntervalNameMajorFourteenth   = IntervalName("MajorFourteenth")
	IntervalNamePerfectFifteenth  = IntervalName("PerfectFifteenth ") // Double octave

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

	IntervalNameMinorNinthShort        = IntervalName("m9")
	IntervalNameMajorNinthShort        = IntervalName("M9")
	IntervalNameMinorTenthShort        = IntervalName("m10")
	IntervalNameMajorTenthShort        = IntervalName("M10")
	IntervalNamePerfectEleventhShort   = IntervalName("P11")
	IntervalNameOctaveWithTritoneShort = IntervalName("OT")
	IntervalNamePerfectTwelfthShort    = IntervalName("P12")
	IntervalNameMinorThirteenthShort   = IntervalName("m13")
	IntervalNameMajorThirteenthShort   = IntervalName("M13")
	IntervalNameMinorFourteenthShort   = IntervalName("m14")
	IntervalNameMajorFourteenthShort   = IntervalName("M14")
	IntervalNamePerfectFifteenthShort  = IntervalName("P15")

	// Diatonic intervals.

	IntervalNameDiminishedSecond  = IntervalName("DiminishedSecond")
	IntervalNameAugmentedUnison   = IntervalName("AugmentedUnison")
	IntervalNameDiminishedThird   = IntervalName("DiminishedThird")
	IntervalNameAugmentedSecond   = IntervalName("AugmentedSecond")
	IntervalNameDiminishedFourth  = IntervalName("DiminishedFourth")
	IntervalNameAugmentedThird    = IntervalName("AugmentedThird")
	IntervalNameDiminishedFifth   = IntervalName("DiminishedFifth")
	IntervalNameAugmentedFourth   = IntervalName("AugmentedFourth")
	IntervalNameDiminishedSixth   = IntervalName("DiminishedSixth")
	IntervalNameAugmentedFifth    = IntervalName("AugmentedFifth")
	IntervalNameDiminishedSeventh = IntervalName("DiminishedSeventh")
	IntervalNameAugmentedSixth    = IntervalName("AugmentedSixth")
	IntervalNameDiminishedOctave  = IntervalName("DiminishedOctave")
	IntervalNameAugmentedSeventh  = IntervalName("AugmentedSeventh")

	IntervalNameDiminishedNinth      = IntervalName("DiminishedNinth")
	IntervalNameAugmentedOctave      = IntervalName("AugmentedOctave")
	IntervalNameDiminishedTenth      = IntervalName("DiminishedTenth")
	IntervalNameAugmentedNinth       = IntervalName("AugmentedNinth")
	IntervalNameDiminishedEleventh   = IntervalName("DiminishedEleventh")
	IntervalNameAugmentedTenth       = IntervalName("AugmentedTenth")
	IntervalNameDiminishedTwelfth    = IntervalName("DiminishedTwelfth")
	IntervalNameAugmentedEleventh    = IntervalName("AugmentedEleventh")
	IntervalNameDiminishedThirteenth = IntervalName("DiminishedThirteenth")
	IntervalNameAugmentedTwelfth     = IntervalName("AugmentedTwelfth")
	IntervalNameDiminishedFourteenth = IntervalName("DiminishedFourteenth")
	IntervalNameAugmentedThirteenth  = IntervalName("AugmentedThirteenth")
	IntervalNameDiminishedFifteenth  = IntervalName("DiminishedFifteenth")
	IntervalNameAugmentedFourteenth  = IntervalName("AugmentedFourteenth") // Double octave
	IntervalNameAugmentedFifteenth   = IntervalName("AugmentedFifteenth")

	IntervalNameDiminishedSecondShort  = IntervalName("d2")
	IntervalNameAugmentedUnisonShort   = IntervalName("A1")
	IntervalNameDiminishedThirdShort   = IntervalName("d3")
	IntervalNameAugmentedSecondShort   = IntervalName("A2")
	IntervalNameDiminishedFourthShort  = IntervalName("d4")
	IntervalNameAugmentedThirdShort    = IntervalName("A3")
	IntervalNameDiminishedFifthShort   = IntervalName("d5")
	IntervalNameAugmentedFourthShort   = IntervalName("A4")
	IntervalNameDiminishedSixthShort   = IntervalName("d6")
	IntervalNameAugmentedFifthShort    = IntervalName("A5")
	IntervalNameDiminishedSeventhShort = IntervalName("d7")
	IntervalNameAugmentedSixthShort    = IntervalName("A6")
	IntervalNameDiminishedOctaveShort  = IntervalName("d8")
	IntervalNameAugmentedSeventhShort  = IntervalName("A7")

	IntervalNameDiminishedNinthShort      = IntervalName("d9")
	IntervalNameAugmentedOctaveShort      = IntervalName("A8")
	IntervalNameDiminishedTenthShort      = IntervalName("d10")
	IntervalNameAugmentedNinthShort       = IntervalName("A9")
	IntervalNameDiminishedEleventhShort   = IntervalName("d11")
	IntervalNameAugmentedTenthShort       = IntervalName("A10")
	IntervalNameDiminishedTwelfthShort    = IntervalName("d12")
	IntervalNameAugmentedEleventhShort    = IntervalName("A11")
	IntervalNameDiminishedThirteenthShort = IntervalName("d13")
	IntervalNameAugmentedTwelfthShort     = IntervalName("A12")
	IntervalNameDiminishedFourteenthShort = IntervalName("d14")
	IntervalNameAugmentedThirteenthShort  = IntervalName("A13")
	IntervalNameDiminishedFifteenthShort  = IntervalName("d15")
	IntervalNameAugmentedFourteenthShort  = IntervalName("A14")
	IntervalNameAugmentedFifteenthShort   = IntervalName("A15")
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

func IntervalPerfectUnison() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectUnison,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectUnison,
			shortName: IntervalNamePerfectUnisonShort,
		},
		halfTones: IntervalHalfTones0,
	}
}

func IntervalMinorSecond() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameMinorSecond,
			shortName: IntervalNameMinorSecondShort,
		},
		halfTones: IntervalHalfTones1,
	}
}

func IntervalMajorSecond() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameMajorSecond,
			shortName: IntervalNameMajorSecondShort,
		},
		halfTones: IntervalHalfTones2,
	}
}

func IntervalMinorThird() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorThird,
		names: &intervalNameExtended{
			name:      IntervalNameMinorThird,
			shortName: IntervalNameMinorThirdShort,
		},
		halfTones: IntervalHalfTones3,
	}
}

func IntervalMajorThird() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorThird,
		names: &intervalNameExtended{
			name:      IntervalNameMajorThird,
			shortName: IntervalNameMajorThirdShort,
		},
		halfTones: IntervalHalfTones4,
	}
}

func IntervalPerfectFourth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectFourth,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectFourth,
			shortName: IntervalNamePerfectFourthShort,
		},
		halfTones: IntervalHalfTones5,
	}
}

func IntervalTritone() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceTritone,
		names: &intervalNameExtended{
			name:      IntervalNameTritone,
			shortName: IntervalNameTritoneShort,
		},
		halfTones: IntervalHalfTones6,
	}
}

func IntervalPerfectFifth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectFifth,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectFifth,
			shortName: IntervalNamePerfectFifthShort,
		},
		halfTones: IntervalHalfTones7,
	}
}

func IntervalMinorSixth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameMinorSixth,
			shortName: IntervalNameMinorSixthShort,
		},
		halfTones: IntervalHalfTones8,
	}
}

func IntervalMajorSixth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameMajorSixth,
			shortName: IntervalNameMajorSixthShort,
		},
		halfTones: IntervalHalfTones9,
	}
}

func IntervalMinorSeventh() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameMinorSeventh,
			shortName: IntervalNameMinorSeventhShort,
		},
		halfTones: IntervalHalfTones10,
	}
}

func IntervalMajorSeventh() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameMajorSeventh,
			shortName: IntervalNameMajorSeventhShort,
		},
		halfTones: IntervalHalfTones11,
	}
}

func IntervalPerfectOctave() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectOctave,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectOctave,
			shortName: IntervalNamePerfectOctaveShort,
		},
		halfTones: IntervalHalfTones12,
	}
}

func IntervalMinorNinth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorNinth,
			shortName: IntervalNameMinorNinthShort,
		},
		halfTones: IntervalHalfTones13,
	}
}

func IntervalMajorNinth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorNinth,
			shortName: IntervalNameMajorNinthShort,
		},
		halfTones: IntervalHalfTones14,
	}
}

func IntervalMinorTenth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorTenth,
			shortName: IntervalNameMinorTenthShort,
		},
		halfTones: IntervalHalfTones15,
	}
}

func IntervalMajorTenth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorTenth,
			shortName: IntervalNameMajorTenthShort,
		},
		halfTones: IntervalHalfTones16,
	}
}

func IntervalPerfectEleventh() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectEleventh,
			shortName: IntervalNamePerfectEleventhShort,
		},
		halfTones: IntervalHalfTones17,
	}
}

func IntervalOctaveWithTritone() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameOctaveWithTritone,
			shortName: IntervalNameOctaveWithTritoneShort,
		},
		halfTones: IntervalHalfTones17,
	}
}

func IntervalPerfectTwelfth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectTwelfth,
			shortName: IntervalNamePerfectTwelfthShort,
		},
		halfTones: IntervalHalfTones19,
	}
}

func IntervalMinorThirteenth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorThirteenth,
			shortName: IntervalNameMinorThirteenthShort,
		},
		halfTones: IntervalHalfTones20,
	}
}

func IntervalMajorThirteenth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorThirteenth,
			shortName: IntervalNameMajorThirteenthShort,
		},
		halfTones: IntervalHalfTones21,
	}
}

func IntervalMinorFourteenth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMinorFourteenth,
			shortName: IntervalNameMinorFourteenthShort,
		},
		halfTones: IntervalHalfTones22,
	}
}

func IntervalMajorFourteenth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNameMajorFourteenth,
			shortName: IntervalNameMajorFourteenthShort,
		},
		halfTones: IntervalHalfTones23,
	}
}

func IntervalPerfectFifteenth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: 0,
		names: &intervalNameExtended{
			name:      IntervalNamePerfectFifteenth,
			shortName: IntervalNamePerfectFifteenthShort,
		},
		halfTones: IntervalHalfTones24,
	}
}

// Compound chromatic (diatonic) intervals

func IntervalDiminishedSecond() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectUnison,
		names: &intervalNameExtended{
			name:      IntervalNameDiminishedSecond,
			shortName: IntervalNameDiminishedSecondShort,
		},
		halfTones: IntervalHalfTones0,
	}
}

func IntervalAugmentedUnison() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameAugmentedUnison,
			shortName: IntervalNameAugmentedUnisonShort,
		},
		halfTones: IntervalHalfTones1,
	}
}

func IntervalDiminishedThird() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorSecond,
		names: &intervalNameExtended{
			name:      IntervalNameDiminishedThird,
			shortName: IntervalNameDiminishedThirdShort,
		},
		halfTones: IntervalHalfTones2,
	}
}

func IntervalAugmentedSecond() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorThird,
		names: &intervalNameExtended{
			name:      IntervalNameAugmentedSecond,
			shortName: IntervalNameAugmentedSecondShort,
		},
		halfTones: IntervalHalfTones3,
	}
}

func IntervalDiminishedFourth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorThird,
		names: &intervalNameExtended{
			name:      IntervalNameDiminishedFourth,
			shortName: IntervalNameDiminishedFourthShort,
		},
		halfTones: IntervalHalfTones4,
	}
}

func IntervalAugmentedThird() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectFourth,
		names: &intervalNameExtended{
			name:      IntervalNameAugmentedThird,
			shortName: IntervalNameAugmentedThirdShort,
		},
		halfTones: IntervalHalfTones5,
	}
}

func IntervalDiminishedFifth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceTritone,
		names: &intervalNameExtended{
			name:      IntervalNameDiminishedFifth,
			shortName: IntervalNameDiminishedFifthShort,
		},
		halfTones: IntervalHalfTones6,
	}
}

func IntervalAugmentedFourth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceTritone,
		names: &intervalNameExtended{
			name:      IntervalNameAugmentedFourth,
			shortName: IntervalNameAugmentedFourthShort,
		},
		halfTones: IntervalHalfTones6,
	}
}

func IntervalDiminishedSixth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectFifth,
		names: &intervalNameExtended{
			name:      IntervalNameDiminishedSixth,
			shortName: IntervalNameDiminishedSixthShort,
		},
		halfTones: IntervalHalfTones7,
	}
}

func IntervalAugmentedFifth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameAugmentedFifth,
			shortName: IntervalNameAugmentedFifthShort,
		},
		halfTones: IntervalHalfTones8,
	}
}

func IntervalDiminishedSeventh() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorSixth,
		names: &intervalNameExtended{
			name:      IntervalNameDiminishedSeventh,
			shortName: IntervalNameDiminishedSeventhShort,
		},
		halfTones: IntervalHalfTones9,
	}
}

func IntervalAugmentedSixth() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMinorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameAugmentedSixth,
			shortName: IntervalNameAugmentedSixthShort,
		},
		halfTones: IntervalHalfTones10,
	}
}

func IntervalDiminishedOctave() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonanceMajorSeventh,
		names: &intervalNameExtended{
			name:      IntervalNameDiminishedOctave,
			shortName: IntervalNameDiminishedOctaveShort,
		},
		halfTones: IntervalHalfTones11,
	}
}

func IntervalAugmentedSeventh() *IntervalChromatic {
	return &IntervalChromatic{
		Sonance: IntervalSonancePerfectOctave,
		names: &intervalNameExtended{
			name:      IntervalNameAugmentedSeventh,
			shortName: IntervalNameAugmentedSeventhShort,
		},
		halfTones: IntervalHalfTones12,
	}
}
