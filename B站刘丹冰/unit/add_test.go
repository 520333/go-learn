package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 方法必须以Test开头
func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect %d,actual %d", 3, re)
	}
}

// 执行命令跳过 go test -short
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下跳过")
	}
	fmt.Println("yes")
	re := add(1, 5)
	if re != 6 {
		t.Errorf("expect %d,actual %d", 6, re)
	}
}

// 基于表格性能测试
func TestAdd3(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{12, 12, 24},
		{-9, 8, -1},
		{0, 0, 0},
	}
	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect:%d,actual:%d", value.out, re)
		}
	}
}

// 性能测试 必须以Benchmark开头 执行命令：test -bench=".*"
func BenchmarkAdd(bb *testing.B) {
	var a, b, c int = 123, 256, 789
	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("%d + %d,except:%d, actual:%d", a, b, c, actual)
		}
	}
}

const numbers = 10000

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
