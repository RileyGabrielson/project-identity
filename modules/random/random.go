package random

import (
	"math/rand"
)

// With 2001 possibilities is will be 2000/2001 chance to fail each roll
// for a 99% chance of success it will take n rolls where n is:
// n = log(1-0.99) / log(2000/2001) ~= 9212 rolls
// 1-(2000/2001)**9212 ~= 0.99
// A computer can run 10s of millions of rolls per second
// This function could return result in anywhere between 1 and infinity rolls
func Identity(num int) int {
	x := 42
	for x != num {
		x = rand.Intn(2001) - 1000
	}
	return x
}
