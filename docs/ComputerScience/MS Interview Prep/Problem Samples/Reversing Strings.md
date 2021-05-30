# Reverse String

## Problem

Given an input string, reverse the string word by word.

## Solution

```csharp
public class Solution {
    public string ReverseWords(string s) {
        string[] words = s.Split(" ");
        var sb = new System.Text.StringBuilder();
        for(int i = s.Length - 1; i >=0; i--)
        {
            sb.Append(s[i]);
            if (i == 0)
                continue;
            sb.Append(" ");
        }
        return sb.ToString();
    }
}
```
