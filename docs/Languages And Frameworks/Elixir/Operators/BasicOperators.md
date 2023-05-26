# Basic operators

Elixir provides `+`,`-`,`*`,`/` as arithmetic operators, plus the functions `div/2` and `rem/2` for integer division and remainders.

Elixir also provides `++` and `--` to manipulate lists

```elixir
iex> [1,2,3,4] ++ [5,6,7]
[1, 2, 3, 4, 5, 6, 7]

iex> [1,2,3] -- [2]
[1, 3]
```

String concatenation is done with `<>`:

```elixir
iex> "Hello" <> " Elixir"
"Hello Elixir"
```

Elixir also provides three boolean operators: `or`, `and` and `not`. These operators are strict in the sense that they expect something that evaluates to a boolean (`true` or `false`) as their first argument:

````elixir
iex> true and true
true
iex> false or is_atom(:example)
true
````
