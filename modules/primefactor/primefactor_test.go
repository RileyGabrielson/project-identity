package primefactor

import (
	"slices"
	"testing"
)

func TestFactor(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "2",
			input:    2,
			expected: []int{2},
		},
		{
			name:     "16",
			input:    16,
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "30",
			input:    30,
			expected: []int{2, 3, 5},
		},
		{
			name:     "210",
			input:    210,
			expected: []int{2, 3, 5, 7},
		},
		{
			name:     "398",
			input:    398,
			expected: []int{2, 199},
		},
		{
			name:     "-398",
			input:    -398,
			expected: []int{-2, 199},
		},
		{
			name:     "677",
			input:    677,
			expected: []int{677},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFactors(tt.input)
			if !compareIntSlices(got, tt.expected) {
				t.Errorf("GetFactors(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
func compareIntSlices(a []int, b []int) bool {
	slices.Sort(a)
	slices.Sort(b)

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
