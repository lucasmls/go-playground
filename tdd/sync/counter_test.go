package main

import (
	"sync"
	"testing"
)

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

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Increment()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounterValue(t, counter, wantedCount)
	})
}
