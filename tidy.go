package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type replace_config struct {
	nonascii   rune
	whitespace rune
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
	new_name := rc.tidy_string(path.Base(entry_path))

	// trailing slashes will cause path.Base to not return the parent dir
	entry_path = strings.TrimRight(entry_path, "/")

	if path.Base(entry_path) == new_name {
		fmt.Fprintf(writer, "%q is already tidy.\n", entry_path)
		return
	}

	err = rename_entry(entry_path, path.Dir(entry_path)+"/"+new_name, dry_run)

	if err != nil {
		return
	}

	if dry_run {
		fmt.Fprintf(writer, "(")
	}

	fmt.Fprintf(writer, "%q -> %q", entry_path, path.Dir(entry_path)+"/"+new_name)

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
