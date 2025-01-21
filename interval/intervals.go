package interval

import (
	"github.com/go-muse/muse/halftone"
)

const (
	// Chromatic intervals.

	NamePerfectUnison = Name("PerfectUnison")
	NameMinorSecond   = Name("MinorSecond")
	NameMajorSecond   = Name("MajorSecond")
	NameMinorThird    = Name("MinorThird")
	NameMajorThird    = Name("MajorThird")
	NamePerfectFourth = Name("PerfectFourth")
	NameTritone       = Name("Tritone")
	NamePerfectFifth  = Name("PerfectFifth")
	NameMinorSixth    = Name("MinorSixth")
	NameMajorSixth    = Name("MajorSixth")
	NameMinorSeventh  = Name("MinorSeventh")
	NameMajorSeventh  = Name("MajorSeventh")
	NamePerfectOctave = Name("PerfectOctave")

	NameMinorNinth        = Name("MinorNinth")
	NameMajorNinth        = Name("MajorNinth")
	NameMinorTenth        = Name("MinorTenth")
	NameMajorTenth        = Name("MajorTenth")
	NamePerfectEleventh   = Name("PerfectEleventh")
	NameOctaveWithTritone = Name("OctaveWithTritone") // no real name for tritone after octave
	NamePerfectTwelfth    = Name("PerfectTwelfth")    // Tritave
	NameMinorThirteenth   = Name("MinorThirteenth")
	NameMajorThirteenth   = Name("MajorThirteenth")
	NameMinorFourteenth   = Name("MinorFourteenth")
	NameMajorFourteenth   = Name("MajorFourteenth")
	NamePerfectFifteenth  = Name("PerfectFifteenth ") // Double octave

	NamePerfectUnisonShort = Name("P1")
	NameMinorSecondShort   = Name("m2")
	NameMajorSecondShort   = Name("M2")
	NameMinorThirdShort    = Name("m3")
	NameMajorThirdShort    = Name("M3")
	NamePerfectFourthShort = Name("P4")
	NameTritoneShort       = Name("TT")
	NamePerfectFifthShort  = Name("P5")
	NameMinorSixthShort    = Name("m6")
	NameMajorSixthShort    = Name("M6")
	NameMinorSeventhShort  = Name("m7")
	NameMajorSeventhShort  = Name("M7")
	NamePerfectOctaveShort = Name("P8")

	NameMinorNinthShort        = Name("m9")
	NameMajorNinthShort        = Name("M9")
	NameMinorTenthShort        = Name("m10")
	NameMajorTenthShort        = Name("M10")
	NamePerfectEleventhShort   = Name("P11")
	NameOctaveWithTritoneShort = Name("OT")
	NamePerfectTwelfthShort    = Name("P12")
	NameMinorThirteenthShort   = Name("m13")
	NameMajorThirteenthShort   = Name("M13")
	NameMinorFourteenthShort   = Name("m14")
	NameMajorFourteenthShort   = Name("M14")
	NamePerfectFifteenthShort  = Name("P15")

	// Diatonic intervals.

	NameDiminishedSecond  = Name("DiminishedSecond")
	NameAugmentedUnison   = Name("AugmentedUnison")
	NameDiminishedThird   = Name("DiminishedThird")
	NameAugmentedSecond   = Name("AugmentedSecond")
	NameDiminishedFourth  = Name("DiminishedFourth")
	NameAugmentedThird    = Name("AugmentedThird")
	NameDiminishedFifth   = Name("DiminishedFifth")
	NameAugmentedFourth   = Name("AugmentedFourth")
	NameDiminishedSixth   = Name("DiminishedSixth")
	NameAugmentedFifth    = Name("AugmentedFifth")
	NameDiminishedSeventh = Name("DiminishedSeventh")
	NameAugmentedSixth    = Name("AugmentedSixth")
	NameDiminishedOctave  = Name("DiminishedOctave")
	NameAugmentedSeventh  = Name("AugmentedSeventh")

	NameDiminishedNinth      = Name("DiminishedNinth")
	NameAugmentedOctave      = Name("AugmentedOctave")
	NameDiminishedTenth      = Name("DiminishedTenth")
	NameAugmentedNinth       = Name("AugmentedNinth")
	NameDiminishedEleventh   = Name("DiminishedEleventh")
	NameAugmentedTenth       = Name("AugmentedTenth")
	NameDiminishedTwelfth    = Name("DiminishedTwelfth")
	NameAugmentedEleventh    = Name("AugmentedEleventh")
	NameDiminishedThirteenth = Name("DiminishedThirteenth")
	NameAugmentedTwelfth     = Name("AugmentedTwelfth")
	NameDiminishedFourteenth = Name("DiminishedFourteenth")
	NameAugmentedThirteenth  = Name("AugmentedThirteenth")
	NameDiminishedFifteenth  = Name("DiminishedFifteenth")
	NameAugmentedFourteenth  = Name("AugmentedFourteenth") // Double octave
	NameAugmentedFifteenth   = Name("AugmentedFifteenth")

	NameDiminishedSecondShort  = Name("d2")
	NameAugmentedUnisonShort   = Name("A1")
	NameDiminishedThirdShort   = Name("d3")
	NameAugmentedSecondShort   = Name("A2")
	NameDiminishedFourthShort  = Name("d4")
	NameAugmentedThirdShort    = Name("A3")
	NameDiminishedFifthShort   = Name("d5")
	NameAugmentedFourthShort   = Name("A4")
	NameDiminishedSixthShort   = Name("d6")
	NameAugmentedFifthShort    = Name("A5")
	NameDiminishedSeventhShort = Name("d7")
	NameAugmentedSixthShort    = Name("A6")
	NameDiminishedOctaveShort  = Name("d8")
	NameAugmentedSeventhShort  = Name("A7")

	NameDiminishedNinthShort      = Name("d9")
	NameAugmentedOctaveShort      = Name("A8")
	NameDiminishedTenthShort      = Name("d10")
	NameAugmentedNinthShort       = Name("A9")
	NameDiminishedEleventhShort   = Name("d11")
	NameAugmentedTenthShort       = Name("A10")
	NameDiminishedTwelfthShort    = Name("d12")
	NameAugmentedEleventhShort    = Name("A11")
	NameDiminishedThirteenthShort = Name("d13")
	NameAugmentedTwelfthShort     = Name("A12")
	NameDiminishedFourteenthShort = Name("d14")
	NameAugmentedThirteenthShort  = Name("A13")
	NameDiminishedFifteenthShort  = Name("d15")
	NameAugmentedFourteenthShort  = Name("A14")
	NameAugmentedFifteenthShort   = Name("A15")
)

