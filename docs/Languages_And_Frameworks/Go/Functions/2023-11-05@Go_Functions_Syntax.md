# Go Functions Zone - Act 1: Overview

This is a short Act, all you have to do to pass is to write your own Go function
that returns the product of two numbers. And as we begin, keep in mind that, as in other programming languages, functions in Go must be as clear as possible, so
remember to use a descriptive name for your func...

> **Player**: Wait, what is a "function"?
>
> **Hint System:** OK, yeah. I suppose we should start with that.

<br />

## What is a function?

Think of a function as a tool. People use and re-use tools to transform something into something else in order to perform a job or activity.
Take a scissor for example. You can re-use a scissor over and over again to transform
a material (paper, plastic, etc.) into a modified version (trimmed, cut, etc.)

In a programming language, a function is very similar to this idea. It is a tool that
an application uses over and over to transform something (the input) into something
else (the output).

<br />

## What does a function in Go look like?

Go has a special keyword that we can use to create a function: `func`.

> **Player**: What is a keyword?
>
> **Hint System**: Software programs are made of a lot of words, some of those words you can choose,
> but others are "reserved", which means you cannot use them to name your own stuff.

This is an example of a function in Go:

```go
func double(p1 int) int {
    // Code to double the input...
}
```

The first thing in a function is the aforementioned keyword `func`.
Right after that comes the name of the function which you need to choose.
For example, `add` or `listFiles`, etc.
Next comes a pair of `()`.

If your function takes parameters (input), you will add them between the parenthesis.

<br />

## Parameters

In Go, to specify parameters for your functions you first...

> **Player:** Sorry, but there's another thing I don't know. What are parameters?
>
> **Hint System:** Ok, let's see... Think of a paper printer as our function. We can
> say that one of the "parameters" to the printer is a paper sheet. You cannot use
> a paper printer to print an apple, or a cat. The only thing that fits is a
> paper sheet. So in programming, a parameter tells other parts of your code what
> can go into the function.

So you first add the name of the parameter then the type, e.g.:
`(parameter1 float64)`. If you have more then one parameter of the same type,
e.g.: `(p1 float64, p2 float64)`, you can used this simplified syntax `(p1, p2 float64)`.

<br />

## Return Values

In the physical world, a tool must yield a result from its work, for example,
a sharpened pencil, or a trimmed paper sheet, etc. In the programming world,
function can (should) return one or more values as a result of its operation.

So after the closing parenthesis where you added the parameters, if your function
returns a value, that's where you specify the type, for example:
`func double(number int) int`. And like I mentioned it's possible to
return more than one value, `func doubleAndHalf(number int) (int, float64)`.

<br />

## The Body

After this closing parenthesis, you will add an opening curly bracket `{`...

> **Player**: What do you mean the "body"?
>
> **Hint System:** Just like your own body, from the outside you don't see it, but
> there is a lot going on in the inside to allow you to breathe, and think, etc.
> And using our paper printer example, if you open it up, you will see that
> inside there is a mechanism that does the actual printing. A function, has code
> inside to perform that actual job. That's the body of the function.

That curly bracket `{` must be in the same line as your function signature (the name
and other parts we just talked about).

In the following lines you add the function body. And at the end, on a new line you add the closing bracket:

```go
func double(p1 int) (int) {
    return p1 * 2 // <-- The body
}
```

> **Player:** This is getting more complex quickly. Can I save my progress here and quit for now?
>
> **Hint System:** Not at this point. If you want to save your progress, you must
> first create a function that returns the product of two numbers as mentioned
> in the beginning of this Act.
>
> Write and run your function here: [The Go Playground](https://go.dev/play/)

<br />
