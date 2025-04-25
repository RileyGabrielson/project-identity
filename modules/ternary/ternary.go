package ternary

import "math"

func Identity(num int) int {
	numUint := uint(math.Abs(float64(num)))
	ch := toTernary(numUint)
	out := fromTernary(ch)
	return int(out) * boolToInt(math.Signbit(float64(num)))
}
func boolToInt(b bool) int {
	if b {
		return -1
	}
	return 1
}

type ternary uint

const (
	Zero ternary = iota
	One
	Two
)

func toTernary(num uint) chan ternary {
	ch := make(chan ternary)
	go func() {
		defer close(ch)
		for num > 0 {
			ch <- ternary(num % 3)
			num /= 3
		}
	}()
	return ch
}
func fromTernary(ch chan ternary) uint {
	num := uint(0)
	multiplier := uint(1)
	for v := range ch {
		num += uint(v) * multiplier
		multiplier *= 3
	}
	return num
}
