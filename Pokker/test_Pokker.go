package Pokker

import (
	"math"
	"math/rand"
)

const (
	m              = 4
	constantPokker = 57.4
)

// Generate a random sequence of bits.
func GenerateRandomSequence(length int) []byte {
	arr := make([]byte, length)
	for i := range arr {
		arr[i] = byte(rand.Intn(2))
	}

	return arr
}

// Divide a sequence of bits into non-overlapping parts of length m.
func divideSequence(bitSequence []byte) [][]byte {
	blocks := make([][]byte, 0)
	for i := 0; i < len(bitSequence); i += m {
		block := bitSequence[i : i+m]
		blocks = append(blocks, block)
	}
	return blocks
}

// Count the number of occurrences of each possible block of length m.
func countBlockOccurrences(blocks [][]byte) map[string]int {
	occurrences := make(map[string]int)
	for _, block := range blocks {
		key := string(block)
		if _, ok := occurrences[key]; !ok {
			occurrences[key] = 0
		}
		occurrences[key]++
	}
	return occurrences
}

// Calculate the X3 parameter for the Pokker test.
func calculateX3(occurrences map[string]int, k int) float64 {
	sum := 0.0
	for _, occurrences := range occurrences {
		sum += math.Pow(float64(occurrences), 2)
	}
	return ((math.Pow(2, float64(m)) / float64(k)) * sum) - float64(k)
}

// Perform the Pokker test.
func PerformPokerTest(sequence []byte) bool {
	blocks := divideSequence(sequence)
	k := len(blocks)

	occurrences := countBlockOccurrences(blocks)

	x3 := calculateX3(occurrences, k)

	return x3 >= 1.03 && x3 <= constantPokker
}
