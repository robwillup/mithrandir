## Named return values

Reference: Official Go documentation

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as "naked" return.

Naked return statements should be used only in short functions. They can harm readability in longer functions.

## If with a short statement

Like `for`, the `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the `if`.

```go
if v := math.Pow(x, n); v < lim {
    return v
}
```

## Pointers

Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

```go
var p *int
```

The `&` operator generates a pointer to its operand.

```go
i := 42
p = &i
```

The `*` operator denotes the pointer's underlying value.

```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as "dereferencing" and "indirecting".

Unlike C, go has no pointer arithmetic.

## Structs

A `struct` is a collection of fields.

Struct fields are accessed using a dot.

## Pointers to structs

Struct fields can be accessed through a struct pointer.

To access the filed `x` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.

## Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-size, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```go
a[low : high]
```

This selects a half-open range which includes the first element, but exclude the last one.

The following expression creates a slice which includes elements 1 through 3 of `a:`

```go
a[1:4]
```

## Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

## Slice length and capacity

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.

You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

## Nil slices

The zero value of a slice is `nil`.

A nil slice has a length and capacity of 0 and has no underlying array.

## Creating a slice with make

Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.

The `make` function allocates a zeroed array and returns a slice that refers to that array:

```go
a := make([]int, 5) // len(a)=5
```

## Maps

A map maps keys to values.

The zero value of a map is `nil`.

The `make` function returns a map of the given type, initialized and ready for use.

```go
func main() {
    age := make(map[string]int)
    age["Rob"] = 28
    fmt.Println(age["Rob"])
    // 28
}
```

To delete an element from a map use

```go
delete(m, key)

// m is a map
```

Test that a key is present with a two-value assignment:

```go
elem, ok = m[key]
// or
elem, ok := m[key]
```

If `key` is in `m`, `ok` is true. If not, `ok` is false.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

## Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

In the example below, the `Abs` method has a receiver of type `Vertex` named `v`.

```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
```

## Method example

```go
package main

type Test struct {
    Name string
}

func (t *Test) getName() (name string) {
    name = t.Name
    return
}

func main() {
    t := Test{}
    t.Name = "Rob"
    print(t.getName())
}
```

## Choosing a value or pointer

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the
receiver is a large struct, for example.

## Interfaces

An *interface type* is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

### Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is not explicit declaration of 
intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could
then appear in any package without prearrangement.

### An example of using interface

```go
package main

import "fmt"

type I interface {
    Add() int
}

type Fruits struct {
    Apples  int
    Oranges int
}

type Objects struct {
    Cars int
    Houses int
}

func (f Fruits) Add() int {
    return f.Apples + f.Oranges
}

func (o *Objects) Add() int {
    return o.Cars + o.Houses
}

func main() {
    var inter I
    f := Fruits{20,25}
    o := Objects{3,2}
    inter = f
    fmt.Printf("%d\n", inter.Add())
    inter = &o
    fmt.Printf("%d\n", inter.Add())    
}
```

### Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concret type:

    (value, type)
    
An interface value holds a value of a specific underlying concret type.

Calling a method on an interface value executes the method of the same name on its unerlying type.

### The empty interface

The interface type that specifies zero methods is known as the *empty interface*:

`interface{}`

An empty interace may hold values of any type.

Empty interfaces are used by code that handles values of unknown type.

## Type assertions

