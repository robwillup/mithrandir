# Common Concurrency Patterns

Learn about techniques for concurrent programming in Go.

- Non-blocking error channels
- Encapsulating goroutings
- Concurrency patterns

## Non-blocking error channel

```go
var (
    in = make(chan string)
    out = make(chan int)
    errCh = make(chan error, 1)
)

func worker(in <-chan string, out chan<- int, errCh chan<- error){
    for msg := range in {
        i, err := strconv.Atoi(msg)
        if err != nil {
            errCh <- err
            return
        }
        out <- i
    }
}
```

Make sure that every Go routine has an exit strategy.
The best practice is that every goroutine is drained, but we can't always guarantee that the consumer of a
goroutine will do that. 
So if a goroutine is not consumed properly we can introduce the risk for a resource leak.

To avoid this situation, we create what is called a non-blocking error channel. And this is actually one
of the only places where using buffered channels is recommended. 

If that error channel never is drained, then nothing is going to have a reference to that channel, and
the GC us going to clean up the channel as well as that message that is sitting in the buffer.

This means that my function will have an exit strategy even if an error occurs and no one is draining a channel.

## Encapsulating Goroutines

What if the output channels are uninitialized?
Channels in Go can have the `nil` value, which means they are uninitialized. And if you try to send a value
to an uninitialized channel in Go, the application panics!

To encapsulate the goroutine, you change it so that it does not take the two output channels as arguments. 
Instead, it creates those channels itself and returns them.

