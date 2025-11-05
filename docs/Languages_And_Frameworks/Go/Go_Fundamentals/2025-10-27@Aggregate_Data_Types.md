# Aggregate Data Types

- Arrays
- Slices
- Maps
- Structs

## Using Arrays

An array is a fixed size collection of a single data type.

Example:

value: |0|1|1|3|7|4|8|9|0|
index: |0|1|2|3|4|5|6|7|8|

```go
var arr [3]int         // array of 3 ints
fmt.Println(arr)       // [0 0 0]
arr = [3]int{1, 3, 5}  // array literal

fmt.Println(arr[1])    // 2
arr[1] = 99            // update value
fmt.Println(arr)       // [1 99 3]

fmt.Println(len(arr))  // 3

arr == arr2            // false - arrays are comparable - Go checks value by value
```

Arrays are copied by value.

## Understanding Slices

A slice is a subset of an underlying array. That's why it's callec `slice`, because you slice the array.

So if you have the following array as the underlying data source:

|1|2|3|4|5|

A slice from it would have to specify the starting point and a length. E.g.,:

Starting point: 2
Length: 2

slice: |3|4|

Slices don't contain their own data. Slices are reference types, they refer to data stored somewhere else.
Slices are like windows to arrays.

So that means if you make a change to the underlying array, it will be simutaneously reflected on the slice.
The other way arround is also true.

Slices are dynamic.

```go
var s []int          // slices of ints
fmt.Println(s)       // [] (nil)
s = []int{1, 2, 3}   // slice literal

fmt.Println(s[1])    // 2
s[1] = 99            // update value
fmt.Println(s)       // [1 99 3]

s = append(s, 5, 7)  // add elements to the slice
fmt.Println(s)       // [1 99 3 5 7]

s = slices.Delete(s, 1, 3) // remove indices 1, 2 from slice

s == s2             // Compile time error. Slices are not comparable in Go. You need to use the slices package.
```

## Understanding Maps

Maps associate values to keys.

```Go
var m map[string]int                               // declare a map. Key is string, values integers
fmt.Println(m)                                     // map[] (nil)
m = map[string]int{"foo": 1, "bar": 2}             // map[foo:1 bar:2]

fmt.Println(m["foo"])                              // lookup value in map
m["bar"] = 99                                      // update value in map

delete(m, "foo")                                   // remove entry from map
m["baz"] = 418                                     // add value to map

fmt.Println(m["foo"])                              // 0 - queries always return results. Even if not there, the zero value for the type is returned.
v, ok := m["foo"]                                  // comma okay syntax verifies presents. True if value came from the map, false to indicate it's a zero value.
```

To copy maps, if you just use the assignment operator that's going to copy by reference.
To copy the map by value, you need to use the `maps.Clone` function.

Because maps are also reference type, they cannot be compared with `==`.

## Understanding Structs

We can associate multiple data types together.

```Go
var s struct{                        // Declare an anonymous struct
    name    string
    id      int
}

fmt.Println(s)                       // {"" 0} -- structs are value types. If it was reference, the zero values for the members would be nil.

s.name = "Arthur"


type myStruct struct{                // create custom type based on struct
    name    string
    id      int
}
```

