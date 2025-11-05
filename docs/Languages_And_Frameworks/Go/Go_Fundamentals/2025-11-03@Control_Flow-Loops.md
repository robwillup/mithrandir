# Control Flow: Loops

```Go
// All loops are 'for' loops in Go

for { ... }                                    // infinite loop

for condition { ... }                          // loop till condition

for initializer; test; post clause { ... }     // counter-based loop

// Looping with Collections

for key, value := range collection { ... }     // we can loop over arrays, slices, and maps

for key := range collection { ... }            // just the indexes

for _, value := range collection { ... }       // just the values
```