const (
	HalfTones0  = halftone.HalfTones(0)
	HalfTones1  = halftone.HalfTones(1)
	HalfTones2  = halftone.HalfTones(2)
	HalfTones3  = halftone.HalfTones(3)
	HalfTones4  = halftone.HalfTones(4)
	HalfTones5  = halftone.HalfTones(5)
	HalfTones6  = halftone.HalfTones(6)
	HalfTones7  = halftone.HalfTones(7)
	HalfTones8  = halftone.HalfTones(8)
	HalfTones9  = halftone.HalfTones(9)
	HalfTones10 = halftone.HalfTones(10)
	HalfTones11 = halftone.HalfTones(11)
	HalfTones12 = halftone.HalfTones(12)
	HalfTones13 = halftone.HalfTones(13)
	HalfTones14 = halftone.HalfTones(14)
	HalfTones15 = halftone.HalfTones(15)
	HalfTones16 = halftone.HalfTones(16)
	HalfTones17 = halftone.HalfTones(17)
	HalfTones18 = halftone.HalfTones(18)
	HalfTones19 = halftone.HalfTones(19)
	HalfTones20 = halftone.HalfTones(20)
	HalfTones21 = halftone.HalfTones(21)
	HalfTones22 = halftone.HalfTones(22)
	HalfTones23 = halftone.HalfTones(23)
	HalfTones24 = halftone.HalfTones(24)
)

