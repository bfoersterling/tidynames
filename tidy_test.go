package main

import (
	"bytes"
	"os"
	"path"
	"regexp"
	"testing"
)

func Test_tidy_bytes(t *testing.T) {
	// 1 - remove whitespace
	rc := replace_config{}

	input := []byte("Foo bar.txt")
	expected_result := []byte("foobar.txt")
	test_result := rc.tidy_bytes(input)

	if bytes.Compare(test_result, expected_result) != 0 {
		t.Fatalf("test_result and expected_result differ.\n"+
			"test_result: %q\n"+
			"expected_result: %q\n", test_result, expected_result)
	}

	// 2 - unicode character
	input = []byte("\u2318foo bar.txt")
	expected_result = []byte("foobar.txt")
	test_result = rc.tidy_bytes(input)

	if bytes.Compare(test_result, expected_result) != 0 {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input, expected_result, test_result)
	}

	// 3 - replace whitespace with underscore
	rc = replace_config{
		whitespace: '_',
	}

	input = []byte("Foo bar.")
	expected_result = []byte("foo_bar.")
	test_result = rc.tidy_bytes(input)

	if bytes.Compare(test_result, expected_result) != 0 {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input, expected_result, test_result)
	}
}

func Benchmark_tidy_bytes(b *testing.B) {
	rc := replace_config{
		whitespace: '_',
	}

	for i := 0; i < b.N; i++ {
		rc.tidy_bytes([]byte("foo Bär.txt"))
	}
}

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

func Benchmark_tidy_string(b *testing.B) {
	rc := replace_config{
		whitespace: '_',
	}

	for i := 0; i < b.N; i++ {
		rc.tidy_string("foo Bär.txt")
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

func Test_replace_whitespace(t *testing.T) {
	// 1
	rc := replace_config{whitespace: '_'}

	input := bytes.NewBuffer([]byte("foo\t  bar.txt"))
	expected_result := bytes.NewBuffer([]byte("foo_bar.txt"))
	test_result := replace_whitespace(input, rc.whitespace)

	if !bytes.Equal(test_result.Bytes(), expected_result.Bytes()) {
		t.Fatalf("test_result.Bytes() and expected_result.Bytes() differ.\n"+
			"test_result.Bytes(): %s\nexpected_result.Bytes(): %s\n",
			test_result.Bytes(), expected_result.Bytes())
	}

	// 2
	rc = replace_config{whitespace: '_'}
	input = bytes.NewBuffer([]byte("foo\t  bar \t foo.txt"))
	expected_result = bytes.NewBuffer([]byte("foo_bar_foo.txt"))
	test_result = replace_whitespace(input, rc.whitespace)

	if !bytes.Equal(test_result.Bytes(), expected_result.Bytes()) {
		t.Fatalf("test_result.Bytes() and expected_result.Bytes() differ.\n"+
			"test_result.Bytes(): %s\nexpected_result.Bytes(): %s\n",
			test_result.Bytes(), expected_result.Bytes())
	}
}

func Benchmark_replace_whitespace(b *testing.B) {
	rc := replace_config{whitespace: '_'}
	input := bytes.NewBuffer([]byte("foo\t  bar \t foo.txt"))

	for i := b.N; i < b.N; i++ {
		replace_whitespace(input, rc.whitespace)
	}
}

func Test_replace_whitespace_fields(t *testing.T) {
	// 1
	rc := replace_config{
		whitespace: '_',
	}

	input := []byte("foo\t  bar \t foo.txt")
	expected_result := []byte("foo_bar_foo.txt")
	test_result := replace_whitespace_fields(input, rc.whitespace)

	if !bytes.Equal(test_result, expected_result) {
		t.Fatalf("test_result is %q, but should be %q.\n",
			test_result, expected_result)
	}

	// 2 - null rune value
	input = []byte("foo\t  bar \t foo.txt")
	expected_result = []byte("foobarfoo.txt")
	test_result = replace_whitespace_fields(input, 0)

	if !bytes.Equal(test_result, expected_result) {
		t.Fatalf("test_result is %q, but should be %q.\n",
			test_result, expected_result)
	}
}

func Benchmark_tidy_substiutes_fields(b *testing.B) {
	// 1
	rc := replace_config{
		whitespace: '_',
	}

	input := []byte("foo\t  bar \t foo.txt")

	for i := 0; i < b.N; i++ {
		replace_whitespace_fields(input, rc.whitespace)
	}
}
