package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("should calculate the perimeter correctly", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("should calculate the area correctly", func(t *testing.T) {
		got := Area(12.0, 6.0)
		want := 72.0

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	})
}
