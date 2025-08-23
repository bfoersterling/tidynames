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
	tc := tidy_config{}

	input := []byte("Foo bar.txt")
	expected_result := []byte("foobar.txt")
	test_result := tc.tidy_bytes(input)

	if bytes.Compare(test_result, expected_result) != 0 {
		t.Fatalf("test_result and expected_result differ.\n"+
			"test_result: %q\n"+
			"expected_result: %q\n", test_result, expected_result)
	}

	// 2 - unicode character
	input = []byte("\u2318foo bar.txt")
	expected_result = []byte("foobar.txt")
	test_result = tc.tidy_bytes(input)

	if bytes.Compare(test_result, expected_result) != 0 {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input, expected_result, test_result)
	}

	// 3 - replace whitespace with underscore
	tc = tidy_config{
		replacement_char: '_',
	}

	input = []byte("Foo bar.")
	expected_result = []byte("foo_bar.")
	test_result = tc.tidy_bytes(input)

	if bytes.Compare(test_result, expected_result) != 0 {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input, expected_result, test_result)
	}
}

func Benchmark_tidy_bytes(b *testing.B) {
	tc := tidy_config{
		replacement_char: '_',
	}

	for i := 0; i < b.N; i++ {
		tc.tidy_bytes([]byte("foo Bär.txt"))
	}
}

func Test_tidy_string(t *testing.T) {
	// 1 - remove whitespace
	tc := tidy_config{}

	input_string := "Foo bar."
	expected_result := "foobar."
	test_result := tc.tidy_string(input_string)

	if test_result != expected_result {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input_string, expected_result, test_result)
	}

	// 2 - unicode character
	input_string = "\u2318foo bar.txt"
	expected_result = "foobar.txt"
	test_result = tc.tidy_string(input_string)

	if test_result != expected_result {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input_string, expected_result, test_result)
	}

	// 3 - replace whitespace with underscore
	tc = tidy_config{
		replacement_char: '_',
	}

	input_string = "Foo bar."
	expected_result = "foo_bar."
	test_result = tc.tidy_string(input_string)

	if test_result != expected_result {
		t.Fatalf("%q should be converted to %q, but was converted to %q\n",
			input_string, expected_result, test_result)
	}
}

func Benchmark_tidy_string(b *testing.B) {
	tc := tidy_config{
		replacement_char: '_',
	}

	for i := 0; i < b.N; i++ {
		tc.tidy_string("foo Bär.txt")
	}
}

func Test_tidy_entry(t *testing.T) {
	// 1 - tidy absolute path
	tc := tidy_config{}

	tmp_dir := t.TempDir()

	var test_buf bytes.Buffer

	tmp_file, err := os.CreateTemp(tmp_dir, "Foo bar")

	if err != nil {
		t.Fatalf("Creating temp file failed!\n")
	}

	err = tc.tidy_entry(tmp_file.Name(), false, &test_buf)

	if err != nil {
		t.Fatalf("tc.tidy_entry failed with err:\n"+
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

	err = tc.tidy_entry(path.Base(tmp_file.Name()), false, &test_buf)

	if err != nil {
		t.Fatalf("tc.tidy_entry failed with err:\n"+
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

func Test_removal_stage(t *testing.T) {
	// 1
	input := bytes.NewBuffer([]byte("l_ö_hne.txt"))
	expected_result := bytes.NewBuffer([]byte("l_hne.txt"))
	tc := tidy_config{replacement_char: '_'}

	tc.removal_stage(input)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input and expected_result differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %q\n",
			input.Bytes(), expected_result.Bytes())
	}

	// 2 - multiple replacement chars at start of name
	input = bytes.NewBuffer([]byte("___hne.txt"))
	expected_result = bytes.NewBuffer([]byte("hne.txt"))
	tc = tidy_config{replacement_char: '_'}

	tc.removal_stage(input)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input and expected_result differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %q\n",
			input.Bytes(), expected_result.Bytes())
	}

	// 3 - multiple replacement chars before period
	input = bytes.NewBuffer([]byte("loehne___.txt"))
	expected_result = bytes.NewBuffer([]byte("loehne.txt"))
	tc = tidy_config{replacement_char: '_'}

	tc.removal_stage(input)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input and expected_result differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %q\n",
			input.Bytes(), expected_result.Bytes())
	}
}

