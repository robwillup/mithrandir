# Java OOP Basics

| Modifier | Visibility | Usable on Classes | Usable on Members |
|:-:|:-:|:-:|:-:|
| No access modifier | Only within its own package (a.k.a. package private)| Y | Y |
| public | Everywhere | Y | Y |
| private | Only within the declaring class | N* | Y |

\* as private applies to top-level classes; private is available to nested-classes;

* When a class is marked as public is simply means that class can be accessed from anywhere, but that does not mean that we can create instances from that class anywhere.

## Special References

* this:
  * implicit reference to current object
    * Useful for reducing ambiguity
    * allows an object to pass itself as a parameter
* null
  * represents an uncreated object
    * can be assigned to any reference variable

## Accessors and Mutatorr

Use the accessor/mutator pattern to control field access (getters and setters)

