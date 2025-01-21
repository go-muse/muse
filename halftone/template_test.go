package halftone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTemplate(t *testing.T) {
	tests := []struct {
		name     string
		input    []HalfTones
		expected Template
	}{
		{
			name:     "Empty template",
			input:    []HalfTones{},
			expected: Template{},
		},
		{
			name:     "Single value template",
			input:    []HalfTones{2},
			expected: Template{2},
		},
		{
			name:     "Multiple values template",
			input:    []HalfTones{2, 2, 1, 2, 2, 2, 1},
			expected: Template{2, 2, 1, 2, 2, 2, 1},
		},
		{
			name:     "Template with zero values",
			input:    []HalfTones{0, 1, 0, 2},
			expected: Template{0, 1, 0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewTemplate(tt.input...)
			assert.Equal(t, tt.expected, result)

			// Check that we have a copy of the input slice.
			if len(result) > 0 {
				result[0] = 99
				assert.NotEqual(t, result[0], tt.input[0], "NewTemplate should return a copy, not the original slice")
			}
		})
	}
}
