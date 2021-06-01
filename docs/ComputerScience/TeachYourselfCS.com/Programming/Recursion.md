# Recursion

Example of recursion to find the factorial of a number:

```csharp
static int Factorial(int n)
{     
    if (n > 1)         
        return n * Factorial(--n);
    return 1;
}
```
