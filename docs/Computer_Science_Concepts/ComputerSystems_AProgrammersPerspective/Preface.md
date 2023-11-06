# Preface

This book (known as CS:APP) is for computer scientists, computer engineers, and others
how want to be able to write better programs by learning what's going on "under the
hood" of a computer system.

The aim of this book is to explain the enduring concepts underlying all computer
systems.

If you study (not just read through) and learn (really grasp) the concepts in this
book, you will be on your way to becoming the rare power programmer who knows how
things work and how to fix them when they break. You will develop skills to write
programs that better use resources and capabilities provided by the operating system
and systems software. Your programs will operate correctly, across a wide range of
operating conditions and runtime parameters. They will run faster and they will not
have the flaws that make programs vulnerable to cyberattack. You will be prepared
to delve deeper into advanced topics such as compilers, computer architecture,
operating systems, embedded systems, networking, and cybersecurity.

**A typical code example:**

*Hello world in C:*

```c
#include <stdio.h>

int main()
{
	printf("Hullo!\n");
	return 0;
}
```

[C Playground](https://www.onlinegdb.com/online_c_compiler)

*Hello World in Rust:*

```rust
fn main() {
	println!("Hullo!");
}
```

[Rust playground](https://play.rust-lang.org/?version=stable&mode=debug&edition=2021)

*Hello World in C#:*

```csharp
System.Console.WriteLine("Hullo");
```

[.NET Playground](https://dotnetfiddle.net/)

*Hello World in Go:*

```go
package main

import "fmt"

func main() {
	fmt.Println("Hullo")
}
```

[Go Playground](https://go.dev/play/)

*Hello World in Elixir:*

```elixir
IO.puts "Hullo"
```

[Elixir Playground](https://www.jdoodle.com/execute-elixir-online/)

*Hello World in Python:*

```python
print("Hullo")
```

To study computer science it is fundamentally important that you do it, meaning,
as you see an opportunity to apply some code, or test something practical on your
computer, or a concept of a system that can be built, **do it!**.

I am including the runnable code/programs in the `./samples` directory under the
directory of this book folder in the `Mithrandir` repository.

As you read, try to solve each problem on your own and then check the solution
to make sure you are on the right track.
