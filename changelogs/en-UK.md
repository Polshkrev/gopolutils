# Changelog
## v1.12.0 - 2025-03-13: The Concurrent Update
For future changes, refer to [TODO.md](../TODO.md).

`Added`
- Added `SafeArray` struct.
- Added `SafeMap` struct.
- Added `SafeQueue` struct.
- Added `SafeStack` struct.
- `Set`
    - Added `NewSafeSet` function.
- `fayl`
    - Added private `readConcurrent` function.
    - Added `runtime` dependency.
- `Docs`
    - Added `SafeArray` documentation.
    - Added `SafeMap` documentation.
    - Added `SafeQueue` documentation.
    - Added `SafeStack` documentation.
    - Added `NewSafeSet` function documentation.

`Changed`
- `fayl`
    - `Read` is now concurrent.
    - `Read` no longer uses the absolute path.
    - The functions `Read`, `readRawObject`, `readRawList`, `readerListDispatch`, `ReadList`, and `ReadObject` now use a pointer to a `Path`.
    - On `Linux`, `Path.Root` simply returns `/`.

`Removed`
- `Set`
    - The `size` property has been removed.

`Fixed`
- `Docs`
    - Fixed `Map` documentation now mentions the implementation of `View`.
## v1.11.0 - 2025-03-07: The Must Update
`Added`
- A new changelog has appeared. lol.
- `Exception`
    - A new `Must` function has been added.
    - A new private `assignMessage` method has been added.
- `Logger`
    - Added a public getter for the `level` property.
    - Added a public getter and setter for the `name` property.
    - Added a deprecation notice for the private `append` method.
- `Docs`
    - Added documentation for the `Must` function in `exception.md`.
    - Added mention of the new getter and setter methods for the `level` and `name` property in `logger.md`.

`Changed`
- `Logger`
    - The `SetLevel` function now simply sets the level of the logger rather than checking if the parametre and property are the same.
## v1.10.1 - 2025-01-12
`Fixed`
- `fayl`
    - `Read` now uses the abosulte path instead of just obtaining the suffix of the path parametre.
## v1.10.0 - 2025-01-11: The Serialization Update
`Added`
- Added the `goserialize` dependency.
- Added `meta.toml`.
- `Exception`
    - A new public getter for the message property has been added.
    - A new public getter for the name property has been added.
- `Docs`
    - Deprecation notices for the types in the `fayl` package were added.

`Changed`
- `fayl`
    - All the public functions now take in a `Path` struct rather than just a string.
    - The private functions `readRawObject`, `readRawList`, and `readerListDispatch` now use the `Path` struct.
    - The private functions `readRawObject` and `readRawList` now take in a `goserialize.Reader` instead of the **now depricated** internal `Reader` type.
    - The private function `readerListDispatch` and the public function `ReadObject` now use the suffix of the given `Path` struct instead of using the **now depricated** `getFileType` internal function that takes in a string.
    - If the `Read` function fails, the error message uses the new `Message` method from the `Exception` struct instead of the full `Error` method. This is simply a formatting clarification.
    - The private function `readerListDispatch` and the public function `ReadObject` now uses the appropriate `goserialize.Reader` within the function. I.e `goserialize.JSONReader` for a JSON file type.

`Deprecated`
- `fayl`
    - Due to the introduction of the `goserialize` dependency, all the type aliases in the `read.go` file have been deprecated.

`Removed`
- `fayl`
    - The `path/filepath` dependency in `read.go` has been removed.
## v1.9.0 - 2025-01-09
`Added`
- `Mapping`
    - The mapping interface now uses the `View` interface instead of the just using the `Size` interface. The `View` interface also implements the `Size` interface.
- `Map`
    - A new `Collect` method has been added. This method has been added to adhere to the `View` interface.

`Changed`
- `Map`
    - The documentation comment for the `Map` struct now says `"A collection of key-value pairs".`
- `Docs`
    - The russian translation of `array.md` and `pair.md` now use the russian translation of the title.

`Fixed`
- `Docs`
    - Fixed a small capitalization typo in `README.md`.
## v1.8.0 - 2025-01-09
`Added`
- `Docs`
    - Added russian documentation.
    - Added `fayl` documentation links to the table of contents in english `README.md`.
    - Added `path.md` english documentation.

`Removed`
- `Docs`
    - Remove `Iterators` documentation link in english `README.md`

`Fixed`
- `Logger`
    - Fixed typos for the documentation comments for the public `Log` and `Close` methods.
- `Docs`
    - Fixed typos in the `array.md` english documentation.
    - Fixed typos in the `collections.md` english documentation.
    - Fixed typos in the `set.md` english documentation.
    - Fixed typos in the `stack.md` english documentation.
    - Fixed typos in the `exception.md` english documentation.
    - Fixed typos in the `version.md` english documentation.
## v1.7.1 - 2025-01-09
`Changed`
- `Queue`
    - The `Remove` method is now implemented.
