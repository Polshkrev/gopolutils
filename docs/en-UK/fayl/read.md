# Read

Functions to aid in the reading of file data.

## Table Of Contents
1. [Types](#types)
    1. [Reader](#reader)
2. [ReadFile](#readfile)
3. [ReadList](#readlist)
4. [ReadObject](#readobject)

### Types
This module defines a few types to standardize reading a file.

#### Reader
The reader type is designed to define an interface standardizing how a file reader function should be typed. As a first parametre, this function takes in a slice of bytes. This represents the file content. Typically, this is obtained through the [Read](#readfile) function in this module. As a second parametre, this function takes in a pointer to either a list or an object to serialize the content of the file. This type is aligned with the [json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) function in the standard library, as the usage is the same.

There are many default readers defined in this module. As stated previously, the `json.Unmarshal` function is avaiable under the `JSONReader` type alias. Under the same paradigm, there is also a `yaml` and `toml` reader available.

### ReadFile
For the puposes of documentation, this function is called `ReadFile`. In the source code, this function is simply called `Read`. This function reads a file from a given path parametre into a slice of bytes. This function uses the absolute path of the given path. If the function can not determine the absolute path of the given path, an `IOError` is returned with a `nil` data pointer. Alternatively, if the file can not be read, an `IOError` is returned with a `nil` data pointer.

### ReadList
This function takes in a filepath &mdash; either relative or absolute &mdash; and returns a [View](/docs/en-UK/collections/view.md) representing the list of serialized data from the file. If the function can not determine the absolute path of the given path, an `IOError` is returned with a `nil` data pointer. Alternatively, if the file can not be read, an `IOError` is returned with a `nil` data pointer. Another exception that can be returned is, if the file with the given path can not be marshalled, an `IOError` is returned with a `nil` data pointer.

### ReadObject
This function takes in a filepath &mdash; either relative or absolute &mdash; and returns a pointer of a type representing the serialized data from the file. If the function can not determine the absolute path of the given path, an `IOError` is returned with a `nil` data pointer. Alternatively, if the file can not be read, an `IOError` is returned with a `nil` data pointer. Another exception that can be returned is, if the file with the given path can not be marshalled, an `IOError` is returned with a `nil` data pointer.