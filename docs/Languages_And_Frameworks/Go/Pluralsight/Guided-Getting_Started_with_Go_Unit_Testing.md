# Getting started with Go unit testing

Tools such as the text editor which can warn about syntax errors, the type system which may warn that parts of the
code don't match up, the compiler which may refuse to compile the program if it encounters issues, help with the
correctness of the program from a *structural* and *syntactical* point of view.

But they have their shortcomings. They can have blind spots and miss some issues. They cannot tell you whether the
program will run correctly, even though it compiles.

For that we need tests (API tests, UI tests, integration tests, etc.). Let's focus on unit tests now.

> Unit tests do not validate the behavior of a program as a whole, but instead it validates its parts (units).
> A unit can be a single expression, or a function, or a class (or struct in the case of Go); an a unit test
> validates such a unit by passing it different kinds of input and checking that the output is as expected,
> ignoring the rest of the application.

By focusing on isolated units, unit tests let you ignore some of the complexity of your program. Unit tests are
generally fast, allowing you to write many of them and run them often.

## Unit tests in Go

Any file that has a name ending in `_test.go` is considered a test file by the Go tooling. Withing these files, any
function that starts with `Test` is considered a test function. Test functions take the shape
`func TextXXX(t *testing.T) {}`. They don't return anything.

For example, your codebase may contain a file `area.go`, containing a function called `GetArea`. If you wanted to write
tests for the `GetArea` function, you would create a file called `area_test.go`. Within this file, you would then
define one or more test functions, such as `TestGetAreaSquare` and `TestGetAreaRectangle`.

To make a test fail you can use one of several methods available on the `t *testing.T` struct.

|Method|Description|
|:----:|:---------:|
|t.Fail()|Marks the test as 'failed', but continues with the test of the test function.|
|t.FailNow()|Marks the test as 'failed' and stops execution of the current test function|
|t.Error(args ...any)|Logs the 'args' parameters to the console and then calls .Fail()|
|t.Fatal(args ...any)|Logs the 'args' parameters to the console and then calls .FailNow()|

### Anatomy of a test

Unit tests, and tests in general, are all about validating expectations or, in other words, making assertions. A unit
test executes a piece of your code (a function), passing in specific values, and checking that the result is what you
expected.

* In all but the simplest of unit tests, this involves a little (or a lot of) prep work; you may need to instantiate
  some values, set up mocks, or otherwise set up the specific conditions under which you want to validate your code.
* You then execute the code that you want to test, passing in the values that you prepared.
* After that, you make some assertions about the result.

The most basic assertion is to check whether your function returned an error or not.

Most tests follow the structure above which is often referred to as `Arrange/Act/Assert`:

* `Arrange:` set up the right preconditions for your test case
* `Act:` execute the code you want to test
* `Assert:` check the results and validate your assumptions about the outcome.

> Alternatively, this structure is referred to as `Given/When/Then`


