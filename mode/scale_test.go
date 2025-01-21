package mode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-muse/muse/note"
	"github.com/go-muse/muse/scale"
)

func TestGenerateScale(t *testing.T) {
	tests := []struct {
		name     string
		mode     *Mode
		desc     bool
		expected scale.Scale
	}{
		{
			name:     "Nil mode",
			mode:     nil,
			desc:     false,
			expected: nil,
		},
		{
			name:     "Nil degree",
			mode:     &Mode{degree: nil},
			desc:     false,
			expected: nil,
		},
		{
			name:     "Minor scale ascending",
			mode:     MustMakeNewMode(NameNaturalMinor, note.A),
			desc:     false,
			expected: scale.Scale{note.A.MustMakeNote(), note.B.MustMakeNote(), note.C.MustMakeNote(), note.D.MustMakeNote(), note.E.MustMakeNote(), note.F.MustMakeNote(), note.G.MustMakeNote()},
		},
		{
			name:     "Minor scale descending",
			mode:     MustMakeNewMode(NameNaturalMinor, note.A),
			desc:     true,
			expected: scale.Scale{note.G.MustMakeNote(), note.F.MustMakeNote(), note.E.MustMakeNote(), note.D.MustMakeNote(), note.C.MustMakeNote(), note.B.MustMakeNote(), note.A.MustMakeNote()},
		},
		{
			name:     "Major scale ascending",
			mode:     MustMakeNewMode(NameNaturalMajor, note.C),
			desc:     false,
			expected: scale.Scale{note.C.MustMakeNote(), note.D.MustMakeNote(), note.E.MustMakeNote(), note.F.MustMakeNote(), note.G.MustMakeNote(), note.A.MustMakeNote(), note.B.MustMakeNote()},
		},
		{
			name:     "Major scale descending",
			mode:     MustMakeNewMode(NameNaturalMajor, note.C),
			desc:     true,
			expected: scale.Scale{note.B.MustMakeNote(), note.A.MustMakeNote(), note.G.MustMakeNote(), note.F.MustMakeNote(), note.E.MustMakeNote(), note.D.MustMakeNote(), note.C.MustMakeNote()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.mode.GenerateScale(tt.desc)
			assert.Equal(t, len(tt.expected), len(result))
			for i := range tt.expected {
				assert.Equal(t, tt.expected[i], result[i])
			}
		})
	}
}
