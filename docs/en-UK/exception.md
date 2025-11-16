# Exception
A wrapper around the `error` type in golang that is based on [Python's exception](https://docs.python.org/3/tutorial/errors.html).
This is to be used as a base class to more easily define custom exceptions.
## Table Of Contents
1. [Usage](#usage)
	1. [Is](#is)
	2. [Must](#must)
## Usage
The exception module should be used as a return value. There are many examples used in the collections package such as the `Collection[Type].At(index uint64) (*Type, *gopolutils.Exception)` method. As an example, the array collection has an `At` method:
```go
func (array Array[Type]) At(index uint64) (*Type, *gopolutils.Exception) {
	if array.IsEmpty() {
		return nil, gopolutils.NewNamedException(gopolutils.ValueError, fmt.Sprintf("Can not access an empty array at index %d.", index))
	} else if index > array.size {
		return nil, gopolutils.NewNamedException(gopolutils.OutOfRangeError, fmt.Sprintf("Can not access array of size %d at index %d.", array.size, index))
	}
	return &array.items[index], nil
}
```
### Is
In the use case of a method returning multiple exceptions, a method is defined to determine which exception is returned. The method `Is` is defined to determine which named exception has been returned.
### Must
Within this module, a `Must` function is defined. This function accepts a tuple result of a type parametre and an `Exception`. If the passed in exception is not nil, the function panics.