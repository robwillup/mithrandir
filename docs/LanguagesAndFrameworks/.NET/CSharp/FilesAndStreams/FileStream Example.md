# Example of Reading and Writing to a File using the FileStream

```csharp
public void Process()
{
    using (var inputFileStream = new FileStream("file/path", FileMode.Open))
    using (var inputStreamReader = new StreamReader(inputFileStream))
    using (var outputFileStream = new FileStream("new/file/path", FileMode.Create))
    using (var ouputStreamWriter = new StreamWriter(outputFileStream))
    {
        while (!inputStreamReader.EndOfStream)
        {
            string line = inputStreamReader.ReadLine();
            string processedLine = line.ToUpperInvariant();
            outputStreamWriter.WriteLine(processedLine);
        }
    }
}
```

### Simplified

```csharp
public void Process()
{    
    using var inputStreamReader = new StreamReader("file/path");   
    using var ouputStreamWriter = new StreamWriter("new/file/path");
    while (!inputStreamReader.EndOfStream)
    {
        string line = inputStreamReader.ReadLine();
        string processedLine = line.ToUpperInvariant();
        outputStreamWriter.WriteLine(processedLine);
    }+-
}
```