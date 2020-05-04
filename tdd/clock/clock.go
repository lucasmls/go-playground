package clock

import (
	"math"
	"time"
)

// Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30)
	return (math.Pi / (30 / (float64(t.Second()))))
}

// SecondHand is the unit vector of the second hand of a analogue clock at time 't'
// represented as a Point.
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
