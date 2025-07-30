package main

import "testing"

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus()
	}
}
func BenchmarkStringFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFmt()
	}
}
