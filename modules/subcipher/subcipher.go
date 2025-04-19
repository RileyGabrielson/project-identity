package subcipher

import (
	"strings"
)

var substitutionWords = []string{
	"Alpha", "Bravo", "Charlie", "Delta", "Echo",
	"Foxtrot", "Golf", "Hotel", "India", "Juliet",
	"Kilo", "Lima", "Mike", "November", "Oscar",
	"Papa", "Quebec", "Romeo", "Sierra", "Tango",
	"Uniform", "Victor", "Whiskey", "Xray", "Yankee",
	"Zulu", "Avalanche", "Blackout", "Cipher", "Darkness",
	"Eclipse", "Falcon", "Ghost", "Havoc", "Inferno",
	"Jaguar", "Kraken", "Lightning", "Midnight", "Nova",
	"Omega", "Phantom", "Quicksilver", "Raven", "Shadow",
	"Thunder", "Umbra", "Viper", "Wolf", "Zenith",
}

func Identity(num int) int {
	// Convert negative numbers to positive for cipher generation
	steps := num
	if num < 0 {
		steps = -num
	}

	// Generate the cipher deterministically
	var cipher strings.Builder
	sum := 0
	wordIndex := 1 // Start with the first word

	for sum < steps {
		// If adding this word would exceed our target, use a smaller word
		if sum + wordIndex > steps {
			wordIndex = steps - sum
		}

		// Add the word to our cipher
		cipher.WriteString(substitutionWords[wordIndex-1])
		cipher.WriteString(" ")
		sum += wordIndex

		// Move to next word, wrapping around if needed
		wordIndex = (wordIndex % 50) + 1
	}

	// Now decipher the cipher to get our original number back
	words := strings.Fields(cipher.String())
	result := 0
	for _, word := range words {
		// Find the index of the word in our list
		for i, substitutionWord := range substitutionWords {
			if word == substitutionWord {
				result += i + 1
				break
			}
		}
	}

	// Restore original sign
	if num < 0 {
		result = -result
	}

	return result
} 