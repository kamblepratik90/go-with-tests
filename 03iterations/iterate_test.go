package iterations

import (
	"fmt"
	"testing"
)

const repeatCount = 5

func Repeat(s string, repeatCount int) (msg string) {
	// for i := 0; i < 5; i++ {
	for i := 0; i < repeatCount; i++ {
		// msg = msg + s
		msg += s
	}
	return
}

func ExampleRepeat() {
	repeated := Repeat("1", repeatCount)
	fmt.Println(repeated)
	// Output: 11111
}

func BenchmarkIterations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
}

func TestIterations(t *testing.T) {
	repeated := Repeat("1", repeatCount)
	expected := "11111"

	if repeated != expected {
		t.Errorf("repeated %q and expected %q ", repeated, expected)
	}
}
