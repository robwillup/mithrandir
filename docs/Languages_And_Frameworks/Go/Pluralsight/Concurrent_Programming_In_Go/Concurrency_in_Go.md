# Concurrency in Go

This version of the application is synchronous:

```go
package main

import (
        "encoding/json"
        "fmt"
        "log"
)

func main() {
        receiveOrders()
        fmt.Println(orders)
}

func receiveOrders() {
        for _, rawOrder := range rawOrders {
                var newOrder order
                err := json.Unmarshal([]byte(rawOrder), &newOrder)
                if err != nil {
                        log.Print(err)
                        continue
                }
                orders = append(orders, newOrder)
        }
}

var rawOrders = []string{
        `{"productCode": 1111, "quantity": 5, "status": 1}`,
        `{"productCode": 2222, "quantity": 42.3, "status": 1}`,
        `{"productCode": 3333, "quantity": 19, "status": 1}`,
        `{"productCode": 4444, "quantity": 8, "status": 1}`,

```

Now, let's make it concurrent:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	// If the goroutine is called without a wait group, the main function will not wait for it and will just terminate.
	var wg sync.WaitGroup
	wg.Add(1) // Indicates how many goroutines need to be waited.
	go receiveOrders(&wg) // This is a goroutine which is managed by the portion of the runtime called 'scheduler'
	wg.Wait() // Blocks this function until all wait groups emit Done.
	fmt.Println(orders)
}

func receiveOrders(wg *sync.WaitGroup) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		orders = append(orders, newOrder)
	}
	wg.Done()

}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

```
