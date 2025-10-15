# Trained with Perplexity

Gauging my initial understanding of Go to taylor this prep effectively.

**Initial Questions**:

1. How familiar are you with Go's concurrency model (goroutines and channels)?
   Can you briefly explain how you've used them?

> No too familar. I know that there is a routine keyword and that channels are the mechanism Go uses to let the caller
> await for the response. I have used it in a CLI application that would make REST requests.

2. Have you worked with Go interfaces and polymorphism? What challenges did you face?

> I worked a little with interfaces in Go, but not with polymorphism.

3. What is your experience with error handling in Go? How do you handle and
   propagate errors in your projects?

> I have worked extensively with `if err != nil` in Go, and to propagate the error I would return it from my function.

**Corrections to the Initial Questions above**:

**Go routines**

- Goroutines are lightweight threads started with the `go` keywork. They run concurrently but share memory.
- Channels are typed conduits to communicate safely between goroutines, often used to synchronize or pass data.

```Go
ch := make(chan string)
go func() {
    ch <- "hello"
}()
msg := <-ch
fmt.Println(msg)
```

**Interfaces and Polymorphism**

- An interface defines method signature that types can implement implicitly.
- Polymorphism lets different types be used interchangeably via a common interface.

```Go
type Speaker interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

func MakeSpeak(s Speaker) {
    fmt.Println(s.Speak())
}
```

**Error Handling**:

It's important to know that `if err != nil` and returning errors is idiomatic ways to handle errors in Go. Another
useful practices is wrapping errors for context using `fmt.Errorf` or the `errors` package.

Example:

```Go
if err != nil {
    return fmt.Errorf("doing X failed: %w", err)
}
```

This helps trace error chains with `%w` and `errors.Is` / `errors.As`.

### Reinforcing about Go Concurrency, Interfaces, and Error Handling

1. **Channels:** What happens if you try to send a value on a channel but no goroutine receives it? What if you try to
    receive from a channel but no value is sent yet?

My first answer:

> ~The value is ignored. The execution will halt and wait for the returned value.~

Correct answer:

> Sending on a channel blocks if no goroutine is receiving. It's not ignored. The sending goroutine is paused until
> another goroutine receives.
> Similarly, receiving from a channel blocks if no value is sent yet, so the goroutine waits.


2. **Interfaces:** Given an interface with one method, how does a Go type implement it? Do you need to explicitly
    declare it implements the interface?

My first answer:

> No, it is implicitly implemented.

Correct answer:

> Go types implement interfaces implicitly by having the required methods.

3. **Error handling:** How would you add context to an error returned by a standard library function to indicate which
    part of your function failed?

My first answer:

> I would use fmt.Errorf or the errors package.

Correct answer:

> `fmt.Errorf("context: %w", err)` is the idiomatic way to wrap errors.

4. **Code exercise:** Write a short Go snippet where you launch two goroutines that each send an integer to a channel,
   then the main goroutine receives and prints both numbers.

```Go
func main() {
    ch := make(chan int)

    go func() { ch <- 1 }()
    go func() { ch <- 2 }()

    val1 := <- ch
    val2 := <- ch

    fmt.Println("First number is %d and second number is %d\n", val1, val2)
}
```

This creates a single channel, launches two goroutines sending values into it, then the main goroutine receives twice.

## More interview questions for Go

1. **Memory Management:** How does Go manage memory? What is garbage collection in Go, and how does it affect
    performance?

Correct answer:

> Go uses a mark-and-sweep garbase collection algorithm, which automatically reclaims memory occupied by objects no
> longer in use.
>
> The GC operates concurrently with your application, minimizing pause times, by marking reachable objects and
> sweeping unreachable ones without stopping the world.
>
> It employs a three-color abstraction (white, gray, black) to efficiently track object reachability during the
> marking phase, avoiding the need for full pause or compaction.
>
> Go's GC does not compact memory (relocate objects), which helps with high performance and reduces overhead
> associated with moving objects.
>
> It's non-generational, handling all objects uniformly regardless of age.
>
> It uses write barriers to ensure consistency during concurrent marking, making it suitable for high-throughput,
> low-latency applications.

2. **Slices vs Arrays:** What is the difference between a slice and an array in Go? How are slices implemented under
    the hood?

> Arrays in Go a fixed size collections of elements of the same type. Their size is part of their type, e.g.: `[3]int`
> is an array of exactly 3 ints. Arrays are value types and copied when assigned or passed to functions.
>
> Slices are references (headers) to a continuous segment of an underlying array. A slice consists of a pointer to the
> array, a length, and a capacity. Slices themselves are mutable but refer to the underlying array.
>
> Modifying a slice mutates the underlying array, not a new one. When you append beyond capacity, Go create a new array
> internally, copies the data there, and updates the slice header to point to this new array.
>
> Thus, slices are mutable views into arrays, not new copies created on every modification.
>
> Arrays are rarely used directly in idiomatic Go; slices provide dynamic sizing and are passed by reference, which is
> more efficient for most use cases.

3. **Goroutine Coordination:** How can you use `select` with channels? Can you give an example where `select`
    improves concurrency handling?

> The `select` statement in Go is used to wait on multiple channel operations simultaneously. It is like a switch
> statement but for channels - it blocks until one of the channels involved is ready for communication (either send or
> receive), and then it executes the corresponding case.
>
> How select works:
> - It evaluates each case, which must be a channel send or receive.
> - It blocks until at least one channel operation can proceed.
> - If multiple channels are ready, one is chosen randomly.
> - It optionally support a default case which runs immediately if no channels are ready, avoiding blocking.

4. **Advanced Interfaces:** How does the empty interface `interface{}` work in Go? How would you assert or convert
    types stored in an empty interface?

> In Go, `interface{}` is the empty interface and can hold values of any type because every type implements zero
> methods automatically, so it satisfies the empty interface.
>
> The concrete type of the value stored inside the interface is called the dynamic type, while the interface itself has
> a static type of `interface{}`
>
> You can assert the dynamic type from the interface using a type assertion with syntax like:
>
> v, ok := i.(int)
> if ok {
>   // v is int
> } else {
>   // assertion failed
> }
>
> You can also use type switches to handle multiple possible dynamic types stored in an interface.
>
> `interface{}` enables flexible, reusable functions and APIs that can accept any type, similar to how "dynamic" types
> work in other languages, but Go's interfaces are sill statically types and checked.

5. **Defer Statement:** What is the purpose of `defer` in Go? When is it executed? Can you give an example of a
    practical use?

> The `defer` keyword schedules a function call to run after the surrounding function returns, regardless of how it
> returns (normal return, error, panic)
>
> Deferred functions are executed in a last-in, last-out (LIFO) order right before the surrounding function exists.
>
> It is typically used to ensure cleanup tasks run like closing files, unlocking mutexes, or releasing resources, even
> if the function returns early or panics.

```Go
func CopyFile(src, dst string) error {
    file, err := os.Open(src)
    if err != nil {
        return err
    }
    defer file.Close() // Ensures file is closed when CopyFile finishes.

    // Copy contents here etc.

    return nil
}
```

`file.Close() is guaranteed to run at the end, ensuring proper resource cleanup.
