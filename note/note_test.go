package note

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-muse/muse/common/fraction"
	"github.com/go-muse/muse/duration"
	"github.com/go-muse/muse/octave"
	"github.com/go-muse/muse/tuplet"
)

func Test_NewNote(t *testing.T) {
	expectedNoteName := C
	assert.Equal(t, expectedNoteName, newNote(C).name)
}

func TestNewNote(t *testing.T) {
	t.Run("TestNewNote: valid note name", func(t *testing.T) {
		noteNames := []Name{
			C,
			CFLAT,
			CFLAT2,
			CSHARP,
			CSHARP2,
			D,
			DFLAT,
			DFLAT2,
			DSHARP,
			DSHARP2,
			E,
			EFLAT,
			EFLAT2,
			ESHARP,
			ESHARP2,
			F,
			FFLAT,
			FFLAT2,
			FSHARP,
			FSHARP2,
			G,
			GFLAT,
			GFLAT2,
			GSHARP,
			GSHARP2,
			A,
			AFLAT,
			AFLAT2,
			ASHARP,
			ASHARP2,
			B,
			BFLAT,
			BFLAT2,
			BSHARP,
			BSHARP2,
		}
		var newNote *Note
		var err error
		for _, noteName := range noteNames {
			// setup: create a Note with a valid name
			newNote, err = New(noteName)
			require.NoError(t, err)
			assert.NotNil(t, newNote, "expected note from note name %s", noteName)

			// assert that the returned name matches the expected name
			if newNote.name != noteName {
				t.Errorf("Expected note name to be '%s', but got '%s'", noteName, newNote.name)
			}
		}
	})

	t.Run("TestNewNote: invalid note name", func(t *testing.T) {
		// setup: create a Note with invalid name
		expectedName := Name("Hello")
		newNote, err := New(expectedName)
		// assert that the returned error matches the expected error
		require.ErrorIs(t, err, ErrNoteNameUnknown)
		assert.Nil(t, newNote)
	})
}

func TestNewNoteWithOctave(t *testing.T) {
	t.Run("TestNewNoteWithOctave: valid note name", func(t *testing.T) {
		noteNames := []Name{
			C,
			CFLAT,
			CFLAT2,
			CSHARP,
			CSHARP2,
			D,
			DFLAT,
			DFLAT2,
			DSHARP,
			DSHARP2,
			E,
			EFLAT,
			EFLAT2,
			ESHARP,
			ESHARP2,
			F,
			FFLAT,
			FFLAT2,
			FSHARP,
			FSHARP2,
			G,
			GFLAT,
			GFLAT2,
			GSHARP,
			GSHARP2,
			A,
			AFLAT,
			AFLAT2,
			ASHARP,
			ASHARP2,
			B,
			BFLAT,
			BFLAT2,
			BSHARP,
			BSHARP2,
		}
		var newNote *Note
		var err error
		for _, noteName := range noteNames {
			// setup: create a Note with a valid name
			newNote, err = NewNoteWithOctave(noteName, octave.NumberDefault)
			require.NoError(t, err)
			assert.NotNil(t, newNote, "expected note from note name %s", noteName)

			// assert that the returned name matches the expected name
			if newNote.name != noteName {
				t.Errorf("Expected note name to be '%s', but got '%s'", noteName, newNote.name)
			}
		}
	})

	t.Run("TestNewNoteWithOctave: invalid note name", func(t *testing.T) {
		// setup: create a Note with invalid name
		expectedName := Name("Hello")
		newNote, err := NewNoteWithOctave(expectedName, octave.NumberDefault)
		// assert that the returned error matches the expected error
		require.ErrorIs(t, err, ErrNoteNameUnknown)
		assert.Nil(t, newNote)
	})

	t.Run("TestNewNoteWithOctave: invalid octave number", func(t *testing.T) {
		// setup: create a Note with invalid octave number
		expectedName := C
		newNote, err := NewNoteWithOctave(expectedName, 15)
		// assert that the returned error matches the expected error
		require.ErrorIs(t, err, octave.ErrOctaveNumberUnknown)
		assert.Nil(t, newNote)
	})
}

