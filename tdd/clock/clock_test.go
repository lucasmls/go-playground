package clock

import (
	"math"
	"testing"
	"time"
)

func Test_SecondHand(t *testing.T) {
	// t.Run("second hand at midnight", func(t *testing.T) {
	// 	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	// 	want := clock.Point{X: 150, Y: 150 - 90}
	// 	got := clock.SecondHand(tm)

	// 	if got != want {
	// 		t.Errorf("Got %v, wanted %v", got, want)
	// 	}
	// })

	// t.Run("second hand at 30 seconds", func(t *testing.T) {
	// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	// 	want := clock.Point{X: 150, Y: 150 + 90}
	// 	got := clock.SecondHand(tm)

	// 	if got != want {
	// 		t.Errorf("Got %v, wanted %v", got, want)
	// 	}
	// })
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(321, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func Test_SecondsInRadians(t *testing.T) {

	cases := []struct {
		time  time.Time
		angle float64
	}{
		{time: simpleTime(0, 0, 0), angle: 0},
		{time: simpleTime(0, 0, 30), angle: math.Pi},
		{time: simpleTime(0, 0, 45), angle: (math.Pi / 2) * 3},
		{time: simpleTime(0, 0, 7), angle: (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		want := c.angle
		got := secondsInRadians(c.time)

		if want != got {
			t.Fatalf("Wanted %v radians, but got %v", want, got)
		}
	}
}
