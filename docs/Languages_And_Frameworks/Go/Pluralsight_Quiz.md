# Pluralsight Skill IQ Quiz

1. You have a slice with 10 elements, as shown:

```Go
arr := []int{54, 23, 12, 52, 10, 65, 45, 75, 95, 10}
```

You must display the numbers 54, 52, 45, and 10. How will you write the code?

```Go
for i, v := range arr {
    if i % 3 == 0 {
        fmt.Println(v)
    }
}
```

---

2. You create a script with 20 functions to handle all email-related operations: validating, saving, analyzing,
logging, and authentication. You upload the script to your internal company server and run it without any errors.
However, your team leader review the script after its test phase completion and suggests that the script has low
cohesion. Why?

My first answer:

> ~You did not use a single function for all email operations.~

I got this wrong, the correct answer was:

> You did not create separate scripts for related operations.

---

3. Your Go script contains an `if / else if / else` chain with 10 `else if` statements. Each `else if` checks two
conditions combined using the `&&` operator. What minimum number of individual conditions must be evaluated as false
for the final `else` block to run?

My answer:

> ~Both conditions from each `else if` statement.~

Correct answer:

> 11 - One `if` condition and at least one condition from each `else if` statement.

---

4. You are monitoring a temperature sensor that emits multiple readings per minute. You must create a function that
accepts these unknown number of readings and returns their average. How will you write the function definition?

```Go
func average(readings ...float64) float64 {
    var sum float64
    for _, reading := range readings {
        sum += reading
    }

    return sum / float64(len(readings))
}
```

5. What is the output of the following program?

```Go
items := []string{"pen", "pencil", "marker"}
newItems := append(items, "eraser")
items[0] = "crayon"
fmt.Println(newItems[0])
```

> "pen"

---

6. You create a function with multiple variable scoping, as shown:

```Go
package foo

import "fmt"

func Foo() {
	s := 20
	k := 0

	if 4 < 5 {
		k = 50
		fmt.Println(k)
	} else {
		k = 90
		fmt.Println(k)
	}
	for x:= 0; x <= 2; x++ {
		if x%2 == 0 {
			fmt.Println(x + s)
		} else {
			fmt.Println(x + s + k)
		}
	}
}

// main.go

package main

import "modulename/internal/foo"

func main() {
    foo.Foo()
}

```

You run the function, and it compiles successfully. What value will you receive on the last iteration of the `for` loop?

> 22

---

7. You declare `price := 19.99`. Later, you try `price = 20` and this compiles without an error. Why?

> Go converts the integer to float64 automatically.

---

8. You write this function:

```Go
func printValue(v interface{}) {
	if v == 42 {
		fmt.Println("int")
	} else if v == 42.0 {
		fmt.Println("float")
	}
}
```

You call it with `printValue(42.0)`, but it prints nothing. Why?

> The type of `v` is `interface{}`, so comparisons fail without type assertion.

---

9. What is the purpose of the `GOPATH` environment variable in Go?

> It defines the root directory for Go workspaces.

---

10. You create a conditional statement to check if a user provided a Boolean value. You must ensure that the input is a
Boolean data type and write the following script to achieve the task:

```Go
func main() {
	var input string
	fmt.Print("Enter value: ")
	fmt.Scanln(&input)
	if input {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
```

You run the script, but before the user can enter the value, the script fails with an error - "non-boolean condition in
if statement". What step will you take to resolve the issue.

```Go
if strings.ToLower(input) == "true" {
    fmt.Println("Pass")
} else {
    fmt.Println("Fail")
}
```

---

11. What is the output of the following code?

```Go
original := []int{1, 2, 3, 4, 5}
sub := original[1:4]
sub[0] = 99
fmt.Println(original)
```

> [1 99 3 4 5]

---

12. You use `if / else if / else` to handle a logical operation, as shown:

```Go
package main
import "fmt"
func main() {
	number := 7
	if number % 2 == 0 && true {
		fmt.Println("Foo")
	} else if !(number % 3) == 0 {
		fmt.Println("Baz")
	} else {
		fmt.Println("Que")
	}
}
```

What happens when you run this script?

> You receive an invalid operation error.



