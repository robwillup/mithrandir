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

## Boolean

Elixir supports `true` and `false` as booleans. Elixir provides a bunch of predicate functions to check for a value type. For example, the `is_boolean/1` function can be used to check if a value is a boolean or not:

```elixir
iex> is_boolean(true)
true
iex> is_boolean(1)
false
```
You can also use `is_integer/1`, `is_float/1` or `is_number/1` to check, respectively, if an argument is an integer, a float, or either.

## Atoms

An atom is a constant whose value is its own name. Some other languages call these symbols. They are often use to enumerate over distinct values, such as:

```elixir
iex> :apple
:apple
iex> :orange
:orange
```

Atoms are equal if their names are equal. Elixir allows you to skip the leading `:` for the atoms `false`, `true` and `nil`.

Finally, Elixir has a construct called aliases which we will explore later. Aliases start in upper case and are also atoms.

## Strings

Strings in Elixir are delimited by double quotes, and they are encoded in UTF-8:

```elixir
iex> "hellö"
"hellö"
```
Elixir also supports string interpolation:

```elixir
iex> string = :world
iex> "hellö #{string}"
"hellö world"
```

Strings can have line breaks in them. You can introduce them using escape sequences:

```elixir
iex> "hello
...> Rob"
"hello\nRob"
iex> "hello\nRob"
"hello\nRob"
```
Strings in Elixir are represented internally by contiguous sequences of bytes known as binaries.

The String module contains a bunch of functions that operate on strings as defined in the Unicode standard:

```elixir
iex> String.upcase("hello")
"HELLO"
```