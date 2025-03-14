# Read
Functions to aid in the reading of file data.

## Table Of Contents
1. [ReadFile](#readfile)
2. [ReadList](#readlist)
3. [ReadObject](#readobject)

### ReadFile
For the puposes of documentation, this function is called `ReadFile`. In the source code, this function is simply called `Read`. This function reads a file from a given path parametre into a slice of bytes. If the file can not be read, an `IOError` is returned with a `nil` data pointer.

### ReadList
This function takes in a filepath &mdash; either relative or absolute &mdash; and returns a [View](/docs/en-UK/collections/view.md) representing the list of serialized data from the file. If the file can not be read, an `IOError` is returned with a `nil` data pointer. Another exception that can be returned is, if the file with the given path can not be unmarshalled, an `IOError` is returned with a `nil` data pointer.

### ReadObject
This function takes in a filepath &mdash; either relative or absolute &mdash; and returns a pointer of a type representing the serialized data from the file. If the file can not be read, an `IOError` is returned with a `nil` data pointer. Another exception that can be returned is, if the file with the given path can not be unmarshalled, an `IOError` is returned with a `nil` data pointer.