func TestMustNewNoteWithOctave(t *testing.T) {
	t.Run("TestMustNewNoteWithOctave: valid note name", func(t *testing.T) {
		noteNames := []Name{
			C,
			CFLAT,
			CFLAT2,
			CSHARP,
			CSHARP2,
			D,
			DFLAT,
			DFLAT2,
			DSHARP,
			DSHARP2,
			E,
			EFLAT,
			EFLAT2,
			ESHARP,
			ESHARP2,
			F,
			FFLAT,
			FFLAT2,
			FSHARP,
			FSHARP2,
			G,
			GFLAT,
			GFLAT2,
			GSHARP,
			GSHARP2,
			A,
			AFLAT,
			AFLAT2,
			ASHARP,
			ASHARP2,
			B,
			BFLAT,
			BFLAT2,
			BSHARP,
			BSHARP2,
		}
		var newNote *Note
		for _, noteName := range noteNames {
			// assert that the function works without panic
			assert.NotPanics(t, func() { newNote = MustNewNoteWithOctave(noteName, octave.NumberDefault) }) //nolint:scopelint

			// assert that the returned name matches the expected name
			if newNote.name != noteName {
				t.Errorf("Expected note name to be '%s', but got '%s'", noteName, newNote.name)
			}
		}
	})

	t.Run("TestMustNewNoteWithOctave: invalid note name", func(t *testing.T) {
		// setup: create a Note with invalid name
		expectedName := Name("Hello")
		// assert that the function works without panic
		assert.Panics(t, func() { _ = MustNewNoteWithOctave(expectedName, octave.NumberDefault) })
	})
}

func TestNoteName(t *testing.T) {
	// setup: create a Note with a known name
	expectedName := C
	note := &Note{name: expectedName}

	// execute the Name() method
	actualName := note.Name()

	// assert that the returned name matches the expected name
	assert.Equal(t, expectedName, actualName, "expected note name: %s, actual: %s", expectedName, actualName)
}

func TestNoteIsEqualByName(t *testing.T) {
	testCases := []struct {
		note1, note2 *Note
		want         bool
	}{
		{
			note1: &Note{name: C},
			note2: &Note{name: C},
			want:  true,
		},
		{
			note1: &Note{name: C},
			note2: &Note{name: D},
			want:  false,
		},
		{
			note1: &Note{name: C},
			note2: nil,
			want:  false,
		},
		{
			note1: nil,
			note2: nil,
			want:  false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note1.IsEqualByName(testCase.note2))
	}
}

func TestNoteIsEqualByOctave(t *testing.T) {
	testCases := []struct {
		note1, note2 *Note
		want         bool
	}{
		{
			note1: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			note2: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			want:  true,
		},
		{
			note1: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			note2: &Note{name: D, octave: octave.MustNewByNumber(octave.Number0)},
			want:  true,
		},
		{
			note1: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			note2: &Note{name: C, octave: octave.MustNewByNumber(octave.Number1)},
			want:  false,
		},
		{
			note1: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			note2: &Note{name: D, octave: octave.MustNewByNumber(octave.Number1)},
			want:  false,
		},
		{
			note1: &Note{name: C, octave: nil},
			note2: &Note{name: C, octave: octave.MustNewByNumber(octave.Number1)},
			want:  false,
		},
		{
			note1: &Note{name: C, octave: nil},
			note2: &Note{name: D, octave: octave.MustNewByNumber(octave.Number1)},
			want:  false,
		},
		{
			note1: &Note{name: C},
			note2: &Note{name: D},
			want:  false,
		},
		{
			note1: &Note{name: C},
			note2: nil,
			want:  false,
		},
		{
			note1: nil,
			note2: nil,
			want:  false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note1.IsEqualByOctave(testCase.note2))
	}
}

func TestNoteIsEqual(t *testing.T) {
	testCases := []struct {
		note1, note2 *Note
		want         bool
	}{
		{
			note1: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			note2: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			want:  true,
		},
		{
			note1: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			note2: &Note{name: D, octave: octave.MustNewByNumber(octave.Number0)},
			want:  false,
		},
		{
			note1: &Note{name: C, octave: octave.MustNewByNumber(octave.Number0)},
			note2: &Note{name: C, octave: octave.MustNewByNumber(octave.Number1)},
			want:  false,
		},
		{
			note1: &Note{name: C, octave: nil},
			note2: &Note{name: C, octave: octave.MustNewByNumber(octave.Number1)},
			want:  false,
		},
		{
			note1: &Note{name: C, octave: nil},
			note2: &Note{name: C, octave: nil},
			want:  false,
		},
		{
			note1: &Note{name: C},
			note2: nil,
			want:  false,
		},
		{
			note1: nil,
			note2: nil,
			want:  false,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note1.IsEqual(testCase.note2))
	}
}

