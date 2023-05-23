package muse

import "sort"

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

func (mts ModeTemplatesStore) Contains(modeName ModeName) bool {
	if _, ok := mts[modeName]; ok {
		return true
	}

	return false
}

// AsSlice returns map of mode names and templates as slice of structs with ste same information.
func (mts ModeTemplatesStore) AsSlice() ModeNamesAndTemplates {
	mnat := make(ModeNamesAndTemplates, 0, len(mts))
	for moteTemplateName, modeTemplate := range mts {
		mnat = append(mnat, ModeNameAndTemplate{moteTemplateName, modeTemplate})
	}

	return mnat
}

// ModeNameAndTemplate is a struct with mode name and template.
type ModeNameAndTemplate struct {
	ModeName     ModeName
	ModeTemplate ModeTemplate
}

// ModeNamesAndTemplates is a slice of structs with mode name and template.
type ModeNamesAndTemplates []ModeNameAndTemplate

// SortByModeName sorts the slice with mode names and templates by mode name.
func (mnat ModeNamesAndTemplates) SortByModeName(desc bool) ModeNamesAndTemplates {
	sort.Slice(mnat, func(i, j int) bool {
		switch desc {
		case true:
			return mnat[i].ModeName > mnat[j].ModeName
		default:
			return mnat[i].ModeName < mnat[j].ModeName
		}
	})

	return mnat
}

// SortByTemplate sorts the slice with mode names and templates by templates.
func (mnat ModeNamesAndTemplates) SortByModeTemplate(desc bool) ModeNamesAndTemplates {
	sort.Slice(mnat[:], func(i, j int) bool {
		for x := range mnat[i].ModeTemplate {
			if mnat[i].ModeTemplate[x] == mnat[j].ModeTemplate[x] {
				continue
			}

			switch desc {
			case true:
				return mnat[i].ModeTemplate[x] > mnat[j].ModeTemplate[x]
			default:
				return mnat[i].ModeTemplate[x] < mnat[j].ModeTemplate[x]
			}
		}

		switch desc {
		case true:
			return mnat[i].ModeName > mnat[j].ModeName
		default:
			return mnat[i].ModeName < mnat[j].ModeName
		}
	})

	return mnat
}

type ModeTemplateWithPrime struct {
	*ModeNameAndTemplate
	PrimeNote *Note
}

type ModeTemplatesWithPrime []ModeTemplateWithPrime

func (mtswp ModeTemplatesWithPrime) Contains(modeName ModeName) bool {
	for _, mtwp := range mtswp {
		if mtwp.ModeName == modeName {
			return true
		}
	}

	return false
}

// FindModeTemplatesByNotes searches for modes in the storage
// that correspond to the given set of notes and returns a new storage with the found modes.
func (mts ModeTemplatesStore) FindModeTemplatesByNotes(notes []*Note) ModeTemplatesWithPrime {
	result := make(ModeTemplatesWithPrime, 0)
	var mode *Mode
	for modeName, modeTemplate := range mts {
		allNotes := GetAllPossibleNotes(1)
		for _, firstNote := range allNotes {
			mode = MustMakeNewMode(modeName, firstNote.Name())
			for _, note := range notes {
				if !mode.Contains(note) {
					goto Next
				}
			}

			result = append(result, ModeTemplateWithPrime{
				ModeNameAndTemplate: &ModeNameAndTemplate{
					ModeName:     modeName,
					ModeTemplate: modeTemplate,
				},
				PrimeNote: firstNote.Copy(),
			})
		Next:
		}
	}

	return result
}

// SortByModeName sorts the slice with mode names and templates with prime notes by mode name.
func (mtswp ModeTemplatesWithPrime) SortByModeName(desc bool) ModeTemplatesWithPrime {
	sort.Slice(mtswp, func(i, j int) bool {
		switch desc {
		case true:
			return mtswp[i].ModeName > mtswp[j].ModeName
		default:
			return mtswp[i].ModeName < mtswp[j].ModeName
		}
	})

	return mtswp
}

// SortByTemplate sorts the slice with mode names and templates with prime notes by templates.
func (mtswp ModeTemplatesWithPrime) SortByModeTemplate(desc bool) ModeTemplatesWithPrime {
	sort.Slice(mtswp[:], func(i, j int) bool {
		for x := range mtswp[i].ModeTemplate {
			if mtswp[i].ModeTemplate[x] == mtswp[j].ModeTemplate[x] {
				continue
			}

			switch desc {
			case true:
				return mtswp[i].ModeTemplate[x] > mtswp[j].ModeTemplate[x]
			default:
				return mtswp[i].ModeTemplate[x] < mtswp[j].ModeTemplate[x]
			}
		}

		switch desc {
		case true:
			return mtswp[i].ModeName > mtswp[j].ModeName
		default:
			return mtswp[i].ModeName < mtswp[j].ModeName
		}
	})

	return mtswp
}

// SortByPrimeNote sorts the slice with mode names and templates with prime notes by Prime note name.
func (mtswp ModeTemplatesWithPrime) SortByPrimeNote(desc bool) ModeTemplatesWithPrime {
	sort.Slice(mtswp, func(i, j int) bool {
		switch desc {
		case true:
			if !mtswp[i].PrimeNote.IsEqualByName(mtswp[j].PrimeNote) {
				return mtswp[i].PrimeNote.Name() > mtswp[j].PrimeNote.Name()
			}

			return mtswp[i].ModeName > mtswp[j].ModeName
		default:
			if !mtswp[i].PrimeNote.IsEqualByName(mtswp[j].PrimeNote) {
				return mtswp[i].PrimeNote.Name() < mtswp[j].PrimeNote.Name()
			}

			return mtswp[i].ModeName < mtswp[j].ModeName
		}
	})

	return mtswp
}
