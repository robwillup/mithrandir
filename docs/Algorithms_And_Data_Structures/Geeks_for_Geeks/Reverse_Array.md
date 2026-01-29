# Reverse array

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[ :]
	array := []rune(args[1])
	fmt.Println(string(array))

	i := 0
	j := len(array) - 1

	for j - i >= 1 {
		temp := array[i]
		array[i] = array[j]
		array[j] = temp
		i++
		j--
	}

	fmt.Println(string(array))
}
```
