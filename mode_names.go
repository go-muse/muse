package muse

// Modename is just a name for a mode.
type ModeName string

// Tonal modes

const (
	ModeNameNaturalMinor  = ModeName("NaturalMinor")
	ModeNameMelodicMinor  = ModeName("MelodicMinor")
	ModeNameHarmonicMinor = ModeName("HarmonicMinor")
	ModeNameNaturalMajor  = ModeName("NaturalMajor")
	ModeNameMelodicMajor  = ModeName("MelodicMajor")
	ModeNameHarmonicMajor = ModeName("HarmonicMajor")
)

// Modes of the Major scale

const (
	ModeNameIonian     = ModeName("Ionian") // Natural Major
	ModeNameDorian     = ModeName("Dorian")
	ModeNameAeolian    = ModeName("Aeolian") // Natural Minor
	ModeNameLydian     = ModeName("Lydian")
	ModeNameMixoLydian = ModeName("MixoLydian")
	ModeNamePhrygian   = ModeName("Phrygian")
	ModeNameLocrian    = ModeName("Locrian")
)

// Modes Of The Melodic Minor scale

const (
	ModeNameIonianFlat3     = ModeName("IonianFlat3")     // Melodic Minor, Jazz minor
	ModeNamePhrygoDorian    = ModeName("PhrygoDorian")    // Phrygian with #6
	ModeNameLydianAugmented = ModeName("LydianAugmented") // The Lydian augmented #4 and #5
	ModeNameLydianDominant  = ModeName("LydianDominant")  // Overtone, acoustic scale; Lydian-MixoLydian; Podgalyansky;
	ModeNameIonianAeolian   = ModeName("IonianAeolian")   // Melodic Major
	ModeNameAeolianLydian   = ModeName("AeolianLydian")   // Locrian #2, Aeolian b5 or Semi Locrian
	ModeNameSuperLocrian    = ModeName("SuperLocrian")    //
)

// Modes of the Harmonic Minor scale

const (
	ModeNameAeolianRais7     = ModeName("AeolianRais7 ")    // Harmonic Minor
	ModeNameLocrianRais6     = ModeName("LocrianRais6")     // Locrian with natural 6th
	ModeNameIonianRais5      = ModeName("IonianRais5")      // Augmented Major
	ModeNameUkrainianDorian  = ModeName("UkrainianDorian")  // Romanian Minor, Altered Dorian
	ModeNamePhrygianDominant = ModeName("PhrygianDominant") // Phrygian Major
	ModeNameLydianRais9      = ModeName("LydianRais9")      // Aeolian Harmonic
	ModeNameUltraLocrian     = ModeName("UltraLocrian")     // Super Locrian Diminished, Super Locrian bb7
)

// Modes Of The Harmonic Major scale

const (
	ModeNameIonianFlat6        = ModeName("IonianFlat6 ")       // Harmonic Major, Ethiopian mode
	ModeNameDorianDiminished   = ModeName("DorianDiminished")   // The Dorian b5 Mode, Locrian #2 #6
	ModeNamePhrygianDiminished = ModeName("PhrygianDiminished") // The Phrygian b4 Mode, Indian mode
	ModeNameLydianDiminished   = ModeName("LydianDiminished")   // Lydian flat 3, Lydian Minor, Melodic Minor #4
	ModeNameMixolydianFlat2    = ModeName("MixolydianFlat2")    // Harmonic Minor Inverse
	ModeNameLydianAugmented2   = ModeName("LydianAugmented2")   // The Lydian augmented #2
	ModeNameLocrianDoubleFlat7 = ModeName("LocrianDoubleFlat7") // Diminished Blues b9
)

// Double Harmonic Major Modes

const (
	ModeNameHungarianMajor                = ModeName("HungarianMajor")                // Double harmonic major, Arabian, Gypsy major
	ModeNameLydianRais2Rais6              = ModeName("LydianRais2Rais6")              // Lydian with raised 2nd and 6th
	ModeNameUltraphrygian                 = ModeName("Ultraphrygian")                 // Phrygian b4 bb7
	ModeNameHungarianMinor                = ModeName("HungarianMinor")                // Double harmonic minor, Gypsy minor
	ModeNameOriental                      = ModeName("Oriental")                      // Asian, Mixolydian with flattened 2nd and flattened 5th
	ModeNameIonianAugmented2              = ModeName("IonianAugmented2")              //
	ModeNameLocrianDoubleFlat3DoubleFlat7 = ModeName("LocrianDoubleFlat3DoubleFlat7") // Locrian bb3 bb7
)

// Pentatonic modes

// Main pentatonics.
const (
	ModeNamePentatonicMajor      = ModeName("PentatonicMajor")
	ModeNamePentatonicSustained  = ModeName("PentatonicSustained")  // African, Malian
	ModeNamePentatonicBluesMinor = ModeName("PentatonicBluesMinor") // The Man Gong
	ModeNamePentatonicBluesMajor = ModeName("PentatonicBluesMajor") // The Ritusen, Scottish pentatonic
	ModeNamePentatonicMinor      = ModeName("PentatonicMinor")
)
