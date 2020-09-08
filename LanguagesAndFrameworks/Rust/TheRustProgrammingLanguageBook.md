# The Rust Programming Language

# Introduction

## Who is Rust for
### Students
Rust is for students and those who are interested in learning about systems concepts. Using Rust, many people have learned about about topics like operating systems development.

### Companies
Companies use Rust in production for a variety of tasks, including:
* Command line tools
* web services 
* DevOps tooling 
* Embedded devices 
* Audio and video analysis and transcoding
* Cryptocurrencies
* Bioinformatics
* Search engines
* Internet of Things
* Machine Learning

### Open Source Developers
Rust is for people who want to build the Rust programming language.

### People who value speed and stability
Rust gives you speed for writing code and Rust programs are fast. The Rust compiler's checks ensure stability. By striving for zero-cost abstractions, higher-level features that compile to lower-level code as fast as code written manually, Rust endeavors to make sage code be fast code as well.

# Getting Started
## Installation
Linux or macOS

            curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh

Make sure you have a C compiler

Windows
https://www.rust-lang.org/tools/install

C++ build tools are required.

## Updating

        rustup update
        or
        rustup update stable


## Checking installation

        rustc --version

## Local documentation

        rustup doc

# Hello world

## rustfmt

        rustfmt .\file_name.rs

```Rust
fn main() {
    println!("Hello, world");
}
```

* Rust style is to ident with 4 spaces, not a tab.
* `println!` calls a Rust macro. If it called a function instead, it would be entered as `println` (without !)

Rust is an *ahead-of-time* compiled language, meaning you can compile a program and give the executable to someone else, and they can run it even without having Rust installed.


# Crate

Crate is a collection of Rust source code files. There are `binary crate` (executable) and `library crate` (to be used by other programs.)

