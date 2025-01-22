package mode

import (
	"sort"

	"github.com/go-muse/muse/note"
)

type TemplatesStore map[Name]Template

func (ts TemplatesStore) AddTemplate(modeName Name, moteTemplate Template) {
	ts[modeName] = moteTemplate
}

func InitTemplatesStore() TemplatesStore {
	ts := make(TemplatesStore)

	// Tonal modes
	ts[NameNaturalMinor] = TemplateNaturalMinor()
	ts[NameMelodicMinor] = TemplateMelodicMinor()
	ts[NameHarmonicMinor] = TemplateHarmonicMinor()
	ts[NameNaturalMajor] = TemplateNaturalMajor()
	ts[NameMelodicMajor] = TemplateMelodicMajor()
	ts[NameHarmonicMajor] = TemplateHarmonicMajor()

	// Modes of the Major scale
	ts[NameIonian] = TemplateIonian()
	ts[NameDorian] = TemplateDorian()
	ts[NameAeolian] = TemplateAeolian()
	ts[NameLydian] = TemplateLydian()
	ts[NameMixoLydian] = TemplateMixoLydian()
	ts[NamePhrygian] = TemplatePhrygian()
	ts[NameLocrian] = TemplateLocrian()

	// Modes Of The Melodic Minor scale
	ts[NameIonianFlat3] = TemplateIonianFlat3()
	ts[NamePhrygoDorian] = TemplatePhrygoDorian()
	ts[NameLydianAugmented] = TemplateLydianAugmented()
	ts[NameLydianDominant] = TemplateLydianDominant()
	ts[NameIonianAeolian] = TemplateIonianAeolian()
	ts[NameAeolianLydian] = TemplateAeolianLydian()
	ts[NameSuperLocrian] = TemplateSuperLocrian()

	// Modes of the Harmonic Minor scale
	ts[NameAeolianRais7] = TemplateAeolianRais7()
	ts[NameLocrianRais6] = TemplateLocrianRais6()
	ts[NameIonianRais5] = TemplateIonianRais5()
	ts[NameUkrainianDorian] = TemplateUkrainianDorian()
	ts[NamePhrygianDominant] = TemplatePhrygianDominant()
	ts[NameLydianRais9] = TemplateLydianRais9()
	ts[NameUltraLocrian] = TemplateUltraLocrian()

	// Modes Of The Harmonic Major scale
	ts[NameIonianFlat6] = TemplateIonianFlat6()
	ts[NameDorianDiminished] = TemplateDorianDiminished()
	ts[NamePhrygianDiminished] = TemplatePhrygianDiminished()
	ts[NameLydianDiminished] = TemplateLydianDiminished()
	ts[NameMixolydianFlat2] = TemplateMixolydianFlat2()
	ts[NameLydianAugmented2] = TemplateLydianAugmented2()
	ts[NameLocrianDoubleFlat7] = TemplateLocrianDoubleFlat7()

	// Double Harmonic Major Modes
	ts[NameHungarianMajor] = TemplateHungarianMajor()
	ts[NameLydianRais2Rais6] = TemplateLydianRais2Rais6()
	ts[NameUltraPhrygian] = TemplateUltraphrygian()
	ts[NameHungarianMinor] = TemplateHungarianMinor()
	ts[NameOriental] = TemplateOriental()
	ts[NameIonianAugmented2] = TemplateIonianAugmented2()
	ts[NameLocrianDoubleFlat3DoubleFlat7] = TemplateLocrianDoubleFlat3DoubleFlat7()

	// Pentatonic modes

	// Main pentatonic modes
	ts[NamePentatonicMajor] = TemplatePentatonicMajor()
	ts[NamePentatonicSustained] = TemplatePentatonicSustained()
	ts[NamePentatonicBluesMinor] = TemplatePentatonicBluesMinor()
	ts[NamePentatonicBluesMajor] = TemplatePentatonicBluesMajor()
	ts[NamePentatonicMinor] = TemplatePentatonicMinor()

	return ts
}

