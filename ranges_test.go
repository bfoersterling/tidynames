package main

import (
	"testing"
)

func Benchmark_get_whitelist_rt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get_whitelist_rt()
	}
}
