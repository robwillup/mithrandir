# More Good Practices for Functions in Go

* Short-circuiting is good
  * Check for errors early
  * If there is an error, return from the function as soon as possible
  * When you have a function that returns an error, DO NOT just ignore it. Go ahead and handle that error properly.

## Variadic Functions

Variadic functions allow you to pass in multiple parameters of the same type:

```go
func sum(values ...float64) float64 {
    total := 0.0
    for _, value := range values {
        total += value
    }

    return total
}
```

It's pretty that you can pass multiple parameters into that function:

```go
total := sum(1, 2, 100, 78)
```

And you can also spread a slice and pass it into the function:

```go
values := []int{1, 2, 100 ,78}

total := sum(values...)
```

> NOTE: The variadic parameter must be the last parameter to a function.

## Private and Public functions

When a function in a package is capitalized that function is public:

```go
// ari/expressions.go

func Add(p1, p2 int) int {
    //...
}

func double(p1 int) int {
    // ...
}

// main.go

value := ari.Add(2, 2) // works

d := ari.double(2) // fails
```

## Named returned values

This allows you to simply assign the values to the return variables and have an
empty `return` at the end of the function:

```go
func Divide(p1, p2 int) (answer float64, err error) {
    if p2 == 0 {
        err = errors.New("Cannot divide by zero")
    }

    answer = p1 / p2
    return
}
```
