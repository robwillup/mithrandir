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

4. We aboid `this.` unless absolutely necessary.

5. We always specify the visibility, even if it's the default (e.g. `private string _foo` not
`string _foo`). Visibility should be the first modifier (e.g. `public abstract` not `abstract public`)

6. Namespace imports should be specified at the top of the file, outside of `namespace` declarations,
and should be sorted alphabetically, with the exception of `System.*` namespaces, which are to be
placed on top of all others.

7. Avoid more than one empty line at any time. For example, do not have two blank lines between
member of a type.

8. Avoid spurious free spaces. For example, avoid `if (someVar == 0)...`, where the dots mark the
spurious free spaces. Consider enabling "View White Space (Ctrl+R, Ctrl+W)" if using Visual Studio
to aid detection.

9. If a file happens to differ in style from these guidelines (e.g. private members are named
`m_member` rather than `_member`), the existing style in that file takes precedence.

10. We only use `var` when the type is explicitly named on the right-hand side, typically due to
either `new` or an explicit cast, e.g. `var stream = new FileStrem(...)` not
`var stream = OpenStandardInput()`.
	1. Similarly, target-typed `new()` can only be used when the type is explicitly named on the
	left-hand side, in a variable definition statement or a field definition statement, e.g.
	`FileStream stream = new(...)`, but not `stream = new(...);`, where the type was specified on
	a previous line.

11. We use language keywords instead of BCL types (e.g. `int, string, float` instead of
`Int32, String, Single`, etc.) for both type references as well as method calls (e.g. `int.Parse`
instead of `Int32.Parse`).

12. We use PascalCasing to name all our constant local variables and fields. The only exception is
for interop code where the constant value should match exactly the name and value of the code you
are calling via interop.

13. We use PascalCasing for all method names, including local functions.

14. We use `nameof(...)` instead of `"..."` whenever possible and relevant.

15. Fields should be specified at the top within type declarations.

16. When including non-ASCII characters in the source code use Unicode escape sequences (\uXXXX)
instead of literal characters. Literal noon-ASCII characters occasionally get garbled by a tool or editor

17. When using labels (for goto), indent the label one less than the current indentation.

18. When using a single-statement if, we follow these conventions:
	1. Never use single-line form (e.g. `if (source == null) return;`)
	2. Using braces is always accepted, and required if any block of an `if/else if/.../else` compound
	statement uses braces of if a single statement body spans multiple lines.
	3. Braces may be omitted only if the body of every block associated with an `if/else if/.../else`
	compound statement is places on a single line.

19. Make all internal and private types static or sealed unless derivation from them is required.
As with any implementation detail, they can be changed if/when derivations is required in the future.
