# C# Coding Style

Coding conventions serve the following purposes:

* They create a consistent look to the code, to that readers can focus on content, not layout.
* They enable readers to understand the code more quickly by making assumptions based on their experience.
* They facilitate copying, changing, and maintaining the code.
* They demonstrate C# best practices

## Source and References

* [https://learn.microsoft.com/en-us/dotnet/csharp/fundamentals/coding-style/coding-conventions]
* [https://github.com/dotnet/runtime/blob/main/docs/coding-guidelines/coding-style.md]
* [https://google.github.io/styleguide/csharp-style.html]

The general rule we follow is "use Visual Studio defaults".

## Braces

We use Allman style braces, where each brace begins on a new line. A single line statement block
can go without braces but the block must be properly indented on its own line and must not be nested
in other statement blocks that use braces. One exception is that a `using` statement is permitted to
be nested within another `using` statement by starting on the following line at the same indentation
level, even if the nested `using` contains a controlled block.

## Naming conventions

In the following examples, any of the guidance pertaining to elements marked `public` is also
applicable when working with `protected` and `protected internal` elements, all of which are intended
to be visible to external callers, so this naming convention is unaffected by modifiers such as const,
static, readonly, etc.

If a file happens to differ in style from these guidelines (e.g. private members are named
`m_member` rather than `_member`), the existing style in that file takes precedence.

### PascalCase

* Names of classes, methods, enumerations, public fields, public properties, namespaces: `PascalCase`.

Use `PascalCase` when naming `class`, `record`, or `struct`:

```csharp
public class DataService
{
}

public record PhysicalAddress(
	string Street,
	string City,
	string StateOrProvince,
	string ZipCode);

public struct ValueCoordinate
{
}
```

When naming an interface, use pascal casing in addition to prefixing the name with an `I`. This clearly
indicates to consumers that it's an `interface`:

```csharp
public interface IWorkerQueue
{
}
```

Public fields should be used sparingly and should use PascalCasing with no prefix when used. When naming
`public` member of types, such as fields, properties, events, methods, and local functions, use pascal
casing.

```csharp
public class ExampleEvents
{
	// A public field, these should be used sparingly
	public bool IsValid;

	// An init-only property
	public IWorkerQueue WorkerQueue { get; init; }

	// An event
	public event Action EventProcessing;

	// Method
	public void StartEventProcessing()
	{
		// Local function
		static int CountQueueItems() => WorkerQueue.Count;
		// ...
	}
}
```

When writing positional records, use pascal casing for parameters as they're the public properties of
the record.

```csharp
public record PhysicalAddress(
	string Street,
	string City,
	string StateOrProvince,
	string ZipCode);
```

We use `PascalCasing` to name all our *constant local variables and fields*. The only exception is
for interop code where the constant value should match exactly the name and value of the code you
are calling via interop.

#### Files

* File names and directory names are `PascalCase`, e.g. `MyFile.cs`.
* Where possiblee the file name should be the same as the name of the main class in the file, e.g.
`MyClass.cs` contains class `public class MyClass {}`.
* In general, prefer one core class per file.

### camelCase

* Names of local variables, parameters: `camelCase`.
* For casing, a "word" is anything written without internal spaces, including acronyms. E.g., `MyRpc`
instead of ~~`MyRPC`~~

Use `camelCasing` when naming `private` or `internal` fields, and prefix them with `_`.

```csharp
	private void Foo(string firstName)
	{
		string greetings = $"Hello, {firstName}";
		Console.WriteLine(greetings);
	}
	private IWorkerQueue _workerQueue;
```

When working with `static` fields that are `private` or `internal`, use the `s_` and for thread static use
`t_`.

```csharp
public class DataServer
{
	private static IWorkerQueue s_workerQueue;

	[ThreadStatic]
	private static TimeSpan t_timeSpan;
}
```

When writing method parameters, use camel casing.

```csharp
public T SomeMethod<T>(int someNumber, bool isValid)
{
}
```

Use `readonly` where possible. When used on static fields, `readonly` should come after `static` (e.g. `static readonly` not `readonly static`).

## Commenting conventions

* Place the comment on a separate line, not at the end of a line of code.
* Begin comment text with an upper case letter.
* End comment text with a period.
* Insert one space between the comment delimiter (//) and the comment text.
* Don't create formatted blocks of asterisks around comment.
* Ensure all public members have the necessary XML comments providing appropriate descriptions about
their behavior.

## Language guidelines

We use language keywords instead of BCL types (e.g. `int, string, float` instead of
`Int32, String, Single`, etc.) for both type references as well as method calls (e.g. `int.Parse`
instead of `Int32.Parse`).

### String data type

* Use string interpolation to concatenate short strings, as shown in the following code:

```csharp
string displayName = $"{lastName}, {firstName}";
```

* To append strings in loops, especially when you're working with large amounts of text, use a
StringBuilder object.

### Implicitly typed local variables

* Only use `var` when the type is explicitly named on the right-hand side, typically due to
either `new` or an explicit cast, e.g. `var stream = new FileStrem(...)` not
`var stream = OpenStandardInput()`.
* Similarly, target-typed `new()` can only be used when the type is explicitly named on the
left-hand side, in a variable definition statement or a field definition statement, e.g.
`FileStream stream = new(...)`, but not `stream = new(...);`, where the type was specified on
a previous line.
* Avoid the use of `var` in place of `dynamic`. Use `dynamic` when you want runtime type inference.
* You may use implicit typing to determine the type of the loop variable in for loops.
* Don't use implicit typing to determine the type of the loop variable in foreach loops. In most cases,
the type of elements in the collection isn't immediately obvious

### Unsigned data types

In general, use `int` rather  than unsigned types. The use of `int` is common throughout C#, and
it is easier to interact with other libraries when you use `int`.

### Arrays

Use the concise syntax when you initialize arrays on the declaration line. In the following example,
note that you can't use `var` instead of `string[]`.

```csharp
string[] vowels1 = { "a", "e", "i", "o", "u" };
```

If you use explicit instantiation, you can use `var`:

```csharp
var vowels2 = new string[] { "a", "e", "i", "o", "u" };
```

### Delegates

Use `Func<>` and `Action<>` instead of defining delegate types. In a class, define the delegate method.

```csharp
using static System.Console;

public static Action<string> ActionExample1 = x => WriteLine($"x is: {x}");

public static Action<string, string> ActionExample2 = (x, y) => WriteLine($"x is: {x}, y is {y}");

public static Func<string, int> FuncExample1 = x => Convert.ToInt32(x);

public static Func<int, int, int> FuncExample2 = (x, y) => x + y;

// Call the method using the signature defined by the Func<> or Action<> delegate.

ActionExample1("string for x");

ActionExample2("string for x", "string for y");

WriteLine($"The value is {FuncExample1("1")}");

WriteLine($"The sum is {FuncExample2(1, 2)}");
```

If you create instances of a delegate type, use the concise syntax. In a class, define the delegate type
and a method that has a matching signature.

```csharp
using static System.Console;

public delegate void Del(string message);

public static void DelMethod(string str)
{
	WriteLine("DelMethod argument: {0}", str);
}

// Create an instance of the delegate type and call it.

// The following declaration shows the condensed syntax.

Del exampleDel2 = DelMethod;
exampleDel2("Hey");

// The following declaration uses the full syntax.

Del exampleDel1 = new Del(DelMethod);
exampleDel1("Hey");
```

## Organization and Modifiers

We always specify the visibility, even if it's the default (e.g. `private string _foo` not
`string _foo`). Visibility should be the first modifier (e.g. `public abstract` not `abstract public`)

Modifiers occur in the following order:

```csharp
public protected internal private new abstract virtual override sealed static readonly extern unsafe volatile async
```

Namespace imports should be specified at the top of the file, outside of `namespace` declarations,
and should be sorted alphabetically, with the exception of `System.*` namespaces, which are to be
placed on top of all others.

### Class member ordering:

* Group class members in the following order:
  * Nested classes, enums, delegates and events.
  * Static, const and readonly fields.
  * Fields and properties.
  * Constructors and finalizers.
  * Methods.
* Within each group, elements should be in the following order:
  * Public.
  * Internal.
  * Protected internal.
  * Protected.
  * Private.
* Where possible, group interface implementations together.

## Layout conventions & Whitespace rules

Good layout uses formatting to emphasize the structure of your code and to make the code easier to read.

* We use four spaces of indentation (to tabs).
* Avoid more than one empty line at any time. For example, do not have two blank lines between
member of a type.
* Use the default Code Editor settings (smart indenting, four-character indents, tabs saved as spaces)
* Write only one statement per line.
* A maximum of one assignment per statement
* Column limit: 100
* Write only one declaration per line.
* If continuation lines are not indented automatically, indent them one tab stop (four spaces)
* Add one blank line between method definitions and property definitions.
* Use parentheses to make clauses in an expression apparent.
* Space after `if/for/while` etc., and after commas
* No space after an opening parenthesis or before a closing parenthesis.
* No space between a unary operator and its operand. One space between the operator and each operand
for all other operators.
* For function definitions and calls, if the arguments do not all fit on one line they should be broken
up onto multiple lines, with each subsequent line aligned with the first argument. If there is not enough
room for this, arguments may instead be placed on subsequent lines with a four space indent.
* Avoid spurious free spaces. For example, avoid `if (someVar == 0)...`, where the dots mark the
spurious free spaces. Consider enabling "View White Space (Ctrl+R, Ctrl+W)" if using Visual Studio
to aid detection.
* When using a single-statement if, we follow these conventions:
	1. Never use single-line form (e.g. `if (source == null) return;`)
	2. Using braces is always accepted, and required if any block of an `if/else if/.../else` compound
	statement uses braces of if a single statement body spans multiple lines.
	3. Braces may be omitted only if the body of every block associated with an `if/else if/.../else`
	compound statement is places on a single line.

## General

* We use `nameof(...)` instead of `"..."` whenever possible and relevant.
* When including non-ASCII characters in the source code use Unicode escape sequences (\uXXXX)
instead of literal characters. Literal noon-ASCII characters occasionally get garbled by a tool or editor
* When using labels (for goto), indent the label one less than the current indentation.
* Make all internal and private types static or sealed unless derivation from them is required.
As with any implementation detail, they can be changed if/when derivations is required in the future.
We avoid `this.` unless absolutely necessary.

## Example

```csharp
using System;                                       // `using` goes at the top, outside the namespace

namespace MyNamespace;                              // Namespaces are PascalCase

public interface IMyInterface                       // Interfaces start with 'I'
{
    public int Calculate(float value, float exp);   // Methods are PascalCase. Space after comma.
}

public enum MyEnum                                  // Enumerations are PascalCase
{
	Yes,                                            // Enumerators are PascalCase
	No
}

public class MyClass                                // Classes are PascalCase.
{
	public int Foo = 0;                             // Public member variables are PascalCase.
	public bool NoCounting = false;                 // Field initializers are encouraged.
	private Results _results;                       // Private member variables are _camelCase;
	private const int _bar = 100;                   // const does not affect naming convention.
}

public int CalculateValue(int mulNumber)
{
	int resultValue = Foo * mulNumber;              // Local variables are camelCase.

	if (!NoCounting)                                // No space after unary operator and space after 'if'
	{
		if (resultValue < 0)
		{											// Braces used even when optional
			_results.NumNegativeResults++;
		}
	}
													// Empty line around blocks.
	return resultValue;
}
```