func TestNoteCopy(t *testing.T) {
	note1 := C.MustNewNote().SetOctave(octave.MustNewByNumber(-1)).SetDurationAbs(0).SetDurationRel(duration.NewRelative(duration.NameHalf).SetDots(1))
	note2 := note1.Copy()

	// Test the pointer address is not the same
	if note1 == note2 {
		t.Error("Expected new note instance, but got the same pointer address")
	}

	// Test the notes have the same name
	if note1.Name() != note2.Name() {
		t.Error("Expected new note with same name, but got different names")
	}

	// Test the notes have the same octave
	if note1.Octave() != note2.Octave() {
		t.Error("Expected new note with same octave, but got different octaves")
	}

	// Test nil input returns nil
	if nilNote := (*Note)(nil).Copy(); nilNote != nil {
		t.Error("Expected nil output for nil input, but got non-nil result")
	}
}

func TestNote_AlterUp(t *testing.T) {
	// Test case when note is nil
	var n *Note
	if n.AlterUp() != nil {
		t.Errorf("AlterUp on nil note should return nil")
	}

	// Test case when Name doesn't end with any symbols
	n = &Note{name: B}
	if n.AlterUp().name != BSHARP {
		t.Errorf("AlterUp on note B should result in B#, got: %s", n.name)
	}

	// Test case when Name ends with AccidentalSharp
	n = &Note{name: BSHARP}
	if n.AlterUp().name != BSHARP2 {
		t.Errorf("AlterUp on note B# should result in B##, got: %s", n.name)
	}

	// Test case when Name ends with AccidentalFlat
	n = &Note{name: BFLAT}
	if n.AlterUp().name != B {
		t.Errorf("AlterUp on note Bb should result in B, got: %s", n.name)
	}

	// Test case when Name ends with AccidentalFlat twice
	n = &Note{name: BFLAT2}
	if n.AlterUp().name != BFLAT {
		t.Errorf("AlterUp on note Bbb should result in Bb, got: %s", n.name)
	}
}

func TestNote_AlterDown(t *testing.T) {
	// Test case when note is nil
	var n *Note
	if n.AlterDown() != nil {
		t.Errorf("AlterDown on nil note should return nil")
	}

	// Test case when Name doesn't end with any symbols
	n = &Note{name: B}
	if n.AlterDown().name != BFLAT {
		t.Errorf("AlterDown on note B should result in Bb, got: %s", n.name)
	}

	// Test case when Name ends with AccidentalSharp
	n = &Note{name: BSHARP}
	if n.AlterDown().name != B {
		t.Errorf("AlterDown on note B# should result in B, got: %s", n.name)
	}

	// Test case when Name ends with AccidentalFlat
	n = &Note{name: BFLAT}
	if n.AlterDown().name != BFLAT2 {
		t.Errorf("AlterDown on note Bb should result in Bbb, got: %s", n.name)
	}

	// Test case when Name ends with AccidentalSharp twice
	n = &Note{name: BSHARP2}
	if n.AlterDown().name != BSHARP {
		t.Errorf("AlterDown on note B## should result in Bb# got: %s", n.name)
	}
}

func TestNote_AlterUpBy(t *testing.T) {
	t.Run("Note_AlterUpBy: positive case 1", func(t *testing.T) {
		n := newNote(C)
		alteredNote := n.AlterUpBy(2)
		expectedNote := newNote(CSHARP2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Error("AlterUpBy should alter the note up by the provided value")
		}
	})

	t.Run("Note_AlterUpBy: positive case 2", func(t *testing.T) {
		n := newNote(CFLAT2)
		alteredNote := n.AlterUpBy(4)
		expectedNote := newNote(CSHARP2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Error("AlterUpBy should alter the note up by the provided value")
		}
	})

	t.Run("Note_AlterUpBy: zero times", func(t *testing.T) {
		n := newNote(C)
		if n.AlterUpBy(0) != n {
			t.Error("AlterUpBy should return the same note for 0 alterations")
		}
	})

	t.Run("Note_AlterUpBy: nil note", func(t *testing.T) {
		var n *Note
		if n.AlterUpBy(1) != nil {
			t.Error("AlterUpBy should return nil for nil note")
		}
	})
}

