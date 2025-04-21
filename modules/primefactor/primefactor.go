package primefactor

import (
	"math"
)

const maxIter = 50

func Identity(num int) int {
	x := num
	if num < 0 {
		x = -1 * num
	}

	factors := GetFactors(x)
	recombined := combineFactors(factors)

	if num < 0 {
		return -1 * recombined
	}
	return recombined
}

func combineFactors(factors []int) int {
	result := 1
	for i := 0; i < len(factors); i++ {
		result = result * factors[i]
	}

	return result
}

func GetFactors(num int) []int {
	factors := []int{num}

	i := 0

	for {
		for j := 0; j < maxIter; j++ {
			var currentFactor = factors[i]
			var newFactor = pollardRho(currentFactor, 2)
			if newFactor == currentFactor {
				break
			} else {
				factors = append(factors, newFactor)
				factors[i] = currentFactor / newFactor
			}
		}

		i += 1
		if i == len(factors) {
			break
		}
	}

	return factors
}

func pollardRho(n int, starting int) int {

	if n < 2 {
		return n
	}

	// With small enough numbers, the algorithm doesn't have enough space to find a factor
	// So, we cover some of the early cases
	if n == 4 {
		return 2
	} else if n == 6 {
		return 3
	} else if n == 8 {
		return 4
	} else if n == 9 {
		return 3
	}

	x := starting
	y := x
	d := 1

	for d == 1 {
		x = g(x, n)
		y = g(g(y, n), n)
		d = greatestCommonDivisor(int(math.Abs(float64(x-y))), n)
	}

	if d == n && starting < n {
		return pollardRho(n, starting+1)
	}

	return d
}

func g(x int, n int) int {
	return (x*x - 1) % n
}

func greatestCommonDivisor(x, y int) int {
	a := x
	b := y

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
