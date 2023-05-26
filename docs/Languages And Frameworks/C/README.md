# C Language

## The Evolution of C

* 1972: The C programming language is created by Dennis Ritchie
  * Dennis Ritchie worked at Bell Labs
  * It was named C because it was the successor of B
  * It was intended to be used to create programs for the Unix operating system, however soon the Unix kernel itself was rewritten in C
* 1978: *The C Programming Language* is published
  * The "hello world" idea started on this book
* 1989: ANSI C ratified
* 1999: "C99" released
  * new data types
  * single line comments
  * others
* 2011: "C11" released
  * generic macros
  * multi threading
* 2018: "C18" released
  * minor technical corrections to the language specification

## The Lowest High-level language

Assembly code is architecture specific, so it's not very portable.

The two problems that C solved that led high adoptions:

* Syntax that works at a higher level of abstraction
* Code portability (which was made possible by the creation of C compilers for different platforms)

```                                  
                                        [--Source Code--]

                    [--C Source Code --]   [--Compiler--]

[--Assembly Code--]   [--C Compiler--]  [-- Runtime/VM--]

[----------------------Machine Code---------------------]
```

## Advantages of using C

* Relatively small: about 40 reserved words
* fast and low level access to memory
* good for building
  * Operating Systems
  * cross-platform and fast apps
    * example: Git
  * embedded systems


## Programming Embedded Systems

C has become the defacto tool for programming embedded systems.

## Writing C Code

You only need two things to start writing C: a compiler and and editor.

### Compilers

* GNU Compiler Collection (GCC)
  * supports: C, C++, Go
* Clang
* CL (Visual Studio)

### Editors and IDEs

There are many options, but these are some of the features you should look for:

* syntax highlighting
* code completion
* for larger projects, IDEs will come in handy:
  * debugging
  * performance profiling
  * test runners


Popular IDEs:

* xCode - macOS
* Visual Studio - Windows
* CLion - Windows, macOS, Linux
* Visual Studio Code 

## Compiling C Program

### C File Types

Header files

```c
// products.h

void display_product(int);
double get_price(int);
```

Header files exist to share declarations between source code files. You include your functions in a header file and then include that file in the source files your you want to call the functions.

Source code files

```c
// inventory.c

void print_inventory(int id) {
    display_product(id);
}
```

### Compiling

program.c ----compiler---> program.obj --linker---> executable

There are tools that automate this process, such as Make.

A `Makefile` contains all the steps required for the build.

### Writing Procedural Code

C is a procedural language that uses *procedures* as the building blocks.

Procedures are also known as methods, functions or routines.

## Some Important Applications of C

* Supercomputers. They are powered by the Unix/Linux operating system, and their kernel is written in C.
* Embedded Systems and Hardware-level Programming
* Mars Curiosity Rover Flight Software
* The Linux kernel is almost entirely written in C.
* Highly-optimized Libraries (e.g. FFTW)

