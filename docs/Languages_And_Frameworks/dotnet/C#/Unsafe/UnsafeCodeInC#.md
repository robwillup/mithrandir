# Unsafe code in C#

There are situations where access to pointer types become necessary,
like interfacing with the underlying OS, accessing a memory-mapped device,
or implementing a time-critical algorithm.

To address this, C# provides the ability to write **unsafe code**.

In a sense, writing unsafe code is like writing C code withing a C# program.

Unsafe code is in fact a "safe" feature from the perspective of both developers
and users. Unsafe code must be clearly marked with the modifier `unsafe`, so
developers can't possibly use unsafe features acidently. And the execution
engine works to ensure that unsafe code cannot be executed in an untrusted
environment.

## Pointers

In an unsafe context, a type maybe a *pointer_type* as well as a *value_type*
or a *reference_type*.

The type specified before the `*` in a pointer type is called the
*reference type* of the pointer type. It represents the type of the variable
to which a value of the pointer type points.

Unlike references (values of reference types), pointers are not tracked by the
garbage collector -- the garbage collector has no knowledge of pointers and
the data to which the point. For this reason a pointer is not permitted to
point to a reference to a struct that contains references, and the
reference type of a pointer must be an *unmanaged_type*.

An *unmanaged_type* is a type that isn't a *reference_type* or a constructed
type and doesn't contain *reference_type* or constructed type fields at any
level of nesting. In other words, an *unmanaged_type* is one of the following:

* sbyte, byte, short, ushort, int, uint, long, ulong, char, float, double,
decimal or bool.
* Any *enum_type*
* Any *pointer_type*
* Any user-defined *struct_type* that is not a constructed type and contains
fields of *unmanaged_types only.

Some examples of pointer are given in the table below
|||