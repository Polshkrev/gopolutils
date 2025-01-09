# Set
Implementation of a set data structure. "Under the hood", this structure uses the [mapping](/docs/en-UK/collections/mapping.md) interface.

## Table Of Contents
1. [Append](#append)
2. [Remove](#remove)
3. [Access](#access)
4. [Utilities](#utilities)
5. [Implements](#implements)

### Append
To append to a set, there are two methods that can be used; there is &mdash; of course &mdash; the `Append` method. This method appends a singular unique item to the set. There is also the `Extend` method. This method appends multiple unique items to the set in the form of a `View`.

### Remove
To remove from a set, the `Remove` method is available. This method takes an index where the item to be removed is located within the set. If the set is empty, or the given index is greater than the size of the set, an `IndexOutOfRange` exception is returned. In a rare case that the data at the given index can not be found, an `IndexError` is returned. A `Discard` method is also available. This method is a removal method that takes in an item to remove and &mdash; instead of returning an exception or just simply crashing &mdash; this method simply does not modify the set and blankly returns if the given item is not found in the set or if the set is empty. In a **rare critical case**, the programme **will crash** if the given item **can not be removed** from the set.

### Access
Accessing a set is not allowed. There is no method defined to allow for access of an item within the set. There is an `At` method defined but will always return a `NotImplementedError`.

### Utilities
This structure has two other methods defined: `Difference` and `Intersection`. As the name suggests, the `Difference` method takes an operand set and returns a pointer to a new set containing all the items in the operand set that are not contained within the original set. The `Intersection` method is the exact opposite of the `Difference` method. The `Intersection` method takes an operand set and returns a pointer to a new set containing all the items that are contained within both sets.

### Implements
This structure implements the `Collection`, `View`, and `Sized` interfaces.