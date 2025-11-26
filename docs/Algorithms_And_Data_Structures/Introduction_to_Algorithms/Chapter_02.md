# Chapter 2 - Getting Started

Analyzing the Insertion sort and Merge sort algorithms.

## Insertion sort

Our first algorithm, insertion sort, solves the sorting problem introduced in Chapter 1:

**Input**: A sequence of n numbers (a1, a2, ..., an).

**Output**: A permutation (reordering) of the input sequence such that a1 <= a2 <= ... <= an.

The numbers that we wish to sort are also known as the keys. Although conceptually we are sorting a sequence, the input
comes to us in the form of an array with n elements.

> This algorithm is similar to sorting a hand of cards.

> pseudocode is not concerned with issues of software engineering, such as data abstraction, modularity,
> error handling, etc.

Insertion soft is an efficient algorithm for sorting a small number of elements. It works the way many people sort a
hand of cards. We start with an empty left hand and the cards face down on the table. We then remove one card at a
time from the table and insert it into the correct position in the left hand. To find the correct position for a card
we compare it with each of the cards already in the hand, from right to left. At all times, the cards held in the left
hand are sorted, and these cards were originally the top cards of the pile on the table.

We present our pseudocode for insertion sort as a procedure called `Insertion-Sort`, which takes as a parameter an
array `A[1...n]` containing a sequence of length n that is to be sorted. The algorithm sorts the input numbers
**in place**: it rearranges the numbers within the array A, with at most a constant number of them stored outside
the array at any time. The input array A contains the sorted output sequence when the `Insertion-Sort` procedure is
finished.

(a) [5][2][4][6][1][3]      (b) [2][5][4][6][1][3]      (c) [2][4][5][6][1][3]

(d) [2][4][5][6][1][3]      (e) [1][2][4][5][6][3]      (f) [1][2][3][4][5][6]

```pseudocode
for j = 2 to A.length
  key = A[j]
  // Insert A[j] into the sorted sequence A[1...j-1].
  i = j - 1
  while i > 0 and A[i] > key
    A[i + 1] = A[i]
    i = i - 1
  A[i + 1] = key
```

First iteration of the above pseudocode on [5][2][4][6][1][3]

* `j` is equal to 2, which is the index for the second position. That position starts with value 2
* `key` is equal to 2
* `i` is equal to 1
* `i` is greater than 0 and `A[i]`, which is 5, is greater than 2
* `A[i + 1]` is equal to `A[1 + 1]` and equal to `A[2]`
* `A[i]`, which is 5, is assigned to `A[2]`
* `A[2]` now contains 5
* `i` becomes 0
* inner loop ends
* `A[i + 1 ]` is equal to `A[0 + 1]` and to `A[1]`
* `key`, which is 2, is assigned to `A[1]`
* `A[1]` now contains 2
* When the outer loop ends, this is the array [2][5][4][6][1][3]

Second iteration of the above pseudocode on [2][5][4][6][1][3]

* `j` is 3
* `key` is 4
* `i` is 2
* `A[i]` 5
* the value (5) in the i index (2) goes ahead to the j index (3)
* `i` is decremented and becomes 1
* the value in the `i` index (2) is now less than the `key` (4)
* inner loop terminates
* The value in `key` is assigned to `A[i + 1]` (inserted)
* When the outer loop ends, this is the array [2][4][5][6][1][3]

