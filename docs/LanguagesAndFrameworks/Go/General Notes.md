# General Notes About Go

* Go was conceived in September 2007
* It was created by Robert Griesemer, Rob Pike, and Ken Thompson.
* All three creators worked at Google
* Go was announced in November 2009
* Go is much more than an upgraded version of C, despite all similarities.
* It borrows and adapts good ideas from other languages but avoids features that created complexity and led to unreliable code.
* Go is especially well suited for building infrastructure like networked servers, and tools and systems from programmers.


## The Origins of Go

Like biological species, successful languages beget offspring that incorporate the advantages of their ancestor; interbreeding sometimes leads to surprising strengths; and, very occasionally, a radical new feature arises without precedent.

* What programming languages have influenced Go?
  * ALGOL 60, Pascal, Modula-2, Oberon, Object-Oberon, Oberon-2, CSP, Squeak, Newsqueak, Alef, C.
  * Modula-2 inspired the package concept.
  * Oberon eliminated the distinction between module interface files and module implementation files.
  * Oberon-2 influenced the syntax for packages, imports, and declarations
  * Object-Oberon provided the syntax for method declarations.

In CSP a program is a parallel composition of processes that have no shared state; the processes communicate and synchronize using channels. 

## Go files

The import declaration must follow the package declaration. 

## Command-Line Arguments

The `os` package provides functions and other values for dealing with the operating system in a platform-independent fashion. Command-line arguments are available to a program in a variable named `Args` that is part of the `os` package; thus its name anywhere outside the `os` package is `os.Args`.