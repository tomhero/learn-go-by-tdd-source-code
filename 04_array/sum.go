package array

// Sum function resuturn summy of array
func Sum(numbers []int) int {
	summy := 0
	for _, number := range numbers {
		summy += number
	}
	return summy
}

// SumAllTails can return sum of number of slices??
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