## v1.7.0 - 2025-01-09
`Added`
- `fayl`
    - Added `Path` struct.

`Changed`
- `Logger`
    - The private global constants `__AVAILABLE_OUPUTS` and `__TIMESTAMP_FORMAT`, as well as the private global variable `__OUTPUT_COUNT` have been renamed to be more idiomatic.
- `Stack`
    - The `Remove` method is now implemented. The documentation comment has been updated to indicate as such.
- `Set`
    - If the `Return` method fails due to the set being empty, the returned error message simply now indicates as such rather than returning a message containing the irrelevant index passed into the method.

`Fixed`
- `Set`
    - Fixed typo in the returned error message in the `Update` method.
## v1.6.0 - 2024-12-26
`Added`
- `Version`
    - Added `NumberString` method.
## v1.5.0 - 2024-12-26
`Added`
- `Logger`
    - Added `SetLevel` method.
- `Version`
    - Added `IsZero` method.
- `Tests`
    - `Version`
        - Added `IsZero` tests.
- `Docs`
    - Added mention of `IsZero` method in the `version.md` english documentation.
    - Added mention of `IsPublic` method in the `version.md` english documentation.
    - Added mentions of numeric compare methods in the `version.md` english documentation.
## v1.4.2 - 2024-12-26
`Added`
- `fayl`
    - Added support for the `yml` file extension in the private `readerListDispatch` function in the `read.go` file.
## v1.4.1 - 2024-12-26
`Changed`
- `Logger`
    - The `outputs` property of the logger is now a sized array, replacing a slice.
`Fixed`
- `fayl`
    - Fixed typo in the documentation comment for `Reader` in the `read.go` file.
## v1.4.0 - 2024-12-25 🎅
`Added`
- Added the `BurntSushi/toml` dependency.
- Added the `gopkg.in/yaml.v2` dependency.
- `fayl`
    - Added standardized file types. I.e `JSONType`...
    - Added `Reader` type.
    - Added each specific reader of each supported type. I.e `JSONReader` for the `JSONType`...
    - Added the private `getFileType` function.
    - Added the private `readRawObject` function.
    - Added the private `readRawList` function.
    - Added the private `readerListDispatch` function.
- `Docs`
    - Added the `read.md` english documentation.

`Changed`
- `fayl`
    - The public `ReadFile` function has been renamed to `Read`.
    - The public `ReadJSONList` function has been renamed to `ReadList`.
    - The public `ReadJSONList` function has been renamed to `ReadObject`.
    - Each of the renamed function have been renamed due to the implementation of each function no longer relying on specifically the file being a `JSON` file.
## v1.3.1 - 2024-12-25 🎅
`Changed`
- `fayl`
    - The public `ReadFile` function now returns an `Exception` along with a slice of bytes.
    - The public `ReadJSONList` function now returns an `Exception` along with a `View`.
    - The public `ReadJSONObject` function now returns an `Exception` along with a type pointer.

`Removed`
- `Tests`
    - `Set`
        - Removed `Remove` tests. LOL.
## v1.3.0 - 2024-12-19
`Added`
- `Array`
    - Added `Update` method.
- `Collection`
    - Added `Update` method.
- `Queue`
    - Added `Update` method.
- `Set`
    - Added `Update` method.
- `Stack`
    - Added `Update` method.
- `Tests`
    - `Array`
        - Added `Update` tests.
    - `Map`
        - Added `Update` tests.
    - `Queue`
        - Added `Update` tests.
    - `Set`
        - Added `Update` tests.
    - `Stack`
        - Added `Update` tests.
- `Docs`
    - Added mentions of updating to `array.md`, `queue.md`, and `stack.md`.
## v1.2.0 - 2024-12-18
`Added`
- Added `TODO.md` to `gitignore`.
- `Mapping`
    - Added `Update` method.
- `Map`
    - Added `Update` method.
- `Docs`
    - Added the `Array` english documentation.
    - Added the `Map` english documentation.
    - Added the `Pair` english documentation.
    - Added the `Queue` english documentation.
    - Added the `Set` english documentation.
    - Added the `Stack` english documentation.
    - Added the `Exception` english documentation.
- `Tests`
    - Added `Array` tests.
    - Added `Iter` tests.
    - Added `Map` tests.
    - Added `Queue` tests.
    - Added `Set` tests.
    - Added `Stack` tests.

`Changed`
- The run command now builds all files in `build.yml`.
- The test command now restricts its tests to everything in the `tests` folder.
## v1.1.1 - 2024-12-18
`Added`
- Added the `-v` flag to the test command in `test.yml`
- `Tests`
    - Added `Iter` tests.
    - Added `Version` tests.

`Changed`
- `Array`
    - The `At` and `Remove` methods now checks if the array is empty. If the array is empty, an `IndexOutOfRangeError` is returned.
