package main

import (
	"project-identity/modules/bigmap"
	"project-identity/modules/casting"
	"project-identity/modules/primefactor"
	"project-identity/modules/random"
	"project-identity/modules/square"
	"project-identity/modules/subcipher"
	"project-identity/modules/ternary"
	"project-identity/modules/trig"
)

func Identity(num int) int {
	result := square.Identity(num)
	result = casting.Identity(result)
	result = subcipher.Identity(result)
	result = bigmap.Identity(result)
	result = primefactor.Identity(result)
	result = ternary.Identity(result)
	result = random.Identity(result)
	result = trig.Identity(result)
	return result
}
