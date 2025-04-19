package main

import (
	"project-identity/modules/casting"
	"project-identity/modules/square"
)

func Identity(num int) int {
	result := square.Identity(num)
	result = casting.Identity(result)
	return result
}
