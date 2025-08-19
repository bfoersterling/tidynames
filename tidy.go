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

type replace_config struct {
	nonascii   rune
	whitespace rune
}

func (rc replace_config) tidy_bytes(name []byte) (proper_name []byte) {
	var name_buffer bytes.Buffer

	for _, r := range name {
		if (32 < r) && (r < 127) {
			name_buffer.WriteByte(r)
			continue
		}

		// whitespace
		if (r <= 32) && (rc.whitespace != 0) {
			name_buffer.WriteByte(byte(rc.whitespace))
			continue
		}

		if (r >= 127) && (rc.nonascii != 0) {
			name_buffer.WriteByte(byte(rc.nonascii))
		}
	}

	proper_name = name_buffer.Bytes()

	if rc.whitespace != 0 {
		proper_name = replace_whitespace(proper_name, rc.whitespace)
	} else {
		proper_name = replace_whitespace(proper_name, 0)
	}

	proper_name = bytes.ToLower(proper_name)

	return
}

func (rc replace_config) tidy_string(name string) (proper_name string) {
	for _, r := range name {
		if (32 < r) && (r < 127) {
			proper_name += string(r)
			continue
		}

		if (r <= 32) && (rc.whitespace != 0) {
			proper_name += string(rc.whitespace)
			continue
		}

		if (r >= 127) && (rc.nonascii != 0) {
			proper_name += string(rc.nonascii)
		}
	}

	proper_name = strings.ToLower(proper_name)

	return
}

func (rc replace_config) tidy_entry(entry_path string, dry_run bool, writer io.Writer) (err error) {
	new_name := rc.tidy_bytes([]byte(path.Base(entry_path)))

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

func (rc replace_config) tidy_entries(args cli_args, entries []string, writer io.Writer) (err error) {
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
		tidy_err := rc.tidy_entry(arg, args.dry_run, os.Stdout)
		if tidy_err != nil {
			fmt.Printf("%v", tidy_err)
		}
	}

	return
}

func replace_whitespace_buffer(name *bytes.Buffer, substitute rune) *bytes.Buffer {
	result := bytes.NewBuffer([]byte(""))
	substitute_written := false

	for b, err := name.ReadByte(); err == nil; b, err = name.ReadByte() {
		if !unicode.IsSpace(rune(b)) {
			result.WriteByte(b)
			substitute_written = false
			continue
		}

		if unicode.IsSpace(rune(b)) && !substitute_written {
			result.WriteByte(byte(substitute))
			substitute_written = true
			continue
		}
	}

	return result
}

func replace_whitespace(name []byte, substitute rune) []byte {
	tokens := bytes.Fields(name)
	var substitute_bytes []byte = []byte("")

	if substitute != 0 && !unicode.IsSpace(substitute) {
		substitute_bytes = append(substitute_bytes, byte(substitute))
	}

	return bytes.Join(tokens, substitute_bytes)
}
