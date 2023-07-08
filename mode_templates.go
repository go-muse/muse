package muse

import "github.com/pkg/errors"

var ErrModeNameUnknown = errors.New("unknown mode name")

// GetTemplateByModeName returns mode template by mode's name.
func GetTemplateByModeName(modeName ModeName) (ModeTemplate, error) {
	switch modeName {
	// Tonal modes
	case ModeNameNaturalMinor:
		return TemplateNaturalMinor(), nil
	case ModeNameHarmonicMinor:
		return TemplateHarmonicMinor(), nil
	case ModeNameMelodicMinor:
		return TemplateMelodicMinor(), nil
	case ModeNameNaturalMajor:
		return TemplateNaturalMajor(), nil
	case ModeNameHarmonicMajor:
		return TemplateHarmonicMajor(), nil
	case ModeNameMelodicMajor:
		return TemplateMelodicMajor(), nil

	// Modes of the Major scale
	case ModeNameIonian:
		return TemplateIonian(), nil
	case ModeNameDorian:
		return TemplateDorian(), nil
	case ModeNameAeolian:
		return TemplateAeolian(), nil
	case ModeNameLydian:
		return TemplateLydian(), nil
	case ModeNameMixoLydian:
		return TemplateMixoLydian(), nil
	case ModeNamePhrygian:
		return TemplatePhrygian(), nil
	case ModeNameLocrian:
		return TemplateLocrian(), nil

	// Modes Of The Melodic Minor scale
	case ModeNameIonianFlat3:
		return TemplateIonianFlat3(), nil
	case ModeNamePhrygoDorian:
		return TemplatePhrygoDorian(), nil
	case ModeNameLydianAugmented:
		return TemplateLydianAugmented(), nil
	case ModeNameLydianDominant:
		return TemplateLydianDominant(), nil
	case ModeNameIonianAeolian:
		return TemplateIonianAeolian(), nil
	case ModeNameAeolianLydian:
		return TemplateAeolianLydian(), nil
	case ModeNameSuperLocrian:
		return TemplateSuperLocrian(), nil

	// Modes of the Harmonic Minor scale
	case ModeNameAeolianRais7:
		return TemplateAeolianRais7(), nil
	case ModeNameLocrianRais6:
		return TemplateLocrianRais6(), nil
	case ModeNameIonianRais5:
		return TemplateIonianRais5(), nil
	case ModeNameUkrainianDorian:
		return TemplateUkrainianDorian(), nil
	case ModeNamePhrygianDominant:
		return TemplatePhrygianDominant(), nil
	case ModeNameLydianRais9:
		return TemplateLydianRais9(), nil
	case ModeNameUltraLocrian:
		return TemplateUltraLocrian(), nil

	// Modes Of The Harmonic Major scale
	case ModeNameIonianFlat6:
		return TemplateIonianFlat6(), nil
	case ModeNameDorianDiminished:
		return TemplateDorianDiminished(), nil
	case ModeNamePhrygianDiminished:
		return TemplatePhrygianDiminished(), nil
	case ModeNameLydianDiminished:
		return TemplateLydianDiminished(), nil
	case ModeNameMixolydianFlat2:
		return TemplateMixolydianFlat2(), nil
	case ModeNameLydianAugmented2:
		return TemplateLydianAugmented2(), nil
	case ModeNameLocrianDoubleFlat7:
		return TemplateLocrianDoubleFlat7(), nil

	// Double Harmonic Major Modes
	case ModeNameHungarianMajor:
		return TemplateHungarianMajor(), nil
	case ModeNameLydianRais2Rais6:
		return TemplateLydianRais2Rais6(), nil
	case ModeNameUltraphrygian:
		return TemplateUltraphrygian(), nil
	case ModeNameHungarianMinor:
		return TemplateHungarianMinor(), nil
	case ModeNameOriental:
		return TemplateOriental(), nil
	case ModeNameIonianAugmented2:
		return TemplateIonianAugmented2(), nil
	case ModeNameLocrianDoubleFlat3DoubleFlat7:
		return TemplateLocrianDoubleFlat3DoubleFlat7(), nil

	// Pentatonics

	case ModeNamePentatonicMajor:
		return TemplatePentatonicMajor(), nil
	case ModeNamePentatonicSustained:
		return TemplatePentatonicSustained(), nil
	case ModeNamePentatonicBluesMinor:
		return TemplatePentatonicBluesMinor(), nil
	case ModeNamePentatonicBluesMajor:
		return TemplatePentatonicBluesMajor(), nil
	case ModeNamePentatonicMinor:
		return TemplatePentatonicMinor(), nil

	case ModeNamePentatonicHirajoshi:
		return TemplatePentatonicHirajoshi(), nil
	case ModeNamePentatonicIwato:
		return TemplatePentatonicIwato(), nil
	case ModeNamePentatonicHonKumoiShiouzhi:
		return TemplatePentatonicHonKumoiShiouzhi(), nil
	case ModeNamePentatonicHonKumoiJoshi:
		return TemplatePentatonicHonKumoiJoshi(), nil
	case ModeNamePentatonicLydianPentatonic:
		return TemplatePentatonicLydianPentatonic(), nil
	}

	return nil, errors.Wrapf(ErrModeNameUnknown, "got: %s", modeName)
}

