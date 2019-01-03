package test

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	var a int64
	for i := 0; i < b.N; i++ {
		a++
	}
}

func BenchmarkMul(b *testing.B) {
	var a int64
	for i := 0; i < b.N; i++ {
		a = (a * int64(i))
	}
}
