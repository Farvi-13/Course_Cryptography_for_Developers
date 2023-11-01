package main

import (
	"Lab_4/Pokker"
	"Lab_4/maximum-length-series"
	"Lab_4/monobit"
	"Lab_4/series-lengths"
	"fmt"
)

func main() {
	// Generate a random bit sequence.
	bitSequence := monobit.GenerateRandomSequence(20000)
	trueOrFalseMonobitTest := monobit.MonobitTest(bitSequence)
	fmt.Println("Monobit Test:", trueOrFalseMonobitTest)

	trueOrFalseMaximumLengthSeriesTest := maximum_length_series.MaxSeriesLengthTest(bitSequence)
	fmt.Println("Maximum Length Series Test:", trueOrFalseMaximumLengthSeriesTest)

	trueOrFalsePokkerTest := Pokker.PerformPokerTest(bitSequence)
	fmt.Println("Pokker Test:", trueOrFalsePokkerTest)

	trueOrFalseSeriesLengthsTest := series_lengths.SeriesLengthTest(bitSequence)
	fmt.Println("Series Lengths Test:", trueOrFalseSeriesLengthsTest)

	if trueOrFalseMonobitTest == true && trueOrFalseMaximumLengthSeriesTest == true &&
		trueOrFalsePokkerTest == true && trueOrFalseSeriesLengthsTest == true {

		fmt.Println("20,000 bits are random enough!")
	} else {
		fmt.Println("The bit sequence is rejected.")
	}
}
