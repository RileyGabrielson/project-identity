package main

import (
	"testing"
)

func TestIdentityChain(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive number", 42, 42},
		{"zero", 0, 0},
		{"negative number", -10, -10},
		{"large number", 1000000, 1000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Identity(tc.input)
			if result != tc.expected {
				t.Errorf("Identity chain failed for %d: expected %d, got %d", tc.input, tc.expected, result)
			}
		})
	}
}
