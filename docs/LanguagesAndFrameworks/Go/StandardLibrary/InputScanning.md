# About Input Scanning

One of the ways to read input from the STDIN is to use `fmt.Scanf`

```go
package main

import (
	"fmt"
)

func main() {
	var firstNumber int
	var secondNumber int
	fmt.Println("What two numbers would you like to add?")
	fmt.Scanf("%d %d", &firstNumber, &secondNumber)
	fmt.Printf("Your total is: %d\n", firstNumber + secondNumber)
}
```

The problem with `fmt.Scanf()` is that we have to rely on the user to enter the values in the correct format.

The other option is to use `fmt.Sscanf()`

```go
package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func main() {
    /* family.csv
    John is 28 years old
    Jane is 27 years old
    */
	file, err := os.Open("family.csv")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var age int
		var name string

		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)
		if err != nil {
			panic(err)
		}
		if n == 2 {
			fmt.Printf("%s,%d\n", name, age)
		}
	}
}
```