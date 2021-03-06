package clock_test

import (
	"math"
	"testing"
	"time"

	clock "github.com/lucasmls/tdd/clockface"
)

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(321, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b clock.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func Test_SecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clock.Point
	}{
		{time: simpleTime(0, 0, 30), point: clock.Point{0, -1}},
		{time: simpleTime(0, 0, 45), point: clock.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.SecondHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("Wanted %v Point, but got %v", want, got)
			}
		})
	}
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
		got := clock.SecondsInRadians(c.time)

		if want != got {
			t.Fatalf("Wanted %v radians, but got %v", want, got)
		}
	}
}

func Test_MinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clock.Point
	}{
		{simpleTime(0, 30, 0), clock.Point{0, -1}},
		{simpleTime(0, 45, 0), clock.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.MinuteHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("Wanted %v Point, but got %v", want, got)
			}
		})
	}
}

func Test_MinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.MinutesInRadians(c.time)
			want := c.angle

			if got != want {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestHourInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.HoursInRadians(c.time)
			want := c.angle

			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("Wanted %v radians, but got %v", want, got)
			}
		})
	}
}

func Test_HourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clock.Point
	}{
		{simpleTime(6, 0, 0), clock.Point{0, -1}},
		{simpleTime(21, 0, 0), clock.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clock.HourHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("Wanted %v Point, but got %v", want, got)
			}
		})
	}
}
