# Tutorial: Get started with Go

Go code is grouped into packages, and packages are grouped into modules.

## About Modules

Your module specifies dependencies needed to run your code, including the Go version and the set of other modules
it requires.

When your code imports packages contained in other modules, you manage those dependencies through your code's own
module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file
stays with your code, included in your source code repository.

In actual development, the module path will typically be the repository location where your source code will be kept.
For example, the module path might be `github.com/mymodule`. If you plan to publish your module for others to use, the
module path must be a location from which Go tools can download your module.

## About Packages

A package is a way to group functions, and it's made up of all the files in the same directory.

