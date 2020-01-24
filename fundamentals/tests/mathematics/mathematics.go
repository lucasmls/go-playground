package mathematics

import (
	"fmt"
	"strconv"
)

// Average ...
func Average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	average := total / float64(len(numbers))
	roundedAverage, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)

	return roundedAverage
}
