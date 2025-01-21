package mode

import (
	"errors"
	"fmt"
)

// ErrNameUnknown is returned when mode name is unknown.
var ErrNameUnknown = errors.New("unknown mode name")

// GetTemplateByName returns mode template by mode's name.
func GetTemplateByName(modeName Name) (Template, error) {
	switch modeName {
	// Tonal modes
	case NameNaturalMinor:
		return TemplateNaturalMinor(), nil
	case NameHarmonicMinor:
		return TemplateHarmonicMinor(), nil
	case NameMelodicMinor:
		return TemplateMelodicMinor(), nil
	case NameNaturalMajor:
		return TemplateNaturalMajor(), nil
	case NameHarmonicMajor:
		return TemplateHarmonicMajor(), nil
	case NameMelodicMajor:
		return TemplateMelodicMajor(), nil

	// Modes of the Major scale
	case NameIonian:
		return TemplateIonian(), nil
	case NameDorian:
		return TemplateDorian(), nil
	case NameAeolian:
		return TemplateAeolian(), nil
	case NameLydian:
		return TemplateLydian(), nil
	case NameMixoLydian:
		return TemplateMixoLydian(), nil
	case NamePhrygian:
		return TemplatePhrygian(), nil
	case NameLocrian:
		return TemplateLocrian(), nil

	// Modes Of The Melodic Minor scale
	case NameIonianFlat3:
		return TemplateIonianFlat3(), nil
	case NamePhrygoDorian:
		return TemplatePhrygoDorian(), nil
	case NameLydianAugmented:
		return TemplateLydianAugmented(), nil
	case NameLydianDominant:
		return TemplateLydianDominant(), nil
	case NameIonianAeolian:
		return TemplateIonianAeolian(), nil
	case NameAeolianLydian:
		return TemplateAeolianLydian(), nil
	case NameSuperLocrian:
		return TemplateSuperLocrian(), nil

	// Modes of the Harmonic Minor scale
	case NameAeolianRais7:
		return TemplateAeolianRais7(), nil
	case NameLocrianRais6:
		return TemplateLocrianRais6(), nil
	case NameIonianRais5:
		return TemplateIonianRais5(), nil
	case NameUkrainianDorian:
		return TemplateUkrainianDorian(), nil
	case NamePhrygianDominant:
		return TemplatePhrygianDominant(), nil
	case NameLydianRais9:
		return TemplateLydianRais9(), nil
	case NameUltraLocrian:
		return TemplateUltraLocrian(), nil

	// Modes Of The Harmonic Major scale
	case NameIonianFlat6:
		return TemplateIonianFlat6(), nil
	case NameDorianDiminished:
		return TemplateDorianDiminished(), nil
	case NamePhrygianDiminished:
		return TemplatePhrygianDiminished(), nil
	case NameLydianDiminished:
		return TemplateLydianDiminished(), nil
	case NameMixolydianFlat2:
		return TemplateMixolydianFlat2(), nil
	case NameLydianAugmented2:
		return TemplateLydianAugmented2(), nil
	case NameLocrianDoubleFlat7:
		return TemplateLocrianDoubleFlat7(), nil

	// Double Harmonic Major Modes
	case NameHungarianMajor:
		return TemplateHungarianMajor(), nil
	case NameLydianRais2Rais6:
		return TemplateLydianRais2Rais6(), nil
	case NameUltraPhrygian:
		return TemplateUltraphrygian(), nil
	case NameHungarianMinor:
		return TemplateHungarianMinor(), nil
	case NameOriental:
		return TemplateOriental(), nil
	case NameIonianAugmented2:
		return TemplateIonianAugmented2(), nil
	case NameLocrianDoubleFlat3DoubleFlat7:
		return TemplateLocrianDoubleFlat3DoubleFlat7(), nil

	// Pentatonics

	case NamePentatonicMajor:
		return TemplatePentatonicMajor(), nil
	case NamePentatonicSustained:
		return TemplatePentatonicSustained(), nil
	case NamePentatonicBluesMinor:
		return TemplatePentatonicBluesMinor(), nil
	case NamePentatonicBluesMajor:
		return TemplatePentatonicBluesMajor(), nil
	case NamePentatonicMinor:
		return TemplatePentatonicMinor(), nil

	case NamePentatonicHirajoshi:
		return TemplatePentatonicHirajoshi(), nil
	case NamePentatonicIwato:
		return TemplatePentatonicIwato(), nil
	case NamePentatonicHonKumoiShiouzhi:
		return TemplatePentatonicHonKumoiShiouzhi(), nil
	case NamePentatonicHonKumoiJoshi:
		return TemplatePentatonicHonKumoiJoshi(), nil
	case NamePentatonicLydianPentatonic:
		return TemplatePentatonicLydianPentatonic(), nil
	}

	return nil, fmt.Errorf("got: '%s: %w", modeName, ErrNameUnknown)
}

// Tonal modes

func TemplateNaturalMinor() Template {
	return Template{2, 1, 2, 2, 1, 2, 2}
}

func TemplateMelodicMinor() Template {
	return Template{2, 1, 2, 2, 2, 2, 1}
}

func TemplateHarmonicMinor() Template {
	return Template{2, 1, 2, 2, 1, 3, 1}
}

func TemplateNaturalMajor() Template {
	return Template{2, 2, 1, 2, 2, 2, 1}
}

func TemplateMelodicMajor() Template {
	return Template{2, 2, 1, 2, 1, 2, 2}
}