// Tonal modes

func TemplateNaturalMinor() ModeTemplate {
	return ModeTemplate{2, 1, 2, 2, 1, 2, 2}
}

func TemplateMelodicMinor() ModeTemplate {
	return ModeTemplate{2, 1, 2, 2, 2, 2, 1}
}

func TemplateHarmonicMinor() ModeTemplate {
	return ModeTemplate{2, 1, 2, 2, 1, 3, 1}
}

func TemplateNaturalMajor() ModeTemplate {
	return ModeTemplate{2, 2, 1, 2, 2, 2, 1}
}

func TemplateMelodicMajor() ModeTemplate {
	return ModeTemplate{2, 2, 1, 2, 1, 2, 2}
}

func TemplateHarmonicMajor() ModeTemplate {
	return ModeTemplate{2, 2, 1, 2, 1, 3, 1}
}

// Modes of the Major scale

func TemplateIonian() ModeTemplate {
	return ModeTemplate{2, 2, 1, 2, 2, 2, 1}
}

func TemplateDorian() ModeTemplate {
	return ModeTemplate{2, 1, 2, 2, 2, 1, 2}
}

func TemplateAeolian() ModeTemplate {
	return ModeTemplate{2, 1, 2, 2, 1, 2, 2}
}

func TemplateLydian() ModeTemplate {
	return ModeTemplate{2, 2, 2, 1, 2, 2, 1}
}

func TemplateMixoLydian() ModeTemplate {
	return ModeTemplate{2, 2, 1, 2, 2, 1, 2}
}

func TemplatePhrygian() ModeTemplate {
	return ModeTemplate{1, 2, 2, 2, 1, 2, 2}
}

func TemplateLocrian() ModeTemplate {
	return ModeTemplate{1, 2, 2, 1, 2, 2, 2}
}

// Modes Of The Melodic Minor scale

func TemplateIonianFlat3() ModeTemplate {
	return ModeTemplate{2, 1, 2, 2, 2, 2, 1}
}

func TemplatePhrygoDorian() ModeTemplate {
	return ModeTemplate{1, 2, 2, 2, 2, 1, 2}
}

func TemplateLydianAugmented() ModeTemplate {
	return ModeTemplate{2, 2, 2, 2, 1, 2, 1}
}

func TemplateLydianDominant() ModeTemplate {
	return ModeTemplate{2, 2, 2, 1, 2, 1, 2}
}

func TemplateIonianAeolian() ModeTemplate {
	return ModeTemplate{2, 2, 1, 2, 1, 2, 2}
}

func TemplateAeolianLydian() ModeTemplate {
	return ModeTemplate{2, 1, 2, 1, 2, 2, 2}
}

func TemplateSuperLocrian() ModeTemplate {
	return ModeTemplate{1, 2, 1, 2, 2, 2, 2}
}

// Modes of the Harmonic Minor scale

func TemplateAeolianRais7() ModeTemplate {
	return ModeTemplate{2, 1, 2, 2, 1, 3, 1}
}

func TemplateLocrianRais6() ModeTemplate {
	return ModeTemplate{1, 2, 2, 1, 3, 1, 2}
}

