package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNotesWithAlterations(t *testing.T) {
	tests := []struct {
		name        string
		notes       Notes
		alterations uint8
		expected    Notes
	}{
		{
			name:        "No alterations",
			notes:       MustNewNotesFromNoteNames(C, D, E, F, G, A, B),
			alterations: 0,
			expected:    Notes{},
		},
		{
			name:        "One alteration",
			notes:       MustNewNotesFromNoteNames(C, D, E, F, G, A, B),
			alterations: 1,
			expected:    MustNewNotesFromNoteNames(CFLAT, CSHARP, DFLAT, DSHARP, EFLAT, ESHARP, FFLAT, FSHARP, GFLAT, GSHARP, AFLAT, ASHARP, BFLAT, BSHARP),
		},
		{
			name:        "Two alterations",
			notes:       MustNewNotesFromNoteNames(C, D, E, F, G, A, B),
			alterations: 2,
			expected:    MustNewNotesFromNoteNames(CFLAT, CFLAT2, CSHARP, CSHARP2, DFLAT, DFLAT2, DSHARP, DSHARP2, EFLAT, EFLAT2, ESHARP, ESHARP2, FFLAT, FFLAT2, FSHARP, FSHARP2, GFLAT, GFLAT2, GSHARP, GSHARP2, AFLAT, AFLAT2, ASHARP, ASHARP2, BFLAT, BFLAT2, BSHARP, BSHARP2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetNotesWithAlterations(tt.notes, tt.alterations)
			assert.ElementsMatch(t, tt.expected, result, "GetNotesWithAlterations returned unexpected result")
		})
	}
}
