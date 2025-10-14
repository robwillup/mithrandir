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

> You did not use a single function for all email operations.

3. Your Go script contains an `if / else if / else` chain with 10 `else if` statements. Each `else if` checks two
conditions combined using the `&&` operator. What minimum number of individual conditions must be evaluated as false
for the final `else` block to run?

> Both conditions from each `else if` statement.

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

6. You declare `price := 19.99`. Later, you try `price = 20` and this compiles without an error. Why?

> Go converts the integer to float64 automatically.