// FindModeTemplatesByPattern searches mode templates by intersection with the given mode template.
func (ts TemplatesStore) FindModeTemplatesByPattern(modeTemplate Template) TemplatesStore {
	mt := modeTemplate
	result := make(TemplatesStore)
	for modeName, storedModeTemplate := range ts {
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

func (ts TemplatesStore) Contains(modeName Name) bool {
	if _, ok := ts[modeName]; ok {
		return true
	}

	return false
}

// AsSlice returns map of mode names and templates as slice of structs with ste same information.
func (ts TemplatesStore) AsSlice() NamesAndTemplates {
	nat := make(NamesAndTemplates, 0, len(ts))
	for moteTemplateName, modeTemplate := range ts {
		nat = append(nat, NameAndTemplate{moteTemplateName, modeTemplate})
	}

	return nat
}

// NameAndTemplate is a struct with mode name and template.
type NameAndTemplate struct {
	Name         Name
	ModeTemplate Template
}

// NamesAndTemplates is a slice of structs with mode name and template.
type NamesAndTemplates []NameAndTemplate

// SortByName sorts the slice with mode names and templates by mode name.
func (nat NamesAndTemplates) SortByName(desc bool) NamesAndTemplates {
	sort.Slice(nat, func(i, j int) bool {
		switch desc {
		case true:
			return nat[i].Name > nat[j].Name
		default:
			return nat[i].Name < nat[j].Name
		}
	})

	return nat
}

// SortByTemplate sorts the slice with mode names and templates by templates.
func (nat NamesAndTemplates) SortByTemplate(desc bool) NamesAndTemplates {
	sort.Slice(nat[:], func(i, j int) bool {
		for x := range nat[i].ModeTemplate {
			if nat[i].ModeTemplate[x] == nat[j].ModeTemplate[x] {
				continue
			}

			switch desc {
			case true:
				return nat[i].ModeTemplate[x] > nat[j].ModeTemplate[x]
			default:
				return nat[i].ModeTemplate[x] < nat[j].ModeTemplate[x]
			}
		}

		switch desc {
		case true:
			return nat[i].Name > nat[j].Name
		default:
			return nat[i].Name < nat[j].Name
		}
	})

	return nat
}

type TemplateWithPrime struct {
	*NameAndTemplate
	PrimeNote *note.Note
}

type TemplatesWithPrime []TemplateWithPrime

func (tswp TemplatesWithPrime) Contains(modeName Name) bool {
	for _, mtwp := range tswp {
		if mtwp.Name == modeName {
			return true
		}
	}

	return false
}

// FindModeTemplatesByNotes searches for modes in the storage
// that correspond to the given set of notes and returns a new storage with the found modes.
func (ts TemplatesStore) FindModeTemplatesByNotes(ns note.Notes) TemplatesWithPrime {
	result := make(TemplatesWithPrime, 0)
	var mode *Mode
	for modeName, modeTemplate := range ts {
		notesWithAlterations := note.GetNotesWithAlterations(ns, 0)
		allNotes := make(note.Notes, len(ns)+len(notesWithAlterations))
		copy(allNotes, ns)
		allNotes = append(allNotes, notesWithAlterations...)

		for _, firstNote := range allNotes.Uniques() {
			mode = MustMakeNewMode(modeName, firstNote.Name())
			for _, note := range ns {
				if !mode.Contains(note) {
					goto Next
				}
			}

			result = append(result, TemplateWithPrime{
				NameAndTemplate: &NameAndTemplate{
					Name:         modeName,
					ModeTemplate: modeTemplate,
				},
				PrimeNote: firstNote.Copy(),
			})
		Next:
		}
	}

	return result
}

// SortByName sorts the slice with mode names and templates with prime notes by mode name.
func (tswp TemplatesWithPrime) SortByName(desc bool) TemplatesWithPrime {
	sort.Slice(tswp, func(i, j int) bool {
		switch desc {
		case true:
			return tswp[i].Name > tswp[j].Name
		default:
			return tswp[i].Name < tswp[j].Name
		}
	})

	return tswp
}

// SortByModeTemplate sorts the slice with mode names and templates with prime notes by templates.
func (tswp TemplatesWithPrime) SortByModeTemplate(desc bool) TemplatesWithPrime {
	sort.Slice(tswp[:], func(i, j int) bool {
		for x := range tswp[i].ModeTemplate {
			if tswp[i].ModeTemplate[x] == tswp[j].ModeTemplate[x] {
				continue
			}

			switch desc {
			case true:
				return tswp[i].ModeTemplate[x] > tswp[j].ModeTemplate[x]
			default:
				return tswp[i].ModeTemplate[x] < tswp[j].ModeTemplate[x]
			}
		}

		switch desc {
		case true:
			return tswp[i].Name > tswp[j].Name
		default:
			return tswp[i].Name < tswp[j].Name
		}
	})

	return tswp
}

// SortByPrimeNote sorts the slice with mode names and templates with prime notes by Prime note name.
func (tswp TemplatesWithPrime) SortByPrimeNote(desc bool) TemplatesWithPrime {
	sort.Slice(tswp, func(i, j int) bool {
		switch desc {
		case true:
			if !tswp[i].PrimeNote.IsEqualByName(tswp[j].PrimeNote) {
				return tswp[i].PrimeNote.Name() > tswp[j].PrimeNote.Name()
			}

			return tswp[i].Name > tswp[j].Name
		default:
			if !tswp[i].PrimeNote.IsEqualByName(tswp[j].PrimeNote) {
				return tswp[i].PrimeNote.Name() < tswp[j].PrimeNote.Name()
			}

			return tswp[i].Name < tswp[j].Name
		}
	})

	return tswp
}
