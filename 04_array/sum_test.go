package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of five number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("want '%d' got '%d' give, %v", want, got, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("want %d got %d from, %v", want, got, numbers)
		}
	})
}
