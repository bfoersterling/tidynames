package main

import (
	"bytes"
	"os"
	"path"
	"regexp"
	"testing"
)

func Test_tidy_string(t *testing.T) {
	// 1 - remove whitespace
	rc := replace_config{}

	input_string := "Foo bar."
	expected_result := "foobar."
	test_result := rc.tidy_string(input_string)

	if test_result != expected_result {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input_string, expected_result, test_result)
	}

	// 2 - unicode character
	input_string = "\u2318foo bar.txt"
	expected_result = "foobar.txt"
	test_result = rc.tidy_string(input_string)

	if test_result != expected_result {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input_string, expected_result, test_result)
	}

	// 3 - replace whitespace with underscore
	rc = replace_config{
		whitespace: '_',
	}

	input_string = "Foo bar."
	expected_result = "foo_bar."
	test_result = rc.tidy_string(input_string)

	if test_result != expected_result {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input_string, expected_result, test_result)
	}
}

func Test_tidy_entry(t *testing.T) {
	// 1 - tidy absolute path
	rc := replace_config{}

	tmp_dir := t.TempDir()

	var test_buf bytes.Buffer

	tmp_file, err := os.CreateTemp(tmp_dir, "Foo bar")

	if err != nil {
		t.Fatalf("Creating temp file failed!\n")
	}

	err = rc.tidy_entry(tmp_file.Name(), false, &test_buf)

	if err != nil {
		t.Fatalf("rc.tidy_entry failed with err:\n"+
			"%v\n", err)
	}

	matches, err := regexp.MatchString(`^"/tmp/Test_tidy_entry.*Foo bar.*-> "/tmp/Test_tidy_entry.*foobar.*"`, test_buf.String())

	if err != nil {
		t.Fatalf("regexp.MatchString returned err:\n"+
			"%v\n", err)
	}

	if !matches {
		t.Fatalf("test_buf.String() does not match the regex.\n"+
			"test_buf.String():\n%q\n", test_buf.String())
	}

	test_buf.Reset()

	// 2 - tidy relative path

	t.Chdir(tmp_dir)

	tmp_file, err = os.CreateTemp(tmp_dir, "Rel File")

	if err != nil {
		t.Fatalf("Creating temp file failed!\n")
	}

	err = rc.tidy_entry(path.Base(tmp_file.Name()), false, &test_buf)

	if err != nil {
		t.Fatalf("rc.tidy_entry failed with err:\n"+
			"%v\n", err)
	}

	matches, err = regexp.MatchString(`^"Rel File.*" -> "./relfile.*"\n$`, test_buf.String())

	if err != nil {
		t.Fatalf("regexp.MatchString returned err:\n"+
			"%v\n", err)
	}

	if !matches {
		t.Fatalf("test_buf.String() does not match the regex.\n"+
			"test_buf.String():\n%v\n", test_buf.String())
	}
}
