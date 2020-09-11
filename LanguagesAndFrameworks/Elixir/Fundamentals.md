# Fundamentals

Checking Elixir version:

            $ elixir -v

## Interactive mode

When you install Elixir you have 3 new executables: `iex`, `elixir`, and `elixirc`. 

For now, let's start by running `iex` (or `iex.bat` if you are on Windows PowerShell, where `iex` is a PowerShell command) which stands for Interactive Elixir. In interactive mode we can type any Elixir expression and get its result. Let's warm up with some basic expressions.

```Elixir
Erlang/OTP 21.0 [64-bit] [smp:2:2] [...]

Interactive Elixir (1.10.4) - press Ctrl+C to exit
iex(1)> 40 + 2
42
iex(2)> "hello" <> " world"
"hello world"
```

On Windows, running the command `iex.bat --werl` opens up an Erlang terminal.

## Running scripts

Create a file with the `.exs` extension:

            $ touch simple.exs

And add the following content:

`IO.puts "Hello world from Elixir"`

Now run the script using:

            $ elixir simple.exs

## Scalability

All Elixir code runs inside lightweight threads of execution (called processes) that are isolated and exchange information via messages.

Due to their lightweight nature, it is not uncommon to have hundreds of thousands of processes running concurrently in the same machine. Isolation allows processes to be garbage collected independently, reducing system-wide pauses, and using all machine resources as efficiently as possible (vertical scaling).

Processes are also able to communicate with other processes running on different machines in the same network. This provides the foundation for distribution, allowing developers to coordinated work across multiple nodes (horizontal scaling).