package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("should calculate the perimeter correctly", func(t *testing.T) {
		rectangle := Rectangle{Width: 10.0, Height: 10.0}

		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}

	t.Run("should calculate the rectangle area correctly", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("should calculate the circle area correctly", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
