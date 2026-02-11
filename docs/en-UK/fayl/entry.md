# Entry
A standardization of an entry on the filesystem. This type is designed to aid in operations pretaining to a concrete file rather than just a path.
## Table of Contents
1. [Usage](#usage)
    1. [Construction](#construction)
    2. [Functionality](#functionality)
### Usage
### Construction
To construct a new entry, the `NewEntry` constructor is defined that takes in a [Path](/docs/en-UK/fayl/path.md), and returns a pointer to a new `Entry` whose type is set to `File` and whose content is an empty initialized slice.
### Functionality
Each of the entries properties have a self-explanatory getter and setter method defined. To create a file there are three methods defined: 
- `Touch`: To create an entry of type `File`.
- `MakeDirectory`: To create an entry of type `Directory`.
- `Create`: A generic wrapper over the previous aforementioned methods that dispatched based on the internal `EntryType` of the entry.


Each of the aforemnetioned methods can fail, if the entry path already exists, or if the entry is not of the corresponding type.

To remove an entry, there are a very similar set of methods defined:
- `RemoveFile`: To remove an entry of type `File`.
- `RemoveDirectory`: To create an entry of type `Directory`.
- `Remove`: A generic wrapper over the previous aforementioned methods that dispatched based on the internal `EntryType` of the entry.


Each of the aforemnetioned methods can fail, if the entry path does not exist, if the entry is not of the corresponding type, or if the entry's type can not be determined.

Each of the creation methods and removal methods are concurrently implemented and utilize concurrent-safe collections.