package main

import "bytes"

func peek_byte_1(buffer *bytes.Buffer) (b byte) {
	b, _ = buffer.ReadByte()

	buffer.UnreadByte()

	return b
}

func peek_byte_2(buffer *bytes.Buffer) byte {
	if buffer.Len() < 1 {
		return 0
	}
	return buffer.Bytes()[0]
}
