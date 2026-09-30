# Diagrams
## Collections
```mermaid
classDiagram
class Collection~Type~ {
    <<interface>>
    +Append(Type Item)*
    +Extend(View~Type~ items)*
    +At(Size index) ~*Type, *Exception~*
    +Update(Size index, Type value) *Exception*
    +Remove(Size index) *Exception*
    +Items() *[]Type*
    +View~Type~$
}

class Mapping~Key, Value~ {
    <<interface>>
    +Insert(Key key, Value value) *Exception*
    +At(Key key) ~*Value, *Exception~*
    +Update(Key key, Value value) *Exception*
    +Keys() []Key*
    +Values() []Value*
    +Remove(Key key) *Exception*
    +HasKey(Key key) bool*
    +View~Pair~Key, Value~~$
}

class SafeCollection~Type~ {
    <<interface>>
    +Lockable$
    +Unlockable$
    +Collection~Type~$
}

class SafeMapping~Key, Value~ {
    <<interface>>
    +Lockable
    +Unlockable
    +Mapping~Key, Value~$
}

class Lockable {
    <<interface>>
    Lock()
    RLock()
}

class Unlockable {
    <<interface>>
    Unlock()
    RUnlock()
}

class View~Type~ {
    <<interface>>
    +Collect() []Type*
    +Sized$
}

class Wrapper~Type~ {
    <<interface>>
    +Into$
    +From$
    +View~Type~$
}

class From~Type~ {
    <<interface>>
    +From(View~Type~)*
}

class Into~Type~ {
    <<interface>>
    +Into() Collection~Type~*
}

class Sized {
    <<interface>>
    +Size() Size*
    +IsEmpty() bool*
}

class Iterable~Type~ {
    <<interface>>
    +Iterator() *Iterator~Type~*
}

class Exception {
    -ExceptionName name
    -String repr
    -String message
    -assignRepr()
    -assignName(ExceptionName name)
    -assignMessage(String message)
    +Name() ExceptionName
    +Message() String
    +Error() String
    +Is(ExceptionName name) bool
}

class ExceptionName {
    <<enumeration>>
    +BaseException
    +ArithmeticError
    +OverflowError
    +UnderflowError
    +ZeroDivisionError
    +AssertionError
    +EOFError
    +LookupError
    +OutOfRangeError
    +IndexError
    +KeyError
    +OSError
    +IOError
    +BlockingIOError
    +ChildProcessError
    +ConnectionError
    +BrokenPipeError
    +ConnectionAbortedError
    +ConnectionRefusedError
    +ConnectionResetError
    +FileExistsError
    +FileNotFoundError
    +IsADirectoryError
    +PermissionError
    +ProcessLookupError
    +TimeoutError
    +RuntimeError
    +NotImplementedError
    +ValueError
    +UnreachableError
}

class Pair~First, Second~ {
    -First first
    -Second second
    +First() First
    +Second() Second
    +SetFirst(First first)
    +SetSecond(Second second)
    +Set(First first, Second second)
    +Swap(Pair~First, Second~)
    +Flip() *Pair~Second, First~
    +Items() ~*First, *Second~
}

class SafePair~First, Second~ {
    -RWMutex firstLock
    -First first
    -RWMutex secondLock
    -Second second
    +First() *First
    +Second() *Second
    +SetFirst(First first)
    +SetSecond(Second second)
    +Set(First first, Second second)
    +Swap(Pair~First, Second~)
    +Flip() *Pair~Second, First~
    +Items() ~*First, *Second~
    +Lock()
    +RLock()
    +Unlock()
    +RUnlock()
}

class Iterator~Value~ {
    +Map(func(Value)Value callback) *Iterator~Value~
    +Filter(func(Value)bool predicate) *Iterator~Value~
    +ForEach(func(Value) callback)
    +Collect() []Value
    +Size() Size
    +IsEmpty() bool
}

class Array~Type~ {
    -[]Type items
    -Size size
    +Append(Type Item)
    +Extend(View~Type~ items)
    +At(Size index) ~*Type, *Exception~
    +Update(Size index, Type value) *Exception
    +Remove(Size index) *Exception
    +Items() *[]Type
    +Collect() []Type
    +Size() Size
    +IsEmpty() bool
}

class Map~Key, Value~ {
    -map~Key, Value~ items
    -Size size
    +Insert(Key key, Value value) *Exception
    +At(Key key) ~*Value, *Exception~
    +Update(Key key, Value value) *Exception
    +Keys() []Key
    +Values() []Value
    +Remove(Key key) *Exception
    +HasKey(Key key) bool
    +Collect() []Pair~Key, Value~
    +Size() Size
    +IsEmpty() bool
}

class Queue~Type~ {
    -[]Type items
    -Size size
    +Append(Type Item)
    +Extend(View~Type~ items)
    +At(Size index) ~*Type, *Exception~
    +Update(Size index, Type value) *Exception
    +Remove(Size index) *Exception
    +Dequeue() ~*Type, *Exception~
    +Peek() ~*Type, *Exception~
    +Items() *[]Type
    +Collect() []Type
    +Size() Size
    +IsEmpty() bool
}

class Stack~Type~ {
    -[]Type items
    -Size size
    +Append(Type Item)
    +Extend(View~Type~ items)
    +At(Size index) ~*Type, *Exception~
    +Update(Size index, Type value) *Exception
    +Remove(Size index) *Exception
    +Pop() ~*Type, *Exception~
    *Peek() ~*Type, *Exception~
    +Items() *[]Type
    +Collect() []Type
    +Size() Size
    +IsEmpty() bool
}

class SafeArray~Type~ {
    -RWMutex itemLock
    -[]Type items
    -RWMutex sizeLock
    -Size size
    +Append(Type Item)
    +Extend(View~Type~ items)
    +At(Size index) ~*Type, *Exception~
    +Update(Size index, Type value) *Exception
    +Remove(Size index) *Exception
    +Items() *[]Type
    +Collect() []Type
    +Size() Size
    +IsEmpty() bool
    +Lock()
    +RLock()
    +Unlock()
    +RUnlock()
}

class SafeMap~Key, Value~ {
    -RWMutex itemLock
    -map~Key, Value~ items
    -RWMutex sizeLock
    -Size size
    +Insert(Key key, Value value) *Exception
    +At(Key key) ~*Value, *Exception~
    +Update(Key key, Value value) *Exception
    +Keys() []Key
    +Values() []Value
    +Remove(Key key) *Exception
    +HasKey(Key key) bool
    +Collect() []SafePair~Key, Value~
    +Size() Size
    +IsEmpty() bool
    +Lock()
    +RLock()
    +Unlock()
    +RUnlock()
}

class SafeQueue~Type~ {
    -RWMutex itemLock
    -[]Type items
    -RWMutex sizeLock
    -Size size
    +Append(Type Item)
    +Extend(View~Type~ items)
    +At(Size index) ~*Type, *Exception~
    +Update(Size index, Type value) *Exception
    +Remove(Size index) *Exception
    +Dequeue() ~*Type, *Exception~
    +Peek() ~*Type, *Exception~
    +Items() *[]Type
    +Collect() []Type
    +Size() Size
    +IsEmpty() bool
    +Lock()
    +RLock()
    +Unlock()
    +RUnlock()
}

class SafeStack~Type~ {
    -RWMutex itemLock
    -[]Type items
    -RWMutex sizeLock
    -Size size
    +Append(Type Item)
    +Extend(View~Type~ items)
    +At(Size index) ~*Type, *Exception~
    +Update(Size index, Type value) *Exception
    +Remove(Size index) *Exception
    +Pop() ~*Type, *Exception~
    *Peek() ~*Type, *Exception~
    +Items() *[]Type
    +Collect() []Type
    +Size() Size
    +IsEmpty() bool
    +Lock()
    +RLock()
    +Unlock()
    +RUnlock()
}

class Set~Type~ {
    -Mapping~Type, None~ items
    +Append(Type Item)
    +Extend(View~Type~ items)
    +At(Size index) ~*Type, *Exception~
    +Update(Size index, Type value) *Exception
    +Remove(Size index) *Exception
    +Discard(Type item) ~*Type, *Exception~
    +Contains(Type item) bool
    +Difference(Set~Type~ other) *Set~Type~
    +Intersection(Set~Type~ other) *Set~Type~
    +Into() Collection~Type~
    +From(View~Type~)
    +String() String
    +Items() *[]Type
    +Collect() []Type
    +Size() Size
    +IsEmpty() bool
}

Collection o-- View : Aggregate
Collection o-- Wrapper : Aggregate
View o-- Sized : Aggregate

Collection <|.. Array : Implements
Collection <|.. Queue : Implements
Collection <|.. Stack : Implements
Collection <|.. Set : Implements
Collection <|.. SafeArray : Implements
Collection <|.. SafeQueue : Implements
Collection <|.. SafeStack : Implements

Iterable <|.. Array : Implements
Iterable <|.. Map : Implements
Iterable <|.. Queue : Implements
Iterable <|.. Stack : Implements
Iterable <|.. Set : Implements
Iterable <|.. SafeArray : Implements
Iterable <|.. SafeMap : Implements
Iterable <|.. SafeQueue : Implements
Iterable <|.. SafeStack : Implements

SafeCollection <|.. SafeArray : Implements
SafeCollection <|.. SafeQueue : Implements
SafeCollection <|.. SafeStack : Implements

Lockable <|.. SafeArray : Implements
Lockable <|.. SafeMap : Implements
Lockable <|.. SafeQueue : Implements
Lockable <|.. SafeStack : Implements

Unlockable <|.. SafeArray : Implements
Unlockable <|.. SafeMap : Implements
Unlockable <|.. SafeQueue : Implements
Unlockable <|.. SafeStack : Implements

View <|.. Array : Implements
View <|.. Map : Implements
View <|.. Queue : Implements
View <|.. Stack : Implements
View <|.. Set : Implements
View <|.. SafeArray : Implements
View <|.. SafeMap : Implements
View <|.. SafeQueue : Implements
View <|.. SafeStack : Implements
View <|.. Iterator : Implements

Wrapper <|.. Set : Implements
Wrapper <.. View : Dependency
Wrapper <.. Collection : Dependency

From o-- Wrapper : Aggregate
From <|.. Set : Implements

Into o-- Wrapper : Aggregate
Into <|.. Set : Implements

Sized <|.. Array : Implements
Sized <|.. Map : Implements
Sized <|.. Queue : Implements
Sized <|.. Stack : Implements
Sized <|.. Set : Implements
Sized <|.. SafeArray : Implements
Sized <|.. SafeMap : Implements
Sized <|.. SafeQueue : Implements
Sized <|.. SafeStack : Implements

Mapping <|.. Map : Implements
Mapping <|.. SafeMap : Implements
SafeMapping <|.. SafeMap : Implements
Mapping ..* Set : Composite
Mapping <.. Pair : Dependency
Map <.. Pair : Dependency
SafeMap <.. SafePair : Dependency

Set <.. Collection : Dependency
Set <.. View : Dependency

Collection <.. Exception : Dependency
Mapping <.. Exception : Dependency

Exception *-- ExceptionName : Composite
```