# C# 11 New Features

## Raw string literals

This is a new format for string literals which may contain arbitrary text, including whitespace, new lines, embedded quotes, and other special characters without requiring
escape sequencies.

```csharp
string text = """
    You
        don't
    have
    "to"
    worry about
    escaping
    """;
```

## Auto-default struct
