# Go Routines

## What is a goroutine?

> A goroutine is a virtualization of a thread.
> Threads are the construction that the OS uses to implement concurrency.
> In Go, we use goroutines as an abstraction over those threads, and rely on the scheduler,
> that piece of the Go runtime, to manage allocating goroutines onto threads.
>
> Benefits:
> As a goroutine runs into a situation where it can't work, the scheduler can remove that goroutine from a thread,
> so the thread resource, the OS resource, can be allocated to something else. We are not tying up that thread,
> waiting for the goroutine to be able to resume its work.
> Since goroutines are created without interaction with the OS, they are very fast to create.
> Since they have a dynamic stack space, they start up very small (2KB), and can grow to have large stack spaces.


A goroutine has its own execution stack which has a variable stack space. Starting at 2KB it can go to a max of 2GB.
This makes goroutines quite lightweight.

Goroutines are managed by the Go runtine, scheduler, which creates an overhead and makes it a little less performant
than OS threads.

But goroutines are also inexpensive because the live directly inside the Go runtime so they don't have to interact with
the OS at all.

## Lifecycle of a Goroutine

## Advices about Goroutines

- Goroutines are cheap - so use them!
- Know how a goroutine will stop when you start it.
  - They can be made to block infenitely
  - You can create a resource leak in Go by creating a goroutine that never terminates, so its local variables are
    never released.
    So make sure you think about the entire lifecycle, making sure the goroutine as an exit strategy.
    Sometimes a goroutine can last for the entire lifecycle, and that's fine as long as you thought about and know
    what you're doing.
- Use channels to communicate between goroutines.
- Use sync.WaitGroup to synchronize completion of tasks.




