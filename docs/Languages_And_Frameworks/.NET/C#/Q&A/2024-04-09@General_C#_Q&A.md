### General C# Q&A

1. **Basic Syntax and Concepts:** What is the purpose of the `using` statement in C#, and can you provide a simple example of its usage?

**My answer without research:** The purpose of the `using` statement is to safely open a resource ensuring that it will be released when no longer needed, that is, by the end of the scope. An example using the newer syntax would be this:

```csharp
using File f = File.Open("path/to/file");
```

**Assessment and corrections:** Your explanation of the `using` statement is correct.

2. **Object-Oriented Programming (OOP):** In C#, what is the difference between an abstract class and an interface? When would you choose one over the other?

**My answer without research:** An abstract class serves an implementation example, but an interface is only the contract defining what is required. Use interfaces when you need to define what consumers should implement. Use an abstract class when you need to provide implementation examples.

**Assessment and corrections:** Your understanding of abstract classes versus interfaces is fundamentally correct. Abstract classes can provide some implementation, while interfaces cannot. Interfaces define a contract without any implementation. The choice between them depends on whether you need to share code among implementations or just enforce a contract.

3. **LINQ (Language Integrated Query):** Can you explain what LINQ is and provide a simple example of how it can be used to filter a list of integers to only include numbers greater than 10?

**MAWR:** It's a C# language feature that allows developers to write queries in a more idiomatic way directly with C#, without having to create strings with a different syntax.

```csharp
List<int> greaterThanTen = lotsOfIntegers.Where(i => i > 10).ToList();
```

**Assessment and corrections:** Your definition and example of LINQ are accurate. LINQ is indeed a powerful feature of C# that allows writing queries natively in C# against collection of objects.

4. **Asynchronous Programming:** What is the difference between `async` and `await` keywords, and how do they work together in C#?

**MAWR:** `async` marks a method or function as asynchronous indicating that it has an operation that will run on a separate thread, the calling code must handle this appropriately using the `await` keyword. The `await` keyword indicates that the following code will block until it is done, so the code execution should wait here.

**Assessment and corrections:** Your explanation captures the essence of `async` and `await`. However, it's important to note that `async` doesn't necessarily indicate that the operation will run on a separate thread but rather that it's an operation that may complete in the future, allowing the current thread to be used for other work in the meantime. `await` is used to asynchronously wait for the task to complete without blocking the thread.

5. **Error Handling:** How does exception handling in C# work? Can you describe a scenario where you would use a `try-catch` block?

**MAWR:** Your code has to be ready for failures because they will happen. One example would be when making a network request. The device where your application is running could lose internet connection, the remote server could go offline, there could be an error, etc. With a `try-catch` block when and stop a normal error from escaping this context and blowing up somewhere else up the stack, and in the catch block you can handle that exception appropriately by logging, retrying, exiting, terminating the app, etc.

**Assessment and corrections:** Your explanation of error handling using `try-catch` is correct. It's a mechanism to gracefully handle exceptions and errors that may occur during the execution of a program.

6. **Generics:** What are generics in C#, and why are they useful? Can you provide an example of a generic method?

**MAWR:** Generics allow a class to have difference types, so the functions in that class can operate on these different types, for examples, `List<string>` has the method `.Add()` and so does `List<int>`.

**Assessment and corrections:** Generics provide a way to write classes, interfaces, and methods with placeholders for the type of data they store or use, allowing for code reuse and type safety. Here's a simple example of a generic method:

```csharp
public T GenericMethod<T>(T param)
{
    return param;
}
```
