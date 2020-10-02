# The Array Type

You can have a collection of multiple values using an *array*. Every element of an array must have the same type and in Rust arrays have a fixed length.

In Rust, use a comma-separated list inside square brackets to declare and initialize an array:

```rust
fn main() {
    let a = [1,2,3,4,5,6,7];
}
```

Arrays are useful when you want your data allocated in the stack rather than the heap or when you want to ensure you always have a fixed number of elements. An array isn't as flexible as the vector type, though. A vector is a similar collection type provided by the standard library that is allowed to grow or shrink in size. If you are unsure whether to use an array or a vector you should probably use a vector.

To define the type and size of an array, use this notation:

```rust
let a: [i32, 5] = [1,2,3,4,5];
```

If you want all the elements of an array to be the same you can write:

```rust
let a: [3,5];
// equals
let a = [3,3,3,3,3];
```

## Accessing Array Elements

An array is a single chunk of memory allocated on the stack. You can access elements of an array using indexing, like this:

```rust
fn main() {
    let a = [1,2,3];
    let first = a[0];
    let second = a[1];
}
```

If you try to access an element outside the bounds of the array, the program will compile but there will be a run time panic.

This is the first example of Rust's safety principles in action. In many low-level languages, this kind of check is not done, and when you provide an incorrect index, invalid memory can be accessed. Rust protects you against this kind of error by immediately exiting instead of allowing the memory access and continuing.