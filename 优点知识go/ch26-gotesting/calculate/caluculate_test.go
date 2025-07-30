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
		{1, 2, 4},
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
