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
		{"min", -1000, -1000},
		{"max", 1000, 1000},
		{"large prime", 677, 677},
		{"perfect square", 16, 16},
		{"product of primes", 210, 210},
		{"power of two", 128, 128},
		{"palindrome", 121, 121},
		{"negative prime", -23, -23},
		{"negative perfect square", -49, -49},
		{"consecutive digits", 123, 123},
		{"repeated digits", 333, 333},
		{"edge case near max", 999, 999},
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
