## Longest Common Prefix

### My First Solution

```go
func longestCommonPrefix(strs []string) string {
    length := len(strs)
    prefix := strs[0]

    // If only one word, that's the prefix.
    if length == 1 {
        return prefix
    }

    // If the first word is empty, no common prefix.
    if len(prefix) < 1 {
        return ""
    }

    for i := 1; i < length; i++ {
        preLen := len(prefix)
        curLen := len(strs[i])

        // If the current word is empty, no common prefix.
        if curLen < 1 {
            return ""
        }

        // Common prefix cannot be longer than any word.
        if curLen < preLen {
            prefix = prefix[0:len(strs[i])]
            preLen = len(prefix)
        }

        for j, v := range strs[i] {
            if j < preLen && string(v) == string(prefix[j]) {
                continue;
            } else {
                prefix = prefix[0:j]
                break;
            }
        }
    }

    return prefix
}
```
