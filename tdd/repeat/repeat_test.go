package repeat

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("expect to receive the character concatenated five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

	t.Run("expect to receive zezezezeze", func(t *testing.T) {
		repeated := Repeat("ze")
		expected := "zezezezeze"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("la")
	}
}
