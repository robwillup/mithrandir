# Command-Line Arguments

The `os` package provides functions and other values for dealing with the operating system in a platform-independent fashion. Command-line arguments are available to a program in a variable named `Args` that is part of the `os` package; thus its name anywhere outside the `os` package is `os.Args`.

The variable `os.Args` is a slice of strings. Slices are a fundamental notion in Go, and we'll talk a lot more about them soon. For now, think of a slice as a dynamically sized sequence s of array elements where individual elements can be accessed as s[i] and a contiguous subsequence as s[m:n]. The number of elements is given by len(s). As in most other programming languages, all indexing in Go uses half-open intervals that include the first index but excludes the last, because it simplifies logic. 

The first element of `o.Args`, `os.Args[0]`, is the name of the command itself; the other elements are the arguments that were preented to the program when it started execution. A slice expression of the form s[m:n] yields a slice that refers to element m through n-1, so the elements we need for our next example are those in the slice `os.Args[1:len(os.Args)]`. If m or n is omitted, it defaults to 0 or len(s) respectively, so we can abbreviated the desired slice as `os.Args[1:]`.

>By convension, we describe each package in a comment immediately preceding its package declaration; for a `main` package, this comment is one or more complete sentences that describe the program as a whole.

```go
// This program prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

The increment statement 1++ adds 1 to i; it's equivalent to i += 1 which is in turn equivalent to i = i + 1. There's a corresponding decrement statement i-- that subtracts 1. These are statements, not expressions as they are in most languages in the C family, so j = i++ is illegal, and they are postfix only, so --i is not legal either.

The `for` loop is the only loop statement in Go. It has a number of forms, one of which is illustrated here:

```go
for initialization; condition; post {
    // zero or more statements
}
```

Parentheses are never ued around the three components of a for loop. The braces are mandatory, however, and the opening brace must be on the same line as the *post* statement.

The *initialization* statement is optional and it is executed before the loop starts. If it is in use, it has to be a *simple statement*, that is, a short variable declaration, an increment or assignment statement, or a function call. The *condition* is a boolean expression that is evaluated at the beginning of each iteration of the loop; if it evaluates to true, the statements controlled by the loop are executed. The *post* statement runs after the body of the loop, then the condition is evaluated again. The loop ends when the condition becomes false.

Any of these parts can be omitted. If there is no *initialization* and no *post*, the semi-colons may also be omitted:

```go
// a traditional "while" loop
for condition {
    // ...
}
```

If the condition is also omitted, we have an infinite loop:

```go
for {
    // this will run 'forever'
}
```

Loops of this form may be terminated in some other way, like a break or a return statement.

Another form of the for loop iterates of a range of values from a data type like a string or a slice. To illustrate, here's a second version of echo:

```go
// This program prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
```

The `_` in the range loop is known as a *blank identifier*. The blank identifier may be used whenever syntax requires a variable name but program logic does not, for instance to discard an unwanted loop index when we require only the element value.

Most Go programmers would likely use `range` and `_` to write the echo program above, since the indexing over `os.Args` is implicit, not explicit, and thus easier to get right.

Here's a simpler and more efficient solution:

```go
// This program prints its command-line arguments
package main

import (
        "fmt"
        "os"
        "strings"
)

func main() {
        fmt.Println(strings.Join(os.Args[1:], " "))
}
```

