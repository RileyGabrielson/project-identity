package trig

import "math"

func Identity(num int) int {
	x, y := splitBits(num)
	// convert to polar coordinates
	r, theta := toPolar(x, y)
	// convert back to cartesian coordinates
	outx, outy := toCartesian(r, theta)
	return joinBits(outx, outy)
}
func toPolar(x int, y int) (float64, float64) {
	r := math.Sqrt(float64(x*x + y*y))
	theta := math.Atan2(float64(y), float64(x))
	return r, theta
}
func toCartesian(r float64, theta float64) (int, int) {
	x := int(math.Round(r * math.Cos(theta)))
	y := int(math.Round(r * math.Sin(theta)))
	return x, y
}
func splitBits(num int) (int, int) {
	bits := num
	if num < 0 {
		bits = -num
	}
	x := bits & 31
	y := bits >> 5
	if num < 0 {
		y += 32
	}
	return x, y
}
func joinBits(x int, y int) int {
	if y >= 32 {
		return -(x + ((y - 32) << 5))
	}
	return x + (y << 5)
}
