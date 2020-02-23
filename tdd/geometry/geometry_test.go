package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("should calculate the perimeter correctly", func(t *testing.T) {
		rectangle := Rectangle{width: 10.0, height: 10.0}

		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("should calculate the area correctly", func(t *testing.T) {
		rectangle := Rectangle{width: 12.0, height: 6.0}

		got := Area(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	})
}
