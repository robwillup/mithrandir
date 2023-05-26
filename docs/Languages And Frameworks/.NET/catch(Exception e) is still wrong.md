# CLR Inside Out - Handling Corrupted State Exceptions

`catch (Exception e)` seems innocent and straightforward, but this little
statement can cause a lot of problems when it fails to do what you
expect.

This is a bad example:

```csharp
public void FileSave(String name)
{ 
   try 
      { 
        FileStream fs = new FileStream(name, FileMode.Create); 
       } 
    catch (Exception)
      { 
        throw new System.IO.IOException("File Open Error!"); 
      } 
}
```

The error is this code is common: it is simple to write code to catch all
exceptions than it is to catch exactly the exceptions that could be
raised by code executing in your try block. But by catching the base of
the exception hierarchy, you will end up swallowing an exception of any
type and converting it to IOException.

Exception handling is one of those areas about which most people have a
working knowledge rather than an intimate understanding.

## What exactly are exceptions?

An exception is the signal raised when a condition is detected that was
not expected in the normal execution of a program thread. Many agents
can detect incorrect conditions and raise exceptions. Program code (or
the library code it uses) can throw types derived from System.Exception,
the CLR execution engine can raise exceptions, and unmanaged code can
raise exceptions as well. Exceptions raised on a thread of execution
follow the thread through native and managed code, across AppDomains,
and, if not handled by the program, are treated as unhandled exceptions
by the operating system.

An exception indicates that something bad has happened. While every
managed exception has a type (such as System.ArgumentException or
System.ArithmeticException), the type is only meaningful in the context
in which the exception is raised. A program can handle an exception if it
understands the conditions that caused the exception to occur. But if the
program doesn't handle the exception, it could indicate any number of bad
things. And once the exception has left the program, it only has one very
general meaning: something bad has happened.

When Windows sees that a program doesn't handle an exception, it tries to
protect the program's persisted data (files on disk, registry settings,
and so on) by terminating the process. Even if the exception originally
indicated some benign, unexpected program state (such as failing to pop
from an empty stack) it appears to be a serious program when Windows sees
it because the operating system doesn't have the context to properly
interpret the exception. A single thread in one AppDomain can bring down
an entire CLR instance by not handling an exception.

If exceptions are so dangerous, why are they so popular? Like
motorcycles and chain saws, the raw power of exceptions makes them very
useful. Normal dataflow on a program thread goes from function to
function calls and returns. Every call to a function creates a frame
of execution on the stack; every return destroys that frame. Aside, from
altering global state, the only flow of data in a program is achieved
by passing data between contiguous frames as function parameters or
return values. In the absence of exception handling, every caller needs
to check for the success of the function that it called (or just assume
everything is always OK).

Most Win32 APIs return a non-zero value to indicate failure because
Windows doesn't use exception handling. **The programmer has to wrap
every function call with code that checks the return value of the called
function.** Fr example, this code from MSDN documentation on how to list
files in a directory explicitly checks each call for success. The call to
`FindNextFile(...)` is wrapped in a check to see if the return is
non-zero. If the call is not successful then a separate function call--
`GetLastError()` -- provides details of the exceptional condition.
Note that every call must be checked for success on the next frame
because return values are necessarily limited to the local function scope:

```csharp
// FindNextFile requires checking for success of each call
while(FindNextFile(hFind, &ffd) != 0);
dwError = GetLastError();
if (dwError != ERROR_NO_MORE_FILES)
{
  ErrorHandler(TEXT("FindFirstFile"));
}
FindClose(hFind);
return dwError;
```
