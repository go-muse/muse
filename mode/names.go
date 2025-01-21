package mode

// Tonal modes

const (
	NameNaturalMinor  = Name("NaturalMinor")
	NameMelodicMinor  = Name("MelodicMinor")
	NameHarmonicMinor = Name("HarmonicMinor")
	NameNaturalMajor  = Name("NaturalMajor")
	NameMelodicMajor  = Name("MelodicMajor")
	NameHarmonicMajor = Name("HarmonicMajor")
)

// Modes of the Major scale

const (
	NameIonian     = Name("Ionian") // Natural Major
	NameDorian     = Name("Dorian")
	NameAeolian    = Name("Aeolian") // Natural Minor
	NameLydian     = Name("Lydian")
	NameMixoLydian = Name("MixoLydian")
	NamePhrygian   = Name("Phrygian")
	NameLocrian    = Name("Locrian")
)

// Modes Of The Melodic Minor scale

const (
	NameIonianFlat3     = Name("IonianFlat3")     // Melodic Minor, Jazz minor
	NamePhrygoDorian    = Name("PhrygoDorian")    // Phrygian with #6
	NameLydianAugmented = Name("LydianAugmented") // The Lydian augmented #4 and #5
	NameLydianDominant  = Name("LydianDominant")  // Overtone, acoustic scale; Lydian-MixoLydian; Podgalyansky;
	NameIonianAeolian   = Name("IonianAeolian")   // Melodic Major
	NameAeolianLydian   = Name("AeolianLydian")   // Locrian #2, Aeolian b5 or Semi Locrian
	NameSuperLocrian    = Name("SuperLocrian")    //
)

// Modes of the Harmonic Minor scale

const (
	NameAeolianRais7     = Name("AeolianRais7")     // Harmonic Minor
	NameLocrianRais6     = Name("LocrianRais6")     // Locrian with natural 6th
	NameIonianRais5      = Name("IonianRais5")      // Augmented Major
	NameUkrainianDorian  = Name("UkrainianDorian")  // Romanian Minor, Altered Dorian
	NamePhrygianDominant = Name("PhrygianDominant") // Phrygian Major
	NameLydianRais9      = Name("LydianRais9")      // Aeolian Harmonic
	NameUltraLocrian     = Name("UltraLocrian")     // Super Locrian Diminished, Super Locrian bb7
)

// Modes Of The Harmonic Major scale

const (
	NameIonianFlat6        = Name("IonianFlat6")        // Harmonic Major, Ethiopian mode
	NameDorianDiminished   = Name("DorianDiminished")   // The Dorian b5 Mode, Locrian #2 #6
	NamePhrygianDiminished = Name("PhrygianDiminished") // The Phrygian b4 Mode, Indian mode
	NameLydianDiminished   = Name("LydianDiminished")   // Lydian flat 3, Lydian Minor, Melodic Minor #4
	NameMixolydianFlat2    = Name("MixolydianFlat2")    // Harmonic Minor Inverse
	NameLydianAugmented2   = Name("LydianAugmented2")   // The Lydian augmented #2
	NameLocrianDoubleFlat7 = Name("LocrianDoubleFlat7") // Diminished Blues b9
)

// Double Harmonic Major Modes

const (
	NameHungarianMajor                = Name("HungarianMajor")                // Double harmonic major, Arabian, Gypsy major
	NameLydianRais2Rais6              = Name("LydianRais2Rais6")              // Lydian with raised 2nd and 6th
	NameUltraPhrygian                 = Name("UltraPhrygian")                 // Phrygian b4 bb7
	NameHungarianMinor                = Name("HungarianMinor")                // Double harmonic minor, Gypsy minor
	NameOriental                      = Name("Oriental")                      // Asian, Mixolydian with flattened 2nd and flattened 5th
	NameIonianAugmented2              = Name("IonianAugmented2")              //
	NameLocrianDoubleFlat3DoubleFlat7 = Name("LocrianDoubleFlat3DoubleFlat7") // Locrian bb3 bb7
)

// Pentatonic modes

// Main pentatonics.
const (
	NamePentatonicMajor      = Name("PentatonicMajor")
	NamePentatonicSustained  = Name("PentatonicSustained")  // African, Malian
	NamePentatonicBluesMinor = Name("PentatonicBluesMinor") // The Man Gong
	NamePentatonicBluesMajor = Name("PentatonicBluesMajor") // The Ritusen, Scottish pentatonic
	NamePentatonicMinor      = Name("PentatonicMinor")
)

// Japanese pentatonics.
const (
	NamePentatonicHirajoshi        = Name("PentatonicHirajoshi")
	NamePentatonicIwato            = Name("PentatonicIwato")
	NamePentatonicHonKumoiShiouzhi = Name("PentatonicHonKumoiShiouzhi") // Raga Bhinna Shadja
	NamePentatonicHonKumoiJoshi    = Name("PentatonicHonKumoiJoshi")
	NamePentatonicLydianPentatonic = Name("PentatonicLydianPentatonic") // Augmented, Raga Amritavarshini (Chinese)
)
