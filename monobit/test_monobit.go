package monobit

import (
	"math/rand"
)

const (
	minOnes = 9654
	maxOnes = 10346
)

func MonobitTest(bitSequence []byte) bool {
	// Count the number of ones in the bit sequence.
	numOnes := 0
	for _, bit := range bitSequence {
		if bit == 1 {
			numOnes++
		}
	}

	// Check if the number of ones is within the acceptable range.
	return minOnes <= numOnes && numOnes <= maxOnes
}

// Generate a random sequence of bits.
func GenerateRandomSequence(length int) []byte {
	arr := make([]byte, length)
	for i := range arr {
		arr[i] = byte(rand.Intn(2))
	}

	return arr
}
