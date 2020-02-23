package operations

// Sum ...
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll ...
func SumAll(listsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range listsToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
