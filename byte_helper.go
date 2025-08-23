package main

// peeks at the next byte
// returns 0 if at end of bytes
func peek_byte(raw_bytes []byte, current_index int) byte {
	if (current_index + 1) >= len(raw_bytes) {
		return 0
	}
	return raw_bytes[current_index+1]
}
