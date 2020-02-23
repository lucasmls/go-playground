package geometry

// Rectangle ...
type Rectangle struct {
	width  float64
	height float64
}

// Perimeter ...
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

// Area ...
func Area(rect Rectangle) float64 {
	return rect.width * rect.height
}
