## stages

For each file name there should be two stages:
1. replace characters
2. remove characters

Problem:\
The removal stage might remove a character between two replacement characters.\
=> the removal stage needs to check that it does not write consecutive \
replacement characters

#### replace stage

- replace whitespace with `_`
- replace umlauts with "ae", "oe", "ue"

#### removal stage

- remove non-ascii characters
- remove shell characters that have a special meaning
- remove characters that are not suitable for file names

The removal stage could be done entirely with a RangeTable.
