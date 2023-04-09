package muse

type ModeTemplatesStore map[ModeName]ModeTemplate

func (mts ModeTemplatesStore) AddTemplate(modeName ModeName, moteTemplate ModeTemplate) {
	mts[modeName] = moteTemplate
}

func InitModeTemplatesStore() ModeTemplatesStore {
	mts := make(ModeTemplatesStore)
	var err error

	// Tonal modes
	mts[ModeNameNaturalMinor], err = GetTemplateByModeName(ModeNameNaturalMinor)
	errCheck(err)
	mts[ModeNameMelodicMinor], err = GetTemplateByModeName(ModeNameMelodicMinor)
	errCheck(err)
	mts[ModeNameHarmonicMinor], err = GetTemplateByModeName(ModeNameHarmonicMinor)
	errCheck(err)
	mts[ModeNameNaturalMajor], err = GetTemplateByModeName(ModeNameNaturalMajor)
	errCheck(err)
	mts[ModeNameMelodicMajor], err = GetTemplateByModeName(ModeNameMelodicMajor)
	errCheck(err)
	mts[ModeNameHarmonicMajor], err = GetTemplateByModeName(ModeNameHarmonicMajor)
	errCheck(err)

	// Modes of the Major scale
	mts[ModeNameIonian], err = GetTemplateByModeName(ModeNameIonian)
	errCheck(err)
	mts[ModeNameDorian], err = GetTemplateByModeName(ModeNameDorian)
	errCheck(err)
	mts[ModeNameAeolian], err = GetTemplateByModeName(ModeNameAeolian)
	errCheck(err)
	mts[ModeNameLydian], err = GetTemplateByModeName(ModeNameLydian)
	errCheck(err)
	mts[ModeNameMixoLydian], err = GetTemplateByModeName(ModeNameMixoLydian)
	errCheck(err)
	mts[ModeNamePhrygian], err = GetTemplateByModeName(ModeNamePhrygian)
	errCheck(err)
	mts[ModeNameLocrian], err = GetTemplateByModeName(ModeNameLocrian)
	errCheck(err)

	// Modes Of The Melodic Minor scale
	mts[ModeNameIonianFlat3], err = GetTemplateByModeName(ModeNameIonianFlat3)
	errCheck(err)
	mts[ModeNamePhrygoDorian], err = GetTemplateByModeName(ModeNamePhrygoDorian)
	errCheck(err)
	mts[ModeNameLydianAugmented], err = GetTemplateByModeName(ModeNameLydianAugmented)
	errCheck(err)
	mts[ModeNameLydianDominant], err = GetTemplateByModeName(ModeNameLydianDominant)
	errCheck(err)
	mts[ModeNameIonianAeolian], err = GetTemplateByModeName(ModeNameIonianAeolian)
	errCheck(err)
	mts[ModeNameAeolianLydian], err = GetTemplateByModeName(ModeNameAeolianLydian)
	errCheck(err)
	mts[ModeNameSuperLocrian], err = GetTemplateByModeName(ModeNameSuperLocrian)
	errCheck(err)

	// Modes of the Harmonic Minor scale
	mts[ModeNameAeolianRais7], err = GetTemplateByModeName(ModeNameAeolianRais7)
	errCheck(err)
	mts[ModeNameLocrianRais6], err = GetTemplateByModeName(ModeNameLocrianRais6)
	errCheck(err)
	mts[ModeNameIonianRais5], err = GetTemplateByModeName(ModeNameIonianRais5)
	errCheck(err)
	mts[ModeNameUkrainianDorian], err = GetTemplateByModeName(ModeNameUkrainianDorian)
	errCheck(err)
	mts[ModeNamePhrygianDominant], err = GetTemplateByModeName(ModeNamePhrygianDominant)
	errCheck(err)
	mts[ModeNameLydianRais9], err = GetTemplateByModeName(ModeNameLydianRais9)
	errCheck(err)
	mts[ModeNameUltraLocrian], err = GetTemplateByModeName(ModeNameUltraLocrian)
	errCheck(err)

	// Modes Of The Harmonic Major scale
	mts[ModeNameIonianFlat6], err = GetTemplateByModeName(ModeNameIonianFlat6)
	errCheck(err)
	mts[ModeNameDorianDiminished], err = GetTemplateByModeName(ModeNameDorianDiminished)
	errCheck(err)
	mts[ModeNamePhrygianDiminished], err = GetTemplateByModeName(ModeNamePhrygianDiminished)
	errCheck(err)
	mts[ModeNameLydianDiminished], err = GetTemplateByModeName(ModeNameLydianDiminished)
	errCheck(err)
	mts[ModeNameMixolydianFlat2], err = GetTemplateByModeName(ModeNameMixolydianFlat2)
	errCheck(err)
	mts[ModeNameLydianAugmented2], err = GetTemplateByModeName(ModeNameLydianAugmented2)
	errCheck(err)
	mts[ModeNameLocrianDoubleFlat7], err = GetTemplateByModeName(ModeNameLocrianDoubleFlat7)
	errCheck(err)

	// Double Harmonic Major Modes
	mts[ModeNameHungarianMajor], err = GetTemplateByModeName(ModeNameHungarianMajor)
	errCheck(err)
	mts[ModeNameLydianRais2Rais6], err = GetTemplateByModeName(ModeNameLydianRais2Rais6)
	errCheck(err)
	mts[ModeNameUltraphrygian], err = GetTemplateByModeName(ModeNameUltraphrygian)
	errCheck(err)
	mts[ModeNameHungarianMinor], err = GetTemplateByModeName(ModeNameHungarianMinor)
	errCheck(err)
	mts[ModeNameOriental], err = GetTemplateByModeName(ModeNameOriental)
	errCheck(err)
	mts[ModeNameIonianAugmented2], err = GetTemplateByModeName(ModeNameIonianAugmented2)
	errCheck(err)
	mts[ModeNameLocrianDoubleFlat3DoubleFlat7], err = GetTemplateByModeName(ModeNameLocrianDoubleFlat3DoubleFlat7)
	errCheck(err)

	// Pentatonic modes

	// Main pentatonic modes
	mts[ModeNamePentatonicMajor], err = GetTemplateByModeName(ModeNamePentatonicMajor)
	errCheck(err)
	mts[ModeNamePentatonicSustained], err = GetTemplateByModeName(ModeNamePentatonicSustained)
	errCheck(err)
	mts[ModeNamePentatonicBluesMinor], err = GetTemplateByModeName(ModeNamePentatonicBluesMinor)
	errCheck(err)
	mts[ModeNamePentatonicBluesMajor], err = GetTemplateByModeName(ModeNamePentatonicBluesMajor)
	errCheck(err)
	mts[ModeNamePentatonicMinor], err = GetTemplateByModeName(ModeNamePentatonicMinor)
	errCheck(err)

	return mts
}

// FindModeTemplatesByPattern searches mode templates by intersection with the given mode template.
func (mts ModeTemplatesStore) FindModeTemplatesByPattern(modeTemplate ModeTemplate) ModeTemplatesStore {
	mt := modeTemplate
	result := make(ModeTemplatesStore)
	for modeName, storedModeTemplate := range mts {
		for i := 0; i <= int(storedModeTemplate.Length()); i++ {
			if i > int(storedModeTemplate.Length())-len(mt) {
				break
			}
			interval := storedModeTemplate[i]
			if mt[0] == interval {
				mt = mt[1:]
			} else {
				if len(mt) < len(modeTemplate) {
					i--
				}
				mt = modeTemplate
			}
			if len(mt) == 0 {
				result[modeName] = storedModeTemplate

				break
			}
		}
		mt = modeTemplate
	}

	return result
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
