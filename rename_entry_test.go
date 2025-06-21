package main

import (
	"os"
	"testing"
)

func Test_rename_entry(t *testing.T) {
	// 1 - rename absolute path to absolute path will work
	tmp_dir := t.TempDir()

	tmp_file, err := os.CreateTemp(tmp_dir, "Foo bar")

	if err != nil {
		t.Fatalf("Creating temp file failed!\n")
	}

	err = rename_entry(tmp_file.Name(), tmp_dir+"gotest", false)

	if err != nil {
		t.Fatalf("rename_entry failed with err:\n"+
			"%v\n", err)
	}

	_, err = os.Stat(tmp_dir + "gotest")

	if err != nil {
		t.Fatalf("os.Stat failed with err:\n"+
			"%v\n", err)
	}

	// 2 - rename_entry() will fail if absolute path will be renamed
	// to relative path (basename)
	tmp_dir = t.TempDir()

	tmp_file, err = os.CreateTemp(tmp_dir, "BAR FOO")

	if err != nil {
		t.Fatalf("Creating temp file failed!\n")
	}

	err = rename_entry(tmp_file.Name(), "gotest", false)

	if err == nil {
		t.Fatalf("rename_entry() should fail here.\n")
	}
}
