- `ä` (228) -> should be replaced by "ae"
- `ö` (246) -> should be replaced by "oe"
- `ü` (252) -> should be replaced by "ue"

When a bytes.Buffer is broken down to bytes \
umlauts seem to be two bytes long.\
An `ö` is `Ã` (195) + `¶` (182).

You need to use `ReadRune()` instead of `ReadByte()`.
