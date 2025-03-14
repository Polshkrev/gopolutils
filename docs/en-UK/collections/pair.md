# Pair

Implementation of a de-facto tuple like pair. Inspired by [std::pair](https://en.cppreference.com/w/cpp/utility/pair) in c++.

## Table Of Contents
1. [Access](#access)
2. [Utilities](#utilities)

### Access
This structure provides multiple access methods. One access method is provided per each of the two properties defined. These methods return pointers to the data stored in each respective property.

### Utilities
As in the original c++ implementation, this structure has a `Swap` method. This method takes a given with the same type generics as the original and swaps each of the values; first with first, second with second. A `Flip` function is also available. This method returns a new pair with the second property of the original pair as its first property and the first property of the original pair as its second property.