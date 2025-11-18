# Concurrent Programming in Go

Concurrency in Go is a fundamental aspect of the language's design, empowering developers to write efficient, scalable,
and responsible software. Go's concurrency model revolves around goroutines and channels, providing a powerful yet
simple way to handle concurrent tasks.

## Goroutines

Goroutines are lightweight threads managed by the Go runtime. They enable concurrent execution of functions, allowing
multiple tasks to run simultaneously within a single Go program. Goroutines are incredibly cheap in terms of memory
footprint and overhead, making it practical to spawn thousands of them within a single application.

To create a goroutine, you simply prefix a function call with the `go` keyword. For example:

```go
func main() {
    // Start a new goroutine
    go sayHello()

    // Continue with main execution
    fmt.Println("Main function")

    // Sleep for a while to allow the goroutine to finish
    time.Sleep(1 * time.Second)
}

func sayHello() {
    fmt.Println("Hello from goroutine")
}
```

## WaitGroup

Wait groups are a synchronization mechanism provided by the sync package in Go. They allow you to wait for a collection
of goroutines to finish their execution before proceeding further in the program. Wait groups are particularly useful
when you have a dynamic number of goroutines and need to ensure they all complete their tasks before continuing.

### How wait groups work

Wait groups are represented by `sync.WaitGroup` type. You create a new wait group using `var wg sync.WaitGroup`, and
then add the number of goroutines you want to wait for using the `Add()` method.

Each goroutine increments the wait group counter using `wg.Add(1)` before starting its task. When a goroutine completes
its task, it decrements the counter using `wg.Done()`.

In the below example, two goroutines are added to the wait group using `wg.Add(2)` before their execution begins.

Finally, the main goroutine, or any other goroutine waiting for the completion of the tasks, calls `wg.Wait()` to block
until all goroutines have finished their tasks.

```go
func main() {
    var wg sync.WaitGroup

    // Add two goroutines to the wait group
    wg.Add(2)

    // Start goroutines
    go doSomeWork(&wg)
    go doSomeWork(&wg)

    // Wait for all goroutines to finish
    wg.Wait()
    fmt.Println("All goroutines have finished")
}

func doSomeWork(wg *sync.WaitGroup) {
    // Signal that the goroutine has finished
    defer wg.Done()
    // Simulate some work
    fmt.Println("Goroutine: Working...")
    time.Sleep(time.Second)

    fmt.Println("Goroutine: Finished")
}
```

In this example, `defer wg.Done()` is used inside the `doSomeWork` function. The `defer` ensures that no matter how the
function exits, whether it returns normally or panics, it will always call `wg.Done()` before exiting. This helps
ensure that the `WaitGroup` counter is decremented appropriately, allowing `wg.Wait()` in the `main` function to block
until all workers are finished.

This pattern ensures that you properly wait for all goroutines to finish before proceeding, even if an error occurs
within a goroutine.

## Mutexes

Let's learn about mutex and how it can be implemented to ensure that variables are only updated only by one goroutine
at a time.

### Mutex

A mutex, short for mutual exclusion, is a synchronization primitive used to control access to shared resources in
concurrent programs. It ensures that only one goroutine can access a shared resource at a time, preventing data races
and ensuring consistency.

In Go, a mutex is represented by the `sync.Mutex` type from the `sync` package. It provides two main methods:

* `Lock()`: Acquires the mutex. If the mutex is already locked by another goroutine, `Lock()` will block until it
  becomes available.
* `Unlock()`: Releases the mutex, allowing other goroutines to acquire it.

To use a mutex in your Go code, follow these steps:

* Declare a variable of type `sync.Mutex`.
* Use `Lock()` to acquire the mutex before accessing the shared resource, and `Unlock()` to release it afterwards.

```go
var (
    counter int
    mutex   sync.Mutex
)

func increment() {
    mutex.Lock()
    counter++
    mutex.Unlock()
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment()
        }()
    }

    wg.Wait()

    fmt.Println("Counter: ", counter)
}

```

## Channels

Compared to wait groups, channels offer a more flexible and expressive way to manage concurrency:

### Synchronization

Channels offer built-in synchronization, ensuring safe data transmission between goroutines without extra synchronization primitives.

### Error handling

Channels facilitate error propagration alongside data, simplifying error management and reducing error susceptibility.

### Ordering

Channels enforce operation sequence, guaranteeing correct result processing without added synchronization overhead.

### Dynamic Scaling

Channels excel in scenarios with unknown or dynamic goroutine counts. They enable on-demand goroutine creation and coordination via channels.

### Creating Channels

Channels are created using the make function with the `chan` keyword followed by the type of data that the channel with transmit.

```go
// Creates an unbuffered channel of type int
ch := make(chan int)
```

### Channel Operations

Channels support two main operations: sending and receiving values. These operations are performed using the `<-` operator:

```go
ch <- value // Send value into the channel
value := <-ch // Receive value from the channel.
```

### Buffered Channels

By default, channels are unbuffered, meaning they only accept a value if there's a corresponding receiver ready
to receive it. Buffered channels, on the other hand, have a fixed capacity and can store a certain number of values
without a corresponding receiver. Here's how you create a buffered channel:

```go
ch := make(chan int, bufferedSize)
```

### Closing Channels

Channels can be closed to indicate that no more values will be sent. Receivers can use the second return value
from a receive operation to determine if the channel has been closed.

```go
close(ch)
```

### Channel Direction

You can specify the direction of a channel in its type signature to restrict its usage to sending or receiving
operations. This helps enforce communication protocols and prevent misuse of channels.

```go
func sendData(ch chan<- int) {
    // Send data into channel
}

func receiveData(ch <-chan int) {
    // Receive data from channel
}
```

Here's how to use channels for communication between two goroutines.

```go
func sendData(ch chan<- int) {
    ch <- 10
    ch <- 20
    ch <- 30
    close(ch)
}

func main() {
    ch := make(chan int)

    go sendData(ch)

    for {
        value, ok := <-ch
        if !ok {
            break
        }

        fmt.Println("Received: ", value)
    }
}
```

## Error Propagation using Channels

In Go, channels can be used not only for communication between goroutines, but also for propagating errors. By
sending error values through channels, goroutines can report errors to other parts of the program or to the main
goroutine responsible for error handling:

```go
errCh := make(chan error)
```

Example:

```go
func doTask(resultCh chan int, errCh chan error) {
    result <- 42
    errCh <- errors.New("something went wrong")
}

func main() {
    resultCh := make(chan int)
    errCh := make(chan error)

    go doTask(resultCh, errCh)

    select {
        case result := <-resultCh:
            fmt.Println("Result:", result)
        case err := <-errCh:
            fmt.Println("Error:", err)
    }
}
```
