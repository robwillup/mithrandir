# Two Sum Problem

This problem gives you an array of integers and a single integer and expects that you write the code
to find the two elements in the array that add up to that single integer called `target`.

So the idea is that you find the most efficient way to search for two elements in an array.

## Brute Force Solution

```go
func twoSum(nums []int, target int) []int {
    for i, a := range nums {
        for j, b := range nums {
            if i != j && a + b == target {
                return []int{i, j}
            }
        }
    }

    return nil
}
```

This solution works, but it's a brute force solution because it compares each number in the array to all other
numbers, and when this patterns of comparing every element to others in the array shows up you can know that you
have a nested loop, and the time complexity is O(n²) and the space complexity is O(1).

## How to make it better

Often times when an algorithm has an O(n²) time complexity and an O(1) space
complexity it can be improved to have the time complexity being O(n) and space
also O(n).

Yeah, the time complexity gets better but space gets worse. But we generally
care more about the time complexity because it's easy to increase the size of
memory, but we can never buy more time.

## Fastest solution

```go
func twoSum(nums []int, target int) []int {
    mapNumberToIndex := make(map[int]int)

    for i, number := range nums {
        mapNumberToIndex[number] = i
    }

    for i, number := range nums {
        if j, isFound := mapNumberToIndex[target - number]; i != j && isFound {
            return []int{i, j}
        }
    }
    return []int{}
}
```
