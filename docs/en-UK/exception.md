# Exception
A wrapper around the `error` type in golang that is based on [Python's exception](https://docs.python.org/3/tutorial/errors.html).
This is to be used as a base class to more easily define custom exceptions.

## Table Of Contents
1. [Usage](#usage)

## Usage

The exception module should be used as a return value. There are many examples used in the collections package such as the `Collection[Type].At(index uint64) (*Type, *gopolutils.Exception)` method. As an exxample, the array collection has an `At` method:
```go
func (array Array[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	var outOfRange *gopolutils.Exception = gopolutils.NewNamedException("IndexOutOfRangeError", fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	if index > array.size {
		return nil, outOfRange
	}
	return &array.items[index], nil
}
```
