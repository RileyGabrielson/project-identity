package square

import "math"

func Identity(num int) int {
	// Store the original sign
	sign := 1
	if num < 0 {
		sign = -1
	}
	
	// Convert to float64 for math operations
	floatNum := float64(num)
	
	// Square the number
	squared := floatNum * floatNum
	
	// Take the square root
	result := math.Sqrt(squared)
	
	// Convert back to integer and restore the sign
	return int(result) * sign
} 