package maximum_length_series

import (
	"math/rand"
)

const maxSeriesLength = 36

func MaxSeriesLengthTest(bitSequence []byte) bool {
	// Count the longest series of zeros and ones.
	longestZeroSeries := 0
	longestOneSeries := 0
	currentZeroSeries := 0
	currentOneSeries := 0
	for _, bit := range bitSequence {
		if bit == 0 {
			currentZeroSeries++
			currentOneSeries = 0
		} else {
			currentOneSeries++
			currentZeroSeries = 0
		}

		if currentZeroSeries > longestZeroSeries {
			longestZeroSeries = currentZeroSeries
		}

		if currentOneSeries > longestOneSeries {
			longestOneSeries = currentOneSeries
		}
	}

	// Check if the longest series is longer than the maximum allowed length.
	if longestZeroSeries > maxSeriesLength || longestOneSeries > maxSeriesLength {
		return false
	}

	return true
}

// Generate a random sequence of bits.
func GenerateRandomSequence(length int) []byte {
	arr := make([]byte, length)
	for i := range arr {
		arr[i] = byte(rand.Intn(2))
	}

	return arr
}
