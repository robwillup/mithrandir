# Data Structures Overview

Data items that we group together are called "compound data".

We have to store compound data into some soft of data structure.

No data structure is perfect. They are all good and bad in different aspects.

The way we measure how good a data structure is at doing different things such as
* add()
* get()
* sort()
* search()

is called the Big O notation.

The Big O notation is pretty much the measure of how well an operation scales.
If you have ten items and then you add a million more, how much longer will each
operation take.

## Linked List

The atomic unit of the linked list is something called "Node",
which contains a value and a pointer.

The value is that data that you are storing and the pointer
contains the address of the next node.

The first node in the linked list is called the *Head* and
the last node which doesn't contain a pointer is called *Tail*.

| Pros | Cons |
|:----:|:----:|
|Adding| retrieval|
|Deleting| searching|


## Array

A continuous block of cells in the computer memory.

| Pros | Cons |
|:----:|:----:|
| retrieval | adding|
| | deleting |

## Hash Table (Dictionary)

| Pros | Cons |
|:----:|:----:|
| add | key collisions |
| remove | |
| retrieve | |

## Stack + Queue

