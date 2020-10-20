# How to start a Rust project?

Run the following command to create a Rust project:

```bash
$ cargo new <app-name>
```

This command creates a directory with the name you provided containing the following:

* src/
* .gitignore
* Cargo.toml

## Command Line Arguments

The text after the name of the program is often called the "command line arguments", or "command line flags" (specially if they look like `--this`). Internally, the operating system usually represents them as a list of strings - roughly speaking, they get separated by spaces.

## Getting Arguments in Rust

The standard library contains the function `std::env::args()` that gives you an iterator of the given arguments. The first entry (at index 0) will be the name of your program (e.g. `grrs`), the ones that follow are what the user wrote afterwards. 

## CLI Arguments as data type

Instead of thinking about them as a bunch of text, it often pays off to think of CLI arguments as a custom data type the represents the inputs to your program.

In Rust it is common to structure programs around the data they handle, so this way of looking at CLI arguments fits very well.

So the following code defines a struct that has two fields to store data in:

```rust
struct Cli {
    pattern: String,
    path: std::path::PathBuf,
}
```
> **Aside:** *PathBuf* is like a String but for file system paths that work cross-platform

## Parsing CLI arguments with StructOpt

To get the actual arguments into we could manually parse the list of strings we get from the operating system and build the structure ourselves. Although this works, it's not very convenient. How would you deal with the requirement to support `--pattern="foo"` or `--pattern "foo"`? How would you implement `--help` ?

A much nicer way is to use one of the many available libraries. The most popular library for parsing command line arguments is called `clap`. It has all the functionality you'd expect, including support for sub-commands, shell completions, and great help messages. 

The `structopt` library builds on `clap` and provides a "derive" macro to generate `clap` code for `struct` definitions. This is quite nice, all we have to do is annotate a `struct` and it will generate the code that parses the arguments into the fields. 

Let
