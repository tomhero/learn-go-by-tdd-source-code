package array

// Sum function resuturn summy of array
func Sum(numbers []int) int {
	summy := 0
	for _, number := range numbers {
		summy += number
	}
	return summy
}
