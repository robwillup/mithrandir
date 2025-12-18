# Two Sum

Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first.

## Problem

A list (array) is provided which contains two numbers which when added sum up to a target number which is also
provided.
Go through every item, comparing with others to find the two that sum to the target.
So I need to iterate over every item comparing with the others, or keep track of the ones I've seen before.

## Pattern

To me this looks like a "Seen/Unseen" or "Hash Set" pattern.

## Solution

To solve this, some basic math is needed.

I know that the array contains two numbers which add up to target. So, I know that

target - array[i] = array[j]

This means that at each iteration I can get the "complement", that is

"having this number, I need this other number (complement) to add up to target".

This other number can be found by iterating over the whole array at each step, or by storing previously seen numbers
and then comparing against them at each step.

with this patther

```
seen = empty set
for each element index in array:
    complement = target - element
    if complement in seen:
        return [seen[complement], index] // return smaller index first
    seen[element] = index
return [] // no result found
```

## Asymptotic analysis

Time complexity is O(N) = time grows linearly with input array size

Space complexity is O(N) = size of new map grows linearly with input array size.

## Go code

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, v := range nums {
        complement := target - v
        if _, ok := seen[complement]; ok {
            return []int{seen[complement], i}
        }
        seen[v] = i
    }
    return []int{}
}
```
