# Review Contains Duplicate Problem

## Problem

Given an integer array `nums`, return true if any of the value appears more than once in the array, otherwise return
false.

Example 1:

    Input: nums = [1,2,3,3]
    Output: True

Example 2:

    Input: nums = [1,2,3,4]
    Output: false

## Recommend Time & Space Complexity

You should aim for a solution with `O(n)` time and `O(n)` space, where n is the size of the input array.

## My solution

I have seen this problem before and so I've learned that I could use the `visited` approach. Which works like this:

As I'm going through each element on the array, I can save the value in a hash map. Any value in that hash map is a
value I have visited or seen before, therefore if in the next iterations I encounter the same value again, it's a
duplicate and I can return true.

The time complexity of this algorithm is O(n) because I only have to go through the array once.

The space complexity of this algorithm is O(n) because in the worst case I may have created a hash map as big as the
original array.

## Pattern

This hash map solution for the Contains Duplicate problem falls into the `Hash Set` pattern, also called
`Seen Unseen` or `Frequency Counting` pattern.

### Why this Pattern

Pattern Trigger: "Check if any element appears more than once" -> Use a set to track seen elements.

    Core Idea: "Have I seen this before?"
    - If yes -> Duplicate found!
    - If no  -> Add to "seen" set

### The Pattern Template

```
seen = empty set
for each element in array:
    if element in seen:
        return true // duplicate!
    seen.add(element)
return false
```


