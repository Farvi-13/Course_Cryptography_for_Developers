package series_lengths

import (
	"math/rand"
)

// Generate a random sequence of bits.
func GenerateRandomSequence(length int) []byte {
	arr := make([]byte, length)
	for i := range arr {
		arr[i] = byte(rand.Intn(2))
	}

	return arr
}

// seriesLengthIntervals is a map of series lengths to the corresponding appropriate intervals.
var seriesLengthIntervals = map[int][]int{
	1: {2267, 2733},
	2: {1079, 1421},
	3: {502, 748},
	4: {223, 402},
	5: {90, 223},
	6: {90, 223},
}

// CountSeries counts the number of series of a given length in a bit sequence.
func countSeries(bitSequence []byte, length int, targetBit byte) int {
	count := 0
	seriesLength := 0

	for i := 0; i < len(bitSequence); i++ {
		if bitSequence[i] == targetBit {
			seriesLength++
		} else {
			if seriesLength == length {
				count++
			}
			seriesLength = 0
		}
	}

	// Checking the latest episode.
	if seriesLength == length {
		count++
	}

	return count
}

// SeriesLengthTest performs the series length test for randomness on a bit sequence.
func SeriesLengthTest(bitSequence []byte) bool {
	// Count the number of series of each length.
	seriesCountsOnes := make(map[int]int)
	seriesCountsTwos := make(map[int]int)

	for length := 1; length <= 6; length++ {
		seriesCountsOnes[length] = countSeries(bitSequence, length, 1)
		seriesCountsTwos[length] = countSeries(bitSequence, length, 0)
	}

	// Check if the series counts fall within the appropriate intervals.
	for length, interval := range seriesLengthIntervals {
		if (seriesCountsOnes[length] < interval[0]) || (seriesCountsOnes[length] > interval[1]) {
			return false
		}
		if (seriesCountsTwos[length] < interval[0]) || (seriesCountsTwos[length] > interval[1]) {
			return false
		}
	}

	return true
}
