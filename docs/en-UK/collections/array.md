# Array

Implementation of a classical dynamic array

## Table Of Contents
1. [Append](#append)
2. [Remove](#remove)
3. [Access](#access)
4. [Implements](#implements)

### Append
To append to an array, there are two methods that can be used; there is &mdash; of course &mdash; the `Append` method. This method appends a singular item to the array. There is also the `Extend` method. This method appends multiple items to the array in the form of a `View`.

### Remove
To remove from an array, there is a simple `Remove` method. This method takes in an index where the value to be removed is located. If the array is empty, or if the given index is greater than the size of the array, then an `IndexOutOfRangeError` is returned.

### Access
To access an item in the array, the `At` method is available. This method will take a given index. This method will either return a pointer to the item stored at the given index, and a nil exception pointer; or this method will return a nil data pointer and an exception. To update the array, an `Update` method is available. This method takes in an index where the item to update is located and a new value to assign at the given index. This method will either update the value value at the given index and return a nil exception pointer or this method will return a pointer to an exception.

### Implements
This structure implements the `Collection`, `View`, and `Sized` interfaces.