func Test_remove_nonascii(t *testing.T) {
	// 1
	input := bytes.NewBuffer([]byte("Löhne.txt"))
	expected_result := bytes.NewBuffer([]byte("Lhne.txt"))
	remove_nonascii(input)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input and expected_result differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %q\n",
			input.Bytes(), expected_result.Bytes())
	}
}

func Benchmark_remove_nonascii(b *testing.B) {
	input := bytes.NewBuffer([]byte("Löhöhöhne.txt"))

	for i := 0; i < b.N; i++ {
		remove_nonascii(input)
	}
}

func Test_remove_special_chars(t *testing.T) {
	// 1
	input := bytes.NewBuffer([]byte("wei,rd!file.txt"))
	expected_result := bytes.NewBuffer([]byte("weirdfile.txt"))
	remove_special_chars(input)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input and expected_result differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %q\n",
			input.Bytes(), expected_result.Bytes())
	}
}

func Benchmark_remove_special_chars(b *testing.B) {
	input := bytes.NewBuffer([]byte("wei,rd!file.txt"))

	for i := 0; i < b.N; i++ {
		remove_special_chars(input)
	}
}

func Test_replace_umlauts(t *testing.T) {
	// 1
	input := bytes.NewBuffer([]byte("überflüssig.txt"))
	expected_result := bytes.NewBuffer([]byte("ueberfluessig.txt"))

	replace_umlauts(input)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input.Bytes and expected_result.Bytes() differ.\n"+
			"input.Bytes: %q\nexpected_result.Bytes(): %q\n", input.Bytes(), expected_result.Bytes())
	}
}

func Benchmark_replace_umlauts(b *testing.B) {
	input := bytes.NewBuffer([]byte("überflüssig.txt"))

	for i := 0; i < b.N; i++ {
		replace_umlauts(input)
	}
}

func Test_replace_whitespace(t *testing.T) {
	// 1
	tc := tidy_config{replacement_char: '_'}

	input := bytes.NewBuffer([]byte("foo\t  bar.txt"))
	expected_result := bytes.NewBuffer([]byte("foo_bar.txt"))
	replace_whitespace(input, tc.replacement_char)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input.Bytes() and expected_result.Bytes() differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %s\n",
			input.Bytes(), expected_result.Bytes())
	}

	// 2
	tc = tidy_config{replacement_char: '_'}
	input = bytes.NewBuffer([]byte("foo\t  bar \t foo.txt"))
	expected_result = bytes.NewBuffer([]byte("foo_bar_foo.txt"))
	replace_whitespace(input, tc.replacement_char)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input.Bytes() and expected_result.Bytes() differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %s\n",
			input.Bytes(), expected_result.Bytes())
	}

	// 3 - file names should not start with a replacement character
	tc = tidy_config{replacement_char: '_'}
	input = bytes.NewBuffer([]byte("\t  bar \t foo.txt"))
	expected_result = bytes.NewBuffer([]byte("bar_foo.txt"))
	replace_whitespace(input, tc.replacement_char)

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input.Bytes() and expected_result.Bytes() differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %s\n",
			input.Bytes(), expected_result.Bytes())
	}

	// 4 - there shall be no replacement chars before periods
	input = bytes.NewBuffer([]byte("foo \t.txt"))
	expected_result = bytes.NewBuffer([]byte("foo.txt"))
	replace_whitespace(input, rune('_'))

	if !bytes.Equal(input.Bytes(), expected_result.Bytes()) {
		t.Fatalf("input.Bytes() and expected_result.Bytes() differ.\n"+
			"input.Bytes(): %s\nexpected_result.Bytes(): %s\n",
			input.Bytes(), expected_result.Bytes())
	}
}

func Benchmark_replace_whitespace(b *testing.B) {
	tc := tidy_config{replacement_char: '_'}
	input := bytes.NewBuffer([]byte("foo\t  bar \t foo.txt"))

	for i := 0; i < b.N; i++ {
		replace_whitespace(input, tc.replacement_char)
	}
}

func Test_replace_whitespace_fields(t *testing.T) {
	// 1
	tc := tidy_config{
		replacement_char: '_',
	}

	input := []byte("foo\t  bar \t foo.txt")
	expected_result := []byte("foo_bar_foo.txt")
	test_result := replace_whitespace_fields(input, tc.replacement_char)

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

func Benchmark_replace_whitespace_fields(b *testing.B) {
	// 1
	tc := tidy_config{
		replacement_char: '_',
	}

	input := []byte("foo\t  bar \t foo.txt")

	for i := 0; i < b.N; i++ {
		replace_whitespace_fields(input, tc.replacement_char)
	}
}
