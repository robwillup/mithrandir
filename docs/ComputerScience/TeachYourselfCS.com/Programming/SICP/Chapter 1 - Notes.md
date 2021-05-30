# Building Abstractions with Procedures

The various kinds of expressions, each with its associated
evaluation rule, constitute the syntax of the programming
language.

## 1.1.4 Compound Procedures

Elements that must appear in any powerful programming language:

* Numbers and arithmetic operations are primitive data and procedures.
* Nesting of combinations provides a means of combining operations.
* Definitions that associate names with values provide a
limited means of abstraction.

*Procedure definition* is a much more powerful abstraction
technique by which a compound operation can be given a name
and then referred to as a unit.

We begin by examining how to express the idea of "squaring".
We might say, "to square something, multiply it by itself".
This is expressed in Schemee as:

```scheme
(define (square x) (* x x))
```

We have here a *compound procedure*, which has been given the name `square`.
The procedure **represents** the operation of multiplying something by itself.
Evaluating the definition creates this compound procedure
and associates it with the name `square`.

> A procedure **represents** an operation

In Lisp / Scheme, compound procedures are used in exactly
the same way as primitive procedures.

* primitive: (+ 5 7)
* compound: (sum 5 7)

