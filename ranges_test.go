package main

import (
	"testing"
)

func Benchmark_get_special_chars_rt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get_special_char_rt()
	}
}
