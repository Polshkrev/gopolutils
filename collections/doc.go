/*
Collections provide interfaces of standardization for operations related to data structures.
Standardizations defined within the collections package include:

Interfaces:

  - [Collection]: An interface for linear data structures.
  - [Mapping]: An interface for key-value pairs.
  - [View]: Interface for read-only data access.
  - [Sized]: Interface defining a data structure with a size.
  - [Wrapper]: Interface defining a data strucuture that can be converted and extracted from another.

Implementations:
  - [Array]
  - [Map]
  - [Queue]
  - [Stack]
  - [Set]

Each of the implementations contain a concurrent-safe variant.

Collections also defines a few function iterators:

  - [Enumerate]: A python-like enumeration over a [View].
  - [In]: A function to determine if a given element is contained within a [View].
  - [Reverse]: Iterate over a [View] in reverse order.
  - [Chain]: Chain together a variadic [View] list.
*/
package collections
