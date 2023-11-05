# Go Functions

As we begin, keep in mind that, as in other programming languages, functions in Go
should be as clear as possible.

What does a function in Go look like?

```go
package main

import "fmt"

func main() {
    fmt.Printf("%d\n", double(2)) // prints 4
}

func double(p1 int) int {
    return p1 * 2
}
```

The first thing in a function is the keyword `func`. After comes the name of the
function which you need to choose, for example, `add` or `listFiles`, etc. Right
after the name there is a pair of `()`. If your function takes parameters, you will
add them between the parenthesis. First the name of the parameter then the type, e.g.: `(parameter1 float64)`. If you have more then one parameter of the same type `(p1 float64, p2 float64)` you can used this simplified syntax `(p1, p2 float64)`.
After the closing parenthesis, if your function returns a value, that's where you
specify the type, for example: `func double(number int) int` and it's possible to
return more than one value, `func doubleAndHalf(number int) (int, float64)`.
After this closing parenthesis, you will add a opening curly bracket `{` and that
must be in the same line. In the following lines you add the function body. Add the
end, on a new line you add the closing bracket:

```go
func doubleAndHalf(p1 int) (int, float64) {
    return p1 * 2, float64(p1) / float64(2)
}
```

## More Good Practives for Functions in Go

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
