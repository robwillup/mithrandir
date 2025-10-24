# Setting Up a Development Environment

**Keep going to https://go.dev**

For tips on how to install and upgrade Go: https://go.dev/doc/install

---

How do IDEs work with Go?

VS Code needs an extension to become aware of the Go language. The extension in turn sends request to libraries
so that it can provide features like auto completion, etc.

VS Code <--> Extension <--> Libraries <--> Go

---

Go is a package-based language.
Every project needs to have a package.

The package can either be `main` or some other word.

* main: executable source code
* other words: library code

---

## Simple Data Types

* Strings
* Numbers
* Booleans
* Errors

### Strings

Series of characters that are linked together.

```go
"this is a string"      // interpreted string - interprets escapes

`this is also a string` // raw string - like a literal string in C#

// Good to note that if you want to break lines in a raw string you can do it with

`this is a line
and this is another` // Go will encode that new line into the raw string.
```

### Numbers

- Integers: int, uint
- Floating point numbers: float32, float64
- Complex numbers: complex64, complex128

### Boolean

In Go you won't find the concept of `falsy` and `truthy` values. Go does not aliases other values such as
0 as false and 1 as true. If you need to work with true and false in Go, Boolean is all you need to worry about.

### Error

The error built-in interface type is the conventional interface for representing an error condition, with the nil value
representing no error.

```go
type error interface {
    Error() string
}
```

## Variables

How to declare them

```go
var myName sring          // declare variable - the zero for string is the empty string ""
var myName string = "Rob" // declare and initialize
var myName = "Rob"        // initialize with inferred type
myName := "Rob"           // short declaration syntax
```

## Constants

```go
const a = 42             // implicitly typed - can be interpreted as different number types int, float32, etc.
const b string = "hello" // explicitly type constants
const c = a              // one constant can be assigned to another
const (
    d = true
    e = 3.14
)

const c = 2 * 5             // constant expression
const d = "hello, " + "Rob" // must be calculable at compile time
const e = someFunction()    // won't work - can't be evaluated at compile time
```
