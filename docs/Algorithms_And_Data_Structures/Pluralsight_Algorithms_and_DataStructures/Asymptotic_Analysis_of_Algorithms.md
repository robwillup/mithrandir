- # Asymptotic Analysis of Algorithms

We are talking about resources. What are resources?

- Operations: The number of times we need to perform some operations
- Memory: How much memory is consumed by the algorithms
- Others: Network transfer, compression ratios, disk usage, temperature, etc.

But in general we're talking about operations.

## Big-O Notation

It represents the upper limit of an algorithm's cost and it's associated with O(n).

## Asymptotic Analysis

> Asymptote
> The asymptote of a curve is a line where the distance between the curve and the line approach zero as they tend
> towards infinity.

In other words, it's the measurement of how the inputs of an algorithm affect the behavior as the inputs approach
some limit.

**What is an upper limit?**

It depends. Look at your domain.

Values that don't affect the overall shape of the curve are ignored. For example:

- O(n+1) is still just O(n)
- O(2n) is still just O(n)

### Examples of Big-O

- O(1): The cost of the algorithm is unchanged by the input size

> Don't confuse fixed with fast!
> A fixed cost algorithm might still do a lot of expensive work.
> The point is: The amount of work that is done is not dependent on the size of the input.

- O(log n): A function whose cost scales logarithmically with the input size.
- O(n): A function whose cost scales linearly with the size of the input.
- O(nm): A function which has two inputs that contribute to the growth (only if n or m is tiny. If similar sizes then it's probably O(n2))
- O(n2) n square: A function that exhibits quadratic growth relative to the input size.

We basically just care about Big-O. 
