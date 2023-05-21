package muse_test

import (
	"fmt"

	"github.com/go-muse/muse"
)

// There is a store with known templates for creating modes.
// It contains mode names and amount of half tones between mode's degrees.
func ExampleInitModeTemplatesStore() {
	mts := muse.InitModeTemplatesStore()

	// Get mode template by mode name
	modeTemplate := mts[muse.ModeNameAeolian]

	fmt.Println(modeTemplate)
	// Output: [2 1 2 2 1 2 2]
}

// You can get the mode templates store as slice and sort it by mode names.
// Also you can specify sorting order.
func ExampleModeNamesAndTemplates_SortByModeName() {
	mts := muse.InitModeTemplatesStore()

	slc := mts.AsSlice()
	slc.SortByModeName(false)

	for _, info := range slc {
		fmt.Println(info.ModeName, info.ModeTemplate)
	}
	// Output: Aeolian [2 1 2 2 1 2 2]
	// AeolianLydian [2 1 2 1 2 2 2]
	// AeolianRais7 [2 1 2 2 1 3 1]
	// Dorian [2 1 2 2 2 1 2]
	// DorianDiminished [2 1 2 1 3 1 2]
	// HarmonicMajor [2 2 1 2 1 3 1]
	// HarmonicMinor [2 1 2 2 1 3 1]
	// HungarianMajor [1 3 1 2 1 3 1]
	// HungarianMinor [2 1 3 1 1 3 1]
	// Ionian [2 2 1 2 2 2 1]
	// IonianAeolian [2 2 1 2 1 2 2]
	// IonianAugmented2 [3 1 1 3 1 2 1]
	// IonianFlat3 [2 1 2 2 2 2 1]
	// IonianFlat6 [2 2 1 2 1 3 1]
	// IonianRais5 [2 2 1 3 1 2 1]
	// Locrian [1 2 2 1 2 2 2]
	// LocrianDoubleFlat3DoubleFlat7 [1 1 3 1 2 1 3]
	// LocrianDoubleFlat7 [1 2 2 1 2 1 3]
	// LocrianRais6 [1 2 2 1 3 1 2]
	// Lydian [2 2 2 1 2 2 1]
	// LydianAugmented [2 2 2 2 1 2 1]
	// LydianAugmented2 [3 1 2 2 1 2 1]
	// LydianDiminished [2 1 3 1 2 2 1]
	// LydianDominant [2 2 2 1 2 1 2]
	// LydianRais2Rais6 [3 1 2 1 3 1 1]
	// LydianRais9 [3 1 2 1 2 2 1]
	// MelodicMajor [2 2 1 2 1 2 2]
	// MelodicMinor [2 1 2 2 2 2 1]
	// MixoLydian [2 2 1 2 2 1 2]
	// MixolydianFlat2 [1 3 1 2 2 1 2]
	// NaturalMajor [2 2 1 2 2 2 1]
	// NaturalMinor [2 1 2 2 1 2 2]
	// Oriental [1 3 1 1 3 1 2]
	// PentatonicBluesMajor [2 3 2 2 3]
	// PentatonicBluesMinor [3 2 3 2 2]
	// PentatonicMajor [2 2 3 2 3]
	// PentatonicMinor [3 2 2 3 2]
	// PentatonicSustained [2 3 2 3 2]
	// Phrygian [1 2 2 2 1 2 2]
	// PhrygianDiminished [1 2 1 3 1 2 2]
	// PhrygianDominant [1 3 1 2 1 2 2]
	// PhrygoDorian [1 2 2 2 2 1 2]
	// SuperLocrian [1 2 1 2 2 2 2]
	// UkrainianDorian [2 1 3 1 2 1 2]
	// UltraLocrian [1 2 1 2 2 1 3]
	// Ultraphrygian [1 2 1 3 1 1 3]
}

