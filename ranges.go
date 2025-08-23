package main

import "unicode"

// special shell characters
// as well as some characters like + or , you do not want in file names
func get_special_char_rt() unicode.RangeTable {
	var ranges []unicode.Range16

	undesirable_chars_1 := unicode.Range16{
		Lo:     33,
		Hi:     44,
		Stride: 1,
	}

	// the slash
	path_separator := unicode.Range16{
		Lo:     47,
		Hi:     47,
		Stride: 1,
	}

	undesirable_chars_2 := unicode.Range16{
		Lo:     59,
		Hi:     63,
		Stride: 1,
	}

	undesirable_chars_3 := unicode.Range16{
		Lo:     91,
		Hi:     94,
		Stride: 1,
	}

	backticks := unicode.Range16{
		Lo:     96,
		Hi:     96,
		Stride: 1,
	}

	undesirable_chars_4 := unicode.Range16{
		Lo:     123,
		Hi:     126,
		Stride: 1,
	}

	ranges = append(ranges, undesirable_chars_1, path_separator,
		undesirable_chars_2, undesirable_chars_3, backticks, undesirable_chars_4)

	return unicode.RangeTable{R16: ranges}
}

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