// Chromatic intervals.

func PerfectUnison() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectUnison,
		names: &nameExtended{
			name:      NamePerfectUnison,
			shortName: NamePerfectUnisonShort,
		},
		halfTones: HalfTones0,
	}
}

func MinorSecond() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorSecond,
		names: &nameExtended{
			name:      NameMinorSecond,
			shortName: NameMinorSecondShort,
		},
		halfTones: HalfTones1,
	}
}

func MajorSecond() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorSecond,
		names: &nameExtended{
			name:      NameMajorSecond,
			shortName: NameMajorSecondShort,
		},
		halfTones: HalfTones2,
	}
}

func MinorThird() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorThird,
		names: &nameExtended{
			name:      NameMinorThird,
			shortName: NameMinorThirdShort,
		},
		halfTones: HalfTones3,
	}
}

func MajorThird() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorThird,
		names: &nameExtended{
			name:      NameMajorThird,
			shortName: NameMajorThirdShort,
		},
		halfTones: HalfTones4,
	}
}

func PerfectFourth() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectFourth,
		names: &nameExtended{
			name:      NamePerfectFourth,
			shortName: NamePerfectFourthShort,
		},
		halfTones: HalfTones5,
	}
}

func Tritone() *Chromatic {
	return &Chromatic{
		Sonance: SonanceTritone,
		names: &nameExtended{
			name:      NameTritone,
			shortName: NameTritoneShort,
		},
		halfTones: HalfTones6,
	}
}

func PerfectFifth() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectFifth,
		names: &nameExtended{
			name:      NamePerfectFifth,
			shortName: NamePerfectFifthShort,
		},
		halfTones: HalfTones7,
	}
}

func MinorSixth() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorSixth,
		names: &nameExtended{
			name:      NameMinorSixth,
			shortName: NameMinorSixthShort,
		},
		halfTones: HalfTones8,
	}
}

func MajorSixth() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorSixth,
		names: &nameExtended{
			name:      NameMajorSixth,
			shortName: NameMajorSixthShort,
		},
		halfTones: HalfTones9,
	}
}

func MinorSeventh() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorSeventh,
		names: &nameExtended{
			name:      NameMinorSeventh,
			shortName: NameMinorSeventhShort,
		},
		halfTones: HalfTones10,
	}
}

func MajorSeventh() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorSeventh,
		names: &nameExtended{
			name:      NameMajorSeventh,
			shortName: NameMajorSeventhShort,
		},
		halfTones: HalfTones11,
	}
}

func PerfectOctave() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectOctave,
		names: &nameExtended{
			name:      NamePerfectOctave,
			shortName: NamePerfectOctaveShort,
		},
		halfTones: HalfTones12,
	}
}

func MinorNinth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMinorNinth,
			shortName: NameMinorNinthShort,
		},
		halfTones: HalfTones13,
	}
}

func MajorNinth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMajorNinth,
			shortName: NameMajorNinthShort,
		},
		halfTones: HalfTones14,
	}
}

func MinorTenth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMinorTenth,
			shortName: NameMinorTenthShort,
		},
		halfTones: HalfTones15,
	}
}

func MajorTenth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMajorTenth,
			shortName: NameMajorTenthShort,
		},
		halfTones: HalfTones16,
	}
}

func PerfectEleventh() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NamePerfectEleventh,
			shortName: NamePerfectEleventhShort,
		},
		halfTones: HalfTones17,
	}
}

func OctaveWithTritone() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameOctaveWithTritone,
			shortName: NameOctaveWithTritoneShort,
		},
		halfTones: HalfTones17,
	}
}

func PerfectTwelfth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NamePerfectTwelfth,
			shortName: NamePerfectTwelfthShort,
		},
		halfTones: HalfTones19,
	}
}

func MinorThirteenth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMinorThirteenth,
			shortName: NameMinorThirteenthShort,
		},
		halfTones: HalfTones20,
	}
}