// You can get the mode templates store as slice and sort it by mode Templates.
// Also you can specify sorting order.
func ExampleModeNamesAndTemplates_SortByModeTemplate() {
	mts := muse.InitModeTemplatesStore()

	slc := mts.AsSlice()
	slc.SortByModeTemplate(false)

	for _, info := range slc {
		fmt.Println(info.ModeTemplate, info.ModeName)
	}
	// Output: [1 1 3 1 2 1 3] LocrianDoubleFlat3DoubleFlat7
	// [1 2 1 2 2 1 3] UltraLocrian
	// [1 2 1 2 2 2 2] SuperLocrian
	// [1 2 1 3 1 1 3] Ultraphrygian
	// [1 2 1 3 1 2 2] PhrygianDiminished
	// [1 2 2 1 2 1 3] LocrianDoubleFlat7
	// [1 2 2 1 2 2 2] Locrian
	// [1 2 2 1 3 1 2] LocrianRais6
	// [1 2 2 2 1 2 2] Phrygian
	// [1 2 2 2 2 1 2] PhrygoDorian
	// [1 3 1 1 3 1 2] Oriental
	// [1 3 1 2 1 2 2] PhrygianDominant
	// [1 3 1 2 1 3 1] HungarianMajor
	// [1 3 1 2 2 1 2] MixolydianFlat2
	// [2 1 2 1 2 2 2] AeolianLydian
	// [2 1 2 1 3 1 2] DorianDiminished
	// [2 1 2 2 1 2 2] Aeolian
	// [2 1 2 2 1 2 2] NaturalMinor
	// [2 1 2 2 1 3 1] AeolianRais7
	// [2 1 2 2 1 3 1] HarmonicMinor
	// [2 1 2 2 2 1 2] Dorian
	// [2 1 2 2 2 2 1] IonianFlat3
	// [2 1 2 2 2 2 1] MelodicMinor
	// [2 1 3 1 1 3 1] HungarianMinor
	// [2 1 3 1 2 1 2] UkrainianDorian
	// [2 1 3 1 2 2 1] LydianDiminished
	// [2 2 1 2 1 2 2] IonianAeolian
	// [2 2 1 2 1 2 2] MelodicMajor
	// [2 2 1 2 1 3 1] HarmonicMajor
	// [2 2 1 2 1 3 1] IonianFlat6
	// [2 2 1 2 2 1 2] MixoLydian
	// [2 2 1 2 2 2 1] Ionian
	// [2 2 1 2 2 2 1] NaturalMajor
	// [2 2 1 3 1 2 1] IonianRais5
	// [2 2 2 1 2 1 2] LydianDominant
	// [2 2 2 1 2 2 1] Lydian
	// [2 2 2 2 1 2 1] LydianAugmented
	// [2 2 3 2 3] PentatonicMajor
	// [2 3 2 2 3] PentatonicBluesMajor
	// [2 3 2 3 2] PentatonicSustained
	// [3 1 1 3 1 2 1] IonianAugmented2
	// [3 1 2 1 2 2 1] LydianRais9
	// [3 1 2 1 3 1 1] LydianRais2Rais6
	// [3 1 2 2 1 2 1] LydianAugmented2
	// [3 2 2 3 2] PentatonicMinor
	// [3 2 3 2 2] PentatonicBluesMinor
}

// It's possible to find mode templates that match a given pattern.
func ExampleModeTemplatesStore_FindModeTemplatesByPattern() {
	mts := muse.InitModeTemplatesStore()

	// Example of incomplete mode template
	myPattern := muse.ModeTemplate{2, 1, 3, 1, 2}

	// result is the same ModeTemplatesStore instance, but with resulting templates only
	result := mts.FindModeTemplatesByPattern(myPattern).AsSlice().SortByModeTemplate(false)

	for _, info := range result {
		fmt.Println(info.ModeTemplate, info.ModeName)
	}
	// Output: [1 2 1 3 1 2 2] PhrygianDiminished
	// [1 2 2 1 3 1 2] LocrianRais6
	// [2 1 2 1 3 1 2] DorianDiminished
	// [2 1 3 1 2 1 2] UkrainianDorian
	// [2 1 3 1 2 2 1] LydianDiminished
	// [2 2 1 3 1 2 1] IonianRais5
}

// In the mode templates store, it is possible to find mode templates that correspond to a given set of notes.
func ExampleModeTemplatesStore_FindModeTemplatesByNotes() {
	mts := muse.InitModeTemplatesStore()

	notes := []*muse.Note{
		muse.C.MustMakeNote(),
		muse.D.MustMakeNote(),
		muse.E.MustMakeNote(),
		muse.F.MustMakeNote(),
		muse.G.MustMakeNote(),
		muse.A.MustMakeNote(),
		muse.B.MustMakeNote(),
		muse.C.MustMakeNote(), // duplicates are ok
		muse.C.MustMakeNote(), // duplicates are ok
	}

	result := mts.FindModeTemplatesByNotes(notes).SortByPrimeNote(false)

	for _, r := range result {
		fmt.Printf("mode name: %s, mode template: %+v, prime note: %+v, scale: %+v\n",
			r.ModeName,
			r.ModeTemplate,
			r.PrimeNote.Name(),
			muse.MustMakeNewMode(r.ModeName, r.PrimeNote.Name()).GenerateScale(false),
		)
	}
	// Output: mode name: Aeolian, mode template: [2 1 2 2 1 2 2], prime note: A, scale: [A B C D E F G]
	// mode name: NaturalMinor, mode template: [2 1 2 2 1 2 2], prime note: A, scale: [A B C D E F G]
	// mode name: Locrian, mode template: [1 2 2 1 2 2 2], prime note: B, scale: [B C D E F G A]
	// mode name: Ionian, mode template: [2 2 1 2 2 2 1], prime note: C, scale: [C D E F G A B]
	// mode name: NaturalMajor, mode template: [2 2 1 2 2 2 1], prime note: C, scale: [C D E F G A B]
	// mode name: Dorian, mode template: [2 1 2 2 2 1 2], prime note: D, scale: [D E F G A B C]
	// mode name: Phrygian, mode template: [1 2 2 2 1 2 2], prime note: E, scale: [E F G A B C D]
	// mode name: Lydian, mode template: [2 2 2 1 2 2 1], prime note: F, scale: [F G A B C D E]
	// mode name: MixoLydian, mode template: [2 2 1 2 2 1 2], prime note: G, scale: [G A B C D E F]
}
