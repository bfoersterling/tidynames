#### tidynames .

What should `tidynames .` do?

Comparing with `mv`:\
You can't rename the current dir with `mv`:
```
$ mv -v . foobar
mv: cannot move '.' to 'foobar': Device or resource busy
```
(using `mv` with absolute path works though)

Tidying up all files in the current dir seems to be the most appropriate.

#### tidynames mydir and tidynames mydir/

`tidynames mydir` and `tidynames mydir/` should both just tidy the name \
of the given dir, not the files inside.\
Bash tab completion will add the trailing slash automatically.\
And the user will probably want to tidy this entry only.

If the user wants to tidy all files in the dir he needs to add an asterix:
```
tidynames mydir/*
```
