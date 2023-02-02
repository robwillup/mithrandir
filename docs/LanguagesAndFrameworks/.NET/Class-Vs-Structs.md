# Class vs. Structs

* Classes are reference types, they are allocated in the heap and are garbage collected.
* Structs are value types, they are allocated in the stack or inline in containing types,
and they are deallocated as the stack unwinds.
* In general, allocation and deallocation or value types are cheaper.

Arrays:

* Arrays of reference types are allocated out-of-line, meaning the array elements are
just references to instances of the reference type residing on the head.
* Value type arrays are allocated inline, meaning that the array elements are the actual
instances of the value type.

Therefore, allocations and deallocations of value type arrays are much cheaper than
allocations and deallocations of reference type arrays. In addition, in a majority of
cases value type arrays exhibit much better locality of reference.

The next difference is related to memory usage. Value types get boxed when cast to
a reference type or one of the interfaces they implement. They get unboxed when cast
back to the value type. Because boxes are objects that are allocated on the heap and
are garbage-collected, too much boxing and unboxing can have a negative impact on the
heap, the garbage collector, and ultimately the performance of the application. In
contrast, no such boxing occurs as reference types are cast.

Next, reference type assignments copy the reference, whereas value type assignments
copy the entire value. Therefore, assignment of large reference types are cheaper than
assignments of large value types.

Since value type copies are not created explicitly by the user but are implicitly
created when arguments are passed or return values are returned, value types that can
be changed can be confusing to many users. Therefore, **value types should be immutable**.

As a rule of thumb, the majority of types in a framework should be classes. There are,
however, some situations in which the characteristics of a value type make it more
appropriate to use structs.

CONSIDER defining a struct instead of a class if instances of the type are small and
commonly short-lived or are commonly embedded in other objects.

AVOID defining a struct unless the type has all of the following characteristics:

* It logically represents a single value, similar to primitive types (int, double, etc.)
* It has an instance size under 16 bytes.
* It is immutable.
* It will not have to be boxed frequently.

In all other cases you should define your types as classes.
