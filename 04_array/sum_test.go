package array

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("test adding positive value", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 9})
		want := []int{3, 12}
		// want := "bob"

		// check if two slices are equal or not??
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v", want, got)
		}
	})
}