func TemplateIonianRais5() ModeTemplate {
	return ModeTemplate{2, 2, 1, 3, 1, 2, 1}
}

func TemplateUkrainianDorian() ModeTemplate {
	return ModeTemplate{2, 1, 3, 1, 2, 1, 2}
}

func TemplatePhrygianDominant() ModeTemplate {
	return ModeTemplate{1, 3, 1, 2, 1, 2, 2}
}

func TemplateLydianRais9() ModeTemplate {
	return ModeTemplate{3, 1, 2, 1, 2, 2, 1}
}

func TemplateUltraLocrian() ModeTemplate {
	return ModeTemplate{1, 2, 1, 2, 2, 1, 3}
}

// Modes Of The Harmonic Major scale

func TemplateIonianFlat6() ModeTemplate {
	return ModeTemplate{2, 2, 1, 2, 1, 3, 1}
}

func TemplateDorianDiminished() ModeTemplate {
	return ModeTemplate{2, 1, 2, 1, 3, 1, 2}
}

func TemplatePhrygianDiminished() ModeTemplate {
	return ModeTemplate{1, 2, 1, 3, 1, 2, 2}
}

func TemplateLydianDiminished() ModeTemplate {
	return ModeTemplate{2, 1, 3, 1, 2, 2, 1}
}

func TemplateMixolydianFlat2() ModeTemplate {
	return ModeTemplate{1, 3, 1, 2, 2, 1, 2}
}

func TemplateLydianAugmented2() ModeTemplate {
	return ModeTemplate{3, 1, 2, 2, 1, 2, 1}
}

func TemplateLocrianDoubleFlat7() ModeTemplate {
	return ModeTemplate{1, 2, 2, 1, 2, 1, 3}
}

// Double Harmonic Major Modes

func TemplateHungarianMajor() ModeTemplate {
	return ModeTemplate{1, 3, 1, 2, 1, 3, 1}
}

func TemplateLydianRais2Rais6() ModeTemplate {
	return ModeTemplate{3, 1, 2, 1, 3, 1, 1}
}

func TemplateUltraphrygian() ModeTemplate {
	return ModeTemplate{1, 2, 1, 3, 1, 1, 3}
}

func TemplateHungarianMinor() ModeTemplate {
	return ModeTemplate{2, 1, 3, 1, 1, 3, 1}
}

func TemplateOriental() ModeTemplate {
	return ModeTemplate{1, 3, 1, 1, 3, 1, 2}
}

func TemplateIonianAugmented2() ModeTemplate {
	return ModeTemplate{3, 1, 1, 3, 1, 2, 1}
}

func TemplateLocrianDoubleFlat3DoubleFlat7() ModeTemplate {
	return ModeTemplate{1, 1, 3, 1, 2, 1, 3}
}

// Pentatonics

// Main pentatonics

func TemplatePentatonicMajor() ModeTemplate {
	return ModeTemplate{2, 2, 3, 2, 3}
}

func TemplatePentatonicSustained() ModeTemplate {
	return ModeTemplate{2, 3, 2, 3, 2}
}

func TemplatePentatonicBluesMinor() ModeTemplate {
	return ModeTemplate{3, 2, 3, 2, 2}
}

func TemplatePentatonicBluesMajor() ModeTemplate {
	return ModeTemplate{2, 3, 2, 2, 3}
}

func TemplatePentatonicMinor() ModeTemplate {
	return ModeTemplate{3, 2, 2, 3, 2}
}

// Japanese pentatonics

func TemplatePentatonicHirajoshi() ModeTemplate {
	return ModeTemplate{2, 1, 4, 1, 4}
}

func TemplatePentatonicIwato() ModeTemplate {
	return ModeTemplate{1, 4, 1, 4, 2}
}

func TemplatePentatonicHonKumoiShiouzhi() ModeTemplate {
	return ModeTemplate{4, 1, 4, 2, 1}
}

func TemplatePentatonicHonKumoiJoshi() ModeTemplate {
	return ModeTemplate{1, 4, 2, 1, 4}
}

func TemplatePentatonicLydianPentatonic() ModeTemplate {
	return ModeTemplate{4, 2, 1, 4, 1}
}

// others
