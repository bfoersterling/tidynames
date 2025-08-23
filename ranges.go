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
