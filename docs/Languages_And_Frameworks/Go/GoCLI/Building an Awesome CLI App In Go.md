# Building an Awesome CLI app in Go

This is a study based on this [article](https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/)

## Topics

* Introduction to UX guidelines for CLIs
* Design our experience
* Build application

## Cobra

Cobra is a Cli library for Go that empowers applications. This application is a tool to generate the needed files to quickly create a
Cobra application.

Usage:

```bash
cobra [command]
```

## UX of CLI

Human Interface Guidlines for CLIs

UNIX Philosophy:

* Simple
* Clear
* Composable
* Extensible
* Modular
* Small

"Many UNIX programs do quite trivial things in isolation, but, combined with other programs, become general and useful tools."

-- Rob Pike

POSIX + GNU

Abbreviated Commands

ls -> list
cp -> copy
cat -> concatenate
cd -> change directory

Short and clear.

* Commands should primarily operate in the current directory.
* They should be verbs: ls, rm, unzip, find
* Args are input
* Commands should accept many args
* Order should matter
* Space is the separator
* An argument is something passed to the verb (command)

```bash
> cp [file] [newFile]
  |     |       |
verb  noun     noun
```

* Options and flags modify (adverb)
* Flags should be stackable

```bash
> rm -r -f
// Same as
> rm -rf
```

* A flag modifies behavior

```bash
> rm     --force      [file]
   |        |           |
 verb    adverb       noun
// Forcefully Remove this file
```

Adverbs express

* manner
* place
* time
* frequency
* degree
* level of certainty
* etc.

Adverbs answer

* How?
* In what way?
* When?
* Where?
* To what extent?

Flags should be input-able and pronounceable

```bash
> ls --width=40

// List the directory with a width of 40 cols
```

Preposition phrase = Adverb

* Consists of a preposition and its object
* Acts as an adverb
* "Speaking at a conference"

## CLI Apps

* httpd
* vi
* emacs
* git

An app is something

* Launch something
* Do more than one thing
* Collection of commands

Sub Commands

* dotnet new
* git clone
* npm install
* apt-get upgrade

* CLI apps do multiple things
* Apps are groups of commands
* (sub) Commands have flags & args
* All rules above still apply

```bash
> dotnet build App.csproj -c Release
    |     |     |          |     |
  app    cmd  input       flag  arg
  noun  verb  object     adverb/adjective

// Pronounceable
// Hey dotnet, build App.csproj for release
```

## About Cobra

* A CLI Command Framework
* A tool to generate CLI apps & commands
* Powers Kubernetes, Docker, Dropbox, Git Lfs, CoreOS, Hugo, Delve...

### Cobra init

```bash
cobra init github.com/robwillup/rosy -a "Robson William"
```
