package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counting three times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})
}
