# Deque

Implementation of a classical double-ended queue using a doubly linked list.

## Table Of Contents
1. Use
2. Implements

### Use
To move an element to the front of the deque, the `Use` method is available. This method takes in an index where the element to be moved is located. The element at the given index is then unlinked from its current position and linked to the front of the deque. If the deque is empty, a `ValueError` is returned. If the given index is greater than or equal to the size of the deque, an `IndexError` is returned. If the given index is zero, no operation is performed.

### Implements
This structure implements the `Collection` interface through its embedded [DoublyLinkedList](/docs/en-UK/collections/doubly_linked_list.md).