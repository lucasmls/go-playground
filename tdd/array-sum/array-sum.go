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

// SumAllTails ...
func SumAllTails(listsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range listsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
