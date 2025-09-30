# Classes, Structs and Records

## Classes in C#

- They are like blueprints of an object
- Defines data and functionality to work on its data
- Created using class keyword
- Foundation of OOP
- In C#, most code will live inside a class.

```CSharp
public class MyClass
{
    public int a;
    public string b;

    public void MyMethod()
    {
        DoSomething();
    }
}
```

- Fields are class-level variables to hold data
- Methods: contain functionality
- Properties: they provide methods to get and set data in backing fields.
- Events: provides a mechanism to notify listeners of events that happen.

C# 12 introduced Primary Constructors

```CSharp
public class Employee(string name, int age)
{
}
```

## Doing more with Classes and Custom Types

### Namespaces

They are used to avoid name collisions.
- Keep class names separate
- Used throughout .NET
- Organize our own classes in custom namespaces
- Make namespace available through the `using` directive.
- Visual Studio creates by default a root namespace which is the name of your project.
- You can use file-scoped Namespaces since C# 10.

### Static types

Defining a field as static means that it's defined on the class level as opposed to on the instance.
Those fields cannot be accessed through the instances. You need to access them through the class.

### Working with Nulls

```CSharp
Employee employee; // employee is null. We created the variable in the stack but it points to null.
employee = new(); // Now it has an object in the heap.
```

### Garbage collection

This is a process that runs automatically and will remove objects from the heap that no longer have active references.
This process is a feature of the .NET CLR.

GC.Collect() manually calls the gargabe collector manually.

### Records

- Since C# 9
- New reference type
- Can replace class
- Aimed at "just" containing data, but can contain other members
- Comes with additional functionality built-in (generated)

```C#
public record Account;
public record class Account;
public record struct Account;
public record Account(string AccountNumber);
Account newAccount = new("123456")
```

### Why Records?

- For types that contain only data
- Records can block changes to this data
- Used for data that shouldn't change after creation.

Advantages:

- Immutability (for positional records)
- Value-based equality
- Concise

```C#
public record Account(string AccountNumber);
```