func TestNote_AlterDownBy(t *testing.T) {
	t.Run("Note_AlterDownBy: positive case 1", func(t *testing.T) {
		n := newNote(C)
		alteredNote := n.AlterDownBy(2)
		expectedNote := newNote(CFLAT2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Errorf("AlterDownBy expected: %s, actual: %s", expectedNote.Name(), alteredNote.Name())
		}
	})

	t.Run("Note_AlterDownBy: positive case 2", func(t *testing.T) {
		n := newNote(CSHARP2)
		alteredNote := n.AlterDownBy(4)
		expectedNote := newNote(CFLAT2)
		if alteredNote != n || !alteredNote.IsEqualByName(expectedNote) {
			t.Errorf("AlterDownBy expected: %s, actual: %s", expectedNote.Name(), alteredNote.Name())
		}
	})

	t.Run("Note_AlterDownBy: zero times", func(t *testing.T) {
		n := newNote(C)
		if n.AlterDownBy(0) != n {
			t.Error("AlterDownBy should return the same note for 0 alterations")
		}
	})

	t.Run("Note_AlterDownBy: nil note", func(t *testing.T) {
		var n *Note
		if n.AlterDownBy(1) != nil {
			t.Error("AlterDownBy should return nil for nil note")
		}
	})
}

func TestNote_BaseName(t *testing.T) {
	testCases := []struct {
		note *Note
		want Name
	}{
		{
			note: newNote(CSHARP2),
			want: C,
		},
		{
			note: newNote(CSHARP),
			want: C,
		},
		{
			note: newNote(C),
			want: C,
		},
		{
			note: newNote(CFLAT),
			want: C,
		},
		{
			note: newNote(CFLAT2),
			want: C,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.BaseName(), "expected note name: %s, result: %s", testCase.want, testCase.note.BaseName())
	}
}

func TestNoteSetOctave(t *testing.T) {
	expectedOctave := octave.MustNewByNumber(octave.NumberDefault)

	t.Run("TestNoteSetOctave: setting octave to the note without octave", func(t *testing.T) {
		// create a note without octave
		note1 := newNote(C)
		// set the octave no the note
		note1.SetOctave(expectedOctave)
		// check that they are the same
		assert.True(t, expectedOctave.IsEqual(note1.Octave()))
	})

	t.Run("TestNoteSetOctave: setting octave to the note that already has an octave", func(t *testing.T) {
		// create a note with an octave
		note1 := newNoteWithOctave(C, octave.MustNewByNumber(octave.Number9))
		// set new octave no the note
		note1.SetOctave(expectedOctave)
		// check that they are the same
		assert.True(t, expectedOctave.IsEqual(note1.Octave()))
	})
}

func TestNoteSetDurationRel(t *testing.T) {
	testCases := []struct {
		note *Note
		want *duration.Relative
	}{
		{
			note: newNote(C),
			want: duration.NewRelative(duration.NameEighth),
		},
		{
			note: C.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameDoubleWhole)).
				SetOctave(octave.MustNewByNumber(-1)).SetDurationAbs(0).SetDurationRel(duration.NewRelative(duration.NameHalf).SetDots(1)),

			want: duration.NewRelative(duration.NameEighth),
		},
		{
			note: newNote(C).SetDurationRel(duration.NewRelative(duration.NameDoubleWhole)),
			want: duration.NewRelative(duration.NameEighth),
		},
		{
			note: C.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameDoubleWhole).SetDots(3).SetTuplet(tuplet.New(1, 2))),
			want: duration.NewRelative(duration.NameEighth),
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.SetDurationRel(testCase.want).durationRel)
	}
}

