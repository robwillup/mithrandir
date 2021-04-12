# General Notes About Go Packages

Packages are Go's code organization strategy.

## What makes up a package?

>A package ... is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package.

- https://golang.org/ref/spec#Packages

* All Go source files in a single directory
  * Subdirectories are excluded

## Types of packages

* Library packages
  * designed to be consumed by other packages
  * Name must match the directory name
  * Should provide a focused set of related features
* Main Packages
  * Application entry point
  * Contains a main() function
  * can be in any directory
  * focus on app setup and initialization

## Naming Packages

Naming recommendations:

* Short and clear
* lowercase
* no underscores
* prefer nouns
* abbreviate judiciously (carefully and consciously)
* don't steal good names from users (event, etc)

**Examples:**

* package utilities -> vague
* package data_layer -> too long, contains underscore
* package dl -> unclear
* package time -> clear anc concise
* package json -> clear and concise

## Naming Package Contents