func TemplateHarmonicMajor() Template {
	return Template{2, 2, 1, 2, 1, 3, 1}
}

// Modes of the Major scale

func TemplateIonian() Template {
	return Template{2, 2, 1, 2, 2, 2, 1}
}

func TemplateDorian() Template {
	return Template{2, 1, 2, 2, 2, 1, 2}
}

func TemplateAeolian() Template {
	return Template{2, 1, 2, 2, 1, 2, 2}
}

func TemplateLydian() Template {
	return Template{2, 2, 2, 1, 2, 2, 1}
}

func TemplateMixoLydian() Template {
	return Template{2, 2, 1, 2, 2, 1, 2}
}

func TemplatePhrygian() Template {
	return Template{1, 2, 2, 2, 1, 2, 2}
}

func TemplateLocrian() Template {
	return Template{1, 2, 2, 1, 2, 2, 2}
}

// Modes Of The Melodic Minor scale

func TemplateIonianFlat3() Template {
	return Template{2, 1, 2, 2, 2, 2, 1}
}

func TemplatePhrygoDorian() Template {
	return Template{1, 2, 2, 2, 2, 1, 2}
}

func TemplateLydianAugmented() Template {
	return Template{2, 2, 2, 2, 1, 2, 1}
}

func TemplateLydianDominant() Template {
	return Template{2, 2, 2, 1, 2, 1, 2}
}

func TemplateIonianAeolian() Template {
	return Template{2, 2, 1, 2, 1, 2, 2}
}

func TemplateAeolianLydian() Template {
	return Template{2, 1, 2, 1, 2, 2, 2}
}

func TemplateSuperLocrian() Template {
	return Template{1, 2, 1, 2, 2, 2, 2}
}

// Modes of the Harmonic Minor scale

func TemplateAeolianRais7() Template {
	return Template{2, 1, 2, 2, 1, 3, 1}
}

func TemplateLocrianRais6() Template {
	return Template{1, 2, 2, 1, 3, 1, 2}
}

func TemplateIonianRais5() Template {
	return Template{2, 2, 1, 3, 1, 2, 1}
}

func TemplateUkrainianDorian() Template {
	return Template{2, 1, 3, 1, 2, 1, 2}
}

func TemplatePhrygianDominant() Template {
	return Template{1, 3, 1, 2, 1, 2, 2}
}

func TemplateLydianRais9() Template {
	return Template{3, 1, 2, 1, 2, 2, 1}
}

func TemplateUltraLocrian() Template {
	return Template{1, 2, 1, 2, 2, 1, 3}
}

// Modes Of The Harmonic Major scale

func TemplateIonianFlat6() Template {
	return Template{2, 2, 1, 2, 1, 3, 1}
}

func TemplateDorianDiminished() Template {
	return Template{2, 1, 2, 1, 3, 1, 2}
}

func TemplatePhrygianDiminished() Template {
	return Template{1, 2, 1, 3, 1, 2, 2}
}

func TemplateLydianDiminished() Template {
	return Template{2, 1, 3, 1, 2, 2, 1}
}

func TemplateMixolydianFlat2() Template {
	return Template{1, 3, 1, 2, 2, 1, 2}
}

func TemplateLydianAugmented2() Template {
	return Template{3, 1, 2, 2, 1, 2, 1}
}

func TemplateLocrianDoubleFlat7() Template {
	return Template{1, 2, 2, 1, 2, 1, 3}
}

// Double Harmonic Major Modes

func TemplateHungarianMajor() Template {
	return Template{1, 3, 1, 2, 1, 3, 1}
}

func TemplateLydianRais2Rais6() Template {
	return Template{3, 1, 2, 1, 3, 1, 1}
}

func TemplateUltraphrygian() Template {
	return Template{1, 2, 1, 3, 1, 1, 3}
}

func TemplateHungarianMinor() Template {
	return Template{2, 1, 3, 1, 1, 3, 1}
}

func TemplateOriental() Template {
	return Template{1, 3, 1, 1, 3, 1, 2}
}

func TemplateIonianAugmented2() Template {
	return Template{3, 1, 1, 3, 1, 2, 1}
}

func TemplateLocrianDoubleFlat3DoubleFlat7() Template {
	return Template{1, 1, 3, 1, 2, 1, 3}
}

// Pentatonics

// Main pentatonics

func TemplatePentatonicMajor() Template {
	return Template{2, 2, 3, 2, 3}
}

func TemplatePentatonicSustained() Template {
	return Template{2, 3, 2, 3, 2}
}

func TemplatePentatonicBluesMinor() Template {
	return Template{3, 2, 3, 2, 2}
}

func TemplatePentatonicBluesMajor() Template {
	return Template{2, 3, 2, 2, 3}
}

func TemplatePentatonicMinor() Template {
	return Template{3, 2, 2, 3, 2}
}

// Japanese pentatonics

func TemplatePentatonicHirajoshi() Template {
	return Template{2, 1, 4, 1, 4}
}

func TemplatePentatonicIwato() Template {
	return Template{1, 4, 1, 4, 2}
}

func TemplatePentatonicHonKumoiShiouzhi() Template {
	return Template{4, 1, 4, 2, 1}
}

func TemplatePentatonicHonKumoiJoshi() Template {
	return Template{1, 4, 2, 1, 4}
}

func TemplatePentatonicLydianPentatonic() Template {
	return Template{4, 2, 1, 4, 1}
}

// others
