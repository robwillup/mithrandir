## Go Strings

As I was trying to solve a problem in LeetCode, I needed to index a string and learned that if you try to get a character from a string using the syntax:

```go
myString[2]
```

that is going to return a byte, not a character. So to work around that I had to do:

```go
string(myString[2])
```

There must be a better way to do this, but I still need to learn that.
