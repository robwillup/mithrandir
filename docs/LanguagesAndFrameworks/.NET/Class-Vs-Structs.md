# Class vs. Structs

* Classes are reference types, they are allocated in the heap and are garbage collected.
* Structs are value types, they are allocated in the stack or inline in containing types,
and they are deallocated as the stack unwinds.
* In general, allocation and deallocation or value types are cheaper.

