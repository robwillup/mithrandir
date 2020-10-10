# Functional Programming - Big Picture

What is functional programming?

> Functional Programming is a paradigm that treats computation as the evaluation of mathematical 
> functions and avoids chanding-state and mutable data.

**Other Paradigms**

* Object-oriented
* Prodedural
* Declarative

> **Pure Function**
> Does not cause side effects, and will always return the same result for a given input.

## Contrasting with OOP

> I fear not the man who has practiced 10,000 kicks once, but I feat the man who has
> practiced one kick 10,000 times.
> - *Bruce Lee*

> It is better to have 100 functions operate on one data structure 
> than 10 functions on 10 data structures
> - *Alan Perils*

## Functional Programming Data Structures

* List []
* Map {}

OOP vs. FP
| OOP | FP |
|:---:|:---:|
| More data sctructures | Fewer data structures|
| Fewer functions per data structures | More functions per data structure |
| 10 functions on 10 data structures | 100 functions on 1 data structure |

**The major difference is how data is processed**

## Do One Thing Well

**The Unix Philosophy**

* Make each program do one thing well
* Expected the output of every program to become the input of another
* Design software to be tried early
* Use tools over unskilled labor

**Function Programming Philosophy**

* Make each function do one and only one thing well
* Expected the output of every function to become the input of another
* Design functions to be tested early

## Generic Type Signature

```
findAuthor :: (a, [a]) -> [a]
```

The signature above denotes a function that receives `a` which is any type, and an array of `a` and returns an array of `a`.

## Why Type Signatures?

* Prevalent in documentation
* 100 functions on 1 data structure
* Simplify second principle

## From Output to Input

*Type signaturess show how functions can be chained*

Several functional programming languages have developed an operator called `pipe`. 

```
funcA([a]) |> funcB()
```

```
loadFile('courses')
|> splitOnNewLine()
|> findAuthorCourses(author)
|> titleCase()
```

*Chaining allows easier decomposition*

*Focused functions lead to higher reuse*

## Test Early

* Always test early

**Function Programming Advantages for Unit Tests**

* Easier to set up
* Easier verification

**Complexity through simplicity**

* Decompose problem
* Reuse building blocks
* Test early

*Learning functional programming involves a shift in perspective*

## Recuding Bugs with Immutable Data

**What is immutable data?**

This is an example of how functional programming could "modify data" without chanding it:

```
def addToCart (cart, item) do
    %{
        items = cart.items.concat(item),
        total = cart.total + item.price
    }
end
```

## Why Functional Programming Matters

### Caching

> **Memoization**
> To speed up computer programs by storing the results of expensive function calls and 
> returning the cached result when the same inputs occur again

*Pure functions are easier to memoize*

### Laziness

> **Lazy Evaluation**
> delays the evaluation of an expression until its value is needed.

 **Lazy evaluation defers execution until the time it's actually needed**

