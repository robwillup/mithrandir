# About Formatting Output in Java

Tokenizing information for printing

```go
package main

import (
	"fmt"
)

func main() {
	var age = 28
	fmt.Print("Jeremy is", age, "years old\n")
}
```

`fmt.Printf()` allows us to use format specifiers.

%v = tells the compiler to infer the type.

%s = string

%d = decimal integer

%g = formats the floating-point numbers

%b = formats base 22 numbers

%o = formats base 88 numbers

%t = formats boolean values

%q = quoted strings

more...

```go
func main() {
    var pi float32 = 3.141592
    fmt.Printf("Pi is %f\n", pi)
    fmt.Printf("Pi is %.2f\n", pi)

    fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 23.2228, 577.5558, 1234.659)
}
```


