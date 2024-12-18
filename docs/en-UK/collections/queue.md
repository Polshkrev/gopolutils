# Queue

Implementation of the queue data structure.

## Table Of Contents
1. [Append](#append)
2. [Remove](#remove)
3. [Access](#access)
4. [Implements](#implements)

### Append
To append to a queue, there are two methods that can be used; there is &mdash; of course &mdash; the `Append` method. This method appends a singular item to the queue. There is also the `Extend` method. This method appends multiple items to the queue in the form of a `View`.

### Remove
To remove from a queue, there is the `Dequeue` method. A `Remove` method is defined, but always returns a `NotImplementedError` and is mainly just to adhere to the `Collection` interface. As the name would suggest, the `Dequeue` method returns the first appended item and removes it from the queue. If the queue is empty, an `Exception` is returned with a nil data pointer.

### Access
To access an item in the queue, there is &mdash; as previously discussed &mdash; the **destructive** `Dequeue` method and the **non-destructive** `Peek` method. The `Peek` method returns the first appended item, but does not remove it from the queue. If the queue is empty, an `Exception` is returned with a nil data pointer.

### Implements
This structure implements the `Collection`, `View`, and `Sized` interfaces.