func TestNote_DurationRel(t *testing.T) {
	testCases := []struct {
		note *Note
		want *duration.Relative
	}{
		{
			note: newNote(C),
			want: nil,
		},
		{
			note: C.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameEighth)),
			want: duration.NewRelative(duration.NameEighth),
		},
		{
			note: newNote(C).SetDurationRel(duration.NewRelative(duration.NameEighth)),
			want: duration.NewRelative(duration.NameEighth),
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.DurationRel())
	}
}

func TestNoteSetDurationAbs(t *testing.T) {
	testCases := []struct {
		note *Note
		want time.Duration
	}{
		{
			note: newNote(C),
			want: time.Second,
		},
		{
			note: newNote(D).SetDurationRel(duration.NewRelative(duration.NameDoubleWhole).SetDots(3).SetTuplet(tuplet.New(1, 2))),
			want: time.Second,
		},
		{
			note: newNote(E).SetDurationRel(duration.NewRelative(duration.NameDoubleWhole)),
			want: time.Second,
		},
		{
			note: F.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameDoubleWhole).SetDots(3).SetTuplet(tuplet.New(1, 2))),
			want: time.Second,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.SetDurationAbs(testCase.want).durationAbs)
	}
}

func TestNoteDurationAbs(t *testing.T) {
	testCases := []struct {
		note *Note
		want time.Duration
	}{
		{
			note: newNote(C),
			want: time.Duration(0),
		},
		{
			note: C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameEighth)),
			want: time.Duration(0),
		},
		{
			note: newNote(C).SetDurationRel(duration.NewRelative(duration.NameEighth)),
			want: time.Duration(0),
		},
		{
			note: C.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameDoubleWhole).SetDots(3).SetTuplet(tuplet.New(1, 2))),
			want: time.Duration(0),
		},
		{
			note: C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameEighth)).SetDurationAbs(time.Second),
			want: time.Second,
		},
		{
			note: newNote(C).SetDurationRel(duration.NewRelative(duration.NameEighth)).SetDurationAbs(time.Second),
			want: time.Second,
		},
		{
			note: C.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameDoubleWhole).SetDots(3).SetTuplet(tuplet.New(1, 2))).SetDurationAbs(time.Second),
			want: time.Second,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.DurationAbs(), "absolute duration: %v", testCase.note.DurationAbs())
	}
}

func TestNoteGetAlterationShift(t *testing.T) {
	baseNotes := GetNotesWithAlterations(GetSetFullChromatic(), 0)

	type testCase struct {
		note *Note
		want int8
	}

	var testCases []testCase

	for _, note := range baseNotes {
		for i := uint8(1); i <= 3; i++ {
			testCases = append(testCases, testCase{note.Copy(), 0})
			testCases = append(testCases, testCase{note.Copy().AlterUpBy(i), int8(i)})
			testCases = append(testCases, testCase{note.Copy().AlterDownBy(i), -int8(i)})
		}
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.want, testCase.note.GetAlterationShift(), testCase.note)
	}
}

func TestNoteGetPartOfBarByRel(t *testing.T) {
	testCases := []struct {
		note          *Note
		timeSignature *fraction.Fraction
		want          decimal.Decimal
	}{
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameWhole)),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromInt(1),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameHalf)),
			timeSignature: fraction.New(1, 2),
			want:          decimal.NewFromInt(1),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameWhole)),
			timeSignature: fraction.New(1, 2),
			want:          decimal.NewFromFloat(0.5),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameHalf)),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromInt(2),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameWhole).SetDots(1)),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1.5),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameWhole).SetDots(2)),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1.75),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletDuplet()),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1.5),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletDuplet().AddDot()),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(2.25),
		},
		{
			note:          C.MustMakeNote().SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletTriplet()),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(0.6666666666666667),
		},
		{
			note:          C.MustNewNote().SetDurationRel(duration.NewRelative(duration.NameWhole).SetTupletTriplet().AddDot()),
			timeSignature: fraction.New(1, 1),
			want:          decimal.NewFromFloat(1),
		},
	}

	for _, testCase := range testCases {
		assert.True(t, testCase.want.Equal(testCase.note.GetPartOfBarByRel(testCase.timeSignature)), "expected: %+v, actual: %+v", testCase.want, testCase.note.durationRel.GetPartOfBar(testCase.timeSignature))
	}
}
