package mathematics

import "testing"

const defaultError = "Expected value was %v, but the calculated result id %v"

func TestAverage(t *testing.T) {
	expectedResult := 5.0
	result := Average(5.0, 5.0, 5.0)

	if result != expectedResult {
		t.Errorf(defaultError, expectedResult, result)
	}
}
