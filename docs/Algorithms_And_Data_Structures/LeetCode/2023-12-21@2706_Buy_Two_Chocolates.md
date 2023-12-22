## 2706. Buy Two Chocolates

Understanding the best solution for this problem.

### Description

You are given an integer array prices representing the prices of various chocolates in a store. You are also given a single integer money, which represents your initial amount of money.

You must buy exactly two chocolates in such a way that you still have some non-negative leftover money. You would like to minimize the sum of the prices of the two chocolates you buy.

Return the amount of money you will have leftover after buying the two chocolates. If there is no way for you to buy two chocolates without ending up in debt, return money. Note that the leftover must be non-negative.

**Example 1:**

Input: prices = [1,2,2], money = 3
Output: 0
Explanation: Purchase the chocolates priced at 1 and 2 units respectively. You will have 3 - 3 = 0 units of money afterwards. Thus, we return 0.

**Example 2:**

Input: prices = [3,2,3], money = 3
Output: 3
Explanation: You cannot buy 2 chocolates without going in debt, so we return 3.

**Constraints:**

* 2 <= prices.length <= 50
* 1 <= prices[i] <= 100
* 1 <= money <= 100

### My solution

When I was first trying to resolve this problem, I assumed that any problem that wants you to get two numbers in an array is best solved by creating a map that you can access the members with some sort of simple math in the key, similar to what is done in the `two sum` problem:

```go
if j, isFound := mapNumberToIndex[target - number]; i != j && isFound {
    return []int{i, j}
}
```

However, that's not the case with this chocolate problem because calculation result isn't exact. The problem does not want you to buy two chocolates so that you have an exact 0 as leftover, it can be any number as long as it's non-negative.

When I realized that a map would not help at all, I brute forced my way and came up with an algorithm that had a time complexity of O(n²) and space complexity of O(1).

```go
import "sort"

func buyChoco(prices []int, money int) int {
    sort.Ints(prices)

    for i, x := range prices {
        for j, y := range prices {
            if i != j {
                leftover := money - (x + y)
                if leftover >= 0 {
                    return leftover
                }
            }
        }
    }

    return money
}
```

Another brute force solution I have seen is to not sort the list, but to keep track of the smallest sum at every iteration.

### Best solution

So my assumption now is that when you need to find numbers whose criteria is not exact, then that cannot be used as an index or key for a map.

The best solution in LeetCode with Go is this:

```go
import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if money < prices[0]+prices[1] {
		return money
	}
	return money - prices[0] - prices[1]
}
```

Let's understand it line by line, then at the end we talk about why it's the best.

```go
import "sort"
```

The `sort` package is imported because in the function it will be necessary to order the prices array.

```go
func buyChoco(prices []int, money int) int {
    // ...
}
```

Obviously, the first line is just the function signature that is provided by the problem, and the last line is just the closing parenthesis.

```go
sort.Ints(prices)
```

This line sorts the `prices` array in ascending order. So after this line, the first elements on the `prices` array will be the cheapest, and the last elements will be the most expensive.

```go
if money < prices[0]+prices[1] {
```

This line is very clever. If the sum of the first and second chocolates, which are the two cheapest chocolates in the list, is greater than the money I have then that means all other chocolates are more expensive and I don't have enough money to buy two chocolates, so I don't even need to keep looking at the rest, I can just:

```go
return money
```

This line `return money` is inside the above if block, which is closed in the nest line. But, if we don't enter this if block, then that means I have enough money to buy the cheapest two chocolates, which means we can return the result/leftover in the next line:

```go
return money - prices[0] - prices[1]
```

And this is so, because the problem description says:

> You would like to minimize the sum of the prices of the two chocolates you buy.

That means the two smallest prices.

### Time and Space Complexity

The time complexity is O(n log n). This may vary depending on the implementation of the sorting algorithm in the programming language.

The space complexity is O(n) or O(log n). When sorting an array in place, some extra space is used. The space complexity depends on the implementation of the sorting algorithm in the programming language.

### Considerations

It is often a good practice to save the length of an array in a variable if it is used multiple times in the code.

It is a good practice to use meaningful variable names even if they are very simple or short-lived, like an index.

It is worth noting the **Greedy** is an algorithmic paradigm as well. It is a way of solving problems by making the locally optimal choice at every step, hoping that it will lead to a globally optimal solution. It is used for optimization problems. Although, it may not always lead to the optimal solution.

Sorting is a common operation in programming. It is used to arrange the elements of a collection in a particular order. There are various sorting algorithms with different time and space complexities. There are two broad categories of sorting algorithms, namely:

* comparison based sorting algorithms, and
* non-comparison based sorting algorithms.

To solve this problem, we are encouraged to implement this greedy approach with sorting and we should find the inbuilt sorting function in their language of choice.
