package array

// Sum function resuturn summy of array
func Sum(numbers []int) int {
	summy := 0
	for _, number := range numbers {
		summy += number
	}
	return summy
}

// SumAll can return sum of number of slices??
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
