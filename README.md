# tidynames

Removes whitespace and non-ascii characters from file and directory names.

## installation

Clone this repo and:
```
sudo just install
```

## usage

Dryrun:
```
$ tidynames -n *
---dry run---
("Foo bar.txt" -> "./foobar.txt")
("This is an aweful ｜ Name for a file ® (file 1) [QJHPlKPOc78].m4a.bak" -> "./thisisanawefulnameforafile(file1)[qjhplkpoc78].m4a.bak")
("Weird Files" -> "./weirdfiles")
"already_tidy.txt" is already tidy.
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

- remove braces `(`,`)` and `[`, `]` - because they force single quotes around file names as they have special meaning in bash
- implement option `-r` to rename entire dir recursively
=> complex feature - need to take renamed parent dirs into account before traversing down
- options to replace whitespace or non-ascii chars with specific chars
(struct `replace_config` already exists)