func MajorThirteenth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMajorThirteenth,
			shortName: NameMajorThirteenthShort,
		},
		halfTones: HalfTones21,
	}
}

func MinorFourteenth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMinorFourteenth,
			shortName: NameMinorFourteenthShort,
		},
		halfTones: HalfTones22,
	}
}

func MajorFourteenth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NameMajorFourteenth,
			shortName: NameMajorFourteenthShort,
		},
		halfTones: HalfTones23,
	}
}

func PerfectFifteenth() *Chromatic {
	return &Chromatic{
		Sonance: 0,
		names: &nameExtended{
			name:      NamePerfectFifteenth,
			shortName: NamePerfectFifteenthShort,
		},
		halfTones: HalfTones24,
	}
}

// Compound chromatic (diatonic) intervals

func DiminishedSecond() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectUnison,
		names: &nameExtended{
			name:      NameDiminishedSecond,
			shortName: NameDiminishedSecondShort,
		},
		halfTones: HalfTones0,
	}
}

func AugmentedUnison() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorSecond,
		names: &nameExtended{
			name:      NameAugmentedUnison,
			shortName: NameAugmentedUnisonShort,
		},
		halfTones: HalfTones1,
	}
}

func DiminishedThird() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorSecond,
		names: &nameExtended{
			name:      NameDiminishedThird,
			shortName: NameDiminishedThirdShort,
		},
		halfTones: HalfTones2,
	}
}

func AugmentedSecond() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorThird,
		names: &nameExtended{
			name:      NameAugmentedSecond,
			shortName: NameAugmentedSecondShort,
		},
		halfTones: HalfTones3,
	}
}

func DiminishedFourth() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorThird,
		names: &nameExtended{
			name:      NameDiminishedFourth,
			shortName: NameDiminishedFourthShort,
		},
		halfTones: HalfTones4,
	}
}

func AugmentedThird() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectFourth,
		names: &nameExtended{
			name:      NameAugmentedThird,
			shortName: NameAugmentedThirdShort,
		},
		halfTones: HalfTones5,
	}
}

func DiminishedFifth() *Chromatic {
	return &Chromatic{
		Sonance: SonanceTritone,
		names: &nameExtended{
			name:      NameDiminishedFifth,
			shortName: NameDiminishedFifthShort,
		},
		halfTones: HalfTones6,
	}
}

func AugmentedFourth() *Chromatic {
	return &Chromatic{
		Sonance: SonanceTritone,
		names: &nameExtended{
			name:      NameAugmentedFourth,
			shortName: NameAugmentedFourthShort,
		},
		halfTones: HalfTones6,
	}
}

func DiminishedSixth() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectFifth,
		names: &nameExtended{
			name:      NameDiminishedSixth,
			shortName: NameDiminishedSixthShort,
		},
		halfTones: HalfTones7,
	}
}

func AugmentedFifth() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorSixth,
		names: &nameExtended{
			name:      NameAugmentedFifth,
			shortName: NameAugmentedFifthShort,
		},
		halfTones: HalfTones8,
	}
}

func DiminishedSeventh() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorSixth,
		names: &nameExtended{
			name:      NameDiminishedSeventh,
			shortName: NameDiminishedSeventhShort,
		},
		halfTones: HalfTones9,
	}
}

func AugmentedSixth() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMinorSeventh,
		names: &nameExtended{
			name:      NameAugmentedSixth,
			shortName: NameAugmentedSixthShort,
		},
		halfTones: HalfTones10,
	}
}

func DiminishedOctave() *Chromatic {
	return &Chromatic{
		Sonance: SonanceMajorSeventh,
		names: &nameExtended{
			name:      NameDiminishedOctave,
			shortName: NameDiminishedOctaveShort,
		},
		halfTones: HalfTones11,
	}
}

func AugmentedSeventh() *Chromatic {
	return &Chromatic{
		Sonance: SonancePerfectOctave,
		names: &nameExtended{
			name:      NameAugmentedSeventh,
			shortName: NameAugmentedSeventhShort,
		},
		halfTones: HalfTones12,
	}
}
