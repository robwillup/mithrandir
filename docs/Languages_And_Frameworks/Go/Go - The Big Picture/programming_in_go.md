# Programming in Go

> **NOTE:**
> Finally nail down what statically and strongly typed means.
> * Statically typed: the type of a variable must be known at compile time.
> * Strongly typed: strict rules about how data types can be used, and generally do not allow
>                   implicit type conversions between unrelated types.

- Go is statically and strongly typed
- C-inspired syntax
- Garbage collected
- Fully compiled
- Rapid compilation

## Go's Evolution

C -> C++ -> Java -> C# | JavaScript | ...
|-> Go

Go runs on faster computers with multiple CPUs and is network aware.

Go does not see everything as an object, such as Java and C#.
Go is not trying to be a new C++.

Go goes back to the days of C where data an behavior is separated.
When necessary, it is possible to combine data and behavior.

Go does not force a lot of cerimonies and features on you upon start. You can simply opt-in if you do need them.

## Development tools

* Visual Studio Code
* Goland
* Vim
* Emacs
* Eclipse
* Others

Gopls is a Go language server. It's maintained by the Go code team.
It's used by the IDEs to provide a standardized experience.

## Common Go frameworks

### Network Services

* Go Kit
  * Comprehensive microservice framework 
  * gokit.io
* Gin
  * Fast, lightweight web framework
  * gin-gonic.com
* Gorilla Toolkit
  * Collection of useful tools without framework lock-in
  * gorillatoolkit.org

### CLIs

* Cobra
  * Framework for building command-line interface applications
  * github.com/spf13/cobra
* Standard library
  * Zero dependency API

### Cloud infrastructure

* Docker
* Kubernetes
* Terraform
  * Cloud infrastructure management platform

## Community Resources

* go.dev
  * Docs
  * **Effectice Go**
  * Play
  * Packages
  * Blog
  * wiki
  * Get started
* Pluralsight
* Wiki/The Go Community
  * Slack channel
  * Reddit
  * Conferences
