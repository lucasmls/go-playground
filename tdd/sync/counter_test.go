package main

import "testing"

func assertCounterValue(t *testing.T, got Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
func TestCounter(t *testing.T) {

	t.Run("incrementing the counting three times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounterValue(t, counter, 3)
	})
}
