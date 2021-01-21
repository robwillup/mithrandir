# Strings in C

C does not have a `String` class like other higher level programming languages. To create a string in C you must create a char array:

```c
char name[] = "Connie";
```

## Memory Layout of C Strings

----Contiguous memory location--->
|C|o|n|n|i|e|null \0|

1 char = 1 byte

Each char is stored in memory using its associated ASCII Value

Recap = String in C is a *Null-terminatedd contiguous sequence of chars*

## The Built-in Char Type

the char type actually stores an integer which is the ASCII number representation of a given character.

## Basic I/O with Strings

When you create a string with 40 elements, for example, you can actually only store 39 characters, because the last element is always the null terminator.

To read a string from command line terminal, you can:

```c
char name[40];

scanf("%39s", name)
```

%s is a placeholder for strings, and the 39 specifies the maximum char count.

The reason we don't use an ampersand (&) in front of the variable in scanf is because in C "the name of an array is equivalent to the address of its first element"

To print a string, you need to use the format specifier in printf("%s \n", name)
