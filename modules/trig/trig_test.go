package trig

import "testing"

func TestIdentity(t *testing.T) {
	// Ensure the rounding is correct
	for i := -1000; i <= 1000; i++ {
		result := Identity(i)
		if result != i {
			t.Errorf("Identity(%d) = %d", i, result)
		}
	}
}
