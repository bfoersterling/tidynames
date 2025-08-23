package main

import "bytes"

func get_last_rune_bb(buffer *bytes.Buffer) rune {
	if buffer.Len() == 0 {
		return 0
	}

	buffer_str := buffer.String()

	return rune(buffer_str[len(buffer_str)-1])
}

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
