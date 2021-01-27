package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) { //基准测试
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() { //实例
	repeated := Repeat("a")
	fmt.Println(repeated)
	// output: aaaaa
}
