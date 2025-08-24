# tidynames

Opinionated file and directory renamer with dry run option.
- replaces whitespace by non consecutive `_`s
- removes non-ascii characters
- removes special shell characters and characters unsuited for naming entries
- converts to lowercase

## installation

Download the latest release
```
wget https://github.com/bfoersterling/tidynames/releases/latest/download/tidynames_linux_x86_64 -O /tmp/tidynames
sudo install -v -m 755 /tmp/tidynames /usr/local/bin
```

Or build it.\
Clone this repo and:
```
sudo just install
```

## usage

Dryrun:
```
$ ./tidynames -n test_files/*
[dry run]
("test_files/  leading and trailing ws  .txt" -> "test_files/leading_and_trailing_ws.txt")
("test_files/Foo bar.txt" -> "test_files/foo_bar.txt")
("test_files/This is an aweful ｜ Name for a file ® (file 1) [QJHPlKPOc78].m4a.bak" -> "test_files/this_is_an_aweful_name_for_a_file_file_1_qjhplkpoc78.m4a.bak")
("test_files/Weird Files" -> "test_files/weird_files")
"test_files/already_tidy.txt" is already tidy.
("test_files/foo   bar.txt" -> "test_files/foo_bar.txt")
("test_files/Übernatürlich .txt" -> "test_files/uebernatuerlich.txt")
```

Rename all files and dirs in the current dir:
```
$ tidynames *
"Foo bar.txt" -> "./foobar.txt"
"This is an aweful ｜ Name for a file ® (file 1) [QJHPlKPOc78].m4a.bak" -> "./thisisanawefulnameforafile(file1)[qjhplkpoc78].m4a.bak"
"Weird Files" -> "./weirdfiles"
"already_tidy.txt" is already tidy.
```

## TODO

- implement option `-r` to rename entire dir recursively
=> complex feature - need to take renamed parent dirs into account before traversing down
- options to replace whitespace or non-ascii chars with specific chars
(struct `replace_config` already exists)
