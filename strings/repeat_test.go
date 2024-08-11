package strings

import (
	"fmt"
	"testing"
)

func TestRepea(t *testing.T) {
	expected := "aaaaa"
	got := Repeat("a", 5)

	if expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
