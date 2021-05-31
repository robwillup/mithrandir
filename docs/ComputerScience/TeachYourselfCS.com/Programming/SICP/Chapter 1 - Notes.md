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

## 1.1.5 The Subtitution Model for Procedure Application

For compound procedures, the application process is as follows:

* To apply a compound procedure to arguments, evaluate the body
of the procedure with each formal parameter replaced by the
corresponding argument.

The *substitution model* for procedure application is a model
that determines the "meaning" of procedure application.

Example of *substitution model*:

```scheme
(+ (square 6) (square 10))
```

If we use the definition of square, this reduces to

```scheme
(+ (* 6 6) (* 10 10))
```

which reduces by multiplication to

```scheme
(+ 36 100)
```

and finally to

136

* The purpose of the substitution is to help us think about
procedure application, not to provide a description of how
the interpreter works.

* The substitution model is a way to get started with thinking
formally about the evaluation process.
In general, when modeling phenomena is science and engineering,
we begin with simplified, incomplete models. As we examine
things in greater detail, these simplified models become
inadequate and must be replace by more refined models.

## 1.1.6 Conditional Expressions and Predicates

```scheme
(define (abs x)
  (cond ((> x 0) x)
        ((= x 0) 0)
        ((< x 0) (- x))))
```

The general form of a conditional expression consists of
the symbol `cond` followed by parenthesized pairs of
expressions called *clauses*. The first expression in each
pair is a *predicate* -- that is, an expression whose value
is interpreted as either true or false.

If none of the predicates is true, the value of the `cond` is undefined.

The word *predicate* is used for procedures that return true or false,
as well as for expressions that evaluate to true or false.

The special form `if` is a restricted type of conditional
that can be used when there are precisely two cases in the
case analysis. The general form of an `if` expression is

```scheme
(if <predicate> <consequent> <alternative>)
```

In addition to primitive predicates such as <, =, and >, there
are logical composition operations, which enable us to
construct compound predicates. The three most frequently used are these:

* (and <e1> .. <en>)

* (or <e1> .. <en>)

* (not <e>)

Notice that `and` and `or` are special forms, not procedures,
because the subexpressions are not necessarily all evaluated.
`Not` is an ordinary procedure.

