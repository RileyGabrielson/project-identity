package main

import (
	"project-identity/modules/casting"
	"project-identity/modules/square"
	"project-identity/modules/subcipher"
)

func Identity(num int) int {
	result := square.Identity(num)
	result = casting.Identity(result)
	result = subcipher.Identity(result)
	return result
}
