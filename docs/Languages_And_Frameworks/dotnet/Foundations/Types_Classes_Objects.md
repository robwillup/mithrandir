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



