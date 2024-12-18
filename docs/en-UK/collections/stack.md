# Stack

Implementation of a stack data structure.

## Table Of Contents
1. [Append](#append)
2. [Remove](#remove)
3. [Access](#access)
4. [Implements](#implements)

### Append
To append to a stack, there are two methods that can be used; there is &mdash; of course &mdash; the `Append` method. This method appends a singular item to the stack. There is also the `Extend` method. This method appends multiple items to the stack in the form of a `View`.

### Remove
To remove from a stack, there is the `Pop` method. A `Remove` method is defined, but always returns a `NotImplementedError` and is mainly just to adhere to the `Collection` interface. As the name would suggest, the `Pop` method returns the last appended item and removes it from the stack. If the stack is empty, an `Exception` is returned with a nil data pointer.

### Access
To access an item in the stack, there is &mdash; as previously discussed &mdash; the **destructive** `Pop` method and the **non-destructive** `Peek` method. The `Peek` method returns the last appended item, but does not remove it from the stack. If the stack is empty, an `Exception` is returned with a nil data pointer.

### Implements
This structure implements the `Collection`, `View`, and `Sized` interfaces.