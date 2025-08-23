package main

// peeks at the next byte
// returns 0 if at end of bytes
func peek_byte(raw_bytes []byte, current_index int) byte {
	if (current_index + 1) >= len(raw_bytes) {
		return 0
	}
	return raw_bytes[current_index+1]
}

func get_last_byte_from_bytes(raw_bytes []byte) byte {
	if len(raw_bytes) == 0 {
		return 0
	}

	return raw_bytes[len(raw_bytes)-1]
}

func get_last_rune_from_bytes(raw_bytes []byte) rune {
	if len(raw_bytes) == 0 {
		return 0
	}

	return rune(raw_bytes[len(raw_bytes)-1])
}
