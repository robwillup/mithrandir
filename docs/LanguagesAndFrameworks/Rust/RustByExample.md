# Formatted Print

Printing is handled by a series of macros defined in `std::fmt` some of which include:

* `format!`: write formatted text to `String`
* `print!`: same as `format!` but the text is printed to the console (io::stdout).
* `println!`: same as `print!` but a newline is appended.
* `eprint!`: same as `format!` but the text is printed to the standard error (io::stderr).
* `eprintln!`: same as `eprint!` but a newline is appended.

> **What is a macro?**
> A macro (short for "macroinstruction", from Greek where *macro* means "long, large") in computer science is a rule or pattern that specifies how a certain input sequence 
> (often a sequence of characters) should be mapped to a replacement output sequence (also often a sequence of characters) according to a defined procedure. 
> *- Wikipedia*

All of the macros above parse text in the same fashion. As a plus, Rust checks formatting correctness at
compile time.


