package note

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/octave"
)

func TestNoteFrequency(t *testing.T) {
	tests := []struct {
		name     string
		note     *Note
		expected float64
	}{
		{
			name:     "С-1",
			note:     &Note{name: C, octave: octave.MustNewByNumber(-1)},
			expected: 8.16,
		},
		{
			name:     "С0",
			note:     &Note{name: C, octave: octave.MustNewByNumber(0)},
			expected: 16.35,
		},
		{
			name:     "A4",
			note:     &Note{name: A, octave: octave.MustNewByNumber(4)},
			expected: 440.0,
		},
		{
			name:     "C4",
			note:     &Note{name: C, octave: octave.MustNewByNumber(4)},
			expected: 261.63,
		},
		{
			name:     "G#4",
			note:     &Note{name: GSHARP, octave: octave.MustNewByNumber(4)},
			expected: 415.30,
		},
		{
			name:     "Bb3",
			note:     &Note{name: BFLAT, octave: octave.MustNewByNumber(3)},
			expected: 233.08,
		},
		{
			name:     "F#5",
			note:     &Note{name: FSHARP, octave: octave.MustNewByNumber(5)},
			expected: 739.99,
		},
		{
			name:     "B8",
			note:     &Note{name: B, octave: octave.MustNewByNumber(8)},
			expected: 7902.08,
		},
		{
			name:     "B9",
			note:     &Note{name: B, octave: octave.MustNewByNumber(9)},
			expected: 15804.26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.note.FrequencyBy440()
			assert.InEpsilon(t, tt.expected, result, 0.01) // Allowing a small margin of error
		})
	}
}
