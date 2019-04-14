package array

// Sum function resuturn summy of array
func Sum(numbers [5]int) int {
	summy := 0
	for i := 0; i < 5; i++ {
		summy += numbers[i]
	}
	return summy
}
