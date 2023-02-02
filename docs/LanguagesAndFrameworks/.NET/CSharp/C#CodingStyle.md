# C# Coding Style

The general rule we follow is "use Visual Studio defaults".

1. We use Allman style braces, where each brace begins on a new line. A single line statement block
can go without braces but the block must be properly indented on its own line and must not be nested
in other statement blocks that use braces. One exception is that a `using` statement is permitted to
be nested within another `using` statement by starting on the following line at the same indentation
level, even if the nested `using` contains a controlled block.

2. We use four spaces of indentation (to tabs).

3. We use `_camelCase` for internal and private fields and use `readonly` where possible.
   1. Prefix internal and private instance fields with `_`
   2. Prefix static fields with `s_`
   3. Prefix thread static fields with `t_`.
   4. When used on static fields, `readonly` should come after `static` (e.g. `static readonly`
   	  not `readonly static`).
   5. Public fields should be used sparingly and should use PascalCasing with no prefix when used.