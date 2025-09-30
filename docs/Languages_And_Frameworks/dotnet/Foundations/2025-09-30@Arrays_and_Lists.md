# Arrays and Lists

## Arrays

- Data structure to store multible variables
- All variables must have the same type
- Accessed through an indes
- Arrays are reference types, even for value types used in the array
- Creation happens upon using new
- Size is set upon creation of the array, but can be at runtime.
- Arrays are zero-based

```C#
int[] employeeIds = new int[4];

int[] ids = new int[] { 1, 2 };
```

Methods

- CopyTo(): actual copy, not just another object pointing to the same array
- Sort()
- Reverse()
- Length - this is a property

## Collection

- List<int> list = new();
