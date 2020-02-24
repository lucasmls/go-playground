package geometry

import "math"

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle ...
type Circle struct {
	Radius float64
}

// Triangle ...
type Triangle struct {
	Base   float64
	Height float64
}

// Shape ...
type Shape interface {
	Area() float64
}

// Perimeter ...
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

// Area ...
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area ...
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area ...
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
