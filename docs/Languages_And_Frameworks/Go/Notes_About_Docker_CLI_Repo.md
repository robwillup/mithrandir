# Notes from Observations of the Docker CLI Repository

The following notes have been taken from observing and studying the Docker CLI repo.
This is a practice recommended in the book `The Complete Software Developer's Career Guide.` by John Sonmez.

## Root of the repository

* There are no `*.go` files in the root of the project.
* Some of the most different files there, that I had not seen before, are:
  * `TESTING.md`: This file has instructions about how testing is done in the project.
  * `MAINTAINERS`: This file contains information about all the people involved in the project.

* This Docker CLI project also uses `cobra-cli`.

## Commands

* Go files are named with full words when short, not abbreviations. For example:
  * `list.go` instead of `ls.go`
* But when the name is long, there are abbreviations:
  * `opts.go`
* When the file has more than one word, they are separated by underscore `_`:
  * `signals_unix.go`, `formatter_stats.go`
* Vast majority of command files (`run.go`, `list.go`, etc.) have a corresponding `_test.go` file.
