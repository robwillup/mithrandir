# Identifying functions and documentation

[Reference](https://elixir-lang.org/getting-started/basic-types.html#identifying-functions-and-documentation)

Functions in Elixir are identified by both their name and their arity. The arity of a function describes the number of arguments that the function takes. `round/1` identifies the function which is named `round` and takes `1` argument.

We can also use this syntax to access documentation. The Elixir shell defines the `h` functions, which you can use to access documentation for any function. For example, typing `h round/1` is going to print the documentation for the `round/1` function:

            iex(1)> h round/1
            * def round(number)

            @spec round(number()) :: integer()

            guard: true

            Rounds a number to the nearest integer.
                        
