# Path

A go implementation of [Python's Pathlib](https://docs.python.org/3/library/pathlib.html).

## Table Of Contents
1. [Usage](#usage)
    1. [Construction](#construction)
        1. [NewPath](#newpath)
        2. [PathFrom](#pathfrom)
        3. [PathFromParts](#pathfromparts)
    2. [Modification](#modification)
## Usage
### Construction
To construct a new path object, there are three functions you can use: `NewPath`, `PathFrom`, and `PathFromParts`. Some of the path methods will return a new path as well, such as `Absolute`, `Parent`, and `Root`. 
#### NewPath
Using `NewPath` will construct a path from the current working directory. If the current working directory can not be obtained, an `OSError` is printed to standard error and the programme exits.
#### PathFrom
Using `PathFrom` will construct a new path from a given path string. This constructor does not fail.
#### PathFromParts
Using `PathFromParts` will construct a path from a given folder name, file name, and file suffix. As an example:
```go
var path *gopolutils.Path = gopolutils.PathFromParts(".", "test", "txt")
// The path string should equal: "{currentDirectory}/test.txt"
```
This constructor uses an operating system dependent path seperator.

### Modification
There are two straightforward ways of modifying the path such as `Append` and `AppendAs`. Using `Append` will append a path object. Using `AppendAs` will append a string.