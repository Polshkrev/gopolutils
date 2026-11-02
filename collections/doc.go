/*
Collections provide interfaces of standardisation for operations related to data structures.
Standardizations defined within the collections package include:

Interfaces:

  - [Collection]: An interface for linear data structures.
  - [Mapping]: An interface for key-value pairs.
  - [View]: Interface for read-only data access.
  - [Sized]: Interface defining a data structure with a size.
  - [Wrapper]: Interface defining a data strucuture that can be converted and extracted from another.
  - [Iterable]: Interface to define a [Collection] from which an [Iterator] can be obtained.
  - [From]: Interface defining a standardisation of a [Wrapper] from a given [View].
  - [Into]: Interface defining a standardisation of a [Wrapper] into a [View].

Implementations:
  - [Array]: Standardisation of a contiguous dynamic array.
  - [Map]: Standardisation of a key-value mapping.
  - [Queue]: Standardisation of a FIFO data structure.
  - [Stack]: Standardisation of a FILO data structure.
  - [Set]: Standardisation of a java-like hash-set.
  - [LinkedList]: Standardisation of a singly-linked list.
  - [DoublyLinkedList]: Standardisation of a doubly-linked list.
  - [Deque]: Standardisation of a double-ended queue.

Each of the implementations contain a concurrent-safe variant.

Collections also defines a few function iterators:

  - [Enumerate]: A python-like enumeration over a [View].
  - [In]: A function to determine if a given element is contained within a [View].
  - [Reverse]: Iterate over a [View] in reverse order.
  - [Chain]: Chain together a variadic [View] list.
*/
package collections
