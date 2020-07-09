# Code Organization

[Official Docs](https://golang.org/doc/code.html#Organization)

Go programs are organized into **packages**. A package is a **collection of source files in the same directory that are compiled together**. Functions, types, variables and constants defined in one source file are **visible** to all other source files within the same package.

A repository contains one or more modules. A module is a collection of related Go packages that are released together. A Go repository typically only contains one module, located at the root of the repository. A file names `go.mod` there declares the *module path*: the import path prefix for all packages within the module. The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another go.mod file (if any).

