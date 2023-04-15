package muse

type ModeTemplatesStore map[ModeName]ModeTemplate

func (mts ModeTemplatesStore) AddTemplate(modeName ModeName, moteTemplate ModeTemplate) {
	mts[modeName] = moteTemplate
}

func InitModeTemplatesStore() ModeTemplatesStore {
	mts := make(ModeTemplatesStore)

	// Tonal modes
	mts[ModeNameNaturalMinor] = TemplateNaturalMinor()
	mts[ModeNameMelodicMinor] = TemplateMelodicMinor()
	mts[ModeNameHarmonicMinor] = TemplateHarmonicMinor()
	mts[ModeNameNaturalMajor] = TemplateNaturalMajor()
	mts[ModeNameMelodicMajor] = TemplateMelodicMajor()
	mts[ModeNameHarmonicMajor] = TemplateHarmonicMajor()

	// Modes of the Major scale
	mts[ModeNameIonian] = TemplateIonian()
	mts[ModeNameDorian] = TemplateDorian()
	mts[ModeNameAeolian] = TemplateAeolian()
	mts[ModeNameLydian] = TemplateLydian()
	mts[ModeNameMixoLydian] = TemplateMixoLydian()
	mts[ModeNamePhrygian] = TemplatePhrygian()
	mts[ModeNameLocrian] = TemplateLocrian()

	// Modes Of The Melodic Minor scale
	mts[ModeNameIonianFlat3] = TemplateIonianFlat3()
	mts[ModeNamePhrygoDorian] = TemplatePhrygoDorian()
	mts[ModeNameLydianAugmented] = TemplateLydianAugmented()
	mts[ModeNameLydianDominant] = TemplateLydianDominant()
	mts[ModeNameIonianAeolian] = TemplateIonianAeolian()
	mts[ModeNameAeolianLydian] = TemplateAeolianLydian()
	mts[ModeNameSuperLocrian] = TemplateSuperLocrian()

	// Modes of the Harmonic Minor scale
	mts[ModeNameAeolianRais7] = TemplateAeolianRais7()
	mts[ModeNameLocrianRais6] = TemplateLocrianRais6()
	mts[ModeNameIonianRais5] = TemplateIonianRais5()
	mts[ModeNameUkrainianDorian] = TemplateUkrainianDorian()
	mts[ModeNamePhrygianDominant] = TemplatePhrygianDominant()
	mts[ModeNameLydianRais9] = TemplateLydianRais9()
	mts[ModeNameUltraLocrian] = TemplateUltraLocrian()

	// Modes Of The Harmonic Major scale
	mts[ModeNameIonianFlat6] = TemplateIonianFlat6()
	mts[ModeNameDorianDiminished] = TemplateDorianDiminished()
	mts[ModeNamePhrygianDiminished] = TemplatePhrygianDiminished()
	mts[ModeNameLydianDiminished] = TemplateLydianDiminished()
	mts[ModeNameMixolydianFlat2] = TemplateMixolydianFlat2()
	mts[ModeNameLydianAugmented2] = TemplateLydianAugmented2()
	mts[ModeNameLocrianDoubleFlat7] = TemplateLocrianDoubleFlat7()

	// Double Harmonic Major Modes
	mts[ModeNameHungarianMajor] = TemplateHungarianMajor()
	mts[ModeNameLydianRais2Rais6] = TemplateLydianRais2Rais6()
	mts[ModeNameUltraphrygian] = TemplateUltraphrygian()
	mts[ModeNameHungarianMinor] = TemplateHungarianMinor()
	mts[ModeNameOriental] = TemplateOriental()
	mts[ModeNameIonianAugmented2] = TemplateIonianAugmented2()
	mts[ModeNameLocrianDoubleFlat3DoubleFlat7] = TemplateLocrianDoubleFlat3DoubleFlat7()

	// Pentatonic modes

	// Main pentatonic modes
	mts[ModeNamePentatonicMajor] = TemplatePentatonicMajor()
	mts[ModeNamePentatonicSustained] = TemplatePentatonicSustained()
	mts[ModeNamePentatonicBluesMinor] = TemplatePentatonicBluesMinor()
	mts[ModeNamePentatonicBluesMajor] = TemplatePentatonicBluesMajor()
	mts[ModeNamePentatonicMinor] = TemplatePentatonicMinor()

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
