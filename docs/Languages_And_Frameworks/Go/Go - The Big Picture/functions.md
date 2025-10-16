# Functions

## Signature

- All functions start with the `func` keyword.
- After that we typically set a name for it.
- Then we have the parameter list `(parameters)`
- And then we have the `(return values)`
- Finally we have the function body enclosed in curly brances

```Go
func greet(name, name2 string) {
    fmt.Println(name)
    fmt.Println(name2)
}
```

### Variadic Parameters

```Go
func greet(names ...string) {
    for _, n := range names {
        fmt.Println(n)
    }
}
```

- Received by the function as a `slice`
- Must be final parameter

### Passing Values and Pointers

```Go
func main() {
    name, otherName := "Name", "Other name"
    fmt.Println(name)        // prints "Name"
    fmt.Println(otherName)   // prints "Other name"
    myFunc(name, &otherName)
    fmt.Println(name)        // still prints "Name"
    fmt.Println(otherName)   // prints "Other new name"
}

func myFunc(name string, otherName *string) {
    name = "New name"
    *otherName = "Other new name"
}
```

> Use pointer to share memory, otherwise use values.

Here's more on this:
    As you use pointers you introduce the possibility of having problems when your program starts running concurrently.
    Use values whenever you can.

