package main

import "unicode"

// receive a RangeTable of allowed characters
// for the use in the removal stage
func get_whitelist_rt() unicode.RangeTable {
	var ranges []unicode.Range16

	dash_dot := unicode.Range16{
		Lo:     45,
		Hi:     46,
		Stride: 1,
	}

	numbers := unicode.Range16{
		Lo:     48,
		Hi:     57,
		Stride: 1,
	}

	at := unicode.Range16{
		Lo:     64,
		Hi:     64,
		Stride: 1,
	}

	underscore := unicode.Range16{
		Lo:     95,
		Hi:     95,
		Stride: 1,
	}

	lowercase_letters := unicode.Range16{
		Lo:     97,
		Hi:     122,
		Stride: 1,
	}

	ranges = append(ranges, dash_dot, numbers, at, underscore, lowercase_letters)

	return unicode.RangeTable{R16: ranges}
}
