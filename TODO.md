- [ ] Create functions to convert bytes to megabytes, megabytes to byte, et cetera.

`collections`
- [ ] Make a `Map`, `Filter`, and `Reduce` iterators.
- [ ] Make an async module.

`fayl`:
- [ ] Make sure when archiving a folder, the path that is passed into the `writer.Create` method is not absolute.
- [ ] Make a `Path.Join` function that returns a new path instead of appending to the original path.
- BUG: In `Entry.Create`, the use of a suffix is wrong. There is already a internal type property. This results in crashes in some cases.
- BUG: The `SuffixFromString` function does not differentiate between a directory and a suffix not defined in the mapping.