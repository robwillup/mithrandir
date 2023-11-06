# Design Patterns

This is my personal journey to trully understand **Design Patterns**.

## Etymology

As usual, it is helpful to understand the etymology behind the terms. The word *design*, late 14c, comes from the Latin *designare* "to mark out, devise, choose". 
The Italian verb *disegnare*, 16c, developed the senses "to contrive, plot, intend" and "to draw, paint, etc." As a noun, the origins of the word *design* mean 
"a scheme or plan in the mind".

As for **pattern**, the meaning "a figure corresponding in outline to an object that is to be fabricated and serving as a guide for its shape and dimensions" is from late 14c, 
but the earliest references to this word were "outline, plan, model, an original proposed for imitation". 

Having the meaning of those terms clearly defined on our heads helps us see what design patterns are and why they exist. Based on the etymology, *design patterns* 
can be understood as *outlines for contriving*, or examples that should be followed in the drawing and creationg of something. Software, in our context.

## Definition in Software Engineering

A **design pattern** is a solution to typical problems that arise when designing software. The solution is general and repeatable, and that means *design patterns* are not 
ready-to-use code snippets, rather they are more like templates that can help solve problems in different situations and scenarios. 

## Benefits of Design Patterns

* Speed up development process
  * This happens because they provide tested, proven development paradigms that guides developers.
* Prevents subtle issues
  * There are issues that might only become visible later in the implementation.
* Improves code readability
  * Other developers who are familiar with the design patterns will find it easier to read the code.


>Often, people only understand how to apply certain software design techniques to certain problems
- https://sourcemaking.com/design_patterns

To examplify the statement above, imagine a programmer that is aware of "DRY" and who has more experience with backends. He probably knows how to avoid recreating the same
functionality in his backend code, but he may not do the same in a frontent page, and end up with lots of duplicated code pieces.

The solutions provided by design patterns are general, which means that they can be applied to a variety of problems.
