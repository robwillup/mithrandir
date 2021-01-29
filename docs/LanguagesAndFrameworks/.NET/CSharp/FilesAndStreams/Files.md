# Files

One way in C# to get information about a directory is to use the following type:

```csharp
string rootDirectoryPath = new System.IO.DirectoryInfo("path/to/dir").Parent.FullName;
System.Console.WriteLine(rootDirectoryPath);
```

