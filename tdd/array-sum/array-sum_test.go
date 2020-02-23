package operations

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum all of the five array items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	})

	t.Run("should sum all of the ten array items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := Sum(numbers)
		want := 55

		if got != want {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("expect to sum all of the items of the array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("expect to sum only the tail of the slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v", got, want)
		}
	})

	t.Run("expect to safely sum the slices tail", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v", got, want)
		}
	})
}
