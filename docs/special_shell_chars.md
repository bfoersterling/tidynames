Characters that should not be allowed in file names:
- `!` (33) (executes history commands, needs to be escaped)
- `"` (34) (file name needs to be enclosed by single quotes)
- `#` (35) (comment, file name needs to be enclosed by single quotes)
- `$` (36) (start of var, will expand)
- (`%` (37) (not harmful, but not helpful as part of a file name either))
- `&` (38) (background command, needs to be enclosed by single quotes)
- `'` (39) (needs to be enclosed by double quotes)
- `(` (40) (needs to be enclosed by single quotes)
- `)` (41) (needs to be enclosed by single quotes)
- `*` (42) (wildcard, needs to be enclosed by single quotes)
- (`+` (43) (not harmful, but is it useful in file names?))
- (`,` (44) (not harmful, but is it useful in file names?))

Should be allowed, but maybe not as the first character:
- `-` (45)

Should not be allowed:
- `/` (46) (part of paths)

48-57 are numbers and should be allowed.

Should not be allowed:
- `;` (59) (end of command, needs to be enclosed by quotes)
- `>` (60) (needs to be enclosed by quotes)
- `=` (61) (assignment, needs to be enclosed by quotes)
- `<` (62) (needs to be enclosed by quotes)
- `?` (63) (needs to be enclosed by quotes)

Should be allowed:
- `@` (64)

65-90 are A-Z and should be allowed but converted to lowercase.
