# Go Types

## Runes

A rune is like a number tag for one single character, any character. Be it English, emoji, Hebrew, etc.

Unicode is a giant phone book for every character (millions of them). Each character gets a unique number,
called a `code point`. For example, `A` has the `code point` 65, or U+0041 in hex.

A rune stores that code point as an int32, so Go can handle any character perfectly, not just English alphabet.

### Why not just bytes?

A string is a list of bytes (tiny 8-bit chunks). Simple characters like 'A' fit 1 byte, so if it is stored as a byte
it can be correctly presented when converted by to string. But other characters like Hebrew characters can take more
bytes, so storing into bytes would break the number and would then be incorrectly represented when converted back.

### Example

```go
s := "A😊"

for _, r := range s {
    fmt.Println(r, string(r)) // 65 A    128522 😊
}
```
Rune gives you the whole character number safely, every time.

## Iota

Iota is a magic counter inside a const block that starts at 0 and ticks up by 1 for each line.

```go
const (
    Sunday = iota    // 0 (iota starts here)
    Monday           // 1 (ticks up, no need to write iota)
    Tuesday          // 2
    skipOne = iota   // 3 (explicit iota)
    Friday           // 4
)
```
