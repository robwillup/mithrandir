# How to work with files in Rust

In this document I share the tips I've been learning about how to work with files in Rust.

## Check if a file exists

This is one of the ways I found to check if a file exists in Rust:

```rust
let exists = std::path::Path::new("path/to/file").exists();

if exists {
    // do something with the file
}
```

## Create a file

```rust
std::fs::File::create("test.txt").expect("Could not create the file");
```

## Appending text to a file

```rust
use std::fs::OpenOptions;
use std::io::prelude::*;

fn main() {
    ...
    let mut file = OpenOptions::new()
        .write(true)
        .append(true)
        .open("test.txt")
        .unwrap();

        if let Err(e) = writeln!(file, "some text") {
            eprintln!("Couldn't write to file: {}", e);
        }
}
```