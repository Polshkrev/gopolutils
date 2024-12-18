# Map

Implementation of a simple key-value pair mapping.

## Table Of Contents
1. [Insert](#insert)
2. [Remove](#remove)
3. [Access](#access)
4. [Implements](#implements)

### Insert
To insert a key-value pair into a map, there is &mdash; of course &mdash; the `Insert` method. This method inserts a given value mapped to a given key into a map.

### Remove
To remove from a map, there is a simple `Remove` method. This method takes in a key mapped to the value to be removed. If the map is empty, or if the given key does not exist in the map, then a `KeyError` is returned.

### Access
To access a value in the map, the `At` method is available. This method will take a given key where the value is mapped. This method will either return a pointer to the mapped value, and a nil exception pointer; or this method will return a nil data pointer and an exception. There is also an `Update` method available. This function takes the key where the value to be updated is mapped and the new value to assign to the given key.

### Implements
This structure implements the `Mapping`, and `Sized` interfaces. This structure does **not** implement the `View` interface.