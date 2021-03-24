# Factory and Abstract Factory

Build code that is more:

* Reusable
* Maintainable
* Testable

It is important to be able to:

* understand the characteristics of the factory patterns
* implement the factory patterns in C#
* understanding of the benefits and tradeoffs 

## What is a Factory?

To understand of the *Factory* pattern is, it helps to thing about what a factory is in the real world. The whole purpose of a factory is to produce a product.

As consumers we don't care about the creation process, all we care about is the finished product. One of the benefits of have a factory is that multiple different consumers can all get their product from the factory.

>A factory is an object for creating objects
- Wikipedia

One of the purposes of a factory is to allow you to reuse code, as well as introducing extensibility.

## Flavours of the Factory Pattern

* Simple Factory
* Factory Method
* Abstract Factory

## Factory Pattern Characteristics

* Client: asks for a created product
* Creator: facilitates a creation
* Product: The product of the creation


## The Simple Factory

## The Factory Method

## The Abstract Factory Pattern

Not all applications will need the *Abstract Factory* pattern. In many cases the **simple factory** and **factory method** will be sufficient.

>The abstract factory pattern provides a way to encapsulate a group of individual factories that have a common theme without specifying their concrete classes.
- Wikipedia