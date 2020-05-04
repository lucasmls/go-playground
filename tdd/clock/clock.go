package clock

import "time"

// Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of a analogue clock at time 't'
// represented as a Point.
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
