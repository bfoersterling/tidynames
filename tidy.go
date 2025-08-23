package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"unicode"
)

type tidy_config struct {
	replacement_char rune
}

func (tc tidy_config) tidy_bytes(name []byte) []byte {
	input_buffer := bytes.NewBuffer(bytes.ToLower(name))

	replace_whitespace(input_buffer, tc.replacement_char)

	replace_umlauts(input_buffer)

	tc.removal_stage(input_buffer)

	return input_buffer.Bytes()
}

func (tc tidy_config) tidy_string(name string) (proper_name string) {
	for _, r := range name {
		if (32 < r) && (r < 127) {
			proper_name += string(r)
			continue
		}

		if (r <= 32) && (tc.replacement_char != 0) {
			proper_name += string(tc.replacement_char)
			continue
		}

		if (r >= 127) && (tc.replacement_char != 0) {
			proper_name += string(tc.replacement_char)
		}
	}

	proper_name = strings.ToLower(proper_name)

	return
}

func (tc tidy_config) tidy_entry(entry_path string, dry_run bool, writer io.Writer) (err error) {
	new_name := tc.tidy_bytes([]byte(path.Base(entry_path)))

	// trailing slashes will cause path.Base to not return the parent dir
	entry_path = strings.TrimRight(entry_path, "/")

	if path.Base(entry_path) == string(new_name) {
		fmt.Fprintf(writer, "%q is already tidy.\n", entry_path)
		return
	}

	err = rename_entry(entry_path, path.Dir(entry_path)+"/"+string(new_name), dry_run)

	if err != nil {
		return
	}

	if dry_run {
		fmt.Fprintf(writer, "(")
	}

	fmt.Fprintf(writer, "%q -> %q", entry_path, path.Dir(entry_path)+"/"+string(new_name))

	if dry_run {
		fmt.Fprintf(writer, ")")
	}

	fmt.Fprintf(writer, "\n")

	return
}

func (tc tidy_config) tidy_entries(args cli_args, entries []string, writer io.Writer) (err error) {
	// this var will only be needed if "." was passed to tidynames
	var dir_entries []os.DirEntry

	if (len(entries) == 1) && (entries[0] == ".") {
		dir_entries, err = os.ReadDir(".")

		if err != nil {
			return
		}

		entries = nil

		for _, dir_entry := range dir_entries {
			entries = append(entries, dir_entry.Name())
		}
	}

	for _, arg := range entries {
		tidy_err := tc.tidy_entry(arg, args.dry_run, os.Stdout)
		if tidy_err != nil {
			fmt.Printf("%v", tidy_err)
		}
	}

	return
}

func (tc tidy_config) removal_stage(name *bytes.Buffer) {
	name_copy := name.String()

	rt := get_whitelist_rt()

	name.Reset()

	for _, r := range name_copy {
		// don't write replacement char at beginning of name
		if r == tc.replacement_char && name.Len() == 0 {
			continue
		}

		// don't write consecutive replacement chars
		if (r == tc.replacement_char) &&
			(get_last_rune_from_bytes(name.Bytes()) == tc.replacement_char) {
			continue
		}

		if unicode.In(r, &rt) {
			name.WriteRune(r)
			continue
		}
	}
}

// remove characters that are not ascii codes between 32 and 127
// 127 is the control character DEL which should not be included
func remove_nonascii(name *bytes.Buffer) {
	name_copy := name.Bytes()

	name.Reset()

	for _, b := range name_copy {
		// printable ascii characters
		if (32 < b) && (b < 127) {
			name.WriteByte(b)
		}
	}
}

func remove_special_chars(name *bytes.Buffer) {
	name_copy := name.String()

	rt := get_special_char_rt()

	name.Reset()

	for _, r := range name_copy {
		if !unicode.In(r, &rt) {
			name.WriteRune(r)
		}
	}
}

// only replaces lower case umlauts
// upper case umlauts should be converted by bytes.Lower() previously
func replace_umlauts(name *bytes.Buffer) {
	name_copy := name.String()
	name.Reset()

	for _, r := range name_copy {
		if r == 228 {
			name.WriteString("ae")
			continue
		}

		if r == 246 {
			name.WriteString("oe")
			continue
		}

		if r == 252 {
			name.WriteString("ue")
			continue
		}

		name.WriteRune(r)
	}
}

// replace whitespace by replacement char
// but do not write consecutive replacement chars
func replace_whitespace(name *bytes.Buffer, substitute rune) {
	name_copy := name.Bytes()
	name.Reset()

	for i, b := range name_copy {
		if !unicode.IsSpace(rune(b)) {
			name.WriteByte(b)
			continue
		}

		// don't write substitute at beginning of the file name
		if name.Len() == 0 {
			continue
		}

		// don't write replacement char before file extensions
		if peek_byte(name_copy, i) == '.' {
			continue
		}

		// don't write replacement char if next byte will be replaced
		if unicode.IsSpace(rune(peek_byte(name_copy, i))) {
			continue
		}

		if unicode.IsSpace(rune(b)) {
			name.WriteByte(byte(substitute))
		}
	}
}

func replace_whitespace_fields(name []byte, substitute rune) []byte {
	tokens := bytes.Fields(name)
	var substitute_bytes []byte = []byte("")

	if substitute != 0 && !unicode.IsSpace(substitute) {
		substitute_bytes = append(substitute_bytes, byte(substitute))
	}

	return bytes.Join(tokens, substitute_bytes)
}
