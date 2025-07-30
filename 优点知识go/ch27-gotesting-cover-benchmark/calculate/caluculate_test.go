package calculate

import (
	"fmt"
	"testing"
)

func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log("10 must be even")
	}
	if Even(12) {
		t.Error("77 is not even")
	}
}

func TestADD(t *testing.T) {
	inputs := []struct{ a, b, c int }{
		{1, 2, 3},
		{4, 5, 9},
		{10, 20, 30},
		{3001, 4001, 7002},
	}
	for _, data := range inputs {
		if result := Add(data.a, data.b); result != data.c {
			t.Errorf("Add(%d, %d) expected result=%d,actual result=%d", data.a, data.b, data.c, result)
		}
		fmt.Println()
	}
}

func TestOdd(t *testing.T) {
	if Odd(11) {
		t.Error("11 must be odd")
	}
	if Odd(70) {
		t.Error("70 is odd")
	}
}

func BenchmarkAdd(b *testing.B) {
	input := struct{ a, b, c int }{
		3001, 4001, 7002,
	}
	for i := 0; i < b.N; i++ {
		if result := Add(input.a, input.b); result != input.c {
			b.Errorf("Add(%d, %d) expected result=%d,actual result=%d", input.a, input.b, input.c, result)
		}
	}

}
