package interval

import (
	"testing"

	"github.com/go-muse/muse/degree"
)

func TestAddFilterByDegreeCharacteristicName(t *testing.T) {
	tests := []struct {
		name     string
		input    []degree.CharacteristicName
		check    degree.CharacteristicName
		expected bool
	}{
		{
			name:     "Single characteristic",
			input:    []degree.CharacteristicName{degree.CharacteristicMajor},
			check:    degree.CharacteristicMajor,
			expected: true,
		},
		{
			name:     "Multiple characteristics",
			input:    []degree.CharacteristicName{degree.CharacteristicMajor, degree.CharacteristicMinor},
			check:    degree.CharacteristicMinor,
			expected: true,
		},
		{
			name:     "Characteristic not added",
			input:    []degree.CharacteristicName{degree.CharacteristicMajor},
			check:    degree.CharacteristicMinor,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := NewIntervalGetOptions(AddFilterByDegreeCharacteristicName(tt.input))
			if got := opts.HasFilterByDegreeCharacteristicName(tt.check); got != tt.expected {
				t.Errorf("HasFilterByDegreeCharacteristicName() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAddFilterByAbsoluteModalPosition(t *testing.T) {
	tests := []struct {
		name     string
		input    []degree.ModalPositionName
		check    degree.ModalPositionName
		expected bool
	}{
		{
			name:     "Single modal position",
			input:    []degree.ModalPositionName{"Position1"},
			check:    "Position1",
			expected: true,
		},
		{
			name:     "Multiple modal positions",
			input:    []degree.ModalPositionName{"Position1", "Position2"},
			check:    "Position2",
			expected: true,
		},
		{
			name:     "Modal position not added",
			input:    []degree.ModalPositionName{"Position1"},
			check:    "Position3",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := NewIntervalGetOptions(AddFilterByAbsoluteModalPosition(tt.input))
			if got := opts.HasFilterByAbsoluteModalPosition(tt.check); got != tt.expected {
				t.Errorf("HasFilterByAbsoluteModalPosition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAddFilterBySonance(t *testing.T) {
	tests := []struct {
		name     string
		input    []Sonance
		check    Sonance
		expected bool
	}{
		{
			name:     "Single sonance",
			input:    []Sonance{SonanceMinorSecond},
			check:    SonanceMinorSecond,
			expected: true,
		},
		{
			name:     "Multiple sonances",
			input:    []Sonance{SonanceMinorSecond, SonanceMajorSixth},
			check:    SonanceMajorSixth,
			expected: true,
		},
		{
			name:     "Sonance not added",
			input:    []Sonance{SonanceMinorSecond},
			check:    SonanceMajorThird,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := NewIntervalGetOptions(AddFilterBySonance(tt.input))
			if got := opts.HasFilterBySonance(tt.check); got != tt.expected {
				t.Errorf("HasFilterBySonance() = %v, want %v", got, tt.expected)
			}
		})
	}
}
