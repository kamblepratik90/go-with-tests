package main

import (
	"fmt"
	"testing"
)

func Sum(a []int) (sum int) {
	for _, value := range a {
		sum += value
	}
	return
}

func ExampleSum() {
	result := Sum([]int{1, 2, 3})
	fmt.Println(result)
	// Output: 6
}

func TestSum(t *testing.T) {
	// result := Sum([]int{1, 2, 3})
	// expected := 6

	// if result != expected {
	// 	t.Errorf("result %d and expected %d ", result, expected)
	// }

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
