# Linked Lists

Like arrays, Linked Lists are a linear data structure.

Unlike arrays, Linked List elements are not stored at a contiguous location;
the elements are linked using pointers.

## Why Linked Lists?

Arrays have the following limitations:

1. The size of the array is fixed, so we must know the upper
limit on the number of elements in advance.
Also, generally, the allocated memory is equal to the upper
limit irrespective of the usage.

2. Inserting a new element in an array of elements is
expensive because the room has to be created for the new
elements and to create room existing elements have to be
shifted.

## Linked Lists Advantages Over Arrays

1. Dynamic size
2. Ease of insertion/deletion

## Drawbacks

1. Random access is not allowed. We have to access element
sequentially starting from the first node. So we cannot do
binary search with linked lists efficiently, with its default
implementation.