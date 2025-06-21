package main

import (
	"errors"
	"fmt"
	"os"
)

func rename_entry(old_name string, new_name string, dry_run bool) (err error) {
	if !dry_run {
		err = os.Rename(old_name, new_name)
	}

	if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("The file or directory %q does not exist.\n", old_name)
		return
	}

	if errors.Is(err, os.ErrExist) {
		err = fmt.Errorf("The file or directory %q already exists.\n", new_name)
		return
	}

	return
}
