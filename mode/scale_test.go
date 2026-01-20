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
			expected: scale.Scale{note.A.NewNote(), note.B.NewNote(), note.C.NewNote(), note.D.NewNote(), note.E.NewNote(), note.F.NewNote(), note.G.NewNote()},
		},
		{
			name:     "Minor scale descending",
			mode:     MustMakeNewMode(NameNaturalMinor, note.A),
			desc:     true,
			expected: scale.Scale{note.G.NewNote(), note.F.NewNote(), note.E.NewNote(), note.D.NewNote(), note.C.NewNote(), note.B.NewNote(), note.A.NewNote()},
		},
		{
			name:     "Major scale ascending",
			mode:     MustMakeNewMode(NameNaturalMajor, note.C),
			desc:     false,
			expected: scale.Scale{note.C.NewNote(), note.D.NewNote(), note.E.NewNote(), note.F.NewNote(), note.G.NewNote(), note.A.NewNote(), note.B.NewNote()},
		},
		{
			name:     "Major scale descending",
			mode:     MustMakeNewMode(NameNaturalMajor, note.C),
			desc:     true,
			expected: scale.Scale{note.B.NewNote(), note.A.NewNote(), note.G.NewNote(), note.F.NewNote(), note.E.NewNote(), note.D.NewNote(), note.C.NewNote()},
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
