package muse_test

import (
	"fmt"
	"sort"

	"github.com/go-muse/muse"
)

func Example() {
	mts := muse.InitModeTemplatesStore()

	type info struct {
		Name     muse.ModeName
		Template muse.ModeTemplate
	}

	m := make([]info, 0, len(mts))

	for moteTemplateName, modeTemplate := range mts {
		m = append(m, info{moteTemplateName, modeTemplate})
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].Name < m[j].Name
	})

	for _, info := range m {
		fmt.Println(info)
	}

	// Output: {Aeolian [2 1 2 2 1 2 2]}
	// {AeolianLydian [2 1 2 1 2 2 2]}
	// {AeolianRais7  [2 1 2 2 1 3 1]}
	// {Dorian [2 1 2 2 2 1 2]}
	// {DorianDiminished [2 1 2 1 3 1 2]}
	// {HarmonicMajor [2 2 1 2 1 3 1]}
	// {HarmonicMinor [2 1 2 2 1 3 1]}
	// {HungarianMajor [1 3 1 2 1 3 1]}
	// {HungarianMinor [2 1 3 1 1 3 1]}
	// {Ionian [2 2 1 2 2 2 1]}
	// {IonianAeolian [2 2 1 2 1 2 2]}
	// {IonianAugmented2 [3 1 1 3 1 2 1]}
	// {IonianFlat3 [2 1 2 2 2 2 1]}
	// {IonianFlat6  [2 2 1 2 1 3 1]}
	// {IonianRais5 [2 2 1 3 1 2 1]}
	// {Locrian [1 2 2 1 2 2 2]}
	// {LocrianDoubleFlat3DoubleFlat7 [1 1 3 1 2 1 3]}
	// {LocrianDoubleFlat7 [1 2 2 1 2 1 3]}
	// {LocrianRais6 [1 2 2 1 3 1 2]}
	// {Lydian [2 2 2 1 2 2 1]}
	// {LydianAugmented [2 2 2 2 1 2 1]}
	// {LydianAugmented2 [3 1 2 2 1 2 1]}
	// {LydianDiminished [2 1 3 1 2 2 1]}
	// {LydianDominant [2 2 2 1 2 1 2]}
	// {LydianRais2Rais6 [3 1 2 1 3 1 1]}
	// {LydianRais9 [3 1 2 1 2 2 1]}
	// {MelodicMajor [2 2 1 2 1 2 2]}
	// {MelodicMinor [2 1 2 2 2 2 1]}
	// {MixoLydian [2 2 1 2 2 1 2]}
	// {MixolydianFlat2 [1 3 1 2 2 1 2]}
	// {NaturalMajor [2 2 1 2 2 2 1]}
	// {NaturalMinor [2 1 2 2 1 2 2]}
	// {Oriental [1 3 1 1 3 1 2]}
	// {PentatonicBluesMajor [2 3 2 2 3]}
	// {PentatonicBluesMinor [3 2 3 2 2]}
	// {PentatonicMajor [2 2 3 2 3]}
	// {PentatonicMinor [3 2 2 3 2]}
	// {PentatonicSustained [2 3 2 3 2]}
	// {Phrygian [1 2 2 2 1 2 2]}
	// {PhrygianDiminished [1 2 1 3 1 2 2]}
	// {PhrygianDominant [1 3 1 2 1 2 2]}
	// {PhrygoDorian [1 2 2 2 2 1 2]}
	// {SuperLocrian [1 2 1 2 2 2 2]}
	// {UkrainianDorian [2 1 3 1 2 1 2]}
	// {UltraLocrian [1 2 1 2 2 1 3]}
	// {Ultraphrygian [1 2 1 3 1 1 3]}
}
