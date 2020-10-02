# Basic Types

Elixir basic types:
* integers
* floats
* booleans
* atoms
* strings
* lists
* tuples

## Basic arithmetic

Note that 10 / 2 returned a float 5.0 instead of an integer 5. This is expected. In Elixir, the operator `/` always returns a float. If you want to do integer division or get the division remainder, you can invoke the `div` and `rem` functions:

```Elixir
$ iex
Interactive Elixir (1.10.3) - press Ctrl+C to exit (type h() ENTER for help)
iex(1)> 1+2
3
iex(2)> 5*5
25
iex(3)> 10/2
5.0
iex(4)> div(10,2)
5
iex(5)> div 10, 2
5
iex(6)> rem 10, 3
1
iex(7)>
```

Notice that Elixir allows you to drop the parentheses when invoking named functions. This feature gives a cleaner syntax when writing declarations and control-flow constructs.

Elixir also supports shortcut notations for entering binary, octal, and hexadecimal numbers.

```Elixir
iex(7)> 0b01
1
iex(8)> 0b11
3
iex(9)> 0xF
15
```

Float numbers require a dot followed by at least one digit and also support `e` for scientific notation:

```Bash
iex> 1.0
1.0
iex> 1.0e-10
1.0e-10
```

Floats in Elixir are 64-bit double precision.

## Identifying functions and documentation