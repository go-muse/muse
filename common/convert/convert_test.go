package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint8Int8(t *testing.T) {
	tests := []struct {
		name     string
		a        uint8
		b        int8
		expected uint8
		panics   bool
	}{
		{"Simple addition", 5, 3, 8, false},
		{"Addition with negative", 10, -3, 7, false},
		{"Max uint8", 255, 0, 255, false},
		{"Min uint8", 0, 0, 0, false},
		{"Overflow", 255, 1, 0, true},
		{"Underflow", 0, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { AddUint8Int8(tt.a, tt.b) })
			} else {
				result := AddUint8Int8(tt.a, tt.b)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestSubUint8Uint8(t *testing.T) {
	tests := []struct {
		name     string
		a        uint8
		b        uint8
		expected int8
		panics   bool
	}{
		{"Simple subtraction", 10, 5, 5, false},
		{"Zero result", 5, 5, 0, false},
		{"Negative result", 5, 10, -5, false},
		{"Max positive result", 255, 128, 127, false},
		{"Max negative result", 0, 128, -128, false},
		{"Overflow positive", 255, 127, 0, true},
		{"Overflow negative", 0, 129, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { SubUint8Uint8(tt.a, tt.b) })
			} else {
				result := SubUint8Uint8(tt.a, tt.b)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
