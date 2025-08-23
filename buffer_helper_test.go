package main

import (
	"bytes"
	"testing"
)

func Test_peek_byte_1(t *testing.T) {
	// 1
	test_buffer := bytes.NewBuffer([]byte("foo"))

	if peek_byte_1(test_buffer) != 'f' {
		t.Fatalf("First byte of test_buffer should be f!\n")
	}

	if peek_byte_1(test_buffer) != 'f' {
		t.Fatalf("First byte of test_buffer should be f!\n")
	}

	// 2 - no bytes left

	test_buffer.ReadByte()
	test_buffer.ReadByte()
	test_buffer.ReadByte()
	test_buffer.ReadByte()
	test_buffer.ReadByte()

	if peek_byte_1(test_buffer) != byte(0) {
		t.Fatalf("First byte of test_buffer should be empty!\n")
	}
}

func Benchmark_peek_byte_1(b *testing.B) {
	test_buffer := bytes.NewBuffer([]byte("foobar.txt"))

	for i := 0; i < b.N; i++ {
		peek_byte_1(test_buffer)
	}
}

func Test_peek_byte_2(t *testing.T) {
	// 1
	test_buffer := bytes.NewBuffer([]byte("foo"))

	if peek_byte_2(test_buffer) != 'f' {
		t.Fatalf("First byte of test_buffer should be f!\n")
	}

	if peek_byte_2(test_buffer) != 'f' {
		t.Fatalf("First byte of test_buffer should be f!\n")
	}

	// 2 - no bytes left

	test_buffer.ReadByte()
	test_buffer.ReadByte()
	test_buffer.ReadByte()
	test_buffer.ReadByte()
	test_buffer.ReadByte()

	if peek_byte_2(test_buffer) != byte(0) {
		t.Fatalf("First byte of test_buffer should be empty!\n")
	}
}

func Benchmark_peek_byte_2(b *testing.B) {
	test_buffer := bytes.NewBuffer([]byte("foobar.txt"))

	for i := 0; i < b.N; i++ {
		peek_byte_2(test_buffer)
	}
}
