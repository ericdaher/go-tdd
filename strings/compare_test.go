package strings

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Run("when the two strings are the same", func(t *testing.T) {
		expected := true
		got := Compare("a", "a")

		if expected != got {
			t.Errorf("expected %t, got %t", expected, got)
		}
	})

	t.Run("when the two strings are different", func(t *testing.T) {
		expected := false
		got := Compare("a", "b")

		if expected != got {
			t.Errorf("expected %t, got %t", expected, got)
		}
	})
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compare("a", "b")
	}
}

func ExampleCompare() {
	result := Compare("a", "a")
	fmt.Println(result)
	// Output: true
}
