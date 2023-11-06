# Go Functions For The 5-Year-Old Me (who already know some programming concepts)

As we begin, keep in mind that, as in other programming languages, functions in Go
must be as clear as possible.

> **5yo me**: What is a "function"?

## What is a function?

Think of a function as a tool. A tool is something that you can use and re-use to
transform something into something else in order to perform a job or activity.
Take a scissor for example. You can re-use a scissor over and over again to transform
a material (paper, plastic, etc.) into a modified version (trimmed, cut, etc.)

In a programming language, a function is very similar to this idea. It is a tool that
an application uses over and over to transform something (the input) into something
else (the output).

## What does a function in Go look like?

Go has a special keyword that we can use to create a function: `func`.

> **5yo me**: What is a keyword?
>
> **Older me**: A software is made of a lot of words, some of those words you can choose,
> but others are "reserved", which means you cannot use them to name your own stuff.

This is an example of a function in Go:

```go
func double(p1 int) int {
    // Does some stuff to double the input...
}
```

The first thing in a function is the aforementioned keyword `func`.
Right after that comes the name of the function which you need to choose.
For example, `add` or `listFiles`, etc.
Next comes a pair of `()`.

If your function takes parameters (input), you will add them between the parenthesis.

## Parameters

In Go, to specify parameters for your functions you first...

> **5yo me:** What is a parameter?
>
> **Older me:** Ok, let's see... Think of a pencil sharpener as our function. We can
> compare the hole in the pencil sharpener as a "parameter". You cannot sharpen an
> apple, or a book. The only thing that fits is a pencil. So in programming, a
> parameter tells other parts of your code what can go into the function.

So you first add the name of the parameter then the type, e.g.:
`(parameter1 float64)`. If you have more then one parameter of the same type,
e.g.: `(p1 float64, p2 float64)`, you can used this simplified syntax `(p1, p2 float64)`.

## Return Values

In the physical world, a tool must yield a result from its work, for example,
a sharpened pencil, or a trimmed paper sheet, etc. In the programming world,
function can (should) return one or more values as a result of its operation.

So after the closing parenthesis where you added the parameters, if your function
returns a value, that's where you specify the type, for example:
`func double(number int) int`. And like I mentioned it's possible to
return more than one value, `func doubleAndHalf(number int) (int, float64)`.

## The Body

After this closing parenthesis, you will add a opening curly bracket `{`...

> **5yo me**: What do you mean the "body"?
>
> **Older me:** Just like your own body, from the outside you don't see it, but
> there is a lot going on in the inside to allow you to breathe, and think, etc.
> A using our pencil sharpener example, if you open it up, you will see that
> inside there is a blade that does the actual sharpening. A function, has code
> inside to perform that actual job. That's the body of the function.

That curly bracker `{` must be in the same line as your function signature (the name and other parts we just talked about).

In the following lines you add the function body. And at the end, on a new line you add the closing bracket:

```go
func double(p1 int) (int) {
    return p1 * 2 // <-- The body
}
```

> **Older me:** I see you're already getting distracted, so that's enough for now.
> Next time, let's talk more about other aspects of functions. Now, go have some func.