- `Map`
    - The `At` and `Remove` methods now checks if the map is empty. If the map is empty, a `KeyError` is returned.
- `Queue`
    - The `At` method now checks if the queue is empty. If the queue is empty, an `IndexOutOfRangeError` is returned.
- `Set`
    - The documentation comments for the `Append` and `Extend` methods have been extended.
    - The `Remove` method now checks if the set is empty. If the set is empty, an `IndexOutOfRangeError` is returned.
    - The `Discard` method now checks if the set is empty. If the set is empty, the method returns with no exception without modifying the set.
- `Stack`
    - The `At` method now checks if the stack is empty. If the stack is empty, an `IndexOutOfRangeError` is returned.
## v1.1.0 - 2024-12-13
`Added`
- `fayl`
    - Added `ReadFile` function.
    - Added `ReadJSONList` function.
    - Added `ReadJSONObject` function.
    - Added private `sliceToCollection` function.
## v1.0.0 - 2024-12-05
`Added`
- Added the `View` interface.
- `Array`
    - Added `Collect` method.
- `Collection`
    - Added `View` interface.
- `Queue`
    - Added `Collect` method.
- `Set`
    - Added `Collect` method.
    - Added `At` method.
- `Stack`
    - Added `Collect` method.
- `Logger`
    - Added documentation comments.
- `Version`
    - Added documentation comments.
- `Docs`
    - Added english documentation.
`Changed`
- `Array`
    - The `Extend` method now takes a `View` replacing a slice.
- `Collection`
    - The `Extend` method now takes a `View` replacing a slice.
- `Iter`
    - The `Enumerate` method now takes a `View` replacing a slice.
    - The `In` method now takes a `View` replacing a slice.
- `Queue`
    - The `Extend` method now takes a `View` replacing a slice.
- `Set`
    - The set now uses the internal `Mapping` interface replacing a built-in map.
    - The `Extend` method now takes a `View` replacing a slice.
    - The `Items` method now returns a pointer to the keys of the mapping instead of returning a pointer to the internal mapping in its entirety. This is more aligned with the expected behaviour of a set.
- `Stack`
    - The `Extend` method now takes a `View` replacing a slice.

`Removed`
- `Set`
    - The `ToSlice` method has been removed.
## v0.0.9 - 2024-12-05
`Added`
- Added `Sized` interface.
- `Array`
    - Added documentation comments.
- `Collection`
    - Added `Sized` interface.
    - Added documentation comments.
- `Iter`
    - Added documentation comments.
- `Mapping`
    - Added `Sized` interface.
    - Added documentation comments.
- `Map`
    - Added documentation comments.
- `Queue`
    - Added documentation comments.
- `Set`
    - Added `IsEmpty` method.
    - Added documentation comments.
- `Stack`
    - Added documentation comments.
- `Exception`
    - Added documentation comments.
- `Tests`
    - Added `Pair` tests.
- `Docs`
    - Added `Go Reference` link to `README.md`.
## v0.0.8 - 2024-12-04
`Added`
- Added `Pair` struct.
- Added `build.yml`.
- Added `test.yml`.
- `Tests`
    - `Version`
        - Added `TestNewVersionZerosOut` test case.
- `Docs`
    - Added `Build` badge link to `README.md`.
    - Added `Test` badge link to `README.md`.
    - Added english documentation link to `README.md`.
    - Added the `README.md` english documentation.
## v0.0.7 - 2024-11-19
`Changed`
- `Logger`
    - The `Log` method now uses the private `publishMessage` function replacing a `bufio.Writer`.
    - The private `buildMessage` function has been replaced by the private `publishMessage` function.
## v0.0.6 - 2024-11-17
`Added`
- `Iter`
    - Added `In` function.

`Changed`
- `Logger`
    - The global constants `AVAILABLE_OUTPUTS` and `TIMESTAMP_FORMAT`, as well as the global variable `OUTPUT_COUNT` are now private.
    - The `Write` function has been renamed to `Log`.
## v0.0.5 - 2024-11-15
`Added`
- Added `Array` struct.
- Added `Collection` interface.
- Added `Enumerate` function.
- Added `Map` struct.
- Added `Mapping` interface.
- Added `Queue` struct.
- Added `Set` struct.
- Added `Stack` struct.
## v0.0.4 - 2024-11-14
`Added`
- `Logger`
    - Added `Close` method.
## v0.0.3 - 2024-11-14
`Added`
- Added `Logger` struct.
- Added `LoggingLevel` enum.
## v0.0.2 - 2024-11-13
`Removed`
- Removed `release.yaml`.
## v0.0.1 - 2024-11-12
`Added`
- Added `Version` struct.
- Added `Exception` struct.
- Added `Collection` interface.
- Added `Array` struct.
- Added `Mapping` interface.
- Added `Map` struct.
- Added `Set` struct.
- Added `Queue` struct.
- Added `Stack` struct.