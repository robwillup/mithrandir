# Channels

## Buffered vs. Unbuffered

Unbuffered:

```go
var ch = make(chan string)

func sender() {
    ch <- "message" // The scheduler first blocks because there are no receivers.
}

func receive() {
    msg := <-ch     // Scheduler blocks this too and checks if sender and receive match. If matched, both functions are unblocked.
    fmt.Println(msg)
}

func main() {
    // synchronization elided for clarity
    go sender()
    go receive()
}
```

Buffered:

```go
var ch = make(chan string, 1) // construct a channel that has an internal buffer, or storage capacity.
                              // In this case the buffer is 1 and allows the channel to store a single
                              // message withing itself.
```

One very interesting way to see how buffered and unbufferend channels work is by creating a channel inside a function
and trying to get data from that channel in the same function. If the channel is unbuffered there will be a 
deadlock error because since functions execute sequentially, when the execution hits the line that is trying to send
into the channel the scheduler will block that line and so the execution cannot continue to the next line which would
receive from the channel. A buffered channel would not have that problem.

```go
package main

import "fmt"

func main() {
    ch := make(chan string)
    ch <- "Hello"
    fmt.Println(<-ch)
}
```

The above throws the error:

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
    main.go:7
exit status 2
```

But the following code with buffered channel does not have this problem:

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 1)
    ch <- "Hello"
    fmt.Println(<-ch)
}
```

## Directional Channels

By default, channels are biderectional.
But that can sometimes hard readability.

Using directional channels help making the intent clearer.

```go
func main() {
    ch := make(chan string)     // bidirectional channel
    
    go func(ch chan<- string) { // send-only channel
        ch <- "message"
    }(ch)

    go func(ch <-chan string) { // receive-only channel
        fmt.Println(<-ch)
    }
}
```
