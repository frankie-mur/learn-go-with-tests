package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Test without optional param", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaa"

		if repeated != expected {
			t.Errorf("Expected %s but got %s", expected, repeated)
		}
	})
	t.Run("Test with optional param", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("Expected %s but got %s", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
