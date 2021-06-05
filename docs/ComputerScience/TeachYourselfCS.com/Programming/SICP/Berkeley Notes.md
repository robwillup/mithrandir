# Computer Science 61A

Computer Science is about controlling complexity.

Layers of Abstraction:

|Application|
|:---------:|
|High level language (e.g. Scheme)|
|Low level language (e.g. C)|
|Machine language / architecture|
|logic gates|
|transistors|
|Quantum physics|

Abstract means 'built of top of other pieces'.

## First Class Data Type

Can be:

* the value of a variable
* an argument to a procedure
* the value returned by a procedure
* a member of an aggregate
* anonymous

## Procedure returned by procedure

```scheme
(define (make-adder num)
(lambda (x) (+ x num)))

(define plus3 (make-adder 3))
```

## Follow ups

* Look up `let` in SICP
